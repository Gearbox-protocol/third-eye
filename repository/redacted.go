package repository

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/utils"
	"math/big"
)

func (repo *Repository) UpdateBalance(eb *core.EventBalance) {
	lastCSS := repo.GetLastCSS(eb.SessionId)
	lastCSS.BlockNum = eb.BlockNumber
	// lastCSS.LogId = eb.Index
	lastCSS.Borrower = eb.Borrower
	if !eb.Clear {
		if eb.BorrowedAmount != nil {
			var newBorrowedAmount *big.Int
			if lastCSS.BorrowedAmountBI != nil {
				newBorrowedAmount = (new(big.Int).Add(lastCSS.BorrowedAmountBI.Convert(), eb.BorrowedAmount))
			} else {
				newBorrowedAmount = eb.BorrowedAmount
			}
			lastCSS.BorrowedAmountBI = (*core.BigInt)(newBorrowedAmount)
			lastCSS.BorrowedAmount = utils.GetFloat64Decimal(newBorrowedAmount, repo.GetUnderlyingDecimal(eb.CreditManager))
		}
		oldBalances := *lastCSS.Balances
		for tokenAddr, amount := range eb.Transfers {
			tokenBStruct := oldBalances[tokenAddr]
			token := repo.GetToken(tokenAddr)
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
		lastCSS.Balances = &oldBalances
	} else {
		if eb.BorrowedAmount == nil {
			lastCSS.BorrowedAmountBI = nil
			lastCSS.BorrowedAmount = 0
		} else {
			lastCSS.BorrowedAmountBI = (*core.BigInt)(eb.BorrowedAmount)
			lastCSS.BorrowedAmount = utils.GetFloat64Decimal(eb.BorrowedAmount, repo.GetUnderlyingDecimal(eb.CreditManager))
		}
		newBalances := core.JsonBalance{}
		for tokenAddr, amount := range eb.Transfers {
			token := repo.GetToken(tokenAddr)
			newBalances[tokenAddr] = &core.BalanceType{
				BI: (*core.BigInt)(amount),
				F:  utils.GetFloat64Decimal(amount, token.Decimals),
			}
		}
		lastCSS.Balances = &newBalances
	}

	newCSS := core.CreditSessionSnapshot{}
	newBalances := core.JsonBalance{}
	for tokenAddr, details := range *lastCSS.Balances {
		amt := *(details.BI.Convert())
		newBalances[tokenAddr] = &core.BalanceType{
			BI: (*core.BigInt)(&amt),
			F:  details.F,
		}
	}
	newCSS.Balances = &newBalances
	// newCSS.LogId = lastCSS.LogId
	newCSS.BlockNum = lastCSS.BlockNum
	newCSS.SessionId = lastCSS.SessionId
	newCSS.СumulativeIndexAtOpen = lastCSS.СumulativeIndexAtOpen
	if lastCSS.BorrowedAmountBI != nil {
		newBorrowBI := *lastCSS.BorrowedAmountBI
		newCSS.BorrowedAmountBI = &newBorrowBI
	}
	newCSS.BorrowedAmount = lastCSS.BorrowedAmount
	repo.AddCreditSessionSnapshot(&newCSS)
}
