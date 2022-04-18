package credit_filter

import (
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditConfigurator"
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditFilter"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/common"
)

type CreditFilter struct {
	*ds.SyncAdapter
	filterContract *creditFilter.CreditFilter
	cfgContract    *creditConfigurator.CreditConfigurator
}

func NewCreditFilter(addr, contractName, creditManager string, discoveredAt int64, client core.ClientI, repo ds.RepositoryI) *CreditFilter {
	syncAdapter := ds.NewSyncAdapter(addr, contractName, discoveredAt, client, repo)
	syncAdapter.Details = map[string]interface{}{"creditManager": creditManager}
	return NewCreditFilterFromAdapter(
		syncAdapter,
	)
}

func NewCreditFilterFromAdapter(adapter *ds.SyncAdapter) *CreditFilter {
	cfContract, err := creditFilter.NewCreditFilter(common.HexToAddress(adapter.Address), adapter.Client)
	if err != nil {
		log.Fatal(err)
	}
	obj := &CreditFilter{
		SyncAdapter:    adapter,
		filterContract: cfContract,
	}
	if adapter.ContractName == ds.CreditConfigurator {
		cfgContract, err := creditConfigurator.NewCreditConfigurator(common.HexToAddress(adapter.Address), adapter.Client)
		log.CheckFatal(err)
		obj.cfgContract = cfgContract
	}
	return obj
}
