package credit_manager

import (
	"github.com/Gearbox-protocol/gearscan/core"
	"github.com/Gearbox-protocol/gearscan/ethclient"
)

type CreditManager struct {
	*core.SyncAdapter
	*core.State
}

func NewCreditManager(addr string, client *ethclient.Client, repo core.RepositoryI, discoveredAt int64) *CreditManager {
	obj := &CreditManager{
		SyncAdapter: &core.SyncAdapter{
			Type:    "CreditManager",
			Address: addr,
			Client:  client,
		},
		State: &core.State{Repo: repo},
	}
	firstDetection := obj.DiscoverFirstLog()
	obj.SyncAdapter.DiscoveredAt = discoveredAt
	obj.SyncAdapter.FirstLogAt = firstDetection
	obj.SyncAdapter.LastSync = firstDetection
	repo.AddCreditManager(&core.CreditManager{
		Sessions: core.NewHstore(),
		Address:  addr,
	})
	return obj
}
