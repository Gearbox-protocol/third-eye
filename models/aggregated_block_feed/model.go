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
	QueryFeeds map[string]ds.QueryPriceFeedI

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
		QueryFeeds:  map[string]ds.QueryPriceFeedI{},
		queryPFdeps: NewQueryPFDepenencies(repo, client),
	}
	wrapper.queryPFdeps.aqf = wrapper
	return wrapper
}

// only called by priceoracle
func (mdl *AQFWrapper) AddQueryPriceFeed(adapter ds.QueryPriceFeedI) {
	mdl.LastSync = utils.Min(adapter.GetLastSync(), mdl.LastSync)
	mdl.QueryFeeds[adapter.GetAddress()] = adapter
}

func (mdl *AQFWrapper) GetQueryFeeds() []ds.QueryPriceFeedI {
	feeds := make([]ds.QueryPriceFeedI, 0, len(mdl.QueryFeeds))
	for _, feed := range mdl.QueryFeeds {
		feeds = append(feeds, feed)
	}
	return feeds
}

func (mdl *AQFWrapper) AddFeedOrToken(token, oracle string, pfType string, discoveredAt int64, pfVersion schemas.PFVersion, underlyingFeeds []string) {
	log.Infof("Add new %s:pfversion(%d) for token(%s): %s discovered at %d", pfType, pfVersion, token, oracle, discoveredAt)
	// MAINNET: yearn yvUSDC has changed over time, previous token was 0x5f18C75AbDAe578b483E5F43f12a39cF75b973a9(only added in gearbox v1 priceOracle) and 0xa354F35829Ae975e850e23e9615b11Da1B3dC4DE, so we can ignore 0xc1 yvUSDC token dependency
	if token != "0x5f18C75AbDAe578b483E5F43f12a39cF75b973a9" {
		mdl.queryPFdeps.checkInDepGraph(token, oracle, discoveredAt)
	}
	if mdl.QueryFeeds[oracle] != nil {
		mdl.QueryFeeds[oracle].AddToken(token, discoveredAt, pfVersion)
	} else {
		mdl.AddQueryPriceFeed(NewQueryPriceFeed(token, oracle, pfType, discoveredAt, mdl.Client, mdl.Repo, pfVersion, underlyingFeeds))
		// MAINNET: old yvUSDC added on gearbox v1
		if token == "0x5f18C75AbDAe578b483E5F43f12a39cF75b973a9" {
			mdl.QueryFeeds[oracle].DisableToken(token, 13856183, pfVersion) // new yvUSDC added on gearbox v1
		}
	}
	// when token is added to the queryPricefeed, add price object at discoveredAt
	// so that  accounts opened just after discoveredAt can get the price from db
	mdl.updateQueryPrices(createPriceFeedOnInit(mdl.QueryFeeds[oracle], mdl.Client, token, discoveredAt))
}

func mergePFVersionAt(blockNum int64, details map[schemas.PFVersion][]int64) schemas.MergedPFVersion {
	var pfVersion schemas.MergedPFVersion = 0
	for version, blockNums := range details {
		// log.Info(version, blockNums, blockNum)
		if blockNums[0] <= blockNum && (len(blockNums) == 1 || blockNum < blockNums[1]) { // 1 is added as price is already added at blockNum
			pfVersion = pfVersion | schemas.MergedPFVersion(version)
		}
	}
	return pfVersion
}
func createPriceFeedOnInit(qpf ds.QueryPriceFeedI, client core.ClientI, token string, discoveredAt int64) []*schemas.PriceFeed {
	mainPFContract, err := priceFeed.NewPriceFeed(common.HexToAddress(qpf.GetAddress()), client)
	log.CheckFatal(err)
	data, err := mainPFContract.LatestRoundData(&bind.CallOpts{BlockNumber: big.NewInt(discoveredAt)})
	log.CheckFatal(err)
	//
	pfVersion := mergePFVersionAt(discoveredAt, qpf.GetTokens()[token])
	return []*schemas.PriceFeed{{
		BlockNumber:     discoveredAt,
		Feed:            qpf.GetAddress(),
		Token:           token,
		RoundId:         data.RoundId.Int64(),
		MergedPFVersion: pfVersion,
		PriceBI:         (*core.BigInt)(data.Answer),
		Price:           utils.GetFloat64Decimal(data.Answer, pfVersion.Decimals()),
	}}
}

func (mdl *AQFWrapper) DisableYearnFeed(token, oracle string, disabledAt int64, pfVersion schemas.PFVersion) {
	mdl.QueryFeeds[oracle].DisableToken(token, disabledAt, pfVersion)
}

func (mdl AQFWrapper) GetDepFetcher() *QueryPFDependencies {
	return mdl.queryPFdeps
}

func (mdl *AQFWrapper) OnLog(txLog types.Log) {
}

// no need to check version of feed, as while adding from chainlink we make sure that the version is more than 1
// and  we can't have version 2 and 3 feed active at the same time.
func (mdl AQFWrapper) getFeedAdapters(blockNum int64, neededTokens map[string]bool) (result []ds.QueryPriceFeedI) {
	for _, adapter := range mdl.QueryFeeds {
		if !adapter.GetVersion().MoreThan(core.NewVersion(1)) {
			continue
		}
		tokensForAdapter := adapter.TokensValidAtBlock(blockNum)
		for _, entry := range tokensForAdapter {
			if neededTokens[entry.Token] {
				result = append(result, adapter)
				break
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
