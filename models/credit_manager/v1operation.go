package credit_manager

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (mdl *CreditManager) onOpenCreditAccount(txLog *types.Log, onBehalfOf, account string,
	collateral, // collateral/user added funds
	borrowAmount,
	referralCode *big.Int) error {
	// manager state
	mdl.CMStatsOnOpenAccount(borrowAmount)
	// other operations
	cmAddr := txLog.Address.Hex()
	sessionId := fmt.Sprintf("%s_%d_%d", account, txLog.BlockNumber, txLog.Index)
	blockNum := int64(txLog.BlockNumber)
	// add account operation
	action, args := mdl.ParseEvent("OpenCreditAccount", txLog)
	accountOperation := &schemas.AccountOperation{
		TxHash:      txLog.TxHash.Hex(),
		BlockNumber: blockNum,
		LogId:       txLog.Index,
		Borrower:    onBehalfOf,
		SessionId:   sessionId,
		AdapterCall: false,
		Action:      action,
		Args:        args,
		Transfers: &core.Transfers{
			mdl.GetUnderlyingToken(): new(big.Int).Add(borrowAmount, collateral), // total amount
		},
		Dapp: cmAddr,
	}
	mdl.PoolBorrow(txLog, sessionId, onBehalfOf, borrowAmount)
	mdl.AddAccountOperation(accountOperation)
	mdl.setUpdateSession(sessionId)

	// add session to manager object
	mdl.AddCreditOwnerSession(onBehalfOf, sessionId)
	// create credit session
	newSession := &schemas.CreditSession{
		ID:             sessionId,
		Status:         schemas.Active,
		Borrower:       onBehalfOf,
		CreditManager:  mdl.Address,
		Account:        account,
		Since:          blockNum,
		InitialAmount:  (*core.BigInt)(collateral),
		BorrowedAmount: (*core.BigInt)(borrowAmount),
		IsDirty:        true,
		Version:        1,
	}
	mdl.Repo.AddCreditSession(newSession, false, txLog.TxHash.Hex(), txLog.Index)
	mdl.setUpdateSession(sessionId)
	mdl.AddCollateralToSession(blockNum, sessionId, mdl.State.UnderlyingToken, collateral)
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
	accountOperation := &schemas.AccountOperation{
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
	mdl.AddAccountOperation(accountOperation)
	mdl.ClosedSessions[sessionId] = &SessionCloseDetails{
		LogId:          txLog.Index,
		RemainingFunds: remainingFunds,
		Status:         schemas.Closed,
		TxHash:         txLog.TxHash.Hex(),
		Borrower:       owner,
	}
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
	accountOperation := &schemas.AccountOperation{
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
	mdl.AddAccountOperation(accountOperation)
	mdl.ClosedSessions[sessionId] = &SessionCloseDetails{
		LogId:          txLog.Index,
		RemainingFunds: remainingFunds,
		Status:         schemas.Liquidated,
		TxHash:         txLog.TxHash.Hex(),
		Borrower:       owner,
	}
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
	accountOperation := &schemas.AccountOperation{
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
		Status:           schemas.Repaid,
		LogId:            txLog.Index,
		AccountOperation: accountOperation,
		TxHash:           txLog.TxHash.Hex(),
		Borrower:         owner,
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
	accountOperation := &schemas.AccountOperation{
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
	mdl.AddAccountOperation(accountOperation)
	mdl.AddCollateralToSession(blockNum, sessionId, token, value)
	mdl.setUpdateSession(sessionId)
	return nil
}

func (mdl *CreditManager) AddCollateralToSession(blockNum int64, sessionId, token string, amount *big.Int) {
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
		log.Info(blockNum, session.CollateralInUSD, utils.GetFloat64Decimal(valueInUSD, 6))
		session.CollateralInUSD = session.CollateralInUSD + utils.GetFloat64Decimal(valueInUSD, 6)
		valueInUnderlyingAsset := mdl.Repo.GetValueInCurrency(blockNum, session.Version, token, mdl.GetUnderlyingToken(), amount)
		session.CollateralInUnderlying += utils.GetFloat64Decimal(valueInUnderlyingAsset, mdl.GetUnderlyingDecimal())
	}
}

func (mdl *CreditManager) onIncreaseBorrowedAmount(txLog *types.Log, borrower string, amount *big.Int) error {
	// manager state
	mdl.addBorrowAmountForBlock(amount)
	// other operations
	sessionId := mdl.GetCreditOwnerSession(borrower)
	blockNum := int64(txLog.BlockNumber)
	action, args := mdl.ParseEvent("IncreaseBorrowedAmount", txLog)
	// add account operation
	accountOperation := &schemas.AccountOperation{
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
	mdl.AddAccountOperation(accountOperation)
	session := mdl.Repo.UpdateCreditSession(sessionId, nil)
	session.BorrowedAmount = (*core.BigInt)(new(big.Int).Add(session.BorrowedAmount.Convert(), amount))
	mdl.setUpdateSession(sessionId)
	return nil
}

func (mdl *CreditManager) onTransferAccount(txLog *types.Log, owner, newOwner string) error {
	sessionId := mdl.GetCreditOwnerSession(owner)
	action, args := mdl.ParseEvent("TransferAccount", txLog)
	// add account operation
	accountOperation := &schemas.AccountOperation{
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
	mdl.Repo.UpdateCreditSession(sessionId, map[string]interface{}{"Borrower": newOwner})
	return nil
}

func (mdl *CreditManager) AddExecuteParams(txLog *types.Log,
	borrower,
	targetContract common.Address) error {
	sessionId := mdl.GetCreditOwnerSession(borrower.Hex())
	blockNum := int64(txLog.BlockNumber)
	mdl.executeParams = append(mdl.executeParams, ds.ExecuteParams{
		SessionId:     sessionId,
		CreditAccount: common.HexToAddress(strings.Split(sessionId, "_")[0]),
		Protocol:      targetContract,
		Borrower:      borrower,
		Index:         txLog.Index,
		BlockNumber:   blockNum,
	})
	return nil
}

func (mdl *CreditManager) saveExecuteEvents(executeParams []ds.ExecuteParams) {
	// credit manager has the execute event
	calls := mdl.Repo.GetExecuteParser().GetExecuteCalls(mdl.LastTxHash, mdl.Address, executeParams)

	for i, call := range calls {
		params := executeParams[i]

		accountOperation := &schemas.AccountOperation{
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
		}
		mdl.AddAccountOperation(accountOperation)
		mdl.setUpdateSession(params.SessionId)
	}
}

func (mdl *CreditManager) AddAccountOperation(accountOperation *schemas.AccountOperation) {
	mdl.Repo.AddAccountOperation(accountOperation)
}

func (mdl *CreditManager) PoolBorrow(txLog *types.Log, sessionId, borrower string, amount *big.Int) {
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

func (mdl *CreditManager) PoolRepay(blockNum int64, logId uint, txHash, sessionId, borrower string, amount *big.Int) {
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
