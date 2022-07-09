package credit_manager

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

func (mdl *CreditManager) FetchFromDCForChangedSessions(blockNum int64) {
	for sessionId := range mdl.UpdatedSessions {
		if mdl.ClosedSessions[sessionId] == nil {
			mdl.updateSession(sessionId, blockNum)
		}
	}
	for sessionId, closeDetails := range mdl.ClosedSessions {
		updates := mdl.UpdatedSessions[sessionId]
		if updates != 0 {
			log.Fatalf("Session: %s updated %d before close %+v in block %d\n", sessionId, updates, closeDetails, blockNum)
		}
		mdl.closeSession(sessionId, blockNum, closeDetails)
	}
	mdl.UpdatedSessions = make(map[string]int)
	mdl.ClosedSessions = make(map[string]*SessionCloseDetails)
}

func (mdl *CreditManager) closeSession(sessionId string, blockNum int64, closeDetails *SessionCloseDetails) {
	mdl.State.OpenedAccountsCount--
	// check the data before credit session was closed by minus 1.
	session := mdl.Repo.UpdateCreditSession(sessionId, map[string]interface{}{})
	// set session fields
	session.ClosedAt = blockNum
	session.Status = closeDetails.Status
	// this checks prevent getting data for credit session that exist only within a block
	// datacompressor query will fail
	if session.Since == session.ClosedAt {
		return
	}
	data := mdl.GetCreditSessionData(blockNum-1, session.Borrower)
	session.BorrowedAmount = (*core.BigInt)(data.BorrowedAmount)
	var amountToPool *big.Int
	switch closeDetails.Status {
	case schemas.Closed,
		schemas.Repaid:
		amountToPool = data.RepayAmount
	case schemas.Liquidated:
		amountToPool = data.LiquidationAmount
	}
	// pool repay
	mdl.PoolRepay(blockNum,
		closeDetails.LogId,
		closeDetails.TxHash,
		sessionId,
		closeDetails.Borrower,
		amountToPool)

	if closeDetails.RemainingFunds == nil && closeDetails.Status == schemas.Repaid {
		closeDetails.RemainingFunds = new(big.Int).Sub(data.TotalValue, data.RepayAmount)
		(*closeDetails.AccountOperation.Args)["repayAmount"] = data.RepayAmount
		mdl.AddAccountOperation(closeDetails.AccountOperation)
	}

	// credit manager state
	mdl.State.TotalRepaidBI = core.AddCoreAndInt(mdl.State.TotalRepaidBI, amountToPool)
	mdl.State.TotalRepaid = utils.GetFloat64Decimal(mdl.State.TotalRepaidBI.Convert(), mdl.GetUnderlyingDecimal())
	//
	// create session snapshot
	css := schemas.CreditSessionSnapshot{}
	mdl.Repo.SetBlock(blockNum - 1)
	css.BlockNum = blockNum - 1
	css.SessionId = sessionId
	css.CollateralInUSD = session.CollateralInUSD
	css.CollateralInUnderlying = session.CollateralInUnderlying
	css.Borrower = session.Borrower
	css.HealthFactor = (*core.BigInt)(data.HealthFactor)
	css.TotalValueBI = (*core.BigInt)(data.TotalValue)
	css.TotalValue = utils.GetFloat64Decimal(data.TotalValue, mdl.GetUnderlyingDecimal())
	mask := mdl.Repo.GetMask(blockNum-1, mdl.GetAddress(), session.Account, session.Version)
	// set balances
	var err error
	css.Balances, err = mdl.Repo.ConvertToBalanceWithMask(data.Balances, mask)
	if err != nil {
		log.Fatalf("DC wrong token values block:%d dc:%s", blockNum, mdl.Repo.GetDCWrapper().ToJson())
	}
	// for close credit account operation on gearbox v2
	// https://github.com/Gearbox-protocol/contracts-v2/blob/main/contracts/credit/CreditFacade.sol#L235
	// there is a skipTokenMask which can be used to skip certain tokens from getting transferred to borrower
	// this can decrease the gas used by credit manager and saving money for borrower
	// as a result, balances fetched from datacompressor on closeBlock-1 will not be valid for remainingFunds
	// calculation.
	if closeDetails.Status != schemas.Closed || session.Version != 2 { // neg( closed on v2)
		session.Balances = css.Balances
	}
	//
	css.BorrowedAmountBI = core.NewBigInt(session.BorrowedAmount)
	css.BorrowedAmount = utils.GetFloat64Decimal(data.BorrowedAmount, mdl.GetUnderlyingDecimal())
	css.СumulativeIndexAtOpen = core.NewBigInt((*core.BigInt)(data.CumulativeIndexAtOpen))
	mdl.Repo.AddCreditSessionSnapshot(&css)
}

func (mdl *CreditManager) updateSession(sessionId string, blockNum int64) {
	session := mdl.Repo.UpdateCreditSession(sessionId, map[string]interface{}{})
	data := mdl.GetCreditSessionData(blockNum, session.Borrower)
	session.BorrowedAmount = (*core.BigInt)(data.BorrowedAmount)

	// create session snapshot
	css := schemas.CreditSessionSnapshot{}
	css.BlockNum = blockNum
	css.SessionId = sessionId
	css.CollateralInUSD = session.CollateralInUSD
	css.CollateralInUnderlying = session.CollateralInUnderlying
	css.Borrower = session.Borrower
	css.HealthFactor = (*core.BigInt)(data.HealthFactor)
	css.TotalValueBI = (*core.BigInt)(data.TotalValue)
	css.TotalValue = utils.GetFloat64Decimal(data.TotalValue, mdl.GetUnderlyingDecimal())
	mask := mdl.Repo.GetMask(blockNum, mdl.GetAddress(), session.Account, session.Version)
	// set balances of css and credit session
	var err error
	css.Balances, err = mdl.Repo.ConvertToBalanceWithMask(data.Balances, mask)
	if err != nil {
		log.Fatalf("DC wrong token values block:%d dc:%s", blockNum, mdl.Repo.GetDCWrapper().ToJson())
	}
	session.Balances = css.Balances
	//
	css.BorrowedAmountBI = core.NewBigInt(session.BorrowedAmount)
	css.BorrowedAmount = utils.GetFloat64Decimal(data.BorrowedAmount, mdl.GetUnderlyingDecimal())
	css.СumulativeIndexAtOpen = core.NewBigInt((*core.BigInt)(data.CumulativeIndexAtOpen))
	mdl.Repo.AddCreditSessionSnapshot(&css)
}
