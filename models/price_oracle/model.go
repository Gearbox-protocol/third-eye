package price_oracle

import (
	"github.com/Gearbox-protocol/gearscan/core"
	"github.com/Gearbox-protocol/gearscan/ethclient"
	"github.com/Gearbox-protocol/gearscan/log"
	"github.com/Gearbox-protocol/gearscan/artifacts/priceOracle"
	"github.com/ethereum/go-ethereum/common"
)

type PriceOracle struct {
	*core.SyncAdapter
	*core.State
	contractETH *priceOracle.PriceOracle
}

func NewPriceOracle(addr string, discoveredAt int64, client *ethclient.Client, repo core.RepositoryI) *PriceOracle {
	return NewPriceOracleFromAdapter(
		repo,
		core.NewSyncAdapter(addr, "PriceOracle", discoveredAt, client),
	)
}

func NewPriceOracleFromAdapter(repo core.RepositoryI, adapter *core.SyncAdapter) *PriceOracle {
	cmContract, err := priceOracle.NewPriceOracle(common.HexToAddress(adapter.Address), adapter.Client)
	if err != nil {
		log.Fatal(err)
	}
	obj := &PriceOracle{
		SyncAdapter: adapter,
		State: &core.State{Repo: repo},
		contractETH:cmContract,
	}
	return obj
}

