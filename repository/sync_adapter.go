package repository

import (
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/account_factory"
	"github.com/Gearbox-protocol/third-eye/models/account_manager"
	"github.com/Gearbox-protocol/third-eye/models/acl"
	"github.com/Gearbox-protocol/third-eye/models/address_provider"
	"github.com/Gearbox-protocol/third-eye/models/aggregated_block_feed"
	"github.com/Gearbox-protocol/third-eye/models/chainlink_price_feed"
	"github.com/Gearbox-protocol/third-eye/models/contract_register"
	"github.com/Gearbox-protocol/third-eye/models/credit_filter"
	"github.com/Gearbox-protocol/third-eye/models/credit_manager"
	"github.com/Gearbox-protocol/third-eye/models/gear_token"
	"github.com/Gearbox-protocol/third-eye/models/pool"
	"github.com/Gearbox-protocol/third-eye/models/price_oracle"
	"github.com/Gearbox-protocol/third-eye/models/treasury"
)

func (repo *Repository) loadSyncAdapters() {
	defer utils.Elapsed("loadSyncAdapters")()
	//
	data := []*ds.SyncAdapter{}
	err := repo.db.Find(&data, "disabled = ? OR type = 'PriceOracle'", false).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, adapter := range data {
		repo.addSyncAdapter(repo.PrepareSyncAdapter(adapter))
	}
}

func (repo *Repository) PrepareSyncAdapter(adapter *ds.SyncAdapter) ds.SyncAdapterI {
	adapter.Client = repo.client
	adapter.Repo = repo
	switch adapter.ContractName {
	case ds.ACL:
		return acl.NewACLFromAdapter(adapter)
	case ds.AddressProvider:
		ap := address_provider.NewAddressProviderFromAdapter(adapter)
		if ap.Details["dc"] != nil {
			repo.dcWrapper.LoadMultipleDC(ap.Details["dc"])
		}
		return ap
	case ds.AccountFactory:
		return account_factory.NewAccountFactoryFromAdapter(adapter)
	case ds.Pool:
		return pool.NewPoolFromAdapter(adapter)
	case ds.CreditManager:
		return credit_manager.NewCreditManagerFromAdapter(adapter)
	case ds.PriceOracle:
		return price_oracle.NewPriceOracleFromAdapter(adapter)
	case ds.ChainlinkPriceFeed:
		return chainlink_price_feed.NewChainlinkPriceFeedFromAdapter(adapter, false)
	case ds.YearnPriceFeed:
		return aggregated_block_feed.NewYearnPriceFeedFromAdapter(adapter)
	case ds.ContractRegister:
		return contract_register.NewContractRegisterFromAdapter(adapter)
	case ds.GearToken:
		return gear_token.NewGearTokenFromAdapter(adapter)
	case ds.Treasury:
		return treasury.NewTreasuryFromAdapter(adapter)
	case ds.AccountManager:
		return account_manager.NewAccountManagerFromAdapter(adapter)
	case ds.CreditConfigurator:
		return credit_filter.NewCreditFilterFromAdapter(adapter)
	case ds.CreditFilter:
		if adapter.Details["creditManager"] != nil {
			cmAddr := adapter.Details["creditManager"].(string)
			repo.AddCreditManagerToFilter(cmAddr, adapter.GetAddress())
		} else {
			log.Fatal("Credit filter doesn't have credit manager", adapter.GetAddress())
		}
		return credit_filter.NewCreditFilterFromAdapter(adapter)
	}
	return nil
}

func (repo *Repository) AddSyncAdapter(newAdapterI ds.SyncAdapterI) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if repo.config.ROLLBACK == "1" {
		return
	}
	if newAdapterI.GetName() == ds.PriceOracle {
		oldPriceOracleAddrs := repo.kit.GetAdapterAddressByName(ds.PriceOracle)
		for _, addr := range oldPriceOracleAddrs {
			oldPriceOracle := repo.kit.GetAdapter(addr)
			if !oldPriceOracle.IsDisabled() {
				oldPriceOracle.SetBlockToDisableOn(newAdapterI.GetDiscoveredAt())
			}
		}
	}
	repo.addSyncAdapter(newAdapterI)
}

func (repo *Repository) addSyncAdapter(adapterI ds.SyncAdapterI) {
	if ds.GearToken == adapterI.GetName() {
		repo.GearTokenAddr = adapterI.GetAddress()
	}
	if adapterI.GetName() == ds.YearnPriceFeed {
		repo.aggregatedFeed.AddYearnFeed(adapterI)
	} else {
		repo.kit.Add(adapterI)
	}
}
