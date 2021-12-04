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
	contractETH *priceFeed.PriceFeed
}

// if oracle and address are same then the normal chainlink interface is not working for this price feed
// it maybe custom price feed of gearbox . so we will disable on vm execution error.
// if oracle and adress are same we try to get the pricefeed.
func NewChainlinkPriceFeed(oracle, token string, discoveredAt int64, client *ethclient.Client, repo core.RepositoryI) *ChainlinkPriceFeed {
	syncAdapter := &core.SyncAdapter{
		Contract: &core.Contract{
			Address:      oracle,
			DiscoveredAt: discoveredAt,
			FirstLogAt:   discoveredAt,
			ContractName: "ChainlinkPriceFeed",
			Client:       client,
		},
		Details:  map[string]string{"oracle": oracle, "token": token},
		LastSync: discoveredAt - 1,
		Repo:     repo,
	}
	return NewChainlinkPriceFeedFromAdapter(
		syncAdapter,
	)
}

func NewChainlinkPriceFeedFromAdapter(adapter *core.SyncAdapter) *ChainlinkPriceFeed {
	oracleAddr := adapter.Details["oracle"]
	pfContract, err := priceFeed.NewPriceFeed(common.HexToAddress(oracleAddr), adapter.Client)
	if err != nil {
		log.Fatal(err)
	}
	obj := &ChainlinkPriceFeed{
		SyncAdapter: adapter,
		contractETH: pfContract,
	}
	if adapter.Address == oracleAddr {
		pfAddr := obj.GetPriceFeedAddr(adapter.DiscoveredAt)
		obj.SetAddress(pfAddr)
	}
	return obj
}

func (mdl *ChainlinkPriceFeed) AfterSyncHook(syncedTill int64) {
	newPriceFeed := mdl.GetPriceFeedAddr(mdl.LastSync)
	if newPriceFeed != mdl.Address {
		mdl.Disable()
		mdl.Repo.AddSyncAdapter(
			NewChainlinkPriceFeed(newPriceFeed, mdl.Details["token"], mdl.LastSync+1, mdl.Client, mdl.Repo),
		)
	}
	mdl.SetLastSync(syncedTill)
}

func (mdl *ChainlinkPriceFeed) GetPriceFeedAddr(blockNum int64) string {
	opts, cancel := utils.GetTimeoutOpts(blockNum)
	defer cancel()
	phaseId, err := mdl.contractETH.PhaseId(opts)
	if err != nil {
		if err.Error() == "execution aborted (timeout = 20s)" {
			log.Fatal(err)
		} else {
			mdl.SetError(err)
			oralceAddr := mdl.Details["oracle"]
			log.Error(oralceAddr, " oracle failed disabling due to ", err)
			return oralceAddr
		}
	}
	opts, cancel = utils.GetTimeoutOpts(blockNum)
	defer cancel()
	newPriceFeed, err := mdl.contractETH.PhaseAggregators(opts, phaseId)
	if err != nil {
		log.Fatal(mdl.Address, err)
	}
	return newPriceFeed.Hex()
}
