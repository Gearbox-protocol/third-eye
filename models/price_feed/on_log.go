package price_feed

import (
	"github.com/Gearbox-protocol/gearscan/core"
	"github.com/Gearbox-protocol/gearscan/log"
	"github.com/Gearbox-protocol/gearscan/utils"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strconv"
	// "github.com/Gearbox-protocol/gearscan/models/price_feed"
)

func (mdl *PriceFeed) OnLog(txLog types.Log) {
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
		blockNum := int64(txLog.BlockNumber)
		mdl.Repo.AddPriceFeed(blockNum, &core.PriceFeed{
			BlockNumber: blockNum,
			Token:       mdl.Details["token"],
			Feed:        mdl.Address,
			RoundId:     roundId,
			PriceETHBI:  (answerBI).String(),
			PriceETH:    utils.GetFloat64Decimal(answerBI, 18),
		})
	}
}
