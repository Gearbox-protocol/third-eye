package yearn_price_feed

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"sync"
	"time"
)

func (mdl *YearnPriceFeed) OnLog(txLog types.Log) {

}

const interval = 25

func (mdl *YearnPriceFeed) Query(queryTill int64, wg *sync.WaitGroup) {
	defer wg.Done()
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
		mdl.query(blockNum)
		if rounds%100 == 0 {
			timeLeft := (time.Now().Sub(loopStartTime).Seconds() * float64(queryTill-blockNum)) /
				float64(blockNum-mdl.GetLastSync())
			timeLeft /= 60
			log.Infof("Synced %d in %d rounds(%fs): TimeLeft %f mins", blockNum, rounds, time.Now().Sub(roundStartTime).Seconds(), timeLeft)
			roundStartTime = time.Now()
		}
		rounds++
	}
	// after sync
	mdl.AfterSyncHook(queryTill)
}

func (mdl *YearnPriceFeed) query(blockNum int64) {
	mdl.Repo.SetBlock(blockNum)
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
	}
	roundData, err := mdl.contractETH.LatestRoundData(opts)
	if err != nil {
		log.Fatal(err)
	}
	tokenAddr, ok := mdl.Details["token"].(string)
	if !ok {
		log.Fatal("Failing in asserting to string: %s", mdl.Details["token"])
	}
	mdl.Repo.AddPriceFeed(blockNum, &core.PriceFeed{
		BlockNumber: blockNum,
		Token:       tokenAddr,
		Feed:        mdl.Address,
		RoundId:     roundData.RoundId.Int64(),
		PriceETHBI:  (*core.BigInt)(roundData.Answer),
		PriceETH:    utils.GetFloat64Decimal(roundData.Answer, 18),
	})
}
