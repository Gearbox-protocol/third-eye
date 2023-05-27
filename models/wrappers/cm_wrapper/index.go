package cm_wrapper

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/wrappers"
)

type CMWrapper struct {
	*wrappers.SyncWrapper
}

func NewCMWrapper(client core.ClientI) *CMWrapper {
	w := &CMWrapper{
		SyncWrapper: wrappers.NewSyncWrapper(ds.CMWrapper, client),
	}
	w.ViaDataProcess = ds.ViaMultipleLogs
	return w
}

func (CMWrapper) SetUnderlyingState(obj interface{}) {

}
