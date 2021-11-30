package repository

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/models/account_factory"
	"github.com/Gearbox-protocol/third-eye/models/acl"
	"github.com/Gearbox-protocol/third-eye/models/address_provider"
	"github.com/Gearbox-protocol/third-eye/models/contract_register"
	"github.com/Gearbox-protocol/third-eye/models/credit_filter"
	"github.com/Gearbox-protocol/third-eye/models/credit_manager"
	"github.com/Gearbox-protocol/third-eye/models/pool"
	"github.com/Gearbox-protocol/third-eye/models/price_feed"
	"github.com/Gearbox-protocol/third-eye/models/price_oracle"
	"strconv"
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
		ap := address_provider.NewAddressProviderFromAdapter(repo, adapter)
		for k, dcAddr := range ap.Details  {
			blockNum, err := strconv.ParseInt(k, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			repo.AddDataCompressor(blockNum, dcAddr)
		}
		return ap
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
	case "CreditFilter":
		return credit_filter.NewCreditFilterFromAdapter(repo, adapter)
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

func (repo *Repository) DisableSyncAdapter(addr string) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	for _, adapter := range repo.syncAdapters {
		if adapter.GetAddress() == addr {
			adapter.Disable()
		}
	}
}
