package pool_lmrewards

import (
	"context"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
)

type PoolLMRewards struct {
	*ds.SyncAdapter
	pendingCalcBlock int64
	chainId          int64
	// diesel symbol to user to balance
	dieselBalances map[string]map[string]*big.Int
	// pool  to user to reward
	rewards map[string]map[string]*big.Int
	// diesel symbol to total supply
	totalSupplies map[string]*big.Int
	// sym to decimals and pool
	decimalsAndPool map[string]*_PoolAndDecimals
}
type _PoolAndDecimals struct {
	decimals int8
	pool     string
}

// func NewPoolLMRewards(addr string, discoveredAt int64, client core.ClientI, repo ds.RepositoryI) *PoolLMRewards {
// 	return NewPoolLMRewardsFromAdapter(
// 		&ds.SyncAdapter{
// 			SyncAdapterSchema: &schemas.SyncAdapterSchema{
// 				LastSync: 0,
// 				Contract: &schemas.Contract{
// 					ContractName: ds.PoolLMRewards,
// 					Address:      addr,
// 					Client:       client,
// 				},
// 			},
// 			Repo: repo,
// 		},
// 	)
// }

func NewPoolLMRewardsFromAdapter(adapter *ds.SyncAdapter) *PoolLMRewards {
	chainId, err := adapter.Client.ChainID(context.Background())
	log.CheckFatal(err)
	obj := &PoolLMRewards{
		SyncAdapter:      adapter,
		pendingCalcBlock: adapter.LastSync + 1,
		chainId:          chainId.Int64(),
		dieselBalances:   map[string]map[string]*big.Int{}, // to DieselBalances for saving in DB
		rewards:          map[string]map[string]*big.Int{}, // to LMRewards for saving in DB
		totalSupplies:    map[string]*big.Int{},            // will be converted to details on syncAdapter
		decimalsAndPool:  map[string]*_PoolAndDecimals{},   // auxillary data
	}
	obj.detailsToTotalSupplies()
	return obj
}

func (mdl *PoolLMRewards) AfterSyncHook(syncedTill int64) {
	mdl.calculateRewards(mdl.pendingCalcBlock, syncedTill)
	mdl.totalSuppliesToDetails() // convert store the supplies in details
	mdl.SyncAdapter.AfterSyncHook(syncedTill)
}
