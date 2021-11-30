package address_provider

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/ethclient"
)

type AddressProvider struct {
	*core.SyncAdapter
	*core.State
}

func NewAddressProvider(addr string, client *ethclient.Client, repo core.RepositoryI) *AddressProvider {
	return NewAddressProviderFromAdapter(
		repo,
		core.NewSyncAdapter(addr, "AddressProvider", -1, client),
	)
}

func NewAddressProviderFromAdapter(repo core.RepositoryI, adapter *core.SyncAdapter) *AddressProvider {
	obj := &AddressProvider{
		SyncAdapter: adapter,
		State:       &core.State{Repo: repo},
	}
	return obj
}
