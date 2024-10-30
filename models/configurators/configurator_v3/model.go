package configurator_v3

import (
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditConfiguratorv3"
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditConfiguratorv310"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/common"
)

type Configuratorv3 struct {
	*ds.SyncAdapter
	cfgContract     *creditConfiguratorv3.CreditConfiguratorv3
	cfgContractv310 *creditConfiguratorv310.CreditConfiguratorv310
}

func NewConfiguratorv3(addr, creditManager string, discoveredAt int64, client core.ClientI, repo ds.RepositoryI) *Configuratorv3 {
	syncAdapter := ds.NewSyncAdapter(addr, ds.CreditConfigurator, discoveredAt, client, repo)
	syncAdapter.Details = map[string]interface{}{"creditManager": creditManager}
	mdl := NewConfiguratorv3FromAdapter(
		syncAdapter,
	)
	return mdl
}

func (mdl *Configuratorv3) GetCM() string {
	return mdl.GetDetailsByKey("creditManager")
}

func NewConfiguratorv3FromAdapter(adapter *ds.SyncAdapter) *Configuratorv3 {
	cfgContract, err := creditConfiguratorv3.NewCreditConfiguratorv3(common.HexToAddress(adapter.Address), adapter.Client)
	log.CheckFatal(err)
	obj := &Configuratorv3{
		SyncAdapter: adapter,
		cfgContract: cfgContract,
	}
	return obj
}
