package acl

import (
	"github.com/Gearbox-protocol/gearscan/core"
	"github.com/Gearbox-protocol/gearscan/ethclient"
	"github.com/ethereum/go-ethereum/core/types"
)

type ACL struct {
	*core.SyncAdapter
	*core.State
}

func NewACL(addr string, client *ethclient.Client, repo core.RepositoryI, discoveredAt int64) *ACL {
	obj := &ACL{
		SyncAdapter: &core.SyncAdapter{
			Type:    "ACL",
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

func (mdl *ACL) OnLog(txLog types.Log) {
}
