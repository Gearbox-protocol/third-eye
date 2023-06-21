package cm_common

import (
	"fmt"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/creditManager"
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditManagerv2"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

// sets underlying state on init
// pool, and underlying token address
func (mdl *CMCommon) CommonInitState(version core.VersionType) {
	// do state changes
	// create underlying token
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(mdl.DiscoveredAt),
	}
	var underlyingToken common.Address
	var err error
	cmContract, err := creditManager.NewCreditManager(common.HexToAddress(mdl.Address), mdl.Client)
	log.CheckFatal(err)

	switch version {
	case 1:
		underlyingToken, err = cmContract.UnderlyingToken(opts)
		log.CheckFatal(err)
	case 2:
		contract, err := creditManagerv2.NewCreditManagerv2(common.HexToAddress(mdl.Address), mdl.Client)
		log.CheckFatal(err)
		underlyingToken, err = contract.Underlying(opts)
		log.CheckFatal(err)
	}
	mdl.Repo.GetToken(underlyingToken.Hex())
	//
	poolAddr, err := cmContract.PoolService(opts)
	if err != nil {
		log.Fatal(err)
	}
	mdl.SetUnderlyingState(&schemas.CreditManagerState{
		Address:         mdl.Address,
		PoolAddress:     poolAddr.Hex(),
		UnderlyingToken: underlyingToken.Hex(),
		Sessions:        map[string]string{},
		Version:         version,
	})
}

// for params
func (mdl *CMCommon) SetParams(params *schemas.Parameters) {
	mdl.params = params
}

// for states
func (mdl *CMCommon) SetUnderlyingState(obj interface{}) {
	mdl.UnderlyingStatePresent = true
	switch underlyingObj := obj.(type) {
	case (*schemas.CreditManagerState):
		mdl.State = underlyingObj
	case (map[string]string):
		mdl.Sessions = underlyingObj
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

func (mdl *CMCommon) AddCreditOwnerSession(owner, sessionId string) {
	mdl.State.Sessions[owner] = sessionId
}

func (mdl *CMCommon) RemoveCreditOwnerSession(owner string) {
	delete(mdl.State.Sessions, owner)
}

func (mdl *CMCommon) GetCreditOwnerSession(owner string, dontFail ...bool) string {
	sessionId := mdl.State.Sessions[owner]
	if (len(dontFail) == 0 || !dontFail[0]) && sessionId == "" {
		panic(
			fmt.Sprintf("session id not found for %s in %+v %s\n", owner, mdl.State.Sessions, mdl.Address),
		)
	}
	return sessionId
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
