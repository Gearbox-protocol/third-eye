package address_provider

import (
	"github.com/Gearbox-protocol/gearscan/models/sync_adapter"
)

type AddressProvider struct {
	sync_adapter.SyncAdapter
}

func NewAddressProvider() sync_adapter.SyncAdapterI {
	return &AddressProvider{

	}
}

func (mdl *AddressProvider) OnLog(){

}

func (mdl *AddressProvider) Sync(){

}

func (mdl *AddressProvider) LoadState(){

}