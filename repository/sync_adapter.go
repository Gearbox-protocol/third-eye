package repository

import (
	"github.com/Gearbox-protocol/gearscan/core"
	"github.com/Gearbox-protocol/gearscan/log"
	"github.com/Gearbox-protocol/gearscan/models/account_factory"
	"github.com/Gearbox-protocol/gearscan/models/acl"
	"github.com/Gearbox-protocol/gearscan/models/address_provider"
	"github.com/Gearbox-protocol/gearscan/models/contract_register"
	"github.com/Gearbox-protocol/gearscan/models/credit_manager"
	"github.com/Gearbox-protocol/gearscan/models/pool"
	"github.com/Gearbox-protocol/gearscan/models/price_oracle"
	"github.com/Gearbox-protocol/gearscan/models/price_feed"
)

func (repo *Repository) loadSyncAdapters() {
	data := []*core.SyncAdapter{}
	err := repo.db.Find(&data, "disabled = ?", false).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, adapter := range data {
		adapter.Client = repo.client
		repo.AddSyncAdapter(prepareSyncAdapter(adapter, repo))
	}
}



func prepareSyncAdapter(adapter *core.SyncAdapter, repo core.RepositoryI) core.SyncAdapterI {
	switch adapter.ContractName {
	case "ACL":
		return acl.NewACLFromAdapter(repo, adapter)
	case "AddressProvider":
		return address_provider.NewAddressProviderFromAdapter(repo,adapter)
	case "AccountFactory":
		return account_factory.NewAccountFactoryFromAdapter(repo, adapter)
	case "Pool":
		return pool.NewPoolFromAdapter(repo, adapter)
	case "CreditManager":
		return credit_manager.NewCreditManagerFromAdapter(repo, adapter)
	case "PriceOracle":
		return price_oracle.NewPriceOracleFromAdapter(repo, adapter)
	case "PriceFeed":
		return price_feed.NewPriceFeedFromAdapter(repo, adapter)
	// case "DataCompressor":
	// 	return data_compressor.NewDataCompressorFromAdapter(repo, adapter)
	case "ContractRegister":
		return contract_register.NewContractRegisterFromAdapter(repo, adapter)
	}
	return nil
}

func (repo *Repository) GetSyncAdapters() []core.SyncAdapterI {
	return repo.syncAdapters
}

func (repo *Repository) AddSyncAdapter(adapterI core.SyncAdapterI) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.syncAdapters = append(repo.syncAdapters, adapterI)
}