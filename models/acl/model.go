package acl

import (
	"github.com/Gearbox-protocol/third-eye/artifacts/aCL"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/ethereum/go-ethereum/common"
)

type ACL struct {
	*core.SyncAdapter
	contractETH *aCL.ACL
}

func NewACL(addr string, discoveredAt int64, client ethclient.ClientI, repo core.RepositoryI) *ACL {
	return NewACLFromAdapter(
		core.NewSyncAdapter(addr, core.ACL, discoveredAt, client, repo),
	)
}

func NewACLFromAdapter(adapter *core.SyncAdapter) *ACL {
	contractETH, err := aCL.NewACL(common.HexToAddress(adapter.Address), adapter.Client)
	log.CheckFatal(err)
	obj := &ACL{
		SyncAdapter: adapter,
		contractETH: contractETH,
	}
	return obj
}
