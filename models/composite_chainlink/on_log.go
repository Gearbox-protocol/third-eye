package composite_chainlink

import (
	"math"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg"
	"github.com/Gearbox-protocol/sdk-go/utils"
	cpf "github.com/Gearbox-protocol/third-eye/models/chainlink_price_feed"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (mdl *CompositeChainlinkPF) breakPoint(tokenType string, mainAgg *cpf.ChainlinkMainAgg) (common.Address, int64) {
	newPhaseAgg, newPhaseId := mainAgg.GetPriceFeedAddr(mdl.WillSyncTill)
	previousPhaseAgg := mdl.getAddrFromDetails(tokenType)
	if previousPhaseAgg != newPhaseAgg && newPhaseAgg != core.NULL_ADDR {
		var discoveredAt int64
		if newPhaseId != -1 {
			discoveredAt = mainAgg.GetFeedUpdateBlockUsingPhaseId(uint16(newPhaseId), mdl.LastSync+1, mdl.WillSyncTill)
		} else {
			discoveredAt = mainAgg.GetFeedUpdateBlockAggregator(newPhaseAgg, mdl.LastSync+1, mdl.WillSyncTill)
		}
		return newPhaseAgg, discoveredAt
	}
	return newPhaseAgg, math.MaxInt64
}

func (mdl *CompositeChainlinkPF) OnLogs(txLogs []types.Log) {
	var breakPoint int64 = math.MaxInt64
	//
	newMainPhaseAgg, bpOne := mdl.breakPoint("targetPhase", mdl.MainAgg)
	if bpOne < breakPoint {
		breakPoint = bpOne
	}
	newBasePhaseAgg, bpTwo := mdl.breakPoint("basePhase", mdl.BaseTokenMainAgg)
	if bpTwo < breakPoint {
		breakPoint = bpTwo
	}
	//
	targetETHPF := mdl.getAddrFromDetails("targetPhase")
	for txLogInd, txLog := range txLogs {
		blockNum := int64(txLog.BlockNumber)
		if breakPoint <= blockNum {
			break
		}
		switch txLog.Topics[0] {
		case core.Topic("AnswerUpdated(int256,uint256,uint256)"):
			// roundId, err := strconv.ParseInt(txLog.Topics[2].Hex()[50:], 16, 64)
			// if err != nil {
			// 	log.Fatal("TxHash", txLog.TxHash.Hex(), "roundid failed", txLog.Topics[2].Hex())
			// }
			answerBI, ok := new(big.Int).SetString(txLog.Topics[1].Hex()[2:], 16)
			if !ok {
				log.Fatal("answer parsing failed", txLog.Topics[1].Hex())
			}
			if txLog.Address == targetETHPF {
				mdl.TokenETHPrice = answerBI
			} else {
				mdl.ETHUSDPrice = answerBI
			}
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
			mdl.addPriceToDB(blockNum)
			//

		}
	}
	if breakPoint != math.MaxInt64 {
		mdl.Details["secAddrs"] = map[string]interface{}{
			"target":      mdl.getAddrFromDetails("target"),
			"base":        mdl.getAddrFromDetails("base"),
			"targetPhase": newMainPhaseAgg,
			"basePhase":   newBasePhaseAgg,
		}
		mdl.setPrices(breakPoint)
		mdl.addPriceToDB(breakPoint) // H1
		//
		txLogs, err := pkg.Node{Client: mdl.Client}.GetLogs(
			breakPoint+1, // bcz price for breakPoint already added at H1
			mdl.WillSyncTill,
			[]common.Address{newMainPhaseAgg, newBasePhaseAgg},
			[][]common.Hash{
				{core.Topic("AnswerUpdated(int256,uint256,uint256)")},
			},
		)
		log.CheckFatal(err)
		mdl.OnLogs(txLogs)
	}
}

func (mdl *CompositeChainlinkPF) addPriceToDB(blockNum int64) {
	answerBI := utils.GetInt64(
		new(big.Int).Mul(mdl.TokenETHPrice, mdl.ETHUSDPrice),
		mdl.decimalsOfBasePF,
	)
	// only usd price feed
	priceFeed := &schemas.PriceFeed{
		BlockNumber:     blockNum,
		Token:           mdl.Token,
		Feed:            mdl.GetDetailsByKey("oracle"),
		RoundId:         0,
		PriceBI:         (*core.BigInt)(answerBI),
		Price:           utils.GetFloat64Decimal(answerBI, 8),
		MergedPFVersion: mdl.GetMergedPFVersion(),
	}
	mdl.Repo.AddPriceFeed(priceFeed)
}

func (mdl *CompositeChainlinkPF) OnLog(types.Log) {

}

func (mdl CompositeChainlinkPF) GetMergedPFVersion() schemas.MergedPFVersion {
	if mdl.Details["mergedPFVersion"] != nil {
		if v, ok := mdl.Details["mergedPFVersion"].(int8); ok {
			return schemas.MergedPFVersion(v)
		}
		if v, ok := mdl.Details["mergedPFVersion"].(float64); ok {
			return schemas.MergedPFVersion(v)
		}
		return schemas.MergedPFVersion(mdl.Details["mergedPFVersion"].(schemas.MergedPFVersion))
	}
	log.Fatal(utils.ToJson(mdl.Details))
	return schemas.MergedPFVersion(0)
}
func (mdl CompositeChainlinkPF) AddToken(token string, pfVersion schemas.PFVersion) {
	if mdl.Details["token"] != nil {
		if mdl.Details["token"].(string) != token {
			log.Fatal("stored token for chainlink is different from new added token", mdl.Details["token"].(string), token)
		}
	}
	mdl.Details["mergedPFVersion"] = mdl.GetMergedPFVersion() | schemas.MergedPFVersion(pfVersion)
}

func (mdl CompositeChainlinkPF) DisableToken(token string, blockNum int64, pfVersion schemas.PFVersion) {
	if mdl.Details["token"] != nil {
		if mdl.Details["token"].(string) != token {
			log.Fatal("stored token for chainlink is different from new added token", mdl.Details["token"].(string), token)
		}
	}
	final := mdl.GetMergedPFVersion() ^ schemas.MergedPFVersion(pfVersion)
	mdl.Details["mergedPFVersion"] = final
	if final == 0 {
		mdl.SetBlockToDisableOn(blockNum)
	}
}
