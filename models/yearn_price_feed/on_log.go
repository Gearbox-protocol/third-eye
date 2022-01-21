package yearn_price_feed

import (
	"fmt"
	"github.com/Gearbox-protocol/third-eye/artifacts/priceFeed"
	"github.com/Gearbox-protocol/third-eye/artifacts/yVault"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strings"
	"time"
)

func (mdl *YearnPriceFeed) OnLog(txLog types.Log) {

}

const interval = 25

func (mdl *YearnPriceFeed) Query(queryTill int64) {
	queryFrom := mdl.GetLastSync() + interval
	log.Infof("Sync %s(%s) from %d to %d", mdl.GetName(), mdl.GetAddress(), queryFrom, queryTill)
	rounds := 0
	loopStartTime := time.Now()
	roundStartTime := time.Now()
	for blockNum := queryFrom; blockNum <= queryTill; blockNum += interval {
		mdl.queryHandler(blockNum)
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

func (mdl *YearnPriceFeed) query(blockNum int64) (*core.PriceFeed, error) {
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
	}
	roundData, err := mdl.contractETH.LatestRoundData(opts)
	if err != nil {
		return nil, err
	}

	return &core.PriceFeed{
		RoundId:    roundData.RoundId.Int64(),
		PriceETHBI: (*core.BigInt)(roundData.Answer),
		PriceETH:   utils.GetFloat64Decimal(roundData.Answer, 18),
	}, nil
}

func (mdl *YearnPriceFeed) queryHandler(blockNum int64) {
	mdl.Repo.SetBlock(blockNum)
	pf, err := mdl.query(blockNum)
	if err != nil {
		if strings.Contains(err.Error(), "execution reverted") {
			pf = mdl.calculatePriceFeedInternally(blockNum)
		} else {
			log.CheckFatal(fmt.Errorf("%s %s", mdl.GetAddress(), err))
		}
	}
	tokenAddr, ok := mdl.Details["token"].(string)
	if !ok {
		log.Fatal("Failing in asserting to string: %s", mdl.Details["token"])
	}
	pf.BlockNumber = blockNum
	pf.Token = tokenAddr
	pf.Feed = mdl.Address
	mdl.Repo.AddPriceFeed(blockNum, pf)
}

func (mdl *YearnPriceFeed) calculatePriceFeedInternally(blockNum int64) *core.PriceFeed {
	if mdl.YVaultContract == nil || mdl.PriceFeedContract == nil || mdl.DecimalDivider == nil {
		mdl.setContracts(blockNum)
	}
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
	}

	roundData, err := mdl.PriceFeedContract.LatestRoundData(opts)
	log.CheckFatal(err)

	pricePerShare, err := mdl.YVaultContract.PricePerShare(opts)
	log.CheckFatal(err)

	newAnswer := new(big.Int).Quo(
		new(big.Int).Mul(pricePerShare, roundData.Answer),
		mdl.DecimalDivider,
	)

	return &core.PriceFeed{
		RoundId:    roundData.RoundId.Int64(),
		PriceETHBI: (*core.BigInt)(newAnswer),
		PriceETH:   utils.GetFloat64Decimal(newAnswer, 18),
	}
}

func (mdl *YearnPriceFeed) setContracts(blockNum int64) {
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
	}
	// set the price feed contract
	priceFeedAddr, err := mdl.contractETH.PriceFeed(opts)
	log.CheckFatal(err)
	priceFeedContract, err := priceFeed.NewPriceFeed(priceFeedAddr, mdl.Client)
	log.CheckFatal(err)
	mdl.PriceFeedContract = priceFeedContract

	// set the yvault contract
	yVaultAddr, err := mdl.contractETH.YVault(opts)
	log.CheckFatal(err)
	yVaultContract, err := yVault.NewYVault(yVaultAddr, mdl.Client)
	log.CheckFatal(err)
	mdl.YVaultContract = yVaultContract

	// set the decimals
	decimals, err := yVaultContract.Decimals(opts)
	log.CheckFatal(err)
	mdl.DecimalDivider = utils.GetExpInt(int8(decimals))
}
