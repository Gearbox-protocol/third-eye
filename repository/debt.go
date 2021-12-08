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
		poolToCumIndex := repo.GetCumulativeIndexForPools(block.Timestamp)
		// update balance of last credit session snapshot and create credit session snapshots
		sessionsToUpdate := make(map[string]bool)
		for _, css := range block.GetCSS() {
			sessionsToUpdate[css.SessionId] = true
		}
		// set the price session list to update
		sessionWithTokens := make(map[string][]string)
		for _, session := range repo.sessions {
			if session.ClosedAt != 0 && session.ClosedAt < blockNum {
				continue
			}
			sessionSnapshot := repo.lastCSS[session.ID]
			for tokenAddr, _ := range sessionSnapshot.Balances {
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
			if session.ClosedAt != 0 && session.ClosedAt < blockNum {
				return
			}
			cmAddr := session.CreditManager
			repo.CalculateSessionDebt(blockNum, sessionId, cmAddr, poolToCumIndex[cmAddr])

		}
	}
}

func (repo *Repository) GetCumulativeIndexForPools(ts uint64) map[string]*big.Int {
	cmAddrs := repo.kit.GetAdapterAddressByName(core.CreditManager)
	poolToCI := make(map[string]*big.Int)
	for _, cmAddr := range cmAddrs {
		poolAddr := repo.GetState(cmAddr).PoolAddress
		poolInterestData := repo.poolLastInterestData[poolAddr]
		tsDiff := new(big.Int).SetInt64(int64(ts - poolInterestData.Timestamp))
		newInterest := new(big.Int).Mul(poolInterestData.BorrowAPYBI.Convert(), tsDiff)
		predicate := new(big.Int).Add(newInterest, big.NewInt(1))
		poolToCI[poolAddr] = new(big.Int).Mul(poolInterestData.CumulativeIndexRAY.Convert(), predicate)
	}
	return poolToCI
}

func (repo *Repository) CalculateSessionDebt(blockNum int64, sessionId string, cmAddr string, cumIndexNow *big.Int) {
	sessionSnapshot := repo.lastCSS[sessionId]
	calThresholdValue := big.NewInt(0)
	for tokenAddr, balance := range sessionSnapshot.Balances {
		decimal := repo.GetToken(tokenAddr).Decimals
		price := utils.StringToInt(repo.tokenLastPrice[tokenAddr].PriceETHBI)
		tokenValue := new(big.Int).Mul(price, balance.BI.Convert())
		tokenValueInDecimal := utils.GetInt64Decimal(tokenValue, decimal)
		tokenLiquidityThreshold := repo.allowedTokensThreshold[cmAddr][tokenAddr]
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
	repo.blocks[blockNum].AddDebt(&core.Debt{
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

func (repo *Repository) GetState(cmAddr string) *core.CreditManagerState {
	state := repo.kit.GetAdapter(cmAddr).GetUnderlyingState()
	cm, ok := state.(*core.CreditManagerState)
	if !ok {
		log.Fatal("Type assertion for credit manager state failed")
	}
	return cm
}
func (repo *Repository) GetUnderlyingDecimal(cmAddr string) uint8 {
	cm := repo.GetState(cmAddr)
	return repo.GetToken(cm.UnderlyingToken).Decimals
}
