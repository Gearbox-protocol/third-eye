package yearn_price_feed

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

func (mdl *YearnPriceFeed) OnLog(txLog types.Log) {

}

const interval = 25

func (mdl *YearnPriceFeed) Query(queryTill int64) {
	for blockNum := (mdl.GetLastSync() + interval); blockNum <= queryTill; blockNum += interval {
		opts := &bind.CallOpts{
			BlockNumber: big.NewInt(blockNum),
		}
		roundData, err := mdl.contractETH.LatestRoundData(opts)
		if err != nil {
			log.Fatal(err)
		}
		mdl.Repo.AddPriceFeed(blockNum, &core.PriceFeed{
			BlockNumber: blockNum,
			Token:       mdl.Details["token"],
			Feed:        mdl.Address,
			RoundId:     roundData.RoundId.Int64(),
			PriceETHBI:  roundData.Answer.String(),
			PriceETH:    utils.GetFloat64Decimal(roundData.Answer, 18),
		})
	}
}
