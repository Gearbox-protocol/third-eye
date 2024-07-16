package v2

import (
	"context"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
)

type LMRewardsv2 struct {
	*ds.SyncAdapter
	pendingCalcBlock int64
	chainId          int64
	// diesel symbol to user to balance
	dieselBalances map[string]map[string]*big.Int
	// pool  to user to reward
	rewards map[string]map[string]*big.Int
	// diesel symbol to total supply
	totalSupplies map[string]*big.Int
	// pool to decimal
	poolToDecimal map[string]symAndDecimals
	hasDataToSave bool
}
type symAndDecimals struct {
	decimals  int8
	dieselsym string
}

func NewLMRewardsv2(addr string, syncedTill int64, client core.ClientI, repo ds.RepositoryI) *LMRewardsv2 {
	return NewLMRewardsv2FromAdapter(
		&ds.SyncAdapter{
			SyncAdapterSchema: &schemas.SyncAdapterSchema{
				LastSync: syncedTill,
				Contract: &schemas.Contract{
					ContractName: ds.LMRewardsv2,
					Address:      addr,
					Client:       client,
				},
				V: core.NewVersion(1),
			},
			Repo: repo,
		},
	)
}

func NewLMRewardsv2FromAdapter(adapter *ds.SyncAdapter) *LMRewardsv2 {
	chainId, err := adapter.Client.ChainID(context.Background())
	log.CheckFatal(err)
	obj := &LMRewardsv2{
		SyncAdapter:      adapter,
		pendingCalcBlock: adapter.LastSync + 1,
		chainId:          chainId.Int64(),
		dieselBalances:   map[string]map[string]*big.Int{}, // to DieselBalances for saving in DB
		rewards:          map[string]map[string]*big.Int{}, // to LMRewards for saving in DB
		totalSupplies:    map[string]*big.Int{},            // will be converted to details on syncAdapter
		poolToDecimal:    map[string]symAndDecimals{},      // auxillary data

	}
	obj.detailsToTotalSupplies()
	return obj
}

func (mdl *LMRewardsv2) AfterSyncHook(syncedTill int64) {
	mdl.calculateRewards(mdl.pendingCalcBlock, syncedTill)
	mdl.pendingCalcBlock = syncedTill + 1
	//
	mdl.totalSuppliesToDetails() // convert store the supplies in details
	mdl.SyncAdapter.AfterSyncHook(syncedTill)
	// sync pool_lm_rewards and diesel_balances if PoolLMrewards has data to save,
	// i.e. it synced over the given range of blocks in the sync engine
	mdl.hasDataToSave = true
}
