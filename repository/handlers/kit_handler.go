package handlers

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/aggregated_block_feed"
	"github.com/Gearbox-protocol/third-eye/models/wrappers/admin_wrapper"
	"github.com/Gearbox-protocol/third-eye/models/wrappers/cf_wrapper"
	"github.com/Gearbox-protocol/third-eye/models/wrappers/pool_wrapper"
)

type AdapterKitHandler struct {
	kit                 *ds.AdapterKit
	aggregatedBlockFeed *aggregated_block_feed.AggregatedBlockFeed
	adminWrapper        *admin_wrapper.AdminWrapper
	cfWrapper           *cf_wrapper.CFWrapper
	poolWrapper         *pool_wrapper.PoolWrapper
}

func NewAdpterKitHandler(client core.ClientI, repo ds.RepositoryI, cfg *config.Config) *AdapterKitHandler {
	obj := &AdapterKitHandler{
		kit:                 ds.NewAdapterKit(),
		aggregatedBlockFeed: aggregated_block_feed.NewAggregatedBlockFeed(client, repo, cfg.Interval),
		cfWrapper:           cf_wrapper.NewCFWrapper(),
		poolWrapper:         pool_wrapper.NewPoolWrapper(client),
		adminWrapper:        admin_wrapper.NewAdminWrapper(client),
	}
	//
	obj.kit.Add(obj.aggregatedBlockFeed)
	obj.kit.Add(obj.adminWrapper)
	obj.kit.Add(obj.cfWrapper)
	obj.kit.Add(obj.poolWrapper)
	//
	return obj
}

// injected in the app itself
// are aggregatedBlockFeed, cfWrapper
//
// whereas adapter with fake address are AccountManager, and PoolLMRewards and CompositeChainlinkPF
func (handler *AdapterKitHandler) addSyncAdapter(adapterI ds.SyncAdapterI) {
	if adapterI.GetName() == ds.QueryPriceFeed {
		handler.aggregatedBlockFeed.AddYearnFeed(adapterI)
		// REVERT ADMIN_WRAPPER
	} else if utils.Contains([]string{
		ds.ContractRegister, ds.ACL,
		ds.AccountFactory, ds.GearToken,
	}, adapterI.GetName()) {
		handler.adminWrapper.AddSyncAdapter(adapterI)
		// REVERT_CF_WRAPPER
	} else if adapterI.GetName() == ds.CreditFilter || adapterI.GetName() == ds.CreditConfigurator {
		handler.cfWrapper.AddSyncAdapter(adapterI)
		// REVERT_POOL_WRAPPER
	} else if adapterI.GetName() == ds.Pool {
		handler.poolWrapper.AddSyncAdapter(adapterI)
	} else {
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
		// REVERT_POOL_WRAPPER
		if adapter := repo.poolWrapper.GetAdapter(addr); adapter != nil {
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

func (repo AdapterKitHandler) getAdapterState() (adapters []*ds.SyncAdapter) {
	for _, adapter := range repo.aggregatedBlockFeed.GetQueryFeeds() {
		adapters = append(adapters, adapter.GetAdapterState()...)
	}
	// REVERT_ADMIN_WRAPPER
	adapters = append(adapters, repo.adminWrapper.GetAdapterState()...)
	// REVERT_CF_WRAPPER
	adapters = append(adapters, repo.cfWrapper.GetAdapterState()...)
	// REVERT_POOL_WRAPPER
	adapters = append(adapters, repo.poolWrapper.GetAdapterState()...)
	return
}

func (repo AdapterKitHandler) GetYearnFeedAddrs() (addrs []string) {
	feeds := repo.aggregatedBlockFeed.GetQueryFeeds()
	addrs = make([]string, 0, len(feeds))
	for _, adapter := range feeds {
		addrs = append(addrs, adapter.GetAddress())
	}
	return
}

// TODO: find eng.repo.GetAdapterAddressByName(ds.CreditManager)
func (repo AdapterKitHandler) GetAdapterAddressByName(name string) []string {
	// REVERT_CM_WRAPPER
	// if name == ds.CreditManager {
	// 	return repo.cmWrapper.GetUnderlyingAdapterAddrs()
	// }
	// REVERT_ADMIN_WRAPPER
	if name == ds.GearToken {
		return repo.adminWrapper.GetAdapterAddrByName(name)
	}
	return repo.kit.GetAdapterAddressByName(name)
}

func (repo AdapterKitHandler) GetAggregatedFeed() *aggregated_block_feed.AggregatedBlockFeed {
	return repo.aggregatedBlockFeed
}
