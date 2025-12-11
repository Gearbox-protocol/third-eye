package debts

import (
	"math/big"
	"sort"
	"time"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/calc"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg/dc"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
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
	// L6 rebaseToken details for stETH
	// L7 pool quota details
	//
	// L1:update params
	for _, params := range block.GetParams() {
		eng.addLastParameters(params)
	}
	// L6: rebaseToken
	for _, params := range block.RebaseDetailsForDB {
		eng.lastRebaseDetails = params
	}
	// L7: poolQuotasDetails
	for _, quotaDetails := range block.QuotaDetails {
		eng.AddPoolQuotaDetails(quotaDetails)
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
			Address:              ps.Address,
			BlockNum:             ps.BlockNum,
			CumulativeIndexRAY:   ps.CumulativeIndexRAY,
			AvailableLiquidityBI: ps.AvailableLiquidityBI,
			ExpectedLiqBI:        ps.ExpectedLiquidityBI,
			BaseBorrowAPYBI:      ps.BaseBorrowAPYBI,
			Timestamp:            block.Timestamp,
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
	for _, ltRamp := range block.GetTokenLTRamp() {
		eng.AddTokenLTRamp(ltRamp)
	}

	for _, to := range block.TokenOracles {
		if !to.Reserve { // only main
			eng.priceHandler.AddTokenOracle(to)
		}
	}
	//
	for _, relation := range block.Relations {
		if relation.Type == "PoolOracle" {
			eng.priceHandler.poolToPriceOracle[relation.Owner] = schemas.PriceOracleT(relation.Dependent)
		}
	}
	// C3.b: updated price
	for _, pf := range block.GetPriceFeeds() {
		// L5
		eng.priceHandler.AddTokenLastPrice(pf)
		// TIMECOMPLEX
		// use mdl.repo.getsyncadapter(pf.Feed).gettokens(block_num)
		for _, oracles := range eng.priceHandler.poTotokenOracle {
			for _, oracle := range oracles {
				if pf.Feed == oracle.Feed {
					tokensUpdated[oracle.Token] = true
				}
			}
		}
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
		eng.currentTs = block.Timestamp
		//
		poolsUpdated, tokensUpdated, sessionsUpdated := eng.updateLocalState(blockNum, block)
		// get pool cumulative interest rate
		cmToPoolDetails := eng.GetCumulativeIndexAndDecimalForCMs(blockNum, block.Timestamp)

		//
		marketToTvl := make(MarketToTvl)
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
			if sessionSnapshot == nil {
				if session.ID != "0x430789365519A006B110727C1c95032475e26Dbf_21786407_69" {
					log.Warnf("Session(%s) not found in last css %d", session.ID, blockNum)
				}
				continue
			}
			caValue := utils.GetFloat64Decimal(
				eng.GetAmountInUSD(
					session.CreditManager,
					cmToPoolDetails[session.CreditManager].Token,
					sessionSnapshot.TotalValueBI.Convert(), session.Version,
				), 8)
			{
				pool := eng.priceHandler.GetPoolFromCM(session.CreditManager)
				adapter := eng.repo.GetAdapter(pool)
				state := adapter.GetUnderlyingState()
				if state == nil {
					log.Fatal("State for pool not found for address: ", pool)
				}
				//
				market := state.(*schemas.PoolState).Market
				marketToTvl.add(market, caValue, 0, 0)
			}

			for token, tokenBalance := range *sessionSnapshot.Balances {
				if tokenBalance.IsEnabled && tokenBalance.HasBalanceMoreThanOne() {
					if tokensUpdated[token] {
						sessionsUpdated[session.ID] = true
					}
				}
			}
			// #C3
			underlyingToken := cmToPoolDetails[session.CreditManager].Token
			if tokensUpdated[underlyingToken] {
				sessionsUpdated[session.ID] = true
			}
		}
		//
		idsToCalculateDebtFor := []string{}
		if len(eng.lastStateOfDebt) == 0 {
			for id := range sessions {
				idsToCalculateDebtFor = append(idsToCalculateDebtFor, id)
			}
		} else {
			for sessionId := range sessionsUpdated {
				idsToCalculateDebtFor = append(idsToCalculateDebtFor, sessionId)
			}
		}
		// calculate each session debt
		for _, sessionId := range idsToCalculateDebtFor {
			session := sessions[sessionId]
			if (session.ClosedAt != 0 && session.ClosedAt <= blockNum) || session.Since > blockNum {
				continue
			}
			if eng.lastCSS[sessionId] == nil {
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
		for _, session := range sessions {
			if (session.ClosedAt != 0 && session.ClosedAt <= blockNum) || session.Since > blockNum {
				continue
			}
			pool := eng.priceHandler.GetPoolFromCM(session.CreditManager)
			adapter := eng.repo.GetAdapter(pool)
			state := adapter.GetUnderlyingState()
			if state == nil {
				log.Fatal("State for pool not found for address: ", pool)
			}
			//
			market := state.(*schemas.PoolState).Market
			if debt := eng.lastStateOfDebt[session.ID]; debt != nil {
				marketToTvl.add(market, 0, 0, 0, debt.TotalValueInUSD)
			}
		}
		//
		eng.createTvlSnapshots(blockNum, marketToTvl)
		if len(sessionsUpdated) > 0 {
			log.Debugf("Calculated %d debts for block %d", len(sessionsUpdated), blockNum)
		}
	}
}

func (eng *DebtEngine) createTvlSnapshots(blockNum int64, marketToTvl MarketToTvl) {
	// if eng.lastTvlSnapshot != nil && blockNum-eng.lastTvlSnapshot.BlockNum < core.NoOfBlocksPerHr { // tvl snapshot every hour
	// 	return
	// }
	for _, entry := range eng.poolLastInterestData {
		adapter := eng.repo.GetAdapter(entry.Address)
		state := adapter.GetUnderlyingState()
		if state == nil {
			log.Fatal("State for pool not found for address: ", entry.Address)
		}
		//
		underlyingToken := state.(*schemas.PoolState).UnderlyingToken
		//
		// latestOracle, version, err := eng.repo.GetActivePriceOracleByBlockNum(blockNum)
		// log.CheckFatal(err)
		//
		fn := func(amount *core.BigInt) float64 {
			if amount.String() == "0" {
				return 0
			}
			return utils.GetFloat64Decimal(
				eng.GetPriceForTvl(
					schemas.PriceOracleT(""),
					state.(*schemas.PoolState).PriceOracle,
					underlyingToken,
					amount.Convert(), state.(*schemas.PoolState).Version), 8)
		}
		availLiq := fn(entry.AvailableLiquidityBI)
		expectedLiqInUSD := fn(entry.ExpectedLiqBI)
		marketToTvl.add(state.(*schemas.PoolState).Market, 0, availLiq, expectedLiqInUSD)
	}
	// save as last tvl snapshot and add to db
	addedMarket := []string{}
	for market, details := range marketToTvl {
		if lastTvlBlock, ok := eng.marketTolastTvlBlock[market]; ok && blockNum-lastTvlBlock < core.BlockPer(core.GetBaseChainId(eng.client), time.Hour) { // only snap her hr.
			continue
		}
		//
		tvl := &schemas.TvlSnapshots{
			BlockNum:           blockNum,
			AvailableLiquidity: details.totalAvailableLiquidity,
			CATotalValue:       details.caTotalValue,
			ExpectedLiq:        details.expectedLiq,
			CATotalValueCalc:   details.caTotalValueCalculated,
			Market:             market,
		}
		addedMarket = append(addedMarket, market)
		eng.tvlSnapshots = append(eng.tvlSnapshots, tvl)
		eng.marketTolastTvlBlock[tvl.Market] = tvl.BlockNum
	}
	if len(addedMarket) > 0 {
		log.Infof("%d:Tvl snapshot added for market %s", blockNum, addedMarket)
	}
}

func (eng *DebtEngine) ifAccountLiquidated(sessionId, cmAddr string, closedAt int64, status int) {
	sessionSnapshot := eng.lastCSS[sessionId]
	if schemas.IsStatusLiquidated(status) {
		account := eng.liquidableBlockTracker[sessionId]
		var liquidableSinceBlockNum int64
		if account != nil {
			liquidableSinceBlockNum = account.BlockNum
		} else {
			log.Warnf("Session(%s) liquidated at block:%d, but liquidable since block not stored", sessionId, closedAt)
		}
		urls := log.NetworkUIUrl(core.GetChainId(eng.client))
		eng.repo.RecentMsgf(log.RiskHeader{
			BlockNumber: closedAt,
			EventCode:   "AMQP",
		}, `Liquidation Alert:
		CreditManager: %s/address/%s
		Tx: %s/tx/%s
		Borrower: %s
		LiquidatedAt: %d 
		Liquidable since: %d
		web: %s/accounts/%s`,
			urls.ExplorerUrl, cmAddr,
			urls.ExplorerUrl, eng.GetLiquidationTx(sessionId),
			sessionSnapshot.Borrower,
			closedAt, liquidableSinceBlockNum,
			urls.ChartUrl, sessionSnapshot.SessionId)
	} else if status != schemas.Active {
		// comment deletion form map
		// this way if the account is liquidable and liquidated in the same debt calculation cycle
		// we will be able to store in db at which block it became liquidable
		// delete(eng.liquidableBlockTracker, sessionId)
	}
}

func (eng *DebtEngine) GetCumulativeIndexAndDecimalForCMs(blockNum int64, ts uint64) map[string]*ds.CumIndexAndUToken {
	// this is assuming that credit managers are not disabled
	cmAddrs := eng.repo.GetAdapterAddressByName(ds.CreditManager)
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
			newInterest := new(big.Int).Mul(poolInterestData.BaseBorrowAPYBI.Convert(), tsDiff)
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

func (eng *DebtEngine) SessionDebtHandler(blockNum int64, session *schemas.CreditSession, cumIndexAndUToken *ds.CumIndexAndUToken) {
	sessionId := session.ID
	sessionSnapshot := eng.lastCSS[sessionId]
	cmAddr := session.CreditManager
	debt, profile := eng.CalculateSessionDebt(blockNum, session, cumIndexAndUToken)
	// if profile is not null
	// yearn price feed might be stale as a result difference btw dc and calculated values
	// solution: fetch price again for all stale yearn feeds
	if profile != nil || debt.CalHealthFactor.Convert().Cmp(big.NewInt(10000)) < 1 {
		retryFeeds := eng.repo.GetRetryFeedForDebts()
		// eng.checkretryFeeds(blockNum, sessionSnapshot.Balances, session, profile != nil)
		for tokenAddr, details := range *sessionSnapshot.Balances {
			if details.IsEnabled && details.HasBalanceMoreThanOne() && tokenAddr != eng.repo.GetWETHAddr() {
				//
				lastPriceEvent := eng.priceHandler.GetLastPriceFeed(session.CreditManager, tokenAddr, session.Version) // don't use reserve
				if lastPriceEvent.BlockNumber != blockNum {
					feedAddr := lastPriceEvent.Feed
					for _, retryFeed := range retryFeeds {
						if retryFeed.GetAddress() == feedAddr {
							log.Infof("retryfeed(%s) for blockNum %d", retryFeed.GetAddress(), blockNum)
							// log.Info("hf ", debt.CalHealthFactor.Convert(), "of", sessionId, "at", blockNum)
							eng.priceHandler.requestPriceFeed(blockNum, eng.client, retryFeed, tokenAddr, profile != nil, eng.db)
						}
					}
				}
			}
		}
		debt, profile = eng.CalculateSessionDebt(blockNum, session, cumIndexAndUToken)
		if profile != nil && (session.ID != "0x36bc4d7f24Ab0e9ACa5a84766152447E4F4B9694_14793256_370" && session.ID != "0x11956F74EA4ac57897e7D34419f9cAD467a868D2_21776080_38") { // this mainnet accounts fails due to some difference in the total vlaue bu the hf is same. why?
			log.Errorf("Debt fields different from data compressor fields: %s", profile)
		}
	}
	// check if data compressor and calculated values match
	eng.liquidationCheck(debt, cmAddr, sessionSnapshot.Borrower, cumIndexAndUToken)
	// for liq v3
	if secStatus := session.StatusAt(blockNum); schemas.IsStatusLiquidated(secStatus) {
		eng.ifAccountLiquidated(sessionId, cmAddr, blockNum, secStatus)
	}
	// for v1/v2 close and liq
	// v3 close
	if session.ClosedAt == blockNum+1 {
		eng.ifAccountLiquidated(sessionId, cmAddr, session.ClosedAt, session.Status)
		eng.addCurrentDebt(debt, cumIndexAndUToken.Decimals)
	}
	eng.AddDebt(debt, sessionSnapshot.BlockNum == blockNum)
}

// todo fix, also check the feed type at this block number
func (eng *DebtEngine) hasRedStoneToken(cm string, version core.VersionType, balances *core.DBBalanceFormat) bool {
	for tokenAddr, details := range *balances {
		if !(details.IsEnabled && details.HasBalanceMoreThanOne()) {
			continue
		}
		priceFeed := eng.priceHandler.GetLastPriceFeed(cm, tokenAddr, version, true)
		if priceFeed == nil { // don't know if reverse
			continue
		}
		feed := eng.repo.GetAdapter(priceFeed.Feed)
		pfType := feed.GetDetails()["pfType"]
		if pfType == "RedStonePF" || pfType == "CompositeRedStonePF" {
			return true
		}
	}
	return false
}

func (eng *DebtEngine) CalculateSessionDebt(blockNum int64, session *schemas.CreditSession, cumIndexAndUToken *ds.CumIndexAndUToken) (*schemas.Debt, *ds.DebtProfile) {
	sessionId := session.ID
	sessionSnapshot := eng.lastCSS[sessionId]
	cmAddr := session.CreditManager
	//
	// calculating account fields
	calculator := calc.Calculator{Store: storeForCalc{inner: eng}}
	// failure true as we are calculating debt
	calHF, calTotalValue, calThresholdValue, debtDetails, profileStr := calculator.CalcAccountFields(
		eng.currentTs,
		blockNum,
		poolDetailsForCalc{
			cumIndexAndUToken: cumIndexAndUToken,
			forQuotas:         eng.v3DebtDetails,
		},
		sessionDetailsForCalc{
			addr:                  session.Account,
			client:                eng.client,
			CreditSessionSnapshot: sessionSnapshot,
			CM:                    session.CreditManager,
			rebaseDetails:         eng.lastRebaseDetails,
			version:               session.Version,
			forQuotas:             eng.v3DebtDetails,
		},
		eng.lastParameters[session.CreditManager].FeeInterest,
		true,
	)
	if calHF.Cmp(big.NewInt(65535)) > 0 && sessionId != "0xc34fef41FA5De8298f6F4a90F3E37708FC4d9447_13938051_134" {
		log.Warn("HF for session", sessionId, "is more than 65535", calHF, "at", blockNum)
		calHF = big.NewInt(65535)
	}

	// the value of credit account is in terms of underlying asset
	// set debt fields
	debt := &schemas.Debt{
		CommonDebtFields: schemas.CommonDebtFields{
			BlockNumber:     blockNum,
			CalHealthFactor: (*core.BigInt)(calHF),
			CalTotalValueBI: (*core.BigInt)(calTotalValue),
			// it has fees too for v2
			CalDebtBI: (*core.BigInt)(debtDetails.Total()),
			// used for calculating the close amount and for comparing with data compressor results as v2 doesn't have borrowedAmountWithInterest
			CalThresholdValueBI:    (*core.BigInt)(calThresholdValue),
			CollateralInUnderlying: sessionSnapshot.CollateralInUnderlying,
		},
		SessionId: sessionId,
	}
	var notMatched bool
	profile := ds.DebtProfile{CreditSessionSnapshot: sessionSnapshot, Tokens: map[string]ds.TokenDetails{}, CalcString: profileStr}
	for tokenAddr, details := range *sessionSnapshot.Balances {
		if details.IsEnabled {
			profile.Tokens[tokenAddr] = ds.TokenDetails{
				Price:             eng.priceHandler.GetLastPrice(session.CreditManager, tokenAddr, session.Version), // don't use reserve
				Decimals:          eng.repo.GetToken(tokenAddr).Decimals,
				TokenLiqThreshold: eng.tokenLTRamp[session.CreditManager][tokenAddr].GetLTForTs(eng.currentTs),
			}
		}
	}

	// use data compressor if debt check is enabled
	if eng.config.DebtDCMatching {
		data := eng.SessionDataFromDC(session.Version, blockNum,
			cmAddr,
			sessionSnapshot.Borrower,
			session.Account,
		)
		utils.ToJson(data)
		// set debt data fetched from dc
		// if healthfactor on diff
		if IsChangeMoreThanFraction(debt.CalTotalValueBI, (*core.BigInt)(data.TotalValue), big.NewFloat(0.0001)) ||
			IsChangeMoreThanFraction(debt.CalDebtBI, (*core.BigInt)(data.Debt), big.NewFloat(0.0001)) ||
			core.ValueDifferSideOf10000(debt.CalHealthFactor, (*core.BigInt)(data.HealthFactor)) {
			profile.DCData = &data
			notMatched = true
		}
		// even if data compressor matching is disabled check the calc values  with session data at block where last credit snapshot was taken
		//  // 20563217 and 0xe8f5F52842D7AF4BbcF5Fe731A336147B51F09D5_19980779_297 on mainnet has creditsessionsnapshot but isSuccessful for dv3 is false.
	} else if sessionSnapshot.BlockNum == blockNum && sessionSnapshot.HealthFactor.Convert().Cmp(new(big.Int)) != 0 &&
		!(utils.Contains([]int64{20563217, 20671148}, blockNum) && sessionSnapshot.TotalValueBI.Convert().Cmp(new(big.Int)) == 0) {
		// it is 0 when the issuccessful is false for redstone credit accounts
		if IsChangeMoreThanFraction(debt.CalTotalValueBI, sessionSnapshot.TotalValueBI, big.NewFloat(0.001)) ||
			// hf value calculated are on different side of 1
			core.ValueDifferSideOf10000(debt.CalHealthFactor, sessionSnapshot.HealthFactor) ||
			// if healhFactor diff by 4 %
			core.DiffMoreThanFraction(debt.CalHealthFactor, sessionSnapshot.HealthFactor, big.NewFloat(0.01)) {
			// log.Info(debt.CalHealthFactor, sessionSnapshot.HealthFactor, blockNum)
			// log.Info(debt.CalTotalValueBI, sessionSnapshot.TotalValueBI, blockNum)
			if log.GetBaseNet(core.GetChainId(eng.client)) == "ARBITRUM" {
				profile.CumIndexAndUToken = cumIndexAndUToken
				profile.Debt = debt
				profile.CreditSessionSnapshot = sessionSnapshot
				profile.UnderlyingDecimals = cumIndexAndUToken.Decimals
				log.Warn(utils.ToJson(profile))
				// } else if eng.hasRedStoneToken(sessionSnapshot.Balances) {
				// 	if IsChangeMoreThanFraction(debt.CalTotalValueBI, sessionSnapshot.TotalValueBI, big.NewFloat(0.003)) { // .3% allowed
				// 		notMatched = true
				// 	}
			} else {
				notMatched = true
			}
			log.Info(
				debt.CalTotalValueBI, sessionSnapshot.TotalValueBI, sessionId,
				IsChangeMoreThanFraction(debt.CalTotalValueBI, sessionSnapshot.TotalValueBI, big.NewFloat(0.001)),
				// hf value calculated are on different side of 1
				core.ValueDifferSideOf10000(debt.CalHealthFactor, sessionSnapshot.HealthFactor),
				// if healhFactor diff by 4 %
				core.DiffMoreThanFraction(debt.CalHealthFactor, sessionSnapshot.HealthFactor, big.NewFloat(0.01)),
			)
		}
	}
	eng.calAmountToPoolAndProfit(debt, session, cumIndexAndUToken, debtDetails)
	// eng.farmingCalc.addFarmingVal(debt, session, eng.lastCSS[session.ID], storeForCalc{inner: eng})
	if notMatched {
		profile.CumIndexAndUToken = cumIndexAndUToken
		profile.Debt = debt
		profile.CreditSessionSnapshot = sessionSnapshot
		profile.UnderlyingDecimals = cumIndexAndUToken.Decimals
		return debt, &profile
	}
	return debt, nil
}

// helper methods
func (eng *DebtEngine) GetPriceForTvl(_ schemas.PriceOracleT, poolPriceOracleT schemas.PriceOracleT, tokenAddr string, amount *big.Int, version core.VersionType) *big.Int {
	priceOracle := poolPriceOracleT
	tokenPrice := eng.priceHandler.GetTokenLastPF(priceOracle, tokenAddr, version)
	tokenDecimals := eng.repo.GetToken(tokenAddr).Decimals
	if version.MoreThan(core.NewVersion(1)) { // than decimals 8
		// log.Info("GetPriceForTvl", tokenPrice, amount, tokenDecimals, tokenAddr, priceOracle)
		return utils.GetInt64(new(big.Int).Mul(tokenPrice.PriceBI.Convert(), amount), tokenDecimals)
	}
	// for v1
	usdcAddr := eng.repo.GetUSD()
	usdcPrice := eng.priceHandler.GetLastPriceFeedByOracle(priceOracle, usdcAddr.Hex(), version)
	// usdcDecimals := eng.repo.GetToken(usdcAddr).Decimals
	var usdcDecimals int8 = 6

	value := new(big.Int).Mul(amount, tokenPrice.PriceBI.Convert())
	value = utils.GetInt64(value, tokenDecimals-usdcDecimals)
	value = new(big.Int).Quo(value, usdcPrice.PriceBI.Convert())
	return new(big.Int).Mul(value, big.NewInt(100))
}
func (eng *DebtEngine) GetAmountInUSD(cm string, tokenAddr string, amount *big.Int, version core.VersionType) *big.Int {
	tokenPrice := eng.priceHandler.GetLastPrice(cm, tokenAddr, version)
	tokenDecimals := eng.repo.GetToken(tokenAddr).Decimals
	if version.MoreThan(core.NewVersion(1)) { // than decimals 8
		return utils.GetInt64(new(big.Int).Mul(tokenPrice, amount), tokenDecimals)
	}
	// for v1
	usdcAddr := eng.repo.GetUSD()
	usdcPrice := eng.priceHandler.GetLastPrice(cm, usdcAddr.Hex(), version)
	// usdcDecimals := eng.repo.GetToken(usdcAddr).Decimals
	var usdcDecimals int8 = 6

	value := new(big.Int).Mul(amount, tokenPrice)
	value = utils.GetInt64(value, tokenDecimals-usdcDecimals)
	value = new(big.Int).Quo(value, usdcPrice)
	return new(big.Int).Mul(value, big.NewInt(100))
}

func (eng *DebtEngine) SessionDataFromDC(version core.VersionType, blockNum int64, cmAddr, borrower, account string) dc.CreditAccountCallData {
	call, resultFn, err := eng.repo.GetDCWrapper().GetCreditAccountData(version, blockNum,
		common.HexToAddress(cmAddr),
		common.HexToAddress(borrower),
		common.HexToAddress(account),
	)
	if err != nil {
		log.Fatalf("Prepaing failed. cm:%s borrower:%s blocknum:%d err:%s", cmAddr, borrower, blockNum, err)
	}
	results := core.MakeMultiCall(eng.client, blockNum, false, []multicall.Multicall2Call{call})
	if !results[0].Success {
		log.Fatalf("Failed multicall. cm:%s borrower:%s blocknum:%d ", cmAddr, borrower, blockNum)
	}
	data, err := resultFn(results[0].ReturnData)
	if err != nil {
		log.Fatalf("cm:%s borrower:%s blocknum:%d err:%s", cmAddr, borrower, blockNum, err)
	}
	return data
}
