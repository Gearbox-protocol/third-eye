package admin_wrapper

import (
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/wrappers"
)

type AdminWrapper struct {
	*wrappers.SyncWrapper
	nameToAddr map[string][]string
}

func NewAdminWrapper() *AdminWrapper {
	w := &AdminWrapper{
		SyncWrapper: wrappers.NewSyncWrapper(ds.AdminWrapper, nil),
		nameToAddr:  make(map[string][]string),
	}
	// not using onBlockChange
	w.ViaDataProcess = ds.ViaLog
	return w
}

func (w *AdminWrapper) AddSyncAdapter(adapter ds.SyncAdapterI) {
	// there can be more than 2 treasuries.
	w.nameToAddr[adapter.GetName()] = append(w.nameToAddr[adapter.GetName()], adapter.GetAddress())
	w.SyncWrapper.AddSyncAdapter(adapter)
}

func (w *AdminWrapper) GetAdapterAddrByName(name string) []string {
	return w.nameToAddr[name]
}
