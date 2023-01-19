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
	bounded     bool
}

func NewMainAgg(client core.ClientI, mainAgg common.Address, bounded ...bool) *ChainlinkMainAgg {
	pfContract, err := priceFeed.NewPriceFeed(mainAgg, client)
	if err != nil {
		log.Fatal(err)
	}
	return &ChainlinkMainAgg{
		contractETH: pfContract,
		Client:      client,
		Addr:        mainAgg,
		bounded:     len(bounded) > 0 && bounded[0],
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
	if mdl.bounded {
		return mdl.getPriceFeedAddrOnBounded(blockNum)
	}
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
	}
	phaseId, err := mdl.contractETH.PhaseId(opts)
	if err != nil {
		log.Fatal("PhaseId not founded for ", mdl.Addr)
	}
	var newPhaseAgg common.Address
	newPhaseAgg, err = mdl.contractETH.PhaseAggregators(opts, phaseId, false)
	if err != nil {
		chainId, err2 := mdl.Client.ChainID(context.TODO())
		log.CheckFatal(err2)
		if chainId.Int64() == 42 || chainId.Int64() == 5 { // for goerli and kovan test the phaseaggregator method is without 's'
			// try with method name phaseAggregator instead of phaseAggregators
			// true is sets typo=true so that phaseAggregator method is used.
			newPhaseAgg, err = mdl.contractETH.PhaseAggregators(opts, phaseId, true)
		}
		if err != nil {
			log.Fatal(mdl.Addr, err)
		}
	}
	return newPhaseAgg, int16(phaseId)
}

func (mdl *ChainlinkMainAgg) getPriceFeedAddrOnBounded(blockNum int64) (common.Address, int16) {
	if mdl.Addr == common.HexToAddress("0xE26FB07da646138553f635c94E2a345270240e30") { // goerli specific case, where bounded MainAgg uses kovan-playground chainlink implementation that doesn't have aggregator method
		return common.HexToAddress("0xd6852347062aB885B6Fb9F7220BedCc5A39CE862"), -1
	}
	//
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
	}
	phaseId, err := mdl.contractETH.PhaseId(opts)
	if err != nil {
		newPhaseAgg, err2 := mdl.contractETH.Aggregator(opts)
		if err2 != nil {
			log.Fatalf("For MainAgg(%s) error: %s", mdl.Addr, err2)
		}
		return newPhaseAgg, -1
	}
	//
	chainId, err2 := mdl.Client.ChainID(context.TODO())
	log.CheckFatal(err2)
	var newPhaseAgg common.Address
	if chainId.Int64() == 1 {
		// get phaseAggregator on the boundedFeed
		newPhaseAgg, err = mdl.contractETH.PhaseAggregators(opts, phaseId, false)
	} else { // goerli and kovan
		// get priceFeed
		underlyingBoundedFeed, err3 := core.CallFuncWithExtraBytes(mdl.Client, "741bef1a",
			mdl.Addr, blockNum, nil) // priceFeed for [bounded chainlink oracle]
		if err != nil {
			log.Fatalf("For bounded oracle(%s) underlying priceFeed not found, err: %s", mdl.Addr, err3)
		}
		// get phaseAggregator on priceFeed
		extras := [32]byte{}
		extras[31] = byte(phaseId)
		var phaseAggregatorData []byte
		// phaseAggregator only on goerli
		phaseAggregatorData, err = core.CallFuncWithExtraBytes(mdl.Client, "d6bcd745",
			common.BytesToAddress(underlyingBoundedFeed), blockNum, extras[:])
		newPhaseAgg = common.BytesToAddress(phaseAggregatorData)
	}
	return newPhaseAgg, int16(phaseId)
}
