package account_factory

import (
	"github.com/Gearbox-protocol/third-eye/artifacts/accountFactory"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/models/account_manager"
	"github.com/ethereum/go-ethereum/common"
)

type AccountFactory struct {
	*core.SyncAdapter
	contractETH *accountFactory.AccountFactory
}

func NewAccountFactory(addr string, discoveredAt int64, client ethclient.ClientI, repo core.RepositoryI) *AccountFactory {
	adapter := account_manager.NewAccountManager(common.Address{}.Hex(), discoveredAt, client, repo)
	repo.AddSyncAdapter(adapter)
	return NewAccountFactoryFromAdapter(
		core.NewSyncAdapter(addr, core.AccountFactory, discoveredAt, client, repo),
	)
}

func NewAccountFactoryFromAdapter(adapter *core.SyncAdapter) *AccountFactory {
	contractETH, err := accountFactory.NewAccountFactory(common.HexToAddress(adapter.Address), adapter.Client)
	if err != nil {
		log.Fatal(err)
	}
	obj := &AccountFactory{
		SyncAdapter: adapter,
		contractETH: contractETH,
	}
	return obj
}
