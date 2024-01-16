package pool_wrapper

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/pool/pool_v3"
	"github.com/Gearbox-protocol/third-eye/models/wrappers"
	"gorm.io/gorm"
)

type PoolWrapper struct {
	*wrappers.SyncWrapper
}

func NewPoolWrapper(client core.ClientI) *PoolWrapper {
	w := &PoolWrapper{
		SyncWrapper: wrappers.NewSyncWrapper(ds.PoolWrapper, client),
	}
	w.ViaDataProcess = ds.ViaMultipleLogs
	return w
}

func (w *PoolWrapper) UpdatePoolv2Ledger(tx *gorm.DB) {
	for _, adapter := range w.Adapters.GetAll() {
		if adapter.GetVersion().MoreThanEq(core.NewVersion(300)) {
			adapter.(*pool_v3.Poolv3).UpdatePoolv2Ledger(tx)
		}
	}
}
