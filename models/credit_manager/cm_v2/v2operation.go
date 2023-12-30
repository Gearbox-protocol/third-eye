package cm_v2

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/credit_manager/cm_common"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// /////////////////////
// Main actions
// /////////////////////
func (mdl *CMv2) onOpenCreditAccountV2(txLog *types.Log, onBehalfOf, account string,
	borrowAmount *big.Int,
	referralCode uint16) error {
	mdl.CMStatsOnOpenAccount(borrowAmount)
	// other operations
	cfAddr := txLog.Address.Hex()
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
			mdl.GetUnderlyingToken(): borrowAmount,
		},
		Dapp: cfAddr,
	}
	mdl.MulticallMgr.AddOpenEvent(accountOperation)
	mdl.PoolBorrow(txLog, sessionId, onBehalfOf, borrowAmount)
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
		BorrowedAmount: (*core.BigInt)(borrowAmount),
		IsDirty:        true,
		Version:        core.NewVersion(2),
	}
	mdl.Repo.AddCreditSession(newSession, false, txLog.TxHash.Hex(), txLog.Index)
	return nil
}

// while closing funds can be transferred from the owner account too
// https://github.com/Gearbox-protocol/contracts-v2/blob/main/contracts/credit/CreditManager.sol#L286-L291
func (mdl *CMv2) onCloseCreditAccountV2(txLog *types.Log, owner, to string) {
	mdl.State.TotalClosedAccounts++ // update totalclosedStats
	sessionId := mdl.GetCreditOwnerSession(owner)
	account := strings.Split(sessionId, "_")[0]
	cfAddr := txLog.Address.Hex()
	blockNum := int64(txLog.BlockNumber)

	//////////
	// get token transfer when account was closed
	txTransfers := mdl.Repo.GetExecuteParser().GetTransfersAtClosev2(txLog.TxHash.Hex(), account, mdl.GetUnderlyingToken(), ds.BorrowerAndTo{
		Borrower: common.HexToAddress(owner),
		To:       common.HexToAddress(to),
	})
	userTransfers := cm_common.ToJsonBalanceWithRepo(txTransfers, mdl.Repo)
	//////////
	// calculate remainingFunds
	var tokens []string
	for token := range userTransfers {
		tokens = append(tokens, token)
	}
	tokens = append(tokens, mdl.GetUnderlyingToken())
	prices := mdl.Repo.GetPricesInUSD(blockNum, tokens)
	//
	remainingFunds := (userTransfers.ValueInUnderlying(
		mdl.GetUnderlyingToken(), mdl.GetUnderlyingDecimal(), prices))
	//////////
	// use remainingFunds
	action, args := mdl.ParseEvent("CloseCreditAccount", txLog)
	(*args)["remainingFunds"] = (*core.BigInt)(remainingFunds)
	accountOperation := &schemas.AccountOperation{ // add account operation
		TxHash:      txLog.TxHash.Hex(),
		BlockNumber: blockNum,
		LogId:       txLog.Index,
		Borrower:    owner,
		SessionId:   sessionId,
		AdapterCall: false,
		Action:      action,
		Args:        args,
		Transfers:   &txTransfers,
		Dapp:        cfAddr,
	}

	////////////////////
	//// update from there
	////////////////////

	session := mdl.Repo.UpdateCreditSession(sessionId, nil) // update session
	session.CloseTransfers = &userTransfers

	mdl.MulticallMgr.AddCloseEvent(accountOperation) // add event to multicall processor
	session.RemainingFunds = (*core.BigInt)(remainingFunds)

	mdl.SetSessionIsClosed(sessionId, &cm_common.SessionCloseDetails{ // update closeSession map with session details
		LogId:          txLog.Index,
		RemainingFunds: remainingFunds,
		Status:         schemas.Closed,
		TxHash:         txLog.TxHash.Hex(),
		Borrower:       owner,
	})

	mdl.RemoveCreditOwnerSession(owner) // remove session to manager object
	mdl.CloseAccount(sessionId, blockNum, txLog.TxHash.Hex(), txLog.Index)
}

func (mdl *CMv2) onLiquidateCreditAccountV2(txLog *types.Log, owner, liquidator string, remainingFunds *big.Int) {
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
	// add event to multicall processor
	mdl.MulticallMgr.AddLiquidateEvent(accountOperation)
	mdl.SetSessionIsClosed(sessionId, &cm_common.SessionCloseDetails{
		LogId:          txLog.Index,
		RemainingFunds: remainingFunds,
		Status:         schemas.Liquidated,
		TxHash:         txLog.TxHash.Hex(),
		Borrower:       owner,
	})
	session := mdl.Repo.GetCreditSession(sessionId)
	session.Liquidator = liquidator
	session.RemainingFunds = (*core.BigInt)(remainingFunds)
	// remove session to manager object
	mdl.RemoveCreditOwnerSession(owner)
	mdl.CloseAccount(sessionId, blockNum, txLog.TxHash.Hex(), txLog.Index)
}

// /////////////////////
// Side actions that can also be used as multicall events
// /////////////////////
func (mdl *CMv2) onAddCollateralV2(txLog *types.Log, onBehalfOf, token string, value *big.Int) {
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
	mdl.MulticallMgr.AddMulticallEvent(accountOperation)
	mdl.AddCollateralToSession(blockNum, sessionId, token, value)
}

// amount can be negative, if decrease borrowamount, add pool repay event
func (mdl *CMv2) onIncreaseBorrowedAmountV2(txLog *types.Log, borrower string, amount *big.Int, eventName string) error {
	// manager state
	if amount.Sign() == 1 {
		mdl.AddBorrowAmountForBlock(amount)
	}
	// other operations
	sessionId := mdl.GetCreditOwnerSession(borrower)
	blockNum := int64(txLog.BlockNumber)
	action, args := mdl.ParseEvent(eventName, txLog)
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
	mdl.MulticallMgr.AddMulticallEvent(accountOperation)
	if amount.Sign() == -1 {
		repayAmount := new(big.Int).Neg(amount)
		// manager state
		mdl.PoolRepay(blockNum, txLog.Index, txLog.TxHash.Hex(), sessionId, borrower, repayAmount)
	} else {
		mdl.PoolBorrow(txLog, sessionId, borrower, amount)
	}
	session := mdl.Repo.UpdateCreditSession(sessionId, nil)
	session.BorrowedAmount = (*core.BigInt)(new(big.Int).Add(session.BorrowedAmount.Convert(), amount))
	return nil
}

func (mdl *CMv2) AddExecuteParamsV2(txLog *types.Log,
	borrower,
	targetContract common.Address) error {
	sessionId := mdl.GetCreditOwnerSession(borrower.Hex(), true) // for borrower = creditfacade, session id is ""
	mdl.MulticallMgr.AddMulticallEvent(&schemas.AccountOperation{
		BlockNumber: int64(txLog.BlockNumber),
		TxHash:      txLog.TxHash.Hex(),
		LogId:       txLog.Index,
		Borrower:    borrower.Hex(),
		Dapp:        targetContract.Hex(),
		AdapterCall: true,
		SessionId:   sessionId,
		Action:      "ExecuteOrder",
	})
	return nil
}

// copied from v1
func (mdl *CMv2) onTransferAccountV2(txLog *types.Log, owner, newOwner string) error {
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

func (mdl *CMv2) enableOrDisableToken(txLog types.Log, action string) {
	borrower := common.BytesToAddress(txLog.Topics[1][:]).Hex()
	token := common.BytesToAddress(txLog.Topics[2][:]).Hex()
	//
	sessionId := mdl.GetCreditOwnerSession(borrower)
	//
	accountOperation := &schemas.AccountOperation{
		TxHash:      txLog.TxHash.Hex(),
		BlockNumber: int64(txLog.BlockNumber),
		LogId:       txLog.Index,
		Borrower:    borrower,
		SessionId:   sessionId,
		AdapterCall: false,
		Action:      action,
		Args:        &core.Json{"token": token},
		Dapp:        txLog.Address.Hex(),
	}
	mdl.MulticallMgr.AddMulticallEvent(accountOperation)
}
