package pool_quota_wrapper

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/wrappers"
)

type PoolQuotaWrapper struct {
	*wrappers.SyncWrapper
}

func NewPoolQuotaWrapper(client core.ClientI) *PoolQuotaWrapper {
	w := &PoolQuotaWrapper{
		SyncWrapper: wrappers.NewSyncWrapper(ds.PoolQuotaWrapper, client),
	}
	w.ViaDataProcess = ds.ViaMultipleLogs
	return w
}
