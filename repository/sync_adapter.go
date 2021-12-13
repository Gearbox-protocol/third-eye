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
	"github.com/Gearbox-protocol/third-eye/models/pool"
	"github.com/Gearbox-protocol/third-eye/models/price_oracle"
	"github.com/Gearbox-protocol/third-eye/models/yearn_price_feed"
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
	case core.ACL:
		return acl.NewACLFromAdapter(adapter)
	case core.AddressProvider:
		ap := address_provider.NewAddressProviderFromAdapter(adapter)
		if ap.Details["dc"] != nil {
			dcMap, ok := (ap.Details["dc"]).(map[string]interface{})
			if !ok {
				log.Fatal("Converting address provider() details for dc to map failed", ap.GetAddress())
			}
			for k, dcAddr := range dcMap {
				blockNum, err := strconv.ParseInt(k, 10, 64)
				if err != nil {
					log.Fatal(err)
				}
				repo.AddDataCompressor(blockNum, dcAddr.(string))
			}
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
	case core.CreditFilter:
		return credit_filter.NewCreditFilterFromAdapter(adapter)
	}
	return nil
}

func (repo *Repository) AddSyncAdapter(adapterI core.SyncAdapterI) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.kit.Add(adapterI)
}
