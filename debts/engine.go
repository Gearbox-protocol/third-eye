package debts

import (
	"math/big"
	"sort"

	"github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressor/mainnet"
	"github.com/Gearbox-protocol/sdk-go/artifacts/yearnPriceFeed"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (eng *DebtEngine) SaveProfile(profile string) {
	err := eng.db.Create(&schemas.ProfileTable{Profile: profile}).Error
	if err != nil {
		log.Fatal(err)
	}
}

func (eng *DebtEngine) updateLocalState(blockNum int64, block *schemas.Block) (poolsUpdated []string, tokensUpdated map[string]bool, sessionsUpdated map[string]bool) {
	//////////////////////////////////////
	// Load data
	//////////////////////////////////////
	// L1 params
	// L2 poolinterestdata
	// L3 credit session snapshots
	// L4 lt
	// L5 pricefeeds
	//
	// L1:update params
	for _, params := range block.GetParams() {
		eng.addLastParameters(params)
	}

	///////////////////////////////////
	// calc debt conditions
	///////////////////////////////////
	// C1 If pool cumIndex changes
	// C2 If new css is present, meaning balance or new token added
	// C3 If account related token(either have balance or underlying token) has some data change, lt or price

	// update pool borrow rate and cumulative index
	// C1: UPDATE BASED ON POOL CUM INDEX
	for _, ps := range block.GetPoolStats() {
		poolsUpdated = append(poolsUpdated, ps.Address)
		// L2
		eng.AddPoolLastInterestData(&schemas.PoolInterestData{
			Address:            ps.Address,
			BlockNum:           ps.BlockNum,
			CumulativeIndexRAY: ps.CumulativeIndexRAY,
			BorrowAPYBI:        ps.BorrowAPYBI,
			Timestamp:          block.Timestamp,
		})
	}

	// update balance of last credit session snapshot and create credit session snapshots
	// C2: UPDATE BASED ON CSS
	sessionsUpdated = make(map[string]bool)
	for _, css := range block.GetCSS() {
		// L3
		eng.AddLastCSS(css)
		sessionsUpdated[css.SessionId] = true
	}

	// C3: UPDATE BASED ON TOKEN
	// C3.a: updated threshold
	tokensUpdated = make(map[string]bool)
	for _, allowedToken := range block.GetAllowedTokens() {
		// L4
		eng.AddAllowedTokenThreshold(allowedToken)
		tokensUpdated[allowedToken.Token] = true
	}

	// C3.b: updated price
	for _, pf := range block.GetPriceFeeds() {
		// if pf.Token == "0x6B175474E89094C44Da98b954EedeAC495271d0F" {
		// 	log.Info(blockNum, utils.ToJson(pf))
		// }
		// L5
		eng.AddTokenLastPrice(pf)
		tokensUpdated[pf.Token] = true
	}
	// updates complete
	return poolsUpdated, tokensUpdated, sessionsUpdated
}

func (eng *DebtEngine) CalculateDebt() {
	blocks := eng.repo.GetBlocks()
	sessions := eng.repo.GetSessions()
	// sort block numbers
	noOfBlock := len(blocks)
	blockNums := make([]int64, 0, noOfBlock)
	for blockNum := range blocks {
		blockNums = append(blockNums, blockNum)
	}
	sort.Slice(blockNums, func(i, j int) bool { return blockNums[i] < blockNums[j] })
	//
	for _, blockNum := range blockNums {
		block := blocks[blockNum]
		poolsUpdated, tokensUpdated, sessionsUpdated := eng.updateLocalState(blockNum, block)
		// get pool cumulative interest rate
		cmToPoolDetails := eng.GetCumulativeIndexAndDecimalForCMs(blockNum, block.Timestamp)

		// check if session's debt needs to be recalculated
		for _, session := range sessions {
			if (session.ClosedAt != 0 && session.ClosedAt <= blockNum) || session.Since > blockNum {
				continue
			}
			// #C1
			if utils.Contains(poolsUpdated, cmToPoolDetails[session.CreditManager].PoolAddr) {
				sessionsUpdated[session.ID] = true
			}
			// #C3
			sessionSnapshot := eng.lastCSS[session.ID]
			for tokenAddr := range *sessionSnapshot.Balances {
				if tokensUpdated[tokenAddr] {
					sessionsUpdated[session.ID] = true
				}
			}
			// #C3
			underlyingToken := cmToPoolDetails[session.CreditManager].Token
			if tokensUpdated[underlyingToken] {
				sessionsUpdated[session.ID] = true
			}
		}
		//
		// calculate each session debt
		for sessionId := range sessionsUpdated {
			session := sessions[sessionId]
			if (session.ClosedAt != 0 && session.ClosedAt <= blockNum) || session.Since > blockNum {
				continue
			}
			cmAddr := session.CreditManager
			// pool cum index is when the pool is not registered
			if cmToPoolDetails[cmAddr] != nil {
				eng.SessionDebtHandler(blockNum, session, cmToPoolDetails[cmAddr])
				// send notification when account is liquidated
			} else {
				log.Fatalf("CM(%s):pool is missing stats at %d, so cumulative index of pool is unknown", cmAddr, blockNum)
			}
		}
		// if len(sessionsUpdated) > 0 {
		// log.Verbosef("Calculated %d debts for block %d", len(sessionsUpdated), blockNum)
		// }
	}
}

func (eng *DebtEngine) ifAccountLiquidated(sessionId, cmAddr string, closedAt int64, status int) {
	sessionSnapshot := eng.lastCSS[sessionId]
	if status == schemas.Liquidated {
		account := eng.liquidableBlockTracker[sessionId]
		var liquidableSinceBlockNum int64
		if account != nil {
			liquidableSinceBlockNum = account.BlockNum
		} else {
			log.Warnf("Session(%s) liquidated at block:%d, but liquidable since block not stored", sessionId, closedAt)
		}
		urls := core.NetworkUIUrl(eng.config.ChainId)
		eng.repo.RecentEventMsg(closedAt-1, `Liquidation Alert:
		CreditManager: %s/address/%s
		Tx: %s/tx/%s
		Borrower: %s
		LiquidatedAt: %d 
		Liquidable since: %d
		web: %s/accounts/%s/%s`,
			urls.ExplorerUrl, cmAddr,
			urls.ExplorerUrl, eng.GetLiquidationTx(sessionId),
			sessionSnapshot.Borrower,
			closedAt, liquidableSinceBlockNum,
			urls.ChartUrl, cmAddr, sessionSnapshot.Borrower)
	} else if status != schemas.Active {
		// comment deletion form map
		// this way if the account is liquidable and liquidated in the same debt calculation cycle
		// we will be able to store in db at which block it became liquidable
		// delete(eng.liquidableBlockTracker, sessionId)
	}
}

func (eng *DebtEngine) GetCumulativeIndexAndDecimalForCMs(blockNum int64, ts uint64) map[string]*ds.CumIndexAndUToken {
	// this is assuming that credit managers are not disabled
	cmAddrs := eng.repo.GetKit().GetAdapterAddressByName(ds.CreditManager)
	cmToCumIndex := make(map[string]*ds.CumIndexAndUToken)
	for _, cmAddr := range cmAddrs {
		cmState := eng.repo.GetCMState(cmAddr)
		// if block_num is before cm exists cmState will be nil
		if cmState == nil {
			continue
		}
		poolAddr := cmState.PoolAddress
		poolInterestData := eng.poolLastInterestData[poolAddr]
		var cumIndexNormalized *big.Int
		if poolInterestData != nil {
			tsDiff := new(big.Int).SetInt64(int64(ts - poolInterestData.Timestamp))
			newInterest := new(big.Int).Mul(poolInterestData.BorrowAPYBI.Convert(), tsDiff)
			newInterestPerSec := new(big.Int).Quo(newInterest, big.NewInt(3600*365*24))
			predicate := new(big.Int).Add(newInterestPerSec, utils.GetExpInt(27))
			cumIndex := new(big.Int).Mul(poolInterestData.CumulativeIndexRAY.Convert(), predicate)
			cumIndexNormalized = utils.GetInt64(cumIndex, 27)
		}
		// set fields
		tokenAddr := cmState.UnderlyingToken
		token := eng.repo.GetToken(tokenAddr)
		cmToCumIndex[cmAddr] = &ds.CumIndexAndUToken{
			CumulativeIndex: cumIndexNormalized,
			Token:           tokenAddr,
			Symbol:          token.Symbol,
			Decimals:        token.Decimals,
			PoolAddr:        poolAddr,
		}
	}
	return cmToCumIndex
}

func (eng *DebtEngine) getTokenPriceFeed(token string, version int16) *schemas.PriceFeed {
	switch version {
	case 1:
		return eng.tokenLastPrice[token]
	case 2:
		return eng.tokenLastPriceV2[token]
	}
	return nil
}
func (eng *DebtEngine) SessionDebtHandler(blockNum int64, session *schemas.CreditSession, cumIndexAndUToken *ds.CumIndexAndUToken) {
	sessionId := session.ID
	sessionSnapshot := eng.lastCSS[sessionId]
	cmAddr := session.CreditManager
	debt, profile := eng.CalculateSessionDebt(blockNum, session, cumIndexAndUToken)
	// if profile is not null
	// yearn price feed might be stale as a result difference btw dc and calculated values
	// solution: fetch price again for all stale yearn feeds
	if profile != nil {
		yearnFeeds := eng.repo.GetYearnFeedAddrs()
		for tokenAddr := range *sessionSnapshot.Balances {
			lastPriceEvent := eng.getTokenPriceFeed(tokenAddr, session.Version)
			// for weth price feed is null
			// if lastPriceEvent != nil {
			// 	log.Infof("%+v\n",lastPriceEvent)
			// 	feed := lastPriceEvent.Feed
			// 	eng.requestPriceFeed(blockNum, feed, tokenAddr)
			// }
			if tokenAddr != eng.repo.GetWETHAddr() && lastPriceEvent.BlockNumber != blockNum {
				feed := lastPriceEvent.Feed
				if utils.Contains(yearnFeeds, feed) {
					eng.requestPriceFeed(blockNum, feed, tokenAddr, lastPriceEvent.IsPriceInUSD)
				}
			}
		}
		debt, profile = eng.CalculateSessionDebt(blockNum, session, cumIndexAndUToken)
		if profile != nil {
			log.Fatalf("Debt fields different from data compressor fields: %s", profile)
			eng.SaveProfile(profile.String())

		}
	}
	// check if data compressor and calculated values match
	eng.liquidationCheck(debt, cmAddr, sessionSnapshot.Borrower, cumIndexAndUToken)
	if session.ClosedAt == blockNum+1 {
		eng.ifAccountLiquidated(sessionId, cmAddr, session.ClosedAt, session.Status)
		eng.addCurrentDebt(debt, cumIndexAndUToken.Decimals)
	}
	eng.AddDebt(debt, sessionSnapshot.BlockNum == blockNum)
}

func (eng *DebtEngine) CalculateSessionDebt(blockNum int64, session *schemas.CreditSession, cumIndexAndUToken *ds.CumIndexAndUToken) (*schemas.Debt, *ds.DebtProfile) {
	sessionId := session.ID
	sessionSnapshot := eng.lastCSS[sessionId]
	cmAddr := session.CreditManager
	accountAddr := session.Account
	calThresholdValue := big.NewInt(0)
	calTotalValue := big.NewInt(0)
	// profiling
	tokenDetails := map[string]ds.TokenDetails{}
	for tokenAddr, balance := range *sessionSnapshot.Balances {
		decimal := eng.repo.GetToken(tokenAddr).Decimals
		price := eng.GetTokenLastPrice(tokenAddr, session.Version)
		tokenLiquidityThreshold := eng.allowedTokensThreshold[cmAddr][tokenAddr]
		// profiling
		tokenDetails[tokenAddr] = ds.TokenDetails{
			Price:             price,
			Decimals:          decimal,
			TokenLiqThreshold: tokenLiquidityThreshold,
			Symbol:            eng.repo.GetToken(tokenAddr).Symbol,
			Version:           session.Version,
		}
		// token not linked continue
		if !balance.Linked {
			continue
		}
		tokenValue := new(big.Int).Mul(price, balance.BI.Convert())
		tokenValueInDecimal := utils.GetInt64(tokenValue, decimal-cumIndexAndUToken.Decimals)
		tokenThresholdValue := new(big.Int).Mul(tokenValueInDecimal, tokenLiquidityThreshold.Convert())
		calThresholdValue = new(big.Int).Add(calThresholdValue, tokenThresholdValue)
		calTotalValue = new(big.Int).Add(calTotalValue, tokenValueInDecimal)

	}
	// the value of credit account is in terms of underlying asset
	underlyingPrice := eng.GetTokenLastPrice(cumIndexAndUToken.Token, session.Version)
	calThresholdValue = new(big.Int).Quo(calThresholdValue, underlyingPrice)
	calTotalValue = new(big.Int).Quo(calTotalValue, underlyingPrice)
	// borrowed + interest and normalized threshold value
	calBorrowWithInterest := big.NewInt(0).Quo(
		big.NewInt(0).Mul(cumIndexAndUToken.CumulativeIndex, sessionSnapshot.BorrowedAmountBI.Convert()),
		sessionSnapshot.СumulativeIndexAtOpen.Convert())
	calReducedThresholdValue := big.NewInt(0).Quo(calThresholdValue, big.NewInt(10000))
	// set debt fields
	debt := &schemas.Debt{
		CommonDebtFields: schemas.CommonDebtFields{
			BlockNumber:                     blockNum,
			CalHealthFactor:                 (*core.BigInt)(big.NewInt(0).Quo(calThresholdValue, calBorrowWithInterest)),
			CalTotalValueBI:                 (*core.BigInt)(calTotalValue),
			CalBorrowedAmountPlusInterestBI: (*core.BigInt)(calBorrowWithInterest),
			CalThresholdValueBI:             (*core.BigInt)(calReducedThresholdValue),
			CollateralInUnderlying:          sessionSnapshot.CollateralInUnderlying,
		},
		SessionId: sessionId,
	}
	var notMatched bool
	profile := ds.DebtProfile{}
	// use data compressor if debt check is enabled
	if eng.config.DebtDCMatching {
		data := eng.SessionDataFromDC(blockNum, cmAddr, sessionSnapshot.Borrower)
		utils.ToJson(data)
		// set debt data fetched from dc
		if !CompareBalance(debt.CalTotalValueBI, (*core.BigInt)(data.TotalValue), cumIndexAndUToken) ||
			!CompareBalance(debt.CalBorrowedAmountPlusInterestBI, (*core.BigInt)(data.BorrowedAmountPlusInterest), cumIndexAndUToken) ||
			core.ValueDifferSideOf10000(debt.CalHealthFactor, (*core.BigInt)(data.HealthFactor)) {
			mask := eng.repo.GetMask(blockNum, cmAddr, accountAddr, session.Version)
			var err error
			profile.RPCBalances, err = eng.repo.ConvertToBalanceWithMask(data.Balances, mask)
			if err != nil {
				log.Fatalf("DC wrong token values block:%d dc:%s", blockNum, eng.repo.GetDCWrapper().ToJson())
			}
			notMatched = true
		}
		// even if data compressor matching is disabled check the values  with values at which last credit snapshot was taken
	} else if sessionSnapshot.BlockNum == blockNum {
		if !CompareBalance(debt.CalTotalValueBI, sessionSnapshot.TotalValueBI, cumIndexAndUToken) ||
			// hf value calculated are on different side of 1
			core.ValueDifferSideOf10000(debt.CalHealthFactor, sessionSnapshot.HealthFactor) ||
			// if healhFactor diff by 4 %
			core.DiffMoreThanFraction(debt.CalHealthFactor, sessionSnapshot.HealthFactor, big.NewFloat(0.04)) {
			profile.RPCBalances = sessionSnapshot.Balances.Copy()
			notMatched = true
		}
	}
	eng.calAmountToPoolAndProfit(debt, session, cumIndexAndUToken)
	if notMatched {
		profile.CumIndexAndUToken = cumIndexAndUToken
		profile.Debt = debt
		profile.CreditSessionSnapshot = sessionSnapshot
		profile.UnderlyingDecimals = cumIndexAndUToken.Decimals
		profile.Tokens = tokenDetails
		return debt, &profile
	}
	return debt, nil
}

// these values are calculated for the borrower not the liquidation, so the calrepayamount takes isliquidated = false
// only for block that the account is liquidated we use isLiquidated set to true so that we can calculate the true amountToPool
//
// amountToPool
// v1 if liquidatecreditaccount is calcClosePayment first return value when isLIquidated= true
// v2 if liquidatecreditaccount is calcClosePayment first return value when isLIquidated= true
//
// remainingfunds is used for profit calculation
// - v1 close or liquidated -- remainingfunds is taken from event
// - v1 repay + open -- remainingfunds is manually calculated with help of calCloseAmount in sdk-go
// - v2 for closecreditaccount, remainingFunds is calculated from the account transfers
// - v2 for liquidatecreditaccount, remainingFunds is taken from event
// - v2 for openedaccounts, totalValue - amountToPool
//
// repayAmount
// v1 - repayAmount = amountToPool, except the blockNum at which account is liquidated
// v2 - close repayAmount is transferred from borrower to account as underlying token
// v2 - liquidation and opened accounts
//      the account might be having some underlying token balance so repayAMount = amountToPool - underlyingToken balance
//
func (eng *DebtEngine) calAmountToPoolAndProfit(debt *schemas.Debt, session *schemas.CreditSession, cumIndexAndUToken *ds.CumIndexAndUToken) {
	var amountToPool, calRemainingFunds *big.Int
	sessionSnapshot := eng.lastCSS[session.ID]
	//
	accountIsLiquidable := session.Status == schemas.Liquidated && session.ClosedAt == debt.BlockNumber+1
	// amount to pool
	amountToPool, calRemainingFunds, _, _ = schemas.CalCloseAmount(eng.lastParameters[session.CreditManager], session.Version, debt.CalTotalValueBI.Convert(), accountIsLiquidable,
		debt.CalBorrowedAmountPlusInterestBI.Convert(),
		sessionSnapshot.BorrowedAmountBI.Convert())

	// calculate profit
	debt.AmountToPoolBI = (*core.BigInt)(amountToPool)
	var remainingFunds *big.Int
	if session.Version == 2 {
		repayAmount := new(big.Int)
		// while close account on v2 we calculate remainingFunds from all the token transfer from the user
		if session.Status == schemas.Closed && session.ClosedAt == debt.BlockNumber+1 {
			prices := core.JsonFloatMap{}
			for token, balance := range *session.Balances {
				tokenPrice := eng.GetTokenLastPrice(token, session.Version)
				price := utils.GetFloat64Decimal(tokenPrice, 8)
				prices[token] = price
				if balance.BI.Convert().Cmp(new(big.Int)) < 0 {
					// assuming there is only one transfer from borrower to account
					// this transfer will be in underlyingtoken. execute_parser.go:246 and
					// https://github.com/Gearbox-protocol/contracts-v2/blob/main/contracts/credit/CreditManager.sol#L286-L291
					repayAmount = new(big.Int).Mul(balance.BI.Convert(), big.NewInt(-1))
				}
			}
			// remainingFunds calculation
			// set price for underlying token
			prices[cumIndexAndUToken.Token] = utils.GetFloat64Decimal(eng.GetTokenLastPrice(cumIndexAndUToken.Token, session.Version), 8)
			remainingFunds = session.Balances.ValueInUnderlying(cumIndexAndUToken.Token, cumIndexAndUToken.Decimals, prices)
		} else if session.ClosedAt == debt.BlockNumber+1 && session.Status == schemas.Liquidated {
			remainingFunds = session.RemainingFunds.Convert()
			repayAmount = new(big.Int)
		} else {
			// repayamount
			// for account not closed yet and account liquidated
			// get underlying balance
			underlying := (*session.Balances)[cumIndexAndUToken.Token]
			underlyingBalance := new(big.Int)
			if underlying != nil {
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
		debt.RepayAmountBI = (*core.BigInt)(repayAmount)
	} else {
		if session.ClosedAt == debt.BlockNumber+1 && (session.Status == schemas.Closed || session.Status == schemas.Liquidated) {
			remainingFunds = (*big.Int)(session.RemainingFunds)
		} else {
			remainingFunds = calRemainingFunds
		}
		if session.Status == schemas.Liquidated && session.ClosedAt == debt.BlockNumber+1 {
			debt.RepayAmountBI = (*core.BigInt)(new(big.Int))
		} else {
			// https://github.com/Gearbox-protocol/gearbox-contracts/blob/master/contracts/credit/CreditManager.sol#L487-L490
			debt.RepayAmountBI = (*core.BigInt)(amountToPool)
		}
	}

	remainingFundsInUSD := eng.GetAmountInUSD(cumIndexAndUToken.Token, remainingFunds, session.Version)
	debt.ProfitInUnderlying = utils.GetFloat64Decimal(remainingFunds, cumIndexAndUToken.Decimals) - debt.CollateralInUnderlying
	debt.CollateralInUnderlying = sessionSnapshot.CollateralInUnderlying
	// fields in USD
	debt.CollateralInUSD = sessionSnapshot.CollateralInUSD
	debt.ProfitInUSD = utils.GetFloat64Decimal(remainingFundsInUSD, 8) - sessionSnapshot.CollateralInUSD
	debt.TotalValueInUSD = utils.GetFloat64Decimal(
		eng.GetAmountInUSD(cumIndexAndUToken.Token, debt.CalTotalValueBI.Convert(), session.Version), 8)
}

func (eng *DebtEngine) GetAmountInUSD(tokenAddr string, amount *big.Int, version int16) *big.Int {
	usdcAddr := eng.repo.GetUSDCAddr()
	tokenPrice := eng.GetTokenLastPrice(tokenAddr, version)
	tokenDecimals := eng.repo.GetToken(tokenAddr).Decimals
	if version == 2 {
		return utils.GetInt64(new(big.Int).Mul(tokenPrice, amount), tokenDecimals)
	}
	usdcPrice := eng.GetTokenLastPrice(usdcAddr, version)
	usdcDecimals := eng.repo.GetToken(usdcAddr).Decimals

	value := new(big.Int).Mul(amount, tokenPrice)
	value = utils.GetInt64(value, tokenDecimals-usdcDecimals)
	value = new(big.Int).Quo(value, usdcPrice)
	return new(big.Int).Mul(value, big.NewInt(100))
}

func (eng *DebtEngine) GetTokenLastPrice(addr string, version int16, dontFail ...bool) *big.Int {
	switch version {
	case 1:
		if eng.tokenLastPrice[addr] != nil {
			return eng.tokenLastPrice[addr].PriceBI.Convert()
		} else if eng.repo.GetWETHAddr() == addr {
			return core.WETHPrice
		}
	case 2:
		if eng.tokenLastPriceV2[addr] != nil {
			return eng.tokenLastPriceV2[addr].PriceBI.Convert()
		}
	}
	if len(dontFail) > 0 && dontFail[0] {
		return nil
	}
	log.Fatalf("Price not found for %s version: %d", addr, version)
	return nil
}

func (eng *DebtEngine) SessionDataFromDC(blockNum int64, cmAddr, borrower string) mainnet.DataTypesCreditAccountDataExtended {
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
	}
	data, err := eng.repo.GetDCWrapper().GetCreditAccountDataExtended(opts,
		common.HexToAddress(cmAddr),
		common.HexToAddress(borrower),
	)
	if err != nil {
		log.Fatalf("cm:%s borrower:%s blocknum:%d err:%s", cmAddr, borrower, blockNum, err)
	}
	return data
}

func (eng *DebtEngine) requestPriceFeed(blockNum int64, feed, token string, isPriceInUSD bool) {
	// PFFIX
	yearnPFContract, err := yearnPriceFeed.NewYearnPriceFeed(common.HexToAddress(feed), eng.client)
	log.CheckFatal(err)
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
	}
	roundData, err := yearnPFContract.LatestRoundData(opts)
	if err != nil {
		log.Fatal(err)
	}
	var decimals int8 = 18 // for eth
	if isPriceInUSD {
		decimals = 8 // for usd
	}
	eng.AddTokenLastPrice(&schemas.PriceFeed{
		BlockNumber:  blockNum,
		Token:        token,
		Feed:         feed,
		RoundId:      roundData.RoundId.Int64(),
		PriceBI:      (*core.BigInt)(roundData.Answer),
		Price:        utils.GetFloat64Decimal(roundData.Answer, decimals),
		IsPriceInUSD: isPriceInUSD,
	})
}
