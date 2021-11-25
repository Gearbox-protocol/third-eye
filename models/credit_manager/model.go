package credit_manager

import (
	"github.com/Gearbox-protocol/gearscan/ethclient"
	"github.com/Gearbox-protocol/gearscan/core"
)

type CreditManager struct {
	*core.SyncAdapter
	*core.State
}

func NewCreditManager(addr string, client *ethclient.Client, repo core.RepositoryI, discoveredAt int64) *CreditManager {
	obj := &CreditManager{
		SyncAdapter: &core.SyncAdapter{
			Type: "CreditManager",
			Address: addr,
			Client: client,
		},
		State: &core.State{Repo: repo},
	}
	firstDetection := obj.DiscoverFirstLog()
	obj.SyncAdapter.DiscoveredAt = discoveredAt
	obj.SyncAdapter.FirstLogAt = firstDetection
	obj.SyncAdapter.LastSync = firstDetection
	return obj
}