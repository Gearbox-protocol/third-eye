package chainlink_price_feed

import (
	"math/big"
	"strconv"

	"math"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/core/types"
)

func (mdl *ChainlinkPriceFeed) OnLog(txLog types.Log) {
}
func (mdl *ChainlinkPriceFeed) OnLogs(txLogs []types.Log) {
	uniPrices := mdl.Repo.GetUniPricesByToken(mdl.Token)
	uniPricesInd := 0
	// filter uniswapool prices that were before the lastsync of contract
	startFrom := mdl.GetLastSync()
	if mdl.prevPriceFeed != nil {
		startFrom = utils.Min(startFrom, mdl.prevPriceFeed.BlockNumber)
	}
	for uniPricesInd < len(uniPrices) && startFrom >= uniPrices[uniPricesInd].BlockNum {
		uniPricesInd++
	}
	for _, txLog := range txLogs {
		var priceFeed *schemas.PriceFeed
		blockNum := int64(txLog.BlockNumber)
		switch txLog.Topics[0] {
		case core.Topic("AnswerUpdated(int256,uint256,uint256)"):
			roundId, err := strconv.ParseInt(txLog.Topics[2].Hex()[2:], 16, 64)
			if err != nil {
				log.Fatal("roundid failed")
			}

			answerBI, ok := new(big.Int).SetString(txLog.Topics[1].Hex()[2:], 16)
			if !ok {
				log.Fatal("answer parsing failed")
			}
			isPriceInUSD := mdl.GetVersion() > 1
			var decimals int8 = 18 // for eth
			if isPriceInUSD {
				decimals = 8 // for usd
			}
			// new(big.Int).SetString(txLog.Data[2:], 16)
			priceFeed = &schemas.PriceFeed{
				BlockNumber:  blockNum,
				Token:        mdl.Token,
				Feed:         mdl.Address,
				RoundId:      roundId,
				PriceBI:      (*core.BigInt)(answerBI),
				Price:        utils.GetFloat64Decimal(answerBI, decimals),
				IsPriceInUSD: isPriceInUSD,
			}
			for uniPricesInd < len(uniPrices) && blockNum > uniPrices[uniPricesInd].BlockNum {
				mdl.compareDiff(mdl.prevPriceFeed, uniPrices[uniPricesInd])
				uniPricesInd++
			}
			if len(uniPrices) > 0 {
				lastValidUniBlock := uniPricesInd
				// if len overflow set value-1
				if lastValidUniBlock == len(uniPrices) ||
					// if the next blockNum for uni is not equal to txLog blockNum
					!(lastValidUniBlock < len(uniPrices) && blockNum == uniPrices[lastValidUniBlock].BlockNum) {
					lastValidUniBlock = lastValidUniBlock - 1
				}
				if lastValidUniBlock >= 0 {
					uniPoolPrice := uniPrices[lastValidUniBlock]
					priceFeed.Uniswapv2Price = uniPoolPrice.PriceV2
					priceFeed.Uniswapv3Price = uniPoolPrice.PriceV3
					priceFeed.Uniswapv3Twap = uniPoolPrice.TwapV3
					priceFeed.UniPriceFetchBlock = uniPoolPrice.BlockNum
				}
			}
			mdl.Repo.AddPriceFeed(blockNum, priceFeed)
			mdl.prevPriceFeed = priceFeed
		}
	}
	// remaining prices are filled
	for uniPricesInd < len(uniPrices) && uniPrices[uniPricesInd].BlockNum < mdl.GetBlockToDisableOn() {
		mdl.compareDiff(mdl.prevPriceFeed, uniPrices[uniPricesInd])
		uniPricesInd++
	}
}

func (mdl *ChainlinkPriceFeed) compareDiff(pf *schemas.PriceFeed, uniPoolPrices *schemas.UniPoolPrices) {
	// previous pricefeed can be nil
	if pf == nil {
		return
	}
	mdl.Repo.AddUniPriceAndChainlinkRelation(&schemas.UniPriceAndChainlink{
		UniBlockNum:          uniPoolPrices.BlockNum,
		ChainlinkBlockNumber: pf.BlockNumber,
		Token:                pf.Token,
		Feed:                 pf.Feed,
	})
	// For usd
	if mdl.GetVersion() <= 1 && (uniPoolPrices.PriceV2Success && greaterFluctuation(uniPoolPrices.PriceV2, pf.Price)) ||
		(uniPoolPrices.PriceV3Success && greaterFluctuation(uniPoolPrices.PriceV3, pf.Price)) ||
		(uniPoolPrices.TwapV3Success && greaterFluctuation(uniPoolPrices.TwapV3, pf.Price)) {
		if !mdl.isNotified() {
			mdl.uniPriceVariationNotify(pf, uniPoolPrices)
			mdl.Details["notified"] = true
		}
	} else {
		mdl.Details["notified"] = false
	}
}

func greaterFluctuation(a, b float64) bool {
	return math.Abs((a-b)/a) > 0.03
}

func (mdl *ChainlinkPriceFeed) uniPriceVariationNotify(pf *schemas.PriceFeed, uniPrices *schemas.UniPoolPrices) {
	symbol := mdl.Repo.GetToken(mdl.Token).Symbol
	log.Infof(`Token:%s(%s) =>
	Chainlink BlockNum:%d %f
	Uni pool prices are at block: %d
	Uniswapv2 Price: %f
	Uniswapv3 Price: %f
	Uniswapv3 Twap: %f\n`, symbol, mdl.Token,
		pf.BlockNumber, pf.Price,
		uniPrices.BlockNum, uniPrices.PriceV2, uniPrices.PriceV3, uniPrices.TwapV3)
}

func (mdl *ChainlinkPriceFeed) isNotified() bool {
	if mdl.Details == nil || mdl.Details["notified"] == nil {
		return false
	}
	value, ok := mdl.Details["notified"].(bool)
	if !ok {
		log.Fatal("Notified not parsed")
	}
	return value
}

func (mdl *ChainlinkPriceFeed) SetUnderlyingState(obj interface{}) {
	switch obj.(type) {
	case (*schemas.PriceFeed):
		state := obj.(*schemas.PriceFeed)
		mdl.prevPriceFeed = state
	default:
		log.Fatal("Type assertion for chainlink state failed")
	}
}
