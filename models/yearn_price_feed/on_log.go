package yearn_price_feed

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"time"
)

func (mdl *YearnPriceFeed) OnLog(txLog types.Log) {

}

const interval = 25

func (mdl *YearnPriceFeed) Query(queryTill int64) {
	queryFrom := mdl.GetLastSync() + interval
	if queryFrom > queryTill {
		return
	}
	log.Infof("Sync %s(%s) from %d to %d", mdl.GetName(), mdl.GetAddress(), queryFrom, queryTill)
	rounds := 0
	loopStartTime := time.Now()
	roundStartTime := time.Now()
	queryTill = utils.Min(mdl.GetBlockToDisableOn(), queryTill)
	// if disable block is set disable after that.
	for blockNum := queryFrom; blockNum <= queryTill; blockNum += interval {
		mdl.Repo.SetBlock(blockNum)
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
