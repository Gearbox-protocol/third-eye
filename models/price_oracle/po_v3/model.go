package po_v3

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/third-eye/ds"
)

type PriceOracle struct {
	*ds.SyncAdapter
}

func NewPriceOracle(addr string, discoveredAt int64, client core.ClientI, repo ds.RepositoryI) *PriceOracle {
	return NewPriceOracleFromAdapter(
		ds.NewSyncAdapter(addr, ds.PriceOracle, discoveredAt, client, repo),
	)
}

func NewPriceOracleFromAdapter(adapter *ds.SyncAdapter) *PriceOracle {
	obj := &PriceOracle{
		SyncAdapter: adapter,
	}
	return obj
}
