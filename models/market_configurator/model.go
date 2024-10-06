package market_configurator

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/third-eye/ds"
)

type MarketConfigurator struct {
	*ds.SyncAdapter
}

func NewMarketConfigurator(addr string, discoveredAt int64, client core.ClientI, repo ds.RepositoryI) *MarketConfigurator {
	return NewMarketConfiguratorFromAdapter(
		ds.NewSyncAdapter(addr, ds.ACL, discoveredAt, client, repo),
	)
}

func NewMarketConfiguratorFromAdapter(adapter *ds.SyncAdapter) *MarketConfigurator {
	obj := &MarketConfigurator{
		SyncAdapter: adapter,
	}
	return obj
}
