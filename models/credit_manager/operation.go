package credit_manager

import (
	"fmt"
	"github.com/Gearbox-protocol/gearscan/core"
	"github.com/Gearbox-protocol/gearscan/utils"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"
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
	mdl.Repo.AddCreditOwnerSession(cmAddr, onBehalfOf, sessionId)
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
	mdl.updateBalances(txLog, sessionId, borrowAmount,  map[string]*big.Int{
		mdl.UToken: new(big.Int).Add(borrowAmount, amount),
	}, false)
	return nil
}

// onCloseCreditAccount handles CloseCreditAccount Event
func (mdl *CreditManager) onCloseCreditAccount(txLog *types.Log, owner, to string, remainingFunds *big.Int) error {
	cmAddr := txLog.Address.Hex()
	sessionId := mdl.Repo.GetCreditOwnerSession(cmAddr, owner)
	blockNum := int64(txLog.BlockNumber)
	action, args := mdl.ParseEvent("CloseCreditAccount", txLog)
	accountOperation := &core.AccountOperation{
		TxHash:          txLog.TxHash.Hex(),
		BlockNumber: blockNum,
		LogId: txLog.Index,
		Borrower: owner,
		SessionId: sessionId,
		AdapterCall: false,
		Action: action,
		Args: args,
		Transfers: "",
		Dapp: cmAddr,
	}
	mdl.Repo.AddAccountOperation(accountOperation)
	mdl.Repo.RemoveCreditOwnerSession(cmAddr, owner)
	mdl.closeSession(sessionId, blockNum, remainingFunds, core.Closed)
	// create credit session snapshot
	mdl.updateBalances(txLog, sessionId, nil, map[string]*big.Int{
		mdl.UToken: remainingFunds,
	}, true)
	return nil
}

func (mdl *CreditManager) onLiquidateCreditAccount(txLog *types.Log, owner, liquidator string, remainingFunds *big.Int) error {
	cmAddr := txLog.Address.Hex()
	sessionId := mdl.Repo.GetCreditOwnerSession(cmAddr, owner)
	blockNum := int64(txLog.BlockNumber)
	action, args := mdl.ParseEvent("LiquidateCreditAccount", txLog)
	accountOperation := &core.AccountOperation{
		TxHash:          txLog.TxHash.Hex(),
		BlockNumber:       blockNum,
		LogId: txLog.Index,
		Borrower: owner,
		SessionId: sessionId,
		AdapterCall: false,
		Action: action,
		Args: args,
		Transfers: "",
		Dapp: txLog.Address.Hex(),
	}
	mdl.Repo.AddAccountOperation(accountOperation)
	mdl.Repo.RemoveCreditOwnerSession(cmAddr, owner)
	mdl.closeSession(sessionId, blockNum, remainingFunds, core.Liquidated)
	// create credit session snapshot
	mdl.updateBalances(txLog, sessionId, nil, map[string]*big.Int{
		mdl.UToken: remainingFunds,
	}, true)
	return nil
}

func (mdl *CreditManager) onRepayCreditAccount(txLog *types.Log, owner, to string) error {
	cmAddr := txLog.Address.Hex()
	sessionId := mdl.Repo.GetCreditOwnerSession(cmAddr, owner)
	blockNum := int64(txLog.BlockNumber)
	action, args := mdl.ParseEvent("RepayCreditAccount", txLog)
	accountOperation := &core.AccountOperation{
		TxHash:          txLog.TxHash.Hex(),
		BlockNumber:        blockNum,
		LogId: txLog.Index,
		Borrower: owner,
		SessionId: sessionId,
		AdapterCall: false,
		Action: action,
		Args: args,
		Transfers: "",
		Dapp: cmAddr,
	}
	mdl.Repo.AddAccountOperation(accountOperation)
	mdl.Repo.RemoveCreditOwnerSession(cmAddr, owner)
	mdl.closeSession(sessionId, blockNum, nil, core.Repaid)
	return nil
}

func (mdl *CreditManager) onAddCollateral(txLog *types.Log, onBehalfOf, token string, value *big.Int) error {
	cmAddr := txLog.Address.Hex()
	sessionId := mdl.Repo.GetCreditOwnerSession(cmAddr, onBehalfOf)
	blockNum := int64(txLog.BlockNumber)
	action, args := mdl.ParseEvent("AddCollateral", txLog)
	accountOperation := &core.AccountOperation{
		TxHash:          txLog.TxHash.Hex(),
		BlockNumber:        blockNum,
		LogId: txLog.Index,
		Borrower: onBehalfOf,
		SessionId: sessionId,
		AdapterCall: false,
		Action: action,
		Args: args,
		Transfers: "",
		Dapp: txLog.Address.Hex(),
	}
	mdl.Repo.AddAccountOperation(accountOperation)
	mdl.Repo.RemoveCreditOwnerSession(cmAddr, onBehalfOf)
	mdl.updateSession(sessionId, blockNum)
	// create credit session snapshot
	mdl.updateBalances(txLog, sessionId, nil, map[string]*big.Int{
		token: value,
	}, false)
	return nil
}

func (mdl *CreditManager) onIncreaseBorrowedAmount(txLog *types.Log, borrower string, amount *big.Int) error {
	cmAddr := txLog.Address.Hex()
	sessionId := mdl.Repo.GetCreditOwnerSession(cmAddr, borrower)
	blockNum := int64(txLog.BlockNumber)
	action, args := mdl.ParseEvent("IncreaseBorrowedAmount", txLog)
	accountOperation := &core.AccountOperation{
		TxHash:          txLog.TxHash.Hex(),
		BlockNumber:        blockNum,
		LogId: txLog.Index,
		Borrower: borrower,
		SessionId: sessionId,
		AdapterCall: false,
		Action: action,
		Args: args,
		Transfers: "",
		Dapp: txLog.Address.Hex(),
	}
	mdl.Repo.AddAccountOperation(accountOperation)
	mdl.updateSession(sessionId, blockNum)
	// create credit session snapshot
	mdl.updateBalances(txLog, sessionId, amount,  map[string]*big.Int{
		mdl.UToken: amount,
	}, false)
	return nil
}

func (mdl *CreditManager) onTransferAccount(txLog *types.Log, owner, newOwner string) error {
	cmAddr := txLog.Address.Hex()
	sessionId := mdl.Repo.GetCreditOwnerSession(cmAddr, owner)
	action, args := mdl.ParseEvent("TransferAccount", txLog)
	accountOperation := &core.AccountOperation{
		TxHash:          txLog.TxHash.Hex(),
		BlockNumber:        int64(txLog.BlockNumber),
		LogId: txLog.Index,
		Borrower: owner,
		SessionId: sessionId,
		AdapterCall: false,
		Action: action,
		Args: args,
		Transfers: "",
		Dapp: txLog.Address.Hex(),
	}
	mdl.Repo.AddAccountOperation(accountOperation)
	mdl.Repo.RemoveCreditOwnerSession(cmAddr, owner)
	mdl.Repo.AddCreditOwnerSession(cmAddr, newOwner, sessionId)
	session := mdl.Repo.GetCreditSession(sessionId)
	session.Borrower = newOwner
	return nil
}

func (mdl *CreditManager) onExecuteOrder(txLog *types.Log,
	borrower,
	targetContract common.Address) error {
	cmAddr := txLog.Address.Hex()
	sessionId := mdl.Repo.GetCreditOwnerSession(cmAddr, borrower.Hex())
	creditAccount := strings.Split(sessionId, "_")[0]
	calls := mdl.Repo.GetExecuteParser().GetExecuteCalls(txLog, common.HexToAddress(creditAccount), targetContract)
	
	for _, call := range calls {
		accountOperation := &core.AccountOperation{
			BlockNumber: int64(txLog.BlockNumber),
			TxHash: txLog.TxHash.Hex(),
			LogId: uint(call.LogId),
			// owner/account data
			Borrower: borrower.Hex(),
			SessionId: sessionId,
			// dapp
			Dapp: targetContract.Hex(),
			// call/events data
			Action: call.Name,
			Args: call.Args,
			AdapterCall: true,
			Transfers: call.Balances.String(),
			// extras
			Depth: call.Depth,
		}
		mdl.Repo.AddAccountOperation(accountOperation)
		// create credit session snapshot
		mdl.updateBalances(txLog, sessionId, nil, (map[string]*big.Int)(call.Balances), false)
	}
	return nil
}

func (mdl *CreditManager) closeSession(sessionId string, blockNum int64, remainingFunds *big.Int, newStatus int) {
	// check the data before credit session was closed by minus 1.
	data := mdl.Repo.GetCreditSessionData(blockNum - 1, sessionId)
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
	data := mdl.Repo.GetCreditSessionData(blockNum, sessionId)
	session := mdl.Repo.GetCreditSession(sessionId)
	extraFunds := new(big.Int).Sub(data.TotalValue, data.BorrowedAmountPlusInterest)
	session.TotalValue = (*core.BigInt)(data.TotalValue)
	session.HealthFactor = data.HealthFactor.Int64()
	session.BorrowedAmount =  (*core.BigInt)(data.BorrowedAmount)
	session.Profit = (*core.BigInt)(new(big.Int).Sub(extraFunds, (*big.Int)(session.InitialAmount)))
	session.ProfitPercentage = float64(new(big.Int).Div(new(big.Int).
			Mul((*big.Int)(session.Profit), big.NewInt(100000)), (*big.Int)(session.InitialAmount)).Int64()) / 1000
}


func (mdl *CreditManager) updateBalances(txLog *types.Log, sessionId string, borrowedAmount *big.Int, balances map[string]*big.Int, clear bool) {
	lastCSS := mdl.Repo.GetLastCSS(sessionId)
	lastCSS.BlockNum = int64(txLog.BlockNumber)
	lastCSS.LogId = int64(txLog.Index)
	if !clear { 
		if borrowedAmount != nil {
			var newBorrowedAmount *big.Int
			if lastCSS.BorrowedAmountBI != nil {
				newBorrowedAmount = (new(big.Int).Add(lastCSS.BorrowedAmountBI.Convert(), borrowedAmount))
			} else {
				newBorrowedAmount = borrowedAmount
			}
			lastCSS.BorrowedAmountBI = (*core.BigInt)(newBorrowedAmount)
			lastCSS.BorrowedAmount = utils.GetFloat64Decimal(newBorrowedAmount, mdl.UDecimals)
		}
		oldBalances := lastCSS.Balances
		for tokenAddr, amount := range balances {
			tokenBStruct := oldBalances[tokenAddr]
			token := mdl.Repo.GetToken(tokenAddr)
			if amount.Sign() != 0 {
				if oldBalances[tokenAddr] != nil {
					newAmt := new(big.Int).Add(tokenBStruct.BI.Convert(), amount)
					oldBalances[tokenAddr] = &core.BalanceType{
						BI: (*core.BigInt)(newAmt),
						F: utils.GetFloat64Decimal(newAmt, token.Decimals),
					}
				} else {
					oldBalances[tokenAddr] = &core.BalanceType{
						BI: (*core.BigInt)(amount),
						F: utils.GetFloat64Decimal(amount, token.Decimals),
					}
				}
			}
		}
		lastCSS.Balances = oldBalances
	} else {
		if borrowedAmount == nil {
			lastCSS.BorrowedAmountBI = nil
			lastCSS.BorrowedAmount = 0
		} else {
			lastCSS.BorrowedAmountBI = (*core.BigInt)(borrowedAmount)
			lastCSS.BorrowedAmount = utils.GetFloat64Decimal(borrowedAmount, mdl.UDecimals)
		}
		newBalances := core.JsonBalance{}
		for tokenAddr, amount := range balances {
			token := mdl.Repo.GetToken(tokenAddr)
			newBalances[tokenAddr] = &core.BalanceType{
				BI: (*core.BigInt)(amount),
				F: utils.GetFloat64Decimal(amount, token.Decimals),
			}
		}
		lastCSS.Balances = newBalances
	}
	

	newCSS := core.CreditSessionSnapshot{}
	newBalances := core.JsonBalance{}
	for tokenAddr, details := range lastCSS.Balances {
		amt := *(details.BI.Convert())
		newBalances[tokenAddr] = &core.BalanceType {
			BI: (*core.BigInt)(&amt),
			F: details.F,
		}
	}
	newCSS.Balances = newBalances
	newCSS.LogId = lastCSS.LogId
	newCSS.BlockNum = lastCSS.BlockNum
	newCSS.SessionId = lastCSS.SessionId
	newBorrowBI := *lastCSS.BorrowedAmountBI
	newCSS.BorrowedAmountBI = &newBorrowBI
	newCSS.BorrowedAmount = lastCSS.BorrowedAmount
	mdl.Repo.AddCreditSessionSnapshot(&newCSS)
}