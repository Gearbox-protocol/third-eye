package base_price_feed

import (
	"sync"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/core/types"
)

type BasePriceFeed struct {
	*ds.SyncAdapter
	mu *sync.Mutex
	// DetailsDS DetailsDS
}

func (feed BasePriceFeed) GetRedstonePF() (ans []*core.RedStonePF) {
	if len(feed.DetailsDS.Underlyings) == 0 {
		return nil
	}
	for _, d := range feed.DetailsDS.Info {
		ans = append(ans, d)
	}
	return ans
}

// single querypricefeed can be valid for multiple tokens so we have to maintain tokens within the details
// details->token is token map to start and end block
func NewBasePriceFeed(token, oracle string, pfType string, discoveredAt int64, client core.ClientI, repo ds.RepositoryI, version core.VersionType) *BasePriceFeed {
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
				"pfType": pfType,
			},
			LastSync: discoveredAt,
			V:        version,
		},
		Repo: repo,
	}

	mdl := NewBasePriceFeedFromAdapter(
		syncAdapter,
	)
	return mdl
}

func (pf *BasePriceFeed) GetPFType() string {
	return pf.GetDetailsByKey("pfType")
}

func NewBasePriceFeedFromAdapter(adapter *ds.SyncAdapter) *BasePriceFeed {
	obj := &BasePriceFeed{
		SyncAdapter: adapter,
		mu:          &sync.Mutex{},
	}
	obj.DataProcessType = ds.ViaQuery
	// obj.DetailsDS.Load(obj.Details, obj.GetVersion())
	return obj
}

func (mdl *BasePriceFeed) OnLog(txLog types.Log) {

}

// func (mdl *BasePriceFeed) GetTokens() map[string]map[schemas.PFVersion][]int64 {
// 	mdl.mu.Lock()
// 	defer mdl.mu.Unlock()
// 	return mdl.DetailsDS.Tokens
// }

///////////////////////
// details for token
///////////////////////

// func (mdl *BasePriceFeed) AddToken(token string, discoveredAt int64, pfVersion schemas.PFVersion) {
// 	tokenDetails := mdl.DetailsDS.Tokens[token]
// 	if tokenDetails == nil {
// 		mdl.DetailsDS.Tokens[token] = map[schemas.PFVersion][]int64{pfVersion: {discoveredAt}}
// 		tokenDetails = mdl.DetailsDS.Tokens[token]
// 	}
// 	blockNums := tokenDetails[pfVersion]
// 	if len(blockNums) == 1 {
// 		log.Debugf("Token/Feed(%s/%s) previously added at %d, again added at %d", token, mdl.Address, blockNums[0], discoveredAt)
// 		return
// 	} else if len(blockNums) == 2 {
// 		mdl.DetailsDS.Logs = append(mdl.DetailsDS.Logs, []interface{}{token, blockNums, pfVersion}) // token blockNums pfVersion
// 	}
// 	tokenDetails[pfVersion] = []int64{discoveredAt}
// }

// // sync till < endBlock number
// func (mdl *BasePriceFeed) DisableToken(token string, disabledAt int64, pfVersion schemas.PFVersion) {
// 	tokenDetails := mdl.DetailsDS.Tokens[token]
// 	if tokenDetails == nil || len(tokenDetails[pfVersion]) != 1 {
// 		log.Fatalf("%s's enable block number for pfVersion %d is malformed: %v, trying with new block: %d", token, pfVersion, tokenDetails, disabledAt)
// 	}
// 	tokenDetails[pfVersion] = append(tokenDetails[pfVersion], disabledAt)
// }

// // read method
// func (mdl *BasePriceFeed) TokensValidAtBlock(blockNum int64) (ans []schemas.TokenAndMergedPFVersion) {
// 	for token, details := range mdl.DetailsDS.Tokens {
// 		mpfVersion := mergePFVersionAt(blockNum, details)
// 		if mpfVersion != 0 {
// 			ans = append(ans, schemas.TokenAndMergedPFVersion{Token: token, MergedPFVersion: mpfVersion, Feed: mdl.GetAddress()})
// 		}
// 	}
// 	return ans
// 	//
// }

// func mergePFVersionAt(blockNum int64, details map[schemas.PFVersion][]int64) schemas.MergedPFVersion {
// 	var pfVersion schemas.MergedPFVersion = 0
// 	for version, blockNums := range details {
// 		// log.Info(version, blockNums, blockNum)
// 		if blockNums[0]+1 <= blockNum && (len(blockNums) == 1 || blockNum < blockNums[1]) { // 1 is added as price is already added at blockNum
// 			pfVersion = pfVersion | schemas.MergedPFVersion(version)
// 		}
// 	}
// 	return pfVersion
// }

func (mdl *BasePriceFeed) AfterSyncHook(b int64) {
	if log.GetBaseNet(core.GetChainId(mdl.Client)) == "MAINNET" {
		var v1CloseBlock int64 = 18577104 // v1 all accounts closed at
		if b >= v1CloseBlock {
			if mdl.GetVersion() == core.NewVersion(1) {
				mdl.SetBlockToDisableOn(v1CloseBlock)
			}
		}
		var v2CloseBlock int64 = 19752044 // v2 all accounts closed at
		if b >= v2CloseBlock {
			if mdl.GetVersion() == core.NewVersion(2) {
				mdl.SetBlockToDisableOn(v2CloseBlock)
			}
		}
	}
	// mdl.Details = mdl.DetailsDS.Save()
	mdl.SyncAdapter.AfterSyncHook(b)
}
