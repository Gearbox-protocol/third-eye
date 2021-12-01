package account_factory

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/ethereum/go-ethereum/core/types"
)

type AccountFactory struct {
	*core.SyncAdapter
}

func NewAccountFactory(addr string, discoveredAt int64, client *ethclient.Client, repo core.RepositoryI) *AccountFactory {
	return NewAccountFactoryFromAdapter(
		core.NewSyncAdapter(addr, "AccountFactory", discoveredAt, client, repo),
	)
}

func NewAccountFactoryFromAdapter(adapter *core.SyncAdapter) *AccountFactory {
	obj := &AccountFactory{
		SyncAdapter: adapter,
	}
	return obj
}

func (mdl *AccountFactory) OnLog(txLog types.Log) {
}
