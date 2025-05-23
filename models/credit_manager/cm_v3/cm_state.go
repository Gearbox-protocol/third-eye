package cm_v3

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/creditManagerv3"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
	mp "github.com/Gearbox-protocol/third-eye/ds/multicall_processor"
	"github.com/Gearbox-protocol/third-eye/models/credit_manager/cm_common"
	"github.com/ethereum/go-ethereum/common"
)

type AccountOwner struct {
	Borrower  string
	SessionId string
}
type Cmv3State struct {
	*cm_common.CommonCMAdapter
	whosAccount      map[string]AccountOwner
	allowedProtocols map[string]bool
}

func NewCmv3State(adapter *ds.SyncAdapter) Cmv3State {
	return Cmv3State{
		CommonCMAdapter:  cm_common.NewCommonCMAdapter(adapter, &mp.MultiCallProcessorv3{}),
		allowedProtocols: map[string]bool{},
		whosAccount:      map[string]AccountOwner{},
	}
}

func (mdl *Cmv3State) SetUnderlyingState(obj interface{}) {
	mdl.UnderlyingStateToSave = true
	switch underlyingObj := obj.(type) {
	case (*schemas.CreditManagerState):
		mdl.State = underlyingObj
	case (map[string]AccountOwner):
		mdl.whosAccount = underlyingObj
	case *schemas.PnlOnRepay:
		mdl.PnlOnCM.Set(underlyingObj)
	case *schemas.Parameters:
		mdl.SetParams(underlyingObj)
	case *big.Int:
		mdl.setTotalDebtLimit(underlyingObj)
	default:
		log.Fatal("Type assertion for credit manager state failed")
	}
}

func (mdl *CMv3) InitState() {
	underlying := func() common.Address {
		data, err := core.CallFuncGetSingleValue(mdl.Client, "6f307dc3", common.HexToAddress(mdl.Address), mdl.DiscoveredAt, nil) // [underlying] on credit manager v2
		log.CheckFatal(err)
		return common.BytesToAddress(data)
	}()

	poolAddr := func() common.Address {
		data, err := core.CallFuncGetSingleValue(mdl.Client, "16f0115b", common.HexToAddress(mdl.Address), mdl.DiscoveredAt, nil)
		// [pool] on creditManager
		log.CheckFatal(err)
		return common.BytesToAddress(data)
	}()

	mdl.SetUnderlyingState(&schemas.CreditManagerState{
		Address:         mdl.Address,
		PoolAddress:     poolAddr.Hex(),
		UnderlyingToken: underlying.Hex(),
		Version:         core.NewVersion(300),
		Name: func() string {
			con, err := creditManagerv3.NewCreditManagerv3(common.HexToAddress(mdl.Address), mdl.Client)
			log.CheckFatal(err)
			name, err := con.Name(nil)
			log.CheckFatal(err)
			return name
		}(),
	})
}

func (mdl Cmv3State) AddCreditAccount(account, sessionId, owner string) {
	mdl.whosAccount[account] = AccountOwner{
		SessionId: sessionId,
		Borrower:  owner,
	}
}

func (mdl Cmv3State) RemoveCreditAccount(account string) {
	delete(mdl.whosAccount, account)
}

func (mdl Cmv3State) GetSessionIdAndBorrower(account string, dontFail ...bool) (string, string) {
	details, ok := mdl.whosAccount[account]
	if (len(dontFail) == 0 || !dontFail[0]) && !ok {
		log.Info(mdl.whosAccount)
		log.Fatalf("session id not found for %s cm(%s)\n", account, mdl.Address)
	}
	return details.SessionId, details.Borrower
}

func (mdl *Cmv3State) setTotalDebtLimit(total *big.Int) {
	mdl.State.TotalDebtLimit = (*core.BigInt)(total)
}
