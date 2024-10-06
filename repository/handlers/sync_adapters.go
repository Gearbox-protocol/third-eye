package handlers

import (
	"fmt"
	"sort"
	"sync"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/account_factory"
	"github.com/Gearbox-protocol/third-eye/models/account_manager"
	"github.com/Gearbox-protocol/third-eye/models/acl"
	"github.com/Gearbox-protocol/third-eye/models/address_provider"
	"github.com/Gearbox-protocol/third-eye/models/aggregated_block_feed"
	"github.com/Gearbox-protocol/third-eye/models/chainlink_price_feed"
	"github.com/Gearbox-protocol/third-eye/models/composite_chainlink"
	"github.com/Gearbox-protocol/third-eye/models/configurators"
	"github.com/Gearbox-protocol/third-eye/models/configurators/credit_filter"
	"github.com/Gearbox-protocol/third-eye/models/contract_register"
	"github.com/Gearbox-protocol/third-eye/models/credit_manager"
	"github.com/Gearbox-protocol/third-eye/models/gear_token"
	"github.com/Gearbox-protocol/third-eye/models/pool"
	lmrewardsv2 "github.com/Gearbox-protocol/third-eye/models/pool_lmrewards/v2"
	lmrewardsv3 "github.com/Gearbox-protocol/third-eye/models/pool_lmrewards/v3"
	"github.com/Gearbox-protocol/third-eye/models/pool_quota_keeper"
	"github.com/Gearbox-protocol/third-eye/models/price_oracle/po_v2"
	"github.com/Gearbox-protocol/third-eye/models/price_oracle/po_v3"
	"github.com/Gearbox-protocol/third-eye/models/rebase_token"
	"github.com/Gearbox-protocol/third-eye/models/treasury"
	"github.com/Gearbox-protocol/third-eye/models/wrappers/pool_wrapper"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SyncAdaptersRepo struct {
	r               ds.RepositoryI
	client          core.ClientI
	extras          *ExtrasRepo
	rollbackAllowed bool
	mu              *sync.Mutex
	*AdapterKitHandler
	cfg *config.Config
}

func NewSyncAdaptersRepo(client core.ClientI, repo ds.RepositoryI, cfg *config.Config, extras *ExtrasRepo) *SyncAdaptersRepo {
	obj := &SyncAdaptersRepo{
		client:            client,
		r:                 repo,
		extras:            extras,
		rollbackAllowed:   cfg.Rollback,
		mu:                &sync.Mutex{},
		AdapterKitHandler: NewAdpterKitHandler(client, repo, cfg),
		cfg:               cfg,
	}
	return obj
}

// load/save
func (repo *SyncAdaptersRepo) LoadSyncAdapters(db *gorm.DB) {
	defer utils.Elapsed("loadSyncAdapters")()
	//
	data := []*ds.SyncAdapter{}
	err := db.Find(&data, "disabled = ? OR disabled is NULL OR type = 'PriceOracle' ORDER BY type", false).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, adapter := range data {
		// adapter.SetDisabled(false)
		p := repo.PrepareSyncAdapter(adapter)
		repo.addSyncAdapter(p)
	}
}

func (repo *SyncAdaptersRepo) Save(tx *gorm.DB) {
	defer utils.Elapsed("sync adapters sql statements")()
	adapters := make([]*ds.SyncAdapter, 0, repo.kit.Len())
	for lvlIndex := 0; lvlIndex < repo.kit.Len(); lvlIndex++ {
		for repo.kit.Next(lvlIndex) {
			adapter := repo.kit.Get(lvlIndex)
			if ds.IsWrapperAdapter(adapter.GetName()) {
				continue
			}
			adapters = append(adapters, adapter.GetAdapterState())
			if adapter.HasUnderlyingStateToSave() {
				err := tx.Clauses(clause.OnConflict{
					UpdateAll: true,
				}).Create(adapter.GetUnderlyingState()).Error
				log.CheckFatal(err)
			}
		}
		repo.kit.Reset(lvlIndex)
	}
	// save wrapper underlying states
	for _, adapter := range repo.AdapterKitHandler.GetAdaptersFromWrapper() {
		if adapter.HasUnderlyingStateToSave() {
			err := tx.Clauses(clause.OnConflict{
				UpdateAll: true,
			}).Create(adapter.GetUnderlyingState()).Error
			log.CheckFatal(err)
		}
		adapters = append(adapters, adapter.GetAdapterState())
	}
	//
	err := tx.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(adapters, 50).Error
	log.CheckFatal(err)

}

// external funcs
// for testing and load from db
func (repo *SyncAdaptersRepo) PrepareSyncAdapter(adapter *ds.SyncAdapter) ds.SyncAdapterI {
	adapter.Client = repo.client
	adapter.Repo = repo.r
	switch adapter.ContractName {
	case ds.ACL:
		return acl.NewACLFromAdapter(adapter)
	case ds.AddressProvider:
		chainId := core.GetChainId(repo.client)
		addrProviders := core.GetAddressProvider(chainId, core.VersionType{})
		ap := address_provider.NewAddressProviderFromAdapter(adapter, addrProviders)
		if ap.Details["dc"] != nil {
			repo.extras.GetDCWrapper().LoadMultipleDC(ap.Details["dc"])
		}
		return ap
	case ds.LMRewardsv2:
		return lmrewardsv2.NewLMRewardsv2FromAdapter(adapter)
	case ds.LMRewardsv3:
		return lmrewardsv3.NewLMRewardsv3FromAdapter(adapter)
	case ds.AccountFactory:
		return account_factory.NewAccountFactoryFromAdapter(adapter)
	case ds.Pool:
		return pool.NewPoolFromAdapter(adapter)
	case ds.PoolQuotaKeeper:
		return pool_quota_keeper.NewPoolQuotaKeeperFromAdapter(adapter)
	case ds.CreditManager:
		return credit_manager.NewCMFromAdapter(adapter)
	case ds.PriceOracle:
		if adapter.GetVersion().LessThan(core.NewVersion(300)) {
			po_v2.NewPriceOracleFromAdapter(adapter)
		} else {
			po_v3.NewPriceOracleFromAdapter(adapter)
		}
	case ds.ChainlinkPriceFeed:
		return chainlink_price_feed.NewChainlinkPriceFeedFromAdapter(adapter, false)
	case ds.CompositeChainlinkPF:
		return composite_chainlink.NewCompositeChainlinkPFFromAdapter(adapter)
	case ds.QueryPriceFeed:
		return aggregated_block_feed.NewQueryPriceFeedFromAdapter(adapter)
	case ds.ContractRegister:
		return contract_register.NewContractRegisterFromAdapter(adapter)
	case ds.GearToken:
		return gear_token.NewGearTokenFromAdapter(adapter)
	case ds.Treasury:
		return treasury.NewTreasuryFromAdapter(adapter)
	case ds.RebaseToken:
		return rebase_token.NewRebaseTokenFromAdapter(adapter)
	case ds.AccountManager:
		return account_manager.NewAccountManagerFromAdapter(adapter)
	case ds.CreditConfigurator:
		return configurators.NewConfiguratorFromAdapter(adapter)
	case ds.CreditFilter:
		if adapter.Details["creditManager"] != nil {
			cmAddr := adapter.Details["creditManager"].(string)
			repo.extras.GetDCWrapper().AddCreditManagerToFilter(cmAddr, adapter.GetAddress())
		} else {
			log.Fatal("Credit filter doesn't have credit manager", adapter.GetAddress())
		}
		return credit_filter.NewCreditFilterFromAdapter(adapter)
	}
	return nil
}

func (repo *SyncAdaptersRepo) AddSyncAdapter(newAdapterI ds.SyncAdapterI) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if repo.rollbackAllowed {
		return
	}
	if repo.GetAdapter(newAdapterI.GetAddress()) != nil {
		return
	}
	if newAdapterI.GetName() == ds.PriceOracle {
		switch  newAdapterI.GetAddress() {
		case "0x6385892aCB085eaa24b745a712C9e682d80FF681": // v2
		oldPriceOracle := repo.GetAdapter("0x0e74a08443c5E39108520589176Ac12EF65AB080") // v1
		oldPriceOracle.SetBlockToDisableOn(newAdapterI.GetDiscoveredAt())
		case "0x599f585D1042A14aAb194AC8031b2048dEFdFB85": // v3
		oldPriceOracle := repo.GetAdapter("0x6385892aCB085eaa24b745a712C9e682d80FF681") // v2
		oldPriceOracle.SetBlockToDisableOn(newAdapterI.GetDiscoveredAt())
		}
	}
	repo.addSyncAdapter(newAdapterI)
}

func (repo *SyncAdaptersRepo) GetKit() *ds.AdapterKit {
	return repo.kit
}
func (repo *SyncAdaptersRepo) GetPoolWrapper() *pool_wrapper.PoolWrapper {
	return repo.poolWrapper
}

////////////////////
// for price oracle
////////////////////

// return the active first oracle under blockNum
// if all disabled return the last one
// blockNum ==0 is latest
func (repo *SyncAdaptersRepo) GetActivePriceOracleByBlockNum(blockNum int64) (latestOracle string, version core.VersionType, err error) {
	oracles := repo.kit.GetAdapterAddressByName(ds.PriceOracle)
	data :=make([]ds.SyncAdapterI, 0, len(oracles))
	for _, addr := range oracles {
		oracleAdapter := repo.GetAdapter(addr)
		if blockNum >= oracleAdapter.GetDiscoveredAt() || blockNum ==0  {
			data = append(data, oracleAdapter)
		}
	}
	sort.Slice(data, func (a,b int) bool {
		return data[a].GetDiscoveredAt() > data[b].GetDiscoveredAt() 
	})
	//
	err = fmt.Errorf("not found")
	var ans ds.SyncAdapterI
	//
	for _, e:= range data {
		if ans != nil && ans.GetVersion() != e.GetVersion() {
			break
		}
		ans = e
		err = nil
	}
	latestOracle= ans.GetAddress()
	version = ans.GetVersion()
	return
}
