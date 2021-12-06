package credit_manager

import (
	"fmt"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/services"
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
	action, args := mdl.ParseEvent("OpenCreditAccount", txLog)
	// add account operation
	accountOperation := &core.AccountOperation{
		TxHash:      txLog.TxHash.Hex(),
		BlockNumber: int64(txLog.BlockNumber),
		LogId:       txLog.Index,
		Borrower:    onBehalfOf,
		SessionId:   sessionId,
		AdapterCall: false,
		Action:      action,
		Args:        args,
		Transfers:   "",
		Dapp:        cmAddr,
	}
	mdl.Repo.AddAccountOperation(accountOperation)
	// add session to manager object
	mdl.AddCreditOwnerSession(onBehalfOf, sessionId)
	// create credit session
	newSession := &core.CreditSession{
		ID:             sessionId,
		Status:         core.Active,
		Borrower:       onBehalfOf,
		CreditManager:  mdl.Address,
		Account:        account,
		Since:          int64(txLog.BlockNumber),
		InitialAmount:  (*core.BigInt)(amount),
		BorrowedAmount: (*core.BigInt)(borrowAmount),
		Profit:         (*core.BigInt)(big.NewInt(0)),
	}
	mdl.Repo.AddCreditSession(newSession)
	// create CSS
	mdl.Repo.AddEventBalance(core.NewEventBalance(
		txLog.BlockNumber,
		txLog.Index,
		sessionId,
		borrowAmount,
		map[string]*big.Int{
			mdl.GetUnderlyingToken(): new(big.Int).Add(borrowAmount, amount),
		},
		false,
		cmAddr,
		onBehalfOf,
	))

	return nil
}

// onCloseCreditAccount handles CloseCreditAccount Event
func (mdl *CreditManager) onCloseCreditAccount(txLog *types.Log, owner, to string, remainingFunds *big.Int) error {
	mdl.State.TotalClosedAccounts++
	cmAddr := txLog.Address.Hex()
	sessionId := mdl.GetCreditOwnerSession(owner)
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
		Transfers:   "",
		Dapp:        cmAddr,
	}
	mdl.Repo.AddAccountOperation(accountOperation)
	// remove session to manager object
	mdl.RemoveCreditOwnerSession(owner)
	// close credit session
	mdl.closeSession(sessionId, blockNum, remainingFunds, core.Closed)
	// create CSS
	mdl.Repo.AddEventBalance(core.NewEventBalance(
		txLog.BlockNumber,
		txLog.Index,
		sessionId,
		nil,
		map[string]*big.Int{
			mdl.GetUnderlyingToken(): remainingFunds,
		},
		true,
		cmAddr,
		owner,
	))
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
		Transfers:   "",
		Dapp:        txLog.Address.Hex(),
	}
	mdl.Repo.AddAccountOperation(accountOperation)
	// remove session to manager object
	mdl.RemoveCreditOwnerSession(owner)
	// close credit session
	mdl.closeSession(sessionId, blockNum, remainingFunds, core.Liquidated)
	// create credit session snapshot
	mdl.Repo.AddEventBalance(core.NewEventBalance(
		txLog.BlockNumber,
		txLog.Index,
		sessionId,
		nil,
		map[string]*big.Int{
			mdl.GetUnderlyingToken(): remainingFunds,
		},
		true,
		mdl.GetAddress(),
		owner,
	))
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
		Transfers:   "",
		Dapp:        cmAddr,
	}
	mdl.Repo.AddAccountOperation(accountOperation)
	// remove session to manager object
	mdl.RemoveCreditOwnerSession(owner)
	// close credit session
	mdl.closeSession(sessionId, blockNum, nil, core.Repaid)
	return nil
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
		Transfers:   "",
		Dapp:        txLog.Address.Hex(),
	}
	mdl.Repo.AddAccountOperation(accountOperation)
	// update credit session
	mdl.updateSession(sessionId, blockNum)
	// create credit session snapshot
	mdl.Repo.AddEventBalance(core.NewEventBalance(
		txLog.BlockNumber,
		txLog.Index,
		sessionId,
		nil,
		map[string]*big.Int{
			token: value,
		},
		false,
		mdl.GetAddress(),
		onBehalfOf,
	))
	return nil
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
		Transfers:   "",
		Dapp:        txLog.Address.Hex(),
	}
	mdl.Repo.AddAccountOperation(accountOperation)
	// update credit session
	mdl.updateSession(sessionId, blockNum)
	// create credit session snapshot
	mdl.Repo.AddEventBalance(core.NewEventBalance(
		txLog.BlockNumber,
		txLog.Index,
		sessionId,
		amount,
		map[string]*big.Int{
			mdl.GetUnderlyingToken(): amount,
		},
		false,
		mdl.GetAddress(),
		borrower,
	))
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
		Transfers:   "",
		Dapp:        txLog.Address.Hex(),
	}
	mdl.Repo.AddAccountOperation(accountOperation)
	// remove session to manager object
	mdl.RemoveCreditOwnerSession(owner)
	mdl.AddCreditOwnerSession(newOwner, sessionId)
	session := mdl.Repo.GetCreditSession(sessionId)
	session.Borrower = newOwner
	return nil
}

func (mdl *CreditManager) AddParams(txLog *types.Log,
	borrower,
	targetContract common.Address) error {
	sessionId := mdl.GetCreditOwnerSession(borrower.Hex())
	mdl.executeParams = append(mdl.executeParams, services.ExecuteParams{
		SessionId:     sessionId,
		CreditAccount: common.HexToAddress(strings.Split(sessionId, "_")[0]),
		Protocol:      targetContract,
		Borrower:      borrower,
		Index:         txLog.Index,
		BlockNumber:   int64(txLog.BlockNumber),
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
			Transfers:   call.Balances.String(),
			// extras
			Depth: call.Depth,
		}
		mdl.Repo.AddAccountOperation(accountOperation)
		// create credit session snapshot
		mdl.Repo.AddEventBalance(core.NewEventBalance(
			uint64(params.BlockNumber),
			params.Index,
			params.SessionId,
			nil,
			(map[string]*big.Int)(call.Balances),
			false,
			mdl.GetAddress(),
			params.Borrower.Hex(),
		))
	}
}

func (mdl *CreditManager) closeSession(sessionId string, blockNum int64, remainingFunds *big.Int, newStatus int) {
	mdl.State.OpenedAccountsCount--
	// check the data before credit session was closed by minus 1.
	data := mdl.GetCreditSessionData(blockNum-1, sessionId)
	session := mdl.Repo.GetCreditSession(sessionId)
	session.ClosedAt = blockNum
	session.TotalValue = (*core.BigInt)(data.TotalValue)
	session.HealthFactor = data.HealthFactor.Int64()
	if remainingFunds == nil && newStatus == core.Repaid {
		remainingFunds = new(big.Int).Sub(data.TotalValue, data.RepayAmount)
	}
	profit := new(big.Int).Sub(remainingFunds, (*big.Int)(session.InitialAmount))
	session.Profit = (*core.BigInt)(profit)
	// credit manager state
	mdl.State.TotalRepaidBI = core.AddCoreAndInt(mdl.State.TotalRepaidBI, profit)
	mdl.State.TotalRepaid = utils.GetFloat64Decimal(mdl.State.TotalRepaidBI.Convert(), mdl.GetUnderlyingDecimal())
	if profit.Sign() < 0 {
		mdl.State.TotalLossesBI = core.AddCoreAndInt(mdl.State.TotalLossesBI, profit)
		mdl.State.TotalLosses = utils.GetFloat64Decimal(mdl.State.TotalLossesBI.Convert(), mdl.GetUnderlyingDecimal())
	} else {
		mdl.State.TotalProfitBI = core.AddCoreAndInt(mdl.State.TotalProfitBI, profit)
		mdl.State.TotalProfit = utils.GetFloat64Decimal(mdl.State.TotalProfitBI.Convert(), mdl.GetUnderlyingDecimal())
	}
	//-- credit manager state

	// ToDo: Change to calculate correct values
	session.ProfitPercentage = float64(new(big.Int).Div(new(big.Int).
		Mul(profit, big.NewInt(100000)), (*big.Int)(session.InitialAmount)).Int64()) / 1000
	session.Status = newStatus
}

func (mdl *CreditManager) updateSession(sessionId string, blockNum int64) {
	data := mdl.GetCreditSessionData(blockNum, sessionId)
	session := mdl.Repo.GetCreditSession(sessionId)
	extraFunds := new(big.Int).Sub(data.TotalValue, data.BorrowedAmountPlusInterest)
	session.TotalValue = (*core.BigInt)(data.TotalValue)
	session.HealthFactor = data.HealthFactor.Int64()
	session.BorrowedAmount = (*core.BigInt)(data.BorrowedAmount)
	session.Profit = (*core.BigInt)(new(big.Int).Sub(extraFunds, (*big.Int)(session.InitialAmount)))
	session.ProfitPercentage = float64(new(big.Int).Div(new(big.Int).
		Mul((*big.Int)(session.Profit), big.NewInt(100000)), (*big.Int)(session.InitialAmount)).Int64()) / 1000
}
