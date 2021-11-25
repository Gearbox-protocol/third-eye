package price_oracle

import (
	"github.com/Gearbox-protocol/gearscan/core"
	"github.com/Gearbox-protocol/gearscan/ethclient"
)

type PriceOracle struct {
	*core.SyncAdapter
	*core.State
}

func NewPriceOracle(addr string, discoveredAt int64, client *ethclient.Client, repo core.RepositoryI) *PriceOracle {
	return NewPriceOracleFromAdapter(
		repo,
		core.NewSyncAdapter(addr, "PriceOracle", discoveredAt, client),
	)
}

func NewPriceOracleFromAdapter(repo core.RepositoryI, adapter *core.SyncAdapter) *PriceOracle {
	obj := &PriceOracle{
		SyncAdapter: adapter,
		State: &core.State{Repo: repo},
	}
	return obj
}

