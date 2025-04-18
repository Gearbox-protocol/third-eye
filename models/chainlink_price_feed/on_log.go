package chainlink_price_feed

import (
	"math/big"
	"strconv"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/core/types"
)

func (mdl *ChainlinkPriceFeed) OnLog(txLog types.Log) {
}
func (mdl *ChainlinkPriceFeed) OnLogs(txLogs []types.Log) {
	var blockNums []int64
	upperLimit := mdl.upperLimit()
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
			tokens := mdl.Repo.TokensValidAtBlock(mdl.Address, blockNum)
			if len(tokens) == 0 {
				return
			}
			// has atleast one valid token.
			if tokens[0].Token == "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" &&
				mdl.Address == "0x37bC7498f4FF12C19678ee8fE19d713b87F6a9e6" && blockNum > 17217055 { // as there is already another chainlink activated 0xE62B71cf983019BFf55bC83B48601ce8419650CC
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
			// for bounded oracle, if answerBI is more than upperLimit, set answer to upperLimit
			if upperLimit.Cmp(new(big.Int)) != 0 && answerBI.Cmp(upperLimit) > 0 {
				answerBI = upperLimit
			}
			// new(big.Int).SetString(txLog.Data[2:], 16)
			priceFeed = &schemas.PriceFeed{
				BlockNumber: blockNum,
				Feed:        mdl.Address,
				RoundId:     roundId,
				PriceBI:     (*core.BigInt)(answerBI),
				Price:       utils.GetFloat64Decimal(answerBI, mdl.GetVersion().Decimals()),
			}
			mdl.pfs = append(mdl.pfs, priceFeed)
			// mdl.Repo.AddPriceFeed(priceFeed)
			blockNums = append(blockNums, blockNum)
		}
	}
	// not supported for v1
	if !mdl.GetVersion().IsGBv1() {
		if len(blockNums) != 0 {
			for token := range mdl.Repo.TokenAddrsValidAtBlock(mdl.Address, blockNums[len(blockNums)-1]) {
				mdl.Repo.ChainlinkPriceUpdatedAt(token, blockNums)
			}
		}
	}

}

// func (mdl *ChainlinkPriceFeed) AddToken(token string, blockNum int64, pfVersion schemas.PriceOracleT) {
// 		//
// 		mdl.mergedPFManager.AddToken(token, blockNum, priceOracle)
// 	data, err :=mdl.MainAgg.contractETH.LatestRoundData(&bind.CallOpts{
// 		BlockNumber: new(big.Int).SetInt64(blockNum),
// 	})
// 	log.CheckFatal(err)
// 	priceFeed := &schemas.PriceFeed{
// 		BlockNumber:     blockNum,
// 		Token:           token,
// 		Feed:            mdl.Address,
// 		RoundId:         data.RoundId.Int64(),
// 		PriceBI:         (*core.BigInt)(data.Answer),
// 		Price:           utils.GetFloat64Decimal(data.Answer, pfVersion.Decimals()),
// 		MergedPFVersion: mdl.mergedPFManager.GetMergedPFVersion(token, blockNum, mdl.Address),
// 	}
// 	mdl.Repo.AddPriceFeed(priceFeed)
// }

// func (mdl ChainlinkPriceFeed) DisableToken(token string, blockNum int64, pfVersion schemas.PriceOracleT) {
// 	mdl.mergedPFManager.DisableToken(blockNum, token, w)
// 	// final := mdl.mergedPFManager.GetMergedPFVersion(token, blockNum, mdl.Address)
// 	// if final == 0 {
// 	// 	mdl.SetBlockToDisableOn(blockNum)
// 	// }
// }
