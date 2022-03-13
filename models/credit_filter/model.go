package credit_filter

import (
	"github.com/Gearbox-protocol/third-eye/artifacts/creditConfigurator"
	"github.com/Gearbox-protocol/third-eye/artifacts/creditFilter"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/ethereum/go-ethereum/common"
)

type CreditFilter struct {
	*core.SyncAdapter
	filterContract *creditFilter.CreditFilter
	cfgContract    *creditConfigurator.CreditConfigurator
}

func NewCreditFilter(addr, contractName, creditManager string, discoveredAt int64, client ethclient.ClientI, repo core.RepositoryI) *CreditFilter {
	syncAdapter := core.NewSyncAdapter(addr, contractName, discoveredAt, client, repo)
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
		SyncAdapter:    adapter,
		filterContract: cfContract,
	}
	if adapter.ContractName == core.CreditConfigurator {
		cfgContract, err := creditConfigurator.NewCreditConfigurator(common.HexToAddress(adapter.Address), adapter.Client)
		log.CheckFatal(err)
		obj.cfgContract = cfgContract
	}
	return obj
}
