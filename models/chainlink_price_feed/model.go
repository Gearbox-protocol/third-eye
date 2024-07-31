package chainlink_price_feed

import (
	"math"
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
	MainAgg         *ChainlinkMainAgg
	mergedPFManager *ds.MergedPFManager
	tokens          []string
	pfs             []*schemas.PriceFeed
}

// if oracle and address are same then the normal chainlink interface is not working for this price feed
// it maybe custom price feed of gearbox . so we will disable on 'vm execution error' or 'execution reverted'.
// if oracle and adress are same we try to get the pricefeed.
func NewChainlinkPriceFeed(client core.ClientI, repo ds.RepositoryI, token, oracle string, discoveredAt int64, mergedPFVersion schemas.MergedPFVersion, bounded bool, includeLastLogBeforeDiscover ...bool) *ChainlinkPriceFeed {
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
			Details:  map[string]interface{}{"oracle": oracle, "token": token, "mergedPFVersion": mergedPFVersion},
			LastSync: discoveredAt - 1,
			V:        mergedPFVersion.MergedPFVersionToList()[0].ToVersion(),
		},
		Repo: repo,
	}
	if upperLimit != "" {
		syncAdapter.Details["upperLimit"] = upperLimit
	}
	x := true
	if len(includeLastLogBeforeDiscover) > 0 {
		x = includeLastLogBeforeDiscover[0]
	}
	adapter := NewChainlinkPriceFeedFromAdapter(
		syncAdapter,
		x,
	)
	return adapter
}
func NewChainlinkPriceFeedFromAdapter(adapter *ds.SyncAdapter, includeLastLogBeforeDiscover bool) *ChainlinkPriceFeed {
	oracleAddr, ok := adapter.Details["oracle"].(string)
	if !ok {
		log.Fatalf("Failed asserting oracle address(%s) as string for chainlink pricefeed(%s) ", adapter.Details["oracle"], adapter.GetAddress())
	}
	obj := &ChainlinkPriceFeed{
		SyncAdapter: adapter,
	}
	obj.MainAgg = NewMainAgg(adapter.Client, common.HexToAddress(oracleAddr), obj.upperLimit().Cmp(new(big.Int)) != 0) // isBounded if upperlimit is not 0

	// feed address is empty
	if adapter.Address == "" {
		pfAddr, _ := obj.MainAgg.GetPriceFeedAddr(adapter.DiscoveredAt)
		obj.SetAddress(pfAddr.Hex())
	}
	// get the last log before the chainlink feed is added to price oracle.
	if includeLastLogBeforeDiscover {
		var lastLogBeforeDiscoverNum int64
		// TODO anvil fork testing
		var err error
		if core.GetChainId(adapter.Client) == 7878 {
			lastLogBeforeDiscoverNum = obj.DiscoveredAt - 3000
		} else {
			lastLogBeforeDiscoverNum, err = obj.FindLastLogBound(1, obj.DiscoveredAt-1, []common.Hash{
				core.Topic("AnswerUpdated(int256,uint256,uint256)"),
			})
			if err != nil {
				log.Fatalf("%s for chainlink(%s) discovered_at %d", err, adapter.GetAddress(), obj.DiscoveredAt)
			}
		}
		if lastLogBeforeDiscoverNum != 0 {
			obj.LastSync = lastLogBeforeDiscoverNum - 1
			obj.FirstLogAt = lastLogBeforeDiscoverNum
		}
	}
	obj.DataProcessType = ds.ViaMultipleLogs
	obj.mergedPFManager = &ds.MergedPFManager{}
	obj.mergedPFManager.Load(obj.Details, obj.FirstLogAt)
	return obj
}

func (mdl *ChainlinkPriceFeed) flushPrices(nextFeedAt int64) {
	for _, pf := range mdl.pfs {
		if pf.BlockNumber < nextFeedAt {
			mdl.Repo.AddPriceFeed(pf)
		} else {
			break
		}
	}
	mdl.pfs = nil
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
		log.Info(mdl.Address, discoveredAt, newPriceFeed)
		mdl.flushPrices(discoveredAt)
		for _, token := range mdl.mergedPFManager.GetTokens(discoveredAt) {
			mdl.Repo.AddNewPriceOracleEvent(&schemas.TokenOracle{
				Token:       token,
				Oracle:      mdl.MainAgg.Addr.Hex(),
				Feed:        mdl.MainAgg.Addr.Hex(), // feed is same as oracle
				BlockNumber: discoveredAt,
				Version:     mdl.GetVersion(),
				Reserve:     schemas.GetReservefromDetails(mdl.Details),
				FeedType:    ds.ChainlinkPriceFeed,
			}, mdl.upperLimit().Cmp(new(big.Int)) != 0, false) // if upperLImit is not zero, then the price is bounded by upperLimit
		}
	}
	mdl.flushPrices(math.MaxInt64)
	mdl.SyncAdapter.AfterSyncHook(syncedTill)
	mdl.mergedPFManager.CloseV2(mdl.Client, syncedTill, mdl.Address)
	mdl.mergedPFManager.Save(&mdl.Details)
}

func (mdl *ChainlinkPriceFeed) upperLimit() *big.Int {
	details := mdl.GetDetails()
	if details == nil || details["upperLimit"] == nil {
		return new(big.Int)
	}
	upperLimit := details["upperLimit"].(string)
	return utils.StringToInt(upperLimit)
}
