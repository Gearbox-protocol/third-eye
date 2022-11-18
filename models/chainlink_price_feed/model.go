package chainlink_price_feed

import (
	"context"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/priceFeed"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/ds/dc_wrapper"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type ChainlinkPriceFeed struct {
	*ds.SyncAdapter
	contractETH   *priceFeed.PriceFeed
	Token         string
	Oracle        string
	prevPriceFeed *schemas.PriceFeed
}

// if oracle and address are same then the normal chainlink interface is not working for this price feed
// it maybe custom price feed of gearbox . so we will disable on 'vm execution error' or 'execution reverted'.
// if oracle and adress are same we try to get the pricefeed.
func NewChainlinkPriceFeed(token, oracle string, discoveredAt int64, client core.ClientI, repo ds.RepositoryI, version int16, bounded bool) *ChainlinkPriceFeed {
	var upperLimit string
	if bounded {
		returnData, err := dc_wrapper.CallFuncWithExtraBytes(client, "b09ad8a0", common.HexToAddress(oracle), discoveredAt, nil) // upperBound
		log.CheckFatal(err)
		upperLimit = new(big.Int).SetBytes(returnData).String()
	}
	syncAdapter := &ds.SyncAdapter{
		SyncAdapterSchema: &schemas.SyncAdapterSchema{
			Contract: &schemas.Contract{
				Address:      "",
				DiscoveredAt: discoveredAt,
				FirstLogAt:   discoveredAt,
				ContractName: ds.ChainlinkPriceFeed,
				Client:       client,
			},
			Details:  map[string]interface{}{"oracle": oracle, "token": token},
			LastSync: discoveredAt - 1,
			V:        version,
		},
		Repo: repo,
	}
	if upperLimit != "" {
		syncAdapter.Details["upperLimit"] = upperLimit
	}
	adapter := NewChainlinkPriceFeedFromAdapter(
		syncAdapter,
		true,
	)
	// repo.AddTokenOracle(&schemas.TokenOracle{
	// 	Token:       token,
	// 	Oracle:      adapter.Oracle,
	// 	Feed:        adapter.Address,
	// 	BlockNumber: discoveredAt,
	// 	Version:     version})
	// repo.AddUniPoolsForToken(adapter.DiscoveredAt, token)
	return adapter
}

func NewChainlinkPriceFeedFromAdapter(adapter *ds.SyncAdapter, includeLastLogBeforeDiscover bool) *ChainlinkPriceFeed {
	oracleAddr, ok := adapter.Details["oracle"].(string)
	if !ok {
		log.Fatal("Failed asserting oracle address(%s) as string for chainlink pricefeed(%s) ", adapter.Details["oracle"], adapter.GetAddress())
	}
	token, ok := adapter.Details["token"].(string)
	if !ok {
		log.Fatal("Get token addr(%v) for oracle(%s) feed(%s)", adapter.Details["token"], oracleAddr, adapter.GetAddress())
	}

	pfContract, err := priceFeed.NewPriceFeed(common.HexToAddress(oracleAddr), adapter.Client)
	if err != nil {
		log.Fatal(err)
	}
	obj := &ChainlinkPriceFeed{
		SyncAdapter: adapter,
		contractETH: pfContract,
		Token:       token,
		Oracle:      oracleAddr,
	}
	// feed address is emptyz
	if adapter.Address == "" {
		pfAddr, _ := obj.GetPriceFeedAddr(adapter.DiscoveredAt)
		obj.SetAddress(pfAddr.Hex())
	}
	// get the last log before the chainlink feed is added to price oracle.
	if includeLastLogBeforeDiscover {
		if lastLogBeforeDiscoverNum, err := obj.FindLastLogBound(1, obj.DiscoveredAt-1, []common.Hash{
			core.Topic("AnswerUpdated(int256,uint256,uint256)"),
		}); err != nil {
			log.Fatal(err, "for chainlink", adapter.GetAddress(), "with discovered_at", obj.DiscoveredAt)
		} else {
			if lastLogBeforeDiscoverNum != 0 {
				obj.LastSync = lastLogBeforeDiscoverNum - 1
				obj.FirstLogAt = lastLogBeforeDiscoverNum
			}
		}
	}
	obj.HasOnLogs = true
	return obj
}

func (mdl *ChainlinkPriceFeed) AfterSyncHook(syncedTill int64) {
	newPriceFeed, newPhaseId := mdl.GetPriceFeedAddr(syncedTill)
	if newPriceFeed != common.HexToAddress(mdl.Address) && newPriceFeed != core.NULL_ADDR {
		var discoveredAt int64
		if newPhaseId != -1 {
			discoveredAt = mdl.GetFeedUpdateBlockUsingPhaseId(uint16(newPhaseId), mdl.LastSync+1, syncedTill)
		} else {
			discoveredAt = mdl.GetFeedUpdateBlockAggregator(newPriceFeed, mdl.LastSync+1, syncedTill)
		}
		mdl.Repo.AddNewPriceOracleEvent(&schemas.TokenOracle{
			Token:       mdl.Token,
			Oracle:      mdl.Oracle,
			Feed:        mdl.Oracle, // feed is same as oracle
			BlockNumber: discoveredAt,
			Version:     mdl.GetVersion(),
			FeedType:    ds.ChainlinkPriceFeed,
		}, mdl.upperLimit().Cmp(new(big.Int)) != 0) // if upperLImit is not zero, then the price is bounded by upperLimit
	}
	mdl.SyncAdapter.AfterSyncHook(syncedTill)
}

func (mdl *ChainlinkPriceFeed) upperLimit() *big.Int {
	details := mdl.GetDetails()
	if details == nil || details["upperLimit"] == nil {
		return new(big.Int)
	}
	upperLimit := details["upperLimit"].(string)
	return utils.StringToInt(upperLimit)
}

func (mdl *ChainlinkPriceFeed) GetFeedUpdateBlockUsingPhaseId(newPhaseId uint16, from, to int64) int64 {
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

func (mdl *ChainlinkPriceFeed) GetFeedUpdateBlockAggregator(newAggAddr common.Address, from, to int64) int64 {
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

func (mdl *ChainlinkPriceFeed) GetPriceFeedAddr(blockNum int64) (common.Address, int16) {
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
	}
	phaseId, err := mdl.contractETH.PhaseId(opts)
	if err != nil {
		newPriceFeed, err2 := mdl.contractETH.Aggregator(opts)
		log.CheckFatal(err2)
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
			log.Fatal(mdl.Address, mdl.Details, mdl.GetVersion(), err)
		}
	}
	return newPriceFeed, int16(phaseId)
}
