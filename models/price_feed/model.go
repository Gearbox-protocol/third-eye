package price_feed

import (
	"github.com/Gearbox-protocol/gearscan/core"
	"github.com/Gearbox-protocol/gearscan/ethclient"
	"github.com/Gearbox-protocol/gearscan/artifacts/priceFeed"
	"github.com/ethereum/go-ethereum/common"
	"github.com/Gearbox-protocol/gearscan/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
)

type PriceFeed struct {
	*core.SyncAdapter
	*core.State
	contractETH *priceFeed.PriceFeed
}


func NewPriceFeed(oracle string, discoveredAt int64, client *ethclient.Client, repo core.RepositoryI) *PriceFeed {
	syncAdapter := core.NewSyncAdapter("", "PriceFeed", discoveredAt, client)
	syncAdapter.Oracle = oracle
	return NewPriceFeedFromAdapter(
		repo, syncAdapter,
	)
}

func NewPriceFeedFromAdapter(repo core.RepositoryI, adapter *core.SyncAdapter) *PriceFeed {
	pfContract, err := priceFeed.NewPriceFeed(common.HexToAddress(adapter.Oracle), adapter.Client)
	if err != nil {
		log.Fatal(err)
	}
	obj := &PriceFeed{
		SyncAdapter: adapter,
		State: &core.State{Repo: repo},
		contractETH: pfContract,
	}
	pfAddr := obj.GetPriceFeed(adapter.DiscoveredAt)
	obj.SetAddress(pfAddr)
	return obj
}

func (mdl *PriceFeed) AfterSyncHook(syncedTill int64)  {
	newPriceFeed := mdl.GetPriceFeed(mdl.LastSync)
	if newPriceFeed != mdl.Address {
		mdl.Disabled = true
		mdl.Repo.AddSyncAdapter(
			NewPriceFeed(newPriceFeed, mdl.LastSync + 1, mdl.Client, mdl.Repo),
		)
	}
	mdl.SetLastSync(syncedTill)
}

func (mdl *PriceFeed) GetPriceFeed(blockNum int64) string {
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
	}
	phaseId, err := mdl.contractETH.PhaseId(opts)
	if err != nil {
		log.Fatal(err)
	}
	newPriceFeed, err := mdl.contractETH.PhaseAggregators(opts, phaseId)
	if err != nil {
		log.Fatal(err)
	}
	return newPriceFeed.Hex()
}