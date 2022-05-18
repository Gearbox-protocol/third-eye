package aggregated_block_feed

import (
	"math/big"
	"sync"

	"github.com/Gearbox-protocol/sdk-go/artifacts/priceFeed"
	"github.com/Gearbox-protocol/sdk-go/artifacts/yVault"
	"github.com/Gearbox-protocol/sdk-go/artifacts/yearnPriceFeed"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type QueryPriceFeed struct {
	*ds.SyncAdapter
	contractETH       *yearnPriceFeed.YearnPriceFeed
	YVaultContract    *yVault.YVault
	PriceFeedContract *priceFeed.PriceFeed
	DecimalDivider    *big.Int
	mu                *sync.Mutex
}

func NewQueryPriceFeed(token, oracle string, pfType string, discoveredAt int64, client core.ClientI, repo ds.RepositoryI, version int16) *QueryPriceFeed {
	syncAdapter := &ds.SyncAdapter{
		SyncAdapterSchema: &schemas.SyncAdapterSchema{
			Contract: &schemas.Contract{
				Address:      oracle,
				DiscoveredAt: discoveredAt,
				FirstLogAt:   discoveredAt,
				ContractName: ds.QueryPriceFeed,
				Client:       client,
			},
			Details:  map[string]interface{}{"token": token, "pfType": pfType},
			LastSync: discoveredAt - 1,
			V:        version,
		},
		Repo: repo,
	}
	return NewQueryPriceFeedFromAdapter(
		syncAdapter,
	)
}

func NewQueryPriceFeedFromAdapter(adapter *ds.SyncAdapter) *QueryPriceFeed {
	obj := &QueryPriceFeed{
		SyncAdapter: adapter,
		mu:          &sync.Mutex{},
	}
	switch adapter.GetDetailsByKey("pfType") {
	case ds.YearnPF:
		var err error
		obj.contractETH, err = yearnPriceFeed.NewYearnPriceFeed(common.HexToAddress(adapter.Address), adapter.Client)
		if err != nil {
			log.Fatal(err)
		}
		// case ds.CurvePF:
		// log.Info("No contract set for curve pf to use internally")
	}
	obj.OnlyQuery = true
	return obj
}

func (mdl *QueryPriceFeed) OnLog(txLog types.Log) {

}
func (mdl *QueryPriceFeed) isNotified() bool {
	mdl.mu.Lock()
	defer mdl.mu.Unlock()
	if mdl.Details == nil || mdl.Details["notified"] == nil {
		return false
	}
	value, ok := mdl.Details["notified"].(bool)
	if !ok {
		log.Fatal("Notified not parsed")
	}
	return value
}

func (mdl *QueryPriceFeed) setNotified(notified bool) {
	mdl.mu.Lock()
	defer mdl.mu.Unlock()
	mdl.Details["notified"] = notified
}

func (mdl *QueryPriceFeed) GetTokenAddr() string {
	mdl.mu.Lock()
	defer mdl.mu.Unlock()
	tokenAddr, ok := mdl.Details["token"].(string)
	if !ok {
		log.Fatal("Failing in asserting to string: %s", mdl.Details["token"])
	}
	return tokenAddr
}

func (mdl *QueryPriceFeed) calculateYearnPFInternally(blockNum int64) *schemas.PriceFeed {
	if mdl.YVaultContract == nil || mdl.PriceFeedContract == nil || mdl.DecimalDivider == nil {
		mdl.setContracts(blockNum)
	}
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
	}

	roundData, err := mdl.PriceFeedContract.LatestRoundData(opts)
	log.CheckFatal(err)
	pricePerShare, err := mdl.YVaultContract.PricePerShare(opts)
	log.CheckFatal(err)

	lowerBound, err := mdl.contractETH.LowerBound(opts)
	log.CheckFatal(err)
	uppwerBound, err := mdl.contractETH.UpperBound(opts)
	log.CheckFatal(err)
	if !(pricePerShare.Cmp(lowerBound) >= 0 && pricePerShare.Cmp(uppwerBound) <= 0) {
		if !mdl.isNotified() {
			mdl.setNotified(true)
			log.Warnf("PricePerShare(%d) is not btw lower limit(%d) and upper limit(%d).", pricePerShare, lowerBound, uppwerBound)
		}
	} else {
		mdl.Details["notified"] = false
	}
	// decimalDivider https://github.com/Gearbox-protocol/contracts-v2/blob/main/contracts/oracles/curve/AbstractCurveLPPriceFeed.sol#L36
	// it is 18 for curve as the lp is denotated in eth.
	// for yearn it is based on the vault. https://github.com/Gearbox-protocol/contracts-v2/blob/main/contracts/oracles/yearn/YearnPriceFeed.sol#L54
	newAnswer := new(big.Int).Quo(
		new(big.Int).Mul(pricePerShare, roundData.Answer),
		mdl.DecimalDivider,
	)
	/// decimals is based on https://github.com/Gearbox-protocol/contracts-v2/blob/main/contracts/oracles/curve/AbstractCurveLPPriceFeed.sol#L22
	// if the feed is usd it is 8 else 18.
	//
	isPriceInUSD := mdl.GetVersion() > 1
	var decimals int8 = 18 // for eth
	if isPriceInUSD {
		decimals = 8 // for usd
	}
	return &schemas.PriceFeed{
		RoundId:      roundData.RoundId.Int64(),
		PriceBI:      (*core.BigInt)(newAnswer),
		Price:        utils.GetFloat64Decimal(newAnswer, decimals),
		IsPriceInUSD: isPriceInUSD,
	}
}

func (mdl *QueryPriceFeed) setContracts(blockNum int64) {
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
	}
	// set the price feed contract
	priceFeedAddr, err := mdl.contractETH.PriceFeed(opts)
	log.CheckFatal(err)
	priceFeedContract, err := priceFeed.NewPriceFeed(priceFeedAddr, mdl.Client)
	log.CheckFatal(err)
	mdl.PriceFeedContract = priceFeedContract

	// set the yvault contract
	yVaultAddr, err := mdl.contractETH.YVault(opts)
	log.CheckFatal(err)
	yVaultContract, err := yVault.NewYVault(yVaultAddr, mdl.Client)
	log.CheckFatal(err)
	mdl.YVaultContract = yVaultContract

	// set the decimals
	decimals, err := yVaultContract.Decimals(opts)
	log.CheckFatal(err)
	mdl.DecimalDivider = utils.GetExpInt(int8(decimals))
}

///////////////////////
// details for token
///////////////////////

func (mdl *QueryPriceFeed) AddToken(token string, discoveredAt int64) {
	if mdl.Details == nil {
		mdl.Details = core.Json{}
	}
	if mdl.Details["token"] != nil {
		obj := map[string]interface{}{}
		switch mdl.Details["token"].(type) {
		case string:
			prevToken := mdl.Details["token"].(string)
			obj[prevToken] = []int64{mdl.DiscoveredAt}
			if prevToken == token {
				log.Warnf("Token/Feed(%s/%s) previously added at %d, again added at %d", token, mdl.Address, mdl.DiscoveredAt, discoveredAt)
				return
			}
		case map[string]interface{}:
			obj, _ = mdl.Details["token"].(map[string]interface{})
			ints := ConvertToListOfInt64(obj[token])
			// token is already in enabled state, we are trying to add again
			if obj[token] != nil && len(ints) == 1 {
				log.Warnf("Token/Feed(%s/%s) previously added at %d, again added at %d", token, mdl.Address, ints[0], discoveredAt)
				return
				// token is disabled so reenable and add to logs
			} else if len(ints) == 2 {
				mdl.Details["logs"] = append(parseLogArray(mdl.Details["logs"]), []interface{}{token, ints})
			}
		}
		obj[token] = []int64{discoveredAt}
		mdl.Details["token"] = obj
	} else {
		log.Fatal("Can't reach this part in the yearn price feed")
	}
}

func parseLogArray(logs interface{}) (parsedLogs [][]interface{}) {
	if logs != nil {
		l, ok := logs.([]interface{})
		if !ok {
			log.Fatal("failed in converting to log array", logs)
		}
		for _, ele := range l {
			obj, ok := ele.([]interface{})
			if !ok {
				log.Fatal("failed in converting to log array element", ele)
			}
			parsedEle := []interface{}{obj[0].(string), ConvertToListOfInt64(obj[1])}
			parsedLogs = append(parsedLogs, parsedEle)
		}
	}
	return
}

func (mdl *QueryPriceFeed) DisableToken(token string, disabledAt int64) {
	obj := map[string]interface{}{}
	switch mdl.Details["token"].(type) {
	case string:
		obj[token] = []int64{mdl.DiscoveredAt, disabledAt}
	case map[string]interface{}:
		obj = mdl.Details["token"].(map[string]interface{})
		ints := ConvertToListOfInt64(obj[token])
		if len(ints) != 1 {
			log.Fatal("%s's enable block number for pricefeed is malformed: %v", ints)
		}
		ints = append(ints, disabledAt)
		obj[token] = ints
	}
	mdl.Details["token"] = obj
}

func (mdl *QueryPriceFeed) TokensValidAtBlock(blockNum int64) []string {
	switch mdl.Details["token"].(type) {
	case string:
		tokens := []string{}
		if mdl.DiscoveredAt <= blockNum {
			tokens = append(tokens, mdl.Details["token"].(string))
		}
		return tokens
	case map[string]interface{}:
		tokens := []string{}
		obj := mdl.Details["token"].(map[string]interface{})
		for token, info := range obj {
			ints := ConvertToListOfInt64(info)
			if ints[0] <= blockNum && (len(ints) == 1 || blockNum < ints[1]) {
				tokens = append(tokens, token)
			}
		}
		return tokens
	}
	return nil
}

func ConvertToListOfInt64(list interface{}) (parsedInts []int64) {
	switch list.(type) {
	case []interface{}:
		ints, ok := list.([]interface{})
		if !ok {
			panic("parsing list of int failed")
		}
		for _, int_ := range ints {
			var parsedInt int64
			switch int_.(type) {
			case int64:
				parsedInt = int_.(int64)
			case float64:
				parsedFloat := int_.(float64)
				parsedInt = int64(parsedFloat)
			default:
				log.Fatalf("QueryPriceFeed token start/end block_num not in int format %v", int_)
			}
			parsedInts = append(parsedInts, parsedInt)
		}
	case []int64:
		ints, ok := list.([]int64)
		if !ok {
			panic("parsing accounts list for token transfer failed")
		}
		parsedInts = ints
	}
	return
}
