package repository

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"sort"
)

func (repo *Repository) SaveProfile(profile string) {
	err := repo.db.Create(&core.ProfileTable{Profile: profile}).Error
	if err != nil {
		log.Fatal(err)
	}
}

func (repo *Repository) calculateDebt() {
	noOfBlock := len(repo.blocks)
	blockNums := make([]int64, 0, noOfBlock)
	log.Info("######################")
	log.Info("Calculting debts for blocks: %d", noOfBlock)
	log.Info("######################")
	for blockNum := range repo.blocks {
		blockNums = append(blockNums, blockNum)
	}
	sort.Slice(blockNums, func(i, j int) bool { return blockNums[i] < blockNums[j] })
	for _, blockNum := range blockNums {
		block := repo.blocks[blockNum]
		// update threshold
		for _, allowedToken := range block.GetAllowedTokens() {
			repo.AddAllowedTokenThreshold(allowedToken)
		}
		// update pool borrow rate and cumulative index
		for _, ps := range block.GetPoolStats() {
			repo.AddPoolLastInterestData(&core.PoolInterestData{
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
			repo.AddLastCSS(css)
			sessionsToUpdate[css.SessionId] = true
		}
		// set the price session list to update
		sessionWithTokens := make(map[string][]string)
		for _, session := range repo.sessions {
			if (session.ClosedAt != 0 && session.ClosedAt <= blockNum) || session.Since > blockNum {
				continue
			}
			sessionSnapshot := repo.lastCSS[session.ID]
			for tokenAddr, _ := range *sessionSnapshot.Balances {
				sessionWithTokens[tokenAddr] = append(sessionWithTokens[tokenAddr], session.ID)
			}
		}
		// update price
		for _, pf := range block.GetPriceFeeds() {
			repo.AddTokenLastPrice(pf)
			repo.addThrottleDetailsFromPriceFeed(pf)
			// set the price session list to update
			for _, sessionId := range sessionWithTokens[pf.Token] {
				sessionsToUpdate[sessionId] = true
			}
		}
		// get pool cumulative interest rate
		cmAddrToCumIndex := repo.GetCumulativeIndexAndDecimalForCMs(blockNum, block.Timestamp)
		// calculate each session debt
		for sessionId, _ := range sessionsToUpdate {
			session := repo.sessions[sessionId]
			if (session.ClosedAt != 0 && session.ClosedAt <= blockNum) || session.Since > blockNum {
				continue
			}
			cmAddr := session.CreditManager
			// pool cum index is when the pool is not registered
			if cmAddrToCumIndex[cmAddr] != nil {
				repo.CalculateSessionDebt(blockNum, sessionId, cmAddr, cmAddrToCumIndex[cmAddr])
			} else {
				log.Fatalf("CM(%s):pool is missing stats at %d, so cumulative index of pool is unknown", cmAddr, blockNum)
			}
		}
		repo.flushDebt(blockNum)
	}
	// if noOfBlock > 0 {
	// 	repo.flushDebt(blockNums[noOfBlock-1])
	// }
}

func (repo *Repository) GetCumulativeIndexAndDecimalForCMs(blockNum int64, ts uint64) map[string]*core.CumIndexAndUToken {
	cmAddrs := repo.kit.GetAdapterAddressByName(core.CreditManager)
	poolToCI := make(map[string]*core.CumIndexAndUToken)
	for _, cmAddr := range cmAddrs {
		poolAddr := repo.GetCMState(cmAddr).PoolAddress
		poolInterestData := repo.poolLastInterestData[poolAddr]
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
				Token:           repo.GetCMState(cmAddr).UnderlyingToken,
			}
			// log.Infof("blockNum%d newInterest:%s tsDiff:%s cumIndexDecimal:%s predicate:%s cumIndex:%s",blockNum ,newInterest, tsDiff, cumIndexNormalized, predicate, cumIndex)
		}
	}
	return poolToCI
}

func (repo *Repository) CalculateSessionDebt(blockNum int64, sessionId string, cmAddr string, cumIndexAndUToken *core.CumIndexAndUToken) {
	sessionSnapshot := repo.lastCSS[sessionId]
	calThresholdValue := big.NewInt(0)
	calTotalValue := big.NewInt(0)
	profile := core.DebtProfile{}
	underlyingtoken := repo.GetToken(cumIndexAndUToken.Token)
	// profiling
	tokenDetails := map[string]core.TokenDetails{}
	for tokenAddr, balance := range *sessionSnapshot.Balances {
		decimal := repo.GetToken(tokenAddr).Decimals
		price := repo.GetTokenPrice(tokenAddr)
		tokenValue := new(big.Int).Mul(price, balance.BI.Convert())
		tokenValueInDecimal := utils.GetInt64Decimal(tokenValue, decimal-underlyingtoken.Decimals)
		tokenLiquidityThreshold := repo.allowedTokensThreshold[cmAddr][tokenAddr]
		tokenThresholdValue := new(big.Int).Mul(tokenValueInDecimal, tokenLiquidityThreshold.Convert())
		calThresholdValue = new(big.Int).Add(calThresholdValue, tokenThresholdValue)
		calTotalValue = new(big.Int).Add(calTotalValue, tokenValueInDecimal)
		// profiling
		tokenDetails[tokenAddr] = core.TokenDetails{
			Price:             price,
			Decimals:          decimal,
			TokenLiqThreshold: tokenLiquidityThreshold,
			Symbol:            repo.GetToken(tokenAddr).Symbol}
	}
	// the value of credit account is in terms of underlying asset
	underlyingPrice := repo.GetTokenPrice(cumIndexAndUToken.Token)
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
		CalTotalValue:                   calTotalValue.String(),
		CalBorrowedAmountPlusInterestBI: calBorrowWithInterest.String(),
		CalThresholdValueBI:             calReducedThresholdValue.String(),
	}
	// use data compressor if debt check is enabled
	if repo.config.DebtDCMatching {
		opts := &bind.CallOpts{
			BlockNumber: big.NewInt(blockNum),
		}
		data, err := repo.dcWrapper.GetCreditAccountDataExtended(opts,
			common.HexToAddress(cmAddr),
			common.HexToAddress(sessionSnapshot.Borrower),
		)
		if err != nil {
			log.Fatalf("cm:%s borrower:%s blocknum:%d err:%s", cmAddr, sessionSnapshot.Borrower, blockNum, err)
		}
		// set debt data fetched from dc
		debt.HealthFactor = data.HealthFactor.Int64()
		debt.TotalValueBI = data.TotalValue.String()
		debt.BorrowedAmountPlusInterestBI = data.BorrowedAmountPlusInterest.String()
		// check if data compressor and calculated values match
		if !core.CompareBalance(calTotalValue, data.TotalValue, underlyingtoken) ||
			!core.CompareBalance(calBorrowWithInterest, data.BorrowedAmountPlusInterest, underlyingtoken) {
			profile.CumIndexAndUToken = cumIndexAndUToken
			profile.Debt = debt
			profile.CreditSessionSnapshot = sessionSnapshot
			profile.RPCBalances = *repo.ConvertToBalance(data.Balances)
			profile.UnderlyingDecimals = underlyingtoken.Decimals
			profile.Tokens = tokenDetails
			log.Infof("Debt fields different from data compressor fields: %s", profile.Json())
			repo.SaveProfile(string(profile.Json()))
		}
	}
	repo.AddDebt(debt)
}

func (repo *Repository) GetCMState(cmAddr string) *core.CreditManagerState {
	state := repo.kit.GetAdapter(cmAddr).GetUnderlyingState()
	cm, ok := state.(*core.CreditManagerState)
	if !ok {
		log.Fatal("Type assertion for credit manager state failed")
	}
	return cm
}
func (repo *Repository) GetUnderlyingDecimal(cmAddr string) int8 {
	cm := repo.GetCMState(cmAddr)
	return repo.GetToken(cm.UnderlyingToken).Decimals
}

func (repo *Repository) GetTokenPrice(addr string) *big.Int {
	if repo.WETHAddr == addr {
		return core.WETHPrice
	} else {
		return repo.tokenLastPrice[addr].PriceETHBI.Convert()
	}
}
