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
	var priceFeeds []*schemas.PriceFeed
	for txLogInd, txLog := range txLogs {
		var priceFeed *schemas.PriceFeed
		blockNum := int64(txLog.BlockNumber)
		switch txLog.Topics[0] {
		case core.Topic("AnswerUpdated(int256,uint256,uint256)"):
			// there might be 2 AnswerUpdated events for same block, use the last one
			// example
			// https://goerli.etherscan.io/tx/0x03308a0b6f024e6c35a92e7708ab5a72322f733d22427d51624862d82ca1983a
			// https://goerli.etherscan.io/tx/0x38e5551ae639d22554072ba1a53e026a0858c2cfedcedb83e5cc63bb1c8b8ea8
			// on mainnet
			// https://etherscan.io/tx/0xb3aaa84cac23a30ab20cbd254b2297840f23057faf1f05e7655304be6cffc19e#eventlog
			// https://etherscan.io/tx/0x3112f0a42f288ca56a2c8f8003355ad20e87e1f23c3ffa991633f6bb25eb8c58#eventlog
			if txLogInd+1 < len(txLogs) && int64(txLogs[txLogInd+1].BlockNumber) == blockNum {
				continue
			}
			//
			roundId, err := strconv.ParseInt(txLog.Topics[2].Hex()[50:], 16, 64)
			if err != nil {
				log.Fatal("TxHash", txLog.TxHash.Hex(), "roundid failed", txLog.Topics[2].Hex())
			}

			answerBI, ok := new(big.Int).SetString(txLog.Topics[1].Hex()[2:], 16)
			if !ok {
				log.Fatal("answer parsing failed", txLog.Topics[1].Hex())
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
			mdl.Repo.AddPriceFeed(priceFeed)
			priceFeeds = append(priceFeeds, priceFeed)
		}
	}

	uniPrices := mdl.Repo.GetUniPricesByToken(mdl.Token)
	uniPricesInd := 0
	// filter uniswapool prices that were before the lastsync of contract
	startFrom := mdl.GetLastSync()
	if mdl.prevPriceFeed != nil {
		startFrom = utils.Min(startFrom, mdl.prevPriceFeed.BlockNumber)
	}
	// finish events btw (lastsync or the prev chainlink price stored) and first log on chainlink
	for uniPricesInd < len(uniPrices) && startFrom >= uniPrices[uniPricesInd].BlockNum {
		uniPricesInd++
	}
	for _, priceFeed := range priceFeeds {
		blockNum := priceFeed.BlockNumber
		//
		for uniPricesInd < len(uniPrices) && blockNum > uniPrices[uniPricesInd].BlockNum {
			mdl.compareDiff(mdl.prevPriceFeed, uniPrices[uniPricesInd])
			uniPricesInd++
		}
		if len(uniPrices) > 0 {
			lastValidUniInd := uniPricesInd
			// if len overflow set value-1
			if lastValidUniInd == len(uniPrices) ||
				// if the next blockNum for uni is not equal to txLog blockNum
				!(lastValidUniInd < len(uniPrices) && blockNum == uniPrices[lastValidUniInd].BlockNum) {
				lastValidUniInd = lastValidUniInd - 1
			}
			if lastValidUniInd >= 0 {
				uniPoolPrice := uniPrices[lastValidUniInd]
				priceFeed.Uniswapv2Price = uniPoolPrice.PriceV2
				priceFeed.Uniswapv3Price = uniPoolPrice.PriceV3
				priceFeed.Uniswapv3Twap = uniPoolPrice.TwapV3
				priceFeed.UniPriceFetchBlock = uniPoolPrice.BlockNum
			}
		}
		mdl.prevPriceFeed = priceFeed
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
