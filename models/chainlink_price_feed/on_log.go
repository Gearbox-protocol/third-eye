package chainlink_price_feed

import (
	"math/big"
	"strconv"

	"math"

	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/utils"
	"github.com/ethereum/go-ethereum/core/types"
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
			log.Info(uniPricesInd, blockNum)
			for uniPricesInd < len(uniPrices) && blockNum > uniPrices[uniPricesInd].BlockNum {
				mdl.compareDiff(prevPriceFeed, uniPrices[uniPricesInd])
				log.Info(uniPrices[uniPricesInd].BlockNum, blockNum)
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
			prevPriceFeed = priceFeed
		}
	}
	// remaining prices are filled
	for uniPricesInd < len(uniPrices) && uniPrices[uniPricesInd].BlockNum < mdl.GetBlockToDisableOn() {
		mdl.compareDiff(prevPriceFeed, uniPrices[uniPricesInd])
		uniPricesInd++
	}
}

func (mdl *ChainlinkPriceFeed) compareDiff(pf *core.PriceFeed, uniPoolPrices *core.UniPoolPrices) {
	// previous pricefeed can be nil
	if pf == nil {
		return
	}
	// set the token and blocknumber of chainlink
	uniPoolPrices.ChainlinkBlockNumber = pf.BlockNumber
	uniPoolPrices.Token = pf.Token
	mdl.Repo.AddUniswapPrices(uniPoolPrices)
	if !mdl.isNotified() &&
		((uniPoolPrices.PriceV2Success && greaterFluctuation(uniPoolPrices.PriceV2, pf.PriceETH)) ||
			(uniPoolPrices.PriceV3Success && greaterFluctuation(uniPoolPrices.PriceV3, pf.PriceETH)) ||
			(uniPoolPrices.TwapV3Success && greaterFluctuation(uniPoolPrices.TwapV3, pf.PriceETH))) {
		mdl.uniPriceVariationNotify(pf, uniPoolPrices)
		mdl.Details["notified"] = true
	} else {
		mdl.Details["notified"] = false
	}
}

func greaterFluctuation(a, b float64) bool {
	return math.Abs((a-b)/a) > 0.01
}

func (mdl *ChainlinkPriceFeed) uniPriceVariationNotify(pf *core.PriceFeed, uniPrices *core.UniPoolPrices) {
	symbol := mdl.Repo.GetToken(mdl.Token).Symbol
	log.Infof(`Token:%s(%s) =>
	Chainlink BlockNum:%d %f
	Uni price at block number: %d
	Uniswapv2 Price: %f
	Uniswapv3 Price: %f
	Uniswapv3 Twap: %f\n`, symbol, mdl.Token,
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
