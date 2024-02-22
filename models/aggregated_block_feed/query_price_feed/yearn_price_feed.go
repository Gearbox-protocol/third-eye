package query_price_feed

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
	DetailsDS DetailsDS
}

// single querypricefeed can be valid for multiple tokens so we have to maintain tokens within the details
// details->token is token map to start and end block
func NewQueryPriceFeed(token, oracle string, pfType string, discoveredAt int64, client core.ClientI, repo ds.RepositoryI, pfVersion schemas.PFVersion) *QueryPriceFeed {
	syncAdapter := &ds.SyncAdapter{
		SyncAdapterSchema: &schemas.SyncAdapterSchema{
			Contract: &schemas.Contract{
				Address:      oracle,
				DiscoveredAt: discoveredAt,
				FirstLogAt:   discoveredAt,
				ContractName: ds.QueryPriceFeed,
				Client:       client,
			},
			Details: map[string]interface{}{
				"tokens": map[string]map[schemas.PFVersion][]int64{token: {pfVersion: {discoveredAt}}},
				"pfType": pfType,
			},
			LastSync: discoveredAt,
			V:        pfVersion.ToVersion(),
		},
		Repo: repo,
	}

	mdl := NewQueryPriceFeedFromAdapter(
		syncAdapter,
	)
	return mdl
}

func (pf *QueryPriceFeed) GetPFType() string {
	return pf.GetDetailsByKey("pfType")
}

func NewQueryPriceFeedFromAdapter(adapter *ds.SyncAdapter) *QueryPriceFeed {
	obj := &QueryPriceFeed{
		SyncAdapter: adapter,
		mu:          &sync.Mutex{},
		yearnPFInternal: yearnPFInternal{
			mainPFAddress: common.HexToAddress(adapter.Address),
			version:       adapter.GetVersion(),
		}, // main price feed
	}
	obj.DataProcessType = ds.ViaQuery
	obj.DetailsDS.Load(obj.Details, obj.GetVersion())
	return obj
}

func (mdl *QueryPriceFeed) OnLog(txLog types.Log) {

}

// only used in v1,v2
func (mdl *QueryPriceFeed) CalculateYearnPFInternally(blockNum int64) (*schemas.PriceFeed, error) {
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

func (mdl *QueryPriceFeed) AddToken(token string, discoveredAt int64, pfVersion schemas.PFVersion) {
	tokenDetails := mdl.DetailsDS.Tokens[token]
	if tokenDetails == nil {
		mdl.DetailsDS.Tokens[token] = map[schemas.PFVersion][]int64{pfVersion: {discoveredAt}}
		tokenDetails = mdl.DetailsDS.Tokens[token]
	}
	blockNums := tokenDetails[pfVersion]
	if len(blockNums) == 1 {
		log.Debugf("Token/Feed(%s/%s) previously added at %d, again added at %d", token, mdl.Address, blockNums[0], discoveredAt)
		return
	} else if len(blockNums) == 2 {
		mdl.DetailsDS.Logs = append(mdl.DetailsDS.Logs, []interface{}{token, blockNums, pfVersion}) // token blockNums pfVersion
	}
	tokenDetails[pfVersion] = []int64{discoveredAt}
}

// sync till < endBlock number
func (mdl *QueryPriceFeed) DisableToken(token string, disabledAt int64, pfVersion schemas.PFVersion) {
	tokenDetails := mdl.DetailsDS.Tokens[token]
	if tokenDetails == nil || len(tokenDetails[pfVersion]) != 1 {
		log.Fatalf("%s's enable block number for pfVersion %d is malformed: %v", token, pfVersion, tokenDetails)
	}
	tokenDetails[pfVersion] = append(tokenDetails[pfVersion], disabledAt)
}

// read method
func (mdl *QueryPriceFeed) TokensValidAtBlock(blockNum int64) (ans []schemas.TokenAndMergedPFVersion) {
	for token, details := range mdl.DetailsDS.Tokens {
		mpfVersion := mergePFVersionAt(blockNum, details)
		if mpfVersion != 0 {
			ans = append(ans, schemas.TokenAndMergedPFVersion{Token: token, MergedPFVersion: mpfVersion, Feed: mdl.GetAddress()})
		}
	}
	return ans
	//
}

func mergePFVersionAt(blockNum int64, details map[schemas.PFVersion][]int64) schemas.MergedPFVersion {
	var pfVersion schemas.MergedPFVersion = 0
	for version, blockNums := range details {
		// log.Info(version, blockNums, blockNum)
		if blockNums[0]+1 <= blockNum && (len(blockNums) == 1 || blockNum < blockNums[1]) { // 1 is added as price is already added at blockNum
			pfVersion = pfVersion | schemas.MergedPFVersion(version)
		}
	}
	return pfVersion
}

func (mdl *QueryPriceFeed) AfterSyncHook(b int64) {
	mdl.Details = mdl.DetailsDS.Save()
	mdl.SyncAdapter.AfterSyncHook(b)
}
