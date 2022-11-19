package composite_chainlink

import (
	"math"
	"math/big"
	"strconv"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	cpf "github.com/Gearbox-protocol/third-eye/models/chainlink_price_feed"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (mdl *CompositeChainlinkPF) breakPoint(tokenType string, mainAgg *cpf.ChainlinkMainAgg) (int64, bool) {
	newPriceFeed, newPhaseId := mainAgg.GetPriceFeedAddr(mdl.WillSyncTill)
	previousPF := mdl.getAddrFromDetails(tokenType)
	if newPriceFeed != previousPF && newPriceFeed != core.NULL_ADDR {
		var discoveredAt int64
		if newPhaseId != -1 {
			discoveredAt = mainAgg.GetFeedUpdateBlockUsingPhaseId(uint16(newPhaseId), mdl.LastSync+1, mdl.WillSyncTill)
		} else {
			discoveredAt = mainAgg.GetFeedUpdateBlockAggregator(newPriceFeed, mdl.LastSync+1, mdl.WillSyncTill)
		}
		return discoveredAt, true
	}
	return 0, false
}

func (mdl *CompositeChainlinkPF) OnLogs(txLogs []types.Log) {
	var breakPoint int64 = math.MaxInt64
	if bp, valid := mdl.breakPoint("target", mdl.MainAgg); valid {
		if bp < breakPoint {
			breakPoint = bp
		}
	}
	if bp, valid := mdl.breakPoint("base", mdl.BaseTokenMainAgg); valid {
		if bp < breakPoint {
			breakPoint = bp
		}
	}
	targetETHPF := mdl.getAddrFromDetails("target")
	lastInds := map[common.Address]int{}
	for txLogInd, txLog := range txLogs {
		blockNum := int64(txLog.BlockNumber)
		if breakPoint <= blockNum {
			break
		}
		switch txLog.Topics[0] {
		case core.Topic("AnswerUpdated(int256,uint256,uint256)"):
			// there might be 2 AnswerUpdated events for same block, use the last one
			// example
			// https://goerli.etherscan.io/tx/0x03308a0b6f024e6c35a92e7708ab5a72322f733d22427d51624862d82ca1983a
			// https://goerli.etherscan.io/tx/0x38e5551ae639d22554072ba1a53e026a0858c2cfedcedb83e5cc63bb1c8b8ea8
			// on mainnet
			// https://etherscan.io/tx/0xb3aaa84cac23a30ab20cbd254b2297840f23057faf1f05e7655304be6cffc19e#eventlog
			// https://etherscan.io/tx/0x3112f0a42f288ca56a2c8f8003355ad20e87e1f23c3ffa991633f6bb25eb8c58#eventlog
			lastInd, exists := lastInds[txLog.Address]
			if exists && int64(txLogs[lastInd].BlockNumber) == blockNum {
				continue
			}
			lastInds[txLog.Address] = txLogInd
			//
			roundId, err := strconv.ParseInt(txLog.Topics[2].Hex()[50:], 16, 64)
			if err != nil {
				log.Fatal("TxHash", txLog.TxHash.Hex(), "roundid failed", txLog.Topics[2].Hex())
			}
			answerBI, ok := new(big.Int).SetString(txLog.Topics[1].Hex()[2:], 16)
			if !ok {
				log.Fatal("answer parsing failed", txLog.Topics[1].Hex())
			}
			if txLog.Address == targetETHPF {
				mdl.TokenETHPrice = answerBI
			} else {
				mdl.ETHUSDPrice = answerBI
			}
			answerBI = utils.GetInt64(
				new(big.Int).Mul(mdl.TokenETHPrice, mdl.ETHUSDPrice),
				18,
			)
			// only usd price feed
			priceFeed := &schemas.PriceFeed{
				BlockNumber:  blockNum,
				Token:        mdl.Token,
				Feed:         mdl.Address,
				RoundId:      roundId,
				PriceBI:      (*core.BigInt)(answerBI),
				Price:        utils.GetFloat64Decimal(answerBI, 8),
				IsPriceInUSD: true,
			}
			mdl.Repo.AddPriceFeed(priceFeed)
		}
	}
	if breakPoint != math.MaxInt64 {
		oracleAddr := common.HexToAddress(mdl.GetDetailsByKey("oracle"))
		targetPF := getAddrFromRPC(mdl.Client, "targetETH", oracleAddr, breakPoint)
		basePF := getAddrFromRPC(mdl.Client, "ETHUSD", oracleAddr, breakPoint)
		txLogs, err := core.Node{Client: mdl.Client}.GetLogs(
			breakPoint,
			mdl.WillSyncTill,
			[]common.Address{targetPF, basePF},
			[][]common.Hash{
				{core.Topic("AnswerUpdated(int256,uint256,uint256)")},
			},
		)
		log.CheckFatal(err)
		mdl.OnLogs(txLogs)
	}
}

func (mdl *CompositeChainlinkPF) OnLog(types.Log) {

}
