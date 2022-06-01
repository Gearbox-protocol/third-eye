package aggregated_block_feed

import (
	"math/big"
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
	"sort"
)

func (mdl *AggregatedBlockFeed) Query(queryTill int64) {
	if len(mdl.UniPoolByToken) == 0 && len(mdl.YearnFeeds) == 0 {
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
		ch <- 1
		wg.Add(1)
		go mdl.queryAsync(blockNum, ch, wg)
		if rounds%100 == 0 {
			timeLeft := (time.Now().Sub(loopStartTime).Seconds() * float64(queryTill-blockNum)) /
				float64(blockNum-mdl.GetLastSync())
			timeLeft /= 60
			log.Infof("Synced %d in %d rounds(%fs): TimeLeft %f mins", blockNum, rounds, time.Now().Sub(roundStartTime).Seconds(), timeLeft)
			roundStartTime = time.Now()
		}
		rounds++
	}
	wg.Wait()

	for _, adapter := range mdl.YearnFeeds {
		// yearn price feed can't be disabled from v2
		if queryTill <= adapter.GetLastSync() || adapter.IsDisabled() {
			continue
		}
		adapter.AfterSyncHook(queryTill)
	}
	for _, prices := range mdl.UniPricesByTokens {
		sort.Sort(prices)
	}
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
	pfs, pricesByToken := mdl.QueryData(blockNum, weth, "all")
	for _, pf := range pfs {
		mdl.Repo.AddPriceFeed(pf)
	}
	mdl.updatePrice(pricesByToken)
	<-ch
	wg.Done()
}

func (mdl *AggregatedBlockFeed) QueryData(blockNum int64, weth, whatToQuery string) ([]*schemas.PriceFeed, map[string]*schemas.UniPoolPrices) {
	calls, queryAbleAdapters := mdl.getRoundDataCalls(blockNum)
	poolCalls, uniTokens := mdl.getUniswapPoolCalls(blockNum, whatToQuery)
	calls = append(calls, poolCalls...)
	//
	result := core.MakeMultiCall(mdl.Client, blockNum, false, calls)
	//
	yearnFeedLen := len(queryAbleAdapters)
	v2ABI := schemas.GetAbi("Uniswapv2Pool")
	v3ABI := schemas.GetAbi("Uniswapv3Pool")
	var pfs []*schemas.PriceFeed
	pricesByToken := map[string]*schemas.UniPoolPrices{}
	for i, entry := range result {
		if i < yearnFeedLen {
			pf := mdl.processPriceData(blockNum, queryAbleAdapters[i], entry)
			pfs = append(pfs, pf...)
		} else {
			tokenInd := (i - yearnFeedLen) / 3
			callInd := i - yearnFeedLen - tokenInd*3
			token := uniTokens[tokenInd]
			tokenDetails := mdl.tokenInfos[token]
			prices := &schemas.UniPoolPrices{BlockNum: blockNum}
			if pricesByToken[token] != nil {
				prices = pricesByToken[token]
			}
			// ignore if failed
			if !entry.Success {
				continue
			}
			pricesByToken[token] = prices
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
				price := utils.GetInt64(squareIt(value[0].(*big.Int)), -tokenDetails.Decimals)
				normalizeFactor := new(big.Int).Exp(big.NewInt(2), big.NewInt(96*2), nil)
				price = new(big.Int).Quo(price, normalizeFactor)
				prices.PriceV3 = utils.GetFloat64Decimal(price, 18)
				// if not sorted use resiprocal
				if !areSorted(token, weth) {
					prices.PriceV3 = 1 / prices.PriceV3
				}
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
	}
	return pfs, pricesByToken
}

func (mdl *AggregatedBlockFeed) updatePrice(pricesByToken map[string]*schemas.UniPoolPrices) {
	mdl.mu.Lock()
	defer mdl.mu.Unlock()
	for token, prices := range pricesByToken {
		prices.Token = token
		mdl.Repo.AddUniswapPrices(prices)
		mdl.UniPricesByTokens[token] = append(mdl.UniPricesByTokens[token], prices)
	}
}

func areSorted(token, weth string) bool {
	return strings.Compare(strings.ToLower(token), strings.ToLower(weth)) == -1
}
func priceInWETH(token, weth string, tokenDecimals int8, r0, r1 *big.Int) *big.Int {
	if !areSorted(token, weth) {
		tmp := r1
		r1 = r0
		r0 = tmp
	}
	amountIn := utils.GetExpInt(tokenDecimals)
	nom := new(big.Int).Mul(r1, amountIn)
	denom := new(big.Int).Add(r0, amountIn)
	return new(big.Int).Quo(nom, denom)
}

func squareIt(a *big.Int) *big.Int {
	return new(big.Int).Mul(a, a)
}
func (mdl *AggregatedBlockFeed) getUniswapPoolCalls(blockNum int64, whatToQuery string) (calls []multicall.Multicall2Call, tokens []string) {
	v2ABI := schemas.GetAbi("Uniswapv2Pool")
	v3ABI := schemas.GetAbi("Uniswapv3Pool")
	for token, pools := range mdl.UniPoolByToken {
		if whatToQuery != "all" && whatToQuery != token {
			continue
		}
		// only sync uniswap pool price for token that have last sync
		if mdl.TokenLastSync[token] >= blockNum {
			continue
		}
		uniswapv2Price, err := v2ABI.Pack("getReserves")
		log.CheckFatal(err)
		calls = append(calls, multicall.Multicall2Call{
			Target:   common.HexToAddress(pools.V2),
			CallData: uniswapv2Price,
		})
		uniswapv3Price, err := v3ABI.Pack("slot0")
		log.CheckFatal(err)
		calls = append(calls, multicall.Multicall2Call{
			Target:   common.HexToAddress(pools.V3),
			CallData: uniswapv3Price,
		})
		uniswapv3Twap, err := v3ABI.Pack("observe", []uint32{0, 600})
		log.CheckFatal(err)
		calls = append(calls, multicall.Multicall2Call{
			Target:   common.HexToAddress(pools.V3),
			CallData: uniswapv3Twap,
		})
		tokens = append(tokens, token)
	}
	return
}

func (mdl *AggregatedBlockFeed) getRoundDataCalls(blockNum int64) (calls []multicall.Multicall2Call, queryAbleAdapters []*QueryPriceFeed) {
	priceFeedABI := schemas.GetAbi("PriceFeed")
	//
	for _, adapter := range mdl.YearnFeeds {
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

func (mdl *AggregatedBlockFeed) processPriceData(blockNum int64, adapter *QueryPriceFeed, entry multicall.Multicall2Result) []*schemas.PriceFeed {
	priceFeedABI := schemas.GetAbi("PriceFeed")
	var priceData *schemas.PriceFeed
	if entry.Success {
		roundData := schemas.LatestRounData{}
		value, err := priceFeedABI.Unpack("latestRoundData", entry.ReturnData)
		log.CheckFatal(err)
		roundData.RoundId = *abi.ConvertType(value[0], new(*big.Int)).(**big.Int)
		roundData.Answer = *abi.ConvertType(value[1], new(*big.Int)).(**big.Int)
		roundData.StartedAt = *abi.ConvertType(value[2], new(*big.Int)).(**big.Int)
		roundData.UpdatedAt = *abi.ConvertType(value[3], new(*big.Int)).(**big.Int)
		roundData.AnsweredInRound = *abi.ConvertType(value[4], new(*big.Int)).(**big.Int)
		isPriceInUSD := adapter.GetVersion() > 1
		var decimals int8 = 18 // for eth
		if isPriceInUSD {
			decimals = 8 // for usd
		}
		priceData = &schemas.PriceFeed{
			RoundId:      roundData.RoundId.Int64(),
			PriceBI:      (*core.BigInt)(roundData.Answer),
			Price:        utils.GetFloat64Decimal(roundData.Answer, decimals),
			IsPriceInUSD: isPriceInUSD, // for 2 and above the prices are in usd
		}
		adapter.setNotified(false)
	} else {
		switch adapter.GetDetailsByKey("pfType") {
		case ds.YearnPF:
			priceData = adapter.calculateYearnPFInternally(blockNum)
		}
	}
	priceFeeds := []*schemas.PriceFeed{}
	for _, token := range adapter.TokensValidAtBlock(blockNum) {
		priceDataCopy := priceData.Clone()
		priceDataCopy.BlockNumber = blockNum
		priceDataCopy.Token = token
		priceDataCopy.Feed = adapter.GetAddress()
		priceFeeds = append(priceFeeds, priceDataCopy)
	}
	return priceFeeds
}

func (mdl *AggregatedBlockFeed) Clear() {
	mdl.UniPricesByTokens = map[string]schemas.SortedUniPoolPrices{}
}

// called on next level in the adapter kit
// so mu is not required as write operation is not performed at that levelAggre
func (mdl *AggregatedBlockFeed) GetUniPricesByToken(token string) []*schemas.UniPoolPrices {
	return mdl.UniPricesByTokens[token]
}
