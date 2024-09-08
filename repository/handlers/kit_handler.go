package handlers

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/aggregated_block_feed"
	"github.com/Gearbox-protocol/third-eye/models/wrappers/admin_wrapper"
	"github.com/Gearbox-protocol/third-eye/models/wrappers/cf_wrapper"
	"github.com/Gearbox-protocol/third-eye/models/wrappers/cm_wrapper"
	"github.com/Gearbox-protocol/third-eye/models/wrappers/pool_quota_wrapper"
	"github.com/Gearbox-protocol/third-eye/models/wrappers/pool_wrapper"
)

type AdapterKitHandler struct {
	kit                 *ds.AdapterKit
	aggregatedBlockFeed *aggregated_block_feed.AQFWrapper
	adminWrapper        *admin_wrapper.AdminWrapper
	cfWrapper           *cf_wrapper.CFWrapper
	cmWrapper           *cm_wrapper.CMWrapper
	poolWrapper         *pool_wrapper.PoolWrapper
	// v3
	poolQuotaWrapper *pool_quota_wrapper.PoolQuotaWrapper
}

func NewAdpterKitHandler(client core.ClientI, repo ds.RepositoryI, cfg *config.Config) *AdapterKitHandler {
	obj := &AdapterKitHandler{
		kit:                 ds.NewAdapterKit(),
		aggregatedBlockFeed: aggregated_block_feed.NewAQFWrapper(client, repo, cfg.Interval),
		adminWrapper:        admin_wrapper.NewAdminWrapper(),
		cfWrapper:           cf_wrapper.NewCFWrapper(),
		poolWrapper:         pool_wrapper.NewPoolWrapper(client),
		cmWrapper:           cm_wrapper.NewCMWrapper(client),
		// v3
		poolQuotaWrapper: pool_quota_wrapper.NewPoolQuotaWrapper(client),
	}
	//
	obj.kit.Add(obj.aggregatedBlockFeed)
	obj.kit.Add(obj.adminWrapper)
	obj.kit.Add(obj.cfWrapper)
	obj.kit.Add(obj.cmWrapper)
	obj.kit.Add(obj.poolWrapper)
	obj.kit.Add(obj.poolQuotaWrapper)
	//
	return obj
}

// injected in the app itself
// are aggregatedBlockFeed, cfWrapper
//
// whereas adapter with fake address are AccountManager, and LMRewardsv2 and CompositeChainlinkPF
func (handler *AdapterKitHandler) addSyncAdapter(adapterI ds.SyncAdapterI) {
	switch adapterI.GetName() {
	case ds.QueryPriceFeed:
		handler.aggregatedBlockFeed.AddQueryPriceFeed(aggregated_block_feed.FromAdapter(adapterI))
		// REVERT_ADMIN_WRAPPER
	case ds.ContractRegister, ds.ACL, ds.AccountFactory, ds.GearToken:
		handler.adminWrapper.AddSyncAdapter(adapterI)
		// REVERT_CF_WRAPPER
	case ds.CreditFilter, ds.CreditConfigurator:
		handler.cfWrapper.AddSyncAdapter(adapterI)
		// REVERT_CM_WRAPPER
	case ds.CreditManager:
		handler.cmWrapper.AddSyncAdapter(adapterI)
		// REVERT_POOL_WRAPPER
	case ds.Pool:
		handler.poolWrapper.AddSyncAdapter(adapterI)
		// v3
	case ds.PoolQuotaKeeper:
		handler.poolQuotaWrapper.AddSyncAdapter(adapterI)
	default:
		handler.kit.Add(adapterI)
	}
}

func (repo *AdapterKitHandler) GetAdapter(addr string) ds.SyncAdapterI {
	adapter := repo.kit.GetAdapter(addr)
	if adapter == nil {
		// REVERT_ADMIN_WRAPPER
		if adapter := repo.adminWrapper.GetAdapter(addr); adapter != nil {
			return adapter
		}
		// REVERT_CF_WRAPPER
		if adapter := repo.cfWrapper.GetAdapter(addr); adapter != nil {
			return adapter
		}
		// REVERT_CM_WRAPPER
		if adapter := repo.cmWrapper.GetAdapter(addr); adapter != nil {
			return adapter
		}
		// REVERT_POOL_WRAPPER
		if adapter := repo.poolWrapper.GetAdapter(addr); adapter != nil {
			return adapter
		}
		// v3
		if adapter := repo.poolQuotaWrapper.GetAdapter(addr); adapter != nil {
			return adapter
		}
		// check if adapter is under aggregated block feed
		feeds := repo.aggregatedBlockFeed.GetQueryFeeds()
		for _, feed := range feeds {
			if feed.GetAddress() == addr {
				return feed
			}
		}
	}
	return adapter
}

func (repo AdapterKitHandler) GetAdaptersFromWrapper() (adapters []ds.SyncAdapterI) {
	for _, adapter := range repo.aggregatedBlockFeed.GetQueryFeeds() {
		adapters = append(adapters, adapter)
	}
	// REVERT_ADMIN_WRAPPER
	adapters = append(adapters, repo.adminWrapper.GetAdapters()...)
	// REVERT_CF_WRAPPER
	adapters = append(adapters, repo.cfWrapper.GetAdapters()...)
	// REVERT_CM_WRAPPER
	adapters = append(adapters, repo.cmWrapper.GetAdapters()...)
	// REVERT_POOL_WRAPPER
	adapters = append(adapters, repo.poolWrapper.GetAdapters()...)
	//
	adapters = append(adapters, repo.poolQuotaWrapper.GetAdapters()...)
	return
}

func (repo AdapterKitHandler) GetRetryFeedForDebts() (addrs []string) {
	feeds := repo.aggregatedBlockFeed.GetQueryFeeds()
	addrs = make([]string, 0, len(feeds))
	for _, adapter := range feeds {
		// ds.RedStonePF
		// ds.CompositeRedStonePF

		if utils.Contains([]string{ds.CurvePF,
			ds.YearnPF,
			ds.SingleAssetPF}, adapter.GetPFType()) {
			addrs = append(addrs, adapter.GetAddress())
		}
	}
	return
}

// TODO: find eng.repo.GetAdapterAddressByName(ds.CreditManager)
func (repo AdapterKitHandler) GetAdapterAddressByName(name string) []string {
	// REVERT_CM_WRAPPER
	if name == ds.CreditManager {
		return repo.cmWrapper.GetUnderlyingAdapterAddrs()
	}
	// REVERT_ADMIN_WRAPPER
	if utils.Contains([]string{ds.ContractRegister, ds.ACL, ds.AccountFactory,
		ds.GearToken}, name) {
		return repo.adminWrapper.GetAdapterAddrByName(name)
	}
	return repo.kit.GetAdapterAddressByName(name)
}

func (repo AdapterKitHandler) GetAggregatedFeed() *aggregated_block_feed.AQFWrapper {
	return repo.aggregatedBlockFeed
}
