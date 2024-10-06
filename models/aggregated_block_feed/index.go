package aggregated_block_feed

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/aggregated_block_feed/composite_redstone_price_feed"
	"github.com/Gearbox-protocol/third-eye/models/aggregated_block_feed/curve_price_feed"
	"github.com/Gearbox-protocol/third-eye/models/aggregated_block_feed/redstone_price_feed"
	"github.com/Gearbox-protocol/third-eye/models/aggregated_block_feed/single_asset_feed"
	"github.com/Gearbox-protocol/third-eye/models/aggregated_block_feed/yearn_price_feed"
)

func NewQueryPriceFeed(token, oracle string, pfType string, discoveredAt int64, client core.ClientI, repo ds.RepositoryI, version core.VersionType) ds.QueryPriceFeedI {
	switch pfType {
	case ds.RedStonePF:
		return redstone_price_feed.NewRedstonePriceFeed(token, oracle, pfType, discoveredAt, client, repo, version)
	case ds.CurvePF:
		return curve_price_feed.NewCurvePriceFeed(token, oracle, pfType, discoveredAt, client, repo, version)
	case ds.CompositeRedStonePF:
		return composite_redstone_price_feed.NewRedstonePriceFeed(token, oracle, pfType, discoveredAt, client, repo, version)
	case ds.YearnPF:
		return yearn_price_feed.NewYearnPriceFeed(token, oracle, pfType, discoveredAt, client, repo, version)
	case ds.SingleAssetPF:
		return single_asset_feed.NewSingleAsset(token, oracle, pfType, discoveredAt, client, repo, version, underlyingFeeds)
	default:
		return nil
	}
}
func NewQueryPriceFeedFromAdapter(adapter *ds.SyncAdapter) ds.QueryPriceFeedI {
	switch adapter.GetDetailsByKey("pfType") {
	case ds.RedStonePF:
		return redstone_price_feed.NewRedstonePriceFeedFromAdapter(adapter)
	case ds.CurvePF:
		return curve_price_feed.NewCurvePriceFeedFromAdapter(adapter)
	case ds.CompositeRedStonePF:
		return composite_redstone_price_feed.NewRedstonePriceFeedFromAdapter(adapter)
		// return curve_price_feed.NewCurvePriceFeedFromAdapter(adapter)
	case ds.YearnPF:
		return yearn_price_feed.NewYearnPriceFeedFromAdapter(adapter)
	case ds.SingleAssetPF:
		return single_asset_feed.NewSingleAssetFromAdapter(adapter)
	default:
		return nil
	}
}
func FromAdapter(obj ds.SyncAdapterI) ds.QueryPriceFeedI {
	switch adapter := obj.(type) {
	case *curve_price_feed.CurvePriceFeed:
		return adapter
	case *yearn_price_feed.YearnPriceFeed:
		return adapter
	case *redstone_price_feed.RedstonePriceFeed:
		return adapter
	case *composite_redstone_price_feed.CompositeRedStonePriceFeed:
		return adapter
	case *single_asset_feed.SingleAssetFeed:
		return adapter
	default:
		return nil
	}
}
