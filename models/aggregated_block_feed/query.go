package aggregated_block_feed

import (
	"math/big"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	// "fmt"
)

func (mdl *AggregatedBlockFeed) Query(queryTill int64) {
	if len(mdl.priceOnUNIFetcher.UniPoolByToken) == 0 && len(mdl.QueryFeeds) == 0 {
		return
	}
	concurrentThreads := 6
	ch := make(chan int, concurrentThreads)
	// msg
	queryFrom := mdl.GetLastSync() + mdl.Interval
	log.Infof("Sync %s from %d to %d", mdl.GetName(), queryFrom, queryTill)
	// timer with query of block
	rounds := 0
	loopStartTime := time.Now()
	roundStartTime := time.Now()
	wg := &sync.WaitGroup{}
	for blockNum := queryFrom; blockNum <= queryTill; blockNum += mdl.Interval {
		mdl.queryPFdeps.aggregatedFetchedBlocks =
			append(mdl.queryPFdeps.aggregatedFetchedBlocks, blockNum)
		ch <- 1
		wg.Add(1)
		go mdl.queryAsync(blockNum, ch, wg)
		if rounds%100 == 0 {
			timeLeft := (time.Since(loopStartTime).Seconds() * float64(queryTill-blockNum)) /
				float64(blockNum-mdl.GetLastSync())
			timeLeft /= 60
			log.Infof("Synced %d in %d rounds(%fs): TimeLeft %f mins", blockNum, rounds, time.Since(roundStartTime).Seconds(), timeLeft)
			roundStartTime = time.Now()
		}
		rounds++
	}
	wg.Wait()
	// set last_sync on querypricefeed
	for _, adapter := range mdl.QueryFeeds {
		// yearn price feed can't be disabled from v2
		if queryTill <= adapter.GetLastSync() || adapter.IsDisabled() {
			continue
		}
		adapter.AfterSyncHook(queryTill)
	}
	mdl.priceOnUNIFetcher.sortUniPrices()
	mdl.addQueryPrices()
}

func (mdl *AggregatedBlockFeed) addQueryPrices() {
	mdl.updateQueryPrices(mdl.queryPFdeps.extraPriceForQueryFeed())
	// query feed prices
	sort.SliceStable(mdl.queryFeedPrices, func(i, j int) bool {
		return mdl.queryFeedPrices[i].BlockNumber < mdl.queryFeedPrices[j].BlockNumber
	})
	for _, queryPrice := range mdl.queryFeedPrices {
		mdl.Repo.AddPriceFeed(queryPrice)
	}
	mdl.queryFeedPrices = nil
}

func powFloat(a *big.Int) *big.Float {
	f := big.NewFloat(1.0001)
	ans := big.NewFloat(1)
	absA := new(big.Int).Abs(a)
	for i := 0; i < absA.BitLen(); i++ {
		if absA.Bit(i) == 1 {
			ans = new(big.Float).Mul(f, ans)
		}
		f = new(big.Float).Mul(f, f)
	}
	if absA == a {
		return ans
	}
	return new(big.Float).Quo(big.NewFloat(1), ans)
}

func (mdl *AggregatedBlockFeed) queryAsync(blockNum int64, ch chan int, wg *sync.WaitGroup) {
	weth := mdl.Repo.GetWETHAddr()
	pfs, uniPrices := mdl.QueryData(blockNum, weth, "all")
	mdl.updateQueryPrices(pfs)
	mdl.updateUniPrices(uniPrices)
	<-ch
	wg.Done()
}

func (mdl *AggregatedBlockFeed) updateUniPrices(pricesByToken map[string]*schemas.UniPoolPrices) {
	mdl.mu.Lock()
	defer mdl.mu.Unlock()
	for token, prices := range pricesByToken {
		mdl.Repo.AddUniswapPrices(prices)
		mdl.priceOnUNIFetcher.UniPricesByTokens[token] =
			append(mdl.priceOnUNIFetcher.UniPricesByTokens[token], prices)
	}
}

func (mdl *AggregatedBlockFeed) updateQueryPrices(pfs []*schemas.PriceFeed) {
	mdl.mu.Lock()
	defer mdl.mu.Unlock()
	mdl.queryFeedPrices = append(mdl.queryFeedPrices, pfs...)
}

func (mdl *AggregatedBlockFeed) QueryData(blockNum int64, weth, whatToQuery string) ([]*schemas.PriceFeed, map[string]*schemas.UniPoolPrices) {
	calls, queryAbleAdapters := mdl.getRoundDataCalls(blockNum)
	poolCalls, uniTokens := mdl.priceOnUNIFetcher.getUniswapPoolCalls(blockNum, whatToQuery)
	calls = append(calls, poolCalls...)
	//
	result := core.MakeMultiCall(mdl.Client, blockNum, false, calls)
	//
	yearnFeedLen := len(queryAbleAdapters)
	v2ABI := core.GetAbi("Uniswapv2Pool")
	v3ABI := core.GetAbi("Uniswapv3Pool")
	//
	var queryFeedPrices []*schemas.PriceFeed
	uniPricesByToken := map[string]*schemas.UniPoolPrices{}
	//
	for ind, entry := range result[:yearnFeedLen] {
		pf := mdl.processRoundData(blockNum, queryAbleAdapters[ind], entry)
		queryFeedPrices = append(queryFeedPrices, pf...)
	}
	//
	for ind, entry := range result[yearnFeedLen:] {
		tokenInd := ind / 3
		callInd := ind - tokenInd*3
		token := uniTokens[tokenInd]
		tokenDetails := mdl.priceOnUNIFetcher.tokenInfos[token]
		prices := &schemas.UniPoolPrices{BlockNum: blockNum, Token: token}
		if uniPricesByToken[token] != nil {
			prices = uniPricesByToken[token]
		}
		// ignore if failed
		if !entry.Success {
			continue
		}
		uniPricesByToken[token] = prices
		switch callInd {
		case 0:
			value, err := v2ABI.Unpack("getReserves", entry.ReturnData)
			log.CheckFatal(err)
			r0 := value[0].(*big.Int)
			r1 := value[1].(*big.Int)
			uniswapv2Price := priceInWETH(token, weth, tokenDetails.Decimals, r0, r1)
			prices.PriceV2 = utils.GetFloat64Decimal(uniswapv2Price, 18)
			prices.PriceV2Success = true
		case 1:
			value, err := v3ABI.Unpack("slot0", entry.ReturnData)
			log.CheckFatal(err)
			//https://docs.uniswap.org/sdk/guides/fetching-prices#understanding-sqrtprice
			// [(slot0**2 *Token0decimals)/2**192], divide by token for getting the float price in WETH
			//
			price := univ3SlotToPriceInBase(value[0].(*big.Int), areSorted(token, weth), tokenDetails.Decimals)
			prices.PriceV3 = utils.GetFloat64Decimal(price, 18)
			prices.PriceV3Success = true
		case 2:
			value, err := v3ABI.Unpack("observe", entry.ReturnData)
			log.CheckFatal(err)
			ticks := value[0].([]*big.Int)
			// https://medium.com/blockchain-development-notes/a-guide-on-uniswap-v3-twap-oracle-2aa74a4a97c5
			// (t1-t0)/interval
			tickDiff := new(big.Int).Quo(new(big.Int).Sub(ticks[1], ticks[0]), big.NewInt(600))
			sqrtPrice := powFloat(tickDiff)
			decimal := 18 - tokenDetails.Decimals
			if decimal != 0 {
				sqrtPrice = new(big.Float).Mul(utils.GetExpFloat(decimal), sqrtPrice)
				sqrtPrice = new(big.Float).Quo(big.NewFloat(1), sqrtPrice)
			}
			twapV3Price, _ := sqrtPrice.Float64()
			prices.TwapV3 = twapV3Price
			// if sorted use resiprocal
			if tokenDetails.Symbol == "YFI" {
				prices.TwapV3 = 1 / prices.TwapV3
			}
			prices.TwapV3Success = true
		}
	}
	return queryFeedPrices, uniPricesByToken
}

func (mdl *AggregatedBlockFeed) getRoundDataCalls(blockNum int64) (calls []multicall.Multicall2Call, queryAbleAdapters []*QueryPriceFeed) {
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

func (mdl *AggregatedBlockFeed) processRoundData(blockNum int64, adapter *QueryPriceFeed, entry multicall.Multicall2Result) []*schemas.PriceFeed {
	var priceData *schemas.PriceFeed
	if entry.Success {
		isPriceInUSD := adapter.GetVersion() > 1
		priceData = parseRoundData(entry.ReturnData, isPriceInUSD, adapter.GetAddress())
		adapter.setNotified(false)
	} else {
		switch adapter.GetDetailsByKey("pfType") {
		case ds.YearnPF:
			_priceData, err := adapter.calculateYearnPFInternally(blockNum)
			log.CheckFatal(err)
			priceData = _priceData
		}
	}
	priceFeeds := []*schemas.PriceFeed{}
	for _, token := range adapter.TokensValidAtBlock(blockNum) {
		priceDataCopy := priceData.Clone()
		//
		priceDataCopy.BlockNumber = blockNum
		priceDataCopy.Token = token
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
		RoundId:      roundData.RoundId.Int64(),
		PriceBI:      (*core.BigInt)(roundData.Answer),
		Price:        utils.GetFloat64Decimal(roundData.Answer, decimals),
		IsPriceInUSD: isPriceInUSD, // for 2 and above the prices are in usd
	}
}

/////////////////////////
// UNI pools related methods
/////////////////////////

func areSorted(token, weth string) bool {
	return strings.Compare(strings.ToLower(token), strings.ToLower(weth)) == -1
}

// r1*x/(r0+x)
func priceInWETH(token, weth string, tokenDecimals int8, r0, r1 *big.Int) *big.Int {
	if !areSorted(token, weth) {
		r1, r0 = r0, r1
	}
	amountIn := utils.GetExpInt(tokenDecimals)
	nom := new(big.Int).Mul(r1, amountIn)
	denom := new(big.Int).Add(r0, amountIn)
	return new(big.Int).Quo(nom, denom)
}

func squareIt(a *big.Int) *big.Int {
	return new(big.Int).Mul(a, a)
}

// uni v3 slot to price in base
// returns price in usdc or weth, base is usdc/weth
// if base is token1: [(slot0**2 *Token0decimals)/2**192]
// if base is token0: [(2**192 *Token1decimals)/slot0**2]
func univ3SlotToPriceInBase(slot0 *big.Int, baseIsToken1 bool, decimals int8) (price *big.Int) {
	normalizeFactor := new(big.Int).Exp(big.NewInt(2), big.NewInt(96*2), nil)
	//
	sqSlot0 := squareIt(slot0)
	if baseIsToken1 {
		price = utils.GetInt64(sqSlot0, -1*decimals)
		price = new(big.Int).Quo(price, normalizeFactor)
	} else {
		price = utils.GetInt64(normalizeFactor, -1*decimals)
		price = new(big.Int).Quo(price, sqSlot0)
	}
	return price
}
