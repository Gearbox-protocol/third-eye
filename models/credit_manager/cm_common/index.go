package cm_common

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/common"
)

type CMCommon struct {
	*ds.SyncAdapter
	//
	State *schemas.CreditManagerState
	// tmp storage
	pnlOnCM                *PnlCM
	borrowedAmountForBlock *big.Int
	params                 *schemas.Parameters
	// calculating credit session stats
	updatedSessions             map[string]int
	closedSessions              map[string]*SessionCloseDetails
	DontGetSessionFromDCForTest bool
	//
	onChangeDetails
}

func NewCMCommon(adapter *ds.SyncAdapter) CMCommon {
	return CMCommon{
		pnlOnCM:         NewPnlCM(),
		SyncAdapter:     adapter,
		updatedSessions: make(map[string]int),
		closedSessions:  make(map[string]*SessionCloseDetails),
	}
}

// get underlyigToken
// get pool
// set state
func (mdl *CMCommon) CommonInitState(version core.VersionType) {

	var underlyingToken common.Address
	if version.IsGBv1() {
		data, err := core.CallFuncWithExtraBytes(mdl.Client, "2495a599", common.HexToAddress(mdl.Address), mdl.DiscoveredAt, nil)
		log.CheckFatal(err) // [underlyingToken] on credit_manager v1
		underlyingToken = common.BytesToAddress(data)
	} else {
		data, err := core.CallFuncWithExtraBytes(mdl.Client, "6f307dc3", common.HexToAddress(mdl.Address), mdl.DiscoveredAt, nil) // [underlying] on credit manager v2
		log.CheckFatal(err)
		underlyingToken = common.BytesToAddress(data)
	}
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

// for params
func (mdl *CMCommon) SetParams(params *schemas.Parameters) {
	mdl.params = params
}

// for states
func (mdl *CMCommon) SetUnderlyingState(obj interface{}) {
	mdl.UnderlyingStateToSave = true
	switch underlyingObj := obj.(type) {
	case (*schemas.CreditManagerState):
		mdl.State = underlyingObj
	case *schemas.PnlOnRepay:
		mdl.pnlOnCM.Set(underlyingObj)
	case *schemas.Parameters:
		mdl.SetParams(underlyingObj)
	default:
		log.Fatal("Type assertion for credit manager state failed")
	}
}

func (mdl *CMCommon) GetUnderlyingState() interface{} {
	return mdl.State
}

func (mdl *CMCommon) GetUnderlyingToken() string {
	return mdl.State.UnderlyingToken
}

func (mdl *CMCommon) GetUnderlyingDecimal() int8 {
	decimals := mdl.Repo.GetToken(mdl.GetUnderlyingToken()).Decimals
	return decimals
}

func (mdl *CMCommon) AddAccountOperation(accountOperation *schemas.AccountOperation) {
	mdl.Repo.AddAccountOperation(accountOperation)
}
