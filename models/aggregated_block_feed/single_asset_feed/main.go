package single_asset_feed

import (
	"encoding/hex"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/artifacts/redstone"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg/priceFetcher"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/aggregated_block_feed/base_price_feed"
	"github.com/ethereum/go-ethereum/common"
)

type SingleAssetFeed struct {
	*base_price_feed.BasePriceFeed
}

func NewSingleAsset(token, oracle string, pfType string, discoveredAt int64, client core.ClientI, repo ds.RepositoryI, pfVersion schemas.PFVersion, underlyingFeeds []string) *SingleAssetFeed {
	adapter := base_price_feed.NewBasePriceFeed(token, oracle, pfType, discoveredAt, client, repo, pfVersion)
	return NewSingleAssetFromAdapter(adapter.SyncAdapter)
}

func NewSingleAssetFromAdapter(adapter *ds.SyncAdapter) *SingleAssetFeed {
	return &SingleAssetFeed{
		BasePriceFeed: base_price_feed.NewBasePriceFeedFromAdapter(adapter),
	}
}

func (feed SingleAssetFeed) GetRedstonePF() *core.RedStonePF {
	if len(feed.DetailsDS.Underlyings) == 0 {
		return nil
	}
	return feed.DetailsDS.Info[feed.DetailsDS.Underlyings[0]]
}

func (mdl *SingleAssetFeed) GetCalls(blockNum int64) (calls []multicall.Multicall2Call, isQueryable bool) {
	updateABI := core.GetAbi("UpdatePriceFeed")
	for _, entry := range mdl.DetailsDS.Underlyings {
		contract, err := redstone.NewRedstone(common.HexToAddress(entry), mdl.Client)
		log.CheckFatal(err)
		var tokenDetails *core.RedStonePF
		if _, ok := mdl.DetailsDS.Info[entry]; ok {
			tokenDetails = mdl.DetailsDS.Info[entry]
		} else if _, err := contract.DataFeedId(nil); err == nil {
			feedToken, signThreshold, dataId := priceFetcher.RedstoneDetails(common.HexToAddress(entry), mdl.Client)
			//
			tokenDetails = &core.RedStonePF{
				Type:             15,
				DataServiceId:    "redstone-primary-prod",
				DataId:           dataId,
				SignersThreshold: signThreshold,
				UnderlyingToken:  feedToken,
			}
			mdl.DetailsDS.Info[entry] = tokenDetails
		}
		if tokenDetails != nil {
			pod := mdl.Repo.GetRedStonemgr().GetPodSignWithRedstoneToken(int64(mdl.Repo.SetAndGetBlock(blockNum).Timestamp), *tokenDetails)
			update, err := updateABI.Pack("updatePrice", pod.CallData)
			log.CheckFatal(err)
			calls = append(calls, multicall.Multicall2Call{
				Target:   common.HexToAddress(entry),
				CallData: update,
			})
		}
	}
	b, err := hex.DecodeString("feaf968c")
	log.CheckFatal(err)
	calls = append(calls, multicall.Multicall2Call{
		Target:   common.HexToAddress(mdl.Address),
		CallData: b,
	})
	return calls, true
}

// same as query price feed
// func (*YearnPriceFeed) GetCalls(blockNum int64) (calls []multicall.Multicall2Call, isQueryable bool) {

func (mdl *SingleAssetFeed) ProcessResult(blockNum int64, results []multicall.Multicall2Result, force ...bool) *schemas.PriceFeed {
	result := results[len(results)-1]
	if !result.Success {
		log.Warnf("Can't get latestRounData for YearnModule in AQFWrapper for %s(%s) at %d",
			mdl.GetDetailsByKey("pfType"), mdl.GetAddress(), blockNum)
		return nil
		//
	}
	isPriceInUSD := mdl.GetVersion().IsPriceInUSD()
	return base_price_feed.ParseQueryRoundData(result.ReturnData, isPriceInUSD, mdl.GetAddress(), blockNum)
}
