package aggregated_block_feed

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/utils"
	"github.com/ethereum/go-ethereum/core/types"
	"math"
)

type UniswapPools struct {
	V2       string
	V3       string
	Decimals int8
}
type AggregatedBlockFeed struct {
	*core.SyncAdapter
	YearnFeeds        []*YearnPriceFeed
	UniswapPools      []string
	UniPoolByToken    map[string]*UniswapPools
	UniPricesByTokens map[string][]*core.PoolPrices
}

func NewAggregatedBlockFeed(client *ethclient.Client, repo core.RepositoryI) *AggregatedBlockFeed {
	syncAdapter := &core.SyncAdapter{
		Contract: &core.Contract{
			// Address:      oracle,
			// DiscoveredAt: discoveredAt,
			// FirstLogAt:   discoveredAt,
			ContractName: core.AggregatedBlockFeed,
			Client:       client,
		},
		LastSync:  math.MaxInt64,
		Repo:      repo,
		OnlyQuery: true,
	}
	return &AggregatedBlockFeed{
		UniPoolByToken:    map[string]*UniswapPools{},
		UniPricesByTokens: map[string][]*core.PoolPrices{},
		SyncAdapter:       syncAdapter,
	}
}

func (mdl *AggregatedBlockFeed) AddYearnFeed(adapter core.SyncAdapterI) {
	yearnFeed, ok := adapter.(*YearnPriceFeed)
	if !ok {
		log.Fatal("Failed in parsing yearn feed for aggregated yearn feed")
	}
	mdl.LastSync = utils.Min(adapter.GetLastSync(), mdl.LastSync)
	mdl.YearnFeeds = append(mdl.YearnFeeds, yearnFeed)
}

func (mdl *AggregatedBlockFeed) OnLog(txLog types.Log) {
}

func (mdl *AggregatedBlockFeed) GetYearnFeeds() []*YearnPriceFeed {
	return mdl.YearnFeeds
}

func (mdl *AggregatedBlockFeed) AddPools(token, poolv2Addr, poolv3Addr string, decimals int8) {
	mdl.UniPoolByToken[token] = &UniswapPools{
		V2: poolv2Addr,
		V3: poolv3Addr,
	}
}
