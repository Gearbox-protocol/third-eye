package pool_quota_keeper

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas/schemas_v3"
	"github.com/Gearbox-protocol/third-eye/ds"
)

// manages the quota of all accounts using account quota manager
// manages current quota details of all tokens in pool
type PoolQuotaKeeper struct {
	*ds.SyncAdapter
	mgr          *ds.AccountQuotaMgr
	quotas       map[string]*schemas_v3.QuotaDetails
	lastBlockNum int64

	// account to token to details
}

func NewPoolQuotaKeeper(addr string, pool string, discoveredAt int64, client core.ClientI, repo ds.RepositoryI) *PoolQuotaKeeper {
	adapter := ds.NewSyncAdapter(addr, ds.PoolQuotaKeeper, discoveredAt, client, repo)
	adapter.Details = core.Json{"pool": pool}
	return NewPoolQuotaKeeperFromAdapter(
		adapter,
	)
}

func NewPoolQuotaKeeperFromAdapter(adapter *ds.SyncAdapter) *PoolQuotaKeeper {
	return &PoolQuotaKeeper{

		SyncAdapter: adapter,
		mgr:         adapter.Repo.GetAccountQuotaMgr(),
		quotas:      map[string]*schemas_v3.QuotaDetails{},
	}
}

func (mdl PoolQuotaKeeper) OnBlockChange(lastBlockNum int64) {
	if mdl.lastBlockNum != 0 && mdl.lastBlockNum == lastBlockNum {
		for _, details := range mdl.quotas {
			if details.IsDirty {
				mdl.Repo.AddQuotaDetails(details.Copy())
				details.IsDirty = false
			}
		}
	}
}
