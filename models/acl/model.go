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

func NewACL(addr string, discoveredAt int64, client *ethclient.Client, repo core.RepositoryI) *ACL {
	return NewACLFromAdapter(
		repo,
		core.NewSyncAdapter(addr, "ACL", discoveredAt, client),
	)
}

func NewACLFromAdapter(repo core.RepositoryI, adapter *core.SyncAdapter) *ACL {
	obj := &ACL{
		SyncAdapter: adapter,
		State: &core.State{Repo: repo},
	}
	return obj
}

func (mdl *ACL) OnLog(txLog types.Log) {
}
