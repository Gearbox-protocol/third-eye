package chainlink_price_feed

import (
	"context"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/priceFeed"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type ChainlinkMainAgg struct {
	contractETH *priceFeed.PriceFeed
	Addr        common.Address
	Client      core.ClientI
}

func NewMainAgg(client core.ClientI, mainAgg common.Address) *ChainlinkMainAgg {
	pfContract, err := priceFeed.NewPriceFeed(mainAgg, client)
	if err != nil {
		log.Fatal(err)
	}
	return &ChainlinkMainAgg{
		contractETH: pfContract,
		Client:      client,
		Addr:        mainAgg,
	}
}

func (mdl *ChainlinkMainAgg) GetFeedUpdateBlockUsingPhaseId(newPhaseId uint16, from, to int64) int64 {
	if from == to {
		return from
	}
	midBlockNum := (from + to) / 2
	phaseId, err := mdl.contractETH.PhaseId(&bind.CallOpts{BlockNumber: big.NewInt(midBlockNum)})
	log.CheckFatal(err)
	if phaseId != newPhaseId {
		return mdl.GetFeedUpdateBlockUsingPhaseId(newPhaseId, midBlockNum+1, to)
	} else {
		return mdl.GetFeedUpdateBlockUsingPhaseId(newPhaseId, from, midBlockNum)
	}
}

func (mdl *ChainlinkMainAgg) GetFeedUpdateBlockAggregator(newAggAddr common.Address, from, to int64) int64 {
	if from == to {
		return from
	}
	midBlockNum := (from + to) / 2
	aggAddr, err := mdl.contractETH.Aggregator(&bind.CallOpts{BlockNumber: big.NewInt(midBlockNum)})
	log.CheckFatal(err)
	if aggAddr != newAggAddr {
		return mdl.GetFeedUpdateBlockAggregator(newAggAddr, midBlockNum+1, to)
	} else {
		return mdl.GetFeedUpdateBlockAggregator(newAggAddr, from, midBlockNum)
	}
}

func (mdl *ChainlinkMainAgg) GetPriceFeedAddr(blockNum int64) (common.Address, int16) {
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
	}
	phaseId, err := mdl.contractETH.PhaseId(opts)
	if err != nil {
		if mdl.Addr == common.HexToAddress("0xE26FB07da646138553f635c94E2a345270240e30") { // goerli specific case, where bounded MainAgg uses kovan-playground chainlink implementation that doesn't have aggregator method
			return common.HexToAddress("0xd6852347062aB885B6Fb9F7220BedCc5A39CE862"), -1
		}
		newPriceFeed, err2 := mdl.contractETH.Aggregator(opts)
		if err2 != nil {
			log.Fatalf("For MainAgg(%s) error: %s", mdl.Addr, err2)
		}
		return newPriceFeed, -1
	}
	var newPriceFeed common.Address
	newPriceFeed, err = mdl.contractETH.PhaseAggregators(opts, phaseId, false)
	if err != nil {
		chainId, err2 := mdl.Client.ChainID(context.TODO())
		log.CheckFatal(err2)
		if chainId.Int64() == 42 || chainId.Int64() == 5 { // for goerli and kovan test the phaseaggregator method is without 's'
			newPriceFeed, err = mdl.contractETH.PhaseAggregators(opts, phaseId, true)
			// try with method name phaseAggregator instead of phaseAggregators
			// true is sets typo=true so that phaseAggregator method is used.
		}
		if err != nil {
			log.Fatal(mdl.Addr, err)
		}
	}
	return newPriceFeed, int16(phaseId)
}
