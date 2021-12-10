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

type DebtSync struct {
	LastCalculatedAt int64 `gorm:"last_calculated_at"`
}

func (DebtSync) TableName() string {
	return "debt_sync"
}

func (repo *Repository) loadLastDebtSync() int64 {
	data := DebtSync{}
	query := "SELECT max(last_calculated_at) as last_calculated_at FROM debt_sync"
	err := repo.db.Raw(query).Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	return data.LastCalculatedAt
}

func (repo *Repository) AddDebt(debt *core.Debt) {
	log.Infof("Debt %#v\n", debt)
	repo.debts = append(repo.debts, debt)
}

func (repo *Repository) calculateDebt() {
	blockNums := make([]int64, 0, len(repo.blocks))
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
		// repo.flushDebt(blockNum)
	}
	blockNumLen := len(blockNums)
	if blockNumLen > 0 {
		repo.flushDebt(blockNums[blockNumLen-1])
	}
}

type CumIndexAndUToken struct {
	CumulativeIndex *big.Int
	Token           string
}

func (repo *Repository) GetCumulativeIndexAndDecimalForCMs(blockNum int64, ts uint64) map[string]*CumIndexAndUToken {
	cmAddrs := repo.kit.GetAdapterAddressByName(core.CreditManager)
	poolToCI := make(map[string]*CumIndexAndUToken)
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
			poolToCI[cmAddr] = &CumIndexAndUToken{
				CumulativeIndex: cumIndexNormalized,
				Token:           repo.GetCMState(cmAddr).UnderlyingToken,
			}
			// log.Infof("blockNum%d newInterest:%s tsDiff:%s cumIndexDecimal:%s predicate:%s cumIndex:%s",blockNum ,newInterest, tsDiff, cumIndexNormalized, predicate, cumIndex)
		}
	}
	return poolToCI
}

func (repo *Repository) CalculateSessionDebt(blockNum int64, sessionId string, cmAddr string, cumIndexAndUToken *CumIndexAndUToken) {
	sessionSnapshot := repo.lastCSS[sessionId]
	calThresholdValue := big.NewInt(0)
	calTotalValue := big.NewInt(0)
	underlyingDecimals := repo.GetToken(cumIndexAndUToken.Token).Decimals
	for tokenAddr, balance := range *sessionSnapshot.Balances {
		decimal := repo.GetToken(tokenAddr).Decimals
		price := repo.GetTokenPrice(tokenAddr)
		tokenValue := new(big.Int).Mul(price, balance.BI.Convert())
		tokenValueInDecimal := utils.GetInt64Decimal(tokenValue, decimal-underlyingDecimals)
		tokenLiquidityThreshold := repo.allowedTokensThreshold[cmAddr][tokenAddr]
		tokenThresholdValue := new(big.Int).Mul(tokenValueInDecimal, tokenLiquidityThreshold.Convert())
		calThresholdValue = new(big.Int).Add(calThresholdValue, tokenThresholdValue)
		calTotalValue = new(big.Int).Add(calTotalValue, tokenValueInDecimal)
	}
	// the value of credit account is in terms of underlying asset
	underlyingPrice := repo.GetTokenPrice(cumIndexAndUToken.Token)
	calThresholdValue = new(big.Int).Quo(calThresholdValue, underlyingPrice)
	calTotalValue = new(big.Int).Quo(calTotalValue, underlyingPrice)

	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
	}
	data, err := repo.GetDataCompressor(blockNum).GetCreditAccountDataExtended(opts,
		common.HexToAddress(cmAddr),
		common.HexToAddress(sessionSnapshot.Borrower),
	)
	if err != nil {
		log.Fatalf("cm:%s borrower:%s blocknum:%d err:%s", cmAddr, sessionSnapshot.Borrower, blockNum, err)
	}
	calBorrowWithInterest := big.NewInt(0).Quo(
		big.NewInt(0).Mul(cumIndexAndUToken.CumulativeIndex, sessionSnapshot.BorrowedAmountBI.Convert()),
		sessionSnapshot.СumulativeIndexAtOpen.Convert())
	calReducedThresholdValue := big.NewInt(0).Quo(calThresholdValue, big.NewInt(10000))
	if !utils.AlmostSameBigInt(calTotalValue, data.TotalValue, underlyingDecimals) || !utils.AlmostSameBigInt(calBorrowWithInterest, data.BorrowedAmountPlusInterest, underlyingDecimals) {
		log.Fatal("calTotalValue:%s totalValue:%s calBorrowWithInterest:%s borrowWithInterest:%s",
			calTotalValue, data.TotalValue, calBorrowWithInterest, data.BorrowedAmountPlusInterest)
	}
	repo.AddDebt(&core.Debt{
		BlockNumber:                     blockNum,
		SessionId:                       sessionId,
		HealthFactor:                    data.HealthFactor.Int64(),
		TotalValueBI:                    data.TotalValue.String(),
		BorrowedAmountPlusInterestBI:    data.BorrowedAmountPlusInterest.String(),
		CalHealthFactor:                 big.NewInt(0).Quo(calThresholdValue, calBorrowWithInterest).Int64(),
		CalTotalValue:                   calTotalValue.String(),
		CalBorrowedAmountPlusInterestBI: calBorrowWithInterest.String(),
		CalThresholdValueBI:             calReducedThresholdValue.String(),
	})
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
	if repo.config.WETHAddr == addr {
		return core.WETHPrice
	} else {
		return repo.tokenLastPrice[addr].PriceETHBI.Convert()
	}
}
