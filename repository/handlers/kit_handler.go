package handlers

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/aggregated_block_feed"
	"github.com/Gearbox-protocol/third-eye/models/wrappers/cf_wrapper"
)

type AdapterKitHandler struct {
	kit                 *ds.AdapterKit
	aggregatedBlockFeed *aggregated_block_feed.AggregatedBlockFeed
	cfWrapper           *cf_wrapper.CFWrapper
}

func NewAdpterKitHandler(client core.ClientI, repo ds.RepositoryI, cfg *config.Config) *AdapterKitHandler {
	obj := &AdapterKitHandler{
		kit:                 ds.NewAdapterKit(),
		aggregatedBlockFeed: aggregated_block_feed.NewAggregatedBlockFeed(client, repo, cfg.Interval),
		cfWrapper:           cf_wrapper.NewCFWrapper(),
	}
	//
	obj.kit.Add(obj.aggregatedBlockFeed)
	obj.kit.Add(obj.cfWrapper)
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
	} else if adapterI.GetName() == ds.CreditFilter || adapterI.GetName() == ds.CreditConfigurator {
		handler.cfWrapper.AddSyncAdapter(adapterI)
	} else {
		handler.kit.Add(adapterI)
	}
}

func (repo *AdapterKitHandler) GetAdapter(addr string) ds.SyncAdapterI {
	adapter := repo.kit.GetAdapter(addr)
	if adapter == nil {
		// check if the adapter is under credit filter wrapper
		if adapter := repo.cfWrapper.GetAdapter(addr); adapter != nil {
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
	// if name == ds.CMWrapper {
	// 	return repo.cmWrapper.GetUnderlyingAdapterAddrs()
	// }
	return repo.kit.GetAdapterAddressByName(name)
}

func (repo AdapterKitHandler) getAdapterState() (adapters []*ds.SyncAdapter) {
	for _, adapter := range repo.aggregatedBlockFeed.GetQueryFeeds() {
		adapters = append(adapters, adapter.GetAdapterState()...)
	}
	adapters = append(adapters, repo.cfWrapper.GetAdapterState()...)
	return
}

func (repo AdapterKitHandler) GetAggregatedFeed() *aggregated_block_feed.AggregatedBlockFeed {
	return repo.aggregatedBlockFeed
}
