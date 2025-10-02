package aggregated_block_feed

import (
	"sort"
	"sync"
	"time"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	// "fmt"
)

func (mdl *AQFWrapper) fetchAllPrices(toSinceTill int64) int64 {
	queryFrom := mdl.GetLastSync() + mdl.Interval
	if queryFrom > toSinceTill {
		return mdl.GetLastSync()
	}
	log.Infof("Sync %s from %d to %d", mdl.GetName(), queryFrom, toSinceTill)
	// for concurrency
	concurrentThreads := 6
	ch := make(chan int, concurrentThreads)
	wg := &sync.WaitGroup{}
	//
	// timer with query of block
	rounds := 0
	loopStartTime := time.Now()
	roundStartTime := time.Now()

	blockNum := mdl.GetLastSync() + mdl.Interval
	for ; blockNum <= toSinceTill; blockNum += mdl.Interval {
		mdl.queryPFdeps.aggregatedFetchedBlocks =
			append(mdl.queryPFdeps.aggregatedFetchedBlocks, blockNum)
		mdl.queryAsync(blockNum, ch, wg)
		if rounds%100 == 0 {
			timeLeft := (time.Since(loopStartTime).Seconds() * float64(toSinceTill-blockNum)) /
				float64(blockNum-mdl.GetLastSync())
			timeLeft /= 60
			log.Infof("Synced %d(inv: %d) in %d rounds(%fs): TimeLeft %f mins", blockNum, mdl.Interval, rounds, time.Since(roundStartTime).Seconds(), timeLeft)
			roundStartTime = time.Now()
		}
		rounds++
	}
	wg.Wait()
	return blockNum - mdl.Interval
}

// update all queryAdapter only if for the lastBlockNumber as we have fetched prices for that block.
// not update for toSinceTill.  if the interval is more than the syncCycle in engin/index.go, the queryAdpater lastsync will be updated but the prices will not be fetched
func (mdl *AQFWrapper) Query(toSinceTill int64) {
	if len(mdl.QueryFeeds) == 0 {
		return
	}

	syncedTill := mdl.fetchAllPrices(toSinceTill)
	//
	// set last_sync on querypricefeed
	for _, adapter := range mdl.QueryFeeds {
		// yearn price feed can't be disabled from v2
		if syncedTill <= adapter.GetLastSync() || adapter.IsDisabled() {
			continue
		}
		adapter.AfterSyncHook(syncedTill)
	}
	mdl.addQueryPrices(mdl.GetLastSync()) // use previous lastSync for getting extra prices
	//
	mdl.LastSync = syncedTill
}

func (mdl *AQFWrapper) addQueryPrices(deleteExtraBefore int64) {
	mdl.updateQueryPrices(mdl.queryPFdeps.extraPriceForQueryFeed(deleteExtraBefore))
	// query feed prices
	sort.SliceStable(mdl.queryFeedPrices, func(i, j int) bool {
		return mdl.queryFeedPrices[i].BlockNumber < mdl.queryFeedPrices[j].BlockNumber
	})
	for _, queryPrice := range mdl.queryFeedPrices {
		mdl.Repo.AddPriceFeed(queryPrice)
	}
	mdl.queryFeedPrices = nil
}

func (mdl *AQFWrapper) queryAsync(blockNum int64, ch chan int, wg *sync.WaitGroup) {
	ch <- 1
	wg.Add(1)
	calls, queryAbleAdapters := mdl.getRoundDataCalls(blockNum)
	go func() {
		pfs := mdl.QueryData(calls, queryAbleAdapters, blockNum)
		mdl.updateQueryPrices(pfs)
		<-ch
		wg.Done()
	}()
}

func (mdl *AQFWrapper) updateQueryPrices(pfs []*schemas.PriceFeed) {
	mdl.mu.Lock()
	defer mdl.mu.Unlock()
	// for _, pf := range pfs {
	// 	if pf.MergedPFVersion == 0 {
	// 		log.Fatalf("MergedPFVersion is 0 for %s", pf)
	// 	}
	// }
	mdl.queryFeedPrices = append(mdl.queryFeedPrices, pfs...)
}

func (mdl *AQFWrapper) QueryData(calls []multicall.Multicall2Call, queryAbleAdapters []adapterAndNoCall, blockNum int64) []*schemas.PriceFeed {

	result := core.MakeMultiCall(mdl.Client, blockNum, false, calls)
	iterator := core.NewMulticallResultIterator(result)
	//
	//
	var queryFeedPrices []*schemas.PriceFeed
	for _, entry := range queryAbleAdapters {
		var results []multicall.Multicall2Result
		for i := 0; i < entry.nocalls; i++ {
			results = append(results, iterator.Next())
		}
		pf := processRoundDataWithAdapterTokens(blockNum, entry.adapter, results, entry.force)
		queryFeedPrices = append(queryFeedPrices, pf...)
	}
	//
	return queryFeedPrices
}

type adapterAndNoCall struct {
	adapter ds.QueryPriceFeedI
	nocalls int
	force   bool
}

func (mdl *AQFWrapper) getRoundDataCalls(blockNum int64) (calls []multicall.Multicall2Call, queryAbleAdapters []adapterAndNoCall) {
	//
	for _, adapter := range mdl.QueryFeeds {
		// log.Info(adapter.GetAddress(), len(mdl.Repo.TokenAddrsValidAtBlock(adapter.GetAddress(), blockNum)))
		if blockNum <= adapter.GetLastSync() || len(mdl.Repo.TokenAddrsValidAtBlock(adapter.GetAddress(), blockNum)) == 0 {
			continue
		}
		moreCalls, isQueryable := adapter.GetCalls(blockNum)
		if isQueryable {
			calls = append(calls, moreCalls...)
			queryAbleAdapters = append(queryAbleAdapters, adapterAndNoCall{
				adapter: adapter,
				nocalls: len(moreCalls),
				force:   getForceForAdapter(mdl.Repo, adapter, blockNum),
			})
			continue
		}
	}
	return
}

func processRoundDataWithAdapterTokens(blockNum int64, adapter ds.QueryPriceFeedI, entries []multicall.Multicall2Result, force bool) []*schemas.PriceFeed {

	// } else if utils.Contains([]string{"0xCbeCfA4017965939805Da5a2150E3DB1BeDD0364", "0x814E6564e8cda436c1ab25041C10bfdb21dEC519"},

	priceData := adapter.ProcessResult(blockNum, entries, "", force)
	if priceData == nil {
		return nil
	}
	priceData.Feed = adapter.GetAddress()
	priceData.BlockNumber = blockNum

	return []*schemas.PriceFeed{priceData}
}

var _lastBlock = core.NewMutexDS[string, int64]()

func getForceForAdapter(repo ds.RepositoryI, adapter ds.QueryPriceFeedI, newblock int64) bool { // so that difference is 1 hr.
	var lastBlock int64
	feed := adapter.GetAddress()
	if price := repo.GetPrevPriceFeed(feed); price != nil {
		lastBlock = price.BlockNumber
	}
	lastBlock = utils.Max(lastBlock, _lastBlock.Get(feed)) // max of prevPriceStore and local store
	if lastBlock == 0 {
		return true
	}

	force := time.Duration(repo.SetAndGetBlock(newblock).Timestamp-repo.SetAndGetBlock(lastBlock).Timestamp)*time.Second > time.Hour
	if force {
		log.Info(feed, " last price at ", lastBlock)
		_lastBlock.Set(feed, utils.Max(_lastBlock.Get(feed), newblock))
	}
	return force
}
