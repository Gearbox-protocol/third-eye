package debts

import (
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
		for _, allowedToken := range block.GetAllowedTokens() {
			eng.AddAllowedTokenThreshold(allowedToken)
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
				eng.CalculateSessionDebt(blockNum, sessionId, cmAddr, cmAddrToCumIndex[cmAddr])
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
			cumIndexNormalized := utils.GetInt64Decimal(cumIndex, 27)
			poolToCI[cmAddr] = &core.CumIndexAndUToken{
				CumulativeIndex: cumIndexNormalized,
				Token:           eng.repo.GetCMState(cmAddr).UnderlyingToken,
			}
			// log.Infof("blockNum%d newInterest:%s tsDiff:%s cumIndexDecimal:%s predicate:%s cumIndex:%s",blockNum ,newInterest, tsDiff, cumIndexNormalized, predicate, cumIndex)
		}
	}
	return poolToCI
}

func (eng *DebtEngine) CalculateSessionDebt(blockNum int64, sessionId string, cmAddr string, cumIndexAndUToken *core.CumIndexAndUToken) {
	sessionSnapshot := eng.lastCSS[sessionId]
	calThresholdValue := big.NewInt(0)
	calTotalValue := big.NewInt(0)
	underlyingtoken := eng.repo.GetToken(cumIndexAndUToken.Token)
	// profiling
	tokenDetails := map[string]core.TokenDetails{}
	for tokenAddr, balance := range *sessionSnapshot.Balances {
		decimal := eng.repo.GetToken(tokenAddr).Decimals
		price := eng.GetTokenPrice(tokenAddr)
		tokenValue := new(big.Int).Mul(price, balance.BI.Convert())
		tokenValueInDecimal := utils.GetInt64Decimal(tokenValue, decimal-underlyingtoken.Decimals)
		tokenLiquidityThreshold := eng.allowedTokensThreshold[cmAddr][tokenAddr]
		tokenThresholdValue := new(big.Int).Mul(tokenValueInDecimal, tokenLiquidityThreshold.Convert())
		calThresholdValue = new(big.Int).Add(calThresholdValue, tokenThresholdValue)
		calTotalValue = new(big.Int).Add(calTotalValue, tokenValueInDecimal)
		// profiling
		tokenDetails[tokenAddr] = core.TokenDetails{
			Price:             price,
			Decimals:          decimal,
			TokenLiqThreshold: tokenLiquidityThreshold,
			Symbol:            eng.repo.GetToken(tokenAddr).Symbol}
	}
	// the value of credit account is in terms of underlying asset
	underlyingPrice := eng.GetTokenPrice(cumIndexAndUToken.Token)
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
		CalHealthFactor:                 big.NewInt(0).Quo(calThresholdValue, calBorrowWithInterest).Int64(),
		CalTotalValueBI:                 (*core.BigInt)(calTotalValue),
		CalBorrowedAmountPlusInterestBI: (*core.BigInt)(calBorrowWithInterest),
		CalThresholdValueBI:             (*core.BigInt)(calReducedThresholdValue),
	}
	var notMatched bool
	profile := core.DebtProfile{}
	// use data compressor if debt check is enabled
	if eng.config.DebtDCMatching {
		opts := &bind.CallOpts{
			BlockNumber: big.NewInt(blockNum),
		}
		data, err := eng.repo.GetDCWrapper().GetCreditAccountDataExtended(opts,
			common.HexToAddress(cmAddr),
			common.HexToAddress(sessionSnapshot.Borrower),
		)
		if err != nil {
			log.Fatalf("cm:%s borrower:%s blocknum:%d err:%s", cmAddr, sessionSnapshot.Borrower, blockNum, err)
		}
		// set debt data fetched from dc
		debt.HealthFactor = data.HealthFactor.Int64()
		debt.TotalValueBI = (*core.BigInt)(data.TotalValue)
		debt.BorrowedAmountPlusInterestBI = (*core.BigInt)(data.BorrowedAmountPlusInterest)
		if !core.CompareBalance(debt.CalTotalValueBI, debt.TotalValueBI, underlyingtoken) ||
			!core.CompareBalance(debt.CalBorrowedAmountPlusInterestBI, debt.BorrowedAmountPlusInterestBI, underlyingtoken) {
			profile.RPCBalances = *eng.repo.ConvertToBalance(data.Balances)
			notMatched = true
		}
		// even if data compressor matching is disabled check the values  with values at which last credit snapshot was taken
	} else if sessionSnapshot.BlockNum == blockNum {
		debt.HealthFactor = sessionSnapshot.HealthFactor
		debt.TotalValueBI = sessionSnapshot.TotalValueBI
		if !core.CompareBalance(debt.CalTotalValueBI, debt.TotalValueBI, underlyingtoken) ||
			// if healhFactor diff by 4 %
			utils.IntDiffMoreThanFraction(debt.CalHealthFactor, debt.HealthFactor, 4) {
			profile.RPCBalances = sessionSnapshot.Balances.Copy()
			notMatched = true
		}
	}
	if notMatched {
		profile.CumIndexAndUToken = cumIndexAndUToken
		profile.Debt = debt
		profile.CreditSessionSnapshot = sessionSnapshot
		profile.UnderlyingDecimals = underlyingtoken.Decimals
		profile.Tokens = tokenDetails
		log.Infof("Debt fields different from data compressor fields: %s", profile.Json())
		eng.SaveProfile(string(profile.Json()))
	}
	// check if data compressor and calculated values match
	eng.AddDebt(debt, sessionSnapshot.BlockNum == blockNum)
}

func (eng *DebtEngine) GetTokenPrice(addr string) *big.Int {
	if eng.repo.GetWETHAddr() == addr {
		return core.WETHPrice
	} else {
		return eng.tokenLastPrice[addr].PriceETHBI.Convert()
	}
}
