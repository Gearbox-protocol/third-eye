package address_provider

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/ethclient"
)

type AddressProvider struct {
	*core.SyncAdapter
}

func NewAddressProvider(addr string, client *ethclient.Client, repo core.RepositoryI) *AddressProvider {
	return NewAddressProviderFromAdapter(
		core.NewSyncAdapter(addr, "AddressProvider", -1, client, repo),
	)
}

func NewAddressProviderFromAdapter(adapter *core.SyncAdapter) *AddressProvider {
	obj := &AddressProvider{
		SyncAdapter: adapter,
	}
	return obj
}
