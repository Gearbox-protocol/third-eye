package cm_common

import (
	"fmt"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/core/types"
)

type SessionCloseDetails struct {
	RemainingFunds   *big.Int
	Status           int
	LogId            uint
	TxHash           string
	Borrower         string
	AccountOperation *schemas.AccountOperation
}

func (x SessionCloseDetails) String() string {
	return fmt.Sprintf("ClosingDetails(Status: %d LogId %d TxHash %s Borrower %s RemainingFunds %s)",
		x.Status, x.LogId, x.TxHash, x.Borrower, x.RemainingFunds)
}

func (mdl *CMCommon) PoolBorrow(txLog *types.Log, sessionId, borrower string, amount *big.Int) {
	mdl.Repo.AddPoolLedger(&schemas.PoolLedger{
		LogId:       txLog.Index,
		BlockNumber: int64(txLog.BlockNumber),
		TxHash:      txLog.TxHash.Hex(),
		Pool:        mdl.State.PoolAddress,
		Event:       "Borrow",
		User:        borrower,
		SessionId:   sessionId,
		AmountBI:    (*core.BigInt)(amount),
		Amount:      utils.GetFloat64Decimal(amount, mdl.GetUnderlyingDecimal()),
	})
}

func (mdl *CMCommon) PoolRepay(blockNum int64, logId uint, txHash, sessionId, borrower string, amount *big.Int) {
	mdl.Repo.AddPoolLedger(&schemas.PoolLedger{
		LogId:       logId,
		BlockNumber: blockNum,
		TxHash:      txHash,
		Pool:        mdl.State.PoolAddress,
		Event:       "Repay",
		User:        borrower,
		SessionId:   sessionId,
		AmountBI:    (*core.BigInt)(amount),
		Amount:      utils.GetFloat64Decimal(amount, mdl.GetUnderlyingDecimal()),
	})
}

func (mdl CMCommon) SetSessionIsUpdated(sessionId string) {
	// log.Info(log.DetectFunc(),sessionId, "increased")
	mdl.updatedSessions[sessionId]++
}

func (mdl CMCommon) SetSessionIsClosed(sessionId string, details *SessionCloseDetails) {
	mdl.closedSessions[sessionId] = details
}
func (mdl CMCommon) UpdateClosedSessionStatus(sessionId string, status int) {
	mdl.closedSessions[sessionId].Status = status
}

func (mdl *CMCommon) CloseAccount(sessionID string, blockNum int64, txHash string, logID uint) {
	session := mdl.Repo.GetCreditSession(sessionID)
	mdl.Repo.GetAccountManager().CloseAccountDetails(session.Account, session.Since, blockNum, txHash, logID)
}

func (mdl *CMCommon) SaveExecuteEvents(lastTxHash string, executeParams []ds.ExecuteParams) {
	// credit manager has the execute event
	calls := mdl.Repo.GetExecuteParser().GetExecuteCalls(lastTxHash, mdl.Address, executeParams)

	for i, call := range calls {
		params := executeParams[i]

		accountOperation := &schemas.AccountOperation{
			BlockNumber: params.BlockNumber,
			TxHash:      lastTxHash,
			LogId:       params.Index,
			// owner/account data
			Borrower:  params.Borrower.Hex(),
			SessionId: params.SessionId,
			// dapp
			Dapp: params.Protocol.Hex(),
			// call/events data
			Action:      call.Name,
			Args:        call.Args,
			AdapterCall: true,
			Transfers:   &call.Transfers,
		}
		mdl.AddAccountOperation(accountOperation)
		mdl.SetSessionIsUpdated(params.SessionId)
	}
}

func (mdl CMCommon) AddCollateralToSession(blockNum int64, sessionId, token string, amount *big.Int) {
	if !mdl.Repo.IsDieselToken(token) && mdl.Repo.GetGearTokenAddr() != token {
		session := mdl.Repo.GetCreditSession(sessionId)
		//
		if session.Collateral == nil {
			session.Collateral = &core.JsonBigIntMap{}
		}
		(*session.Collateral)[token] = (*core.BigInt)(new(big.Int).Add(
			core.NewBigInt((*session.Collateral)[token]).Convert(),
			amount,
		))
		//
		valueInUSD := mdl.Repo.GetValueInCurrency(blockNum, session.Version, token, "USDC", amount)
		session.CollateralInUSD = session.CollateralInUSD + utils.GetFloat64Decimal(valueInUSD, 6)
		valueInUnderlyingAsset := mdl.Repo.GetValueInCurrency(blockNum, session.Version, token, mdl.GetUnderlyingToken(), amount)
		session.CollateralInUnderlying += utils.GetFloat64Decimal(valueInUnderlyingAsset, mdl.GetUnderlyingDecimal())
	}
}
