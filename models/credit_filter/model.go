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
	*core.State
	contractETH *creditFilter.CreditFilter
}

func NewCreditFilter(addr, creditManager string, discoveredAt int64, client *ethclient.Client, repo core.RepositoryI) *CreditFilter {
	syncAdapter := core.NewSyncAdapter(addr, "CreditFilter", discoveredAt, client)
	syncAdapter.Details = map[string]string{"creditManager": creditManager}
	return NewCreditFilterFromAdapter(
		repo,
		syncAdapter,
	)
}

func NewCreditFilterFromAdapter(repo core.RepositoryI, adapter *core.SyncAdapter) *CreditFilter {
	cfContract, err := creditFilter.NewCreditFilter(common.HexToAddress(adapter.Address), adapter.Client)
	if err != nil {
		log.Fatal(err)
	}
	obj := &CreditFilter{
		SyncAdapter: adapter,
		State:       &core.State{Repo: repo},
		contractETH: cfContract,
	}
	return obj
}
