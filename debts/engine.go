package debts

import (
	"github.com/Gearbox-protocol/third-eye/artifacts/dataCompressor/mainnet"
	"github.com/Gearbox-protocol/third-eye/artifacts/yearnPriceFeed"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"sort"
)

func (eng *DebtEngine) SaveProfile(profile string) {
	err := eng.db.Create(&core.ProfileTable{Profile: profile}).Error
	if err != nil {
		log.Fatal(err)
	}
}

func (eng *DebtEngine) calculateDebt() {
	blocks := eng.repo.GetBlocks()
	sessions := eng.repo.GetSessions()
	noOfBlock := len(blocks)
	blockNums := make([]int64, 0, noOfBlock)
	for blockNum := range blocks {
		blockNums = append(blockNums, blockNum)
	}
	sort.Slice(blockNums, func(i, j int) bool { return blockNums[i] < blockNums[j] })
	for _, blockNum := range blockNums {
		block := blocks[blockNum]
		// update threshold
		ltChangedTokens := []string{}
		for _, allowedToken := range block.GetAllowedTokens() {
			eng.AddAllowedTokenThreshold(allowedToken)
			ltChangedTokens = append(ltChangedTokens, allowedToken.Token)
		}
		// update pool borrow rate and cumulative index
		for _, ps := range block.GetPoolStats() {
			eng.AddPoolLastInterestData(&core.PoolInterestData{
				Address:            ps.Address,
				BlockNum:           ps.BlockNum,
				CumulativeIndexRAY: ps.CumulativeIndexRAY,
				BorrowAPYBI:        ps.BorrowAPYBI,
				Timestamp:          block.Timestamp,
			})
		}

		// update balance of last credit session snapshot and create credit session snapshots
		sessionsToUpdate := make(map[string]bool)
		for _, css := range block.GetCSS() {
			eng.AddLastCSS(css)
			sessionsToUpdate[css.SessionId] = true
		}
		// set the price session list to update
		sessionWithTokens := make(map[string][]string)
		for _, session := range sessions {
			if (session.ClosedAt != 0 && session.ClosedAt <= blockNum) || session.Since > blockNum {
				continue
			}
			sessionSnapshot := eng.lastCSS[session.ID]
			for tokenAddr := range *sessionSnapshot.Balances {
				sessionWithTokens[tokenAddr] = append(sessionWithTokens[tokenAddr], session.ID)
			}
		}
		// update price
		for _, pf := range block.GetPriceFeeds() {
			eng.AddTokenLastPrice(pf)
			// set the price session list to update
			for _, sessionId := range sessionWithTokens[pf.Token] {
				sessionsToUpdate[sessionId] = true
			}
		}
		for _, token := range ltChangedTokens {
			for _, sessionId := range sessionWithTokens[token] {
				sessionsToUpdate[sessionId] = true
			}
		}

		for _, params := range block.GetParams() {
			eng.addLastParameters(params)
		}
		// get pool cumulative interest rate
		cmAddrToCumIndex := eng.GetCumulativeIndexAndDecimalForCMs(blockNum, block.Timestamp)
		// calculate each session debt
		for sessionId := range sessionsToUpdate {
			session := sessions[sessionId]
			if (session.ClosedAt != 0 && session.ClosedAt <= blockNum) || session.Since > blockNum {
				continue
			}
			cmAddr := session.CreditManager
			// pool cum index is when the pool is not registered
			if cmAddrToCumIndex[cmAddr] != nil {
				eng.SessionDebtHandler(blockNum, session, cmAddrToCumIndex[cmAddr])
				// send notification when account is liquidated
			} else {
				log.Fatalf("CM(%s):pool is missing stats at %d, so cumulative index of pool is unknown", cmAddr, blockNum)
			}
		}
		if len(sessionsToUpdate) > 0 {
			log.Verbosef("Calculated %d debts for block %d", len(sessionsToUpdate), blockNum)
		}
		eng.flushDebt(blockNum)
	}
	// if noOfBlock > 0 {
	// 	eng.flushDebt(blockNums[noOfBlock-1])
	// }
}

func (eng *DebtEngine) ifAccountClosed(sessionId, cmAddr string, closedAt int64, status int) {
	sessionSnapshot := eng.lastCSS[sessionId]
	if status == core.Liquidated {
		account := eng.liquidableBlockTracker[sessionId]
		var liquidableSinceBlockNum int64
		if account != nil {
			liquidableSinceBlockNum = account.BlockNum
		} else {
			log.Warnf("Session(%s) liquidated at block:%d, but liquidable since block not stored", sessionId, closedAt)
		}
		urls := eng.networkUIUrl()
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
	} else if status != core.Active {
		// comment deletion form map
		// this way if the account is liquidable and liquidated in the same debt calculation cycle
		// we will be able to store in db at which block it became liquidable
		// delete(eng.liquidableBlockTracker, sessionId)
	}
}

func (eng *DebtEngine) GetCumulativeIndexAndDecimalForCMs(blockNum int64, ts uint64) map[string]*core.CumIndexAndUToken {
	cmAddrs := eng.repo.GetKit().GetAdapterAddressByName(core.CreditManager)
	poolToCI := make(map[string]*core.CumIndexAndUToken)
	for _, cmAddr := range cmAddrs {
		poolAddr := eng.repo.GetCMState(cmAddr).PoolAddress
		poolInterestData := eng.poolLastInterestData[poolAddr]
		if poolInterestData == nil {
			poolToCI[poolAddr] = nil
		} else {
			tsDiff := new(big.Int).SetInt64(int64(ts - poolInterestData.Timestamp))
			newInterest := new(big.Int).Mul(poolInterestData.BorrowAPYBI.Convert(), tsDiff)
			newInterestPerSec := new(big.Int).Quo(newInterest, big.NewInt(3600*365*24))
			predicate := new(big.Int).Add(newInterestPerSec, utils.GetExpInt(27))
			cumIndex := new(big.Int).Mul(poolInterestData.CumulativeIndexRAY.Convert(), predicate)
			cumIndexNormalized := utils.GetInt64(cumIndex, 27)
			tokenAddr := eng.repo.GetCMState(cmAddr).UnderlyingToken
			token := eng.repo.GetToken(tokenAddr)
			poolToCI[cmAddr] = &core.CumIndexAndUToken{
				CumulativeIndex: cumIndexNormalized,
				Token:           tokenAddr,
				Symbol:          token.Symbol,
				Decimals:        token.Decimals,
				Price:           eng.GetTokenLastPrice(tokenAddr),
			}
			// log.Infof("blockNum%d newInterest:%s tsDiff:%s cumIndexDecimal:%s predicate:%s cumIndex:%s",blockNum ,newInterest, tsDiff, cumIndexNormalized, predicate, cumIndex)
		}
	}
	return poolToCI
}

func (eng *DebtEngine) SessionDebtHandler(blockNum int64, session *core.CreditSession, cumIndexAndUToken *core.CumIndexAndUToken) {
	sessionId := session.ID
	sessionSnapshot := eng.lastCSS[sessionId]
	cmAddr := session.CreditManager
	debt, profile := eng.CalculateSessionDebt(blockNum, session, cumIndexAndUToken)
	// if profile is not null
	// yearn price feed might be stale as a result difference btw dc and calculated values
	// solution: fetch price again for all stale yearn feeds
	if profile != nil {
		kit := eng.repo.GetKit()
		yearnFeeds := kit.GetAdapterAddressByName(core.YearnPriceFeed)
		for tokenAddr := range *sessionSnapshot.Balances {
			lastPriceEvent := eng.tokenLastPrice[tokenAddr]
			// for weth price feed is null
			if lastPriceEvent != nil {
				feed := lastPriceEvent.Feed
				eng.requestPriceFeed(blockNum, feed, tokenAddr)
			}
			if tokenAddr != eng.repo.GetWETHAddr() && lastPriceEvent.BlockNumber != blockNum {
				feed := lastPriceEvent.Feed
				if utils.Contains(yearnFeeds, feed) {
					eng.requestPriceFeed(blockNum, feed, tokenAddr)
				}
			}
		}
		debt, profile = eng.CalculateSessionDebt(blockNum, session, cumIndexAndUToken)
		if profile != nil {
			log.Fatalf("Debt fields different from data compressor fields: %s", profile.Json())
			eng.SaveProfile(string(profile.Json()))

		}
	}
	// check if data compressor and calculated values match
	eng.liquidationCheck(debt, cmAddr, sessionSnapshot.Borrower, cumIndexAndUToken)
	if session.ClosedAt == blockNum+1 {
		eng.ifAccountClosed(sessionId, cmAddr, session.ClosedAt, session.Status)
		eng.addCurrentDebt(debt, cumIndexAndUToken.Decimals)
	}
	eng.AddDebt(debt, sessionSnapshot.BlockNum == blockNum)
}

func (eng *DebtEngine) CalculateSessionDebt(blockNum int64, session *core.CreditSession, cumIndexAndUToken *core.CumIndexAndUToken) (*core.Debt, *core.DebtProfile) {
	sessionId := session.ID
	sessionSnapshot := eng.lastCSS[sessionId]
	cmAddr := session.CreditManager
	accountAddr := session.Account
	calThresholdValue := big.NewInt(0)
	calTotalValue := big.NewInt(0)
	// profiling
	tokenDetails := map[string]core.TokenDetails{}
	for tokenAddr, balance := range *sessionSnapshot.Balances {
		decimal := eng.repo.GetToken(tokenAddr).Decimals
		price := eng.GetTokenLastPrice(tokenAddr)
		tokenLiquidityThreshold := eng.allowedTokensThreshold[cmAddr][tokenAddr]
		// profiling
		tokenDetails[tokenAddr] = core.TokenDetails{
			Price:             price,
			Decimals:          decimal,
			TokenLiqThreshold: tokenLiquidityThreshold,
			Symbol:            eng.repo.GetToken(tokenAddr).Symbol}
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
	underlyingPrice := cumIndexAndUToken.Price
	calThresholdValue = new(big.Int).Quo(calThresholdValue, underlyingPrice)
	calTotalValue = new(big.Int).Quo(calTotalValue, underlyingPrice)
	// borrowed + interest and normalized threshold value
	calBorrowWithInterest := big.NewInt(0).Quo(
		big.NewInt(0).Mul(cumIndexAndUToken.CumulativeIndex, sessionSnapshot.BorrowedAmountBI.Convert()),
		sessionSnapshot.СumulativeIndexAtOpen.Convert())
	calReducedThresholdValue := big.NewInt(0).Quo(calThresholdValue, big.NewInt(10000))
	// set debt fields
	debt := &core.Debt{
		BlockNumber:                     blockNum,
		SessionId:                       sessionId,
		CalHealthFactor:                 (*core.BigInt)(big.NewInt(0).Quo(calThresholdValue, calBorrowWithInterest)),
		CalTotalValueBI:                 (*core.BigInt)(calTotalValue),
		CalBorrowedAmountPlusInterestBI: (*core.BigInt)(calBorrowWithInterest),
		CalThresholdValueBI:             (*core.BigInt)(calReducedThresholdValue),
		CollateralInUnderlying:          utils.GetFloat64Decimal(sessionSnapshot.CollateralInUnderlying.Convert(), cumIndexAndUToken.Decimals),
	}
	var notMatched bool
	profile := core.DebtProfile{}
	// use data compressor if debt check is enabled
	if eng.config.DebtDCMatching {
		data := eng.SessionDataFromDC(blockNum, cmAddr, sessionSnapshot.Borrower)
		utils.ToJson(data)
		// set debt data fetched from dc
		debt.HealthFactor = (*core.BigInt)(data.HealthFactor)
		debt.TotalValueBI = (*core.BigInt)(data.TotalValue)
		debt.BorrowedAmountPlusInterestBI = (*core.BigInt)(data.BorrowedAmountPlusInterest)
		if !core.CompareBalance(debt.CalTotalValueBI, debt.TotalValueBI, cumIndexAndUToken) ||
			!core.CompareBalance(debt.CalBorrowedAmountPlusInterestBI, debt.BorrowedAmountPlusInterestBI, cumIndexAndUToken) ||
			core.ValueDifferSideOf10000(debt.CalHealthFactor, debt.HealthFactor) {
			mask := eng.repo.GetMask(blockNum, cmAddr, accountAddr)
			var err error
			profile.RPCBalances, err = eng.repo.ConvertToBalanceWithMask(data.Balances, mask)
			if err != nil {
				log.Fatalf("DC wrong token values block:%d dc:%s", blockNum, eng.repo.GetDCWrapper().ToJson())
			}
			notMatched = true
		}
		// even if data compressor matching is disabled check the values  with values at which last credit snapshot was taken
	} else if sessionSnapshot.BlockNum == blockNum {
		debt.HealthFactor = sessionSnapshot.HealthFactor
		debt.TotalValueBI = sessionSnapshot.TotalValueBI
		if !core.CompareBalance(debt.CalTotalValueBI, debt.TotalValueBI, cumIndexAndUToken) ||
			// hf value calculated are on different side of 1
			core.ValueDifferSideOf10000(debt.CalHealthFactor, debt.HealthFactor) ||
			// if healhFactor diff by 4 %
			core.DiffMoreThanFraction(debt.CalHealthFactor, debt.HealthFactor, big.NewFloat(0.04)) {
			profile.RPCBalances = sessionSnapshot.Balances.Copy()
			notMatched = true
		}
	}
	if notMatched {
		profile.CumIndexAndUToken = cumIndexAndUToken
		profile.Debt = debt
		profile.CreditSessionSnapshot = sessionSnapshot
		profile.UnderlyingDecimals = cumIndexAndUToken.Decimals
		profile.Tokens = tokenDetails
		return debt, &profile
	}
	eng.calAmountToPoolAndProfit(debt, session, cumIndexAndUToken)
	return debt, nil
}

func (eng *DebtEngine) calAmountToPoolAndProfit(debt *core.Debt, session *core.CreditSession, cumIndexAndUToken *core.CumIndexAndUToken) {
	var amountToPool *big.Int
	sessionSnapshot := eng.lastCSS[session.ID]
	// amount to pool
	// if account is liquidable
	if debt.CalHealthFactor.Convert().Cmp(big.NewInt(10000)) < 0 {
		amountToPool, _, _ = eng.calCloseAmount(session.CreditManager, debt.CalTotalValueBI, true,
			debt.CalBorrowedAmountPlusInterestBI.Convert(),
			sessionSnapshot.BorrowedAmountBI.Convert())
	} else {
		amountToPool, _, _ = eng.calCloseAmount(session.CreditManager, debt.CalTotalValueBI, false,
			debt.CalBorrowedAmountPlusInterestBI.Convert(),
			sessionSnapshot.BorrowedAmountBI.Convert())
	}
	debt.AmountToPoolBI = (*core.BigInt)(amountToPool)
	// calculate profit
	var remainingFunds *big.Int
	if (session.ClosedAt == 0 || session.ClosedAt != debt.BlockNumber+1) || session.Status == core.Repaid {
		remainingFunds = new(big.Int).Sub(debt.CalTotalValueBI.Convert(), debt.AmountToPoolBI.Convert())
	} else {
		remainingFunds = (*big.Int)(session.RemainingFunds)
	}
	remainingFundsInUSD := eng.GetAmountInUSD(cumIndexAndUToken.Token, remainingFunds)
	profitInUSD := new(big.Int).Sub(remainingFundsInUSD, sessionSnapshot.CollateralInUSD.Convert())
	debt.ProfitInUnderlying = utils.GetFloat64Decimal(remainingFunds, cumIndexAndUToken.Decimals) - debt.CollateralInUnderlying
	// fields in USD
	debt.CollateralInUSD = utils.GetFloat64Decimal(sessionSnapshot.CollateralInUSD.Convert(), 8)
	debt.ProfitInUSD = utils.GetFloat64Decimal(profitInUSD, 8)
	debt.TotalValueInUSD = utils.GetFloat64Decimal(eng.GetAmountInUSD(cumIndexAndUToken.Token, debt.CalTotalValueBI.Convert()), 8)
}

func (eng *DebtEngine) GetAmountInUSD(tokenAddr string, amount *big.Int) *big.Int {
	usdcAddr := eng.repo.GetUSDCAddr()
	tokenPrice := eng.GetTokenLastPrice(tokenAddr)
	usdcPrice := eng.GetTokenLastPrice(usdcAddr)
	tokenDecimals := eng.repo.GetToken(tokenAddr).Decimals
	usdcDecimals := eng.repo.GetToken(usdcAddr).Decimals

	value := new(big.Int).Mul(amount, tokenPrice)
	value = utils.GetInt64(value, tokenDecimals-usdcDecimals)
	value = new(big.Int).Quo(value, usdcPrice)
	return new(big.Int).Mul(value, big.NewInt(100))
}

func (eng *DebtEngine) GetTokenLastPrice(addr string) *big.Int {
	if eng.repo.GetWETHAddr() == addr {
		return core.WETHPrice
	} else {
		return eng.tokenLastPrice[addr].PriceETHBI.Convert()
	}
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

func (eng *DebtEngine) requestPriceFeed(blockNum int64, feed, token string) {
	yearnPFContract, err := yearnPriceFeed.NewYearnPriceFeed(common.HexToAddress(feed), eng.client)
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
	}
	roundData, err := yearnPFContract.LatestRoundData(opts)
	if err != nil {
		log.Fatal(err)
	}
	eng.AddTokenLastPrice(&core.PriceFeed{
		BlockNumber: blockNum,
		Token:       token,
		Feed:        feed,
		RoundId:     roundData.RoundId.Int64(),
		PriceETHBI:  (*core.BigInt)(roundData.Answer),
		PriceETH:    utils.GetFloat64Decimal(roundData.Answer, 18),
	})
}
