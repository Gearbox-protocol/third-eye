package credit_filter

import (
	"github.com/Gearbox-protocol/third-eye/artifacts/creditFilter"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/ethereum/go-ethereum/common"
)

type CreditFilter struct {
	*core.SyncAdapter
	contractETH *creditFilter.CreditFilter
}

func NewCreditFilter(addr, creditManager string, discoveredAt int64, client ethclient.ClientI, repo core.RepositoryI) *CreditFilter {
	syncAdapter := core.NewSyncAdapter(addr, core.CreditFilter, discoveredAt, client, repo)
	syncAdapter.Details = map[string]interface{}{"creditManager": creditManager}
	return NewCreditFilterFromAdapter(
		syncAdapter,
	)
}

func NewCreditFilterFromAdapter(adapter *core.SyncAdapter) *CreditFilter {
	cfContract, err := creditFilter.NewCreditFilter(common.HexToAddress(adapter.Address), adapter.Client)
	if err != nil {
		log.Fatal(err)
	}
	obj := &CreditFilter{
		SyncAdapter: adapter,
		contractETH: cfContract,
	}
	return obj
}
