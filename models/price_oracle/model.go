package price_oracle

import (
	"github.com/Gearbox-protocol/third-eye/artifacts/priceOracle"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/ethereum/go-ethereum/common"
)

type PriceOracle struct {
	*core.SyncAdapter
	contractETH *priceOracle.PriceOracle
}

func NewPriceOracle(addr string, discoveredAt int64, client *ethclient.Client, repo core.RepositoryI) *PriceOracle {
	return NewPriceOracleFromAdapter(
		core.NewSyncAdapter(addr, core.PriceOracle, discoveredAt, client, repo),
	)
}

func NewPriceOracleFromAdapter(adapter *core.SyncAdapter) *PriceOracle {
	cmContract, err := priceOracle.NewPriceOracle(common.HexToAddress(adapter.Address), adapter.Client)
	if err != nil {
		log.Fatal(err)
	}
	obj := &PriceOracle{
		SyncAdapter: adapter,
		contractETH: cmContract,
	}
	return obj
}
