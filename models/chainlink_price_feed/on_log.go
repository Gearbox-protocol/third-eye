package chainlink_price_feed

import (
	"math/big"
	"strconv"

	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/utils"
	"github.com/ethereum/go-ethereum/core/types"
	"math"
)

func (mdl *ChainlinkPriceFeed) OnLog(txLog types.Log) {
}
func (mdl *ChainlinkPriceFeed) OnLogs(txLogs []types.Log) {
	uniPrices := mdl.Repo.GetUniPricesByToken(mdl.Token)
	uniPricesInd := 0
	// filter uniswapool prices that were before the lastsync of contract
	for uniPricesInd < len(uniPrices) && mdl.GetLastSync() >= uniPrices[uniPricesInd].BlockNum {
		uniPricesInd++
	}
	var prevPriceFeed *core.PriceFeed
	for _, txLog := range txLogs {
		var priceFeed *core.PriceFeed
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
			// new(big.Int).SetString(txLog.Data[2:], 16)
			priceFeed = &core.PriceFeed{
				BlockNumber: blockNum,
				Token:       mdl.Token,
				Feed:        mdl.Address,
				RoundId:     roundId,
				PriceETHBI:  (*core.BigInt)(answerBI),
				PriceETH:    utils.GetFloat64Decimal(answerBI, 18),
			}
		}
		for uniPricesInd < len(uniPrices) && blockNum > uniPrices[uniPricesInd].BlockNum {
			mdl.compareDiff(prevPriceFeed, uniPrices[uniPricesInd])
			uniPricesInd++
		}
		lastUniPrice := uniPricesInd
		if lastUniPrice == len(uniPrices) && lastUniPrice > 0 {
			lastUniPrice = lastUniPrice - 1
		}
		if len(uniPrices) != 0 {
			uniPoolPrices := uniPrices[lastUniPrice]
			priceFeed.Uniswapv2Price = uniPoolPrices.PriceV2
			priceFeed.Uniswapv3Price = uniPoolPrices.PriceV3
			priceFeed.Uniswapv3Twap = uniPoolPrices.TwapV3
			priceFeed.UniPriceFetchBlock = uniPoolPrices.BlockNum
		}
		mdl.Repo.AddPriceFeed(blockNum, priceFeed)
		prevPriceFeed = priceFeed
	}
	// remaining prices are filled
	for uniPricesInd < len(uniPrices) {
		mdl.compareDiff(prevPriceFeed, uniPrices[uniPricesInd])
		uniPricesInd++
	}
}

func (mdl *ChainlinkPriceFeed) compareDiff(pf *core.PriceFeed, uniPrices *core.PoolPrices) {
	// previous pricefeed can be nil
	if pf == nil {
		return
	}
	if !mdl.isNotified() && 
		((uniPrices.PriceV2Success && greaterFluctuation(uniPrices.PriceV2, pf.PriceETH)) ||
		(uniPrices.PriceV3Success && greaterFluctuation(uniPrices.PriceV3, pf.PriceETH)) ||
		(uniPrices.TwapV3Success && greaterFluctuation(uniPrices.TwapV3, pf.PriceETH))) {
		mdl.uniPriceVariationNotify(pf, uniPrices)
		mdl.Details["notified"] = true
	} else {
		mdl.Details["notified"] = false
	}
}

func greaterFluctuation(a, b float64) bool {
	return math.Abs((a-b)/a) > 0.03
}

func (mdl *ChainlinkPriceFeed) uniPriceVariationNotify(pf *core.PriceFeed, uniPrices *core.PoolPrices) {
	symbol := mdl.Repo.GetToken(mdl.Token).Symbol
	log.Info(`Token:%s(%s) =>
	Chainlink BlockNum:%d %f
	Uni price at block number: %d
	Uniswapv2 Price: %f
	Uniswapv3 Price: %f
	Uniswapv3 Twap: %f`, symbol, mdl.Token,
		pf.BlockNumber, pf.PriceETH,
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
