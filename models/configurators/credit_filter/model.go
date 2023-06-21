package credit_filter

import (
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditFilter"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/common"
)

type CreditFilter struct {
	*ds.SyncAdapter
	filterContract *creditFilter.CreditFilter
}

func NewCreditFilter(addr, creditManager string, discoveredAt int64, client core.ClientI, repo ds.RepositoryI) *CreditFilter {
	syncAdapter := ds.NewSyncAdapter(addr, ds.CreditFilter, discoveredAt, client, repo)
	syncAdapter.Details = map[string]interface{}{"creditManager": creditManager}
	mdl := NewCreditFilterFromAdapter(
		syncAdapter,
	)
	return mdl
}

func NewCreditFilterFromAdapter(adapter *ds.SyncAdapter) *CreditFilter {
	cfContract, err := creditFilter.NewCreditFilter(common.HexToAddress(adapter.Address), adapter.Client)
	log.CheckFatal(err)
	obj := &CreditFilter{
		SyncAdapter:    adapter,
		filterContract: cfContract,
	}
	return obj
}

func (mdl *CreditFilter) GetCM() string {
	return mdl.GetDetailsByKey("creditManager")
}
