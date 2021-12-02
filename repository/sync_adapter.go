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
		adapter.Repo = repo
		repo.AddSyncAdapter(repo.prepareSyncAdapter(adapter))
	}
}

func (repo *Repository) prepareSyncAdapter(adapter *core.SyncAdapter) core.SyncAdapterI {
	switch adapter.ContractName {
	case "ACL":
		return acl.NewACLFromAdapter(adapter)
	case "AddressProvider":
		ap := address_provider.NewAddressProviderFromAdapter(adapter)
		for k, dcAddr := range ap.Details {
			blockNum, err := strconv.ParseInt(k, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			repo.AddDataCompressor(blockNum, dcAddr)
		}
		return ap
	case "AccountFactory":
		return account_factory.NewAccountFactoryFromAdapter(adapter)
	case "Pool":
		return pool.NewPoolFromAdapter(adapter)
	case "CreditManager":
		return credit_manager.NewCreditManagerFromAdapter(adapter)
	case "PriceOracle":
		return price_oracle.NewPriceOracleFromAdapter(adapter)
	case "PriceFeed":
		return price_feed.NewPriceFeedFromAdapter(adapter)
	case "ContractRegister":
		return contract_register.NewContractRegisterFromAdapter(adapter)
	case "CreditFilter":
		return credit_filter.NewCreditFilterFromAdapter(adapter)
	}
	return nil
}

func (repo *Repository) AddSyncAdapter(adapterI core.SyncAdapterI) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.kit.Add(adapterI)
}
