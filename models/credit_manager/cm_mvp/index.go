package cm_mvp

import (
	"fmt"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
	mp "github.com/Gearbox-protocol/third-eye/ds/multicall_processor"
	"github.com/Gearbox-protocol/third-eye/models/credit_manager/cm_common"
	"github.com/ethereum/go-ethereum/common"
)

// It is only used by v1 and v2 cm
type Cmv1v2 struct {
	*cm_common.CommonCMAdapter
	Sessions map[string]string // borrower to sessionId
	//
}

func NewCMv1v2(adapter *ds.SyncAdapter) *Cmv1v2 {
	return &Cmv1v2{
		CommonCMAdapter: cm_common.NewCommonCMAdapter(adapter, &mp.MultiCallProcessorv2{}),
		Sessions:        map[string]string{},
	}
}

// get underlyigToken
// get pool
// set state
func (mdl *Cmv1v2) CommonInitState(version core.VersionType) {

	underlyingToken := func() common.Address {
		if version.IsGBv1() {
			data, err := core.CallFuncWithExtraBytes(mdl.Client, "2495a599", common.HexToAddress(mdl.Address), mdl.DiscoveredAt, nil)
			log.CheckFatal(err) // [underlyingToken] on credit_manager v1
			return common.BytesToAddress(data)
		} else {
			data, err := core.CallFuncWithExtraBytes(mdl.Client, "6f307dc3", common.HexToAddress(mdl.Address), mdl.DiscoveredAt, nil) // [underlying] on credit manager v2
			log.CheckFatal(err)
			return common.BytesToAddress(data)
		}
	}()
	mdl.Repo.GetToken(underlyingToken.Hex())
	//

	data, err := core.CallFuncWithExtraBytes(mdl.Client, "570a7af2", common.HexToAddress(mdl.Address), mdl.DiscoveredAt, nil)
	// [PoolService] on creditManager
	log.CheckFatal(err)
	poolAddr := common.BytesToAddress(data)
	mdl.SetUnderlyingState(&schemas.CreditManagerState{
		Address:         mdl.Address,
		PoolAddress:     poolAddr.Hex(),
		UnderlyingToken: underlyingToken.Hex(),
		Version:         version,
	})
}

// for states
func (mdl *Cmv1v2) SetUnderlyingState(obj interface{}) {
	mdl.UnderlyingStateToSave = true
	switch underlyingObj := obj.(type) {
	case (*schemas.CreditManagerState):
		mdl.State = underlyingObj
	case (map[string]string):
		mdl.Sessions = underlyingObj
	case *schemas.PnlOnRepay:
		mdl.PnlOnCM.Set(underlyingObj)
	case *schemas.Parameters:
		mdl.SetParams(underlyingObj)
	default:
		log.Fatal("Type assertion for credit manager state failed")
	}
}

func (mdl *Cmv1v2) AddCreditOwnerSession(owner, sessionId string) {
	mdl.Sessions[owner] = sessionId
}

func (mdl *Cmv1v2) RemoveCreditOwnerSession(owner string) {
	delete(mdl.Sessions, owner)
}

func (mdl *Cmv1v2) GetCreditOwnerSession(owner string, dontFail ...bool) string {
	sessionId := mdl.Sessions[owner]
	if (len(dontFail) == 0 || !dontFail[0]) && sessionId == "" {
		panic(
			fmt.Sprintf("session id not found for %s in %+v %s\n", owner, mdl.Sessions, mdl.Address),
		)
	}
	return sessionId
}
