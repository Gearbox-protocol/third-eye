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
	repo.debts = append(repo.debts, debt)
}

func (repo *Repository) CalculateDebt() {
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
		// get pool cumulative interest rate
		cmAddrToCumIndex := repo.GetCumulativeIndexForCMs(blockNum, block.Timestamp)
		// update balance of last credit session snapshot and create credit session snapshots
		sessionsToUpdate := make(map[string]bool)
		for _, css := range block.GetCSS() {
			repo.AddLastCSS(css)
			log.Info("add last css", css.SessionId)
			sessionsToUpdate[css.SessionId] = true
		}
		// set the price session list to update
		sessionWithTokens := make(map[string][]string)
		for _, session := range repo.sessions {
			if (session.ClosedAt != 0 && session.ClosedAt <= blockNum) || session.Since > blockNum {
				continue
			}
			sessionSnapshot := repo.lastCSS[session.ID]
			log.Info("session getting last css", session.ID)
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
	}
	blockNumLen := len(blockNums)
	if blockNumLen > 0 {
		repo.flushDebt(blockNums[blockNumLen-1])
	}
}

func (repo *Repository) GetCumulativeIndexForCMs(blockNum int64, ts uint64) map[string]*big.Int {
	cmAddrs := repo.kit.GetAdapterAddressByName(core.CreditManager)
	poolToCI := make(map[string]*big.Int)
	for _, cmAddr := range cmAddrs {
		poolAddr := repo.GetCMState(cmAddr).PoolAddress
		poolInterestData := repo.poolLastInterestData[poolAddr]
		if poolInterestData == nil {
			poolToCI[poolAddr] = nil
		} else {
			tsDiff := new(big.Int).SetInt64(int64(ts - poolInterestData.Timestamp))
			newInterest := new(big.Int).Mul(poolInterestData.BorrowAPYBI.Convert(), tsDiff)
			predicate := new(big.Int).Add(newInterest, big.NewInt(1))
			poolToCI[cmAddr] = new(big.Int).Mul(poolInterestData.CumulativeIndexRAY.Convert(), predicate)
		}
	}
	return poolToCI
}

func (repo *Repository) CalculateSessionDebt(blockNum int64, sessionId string, cmAddr string, cumIndexNow *big.Int) {
	sessionSnapshot := repo.lastCSS[sessionId]
	calThresholdValue := big.NewInt(0)
	for tokenAddr, balance := range *sessionSnapshot.Balances {
		decimal := repo.GetToken(tokenAddr).Decimals
		log.Infof("tokenprice %#v %s",repo.tokenLastPrice, tokenAddr)
		price := utils.StringToInt(repo.tokenLastPrice[tokenAddr].PriceETHBI)
		tokenValue := new(big.Int).Mul(price, balance.BI.Convert())
		tokenValueInDecimal := utils.GetInt64Decimal(tokenValue, decimal)
		tokenLiquidityThreshold := repo.allowedTokensThreshold[cmAddr][tokenAddr]
		// log.Info("TLT %#v %s", repo.allowedTokensThreshold, tokenAddr)
		tokenThresholdValue := new(big.Int).Mul(tokenValueInDecimal, tokenLiquidityThreshold.Convert())
		calThresholdValue = new(big.Int).Add(calThresholdValue, tokenThresholdValue)
	}
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
	}
	data, err := repo.GetDataCompressor(blockNum).GetCreditAccountDataExtended(opts,
		common.HexToAddress(cmAddr),
		common.HexToAddress(sessionSnapshot.Borrower),
	)
	if err != nil {
		log.Fatal(err)
	}
	calBorrowWithInterest := big.NewInt(0).Quo(
		big.NewInt(0).Mul(cumIndexNow, sessionSnapshot.BorrowedAmountBI.Convert()),
		sessionSnapshot.СumulativeIndexAtOpen.Convert())
	repo.AddDebt(&core.Debt{
		BlockNumber:                     blockNum,
		SessionId:                       sessionId,
		CalThresholdValueBI:             calThresholdValue.String(),
		HealthFactor:                    data.HealthFactor.Int64(),
		ThresholdValueBI:                data.TotalValue.String(),
		BorrowedAmountPlusInterestBI:    data.BorrowedAmountPlusInterest.String(),
		CalBorrowedAmountPlusInterestBI: calBorrowWithInterest.String(),
		CalHealthFactor:                 big.NewInt(0).Quo(calThresholdValue, calBorrowWithInterest).Int64(),
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
func (repo *Repository) GetUnderlyingDecimal(cmAddr string) uint8 {
	cm := repo.GetCMState(cmAddr)
	return repo.GetToken(cm.UnderlyingToken).Decimals
}
