package aggregated_block_feed

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/core/types"
	"math"
	"sync"
)

type AggregatedBlockFeed struct {
	*ds.SyncAdapter
	YearnFeeds        []*YearnPriceFeed
	UniswapPools      []string
	UniPoolByToken    map[string]*schemas.UniswapPools
	TokenLastSync     map[string]int64
	UniPricesByTokens map[string]schemas.SortedUniPoolPrices
	Interval          int64
	mu                *sync.Mutex
	tokenInfos        map[string]*schemas.Token
}

func NewAggregatedBlockFeed(client core.ClientI, repo ds.RepositoryI, interval int64) *AggregatedBlockFeed {
	syncAdapter := &ds.SyncAdapter{
		SyncAdapterSchema: &schemas.SyncAdapterSchema{
			Contract: &schemas.Contract{
				// Address:      oracle,
				// DiscoveredAt: discoveredAt,
				// FirstLogAt:   discoveredAt,
				ContractName: ds.AggregatedBlockFeed,
				Client:       client,
			},
			LastSync: math.MaxInt64,
		},
		Repo:      repo,
		OnlyQuery: true,
	}
	return &AggregatedBlockFeed{
		UniPoolByToken:    map[string]*schemas.UniswapPools{},
		UniPricesByTokens: map[string]schemas.SortedUniPoolPrices{},
		SyncAdapter:       syncAdapter,
		TokenLastSync:     map[string]int64{},
		Interval:          interval,
		mu:                &sync.Mutex{},
		tokenInfos:        map[string]*schemas.Token{},
	}
}

func (mdl *AggregatedBlockFeed) AddYearnFeed(adapter ds.SyncAdapterI) {
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

func (mdl *AggregatedBlockFeed) AddUniPools(token *schemas.Token, uniswapPools *schemas.UniswapPools) {
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

func (mdl *AggregatedBlockFeed) GetUniswapPools() (updatedPools []*schemas.UniswapPools) {
	for _, entry := range mdl.UniPoolByToken {
		if entry.Updated {
			updatedPools = append(updatedPools, entry)
		}
		entry.Updated = false
	}
	return
}
