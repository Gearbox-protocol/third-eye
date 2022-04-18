package price_oracle

import (
	"github.com/Gearbox-protocol/sdk-go/artifacts/priceOracle"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/common"
)

type PriceOracle struct {
	*ds.SyncAdapter
	contractETH *priceOracle.PriceOracle
}

func NewPriceOracle(addr string, discoveredAt int64, client core.ClientI, repo ds.RepositoryI) *PriceOracle {
	return NewPriceOracleFromAdapter(
		ds.NewSyncAdapter(addr, ds.PriceOracle, discoveredAt, client, repo),
	)
}

func NewPriceOracleFromAdapter(adapter *ds.SyncAdapter) *PriceOracle {
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
