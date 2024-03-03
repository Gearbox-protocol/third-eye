package aggregated_block_feed

import (
	"fmt"
	"math/big"
	"sort"
	"sync"
	"time"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/aggregated_block_feed/query_price_feed"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
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
		ch <- 1
		wg.Add(1)
		go mdl.queryAsync(blockNum, ch, wg)
		mdl.queryRedStone(blockNum)
		if rounds%100 == 0 {
			timeLeft := (time.Since(loopStartTime).Seconds() * float64(toSinceTill-blockNum)) /
				float64(blockNum-mdl.GetLastSync())
			timeLeft /= 60
			log.Infof("Synced %d in %d rounds(%fs): TimeLeft %f mins", blockNum, rounds, time.Since(roundStartTime).Seconds(), timeLeft)
			roundStartTime = time.Now()
		}
		rounds++
	}
	wg.Wait()
	return blockNum - mdl.Interval
}
func (mdl *AQFWrapper) queryRedStone(blockNum int64) {
	// for redstone
	for _, adapter := range mdl.QueryFeeds {
		// fetch from redstone
		validTokens := adapter.TokensValidAtBlock(blockNum)
		if adapter.GetPFType() == ds.RedStonePF &&
			adapter.GetLastSync() < blockNum &&
			len(validTokens) > 0 && adapter.GetPFType() == ds.RedStonePF { // if adapter has redstone token, then fetch from redstone
			priceBI := mdl.Repo.GetRedStonemgr().GetPrice(int64(mdl.Repo.SetAndGetBlock(blockNum).Timestamp), validTokens[0].Token)
			//
			isPriceInUSD := adapter.GetVersion().IsPriceInUSD() // should be always true

			priceData := parsePriceForRedStone(priceBI, isPriceInUSD, adapter.GetAddress())
			log.Infof("RedStone price for %s at %d is %f", mdl.Repo.GetToken(validTokens[0].Token).Symbol, blockNum, priceData.Price)
			//
			priceFeeds := []*schemas.PriceFeed{}
			for _, entry := range adapter.TokensValidAtBlock(blockNum) {
				priceDataCopy := priceData.Clone()
				//
				priceDataCopy.BlockNumber = blockNum
				priceDataCopy.Token = entry.Token
				priceDataCopy.MergedPFVersion = entry.MergedPFVersion
				priceDataCopy.Feed = adapter.GetAddress()
				//
				priceFeeds = append(priceFeeds, priceDataCopy)
			}
			mdl.updateQueryPrices(priceFeeds)
		}
	}
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

func (mdl *AQFWrapper) addQueryPrices(clearExtraBefore int64) {
	mdl.updateQueryPrices(mdl.queryPFdeps.extraPriceForQueryFeed(clearExtraBefore))
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
	pfs := mdl.QueryData(blockNum)
	mdl.updateQueryPrices(pfs)
	<-ch
	wg.Done()
}

func (mdl *AQFWrapper) updateQueryPrices(pfs []*schemas.PriceFeed) {
	mdl.mu.Lock()
	defer mdl.mu.Unlock()
	for _, pf := range pfs {
		if pf.MergedPFVersion == 0 {
			log.Fatalf("MergedPFVersion is 0 for %s", pf)
		}
	}
	mdl.queryFeedPrices = append(mdl.queryFeedPrices, pfs...)
}

func (mdl *AQFWrapper) QueryData(blockNum int64) []*schemas.PriceFeed {
	calls, queryAbleAdapters := mdl.getRoundDataCalls(blockNum)
	result := core.MakeMultiCall(mdl.Client, blockNum, false, calls)
	//
	//
	var queryFeedPrices []*schemas.PriceFeed
	for ind, entry := range result {
		pf := mdl.processRoundData(blockNum, queryAbleAdapters[ind], entry)
		queryFeedPrices = append(queryFeedPrices, pf...)
	}
	//
	return queryFeedPrices
}

func (mdl *AQFWrapper) getRoundDataCalls(blockNum int64) (calls []multicall.Multicall2Call, queryAbleAdapters []*query_price_feed.QueryPriceFeed) {
	priceFeedABI := core.GetAbi("PriceFeed")
	//
	for _, adapter := range mdl.QueryFeeds {
		if blockNum <= adapter.GetLastSync() || len(adapter.TokensValidAtBlock(blockNum)) == 0 {
			continue
		}
		data, err := priceFeedABI.Pack("latestRoundData")
		log.CheckFatal(err)
		call := multicall.Multicall2Call{
			Target:   common.HexToAddress(adapter.GetAddress()),
			CallData: data,
		}
		calls = append(calls, call)
		queryAbleAdapters = append(queryAbleAdapters, adapter)
	}
	return
}

var curvePFLatestRoundDataTimer = map[string]log.TimerFn{}

func (mdl *AQFWrapper) processRoundData(blockNum int64, adapter *query_price_feed.QueryPriceFeed, entry multicall.Multicall2Result) []*schemas.PriceFeed {
	var priceData *schemas.PriceFeed

	if adapter.GetPFType() == ds.RedStonePF { // 20 blocks
		return nil
	} else if entry.Success {
		isPriceInUSD := adapter.GetVersion().IsPriceInUSD()
		priceData = parseRoundData(entry.ReturnData, isPriceInUSD, adapter.GetAddress())
	} else if adapter.GetVersion().MoreThanEq(core.NewVersion(300)) {
		if core.GetChainId(mdl.Client) == 7878 {
			return nil
		} else {
			log.Fatalf("Can't get latestRounData in AQFWrapper for %s(%s) at %d",
				adapter.GetDetailsByKey("pfType"), adapter.GetAddress(), blockNum)
		}
	} else {
		switch adapter.GetDetailsByKey("pfType") {
		case ds.YearnPF:
			// fail on err, since we only sync for block_num which is more than discovered_at, we can assume that underlying price feed will be set for given block_num
			_priceData, err := adapter.CalculateYearnPFInternally(blockNum)
			if err != nil {
				log.Fatal(fmt.Errorf("At %d can't calculate yearnfeed(%s)'s price internally: %s",
					blockNum,
					adapter.GetAddress(), err.Error()))
			}
			priceData = _priceData
		case ds.CurvePF:
			// if virtualprice of pool for this oracle is not within lowerBound and upperBound , ignore the price
			oracleAddr := common.HexToAddress(adapter.GetAddress())
			virtualPrice := GetCurveVirtualPrice(blockNum, oracleAddr, adapter.GetVersion(), mdl.Client)
			//
			withinLimits := func() bool {
				lowerLimit, err := core.CallFuncWithExtraBytes(mdl.Client, "a384d6ff", oracleAddr, blockNum, nil) // lowerBound
				log.CheckFatal(err)
				upperLimit, err := core.CallFuncWithExtraBytes(mdl.Client, "b09ad8a0", oracleAddr, blockNum, nil) // upperBound
				log.CheckFatal(err)
				return new(big.Int).SetBytes(lowerLimit).Cmp(virtualPrice) < 0 &&
					new(big.Int).SetBytes(upperLimit).Cmp(virtualPrice) > 0
			}()
			if curvePFLatestRoundDataTimer[adapter.GetAddress()] == nil {
				curvePFLatestRoundDataTimer[adapter.GetAddress()] = log.GetRiskMsgTimer()
			}
			var msg string
			if !withinLimits {
				msg = "virtual price is not within limits for " + adapter.GetAddress()
			} else {
				msg = "failing due to unknown reason maybe underlying pricefeed of curve pool token is failing for curve adapter" + adapter.GetAddress()
			}
			log.SendRiskAlertPerTimer(
				log.RiskAlert{
					Msg: msg,
					RiskHeader: log.RiskHeader{
						BlockNumber: blockNum,
						EventCode:   "CURVE_LATEST_ROUNDDATA_FAIL",
					},
				},
				curvePFLatestRoundDataTimer[adapter.GetAddress()],
				86400*time.Second,
			)
			return nil
		default:
			log.Fatalf("Can't get latestRounData in AQFWrapper for %s(%s) at %d", adapter.GetDetailsByKey("pfType"), adapter.GetAddress(), blockNum)
		}
	}
	priceFeeds := []*schemas.PriceFeed{}
	for _, entry := range adapter.TokensValidAtBlock(blockNum) {
		priceDataCopy := priceData.Clone()
		//
		priceDataCopy.BlockNumber = blockNum
		priceDataCopy.Token = entry.Token
		priceDataCopy.MergedPFVersion = entry.MergedPFVersion
		priceDataCopy.Feed = adapter.GetAddress()
		//
		priceFeeds = append(priceFeeds, priceDataCopy)
	}
	return priceFeeds
}

func parseRoundData(returnData []byte, isPriceInUSD bool, feed string) *schemas.PriceFeed {
	priceFeedABI := core.GetAbi("PriceFeed")
	roundData := schemas.LatestRounData{}
	value, err := priceFeedABI.Unpack("latestRoundData", returnData)
	if err != nil {
		log.Fatalf("For feed(%s) can't get the lastestRounData: %s", feed, err)
	}
	roundData.RoundId = *abi.ConvertType(value[0], new(*big.Int)).(**big.Int)
	roundData.Answer = *abi.ConvertType(value[1], new(*big.Int)).(**big.Int)
	// roundData.StartedAt = *abi.ConvertType(value[2], new(*big.Int)).(**big.Int)
	// roundData.UpdatedAt = *abi.ConvertType(value[3], new(*big.Int)).(**big.Int)
	// roundData.AnsweredInRound = *abi.ConvertType(value[4], new(*big.Int)).(**big.Int)
	var decimals int8 = 18 // for eth
	if isPriceInUSD {
		decimals = 8 // for usd
	}
	return &schemas.PriceFeed{
		RoundId: roundData.RoundId.Int64(),
		PriceBI: (*core.BigInt)(roundData.Answer),
		Price:   utils.GetFloat64Decimal(roundData.Answer, decimals),
	}
}

func parsePriceForRedStone(price *big.Int, isPriceInUSD bool, feed string) *schemas.PriceFeed {
	var decimals int8 = 18 // for eth
	if isPriceInUSD {
		decimals = 8 // for usd
	}
	return &schemas.PriceFeed{
		RoundId: 0,
		PriceBI: (*core.BigInt)(price),
		Price:   utils.GetFloat64Decimal(price, decimals),
	}
}
