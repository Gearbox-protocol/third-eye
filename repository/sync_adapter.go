package repository

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/models/account_factory"
	"github.com/Gearbox-protocol/third-eye/models/acl"
	"github.com/Gearbox-protocol/third-eye/models/address_provider"
	"github.com/Gearbox-protocol/third-eye/models/chainlink_price_feed"
	"github.com/Gearbox-protocol/third-eye/models/contract_register"
	"github.com/Gearbox-protocol/third-eye/models/credit_filter"
	"github.com/Gearbox-protocol/third-eye/models/credit_manager"
	"github.com/Gearbox-protocol/third-eye/models/gear_token"
	"github.com/Gearbox-protocol/third-eye/models/pool"
	"github.com/Gearbox-protocol/third-eye/models/price_oracle"
	"github.com/Gearbox-protocol/third-eye/models/treasury"
	"github.com/Gearbox-protocol/third-eye/models/yearn_price_feed"
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
		repo.addSyncAdapter(repo.prepareSyncAdapter(adapter))
	}
}

func (repo *Repository) prepareSyncAdapter(adapter *core.SyncAdapter) core.SyncAdapterI {
	switch adapter.ContractName {
	case core.ACL:
		return acl.NewACLFromAdapter(adapter)
	case core.AddressProvider:
		ap := address_provider.NewAddressProviderFromAdapter(adapter)
		if ap.Details["dc"] != nil {
			repo.dcWrapper.LoadMultipleDC(ap.Details["dc"])
		}
		if ap.Details["weth"] != nil {
			weth, ok := (ap.Details["weth"]).(string)
			if !ok {
				log.Fatalf("weth is set in addressprovider sync adapter but it is not string %v", ap.Details["weth"])
			}
			repo.SetWETHAddr(weth)
		}
		return ap
	case core.AccountFactory:
		return account_factory.NewAccountFactoryFromAdapter(adapter)
	case core.Pool:
		return pool.NewPoolFromAdapter(adapter)
	case core.CreditManager:
		return credit_manager.NewCreditManagerFromAdapter(adapter)
	case core.PriceOracle:
		return price_oracle.NewPriceOracleFromAdapter(adapter)
	case core.ChainlinkPriceFeed:
		return chainlink_price_feed.NewChainlinkPriceFeedFromAdapter(adapter, false)
	case core.YearnPriceFeed:
		return yearn_price_feed.NewYearnPriceFeedFromAdapter(adapter)
	case core.ContractRegister:
		return contract_register.NewContractRegisterFromAdapter(adapter)
	case core.GearToken:
		return gear_token.NewGearTokenFromAdapter(adapter)
	case core.Treasury:
		return treasury.NewTreasuryFromAdapter(adapter)
	case core.CreditFilter:
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

func (repo *Repository) AddSyncAdapter(adapterI core.SyncAdapterI) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if repo.config.ROLLBACK == "1" {
		return
	}
	repo.addSyncAdapter(adapterI)
}

func (repo *Repository) addSyncAdapter(adapterI core.SyncAdapterI) {
	if core.GearToken == adapterI.GetName() {
		repo.GearTokenAddr = adapterI.GetAddress()
	}
	repo.kit.Add(adapterI)
}
