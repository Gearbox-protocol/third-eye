package cm_v2

import (
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditFacade"
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditManagerv2"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/credit_manager/cm_common"
	"github.com/ethereum/go-ethereum/common"
)

type CMv2 struct {
	cm_common.CMCommon
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
	mdl.addCreditConfiguratorAdapter(mdl.GetDetailsByKey("configurator"))
	return mdl
}

func NewCMv2FromAdapter(adapter *ds.SyncAdapter) *CMv2 {
	//
	obj := &CMv2{
		CMCommon:         cm_common.NewCMCommon(adapter),
		CMv2Fields:       CMv2Fields{},
		allowedProtocols: map[string]bool{},
	}

	// cm is registered with dataCompressor after discoveredAt, so we can get adapters for blockNum more than discoveredAt
	blockToFetchCMData := obj.DiscoveredAt
	if blockToFetchCMData < obj.LastSync {
		blockToFetchCMData = obj.LastSync
	}
	obj.addProtocolAdaptersLocally(blockToFetchCMData)
	obj.GetAbi()
	obj.SetOnChange()

	// v2 logic
	//
	// set credit manager
	cmContract, err := creditManagerv2.NewCreditManagerv2(common.HexToAddress(adapter.Address), adapter.Client)
	if err != nil {
		log.Fatal(err)
	}
	obj.cmContractv2 = cmContract
	// set credit facade contract
	obj.facadeContractv2, err = creditFacade.NewCreditFacade(core.NULL_ADDR, nil)
	log.CheckFatal(err)

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
