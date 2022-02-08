package aggregated_block_feed

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/utils"
	"github.com/ethereum/go-ethereum/core/types"
	"math"
	"sync"
)

type AggregatedBlockFeed struct {
	*core.SyncAdapter
	YearnFeeds        []*YearnPriceFeed
	UniswapPools      []string
	UniPoolByToken    map[string]*core.UniswapPools
	TokenLastSync     map[string]int64
	UniPricesByTokens map[string]core.SortedUniPoolPrices
	Interval          int64
	mu                *sync.Mutex
	tokenInfos        map[string]*core.Token
}

func NewAggregatedBlockFeed(client ethclient.ClientI, repo core.RepositoryI, interval int64) *AggregatedBlockFeed {
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
		UniPoolByToken:    map[string]*core.UniswapPools{},
		UniPricesByTokens: map[string]core.SortedUniPoolPrices{},
		SyncAdapter:       syncAdapter,
		TokenLastSync:     map[string]int64{},
		Interval:          interval,
		mu:                &sync.Mutex{},
		tokenInfos:        map[string]*core.Token{},
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

func (mdl *AggregatedBlockFeed) AddPools(token *core.Token, uniswapPools *core.UniswapPools) {
	if mdl.UniPoolByToken[uniswapPools.Token] == nil {
		mdl.UniPoolByToken[uniswapPools.Token] = uniswapPools
	}
	mdl.tokenInfos[token.Address] = token
}

func (mdl *AggregatedBlockFeed) AddLastSyncForToken(token string, lastSync int64) {
	mdl.LastSync = utils.Min(lastSync, mdl.LastSync)
	// there is new oracle/feed added for a token
	if mdl.TokenLastSync[token] == 0 {
		mdl.TokenLastSync[token] = lastSync
	}
	mdl.TokenLastSync[token] = utils.Min(mdl.TokenLastSync[token], lastSync)
}

func (mdl *AggregatedBlockFeed) GetUniswapPools() (updatedPools []*core.UniswapPools) {
	for _, entry := range mdl.UniPoolByToken {
		if entry.Updated {
			updatedPools = append(updatedPools, entry)
		}
		entry.Updated = false
	}
	return
}
