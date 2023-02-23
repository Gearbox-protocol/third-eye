package chainlink_price_feed

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/common"
)

type ChainlinkPriceFeed struct {
	*ds.SyncAdapter
	Token   string
	MainAgg *ChainlinkMainAgg
}

// if oracle and address are same then the normal chainlink interface is not working for this price feed
// it maybe custom price feed of gearbox . so we will disable on 'vm execution error' or 'execution reverted'.
// if oracle and adress are same we try to get the pricefeed.
func NewChainlinkPriceFeed(token, oracle string, discoveredAt int64, client core.ClientI, repo ds.RepositoryI, version int16, bounded bool) *ChainlinkPriceFeed {
	var upperLimit string
	if bounded {
		returnData, err := core.CallFuncWithExtraBytes(client, "b09ad8a0", common.HexToAddress(oracle), discoveredAt, nil) // upperBound
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
	return adapter
}

func NewChainlinkPriceFeedFromAdapter(adapter *ds.SyncAdapter, includeLastLogBeforeDiscover bool) *ChainlinkPriceFeed {
	oracleAddr, ok := adapter.Details["oracle"].(string)
	if !ok {
		log.Fatalf("Failed asserting oracle address(%s) as string for chainlink pricefeed(%s) ", adapter.Details["oracle"], adapter.GetAddress())
	}
	token, ok := adapter.Details["token"].(string)
	if !ok {
		log.Fatalf("Get token addr(%v) for oracle(%s) feed(%s)", adapter.Details["token"], oracleAddr, adapter.GetAddress())
	}
	obj := &ChainlinkPriceFeed{
		SyncAdapter: adapter,
		Token:       token,
	}
	obj.MainAgg = NewMainAgg(adapter.Client, common.HexToAddress(oracleAddr), obj.upperLimit().Cmp(new(big.Int)) != 0) // isBounded if upperlimit is not 0

	// feed address is empty
	if adapter.Address == "" {
		pfAddr, _ := obj.MainAgg.GetPriceFeedAddr(adapter.DiscoveredAt)
		obj.SetAddress(pfAddr.Hex())
	}
	// get the last log before the chainlink feed is added to price oracle.
	if includeLastLogBeforeDiscover {
		if lastLogBeforeDiscoverNum, err := obj.FindLastLogBound(1, obj.DiscoveredAt-1, []common.Hash{
			core.Topic("AnswerUpdated(int256,uint256,uint256)"),
		}); err != nil {
			log.Fatalf("%s for chainlink(%s) discovered_at %d", err, adapter.GetAddress(), obj.DiscoveredAt)
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
	newPriceFeed, newPhaseId := mdl.MainAgg.GetPriceFeedAddr(syncedTill)
	if newPriceFeed != common.HexToAddress(mdl.Address) && newPriceFeed != core.NULL_ADDR {
		var discoveredAt int64
		if newPhaseId != -1 {
			discoveredAt = mdl.MainAgg.GetFeedUpdateBlockUsingPhaseId(uint16(newPhaseId), mdl.LastSync+1, syncedTill)
		} else {
			discoveredAt = mdl.MainAgg.GetFeedUpdateBlockAggregator(newPriceFeed, mdl.LastSync+1, syncedTill)
		}
		mdl.Repo.AddNewPriceOracleEvent(&schemas.TokenOracle{
			Token:       mdl.Token,
			Oracle:      mdl.MainAgg.Addr.Hex(),
			Feed:        mdl.MainAgg.Addr.Hex(), // feed is same as oracle
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
