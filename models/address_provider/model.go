package address_provider

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/third-eye/ds"
)

type AddressProvider struct {
	*ds.SyncAdapter
}

func NewAddressProvider(addr string, client core.ClientI, repo ds.RepositoryI) *AddressProvider {
	return NewAddressProviderFromAdapter(
		ds.NewSyncAdapter(addr, ds.AddressProvider, -1, client, repo),
	)
}

func NewAddressProviderFromAdapter(adapter *ds.SyncAdapter) *AddressProvider {
	obj := &AddressProvider{
		SyncAdapter: adapter,
	}
	return obj
}
