package credit_manager

import (
	"fmt"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strings"
)

func (mdl *CreditManager) onOpenCreditAccount(txLog *types.Log, sender, onBehalfOf, account string,
	amount,
	borrowAmount,
	referralCode *big.Int) error {
	// manager state
	mdl.State.TotalOpenedAccounts++
	mdl.State.OpenedAccountsCount++
	mdl.State.TotalBorrowedBI = core.AddCoreAndInt(mdl.State.TotalBorrowedBI, borrowAmount)
	mdl.State.TotalBorrowed = utils.GetFloat64Decimal(mdl.State.TotalBorrowedBI.Convert(), mdl.GetUnderlyingDecimal())
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
			mdl.GetUnderlyingToken(): new(big.Int).Add(borrowAmount, amount),
		},
		Dapp: cmAddr,
	}
	mdl.PoolBorrow(txLog, sessionId, onBehalfOf, borrowAmount)
	mdl.AddEventBasedAccountOperationAndState(accountOperation,
		borrowAmount,
		false,
		cmAddr)
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
		InitialAmount:  (*core.BigInt)(amount),
		BorrowedAmount: (*core.BigInt)(borrowAmount),
		IsDirty:        true,
	}
	mdl.Repo.AddCreditSession(newSession, false, txLog.TxHash.Hex(), txLog.Index)
	mdl.AddCollateralToSession(blockNum, sessionId, mdl.State.UnderlyingToken, amount)
	return nil
}

// onCloseCreditAccount handles CloseCreditAccount Event
func (mdl *CreditManager) onCloseCreditAccount(txLog *types.Log, owner, to string, remainingFunds *big.Int) error {
	mdl.State.TotalClosedAccounts++
	cmAddr := txLog.Address.Hex()
	sessionId := mdl.GetCreditOwnerSession(owner)
	session := mdl.Repo.GetCreditSession(sessionId)
	session.RemainingFunds = (*core.BigInt)(remainingFunds)
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
		Transfers: &core.Transfers{
			mdl.GetUnderlyingToken(): remainingFunds,
		},
		Dapp: cmAddr,
	}
	mdl.AddEventBasedAccountOperationAndState(accountOperation,
		nil,
		true,
		cmAddr)
	mdl.ClosedSessions[sessionId] = &SessionCloseDetails{RemainingFunds: remainingFunds, Status: core.Closed}
	// remove session to manager object
	mdl.RemoveCreditOwnerSession(owner)
	mdl.closeAccount(sessionId, blockNum, txLog.TxHash.Hex(), txLog.Index)
	return nil
}

func (mdl *CreditManager) onLiquidateCreditAccount(txLog *types.Log, owner, liquidator string, remainingFunds *big.Int) error {
	mdl.State.TotalLiquidatedAccounts++
	sessionId := mdl.GetCreditOwnerSession(owner)
	blockNum := int64(txLog.BlockNumber)
	action, args := mdl.ParseEvent("LiquidateCreditAccount", txLog)
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
		Transfers: &core.Transfers{
			mdl.GetUnderlyingToken(): remainingFunds,
		},
		Dapp: txLog.Address.Hex(),
	}
	mdl.AddEventBasedAccountOperationAndState(accountOperation,
		nil,
		true,
		mdl.GetAddress())
	mdl.ClosedSessions[sessionId] = &SessionCloseDetails{RemainingFunds: remainingFunds, Status: core.Liquidated}
	session := mdl.Repo.GetCreditSession(sessionId)
	session.Liquidator = liquidator
	session.RemainingFunds = (*core.BigInt)(remainingFunds)
	// remove session to manager object
	mdl.RemoveCreditOwnerSession(owner)
	mdl.closeAccount(sessionId, blockNum, txLog.TxHash.Hex(), txLog.Index)
	return nil
}

func (mdl *CreditManager) onRepayCreditAccount(txLog *types.Log, owner, to string) error {
	mdl.State.TotalRepaidAccounts++
	cmAddr := txLog.Address.Hex()
	sessionId := mdl.GetCreditOwnerSession(owner)
	blockNum := int64(txLog.BlockNumber)
	action, args := mdl.ParseEvent("RepayCreditAccount", txLog)
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
		Transfers:   nil,
		Dapp:        cmAddr,
	}
	// Since remainingFunds is not known for repay, we get it from datacompressor at end of each block
	mdl.ClosedSessions[sessionId] = &SessionCloseDetails{RemainingFunds: nil,
		Status:           core.Repaid,
		LogId:            txLog.Index,
		AccountOperation: accountOperation,
	}
	// remove session to manager object
	mdl.RemoveCreditOwnerSession(owner)
	mdl.closeAccount(sessionId, blockNum, txLog.TxHash.Hex(), txLog.Index)
	return nil
}

func (mdl *CreditManager) closeAccount(sessionID string, blockNum int64, txHash string, logID uint) {
	session := mdl.Repo.GetCreditSession(sessionID)
	mdl.Repo.GetAccountManager().CloseAccountDetails(session.Account, session.Since, blockNum, txHash, logID)
}

func (mdl *CreditManager) onAddCollateral(txLog *types.Log, onBehalfOf, token string, value *big.Int) error {
	sessionId := mdl.GetCreditOwnerSession(onBehalfOf)
	blockNum := int64(txLog.BlockNumber)
	action, args := mdl.ParseEvent("AddCollateral", txLog)
	// add account operation
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
			token: value,
		},
		Dapp: txLog.Address.Hex(),
	}
	mdl.AddEventBasedAccountOperationAndState(accountOperation,
		nil,
		false,
		mdl.GetAddress())
	mdl.AddCollateralToSession(blockNum, sessionId, token, value)
	mdl.UpdatedSessions[sessionId]++
	return nil
}

func (mdl *CreditManager) AddCollateralToSession(blockNum int64, sessionId, token string, amount *big.Int) {
	if !mdl.Repo.IsDieselToken(token) && mdl.Repo.GetGearTokenAddr() != token {
		session := mdl.Repo.GetCreditSession(sessionId)
		valueInUSD := mdl.Repo.GetValueInCurrency(blockNum, token, "USDC", amount)
		session.CollateralInUSD = core.AddCoreAndInt(session.CollateralInUSD, valueInUSD)
		valueInUnderlyingAsset := mdl.Repo.GetValueInCurrency(blockNum, token, mdl.GetUnderlyingToken(), amount)
		session.CollateralInUnderlying = core.AddCoreAndInt(session.CollateralInUnderlying, valueInUnderlyingAsset)
	}
}

func (mdl *CreditManager) onIncreaseBorrowedAmount(txLog *types.Log, borrower string, amount *big.Int) error {
	// manager state
	mdl.State.TotalBorrowedBI = core.AddCoreAndInt(mdl.State.TotalBorrowedBI, amount)
	mdl.State.TotalBorrowed = utils.GetFloat64Decimal(mdl.State.TotalBorrowedBI.Convert(), mdl.GetUnderlyingDecimal())
	// other operations
	sessionId := mdl.GetCreditOwnerSession(borrower)
	blockNum := int64(txLog.BlockNumber)
	action, args := mdl.ParseEvent("IncreaseBorrowedAmount", txLog)
	// add account operation
	accountOperation := &core.AccountOperation{
		TxHash:      txLog.TxHash.Hex(),
		BlockNumber: blockNum,
		LogId:       txLog.Index,
		Borrower:    borrower,
		SessionId:   sessionId,
		AdapterCall: false,
		Action:      action,
		Args:        args,
		Transfers: &core.Transfers{
			mdl.GetUnderlyingToken(): amount,
		},
		Dapp: txLog.Address.Hex(),
	}
	mdl.PoolBorrow(txLog, sessionId, borrower, amount)
	mdl.AddEventBasedAccountOperationAndState(accountOperation,
		amount,
		false,
		mdl.GetAddress())
	mdl.UpdatedSessions[sessionId]++
	return nil
}

func (mdl *CreditManager) onTransferAccount(txLog *types.Log, owner, newOwner string) error {
	sessionId := mdl.GetCreditOwnerSession(owner)
	action, args := mdl.ParseEvent("TransferAccount", txLog)
	// add account operation
	accountOperation := &core.AccountOperation{
		TxHash:      txLog.TxHash.Hex(),
		BlockNumber: int64(txLog.BlockNumber),
		LogId:       txLog.Index,
		Borrower:    owner,
		SessionId:   sessionId,
		AdapterCall: false,
		Action:      action,
		Args:        args,
		Transfers:   nil,
		Dapp:        txLog.Address.Hex(),
	}
	mdl.AddAccountOperation(accountOperation)
	// remove session to manager object
	mdl.RemoveCreditOwnerSession(owner)
	mdl.AddCreditOwnerSession(newOwner, sessionId)
	session := mdl.Repo.GetCreditSession(sessionId)
	session.Borrower = newOwner
	session.IsDirty = true
	return nil
}

func (mdl *CreditManager) AddExecuteParams(txLog *types.Log,
	borrower,
	targetContract common.Address) error {
	sessionId := mdl.GetCreditOwnerSession(borrower.Hex())
	blockNum := int64(txLog.BlockNumber)
	mdl.executeParams = append(mdl.executeParams, core.ExecuteParams{
		SessionId:     sessionId,
		CreditAccount: common.HexToAddress(strings.Split(sessionId, "_")[0]),
		Protocol:      targetContract,
		Borrower:      borrower,
		Index:         txLog.Index,
		BlockNumber:   blockNum,
	})
	return nil
}

func (mdl *CreditManager) handleExecuteEvents() {

	calls := mdl.Repo.GetExecuteParser().GetExecuteCalls(mdl.LastTxHash, mdl.Address, mdl.executeParams)

	for i, call := range calls {
		params := mdl.executeParams[i]
		// add account operation
		accountOperation := &core.AccountOperation{
			BlockNumber: params.BlockNumber,
			TxHash:      mdl.LastTxHash,
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
			// extras
			Depth: call.Depth,
		}
		mdl.AddEventBasedAccountOperationAndState(accountOperation,
			nil,
			false,
			mdl.GetAddress())
		mdl.UpdatedSessions[params.SessionId]++
	}
}

func (mdl *CreditManager) AddEventBasedAccountOperationAndState(
	accountOperation *core.AccountOperation,
	borrowAmount *big.Int,
	clear bool,
	cmAddr string) {
	mdl.Repo.AddEventBalance(core.NewEventBalance(
		accountOperation.BlockNumber,
		accountOperation.LogId,
		accountOperation.SessionId,
		borrowAmount,
		*accountOperation.Transfers,
		false,
		cmAddr,
	))
	mdl.AddAccountOperation(accountOperation)
}

func (mdl *CreditManager) AddAccountOperation(accountOperation *core.AccountOperation) {
	mdl.Repo.AddAccountOperation(accountOperation)
}

func (mdl *CreditManager) PoolBorrow(txLog *types.Log, sessionId, borrower string, amount *big.Int) {
	mdl.Repo.AddPoolLedger(&core.PoolLedger{
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

func (mdl *CreditManager) PoolRepay(blockNum int64, logId uint, txHash, sessionId, borrower string, amount *big.Int) {
	mdl.Repo.AddPoolLedger(&core.PoolLedger{
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
