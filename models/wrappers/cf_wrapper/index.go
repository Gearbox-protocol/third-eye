package cf_wrapper

import (
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/wrappers"
)

type CFWrapper struct {
	*wrappers.SyncWrapper
}

func NewCFWrapper() *CFWrapper {
	w := &CFWrapper{
		SyncWrapper: wrappers.NewSyncWrapper(ds.CFWrapper, nil),
	}
	w.ViaDataProcess = ds.ViaLog
	return w
}
