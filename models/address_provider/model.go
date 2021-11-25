package address_provider

import (
	"github.com/Gearbox-protocol/gearscan/core"
	"github.com/Gearbox-protocol/gearscan/ethclient"
)

type AddressProvider struct {
	*core.SyncAdapter
	*core.State
}

func NewAddressProvider(addr string, client *ethclient.Client, repo core.RepositoryI) *AddressProvider {
	obj := &AddressProvider{
		SyncAdapter: &core.SyncAdapter{
			Type:    "AddressProvider",
			Address: addr,
			Client:  client,
		},
		State: &core.State{Repo: repo},
	}
	firstDetection := obj.DiscoverFirstLog()
	obj.SyncAdapter.DiscoveredAt = firstDetection
	obj.SyncAdapter.FirstLogAt = firstDetection
	obj.SyncAdapter.LastSync = firstDetection
	return obj
}
