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
	mdl.addEventBalance(newEventBalance(
		txLog.BlockNumber,
		txLog.Index,
		sessionId,
		borrowAmount,
		map[string]*big.Int{
			mdl.GetUnderlyingToken(): new(big.Int).Add(borrowAmount, amount),
		},
		false,
	))

	return nil
}

// onCloseCreditAccount handles CloseCreditAccount Event
func (mdl *CreditManager) onCloseCreditAccount(txLog *types.Log, owner, to string, remainingFunds *big.Int) error {
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
	mdl.addEventBalance(newEventBalance(
		txLog.BlockNumber,
		txLog.Index,
		sessionId,
		nil,
		map[string]*big.Int{
			mdl.GetUnderlyingToken(): remainingFunds,
		},
		true,
	))
	return nil
}

func (mdl *CreditManager) onLiquidateCreditAccount(txLog *types.Log, owner, liquidator string, remainingFunds *big.Int) error {
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
	mdl.addEventBalance(newEventBalance(
		txLog.BlockNumber,
		txLog.Index,
		sessionId,
		nil,
		map[string]*big.Int{
			mdl.GetUnderlyingToken(): remainingFunds,
		},
		true,
	))
	return nil
}

func (mdl *CreditManager) onRepayCreditAccount(txLog *types.Log, owner, to string) error {
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
	mdl.addEventBalance(newEventBalance(
		txLog.BlockNumber,
		txLog.Index,
		sessionId,
		nil,
		map[string]*big.Int{
			token: value,
		},
		false,
	))
	return nil
}

func (mdl *CreditManager) onIncreaseBorrowedAmount(txLog *types.Log, borrower string, amount *big.Int) error {
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
	mdl.addEventBalance(newEventBalance(
		txLog.BlockNumber,
		txLog.Index,
		sessionId,
		amount,
		map[string]*big.Int{
			mdl.GetUnderlyingToken(): amount,
		},
		false,
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
		mdl.addEventBalance(newEventBalance(
			uint64(params.BlockNumber),
			params.Index,
			params.SessionId,
			nil,
			(map[string]*big.Int)(call.Balances),
			false,
		))
	}
}

func (mdl *CreditManager) closeSession(sessionId string, blockNum int64, remainingFunds *big.Int, newStatus int) {
	// check the data before credit session was closed by minus 1.
	data := mdl.GetCreditSessionData(blockNum-1, sessionId)
	session := mdl.Repo.GetCreditSession(sessionId)
	session.ClosedAt = blockNum
	session.TotalValue = (*core.BigInt)(data.TotalValue)
	session.HealthFactor = data.HealthFactor.Int64()
	if remainingFunds == nil && newStatus == core.Repaid {
		remainingFunds = new(big.Int).Sub(data.TotalValue, data.RepayAmount)
	}
	session.Profit = (*core.BigInt)(new(big.Int).Sub(remainingFunds, (*big.Int)(session.InitialAmount)))
	// ToDo: Change to calculate correct values
	session.ProfitPercentage = float64(new(big.Int).Div(new(big.Int).
		Mul((*big.Int)(session.Profit), big.NewInt(100000)), (*big.Int)(session.InitialAmount)).Int64()) / 1000
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

func (mdl *CreditManager) updateBalance(eb *EventBalance) {
	lastCSS := mdl.Repo.GetLastCSS(eb.SessionId)
	lastCSS.BlockNum = eb.BlockNumber
	lastCSS.LogId = eb.Index
	if !eb.Clear {
		if eb.BorrowedAmount != nil {
			var newBorrowedAmount *big.Int
			if lastCSS.BorrowedAmountBI != nil {
				newBorrowedAmount = (new(big.Int).Add(lastCSS.BorrowedAmountBI.Convert(), eb.BorrowedAmount))
			} else {
				newBorrowedAmount = eb.BorrowedAmount
			}
			lastCSS.BorrowedAmountBI = (*core.BigInt)(newBorrowedAmount)
			lastCSS.BorrowedAmount = utils.GetFloat64Decimal(newBorrowedAmount, mdl.GetUnderlyingDecimal())
		}
		oldBalances := lastCSS.Balances
		for tokenAddr, amount := range eb.Transfers {
			tokenBStruct := oldBalances[tokenAddr]
			token := mdl.Repo.GetToken(tokenAddr)
			if amount.Sign() != 0 {
				if oldBalances[tokenAddr] != nil {
					newAmt := new(big.Int).Add(tokenBStruct.BI.Convert(), amount)
					oldBalances[tokenAddr] = &core.BalanceType{
						BI: (*core.BigInt)(newAmt),
						F:  utils.GetFloat64Decimal(newAmt, token.Decimals),
					}
				} else {
					oldBalances[tokenAddr] = &core.BalanceType{
						BI: (*core.BigInt)(amount),
						F:  utils.GetFloat64Decimal(amount, token.Decimals),
					}
				}
			}
		}
		lastCSS.Balances = oldBalances
	} else {
		if eb.BorrowedAmount == nil {
			lastCSS.BorrowedAmountBI = nil
			lastCSS.BorrowedAmount = 0
		} else {
			lastCSS.BorrowedAmountBI = (*core.BigInt)(eb.BorrowedAmount)
			lastCSS.BorrowedAmount = utils.GetFloat64Decimal(eb.BorrowedAmount, mdl.GetUnderlyingDecimal())
		}
		newBalances := core.JsonBalance{}
		for tokenAddr, amount := range eb.Transfers {
			token := mdl.Repo.GetToken(tokenAddr)
			newBalances[tokenAddr] = &core.BalanceType{
				BI: (*core.BigInt)(amount),
				F:  utils.GetFloat64Decimal(amount, token.Decimals),
			}
		}
		lastCSS.Balances = newBalances
	}

	newCSS := core.CreditSessionSnapshot{}
	newBalances := core.JsonBalance{}
	for tokenAddr, details := range lastCSS.Balances {
		amt := *(details.BI.Convert())
		newBalances[tokenAddr] = &core.BalanceType{
			BI: (*core.BigInt)(&amt),
			F:  details.F,
		}
	}
	newCSS.Balances = newBalances
	newCSS.LogId = lastCSS.LogId
	newCSS.BlockNum = lastCSS.BlockNum
	newCSS.SessionId = lastCSS.SessionId
	if lastCSS.BorrowedAmountBI != nil {
		newBorrowBI := *lastCSS.BorrowedAmountBI
		newCSS.BorrowedAmountBI = &newBorrowBI
	}
	newCSS.BorrowedAmount = lastCSS.BorrowedAmount
	mdl.Repo.AddCreditSessionSnapshot(&newCSS)
}
