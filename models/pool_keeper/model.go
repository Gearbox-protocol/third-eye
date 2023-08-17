package pool_keeper

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas/schemas_v3"
	"github.com/Gearbox-protocol/third-eye/ds"
)

type PoolKeeper struct {
	*ds.SyncAdapter
	quotas       map[string]*schemas_v3.QuotaDetails
	lastBlockNum int64
}

func NewPoolKeeper(addr string, pool string, discoveredAt int64, client core.ClientI, repo ds.RepositoryI) *PoolKeeper {
	adapter := ds.NewSyncAdapter(addr, ds.PoolKeeper, discoveredAt, client, repo)
	adapter.Details = core.Json{"pool": pool}
	return NewPoolKeeperFromAdapter(
		adapter,
	)
}

func NewPoolKeeperFromAdapter(adapter *ds.SyncAdapter) *PoolKeeper {
	return &PoolKeeper{
		SyncAdapter: adapter,
		quotas:      map[string]*schemas_v3.QuotaDetails{},
	}
}

func (mdl PoolKeeper) OnBlockChange(lastBlockNum int64) {
	if mdl.lastBlockNum != 0 && mdl.lastBlockNum == lastBlockNum {
		for _, details := range mdl.quotas {
			if details.IsDirty {
				mdl.Repo.AddQuotaDetails(details)
				details.IsDirty = false
			}
		}
	}
}
