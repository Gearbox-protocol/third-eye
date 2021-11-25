package price_oracle

import (
	"github.com/Gearbox-protocol/gearscan/core"
	"github.com/Gearbox-protocol/gearscan/ethclient"
)

type PriceOracle struct {
	*core.SyncAdapter
	*core.State
}

func NewPriceOracle(addr string, client *ethclient.Client, repo core.RepositoryI, discoveredAt int64) *PriceOracle {
	obj := &PriceOracle{
		SyncAdapter: &core.SyncAdapter{
			Type:    "PriceOracle",
			Address: addr,
			Client:  client,
		},
		State: &core.State{Repo: repo},
	}
	firstDetection := obj.DiscoverFirstLog()
	obj.SyncAdapter.DiscoveredAt = discoveredAt
	obj.SyncAdapter.FirstLogAt = firstDetection
	obj.SyncAdapter.LastSync = firstDetection
	return obj
}
