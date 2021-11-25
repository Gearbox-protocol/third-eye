package contract_register

import (
	"github.com/Gearbox-protocol/gearscan/core"
	"github.com/Gearbox-protocol/gearscan/ethclient"
)

type ContractRegister struct {
	*core.SyncAdapter
	*core.State
}


func NewContractRegister(addr string, discoveredAt int64, client *ethclient.Client, repo core.RepositoryI) *ContractRegister {
	return NewContractRegisterFromAdapter(
		repo,
		core.NewSyncAdapter(addr, "ContractRegister", discoveredAt, client),
	)
}

func NewContractRegisterFromAdapter(repo core.RepositoryI, adapter *core.SyncAdapter) *ContractRegister {
	obj := &ContractRegister{
		SyncAdapter: adapter,
		State: &core.State{Repo: repo},
	}
	return obj
}