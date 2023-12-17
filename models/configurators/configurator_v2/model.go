package configurator_v2

import (
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditConfigurator"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/common"
)

type Configuratorv2 struct {
	*ds.SyncAdapter
	cfgContract *creditConfigurator.CreditConfigurator
}

func NewConfiguratorv2(addr, creditManager string, discoveredAt int64, client core.ClientI, repo ds.RepositoryI) *Configuratorv2 {
	syncAdapter := ds.NewSyncAdapter(addr, ds.CreditConfigurator, discoveredAt, client, repo)
	syncAdapter.Details = map[string]interface{}{"creditManager": creditManager}
	mdl := NewConfiguratorv2FromAdapter(
		syncAdapter,
	)
	return mdl
}

func (mdl *Configuratorv2) GetCM() string {
	return mdl.GetDetailsByKey("creditManager")
}

func NewConfiguratorv2FromAdapter(adapter *ds.SyncAdapter) *Configuratorv2 {
	cfgContract, err := creditConfigurator.NewCreditConfigurator(common.HexToAddress(adapter.Address), adapter.Client)
	log.CheckFatal(err)
	obj := &Configuratorv2{
		SyncAdapter: adapter,
		cfgContract: cfgContract,
	}
	return obj
}
