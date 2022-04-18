package account_factory

import (
	"github.com/Gearbox-protocol/sdk-go/artifacts/accountFactory"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/account_manager"
	"github.com/ethereum/go-ethereum/common"
)

type AccountFactory struct {
	*ds.SyncAdapter
	contractETH *accountFactory.AccountFactory
}

func NewAccountFactory(addr string, discoveredAt int64, client core.ClientI, repo ds.RepositoryI) *AccountFactory {
	adapter := account_manager.NewAccountManager(common.Address{}.Hex(), discoveredAt, client, repo)
	repo.AddSyncAdapter(adapter)
	return NewAccountFactoryFromAdapter(
		ds.NewSyncAdapter(addr, ds.AccountFactory, discoveredAt, client, repo),
	)
}

func NewAccountFactoryFromAdapter(adapter *ds.SyncAdapter) *AccountFactory {
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
