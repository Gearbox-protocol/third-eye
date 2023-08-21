package cm_v3

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/third-eye/ds"
)

type CMv3 struct {
	CMv3Fields
	//
	Cmv3State
	allowedProtocols map[string]bool
}

func NewCMv3(addr string, client core.ClientI, repo ds.RepositoryI, discoveredAt int64) *CMv3 {
	// credit manager
	mdl := NewCMv3FromAdapter(
		ds.NewSyncAdapter(addr, ds.CreditManager, discoveredAt, client, repo),
	)
	mdl.InitState()
	mdl.addCreditConfiguratorAdapter(mdl.GetDetailsByKey("configurator"))
	return mdl
}

func NewCMv3FromAdapter(adapter *ds.SyncAdapter) *CMv3 {
	obj := &CMv3{
		Cmv3State:        NewCmv3State(adapter),
		allowedProtocols: map[string]bool{},
	}

	// obj.addProtocolAdaptersLocally()
	obj.GetAbi()
	// obj.SetOnChangeFn()

	obj.setv3AddrIfNotPresent()
	// credit facade syncer
	obj.setCreditFacadeSyncer(obj.GetDetailsByKey("facade"))
	// set credit cofigurator syncer
	obj.setConfiguratorSyncer(obj.GetDetailsByKey("configurator"))

	return obj
}

func (mdl *CMv3) GetAbi() {
	mdl.ABI = core.GetAbi("CreditFacadev3")
}
