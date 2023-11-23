package debts

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/calc"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
)

// these values are calculated for the borrower not the liquidation, so the calrepayamount takes isliquidated = false
// only for block that the account is liquidated we use isLiquidated set to true so that we can calculate the true amountToPool
//
// amountToPool
// for v1,v2 amountToPool is calculated by calc.CalCloseAmount
//
// remainingfunds => is used for profit calculation, assets that user gets back.
// - v1 close or liquidated -- remainingfunds is taken from event
// - v1 repay + open -- remainingfunds is manually calculated with help of calCloseAmount in sdk-go
// - v2 for closeCreditAccount, remainingFunds is calculated from the account transfers
// - v2 for liquidateCreditAccount, remainingFunds is taken from event
// - v2 for openedAccounts, totalValue - amountToPool
//
// repayAmount => transfer from owner to account needed to close the account
// v1 - repayAmount = amountToPool, except the blockNum at which account is liquidated
//   - for liquidated account repay amount is amountToPool+ calcRemainginFunds https://github.com/Gearbox-protocol/gearbox-contracts/blob/master/contracts/credit/CreditManager.sol#L999
//   - NIT, closeAmount doesn't need repayAmount as all assets are converted to underlying token
//     so repayAmount is zero, https://github.com/Gearbox-protocol/gearbox-contracts/blob/master/contracts/credit/CreditManager.sol#L448-L465
//
// v2 - close repayAmount is transferred from borrower to account as underlying token
// v2 - for liquidation, repayAmount is zero.
// v2 - opened accounts
//
//	the account might be having some underlying token balance so repayAMount = amountToPool - underlyingToken balance
func (eng *DebtEngine) calAmountToPoolAndProfit(debt *schemas.Debt, session *schemas.CreditSession, cumIndexAndUToken *ds.CumIndexAndUToken, debtDetails *calc.DebtDetails) {
	var amountToPool, calRemainingFunds *big.Int
	sessionSnapshot := eng.lastCSS[session.ID]
	//
	// closedAt set and debt.BlockNumber is 1 before closure then status from session.Status
	// if liq v3 at blockNum +1 then status is Liquidated since we need to calculate remainingFunds and profit
	// if active
	status := session.StatusAt(debt.BlockNumber + 1)

	// amount to pool
	amountToPool, calRemainingFunds, _, _ = calc.CalCloseAmount(eng.lastParameters[session.CreditManager],
		session.Version, debt.CalTotalValueBI.Convert(), status,
		debtDetails)

	// calculate profit
	debt.AmountToPoolBI = (*core.BigInt)(amountToPool)
	var remainingFunds *big.Int
	if session.Version.Eq(300) {
		var repayAmount *big.Int
		repayAmount, remainingFunds = eng.remainingFundsv2(session, debt, cumIndexAndUToken, amountToPool, calRemainingFunds)
		debt.RepayAmountBI = (*core.BigInt)(repayAmount)
	} else if session.Version.Eq(2) {
		var repayAmount *big.Int
		repayAmount, remainingFunds = eng.remainingFundsv2(session, debt, cumIndexAndUToken, amountToPool, calRemainingFunds)
		debt.RepayAmountBI = (*core.BigInt)(repayAmount)
	} else {
		var repayAmount *big.Int
		repayAmount, remainingFunds = eng.remainingFundsv1(session, debt, amountToPool, calRemainingFunds)
		debt.RepayAmountBI = (*core.BigInt)(repayAmount)
	}

	remainingFundsInUSD := eng.GetAmountInUSD(cumIndexAndUToken.Token, remainingFunds, session.Version)
	debt.ProfitInUnderlying = utils.GetFloat64Decimal(remainingFunds, cumIndexAndUToken.Decimals) - debt.CollateralInUnderlying
	// debt.CollateralInUnderlying = sessionSnapshot.CollateralInUnderlying
	// fields in USD
	debt.CollateralInUSD = sessionSnapshot.CollateralInUSD
	debt.ProfitInUSD = utils.GetFloat64Decimal(remainingFundsInUSD, 8) - sessionSnapshot.CollateralInUSD
	debt.TotalValueInUSD = utils.GetFloat64Decimal(
		eng.GetAmountInUSD(cumIndexAndUToken.Token, debt.CalTotalValueBI.Convert(), session.Version), 8)
}

func (eng *DebtEngine) remainingFundsv3(session *schemas.CreditSession, debt *schemas.Debt, cumIndexAndUToken *ds.CumIndexAndUToken,
	amountToPool, calRemainingFunds *big.Int) (repayAmount, remainingFunds *big.Int) {
	// while close account on v2 we calculate remainingFunds from all the token transfer from the user
	if session.Status == schemas.Closed && session.ClosedAt == debt.BlockNumber+1 {
		prices := core.JsonFloatMap{}
		for token, transferAmt := range *session.CloseTransfers {
			tokenPrice := eng.GetTokenLastPrice(token, session.Version)
			price := utils.GetFloat64Decimal(tokenPrice, 8)
			prices[token] = price
			if transferAmt < 0 {
				// assuming there is only one transfer from borrower to account
				// this transfer will be in underlyingtoken. execute_parser.go:246 and
				// https://github.com/Gearbox-protocol/core-v2/blob/main/contracts/credit/CreditManager.sol#L359-L363
				amt := new(big.Float).Mul(big.NewFloat(transferAmt*-1), utils.GetExpFloat(eng.repo.GetToken(token).Decimals))
				repayAmount, _ = amt.Int(nil)
			}
		}
		// remainingFunds calculation
		// set price for underlying token
		prices[cumIndexAndUToken.Token] = utils.GetFloat64Decimal(
			eng.GetTokenLastPrice(cumIndexAndUToken.Token, session.Version), 8)
		remainingFunds = session.CloseTransfers.ValueInUnderlying(cumIndexAndUToken.Token, cumIndexAndUToken.Decimals, prices)
	} else if secStatus := session.StatusAt(debt.BlockNumber + 1); schemas.IsStatusLiquidated(secStatus) {
		//
		remainingFunds = calRemainingFunds
		repayAmount = new(big.Int) // TODO can be set by tracking all transfers on the liquidation and then checking the underlyingToken transfer to the account
	} else {
		// repayamount
		// for account not closed or liquidated yet
		// get underlying balance
		underlying := (*eng.lastCSS[session.ID].Balances)[cumIndexAndUToken.Token]
		underlyingBalance := new(big.Int)
		if underlying.BI != nil {
			underlyingBalance = underlying.BI.Convert()
		}
		if new(big.Int).Sub(underlyingBalance, new(big.Int).Add(amountToPool, calRemainingFunds)).Cmp(big.NewInt(1)) > 0 {
			repayAmount = new(big.Int)
		} else {
			repayAmount = new(big.Int).Sub(new(big.Int).Add(amountToPool, calRemainingFunds), underlyingBalance)
		}

		// remainingfunds
		remainingFunds = new(big.Int).Sub(debt.CalTotalValueBI.Convert(), amountToPool)
	}
	return
}

func (eng *DebtEngine) remainingFundsv2(session *schemas.CreditSession, debt *schemas.Debt, cumIndexAndUToken *ds.CumIndexAndUToken,
	amountToPool, calRemainingFunds *big.Int) (repayAmount, remainingFunds *big.Int) {
	repayAmount = new(big.Int)
	// while close account on v2 we calculate remainingFunds from all the token transfer from the user
	if session.Status == schemas.Closed && session.ClosedAt == debt.BlockNumber+1 {
		prices := core.JsonFloatMap{}
		for token, transferAmt := range *session.CloseTransfers {
			tokenPrice := eng.GetTokenLastPrice(token, session.Version)
			price := utils.GetFloat64Decimal(tokenPrice, 8)
			prices[token] = price
			if transferAmt < 0 {
				// assuming there is only one transfer from borrower to account
				// this transfer will be in underlyingtoken. execute_parser.go:246 and
				// https://github.com/Gearbox-protocol/core-v2/blob/main/contracts/credit/CreditManager.sol#L359-L363
				amt := new(big.Float).Mul(big.NewFloat(transferAmt*-1), utils.GetExpFloat(eng.repo.GetToken(token).Decimals))
				repayAmount, _ = amt.Int(nil)
			}
		}
		// remainingFunds calculation
		// set price for underlying token
		prices[cumIndexAndUToken.Token] = utils.GetFloat64Decimal(
			eng.GetTokenLastPrice(cumIndexAndUToken.Token, session.Version), 8)
		remainingFunds = session.CloseTransfers.ValueInUnderlying(cumIndexAndUToken.Token, cumIndexAndUToken.Decimals, prices)
	} else if session.ClosedAt == debt.BlockNumber+1 && schemas.IsStatusLiquidated(session.Status) {
		remainingFunds = session.RemainingFunds.Convert()
		repayAmount = new(big.Int)
	} else {
		// repayamount
		// for account not closed or liquidated yet
		// get underlying balance
		underlying := (*eng.lastCSS[session.ID].Balances)[cumIndexAndUToken.Token]
		underlyingBalance := new(big.Int)
		if underlying.BI != nil {
			underlyingBalance = underlying.BI.Convert()
		}
		if new(big.Int).Sub(underlyingBalance, new(big.Int).Add(amountToPool, calRemainingFunds)).Cmp(big.NewInt(1)) > 0 {
			repayAmount = new(big.Int)
		} else {
			repayAmount = new(big.Int).Sub(new(big.Int).Add(amountToPool, calRemainingFunds), underlyingBalance)
		}

		// remainingfunds
		remainingFunds = new(big.Int).Sub(debt.CalTotalValueBI.Convert(), amountToPool)
	}
	return
}

func (eng *DebtEngine) remainingFundsv1(session *schemas.CreditSession, debt *schemas.Debt, amountToPool, calRemainingFunds *big.Int) (repayAmount, remainingFunds *big.Int) {
	if session.ClosedAt == debt.BlockNumber+1 && (session.Status == schemas.Closed || session.Status == schemas.Liquidated) {
		remainingFunds = (*big.Int)(session.RemainingFunds)
	} else {
		remainingFunds = calRemainingFunds
	}
	if session.Status == schemas.Liquidated && session.ClosedAt == debt.BlockNumber+1 {
		repayAmount = new(big.Int).Add(amountToPool, remainingFunds)
	} else if session.Status == schemas.Closed && session.ClosedAt == debt.BlockNumber+1 {
		repayAmount = new(big.Int)
	} else {
		// https://github.com/Gearbox-protocol/gearbox-contracts/blob/master/contracts/credit/CreditManager.sol#L487-L490
		repayAmount = amountToPool
	}
	return
}
