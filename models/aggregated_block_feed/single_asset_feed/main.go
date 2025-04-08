package single_asset_feed

import (
	"fmt"
	"time"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/aggregated_block_feed/base_price_feed"
)

type SingleAssetFeed struct {
	*base_price_feed.BasePriceFeed
}

func NewSingleAsset(token, oracle string, pfType string, discoveredAt int64, client core.ClientI, repo ds.RepositoryI, version core.VersionType, underlyingFeeds []string) *SingleAssetFeed {
	adapter := base_price_feed.NewBasePriceFeed(token, oracle, pfType, discoveredAt, client, repo, version)
	return NewSingleAssetFromAdapter(adapter.SyncAdapter)
}

func NewSingleAssetFromAdapter(adapter *ds.SyncAdapter) *SingleAssetFeed {
	return &SingleAssetFeed{
		BasePriceFeed: base_price_feed.NewBasePriceFeedFromAdapter(adapter),
	}
}

func (feed SingleAssetFeed) GetCalls(blockNum int64) (calls []multicall.Multicall2Call, isQueryable bool) {
	return feed.GetUnderlyingCalls(blockNum)
}

var counter = log.SendMsgIfCountMoreThan(24*time.Hour, 10)

// same as query price feed
// func (*YearnPriceFeed) GetCalls(blockNum int64) (calls []multicall.Multicall2Call, isQueryable bool) {

func (mdl *SingleAssetFeed) ProcessResult(blockNum int64, results []multicall.Multicall2Result, force ...bool) *schemas.PriceFeed {
	result := results[len(results)-1]
	if !result.Success {
		counter(mdl.GetAddress(), fmt.Sprintf("Can't get latestRounData for YearnModule in AQFWrapper for %s(%s) at %d",
			mdl.GetDetailsByKey("pfType"), mdl.GetAddress(), blockNum))
		return nil
		//
	}
	isPriceInUSD := mdl.GetVersion().IsPriceInUSD()
	return base_price_feed.ParseQueryRoundData(result.ReturnData, isPriceInUSD, mdl.GetAddress(), blockNum)
}
