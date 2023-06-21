package cm_common

import (
	"math/big"

	dcv2 "github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressor/dataCompressorv2"
	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/calc"
	"github.com/ethereum/go-ethereum/common"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

func (mdl *CMCommon) FetchFromDCForChangedSessions(blockNum int64) (calls []multicall.Multicall2Call, processFns []func(multicall.Multicall2Result)) {
	for sessionId := range mdl.updatedSessions {
		if mdl.closedSessions[sessionId] == nil {
			call, processFn := mdl.updateSessionCallAndProcessFn(sessionId, blockNum)
			if processFn != nil {
				calls = append(calls, call)
				processFns = append(processFns, processFn)
			}
		}
	}
	{
		calls := make([]multicall.Multicall2Call, 0)
		processFns := make([]func(multicall.Multicall2Result), 0)
		for sessionId, closeDetails := range mdl.closedSessions {
			updates := mdl.updatedSessions[sessionId]
			if updates != 0 {
				// in this case, affected fields are data that we fetch from datacompressor:
				//
				// borrowAmount, for amountToPool calc to add PoolRepay[affected]
				//
				// totalValue, [not affected]
				// BorrowedAmountPlusInterest, [not affected]
				// CumulativeIndexAtOpen, [not affected]
				// HealthFactor, [not affected]
				// balances, [not affected]
				//
				// remainingFunds, collateral, won't be affected, dependent
				// on (close, liquidate adding remainingFunds and collateral)
				log.Infof("Warn Session: %s updated %d before close %+v in same block %d\n", sessionId, updates, closeDetails, blockNum)
			}
			call, processFn := mdl.closeSessionCallAndResultFn(blockNum, sessionId, closeDetails)
			if processFn != nil {
				calls = append(calls, call)
				processFns = append(processFns, processFn)
			}
		}
		results := core.MakeMultiCall(mdl.Client, blockNum-1, false, calls)
		for i, result := range results {
			processFns[i](result)
		}
	}
	mdl.updatedSessions = make(map[string]int)
	mdl.closedSessions = make(map[string]*SessionCloseDetails)
	return
}

func (mdl *CMCommon) closeSessionCallAndResultFn(blockNum int64, sessionId string, closeDetails *SessionCloseDetails) (call multicall.Multicall2Call, processFn func(multicall.Multicall2Result)) {
	mdl.State.OpenedAccountsCount--
	// check the data before credit session was closed by minus 1.
	session := mdl.Repo.UpdateCreditSession(sessionId, nil)
	// set session fields
	session.ClosedAt = blockNum
	session.Status = closeDetails.Status
	// this checks prevent getting data for credit session that exist only within a block
	// datacompressor query will fail
	if session.Since == session.ClosedAt {
		return
	}
	// get call and processFn
	call, resultFn, err := mdl.Repo.GetDCWrapper().GetCreditAccountData(blockNum-1,
		common.HexToAddress(mdl.GetAddress()),
		common.HexToAddress(session.Borrower))
	if err != nil {
		log.Fatalf("Failing preparing GetAccount for CM:%s Borrower:%s: %v", mdl.GetAddress(), session.Borrower, err)
	}
	return call, func(result multicall.Multicall2Result) {
		if !result.Success {
			log.Fatalf("Failing GetAccount for CM:%s Borrower:%s: %v", mdl.GetAddress(), session.Borrower, result.ReturnData)
		}
		dcAccountData, err := resultFn(result.ReturnData)
		if err != nil {
			log.Fatalf("For blockNum %d CM:%s Borrower:%s %v", blockNum, mdl.GetAddress(), session.Borrower, err)
		}
		mdl.closeSession(blockNum, session, dcAccountData, closeDetails)
	}
}

func (mdl *CMCommon) closeSession(blockNum int64, session *schemas.CreditSession, data dcv2.CreditAccountData, closeDetails *SessionCloseDetails) {
	session.BorrowedAmount = (*core.BigInt)(data.BorrowedAmount)
	// log.Info(mdl.params, session.Version,
	// "totalvalue", data.TotalValue, closeDetails.Status,
	// "borrow", data.BorrowedAmountPlusInterest, data.BorrowedAmount)
	amountToPool, _, _, _ := calc.CalCloseAmount(mdl.params,
		session.Version, data.TotalValue,
		closeDetails.Status,
		data.BorrowedAmountPlusInterest,
		data.BorrowedAmount,
	)
	// switch closeDetails.Status {
	// case schemas.Closed,
	// 	schemas.Repaid:
	// 	amountToPool = data.RepayAmount
	// case schemas.Liquidated:
	// 	amountToPool = data.LiquidationAmount
	// }
	// pool repay
	// check for avoiding db errors
	mdl.PoolRepay(blockNum,
		closeDetails.LogId,
		closeDetails.TxHash,
		session.ID,
		closeDetails.Borrower,
		amountToPool)

	if closeDetails.RemainingFunds == nil && closeDetails.Status == schemas.Repaid {
		closeDetails.RemainingFunds = new(big.Int).Sub(data.TotalValue, data.RepayAmount)
		session.RemainingFunds = (*core.BigInt)(closeDetails.RemainingFunds)
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
	css.SessionId = session.ID
	css.CollateralInUSD = session.CollateralInUSD
	css.CollateralInUnderlying = session.CollateralInUnderlying
	css.Borrower = session.Borrower
	css.HealthFactor = (*core.BigInt)(data.HealthFactor)
	css.TotalValueBI = (*core.BigInt)(data.TotalValue)
	css.TotalValue = utils.GetFloat64Decimal(data.TotalValue, mdl.GetUnderlyingDecimal())
	// set balances
	css.Balances = mdl.addFloatValue(session.Account, blockNum-1, data.Balances)
	// for close credit account operation on gearbox v2
	// https://github.com/Gearbox-protocol/contracts-v2/blob/main/contracts/credit/CreditFacade.sol#L235
	// there is a skipTokenMask which can be used to skip certain tokens from getting transferred to borrower
	// this can decrease the gas used by credit manager and saving money for borrower
	// as a result, balances fetched from datacompressor on closeBlock-1 will not be valid for remainingFunds
	// calculation.
	session.Balances = css.Balances
	//
	css.BorrowedAmountBI = core.NewBigInt(session.BorrowedAmount)
	css.BorrowedAmount = utils.GetFloat64Decimal(data.BorrowedAmount, mdl.GetUnderlyingDecimal())
	css.СumulativeIndexAtOpen = core.NewBigInt((*core.BigInt)(data.CumulativeIndexAtOpen))
	mdl.Repo.AddCreditSessionSnapshot(&css)
}

func (mdl *CMCommon) updateSessionCallAndProcessFn(sessionId string, blockNum int64) (
	multicall.Multicall2Call, func(multicall.Multicall2Result)) {
	if mdl.DontGetSessionFromDCForTest {
		return multicall.Multicall2Call{}, nil
	}
	session := mdl.Repo.UpdateCreditSession(sessionId, nil)
	call, resultFn, err := mdl.Repo.GetDCWrapper().GetCreditAccountData(blockNum,
		common.HexToAddress(mdl.GetAddress()),
		common.HexToAddress(session.Borrower))
	if err != nil {
		log.Fatalf("Failing preparing GetAccount for CM:%s Borrower:%s: %v", mdl.GetAddress(), session.Borrower, err)
	}
	return call, func(result multicall.Multicall2Result) {
		if !result.Success {
			log.Fatalf("Failing GetAccount for CM:%s Borrower:%s: %v", mdl.GetAddress(), session.Borrower, result.ReturnData)
		}
		dcAccountData, err := resultFn(result.ReturnData)
		if err != nil {
			log.Fatalf("For blockNum %d CM:%s Borrower:%s %v", blockNum, mdl.GetAddress(), session.Borrower, err)
		}
		mdl.updateSession(blockNum, session, dcAccountData)
	}
}

func (mdl *CMCommon) updateSession(blockNum int64, session *schemas.CreditSession, data dcv2.CreditAccountData) {
	session.BorrowedAmount = (*core.BigInt)(data.BorrowedAmount)

	// create session snapshot
	css := schemas.CreditSessionSnapshot{}
	css.BlockNum = blockNum
	css.SessionId = session.ID
	css.CollateralInUSD = session.CollateralInUSD
	css.CollateralInUnderlying = session.CollateralInUnderlying
	css.Borrower = session.Borrower
	css.HealthFactor = (*core.BigInt)(data.HealthFactor)
	css.TotalValueBI = (*core.BigInt)(data.TotalValue)
	css.TotalValue = utils.GetFloat64Decimal(data.TotalValue, mdl.GetUnderlyingDecimal())
	// set balances of css and credit session
	css.Balances = mdl.addFloatValue(session.Account, blockNum, data.Balances)
	session.Balances = css.Balances
	//
	css.BorrowedAmountBI = core.NewBigInt(session.BorrowedAmount)
	css.BorrowedAmount = utils.GetFloat64Decimal(data.BorrowedAmount, mdl.GetUnderlyingDecimal())
	css.СumulativeIndexAtOpen = core.NewBigInt((*core.BigInt)(data.CumulativeIndexAtOpen))
	mdl.Repo.AddCreditSessionSnapshot(&css)
}

func (mdl *CMCommon) addFloatValue(account string, blockNum int64, dcv2Balances []dcv2.TokenBalance) *core.DBBalanceFormat {
	dbFormat := core.DBBalanceFormat{}
	for ind, balance := range dcv2Balances {
		token := balance.Token.Hex()
		if balance.IsEnabled && balance.Balance.Cmp(new(big.Int)) > 0 {
			dbFormat[token] = core.CoreIntBalance{
				IsAllowed: balance.IsAllowed,
				IsEnabled: balance.IsEnabled,
				BI:        (*core.BigInt)(balance.Balance),
				F:         utils.GetFloat64Decimal(balance.Balance, mdl.Repo.GetToken(token).Decimals),
				Ind:       ind,
			}
			if mdl.Repo.GetTokenFromSdk("stETH") == token {
				accountData := common.HexToHash(account)
				_v, err := core.CallFuncWithExtraBytes(
					mdl.Client, "f5eb42dc", // shareOf, https://etherscan.io/token/0xae7ab96520de3a18e5e111b5eaab095312d7fe84#readProxyContract
					common.HexToAddress(token), blockNum, accountData[:],
				)
				log.CheckFatal(err)
				amt := new(big.Int).SetBytes(_v)
				//
				dbFormat[core.NULL_ADDR.Hex()] = core.CoreIntBalance{
					IsAllowed: false,
					IsEnabled: false,
					BI:        (*core.BigInt)(amt),
					F:         utils.GetFloat64Decimal(amt, 18),
					Ind:       -1,
				}
			}
		}
	}
	return &dbFormat
}
