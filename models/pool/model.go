package pool

import (
	"github.com/Gearbox-protocol/gearscan/ethclient"
	"github.com/Gearbox-protocol/gearscan/core"
	"github.com/ethereum/go-ethereum/core/types"
)

type Pool struct {
	*core.SyncAdapter
	*core.State
}

func NewPool(addr string, client *ethclient.Client, repo core.RepositoryI, discoveredAt int64) *Pool {
	obj := &Pool{
		SyncAdapter: &core.SyncAdapter{
			Type: "Pool",
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

func (mdl *Pool) OnLog(txLog types.Log){
}