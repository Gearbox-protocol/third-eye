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
	DecimalDivider    *big.Int
	mu                *sync.Mutex
	contractETH       *yearnPriceFeed.YearnPriceFeed // for yearn manual price calculation
	YVaultContract    *yVault.YVault                 //f or yearn manual price calculation
	PriceFeedContract *priceFeed.PriceFeed
}

// single querypricefeed can be valid for multiple tokens so we have to maintain tokens within the details
// details->token is token map to start and end block
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
			Details:  map[string]interface{}{"token": map[string]interface{}{token: []int64{discoveredAt}}, "pfType": pfType},
			LastSync: discoveredAt - 1,
			V:        version,
		},
		Repo: repo,
	}
	mdl := NewQueryPriceFeedFromAdapter(
		syncAdapter,
	)
	return mdl
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

func (mdl *QueryPriceFeed) calculateYearnPFInternally(blockNum int64) (*schemas.PriceFeed, error) {
	if mdl.YVaultContract == nil || mdl.PriceFeedContract == nil || mdl.DecimalDivider == nil {
		if err := mdl.setContracts(blockNum); err != nil {
			return nil, err
		}
	}
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
	}
	pricePerShare, err := mdl.YVaultContract.PricePerShare(opts)
	if err != nil {
		return nil, err
	}
	roundData, err := mdl.PriceFeedContract.LatestRoundData(opts)
	if err != nil {
		return nil, err
	}

	lowerBound, err := mdl.contractETH.LowerBound(opts)
	if err != nil {
		return nil, err
	}
	uppwerBound, err := mdl.contractETH.UpperBound(opts)
	if err != nil {
		return nil, err
	}
	if !(pricePerShare.Cmp(lowerBound) >= 0 && pricePerShare.Cmp(uppwerBound) <= 0) {
		if !mdl.isNotified() {
			mdl.setNotified(true)
			mdl.Repo.RecentEventMsg(blockNum, "PricePerShare(%s) %d is not btw lower limit(%d) and upper limit(%d).", mdl.Address, pricePerShare, lowerBound, uppwerBound)
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
	}, nil
}

func (mdl *QueryPriceFeed) setContracts(blockNum int64) error {
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
	}
	// set the price feed contract
	if priceFeedAddr, err := mdl.contractETH.PriceFeed(opts); err != nil {
		return err
	} else if priceFeedContract, err := priceFeed.NewPriceFeed(priceFeedAddr, mdl.Client); err != nil {
		return err
	} else {
		mdl.PriceFeedContract = priceFeedContract
	}

	// set the yvault contract
	if yVaultAddr, err := mdl.contractETH.YVault(opts); err != nil {
		return err
	} else if yVaultContract, err := yVault.NewYVault(yVaultAddr, mdl.Client); err != nil {
		return err
	} else {
		mdl.YVaultContract = yVaultContract
	}

	// set the decimals
	if decimals, err := mdl.YVaultContract.Decimals(opts); err != nil {
		return err
	} else {
		mdl.DecimalDivider = utils.GetExpInt(int8(decimals))
	}
	//
	return nil
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
		case map[string]interface{}:
			obj, _ = mdl.Details["token"].(map[string]interface{})
			ints := ConvertToListOfInt64(obj[token])
			// token is already in enabled state, we are trying to add again
			if obj[token] != nil && len(ints) == 1 {
				log.Verbosef("Token/Feed(%s/%s) previously added at %d, again added at %d", token, mdl.Address, ints[0], discoveredAt)
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
	case map[string]interface{}:
		obj = mdl.Details["token"].(map[string]interface{})
		ints := ConvertToListOfInt64(obj[token])
		if len(ints) != 1 {
			log.Fatalf("%s's enable block number for pricefeed is malformed: %v", token, ints)
		}
		ints = append(ints, disabledAt)
		obj[token] = ints
	}
	mdl.Details["token"] = obj
}

// read method
func (mdl *QueryPriceFeed) TokensValidAtBlock(blockNum int64) []string {
	switch mdl.Details["token"].(type) {
	case map[string]interface{}:
		tokens := []string{}
		obj := mdl.Details["token"].(map[string]interface{})
		for token, info := range obj {
			ints := ConvertToListOfInt64(info)
			// when token is added to the feed, price at discoveredAt is added for that token
			// so we can ignore that token is valid at discoveredAt, hence added one to discoveredAt
			if ints[0]+1 <= blockNum && (len(ints) == 1 || blockNum < ints[1]) {
				tokens = append(tokens, token)
			}
		}
		return tokens
	}
	return nil
}

func ConvertToListOfInt64(list interface{}) (parsedInts []int64) {
	switch ints := list.(type) {
	case []interface{}:
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
		parsedInts = ints
	}
	return
}
