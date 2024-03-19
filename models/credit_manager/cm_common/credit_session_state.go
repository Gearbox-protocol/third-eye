package cm_common

import (
	"math/big"
	"strconv"
	"strings"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/calc"
	"github.com/Gearbox-protocol/sdk-go/pkg/dc"
	"github.com/ethereum/go-ethereum/common"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

// for update actions fetch at blockNum
// for closure and liquidation (liq except v3) fetch at blockNum-1
// for liq v3 fetch at blockNum -1 and blockNum
func (mdl *CommonCMAdapter) FetchFromDCForChangedSessions(blockNum int64) (calls []multicall.Multicall2Call, processFns []func(multicall.Multicall2Result)) {
	for sessionId := range mdl.updatedSessions { // for liq v3 updateSessions is also updated
		if mdl.closedSessions[sessionId] == nil {
			call, processFn := mdl.updateSessionCallAndProcessFn(sessionId, blockNum)
			if processFn != nil {
				calls = append(calls, call)
				processFns = append(processFns, processFn)
			}
		}
	}
	{ // these calls are made internally, wrapper doesn't handle then since they are on blockNum -1
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
		{ // liq v3
			for sessionId, liqv3Details := range mdl.liqv3Sessions {
				call, processFn := mdl.liqv3SessionCallAndResultFn(blockNum, sessionId, liqv3Details)
				if processFn != nil {
					calls = append(calls, call)
					processFns = append(processFns, processFn)
				}
			}
		}
		results := core.MakeMultiCall(mdl.Client, blockNum-1, false, calls)
		for i, result := range results {
			processFns[i](result)
		}
	}
	//
	mdl.updatedSessions = make(map[string]int)
	mdl.liqv3Sessions = make(map[string]*SessionLiqUpdatev3Details)
	mdl.closedSessions = make(map[string]*SessionCloseDetails)
	return
}

func (mdl *CommonCMAdapter) liqv3SessionCallAndResultFn(liquidatedAt int64, sessionId string, liqv3Details *SessionLiqUpdatev3Details) (call multicall.Multicall2Call, processFn func(multicall.Multicall2Result)) {
	session := mdl.Repo.UpdateCreditSession(sessionId, nil)
	// set secondary status
	if session.TeritaryStatus == nil || *session.TeritaryStatus == nil {
		session.TeritaryStatus = &core.Json{"secStatus": [][]int64{}}
	}
	secStatus := utils.ListOfInt64List((*session.TeritaryStatus)["secStatus"])
	secStatus = append(secStatus, []int64{liquidatedAt, int64(liqv3Details.Status)})
	(*session.TeritaryStatus)["secStatus"] = secStatus
	//
	session.IsDirty = true
	// get call and processFn
	call, resultFn, err := mdl.Repo.GetDCWrapper().GetCreditAccountData(mdl.GetVersion(), liquidatedAt-1,
		common.HexToAddress(mdl.GetAddress()),
		common.HexToAddress(session.Borrower),
		common.HexToAddress(session.Account),
	)
	if err != nil {
		log.Fatalf("Failing preparing GetAccount for CM:%s Borrower:%s: %v", mdl.GetAddress(), session.Borrower, err)
	}
	return call, func(result multicall.Multicall2Result) {
		if !result.Success {
			log.Fatalf("Failing GetAccount for CM:%s Borrower:%s at %d: %v", mdl.GetAddress(), session.Borrower, liquidatedAt-1, result.ReturnData)
		}
		dcAccountData, err := resultFn(result.ReturnData)
		if err != nil {
			log.Fatalf("For blockNum %d CM:%s Borrower:%s %v", liquidatedAt, mdl.GetAddress(), session.Borrower, err)
		}
		mdl.liqv3Session(liquidatedAt, session, dcAccountData, liqv3Details)
	}
}

// data is for liquidatedAt -1
func (mdl *CommonCMAdapter) liqv3Session(liquidatedAt int64, session *schemas.CreditSession, data dc.CreditAccountCallData, liqv3Details *SessionLiqUpdatev3Details) {
	session.BorrowedAmount = (*core.BigInt)(data.BorrowedAmount)
	session.IsDirty = true

	mdl.createCSSnapshot(liquidatedAt-1, session, data)
}

func (mdl *CommonCMAdapter) closeSessionCallAndResultFn(closedAt int64, sessionId string, closeDetails *SessionCloseDetails) (call multicall.Multicall2Call, processFn func(multicall.Multicall2Result)) {
	mdl.State.OpenedAccountsCount--
	// check the data before credit session was closed by minus 1.
	session := mdl.Repo.UpdateCreditSession(sessionId, nil)
	// set session fields
	session.ClosedAt = closedAt
	session.Status = closeDetails.Status
	session.IsDirty = true
	// this checks prevent getting data for credit session that exist only within a block
	// datacompressor query will fail
	if session.Since == session.ClosedAt {
		return
	}
	// get call and processFn
	call, resultFn, err := mdl.Repo.GetDCWrapper().GetCreditAccountData(mdl.GetVersion(), closedAt-1,
		common.HexToAddress(mdl.GetAddress()),
		common.HexToAddress(session.Borrower),
		common.HexToAddress(session.Account),
	)
	if err != nil {
		log.Fatalf("Failing preparing GetAccount for CM:%s Borrower:%s: %v", mdl.GetAddress(), session.Borrower, err)
	}
	return call, func(result multicall.Multicall2Result) {
		if !result.Success {
			key, dc := mdl.Repo.GetDCWrapper().GetKeyAndAddress(session.Version, closedAt-1)
			log.Fatalf("Failing GetAccount for CM:%s Borrower:%s at %d: %v, dc: %s(%s)", mdl.GetAddress(), session.Borrower, closedAt-1, result.ReturnData, key, dc)
		}
		dcAccountData, err := resultFn(result.ReturnData)
		if err != nil {
			log.Fatalf("For blockNum %d CM:%s Borrower:%s %v", closedAt, mdl.GetAddress(), session.Borrower, err)
		}
		mdl.closeSession(closedAt, session, dcAccountData, closeDetails)
		// // set close price for v3
		// if session.Version.MoreThanEq(core.NewVersion(300)) {
		// 	mdl.Repo.GetCreditSession(sessionId).ClosePrice = ds.GetOneInchUpdater().GetCurrentPriceAtBlockNum(closedAt-1, *session.Balances, mdl.GetUnderlyingToken())
		// }
	}
}

// used for v1/v2 close and liquidate
// used for v3 close
func (mdl *CommonCMAdapter) closeSession(closedAt int64, session *schemas.CreditSession, data dc.CreditAccountCallData, closeDetails *SessionCloseDetails) {
	session.BorrowedAmount = (*core.BigInt)(data.BorrowedAmount)
	session.IsDirty = true

	if !session.Version.MoreThanEq(core.NewVersion(300)) {
		mdl.poolRepay(closedAt, session, closeDetails, data)
	}

	// v1 repayment
	if closeDetails.RemainingFunds == nil && closeDetails.Status == schemas.Repaid {
		closeDetails.RemainingFunds = new(big.Int).Sub(data.TotalValue.Convert(), data.RepayAmountv1v2.Convert())
		session.RemainingFunds = (*core.BigInt)(closeDetails.RemainingFunds)
		(*closeDetails.AccountOperation.Args)["repayAmount"] = data.RepayAmountv1v2
		mdl.AddAccountOperation(closeDetails.AccountOperation)
	}

	mdl.createCSSnapshot(closedAt-1, session, data)
}

func (mdl *CommonCMAdapter) setCSSCollateralFields(blockNum int64, session *schemas.CreditSession, css *schemas.CreditSessionSnapshot) {
	css.CollateralInUSD = session.CollateralInUSD
	css.CollateralInUnderlying = session.CollateralInUnderlying
	css.Collateral = session.Collateral

	collateral := new(big.Int)
	// log.Info(session.ID, session.Collateral)
	for token, amount := range *session.Collateral {
		valueInUnderlyingAsset := mdl.Repo.GetValueInCurrency(blockNum, session.Version, token, mdl.GetUnderlyingToken(), amount.Convert())
		collateral = new(big.Int).Add(collateral, valueInUnderlyingAsset)
	}
	css.InstCollteralUnderlying = utils.GetFloat64Decimal(collateral, mdl.GetUnderlyingDecimal())
	//
	valueInUSD := mdl.Repo.GetValueInCurrency(blockNum, session.Version, mdl.GetUnderlyingToken(), "USDC", collateral)
	css.InstCollteralUSD = utils.GetFloat64Decimal(valueInUSD, 6)
}

// for liquidatev3 it is called at blockNum -1 due to `liqv3Sessions` and then at blockNum due to `updatedSessions`
func (mdl *CommonCMAdapter) createCSSnapshot(blockNum int64, session *schemas.CreditSession, data dc.CreditAccountCallData) {
	// create session snapshot
	css := &schemas.CreditSessionSnapshot{}
	mdl.Repo.SetBlock(blockNum)
	css.BlockNum = blockNum
	css.SessionId = session.ID
	mdl.setCSSCollateralFields(blockNum, session, css)
	css.Borrower = session.Borrower
	css.HealthFactor = (*core.BigInt)(data.HealthFactor)
	css.TotalValueBI = (*core.BigInt)(data.TotalValue)
	css.TotalValue = utils.GetFloat64Decimal(data.TotalValue, mdl.GetUnderlyingDecimal())
	//
	css.CumulativeQuotaInterest = data.CumulativeQuotaInterest
	// quota fees
	css.QuotaFees = (*core.BigInt)(data.GetQuotaFees(mdl.params.FeeInterest))
	//
	// set balances
	css.Balances = mdl.addFloatValue(session.Account, blockNum, data.Balances)
	css.ExtraQuotaAPY = schemas.QuotaBorrowRate(*css.Balances, css.TotalValueBI)
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
	mdl.Repo.AddCreditSessionSnapshot(css)
}
func (mdl *CommonCMAdapter) poolRepay(blockNum int64, session *schemas.CreditSession, details *SessionCloseDetails, data dc.CreditAccountCallData) {
	// log.Info(mdl.params, session.Version,
	// "totalvalue", data.TotalValue, closeDetails.Status,
	// "borrow", data.BorrowedAmountPlusInterest, data.BorrowedAmount)
	amountToPool, _, _, _ := calc.CalCloseAmount(mdl.params,
		session.Version, data.TotalValue.Convert(),
		details.Status,
		calc.NewDebtDetails(
			data.Debt.Convert(),
			data.AccruedInterest.Convert(),
			data.BorrowedAmount.Convert(),
		),
	)
	// pool repay
	// check for avoiding db errors
	mdl.PoolRepay(blockNum,
		details.LogId,
		details.TxHash,
		session.ID,
		details.Borrower,
		amountToPool)

}

func (mdl *CommonCMAdapter) updateSessionCallAndProcessFn(sessionId string, blockNum int64) (
	multicall.Multicall2Call, func(multicall.Multicall2Result)) {
	if mdl.DontGetSessionFromDCForTest {
		return multicall.Multicall2Call{}, nil
	}
	session := mdl.Repo.UpdateCreditSession(sessionId, nil)
	call, resultFn, err := mdl.Repo.GetDCWrapper().GetCreditAccountData(mdl.GetVersion(), blockNum,
		common.HexToAddress(mdl.GetAddress()),
		common.HexToAddress(session.Borrower),
		common.HexToAddress(session.Account),
	)
	if err != nil {
		log.Fatalf("Failing preparing GetAccount for CM:%s Borrower:%s: %v", mdl.GetAddress(), session.Borrower, err)
	}
	return call, func(result multicall.Multicall2Result) {
		if !result.Success {
			log.Fatalf("Failing GetAccount for CM:%s Borrower:%s at %d: %v. account: %s", mdl.GetAddress(), session.Borrower, blockNum, result.ReturnData, session.Account)
		}
		dcAccountData, err := resultFn(result.ReturnData)
		if err != nil {
			log.Fatalf("For blockNum %d CM:%s Borrower:%s %v", blockNum, mdl.GetAddress(), session.Borrower, err)
		}
		mdl.updateSession(blockNum, session, dcAccountData)
		// set entry price for v3
		// if session.Since == blockNum && session.Version.MoreThanEq(core.NewVersion(300)) {
		// 	mdl.Repo.GetCreditSession(sessionId).EntryPrice = ds.GetOneInchUpdater().GetCurrentPriceAtBlockNum(blockNum, *session.Balances, mdl.GetUnderlyingToken())
		// }
	}
}

func sessionIdToSince(sessionId string) int64 {
	_since := strings.Split(sessionId, "_")[1]
	since, err := strconv.ParseInt(_since, 10, 64)
	log.CheckFatal(err)
	return since
}

func (mdl *CommonCMAdapter) updateSession(blockNum int64, session *schemas.CreditSession, data dc.CreditAccountCallData) {
	session.BorrowedAmount = (*core.BigInt)(data.BorrowedAmount)
	session.IsDirty = true
	mdl.createCSSnapshot(blockNum, session, data)
}

func (mdl *CommonCMAdapter) addFloatValue(account string, blockNum int64, dcv2Balances []core.TokenBalanceCallData) *core.DBBalanceFormat {
	return AddStETHBalance(account, blockNum, dcv2Balances, mdl.Client, mdl, mdl.Repo.GetTokenFromSdk("stETH"))
}

func (mdl *CommonCMAdapter) GetDecimals(token common.Address) int8 {
	return mdl.Repo.GetToken(token.Hex()).Decimals
}

type DecimalStoreI interface {
	GetDecimals(tokenAddr common.Address) int8
}

func AddStETHBalance(account string, blockNum int64, dcv2Balances []core.TokenBalanceCallData, client core.ClientI, tStore DecimalStoreI, stETH string) *core.DBBalanceFormat {
	dbFormat := core.DBBalanceFormat{}
	for _, balance := range dcv2Balances {
		token := balance.Token
		if balance.HasBalanceMoreThanOne() { // is enabled not needed.
			balance.F = utils.GetFloat64Decimal(balance.BI, tStore.GetDecimals(common.HexToAddress(token)))
			dbFormat[token] = balance.DBTokenBalance
			//
			if stETH == token {
				accountData := common.HexToHash(account)
				_v, err := core.CallFuncWithExtraBytes(
					client, "f5eb42dc", // shareOf, https://etherscan.io/token/0xae7ab96520de3a18e5e111b5eaab095312d7fe84#readProxyContract
					common.HexToAddress(token), blockNum, accountData[:],
				)
				log.CheckFatal(err)
				amt := new(big.Int).SetBytes(_v)
				//
				dbFormat[core.NULL_ADDR.Hex()] = core.DBTokenBalance{
					IsForbidden: false,
					IsEnabled:   false,
					BI:          (*core.BigInt)(amt),
					F:           utils.GetFloat64Decimal(amt, 18),
					Ind:         -1,
				}
			}
		}
	}
	return &dbFormat
}
