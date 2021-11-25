package account_factory

import (
	"github.com/Gearbox-protocol/gearscan/core"
	"github.com/Gearbox-protocol/gearscan/ethclient"
	"github.com/ethereum/go-ethereum/core/types"
)

type AccountFactory struct {
	*core.SyncAdapter
	*core.State
}

func NewAccountFactory(addr string, client *ethclient.Client, repo core.RepositoryI, discoveredAt int64) *AccountFactory {
	obj := &AccountFactory{
		SyncAdapter: &core.SyncAdapter{
			Type:    "AccountFactory",
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

func (mdl *AccountFactory) OnLog(txLog types.Log) {
}
