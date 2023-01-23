package aggregated_block_feed

import (
	"math"
	"math/big"
	"sync"

	"github.com/Gearbox-protocol/sdk-go/artifacts/priceFeed"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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
			// if no yearn feed is added , then lastSync = math.MaxInt64 can overflow.
			LastSync: math.MaxInt64 - 1,
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
	log.Infof("Add new %s for token(%s): %s discovered at %d", pfType, token, oracle, discoveredAt)
	// MAINNET: yearn yvUSDC has changed over time, previous token was 0x5f18C75AbDAe578b483E5F43f12a39cF75b973a9(only added in gearbox v1 priceOracle) and 0xa354F35829Ae975e850e23e9615b11Da1B3dC4DE, so we can ignore 0xc1 yvUSDC token dependency
	if token != "0x5f18C75AbDAe578b483E5F43f12a39cF75b973a9" {
		mdl.queryPFdeps.checkInDepGraph(token, oracle, discoveredAt)
	}
	if mdl.QueryFeeds[oracle] != nil {
		mdl.QueryFeeds[oracle].AddToken(token, discoveredAt)
	} else {
		mdl.QueryFeeds[oracle] = NewQueryPriceFeed(token, oracle, pfType, discoveredAt, mdl.Client, mdl.Repo, version)
	}
	// when token is added to the queryPricefeed, add price object at discoveredAt
	// so that  accounts opened just after discoveredAt can get the price from db
	mdl.addPriceForToken(mdl.QueryFeeds[oracle], token, discoveredAt)
}

func (mdl *AggregatedBlockFeed) addPriceForToken(qpf *QueryPriceFeed, token string, discoveredAt int64) {
	mainPFContract, err := priceFeed.NewPriceFeed(common.HexToAddress(qpf.Address), qpf.Client)
	log.CheckFatal(err)
	data, err := mainPFContract.LatestRoundData(&bind.CallOpts{BlockNumber: big.NewInt(discoveredAt)})
	log.CheckFatal(err)
	var decimals int8 = 8
	if mdl.GetVersion() == 1 {
		decimals = 18
	}
	mdl.updateQueryPrices([]*schemas.PriceFeed{{
		BlockNumber:  discoveredAt,
		Feed:         mdl.Address,
		Token:        token,
		RoundId:      data.RoundId.Int64(),
		IsPriceInUSD: mdl.GetVersion() > 1, // for version more than 1
		PriceBI:      (*core.BigInt)(data.Answer),
		Price:        utils.GetFloat64Decimal(data.Answer, decimals),
	}})
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
