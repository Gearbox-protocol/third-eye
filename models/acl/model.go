package acl

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/ethereum/go-ethereum/core/types"
)

type ACL struct {
	*core.SyncAdapter
}

func NewACL(addr string, discoveredAt int64, client *ethclient.Client, repo core.RepositoryI) *ACL {
	return NewACLFromAdapter(
		core.NewSyncAdapter(addr, core.ACL, discoveredAt, client, repo),
	)
}

func NewACLFromAdapter(adapter *core.SyncAdapter) *ACL {
	obj := &ACL{
		SyncAdapter: adapter,
	}
	return obj
}

func (mdl *ACL) OnLog(txLog types.Log) {
}
