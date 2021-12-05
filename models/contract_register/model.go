package contract_register

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/ethclient"
)

type ContractRegister struct {
	*core.SyncAdapter
}

func NewContractRegister(addr string, discoveredAt int64, client *ethclient.Client, repo core.RepositoryI) *ContractRegister {
	return NewContractRegisterFromAdapter(
		core.NewSyncAdapter(addr, core.ContractRegister, discoveredAt, client, repo),
	)
}

func NewContractRegisterFromAdapter(adapter *core.SyncAdapter) *ContractRegister {
	obj := &ContractRegister{
		SyncAdapter: adapter,
	}
	return obj
}
