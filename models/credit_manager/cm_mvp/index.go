package cm_mvp

import (
	"fmt"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/credit_manager/cm_common"
	"github.com/ethereum/go-ethereum/common"
)

type CmMVP struct {
	cm_common.CommonCMAdapter
	Sessions map[string]string // borrower to sessionId
	//
}

func NewCMCommon(adapter *ds.SyncAdapter) CmMVP {
	return CmMVP{
		CommonCMAdapter: cm_common.NewCommonCMAdapter(adapter),
		Sessions:        map[string]string{},
	}
}

// get underlyigToken
// get pool
// set state
func (mdl *CmMVP) CommonInitState(version core.VersionType) {

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
func (mdl *CmMVP) SetUnderlyingState(obj interface{}) {
	mdl.UnderlyingStatePresent = true
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

func (mdl *CmMVP) AddCreditOwnerSession(owner, sessionId string) {
	mdl.Sessions[owner] = sessionId
}

func (mdl *CmMVP) RemoveCreditOwnerSession(owner string) {
	delete(mdl.Sessions, owner)
}

func (mdl *CmMVP) GetCreditOwnerSession(owner string, dontFail ...bool) string {
	sessionId := mdl.Sessions[owner]
	if (len(dontFail) == 0 || !dontFail[0]) && sessionId == "" {
		panic(
			fmt.Sprintf("session id not found for %s in %+v %s\n", owner, mdl.Sessions, mdl.Address),
		)
	}
	return sessionId
}
