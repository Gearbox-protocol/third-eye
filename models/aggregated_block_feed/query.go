package aggregated_block_feed

import (
	"math/big"
	"strings"
	"time"

	"github.com/Gearbox-protocol/third-eye/artifacts/multicall"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	// "fmt"
)

const interval = 25

func (mdl *AggregatedBlockFeed) Query(queryTill int64) {
	if len(mdl.UniPoolByToken) == 0 && len(mdl.YearnFeeds) == 0 {
		return
	}
	// msg
	queryFrom := mdl.GetLastSync() + interval
	log.Infof("Sync %s from %d to %d", mdl.GetName(), queryFrom, queryTill)
	// timer with query of block
	rounds := 0
	loopStartTime := time.Now()
	roundStartTime := time.Now()
	for blockNum := queryFrom; blockNum <= queryTill; blockNum += interval {
		mdl.query(blockNum)
		if rounds%100 == 0 {
			timeLeft := (time.Now().Sub(loopStartTime).Seconds() * float64(queryTill-blockNum)) /
				float64(blockNum-mdl.GetLastSync())
			timeLeft /= 60
			log.Infof("Synced %d in %d rounds(%fs): TimeLeft %f mins", blockNum, rounds, time.Now().Sub(roundStartTime).Seconds(), timeLeft)
			roundStartTime = time.Now()
		}
		rounds++
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

func (mdl *AggregatedBlockFeed) query(blockNum int64) {
	mdl.Repo.SetBlock(blockNum)
	calls, queryAbleAdapters := mdl.getRoundDataCalls(blockNum)
	poolCalls, uniTokens := mdl.getUniswapPoolCalls(blockNum)
	calls = append(calls, poolCalls...)
	//
	result := mdl.Repo.MakeMultiCall(blockNum, false, calls)
	//
	yearnFeedLen := len(queryAbleAdapters)
	v2ABI := core.GetAbi("Uniswapv2Pool")
	v3ABI := core.GetAbi("Uniswapv3Pool")
	pricesByToken := map[string]*core.UniPoolPrices{}
	weth := mdl.Repo.GetWETHAddr()
	for i, entry := range result {
		if i < yearnFeedLen {
			mdl.processPriceData(blockNum, queryAbleAdapters[i], entry)
		} else {
			tokenInd := (i - yearnFeedLen) / 3
			callInd := i - yearnFeedLen - tokenInd*3
			token := uniTokens[tokenInd]
			tokenDetails := mdl.Repo.GetToken(token)
			prices := &core.UniPoolPrices{BlockNum: blockNum}
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
	for token, prices := range pricesByToken {
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
func (mdl *AggregatedBlockFeed) getUniswapPoolCalls(blockNum int64) (calls []multicall.Multicall2Call, tokens []string) {
	v2ABI := core.GetAbi("Uniswapv2Pool")
	v3ABI := core.GetAbi("Uniswapv3Pool")
	for token, pools := range mdl.UniPoolByToken {
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

func (mdl *AggregatedBlockFeed) getRoundDataCalls(blockNum int64) (calls []multicall.Multicall2Call, queryAbleAdapters []*YearnPriceFeed) {
	priceFeedABI := core.GetAbi(core.YearnPriceFeed)
	//
	for _, adapter := range mdl.YearnFeeds {
		if blockNum <= adapter.GetLastSync() || adapter.IsDisabled() {
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
		adapter.AfterSyncHook(blockNum)
	}
	return
}

func (mdl *AggregatedBlockFeed) processPriceData(blockNum int64, adapter *YearnPriceFeed, entry multicall.Multicall2Result) {
	priceFeedABI := core.GetAbi(core.YearnPriceFeed)
	var priceData *core.PriceFeed
	if entry.Success {
		roundData := core.LatestRounData{}
		value, err := priceFeedABI.Unpack("latestRoundData", entry.ReturnData)
		log.CheckFatal(err)
		roundData.RoundId = *abi.ConvertType(value[0], new(*big.Int)).(**big.Int)
		roundData.Answer = *abi.ConvertType(value[1], new(*big.Int)).(**big.Int)
		roundData.StartedAt = *abi.ConvertType(value[2], new(*big.Int)).(**big.Int)
		roundData.UpdatedAt = *abi.ConvertType(value[3], new(*big.Int)).(**big.Int)
		roundData.AnsweredInRound = *abi.ConvertType(value[4], new(*big.Int)).(**big.Int)
		priceData = &core.PriceFeed{
			RoundId:    roundData.RoundId.Int64(),
			PriceETHBI: (*core.BigInt)(roundData.Answer),
			PriceETH:   utils.GetFloat64Decimal(roundData.Answer, 18),
		}
		adapter.Details["notified"] = false
	} else {
		priceData = adapter.calculatePriceFeedInternally(blockNum)
	}
	tokenAddr, ok := adapter.Details["token"].(string)
	if !ok {
		log.Fatal("Failing in asserting to string: %s", mdl.Details["token"])
	}
	priceData.BlockNumber = blockNum
	priceData.Token = tokenAddr
	priceData.Feed = mdl.GetAddress()
	mdl.Repo.AddPriceFeed(blockNum, priceData)
}

func (mdl *AggregatedBlockFeed) Clear() {
	mdl.UniPricesByTokens = map[string][]*core.UniPoolPrices{}
}

func (mdl *AggregatedBlockFeed) GetUniPricesByToken(token string) []*core.UniPoolPrices {
	return mdl.UniPricesByTokens[token]
}
