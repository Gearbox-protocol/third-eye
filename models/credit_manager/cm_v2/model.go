package cm_v2

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/credit_manager/cm_mvp"
)

type CMv2 struct {
	cm_mvp.CmMVP
	CMv2Fields
	//
	allowedProtocols map[string]bool
}

func NewCMv2(addr string, client core.ClientI, repo ds.RepositoryI, discoveredAt int64) *CMv2 {
	// credit manager
	mdl := NewCMv2FromAdapter(
		ds.NewSyncAdapter(addr, ds.CreditManager, discoveredAt, client, repo),
	)
	mdl.CommonInitState(mdl.GetVersion())
	// add contract to the syncadapter on creation
	mdl.addCreditConfiguratorAdapter(mdl.GetDetailsByKey("configurator"))
	return mdl
}

func NewCMv2FromAdapter(adapter *ds.SyncAdapter) *CMv2 {
	//
	obj := &CMv2{
		CmMVP:            cm_mvp.NewCMCommon(adapter),
		CMv2Fields:       CMv2Fields{},
		allowedProtocols: map[string]bool{},
	}

	obj.addProtocolAdaptersLocally()
	obj.GetAbi()
	obj.SetOnChangeFn()

	// v2 logic
	// set facade and configurator in map
	obj.setv2AddrIfNotPresent()
	// credit facade syncer
	obj.setCreditFacadeSyncer(obj.GetDetailsByKey("facade"))
	// set credit cofigurator syncer
	obj.setConfiguratorSyncer(obj.GetDetailsByKey("configurator"))
	//
	// v2 logic
	return obj
}

func (mdl *CMv2) GetAbi() {
	mdl.ABI = core.GetAbi("CreditFacade")
}
