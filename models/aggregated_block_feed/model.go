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

type AQFWrapper struct {
	*ds.SyncAdapter
	mu *sync.Mutex
	// yearn feed
	QueryFeeds map[string]*QueryPriceFeed

	// for dependency based fetching price
	queryPFdeps *QueryPFDependencies
	//
	queryFeedPrices []*schemas.PriceFeed
	// intervel from config
	Interval int64
}

// not present in db , manaully added in syncadapter repository handler
// last_sync is dependent on min(QueryPriceFeed's last_sync)
func NewAQFWrapper(client core.ClientI, repo ds.RepositoryI, interval int64) *AQFWrapper {
	syncAdapter := &ds.SyncAdapter{
		SyncAdapterSchema: &schemas.SyncAdapterSchema{
			Contract: &schemas.Contract{
				// DiscoveredAt: discoveredAt,
				// FirstLogAt:   discoveredAt,
				Address:      ds.AggregatedQueryFeedWrapper,
				ContractName: ds.AggregatedQueryFeedWrapper,
				Client:       client,
			},
			// if no yearn feed is added , then lastSync = math.MaxInt64 can overflow.
			LastSync: math.MaxInt64 - 1,
		},
		Repo:            repo,
		DataProcessType: ds.ViaQuery,
	}
	wrapper := &AQFWrapper{
		SyncAdapter: syncAdapter,
		Interval:    interval,
		mu:          &sync.Mutex{},
		QueryFeeds:  map[string]*QueryPriceFeed{},
		queryPFdeps: NewQueryPFDepenencies(repo, client),
	}
	wrapper.queryPFdeps.aqf = wrapper
	return wrapper
}

// only called by priceoracle
func (mdl *AQFWrapper) AddYearnFeed(adapter ds.SyncAdapterI) {
	yearnFeed, ok := adapter.(*QueryPriceFeed)
	if !ok {
		log.Fatal("Failed in parsing yearn feed for aggregated yearn feed")
	}
	mdl.LastSync = utils.Min(adapter.GetLastSync(), mdl.LastSync)
	mdl.QueryFeeds[adapter.GetAddress()] = yearnFeed
}

func (mdl *AQFWrapper) GetQueryFeeds() []*QueryPriceFeed {
	feeds := make([]*QueryPriceFeed, 0, len(mdl.QueryFeeds))
	for _, feed := range mdl.QueryFeeds {
		feeds = append(feeds, feed)
	}
	return feeds
}

func (mdl *AQFWrapper) AddFeedOrToken(token, oracle string, pfType string, discoveredAt int64, version core.VersionType) {
	log.Infof("Add new %s for token(%s): %s discovered at %d", pfType, token, oracle, discoveredAt)
	// MAINNET: yearn yvUSDC has changed over time, previous token was 0x5f18C75AbDAe578b483E5F43f12a39cF75b973a9(only added in gearbox v1 priceOracle) and 0xa354F35829Ae975e850e23e9615b11Da1B3dC4DE, so we can ignore 0xc1 yvUSDC token dependency
	if token != "0x5f18C75AbDAe578b483E5F43f12a39cF75b973a9" {
		mdl.queryPFdeps.checkInDepGraph(token, oracle, discoveredAt)
	}
	if mdl.QueryFeeds[oracle] != nil {
		mdl.QueryFeeds[oracle].AddToken(token, discoveredAt)
	} else {
		mdl.AddYearnFeed(NewQueryPriceFeed(token, oracle, pfType, discoveredAt, mdl.Client, mdl.Repo, version))
		// MAINNET: old yvUSDC added on gearbox v1
		if token == "0x5f18C75AbDAe578b483E5F43f12a39cF75b973a9" {
			mdl.QueryFeeds[oracle].DisableToken(token, 13856183) // new yvUSDC added on gearbox v1
		}
	}
	// when token is added to the queryPricefeed, add price object at discoveredAt
	// so that  accounts opened just after discoveredAt can get the price from db
	mdl.updateQueryPrices(createPriceFeedOnInit(mdl.QueryFeeds[oracle], token, discoveredAt))
}

func createPriceFeedOnInit(qpf *QueryPriceFeed, token string, discoveredAt int64) []*schemas.PriceFeed {
	mainPFContract, err := priceFeed.NewPriceFeed(common.HexToAddress(qpf.Address), qpf.Client)
	log.CheckFatal(err)
	data, err := mainPFContract.LatestRoundData(&bind.CallOpts{BlockNumber: big.NewInt(discoveredAt)})
	log.CheckFatal(err)
	return []*schemas.PriceFeed{{
		BlockNumber:  discoveredAt,
		Feed:         qpf.Address,
		Token:        token,
		RoundId:      data.RoundId.Int64(),
		IsPriceInUSD: qpf.GetVersion().IsPriceInUSD(), // for version more than 1
		PriceBI:      (*core.BigInt)(data.Answer),
		Price:        utils.GetFloat64Decimal(data.Answer, qpf.GetVersion().Decimals()),
	}}
}

func (mdl *AQFWrapper) DisableYearnFeed(token, oracle string, disabledAt int64) {
	mdl.QueryFeeds[oracle].DisableToken(token, disabledAt)
}

func (mdl AQFWrapper) GetDepFetcher() *QueryPFDependencies {
	return mdl.queryPFdeps
}

func (mdl *AQFWrapper) OnLog(txLog types.Log) {
}

type local struct {
	Feed  string
	Token string
}

// no need to check version of feed, as while adding from chainlink we make sure that the version is more than 1
// and  we can't have version 2 and 3 feed active at the same time.
func (mdl AQFWrapper) getFeeds(blockNum int64, neededTokens map[string]bool) (result []local) {
	for _, adapter := range mdl.QueryFeeds {
		if !adapter.GetVersion().MoreThan(core.NewVersion(1)) {
			continue
		}
		tokensForAdapter := adapter.TokensValidAtBlock(blockNum)
		for _, token := range tokensForAdapter {
			if neededTokens[token] {
				result = append(result, local{
					Feed:  adapter.GetAddress(),
					Token: token,
				})
			}
		}
	}
	return
}

func (mdl AQFWrapper) ChainlinkPriceUpdatedAt(token string, blockNums []int64) {
	mdl.queryPFdeps.chainlinkPriceUpdatedAt(token, blockNums)
}

func (mdl AQFWrapper) AfterSyncHook(blockNum int64) {
	// don't do any thing as the lastSync should not be updated from outside.
	// It is  the min of the lastsync of all the interally managed queryPriceFeeds
}
