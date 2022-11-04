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

	// for getting uni price on v2, v3 and uni oracle
	priceOnUNIFetcher *PriceOnUNIFetcher
	// for dependency based fetching price
	queryPFdeps *QueryPFDependencies
	//
	queryFeedPrices []*schemas.PriceFeed
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
		SyncAdapter:       syncAdapter,
		Interval:          interval,
		mu:                &sync.Mutex{},
		QueryFeeds:        map[string]*QueryPriceFeed{},
		priceOnUNIFetcher: NewPriceOnUNIFetcher(),
		queryPFdeps:       NewQueryPFDepenencies(repo, client),
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

func (mdl *AggregatedBlockFeed) GetQueryFeeds() []*QueryPriceFeed {
	feeds := []*QueryPriceFeed{}
	for _, feed := range mdl.QueryFeeds {
		feeds = append(feeds, feed)
	}
	return feeds
}

func (mdl *AggregatedBlockFeed) AddFeedOrToken(token, oracle string, pfType string, discoveredAt int64, version int16) {
	mdl.queryPFdeps.checkInDepGraph(token, oracle, discoveredAt)
	if mdl.QueryFeeds[oracle] != nil {
		mdl.QueryFeeds[oracle].AddToken(token, discoveredAt)
	} else {
		mdl.QueryFeeds[oracle] = NewQueryPriceFeed(token, oracle, pfType, discoveredAt, mdl.Client, mdl.Repo, version)
	}
}

func (mdl *AggregatedBlockFeed) DisableYearnFeed(token, oracle string, disabledAt int64) {
	mdl.QueryFeeds[oracle].DisableToken(token, disabledAt)
}

func (mdl AggregatedBlockFeed) UNIFetcher() *PriceOnUNIFetcher {
	return mdl.priceOnUNIFetcher
}
func (mdl AggregatedBlockFeed) GetDepFetcher() *QueryPFDependencies {
	return mdl.queryPFdeps
}

func (mdl *AggregatedBlockFeed) OnLog(txLog types.Log) {
}

func (mdl AggregatedBlockFeed) Clear() {
	mdl.priceOnUNIFetcher.Clear()
}
