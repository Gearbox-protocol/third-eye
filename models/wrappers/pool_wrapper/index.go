package pool_wrapper

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/wrappers"
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

func (PoolWrapper) SetUnderlyingState(obj interface{}) {

}
