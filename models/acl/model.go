package acl

import (
	"github.com/Gearbox-protocol/sdk-go/artifacts/aCL"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/common"
)

type ACL struct {
	*ds.SyncAdapter
	contractETH *aCL.ACL
}

func NewACL(addr string, discoveredAt int64, client core.ClientI, repo ds.RepositoryI) *ACL {
	return NewACLFromAdapter(
		ds.NewSyncAdapter(addr, ds.ACL, discoveredAt, client, repo),
	)
}

func NewACLFromAdapter(adapter *ds.SyncAdapter) *ACL {
	contractETH, err := aCL.NewACL(common.HexToAddress(adapter.Address), adapter.Client)
	log.CheckFatal(err)
	obj := &ACL{
		SyncAdapter: adapter,
		contractETH: contractETH,
	}
	return obj
}
