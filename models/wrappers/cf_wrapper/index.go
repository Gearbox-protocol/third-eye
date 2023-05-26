package cf_wrapper

import (
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/wrappers"
)

type CFWrapper struct {
	*wrappers.SyncWrapper
}

func NewCFWrapper() *CFWrapper {
	return &CFWrapper{
		SyncWrapper: wrappers.NewSyncWrapper(ds.CFWrapper),
	}
}

func (CFWrapper) SetUnderlyingState(obj interface{}) {

}
