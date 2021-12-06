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
		// update balance of last credit session snapshot and create credit session snapshots
		sessionsToUpdate := make(map[string]bool)
		for _, eb := range block.GetEventBalances() {
			repo.UpdateBalance(eb)
			sessionsToUpdate[eb.SessionId] = true
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
			repo.CalculateSessionDebt(blockNum, sessionId)

		}
	}
}

func (repo *Repository) CalculateSessionDebt(blockNum int64, sessionId string) {
	session := repo.sessions[sessionId]
	if session.ClosedAt != 0 && session.ClosedAt < blockNum {
		return
	}
	sessionSnapshot := repo.lastCSS[sessionId]
	cmAddr := session.CreditManager
	calTotalValue := big.NewInt(0)
	for tokenAddr, balance := range sessionSnapshot.Balances {
		decimal := repo.GetToken(tokenAddr).Decimals
		price := utils.StringToInt(repo.tokenLastPrice[tokenAddr].PriceETHBI)
		tokenValue := new(big.Int).Mul(price, balance.BI.Convert())
		tokenValueInDecimal := utils.GetInt64Decimal(tokenValue, decimal)
		tokenLiquidityThreshold := repo.allowedTokensThreshold[cmAddr][tokenAddr]
		tokenThresholdValue := new(big.Int).Mul(tokenValueInDecimal, tokenLiquidityThreshold.Convert())
		calTotalValue = new(big.Int).Add(calTotalValue, tokenThresholdValue)
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
	repo.blocks[blockNum].AddDebt(&core.Debt{
		BlockNumber:                  blockNum,
		SessionId:                    sessionId,
		CalculatedTotalValueBI:       calTotalValue.String(),
		HealthFactor:                 data.HealthFactor.Int64(),
		TotalValueBI:                 data.TotalValue.String(),
		BorrowedAmountPlusInterestBI: data.BorrowedAmountPlusInterest.String(),
	})
}

func (repo *Repository) GetUnderlyingDecimal(cmAddr string) uint8 {
	state := repo.kit.GetAdapter(cmAddr).GetUnderlyingState()
	cm, ok := state.(*core.CreditManagerState)
	if !ok {
		log.Fatal("Type assertion for credit manager state failed")
	}
	return repo.GetToken(cm.UnderlyingToken).Decimals
}

func (repo *Repository) UpdateBalance(eb *core.EventBalance) {
	lastCSS := repo.GetLastCSS(eb.SessionId)
	lastCSS.BlockNum = eb.BlockNumber
	lastCSS.LogId = eb.Index
	lastCSS.Borrower = eb.Borrower
	if !eb.Clear {
		if eb.BorrowedAmount != nil {
			var newBorrowedAmount *big.Int
			if lastCSS.BorrowedAmountBI != nil {
				newBorrowedAmount = (new(big.Int).Add(lastCSS.BorrowedAmountBI.Convert(), eb.BorrowedAmount))
			} else {
				newBorrowedAmount = eb.BorrowedAmount
			}
			lastCSS.BorrowedAmountBI = (*core.BigInt)(newBorrowedAmount)
			lastCSS.BorrowedAmount = utils.GetFloat64Decimal(newBorrowedAmount, repo.GetUnderlyingDecimal(eb.CreditManager))
		}
		oldBalances := lastCSS.Balances
		for tokenAddr, amount := range eb.Transfers {
			tokenBStruct := oldBalances[tokenAddr]
			token := repo.GetToken(tokenAddr)
			if amount.Sign() != 0 {
				if oldBalances[tokenAddr] != nil {
					newAmt := new(big.Int).Add(tokenBStruct.BI.Convert(), amount)
					oldBalances[tokenAddr] = &core.BalanceType{
						BI: (*core.BigInt)(newAmt),
						F:  utils.GetFloat64Decimal(newAmt, token.Decimals),
					}
				} else {
					oldBalances[tokenAddr] = &core.BalanceType{
						BI: (*core.BigInt)(amount),
						F:  utils.GetFloat64Decimal(amount, token.Decimals),
					}
				}
			}
		}
		lastCSS.Balances = oldBalances
	} else {
		if eb.BorrowedAmount == nil {
			lastCSS.BorrowedAmountBI = nil
			lastCSS.BorrowedAmount = 0
		} else {
			lastCSS.BorrowedAmountBI = (*core.BigInt)(eb.BorrowedAmount)
			lastCSS.BorrowedAmount = utils.GetFloat64Decimal(eb.BorrowedAmount, repo.GetUnderlyingDecimal(eb.CreditManager))
		}
		newBalances := core.JsonBalance{}
		for tokenAddr, amount := range eb.Transfers {
			token := repo.GetToken(tokenAddr)
			newBalances[tokenAddr] = &core.BalanceType{
				BI: (*core.BigInt)(amount),
				F:  utils.GetFloat64Decimal(amount, token.Decimals),
			}
		}
		lastCSS.Balances = newBalances
	}

	newCSS := core.CreditSessionSnapshot{}
	newBalances := core.JsonBalance{}
	for tokenAddr, details := range lastCSS.Balances {
		amt := *(details.BI.Convert())
		newBalances[tokenAddr] = &core.BalanceType{
			BI: (*core.BigInt)(&amt),
			F:  details.F,
		}
	}
	newCSS.Balances = newBalances
	newCSS.LogId = lastCSS.LogId
	newCSS.BlockNum = lastCSS.BlockNum
	newCSS.SessionId = lastCSS.SessionId
	if lastCSS.BorrowedAmountBI != nil {
		newBorrowBI := *lastCSS.BorrowedAmountBI
		newCSS.BorrowedAmountBI = &newBorrowBI
	}
	newCSS.BorrowedAmount = lastCSS.BorrowedAmount
	repo.AddCreditSessionSnapshot(&newCSS)
}
