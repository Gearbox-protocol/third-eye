package chainlink_price_feed

import (
	"github.com/Gearbox-protocol/third-eye/artifacts/priceFeed"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/utils"
	"github.com/ethereum/go-ethereum/common"
)

type ChainlinkPriceFeed struct {
	*core.SyncAdapter
	contractETH   *priceFeed.PriceFeed
	Token         string
	Oracle        string
	prevPriceFeed *core.PriceFeed
}

// if oracle and address are same then the normal chainlink interface is not working for this price feed
// it maybe custom price feed of gearbox . so we will disable on 'vm execution error' or 'execution reverted'.
// if oracle and adress are same we try to get the pricefeed.
func NewChainlinkPriceFeed(token, oracle, feed string, discoveredAt int64, client *ethclient.Client, repo core.RepositoryI) *ChainlinkPriceFeed {
	syncAdapter := &core.SyncAdapter{
		Contract: &core.Contract{
			Address:      feed,
			DiscoveredAt: discoveredAt,
			FirstLogAt:   discoveredAt,
			ContractName: core.ChainlinkPriceFeed,
			Client:       client,
		},
		Details:  map[string]interface{}{"oracle": oracle, "token": token},
		LastSync: discoveredAt - 1,
		Repo:     repo,
	}
	adapter := NewChainlinkPriceFeedFromAdapter(
		syncAdapter,
		true,
	)
	repo.AddTokenOracle(token, oracle, feed, discoveredAt)
	return adapter
}

func NewChainlinkPriceFeedFromAdapter(adapter *core.SyncAdapter, includeLastLogBeforeDiscover bool) *ChainlinkPriceFeed {
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
	if adapter.Address == oracleAddr {
		pfAddr := obj.GetPriceFeedAddr(adapter.DiscoveredAt)
		obj.SetAddress(pfAddr)
	}
	if includeLastLogBeforeDiscover {
		if lastLogBeforeDiscoverNum, err := obj.FindLastLogBound(1, obj.DiscoveredAt-1, []common.Hash{
			core.Topic("AnswerUpdated(int256,uint256,uint256)"),
		}); err != nil {
			log.Fatal(err)
		} else {
			if lastLogBeforeDiscoverNum != 0 {
				obj.LastSync = lastLogBeforeDiscoverNum - 1
				obj.FirstLogAt = lastLogBeforeDiscoverNum
			}
		}
	}
	obj.HasOnLogs = true
	adapter.Repo.AddPoolsForToken(adapter.DiscoveredAt, token)
	obj.Repo.AddLastSyncForToken(token, obj.GetLastSync())
	return obj
}

func (mdl *ChainlinkPriceFeed) AfterSyncHook(syncedTill int64) {
	newPriceFeed := mdl.GetPriceFeedAddr(mdl.LastSync)
	if newPriceFeed != mdl.Address && newPriceFeed != "" {
		mdl.Repo.AddSyncAdapter(
			NewChainlinkPriceFeed(mdl.Token, mdl.Oracle, newPriceFeed, mdl.LastSync+1, mdl.Client, mdl.Repo),
		)
	}
	mdl.SyncAdapter.AfterSyncHook(syncedTill)
}

func (mdl *ChainlinkPriceFeed) GetPriceFeedAddr(blockNum int64) string {
	opts, cancel := utils.GetTimeoutOpts(blockNum)
	defer cancel()
	phaseId, err := mdl.contractETH.PhaseId(opts)
	if err != nil {
		log.Fatal(err)
		// if err.Error() == "execution aborted (timeout = 20s)" {
		// } else {
		// 	mdl.SetError(err)
		// 	log.Error(mdl.GetAddress(), " feed failed disabling due to ", err)
		// 	return ""
		// }
	}
	opts, cancel = utils.GetTimeoutOpts(blockNum)
	defer cancel()
	newPriceFeed, err := mdl.contractETH.PhaseAggregators(opts, phaseId)
	if err != nil {
		log.Fatal(mdl.Address, err)
	}
	return newPriceFeed.Hex()
}
