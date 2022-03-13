package credit_manager

import (
	"fmt"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/utils"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

func (mdl *CreditManager) CMStatsOnOpenAccount(borrowAmount *big.Int) {
	// manager state
	mdl.State.TotalOpenedAccounts++
	mdl.State.OpenedAccountsCount++
	mdl.State.TotalBorrowedBI = core.AddCoreAndInt(mdl.State.TotalBorrowedBI, borrowAmount)
	mdl.State.TotalBorrowed = utils.GetFloat64Decimal(mdl.State.TotalBorrowedBI.Convert(), mdl.GetUnderlyingDecimal())
}
func (mdl *CreditManager) onOpenCreditAccountV2(txLog *types.Log, onBehalfOf, account string,
	borrowAmount,
	referralCode *big.Int) error {
	mdl.CMStatsOnOpenAccount(borrowAmount)
	// other operations
	cmAddr := txLog.Address.Hex()
	sessionId := fmt.Sprintf("%s_%d_%d", account, txLog.BlockNumber, txLog.Index)
	blockNum := int64(txLog.BlockNumber)
	// add account operation
	action, args := mdl.ParseEvent("OpenCreditAccount", txLog)
	accountOperation := &core.AccountOperation{
		TxHash:      txLog.TxHash.Hex(),
		BlockNumber: blockNum,
		LogId:       txLog.Index,
		Borrower:    onBehalfOf,
		SessionId:   sessionId,
		AdapterCall: false,
		Action:      action,
		Args:        args,
		Transfers: &core.Transfers{
			mdl.GetUnderlyingToken(): borrowAmount,
		},
		Dapp: cmAddr,
	}
	mdl.PoolBorrow(txLog, sessionId, onBehalfOf, borrowAmount)
	mdl.AddAccountOperation(accountOperation)
	mdl.UpdatedSessions[sessionId]++
	// add session to manager object
	mdl.AddCreditOwnerSession(onBehalfOf, sessionId)
	// create credit session
	newSession := &core.CreditSession{
		ID:             sessionId,
		Status:         core.Active,
		Borrower:       onBehalfOf,
		CreditManager:  mdl.Address,
		Account:        account,
		Since:          blockNum,
		BorrowedAmount: (*core.BigInt)(borrowAmount),
		IsDirty:        true,
		Version:        1,
	}
	mdl.Repo.AddCreditSession(newSession, false, txLog.TxHash.Hex(), txLog.Index)
	// mdl.AddCollateralToSession(blockNum, sessionId, mdl.State.UnderlyingToken, amount)
	return nil
}

func (mdl *CreditManager) onCloseCreditAccountV2(txLog *types.Log, owner, to string) error {
	mdl.State.TotalClosedAccounts++
	cmAddr := txLog.Address.Hex()
	sessionId := mdl.GetCreditOwnerSession(owner)
	// session := mdl.Repo.GetCreditSession(sessionId)
	// session.RemainingFunds = (*core.BigInt)(remainingFunds)
	blockNum := int64(txLog.BlockNumber)
	action, args := mdl.ParseEvent("CloseCreditAccount", txLog)
	// add account operation
	accountOperation := &core.AccountOperation{
		TxHash:      txLog.TxHash.Hex(),
		BlockNumber: blockNum,
		LogId:       txLog.Index,
		Borrower:    owner,
		SessionId:   sessionId,
		AdapterCall: false,
		Action:      action,
		Args:        args,
		// Transfers: &core.Transfers{
		// 	// mdl.GetUnderlyingToken(): remainingFunds,
		// },
		Dapp: cmAddr,
	}
	mdl.AddAccountOperation(accountOperation)
	mdl.ClosedSessions[sessionId] = &SessionCloseDetails{
		// RemainingFunds: remainingFunds,
		Status:   core.Closed,
		TxHash:   txLog.TxHash.Hex(),
		Borrower: owner,
	}
	// remove session to manager object
	mdl.RemoveCreditOwnerSession(owner)
	mdl.closeAccount(sessionId, blockNum, txLog.TxHash.Hex(), txLog.Index)
	return nil
}
