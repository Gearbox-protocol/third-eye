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

func NewAccountFactory(addr string, discoveredAt int64, client *ethclient.Client, repo core.RepositoryI) *AccountFactory {
	return NewAccountFactoryFromAdapter(
		repo,
		core.NewSyncAdapter(addr, "AccountFactory", discoveredAt, client),
	)
}

func NewAccountFactoryFromAdapter(repo core.RepositoryI, adapter *core.SyncAdapter) *AccountFactory {
	obj := &AccountFactory{
		SyncAdapter: adapter,
		State: &core.State{Repo: repo},
	}
	return obj
}

func (mdl *AccountFactory) OnLog(txLog types.Log) {
}
