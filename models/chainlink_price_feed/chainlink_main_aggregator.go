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
	pfContract, err := priceFeed.NewPriceFeed(mainAgg, client) // on oracle
	log.CheckFatal(err)
	return &ChainlinkMainAgg{
		contractETH: pfContract,
		Client:      client,
		Addr:        mainAgg,
	}
}

// get blockNum at which the phaseId was changed.
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

// get blockNum at which the aggregator in the oracle was changed.
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

// chainlink oracle has priceFeed(phaseAggregator)
// get priceFeed at provided blockNum using PhaseId and phaseAggregator/phaseAggregators
func (mdl *ChainlinkMainAgg) GetPriceFeedAddr(blockNum int64) common.Address {
	// if mdl.bounded {
	// 	return mdl.getPriceFeedAddrOnBounded(blockNum)
	// }
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
	}
	// phaseId, err := mdl.contractETH.PhaseId(opts)
	// if err != nil {
	// 	log.Fatal("PhaseId not founded for ", mdl.Addr)
	// }
	newPhaseAgg, err := mdl.contractETH.Aggregator(opts) // typo check for phaseAggregator for only for kovan and goerli so remove that token.
	log.CheckFatal(err)
	return newPhaseAgg
}

// - if oracle is 0xE2, phaseAggregator(priceFeed) is 0xd6
// - if phaseId method is missing on boundedOracle, check for Aggregator if method is missing, fail Fatal
// - if mainnet(chainId=1), check phaseAggregators on boundedOracle
// - if goerli/kovan, call priceFeed for getting underlying aggregator, then on that aggregator call phaseAggregator(without s)
func (mdl *ChainlinkMainAgg) getPriceFeedAddrOnBounded(blockNum int64) (common.Address, int16) {
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
	// TODO anvil fork
	if chainId.Int64() == 1 || chainId.Int64() == 7878 {
		// get phaseAggregator on the boundedFeed
		newPhaseAgg, err := mdl.contractETH.PhaseAggregators(opts, phaseId)
		log.CheckFatal(err)
		return newPhaseAgg, int16(phaseId)
	} else { // goerli and kovan
		// get priceFeed
		underlyingBoundedFeed, err3 := core.CallFuncGetSingleValue(mdl.Client, "741bef1a",
			mdl.Addr, blockNum, nil) // priceFeed for [bounded chainlink oracle]
		if err != nil {
			log.Fatalf("For bounded oracle(%s) underlying priceFeed not found, err: %s", mdl.Addr, err3)
		}
		// get phaseAggregator on priceFeed
		extras := [32]byte{}
		extras[31] = byte(phaseId)
		var phaseAggregatorData []byte
		// phaseAggregator only on goerli
		phaseAggregatorData, err := core.CallFuncGetSingleValue(mdl.Client, "d6bcd745",
			common.BytesToAddress(underlyingBoundedFeed), blockNum, extras[:])
		log.CheckFatal(err)
		newPhaseAgg := common.BytesToAddress(phaseAggregatorData)
		return newPhaseAgg, int16(phaseId)
	}
}
