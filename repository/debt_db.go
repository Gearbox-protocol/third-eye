package repository

import (
	"math"

	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
)

func (repo *Repository) loadLastDebtSync() int64 {
	data := core.DebtSync{}
	query := "SELECT max(last_calculated_at) as last_calculated_at FROM debt_sync"
	err := repo.db.Raw(query).Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	return data.LastCalculatedAt
}

func (repo *Repository) loadLastAdapterSync() int64 {
	data := core.DebtSync{}
	query := "SELECT max(last_sync) as last_calculated_at FROM sync_adapters"
	err := repo.db.Raw(query).Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	return data.LastCalculatedAt
}

func (repo *Repository) AddDebt(debt *core.Debt) {
	if repo.config.ThrottleDebtCal {
		lastDebt := repo.lastDebts[debt.SessionId]
		if lastDebt == nil {
			repo.addLastDebt(debt)
			repo.debts = append(repo.debts, debt)
		} else if (debt.BlockNumber-lastDebt.BlockNumber) >= repo.sessionBasedThrottleLimit(debt.SessionId) ||
			// add debt when the health factor is on different side of 10000 from the lastdebt
			(debt.CalHealthFactor >= 10000) != (lastDebt.CalHealthFactor >= 10000) {
			repo.addLastDebt(debt)
			repo.debts = append(repo.debts, debt)
		}
	} else {
		repo.debts = append(repo.debts, debt)
	}
}

func (repo *Repository) loadLastDebts() {
	data := []*core.Debt{}
	query := `SELECT debts.* FROM 
			(SELECT max(block_num), session_id FROM debts GROUP BY session_id) debt_max_block
			JOIN debts ON debt_max_block.max = debts.block_num AND debt_max_block.session_id = debts.session_id`
	err := repo.db.Raw(query).Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, debt := range data {
		repo.addLastDebt(debt)
	}
}

func (repo *Repository) addLastDebt(debt *core.Debt) {
	repo.lastDebts[debt.SessionId] = debt
}

func (repo *Repository) loadThrottleDetails(syncedTill int64) {
	data := []*core.ThrottleDetail{}
	query := `SELECT token, count(block_num), min(block_num) as min_block_num, max(block_num) as current_block_num FROM 
		price_feeds where block_num <= ? group by token`
	err := repo.db.Raw(query, syncedTill).Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, throttleDetails := range data {
		repo.addThrottleDetails(throttleDetails)
	}
}

func (repo *Repository) addThrottleDetails(td *core.ThrottleDetail) {
	repo.tokenThrottleDetails[td.Token] = td
}

func (repo *Repository) addThrottleDetailsFromPriceFeed(pf *core.PriceFeed) {
	td := repo.tokenThrottleDetails[pf.Token]
	if td == nil {
		td = &core.ThrottleDetail{Token: pf.Token, MinBlockNum: pf.BlockNumber}
		td = repo.tokenThrottleDetails[pf.Token]
	}
	td.CurrentBlockNum = pf.BlockNumber
	td.Count++
}

func (repo *Repository) sessionBasedMinUpdatePeriod(sessionId string) int64 {
	css := repo.lastCSS[sessionId]
	var minPriceFeedUpdatePeriod int64 = math.MaxInt64
	for token, balance := range *css.Balances {
		if balance.BI.Convert().Sign() == 0 {
			continue
		}
		td := repo.tokenThrottleDetails[token]
		tokenUpdatePeriod := (td.CurrentBlockNum - td.MinBlockNum) / td.Count
		if minPriceFeedUpdatePeriod > tokenUpdatePeriod {
			minPriceFeedUpdatePeriod = tokenUpdatePeriod
		}
	}
	return minPriceFeedUpdatePeriod
}

func (repo *Repository) sessionBasedThrottleLimit(sessionId string) int64 {
	minPeriod := repo.sessionBasedMinUpdatePeriod(sessionId)
	if minPeriod > 60*core.NoOfBlocksPerMin {
		return minPeriod
	} else {
		return 60 * core.NoOfBlocksPerMin
	}
}
