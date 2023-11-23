package cm_v3

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/common"
)

type CMv3 struct {
	expirationDate uint64
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
	obj.SetOnChangeFn()

	// set facade and contract address if not present
	obj.setv3AddrIfNotPresent()
	//
	///// checks if address is changed and cm wrapper will be nofitied
	// credit facade syncer
	obj.setCreditFacadeSyncer(obj.GetDetailsByKey("facade"))
	// set credit cofigurator syncer
	obj.setConfiguratorSyncer(obj.GetDetailsByKey("configurator"))
	/////
	//
	// get expiration date for liquidation call and setting expired state
	obj.expirationDate = func() uint64 {
		data, err := core.CallFuncWithExtraBytes(obj.Client, "8f620487",
			common.HexToAddress(obj.GetDetailsByKey("facade")), obj.LastSync, nil)
		log.CheckFatal(err)
		return uint64(new(big.Int).SetBytes(data).Int64())
	}()
	return obj
}

func (mdl *CMv3) GetAbi() {
	mdl.ABI = core.GetAbi("CreditFacadev3")
}
