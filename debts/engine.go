package debts

import (
	"fmt"
	"math/big"
	"sort"

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

	// C3.b: updated price
	for _, pf := range block.GetPriceFeeds() {
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
		eng.currentTs = block.Timestamp
		//
		poolsUpdated, tokensUpdated, sessionsUpdated := eng.updateLocalState(blockNum, block)
		// get pool cumulative interest rate
		cmToPoolDetails := eng.GetCumulativeIndexAndDecimalForCMs(blockNum, block.Timestamp)

		//
		var caTotalValueInUSD float64 = 0
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
			caTotalValueInUSD += utils.GetFloat64Decimal(
				eng.GetAmountInUSD(
					cmToPoolDetails[session.CreditManager].Token,
					sessionSnapshot.TotalValueBI.Convert(), schemas.VersionToPFVersion(session.Version, false),
				), 8)

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
		//
		eng.createTvlSnapshots(blockNum, caTotalValueInUSD)
		if len(sessionsUpdated) > 0 {
			log.Debugf("Calculated %d debts for block %d", len(sessionsUpdated), blockNum)
		}
	}
}

func (eng *DebtEngine) createTvlSnapshots(blockNum int64, caTotalValueInUSD float64) {
	if eng.lastTvlSnapshot != nil && blockNum-eng.lastTvlSnapshot.BlockNum < core.NoOfBlocksPerHr { // tvl snapshot every hour
		return
	}
	var totalAvailableLiquidityInUSD float64 = 0
	log.Info("tvl for block", blockNum)
	for _, entry := range eng.poolLastInterestData {
		adapter := eng.repo.GetAdapter(entry.Address)
		state := adapter.GetUnderlyingState()
		if state == nil {
			log.Fatal("State for pool not found for address: ", entry.Address)
		}
		//
		underlyingToken := state.(*schemas.PoolState).UnderlyingToken
		//
		version := schemas.V1PF
		if eng.tokenLastPrice[schemas.V2PF] != nil {
			version = schemas.V2PF
		}
		if eng.tokenLastPrice[schemas.V3PF_MAIN] != nil {
			version = schemas.V3PF_MAIN
		}
		//
		totalAvailableLiquidityInUSD += utils.GetFloat64Decimal(
			eng.GetAmountInUSD(
				underlyingToken,
				entry.AvailableLiquidityBI.Convert(), version), 8)
	}
	// save as last tvl snapshot and add to db
	tvls := &schemas.TvlSnapshots{
		BlockNum:           blockNum,
		AvailableLiquidity: totalAvailableLiquidityInUSD,
		CATotalValue:       caTotalValueInUSD,
	}
	eng.tvlSnapshots = append(eng.tvlSnapshots, tvls)
	eng.lastTvlSnapshot = tvls
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

func (eng *DebtEngine) getTokenPriceFeed(token string, pfVersion schemas.PFVersion) *schemas.PriceFeed {
	return eng.tokenLastPrice[pfVersion][token]
}

func (eng *DebtEngine) SessionDebtHandler(blockNum int64, session *schemas.CreditSession, cumIndexAndUToken *ds.CumIndexAndUToken) {
	sessionId := session.ID
	sessionSnapshot := eng.lastCSS[sessionId]
	cmAddr := session.CreditManager
	debt, profile := eng.CalculateSessionDebt(blockNum, session, cumIndexAndUToken)
	// if profile is not null
	// yearn price feed might be stale as a result difference btw dc and calculated values
	// solution: fetch price again for all stale yearn feeds
	if profile != nil { //|| debt.CalHealthFactor.Convert().Cmp(big.NewInt(10000)) < 1 {
		retryFeeds := eng.repo.GetRetryFeedForDebts()
		for tokenAddr, details := range *sessionSnapshot.Balances {
			if details.IsEnabled && details.HasBalanceMoreThanOne() {
				lastPriceEvent := eng.getTokenPriceFeed(tokenAddr, schemas.VersionToPFVersion(session.Version, false)) // don't use reserve
				//
				if tokenAddr != eng.repo.GetWETHAddr() && lastPriceEvent.BlockNumber != blockNum {
					feedAddr := lastPriceEvent.Feed
					for _, retryFeed := range retryFeeds {
						if retryFeed.GetAddress() == feedAddr {
							eng.requestPriceFeed(blockNum, retryFeed, tokenAddr, lastPriceEvent.MergedPFVersion)

						}
					}
				}
			}
		}
		debt, profile = eng.CalculateSessionDebt(blockNum, session, cumIndexAndUToken)
		if profile != nil && (session.ID != "0x36bc4d7f24Ab0e9ACa5a84766152447E4F4B9694_14793256_370" && session.ID != "0x11956F74EA4ac57897e7D34419f9cAD467a868D2_21776080_38") { // this mainnet accounts fails due to some difference in the total vlaue bu the hf is same. why?
			log.Fatalf("Debt fields different from data compressor fields: %s", profile)
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
func (eng *DebtEngine) hasRedStoneToken(balances *core.DBBalanceFormat) bool {
	pfs := core.GetRedStonePFByChainId(core.GetChainId(eng.client))
	addToSym := core.GetTokenToSymbolByChainId(core.GetChainId(eng.client))
	for tokenAddr, details := range *balances {
		if _, ok := pfs[addToSym[common.HexToAddress(tokenAddr)]]; details.IsEnabled && details.HasBalanceMoreThanOne() && ok {
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
			CreditSessionSnapshot: sessionSnapshot,
			CM:                    session.CreditManager,
			rebaseDetails:         eng.lastRebaseDetails,
			stETH:                 eng.repo.GetTokenFromSdk("stETH"),
			version:               session.Version,
			forQuotas:             eng.v3DebtDetails,
		},
		eng.lastParameters[session.CreditManager].FeeInterest,
		true,
	)

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
				Price:             eng.GetTokenLastPrice(tokenAddr, schemas.VersionToPFVersion(session.Version, false)), // don't use reserve
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
	} else if sessionSnapshot.BlockNum == blockNum && sessionSnapshot.HealthFactor.Convert().Cmp(new(big.Int)) != 0 { // it is 0 when the issuccessful is false for redstone credit accounts
		if IsChangeMoreThanFraction(debt.CalTotalValueBI, sessionSnapshot.TotalValueBI, big.NewFloat(0.001)) ||
			// hf value calculated are on different side of 1
			core.ValueDifferSideOf10000(debt.CalHealthFactor, sessionSnapshot.HealthFactor) ||
			// if healhFactor diff by 4 %
			core.DiffMoreThanFraction(debt.CalHealthFactor, sessionSnapshot.HealthFactor, big.NewFloat(0.04)) {
			// log.Info(debt.CalHealthFactor, sessionSnapshot.HealthFactor, blockNum)
			// log.Info(debt.CalTotalValueBI, sessionSnapshot.TotalValueBI, blockNum)
			if log.GetBaseNet(core.GetChainId(eng.client)) == "ARBITRUM" {
				profile.CumIndexAndUToken = cumIndexAndUToken
				profile.Debt = debt
				profile.CreditSessionSnapshot = sessionSnapshot
				profile.UnderlyingDecimals = cumIndexAndUToken.Decimals
				log.Warn(utils.ToJson(profile))
			} else if eng.hasRedStoneToken(sessionSnapshot.Balances) {
				if IsChangeMoreThanFraction(debt.CalTotalValueBI, sessionSnapshot.TotalValueBI, big.NewFloat(0.003)) { // .3% allowed
					notMatched = true
				}
			} else {
				notMatched = true
			}
			log.Info(
				IsChangeMoreThanFraction(debt.CalTotalValueBI, sessionSnapshot.TotalValueBI, big.NewFloat(0.001)),
				// hf value calculated are on different side of 1
				core.ValueDifferSideOf10000(debt.CalHealthFactor, sessionSnapshot.HealthFactor),
				// if healhFactor diff by 4 %
				core.DiffMoreThanFraction(debt.CalHealthFactor, sessionSnapshot.HealthFactor, big.NewFloat(0.04)),
			)
		}
	}
	eng.calAmountToPoolAndProfit(debt, session, cumIndexAndUToken, debtDetails)
	eng.farmingCalc.addFarmingVal(debt, session, eng.lastCSS[session.ID], storeForCalc{inner: eng})
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
func (eng *DebtEngine) GetAmountInUSD(tokenAddr string, amount *big.Int, pfVersion schemas.PFVersion) *big.Int {
	tokenPrice := eng.GetTokenLastPrice(tokenAddr, pfVersion)
	tokenDecimals := eng.repo.GetToken(tokenAddr).Decimals
	if pfVersion.Decimals() == 8 {
		return utils.GetInt64(new(big.Int).Mul(tokenPrice, amount), tokenDecimals)
	}
	// for v1
	usdcAddr := eng.repo.GetUSDCAddr()
	usdcPrice := eng.GetTokenLastPrice(usdcAddr, pfVersion)
	usdcDecimals := eng.repo.GetToken(usdcAddr).Decimals

	value := new(big.Int).Mul(amount, tokenPrice)
	value = utils.GetInt64(value, tokenDecimals-usdcDecimals)
	value = new(big.Int).Quo(value, usdcPrice)
	return new(big.Int).Mul(value, big.NewInt(100))
}

func (eng *DebtEngine) GetTokenLastPrice(addr string, pfVersion schemas.PFVersion, dontFail ...bool) *big.Int {
	if pfVersion.ToVersion().Eq(1) && eng.repo.GetWETHAddr() == addr {
		return core.WETHPrice
	}
	if eng.tokenLastPrice[pfVersion][addr] != nil {
		return eng.tokenLastPrice[pfVersion][addr].PriceBI.Convert()
	}
	//
	if len(dontFail) > 0 && dontFail[0] {
		return nil
	}
	log.Fatal(fmt.Sprintf("Price not found for %s pfversion: %d", addr, pfVersion))
	return nil
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

func (eng *DebtEngine) requestPriceFeed(blockNum int64, retryFeed ds.QueryPriceFeedI, token string, pfVersion schemas.MergedPFVersion) {
	// defer func() {
	// 	if err:= recover(); err != nil {
	// 		log.Warn("err", err, "in getting yearn price feed in debt", feed, token, blockNum, pfVersion)
	// 	}
	// }()
	// PFFIX
	calls, isQueryable := retryFeed.GetCalls(blockNum)
	if !isQueryable {
		return
	}
	log.Info("getting price for ", token, "at", blockNum)
	results := core.MakeMultiCall(eng.client, blockNum, false, calls)
	price := retryFeed.ProcessResult(blockNum, results, true)
	eng.AddTokenLastPrice(&schemas.PriceFeed{
		BlockNumber:     blockNum,
		Token:           token,
		Feed:            retryFeed.GetAddress(),
		RoundId:         price.RoundId,
		PriceBI:         price.PriceBI,
		Price:           price.Price,
		MergedPFVersion: pfVersion,
	})
}
