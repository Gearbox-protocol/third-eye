package aggregated_block_feed

import (
	"math"
	"sync"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/core/types"
)

type AggregatedBlockFeed struct {
	*ds.SyncAdapter
	mu *sync.Mutex
	// yearn feed
	QueryFeeds map[string]*QueryPriceFeed
	// yearn feed prices to be ordered
	queryFeedPrices []*schemas.PriceFeed
	// uniswap price related data structures
	UniswapPools      []string
	UniPoolByToken    map[string]*schemas.UniswapPools
	UniPricesByTokens map[string]schemas.SortedUniPoolPrices
	tokenInfos        map[string]*schemas.Token
	//
	TokenLastSync map[string]int64
	// intervel from config
	Interval int64
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
		QueryFeeds:        map[string]*QueryPriceFeed{},
	}
}

// only called by priceoracle
func (mdl *AggregatedBlockFeed) AddYearnFeed(adapter ds.SyncAdapterI) {
	yearnFeed, ok := adapter.(*QueryPriceFeed)
	if !ok {
		log.Fatal("Failed in parsing yearn feed for aggregated yearn feed")
	}
	mdl.LastSync = utils.Min(adapter.GetLastSync(), mdl.LastSync)
	// log.Info(adapter.GetAddress(), "added to aggregatedpricefeed has last_sync", adapter.GetLastSync())
	mdl.QueryFeeds[adapter.GetAddress()] = yearnFeed
}

func (mdl *AggregatedBlockFeed) OnLog(txLog types.Log) {
}

func (mdl *AggregatedBlockFeed) GetQueryFeeds() []*QueryPriceFeed {
	feeds := []*QueryPriceFeed{}
	for _, feed := range mdl.QueryFeeds {
		feeds = append(feeds, feed)
	}
	return feeds
}

func (mdl *AggregatedBlockFeed) AddUniPools(token *schemas.Token, uniswapPools *schemas.UniswapPools) {
	if mdl.UniPoolByToken[uniswapPools.Token] == nil {
		mdl.UniPoolByToken[uniswapPools.Token] = uniswapPools
	}
	mdl.tokenInfos[token.Address] = token
}

// for getting the uniswap prices for chainlink token/usdc uniswap pairs.
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

func (mdl *AggregatedBlockFeed) AddFeedOrToken(token, oracle string, pfType string, discoveredAt int64, version int16) {
	if mdl.QueryFeeds[oracle] != nil {
		mdl.QueryFeeds[oracle].AddToken(token, discoveredAt)
	} else {
		mdl.QueryFeeds[oracle] = NewQueryPriceFeed(token, oracle, pfType, discoveredAt, mdl.Client, mdl.Repo, version)
	}
}

func (mdl *AggregatedBlockFeed) DisableYearnFeed(token, oracle string, disabledAt int64) {
	mdl.QueryFeeds[oracle].DisableToken(token, disabledAt)
}
