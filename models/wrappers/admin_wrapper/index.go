package admin_wrapper

import (
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/wrappers"
)

type AdminWrapper struct {
	*wrappers.SyncWrapper
	nameToAddr map[string]string
}

func NewAdminWrapper() *AdminWrapper {
	w := &AdminWrapper{
		SyncWrapper: wrappers.NewSyncWrapper(ds.AdminWrapper, nil),
		nameToAddr:  make(map[string]string),
	}
	w.ViaDataProcess = ds.ViaLog
	return w
}

func (AdminWrapper) SetUnderlyingState(obj interface{}) {

}

func (w *AdminWrapper) AddSyncAdapter(adapter ds.SyncAdapterI) {
	w.nameToAddr[adapter.GetName()] = adapter.GetAddress()
	w.SyncWrapper.AddSyncAdapter(adapter)
}

func (w *AdminWrapper) GetAdapterAddrByName(name string) []string {
	return []string{w.nameToAddr[name]}
}
