package cm_v3

import (
	"fmt"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/credit_manager/cm_common"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (mdl *CMv3) onOpenCreditAccountV3(txLog *types.Log, onBehalfOf, account string,
	referralCode *big.Int) {
	// default is zero in v3
	borrowAmount := new(big.Int)
	//
	mdl.CMStatsOnOpenAccount(borrowAmount)
	//
	cfAddr := txLog.Address.Hex()
	sessionId := fmt.Sprintf("%s_%d_%d", account, txLog.BlockNumber, txLog.Index)
	blockNum := int64(txLog.BlockNumber)

	//
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
	// add account
	mdl.AddCreditAccount(account, sessionId, onBehalfOf)

	newSession := &schemas.CreditSession{
		ID:             sessionId,
		Status:         schemas.Active,
		Borrower:       onBehalfOf,
		CreditManager:  mdl.Address,
		Account:        account,
		Since:          blockNum,
		BorrowedAmount: (*core.BigInt)(borrowAmount),
		IsDirty:        true,
		Version:        core.NewVersion(300),
	}
	// direct token manager
	mdl.Repo.AddCreditSession(newSession, false, txLog.TxHash.Hex(), txLog.Index)
}

func (mdl *CMv3) onCloseCreditAccountV3(txLog *types.Log, creditAccount, to string) {
	mdl.State.TotalClosedAccounts++ // update totalclosedStats
	sessionId, owner := mdl.GetSessionIdAndBorrower(creditAccount)
	cfAddr := txLog.Address.Hex()
	blockNum := int64(txLog.BlockNumber)

	//////////
	// get token transfer when account was closed
	txTransfers := mdl.Repo.GetExecuteParser().GetTransfersAtClosev2(txLog.TxHash.Hex(), creditAccount, mdl.GetUnderlyingToken(), ds.BorrowerAndTo{
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

	mdl.MulticallMgr.AddCloseOrLiquidateEvent(accountOperation) // add event to multicall processor
	session.RemainingFunds = (*core.BigInt)(remainingFunds)

	mdl.SetSessionIsClosed(sessionId, &cm_common.SessionCloseDetails{ // update closeSession map with session details
		LogId:          txLog.Index,
		RemainingFunds: remainingFunds,
		Status:         schemas.Closed,
		TxHash:         txLog.TxHash.Hex(),
		Borrower:       owner,
	})

	mdl.RemoveCreditAccount(creditAccount) // remove session to manager object
	mdl.CloseAccount(sessionId, blockNum, txLog.TxHash.Hex(), txLog.Index)
}

func (mdl *CMv3) onLiquidateCreditAccountV3(txLog *types.Log, creditAccount, liquidator string, closeAction uint8, remainingFunds *big.Int) {
	mdl.State.TotalLiquidatedAccounts++
	sessionId, owner := mdl.GetSessionIdAndBorrower(creditAccount)
	blockNum := int64(txLog.BlockNumber)

	//
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
	status := func() int {
		if mdl.State.Paused {
			return schemas.LiquidateExpired
		}
		if closeAction == 1 {
			return schemas.Liquidated
		} else if closeAction == 2 {
			return schemas.LiquidateExpired
		}
		log.Fatal("Wrong status")
		return 0
	}()
	mdl.MulticallMgr.AddCloseOrLiquidateEvent(accountOperation)
	mdl.SetSessionIsClosed(sessionId, &cm_common.SessionCloseDetails{
		LogId:          txLog.Index,
		RemainingFunds: remainingFunds,
		Status:         status,
		TxHash:         txLog.TxHash.Hex(),
		Borrower:       owner,
	})
	session := mdl.Repo.GetCreditSession(sessionId)
	session.Liquidator = liquidator
	session.RemainingFunds = (*core.BigInt)(remainingFunds)
	// remove session to manager object
	mdl.RemoveCreditAccount(owner)
	mdl.CloseAccount(sessionId, blockNum, txLog.TxHash.Hex(), txLog.Index) // for direct token transfer manager
}

// /////////////////////
// Side actions that can also be used as multicall events
// /////////////////////
func (mdl *CMv3) onAddCollateralV3(txLog *types.Log, creditAccount, token string, value *big.Int) {
	sessionId, owner := mdl.GetSessionIdAndBorrower(creditAccount)
	blockNum := int64(txLog.BlockNumber)
	action, args := mdl.ParseEvent("AddCollateral", txLog)
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
			token: value,
		},
		Dapp: txLog.Address.Hex(),
	}
	mdl.MulticallMgr.AddMulticallEvent(accountOperation)
	mdl.AddCollateralToSession(blockNum, sessionId, token, value)
}

// amount can be negative, if decrease borrowamount, add pool repay event
func (mdl *CMv3) onIncreaseBorrowedAmountV3(txLog *types.Log, creditAccount string, amount *big.Int, eventName string) error {
	// manager state
	if amount.Sign() == 1 {
		mdl.AddBorrowAmountForBlock(amount)
	}
	// other operations
	sessionId, borrower := mdl.GetSessionIdAndBorrower(creditAccount)
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
		mdl.State.TotalRepaidBI = core.AddCoreAndInt(mdl.State.TotalRepaidBI, repayAmount)
		mdl.State.TotalRepaid = utils.GetFloat64Decimal(mdl.State.TotalRepaidBI.Convert(), mdl.GetUnderlyingDecimal())
		mdl.PoolRepay(blockNum, txLog.Index, txLog.TxHash.Hex(), sessionId, borrower, repayAmount)
	} else {
		mdl.PoolBorrow(txLog, sessionId, borrower, amount)
	}
	session := mdl.Repo.UpdateCreditSession(sessionId, nil)
	session.BorrowedAmount = (*core.BigInt)(new(big.Int).Add(session.BorrowedAmount.Convert(), amount))
	return nil
}

func (mdl *CMv3) AddExecuteParamsV3(txLog *types.Log,
	creditAccount,
	targetContract common.Address) error {
	sessionId, borrower := mdl.GetSessionIdAndBorrower(creditAccount.Hex(), true) // for borrower = creditfacade, session id is ""
	mdl.MulticallMgr.AddMulticallEvent(&schemas.AccountOperation{
		BlockNumber: int64(txLog.BlockNumber),
		TxHash:      txLog.TxHash.Hex(),
		LogId:       txLog.Index,
		Borrower:    borrower,
		Dapp:        targetContract.Hex(),
		AdapterCall: true,
		SessionId:   sessionId,
		Action:      "ExecuteOrder",
	})
	return nil
}
