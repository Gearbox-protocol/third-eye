package credit_manager

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/utils"
	"math/big"
)

func (mdl *CreditManager) FetchFromDCForChangedSessions(blockNum int64) {
	for sessionId, _ := range mdl.UpdatedSessions {
		if mdl.ClosedSessions[sessionId] == nil {
			mdl.updateSession(sessionId, blockNum)
		}
	}
	for sessionId, closeDetails := range mdl.ClosedSessions {
		updates := mdl.UpdatedSessions[sessionId]
		if updates != 0 {
			log.Fatal("Session: %s updated %d before close %+v in block %d\n", sessionId, updates, closeDetails, blockNum)
		}
		mdl.closeSession(sessionId, blockNum, closeDetails)
	}
	mdl.UpdatedSessions = make(map[string]int)
	mdl.ClosedSessions = make(map[string]*SessionCloseDetails)
}

func (mdl *CreditManager) closeSession(sessionId string, blockNum int64, closeDetails *SessionCloseDetails) {
	mdl.State.OpenedAccountsCount--
	// check the data before credit session was closed by minus 1.
	session := mdl.Repo.GetCreditSession(sessionId)
	// set session fields
	session.ClosedAt = blockNum
	// this checks prevent getting data for credit session that exist only within a block
	// datacompressor query will fail
	if session.Since == session.ClosedAt {
		return
	}
	data := mdl.GetCreditSessionData(blockNum-1, session.Borrower)
	session.Status = closeDetails.Status
	session.HealthFactor = data.HealthFactor.Int64()
	session.BorrowedAmount = (*core.BigInt)(data.BorrowedAmount)
	session.TotalValueBI = (*core.BigInt)(data.TotalValue)
	if closeDetails.RemainingFunds == nil && closeDetails.Status == core.Repaid {
		closeDetails.RemainingFunds = new(big.Int).Sub(data.TotalValue, data.RepayAmount)
		mdl.Repo.AddEventBalance(core.NewEventBalance(blockNum,
			closeDetails.LogId,
			sessionId,
			nil,
			core.Transfers{
				mdl.GetUnderlyingToken(): closeDetails.RemainingFunds,
			},
			true,
			mdl.GetAddress()))
	}
	profit := new(big.Int).Sub(closeDetails.RemainingFunds, (*big.Int)(session.InitialAmount))
	session.Profit = (*core.BigInt)(profit)
	session.ProfitPercentage = float64(new(big.Int).Div(new(big.Int).
		Mul(profit, big.NewInt(100000)), (*big.Int)(session.InitialAmount)).Int64()) / 1000

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

	// create session snapshot
	css := core.CreditSessionSnapshot{}
	css.BlockNum = blockNum - 1
	css.SessionId = sessionId
	css.Borrower = session.Borrower
	css.HealthFactor = session.HealthFactor
	css.TotalValueBI = core.NewBigInt(session.TotalValueBI)
	css.TotalValue = utils.GetFloat64Decimal(data.TotalValue, mdl.GetUnderlyingDecimal())
	css.Balances = mdl.Repo.ConvertToBalance(data.Balances)
	css.BorrowedAmountBI = core.NewBigInt(session.BorrowedAmount)
	css.BorrowedAmount = utils.GetFloat64Decimal(data.BorrowedAmount, mdl.GetUnderlyingDecimal())
	css.СumulativeIndexAtOpen = core.NewBigInt((*core.BigInt)(data.CumulativeIndexAtOpen))
	log.Info(css)
	mdl.Repo.AddCreditSessionSnapshot(&css)
}

func (mdl *CreditManager) updateSession(sessionId string, blockNum int64) {
	session := mdl.Repo.GetCreditSession(sessionId)
	data := mdl.GetCreditSessionData(blockNum, session.Borrower)
	session.HealthFactor = data.HealthFactor.Int64()
	session.BorrowedAmount = (*core.BigInt)(data.BorrowedAmount)
	session.TotalValueBI = (*core.BigInt)(data.TotalValue)
	extraFunds := new(big.Int).Sub(data.TotalValue, data.BorrowedAmountPlusInterest)
	session.Profit = (*core.BigInt)(new(big.Int).Sub(extraFunds, (*big.Int)(session.InitialAmount)))
	session.ProfitPercentage = float64(new(big.Int).Div(new(big.Int).
		Mul((*big.Int)(session.Profit), big.NewInt(100000)), (*big.Int)(session.InitialAmount)).Int64()) / 1000

	// create session snapshot
	css := core.CreditSessionSnapshot{}
	css.BlockNum = blockNum
	css.SessionId = sessionId
	css.Borrower = session.Borrower
	css.HealthFactor = session.HealthFactor
	css.TotalValueBI = core.NewBigInt(session.TotalValueBI)
	css.TotalValue = utils.GetFloat64Decimal(data.TotalValue, mdl.GetUnderlyingDecimal())
	css.Balances = mdl.Repo.ConvertToBalance(data.Balances)
	css.BorrowedAmountBI = core.NewBigInt(session.BorrowedAmount)
	css.BorrowedAmount = utils.GetFloat64Decimal(data.BorrowedAmount, mdl.GetUnderlyingDecimal())
	css.СumulativeIndexAtOpen = core.NewBigInt((*core.BigInt)(data.CumulativeIndexAtOpen))
	mdl.Repo.AddCreditSessionSnapshot(&css)
}
