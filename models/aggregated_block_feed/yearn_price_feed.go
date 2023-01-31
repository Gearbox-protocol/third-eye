package aggregated_block_feed

import (
	"sync"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type QueryPriceFeed struct {
	*ds.SyncAdapter
	mu *sync.Mutex
	yearnPFInternal
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
		SyncAdapter:     adapter,
		mu:              &sync.Mutex{},
		yearnPFInternal: yearnPFInternal{mainPFAddress: common.HexToAddress(adapter.Address)}, // main price feed
	}
	obj.OnlyQuery = true
	return obj
}

func (mdl *QueryPriceFeed) OnLog(txLog types.Log) {

}
func (mdl *QueryPriceFeed) calculateYearnPFInternally(blockNum int64) (*schemas.PriceFeed, error) {
	return mdl.yearnPFInternal.calculatePrice(blockNum, mdl.Client, mdl.GetVersion())
}

func (mdl *QueryPriceFeed) GetTokenAddr() string {
	mdl.mu.Lock()
	defer mdl.mu.Unlock()
	tokenAddr, ok := mdl.Details["token"].(string)
	if !ok {
		log.Fatalf("Failing in asserting to string: %s", mdl.Details["token"])
	}
	return tokenAddr
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

// sync till < endBlock number
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
