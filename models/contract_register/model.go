package contract_register

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/third-eye/ds"
)

type ContractRegister struct {
	*ds.SyncAdapter
}

func NewContractRegister(addr string, discoveredAt int64, client core.ClientI, repo ds.RepositoryI) *ContractRegister {
	return NewContractRegisterFromAdapter(
		ds.NewSyncAdapter(addr, ds.ContractRegister, discoveredAt, client, repo),
	)
}

func NewContractRegisterFromAdapter(adapter *ds.SyncAdapter) *ContractRegister {
	obj := &ContractRegister{
		SyncAdapter: adapter,
	}
	return obj
}
