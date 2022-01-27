package aggregated_block_feed

import (
	"github.com/Gearbox-protocol/third-eye/artifacts/multicall"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
	"time"
	// "fmt"
)

const interval = 25

func (mdl *AggregatedBlockFeed) Query(queryTill int64) {
	queryFrom := mdl.GetLastSync() + interval
	log.Infof("Sync %s from %d to %d", mdl.GetName(), queryFrom, queryTill)
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

func pow(a *big.Int) *big.Int {
	f := big.NewFloat(1.0001)
	ans := big.NewFloat(1)
	for i := 0; i < a.BitLen(); i++ {
		if a.Bit(i) == 1 {
			ans = new(big.Float).Mul(f, ans)
		}
		f = new(big.Float).Mul(f, f)
	}
	integer := new(big.Int)
	f.Int(integer)
	return integer
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
	pricesByToken := map[string]*core.PoolPrices{}
	weth := mdl.Repo.GetWETHAddr()
	for i, entry := range result {
		if i < yearnFeedLen {
			mdl.processPriceData(blockNum, queryAbleAdapters[i], entry)
		} else {
			tokenInd := (i - yearnFeedLen) / 3
			callInd := i - yearnFeedLen - tokenInd*3
			token := uniTokens[tokenInd]
			pools := mdl.UniPoolByToken[token]
			prices := &core.PoolPrices{BlockNum: blockNum}
			if pricesByToken[token] != nil {
				prices = pricesByToken[token]
			}
			pricesByToken[token] = prices
			switch callInd {
			case 0:
				value, err := v2ABI.Unpack("getReserves", entry.ReturnData)
				log.CheckFatal(err)
				r0 := value[0].(*big.Int)
				r1 := value[1].(*big.Int)
				uniswapv2Price := priceInWETH(token, weth, pools.Decimals, r0, r1)
				prices.PriceV2 = utils.GetFloat64Decimal(uniswapv2Price, 18)
			case 1:
				value, err := v3ABI.Unpack("slot0", entry.ReturnData)
				log.CheckFatal(err)
				uniswapv3Price := squareIt(value[0].(*big.Int))
				prices.PriceV3 = utils.GetFloat64Decimal(uniswapv3Price, 18)
			case 2:
				value, err := v3ABI.Unpack("observe", entry.ReturnData)
				log.CheckFatal(err)
				ticks := value[0].([]*big.Int)
				tickDiff := new(big.Int).Sub(ticks[1], ticks[0])
				sqrtPrice := pow(tickDiff)
				twapV3 := squareIt(sqrtPrice)
				prices.TwapV3 = utils.GetFloat64Decimal(twapV3, 18)
			}
		}
	}
	for token, prices := range pricesByToken {
		mdl.UniPricesByTokens[token] = append(mdl.UniPricesByTokens[token], prices)
	}
}

func priceInWETH(token, weth string, tokenDecimals int8, r0, r1 *big.Int) *big.Int {
	if strings.Compare(token, weth) == 1 {
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
		uniswapv3Twap, err := v3ABI.Pack("observe", []int32{0, 600})
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
	mdl.UniPricesByTokens = map[string][]*core.PoolPrices{}
}

func (mdl *AggregatedBlockFeed) GetUniPricesByToken(token string) []*core.PoolPrices {
	return mdl.UniPricesByTokens[token]
}
