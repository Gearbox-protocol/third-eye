package debts

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"gorm.io/gorm/clause"
)

func (eng *DebtEngine) liquidationCheck(debt *schemas.Debt, cmAddr, borrower string, token *ds.CumIndexAndUToken) {
	lastDebt := eng.lastDebts[debt.SessionId]
	if lastDebt != nil {
		if !core.IntGreaterThanEqualTo(lastDebt.CalHealthFactor, 10000) &&
			core.IntGreaterThanEqualTo(debt.CalHealthFactor, 10000) {
			if eng.liquidableBlockTracker[debt.SessionId] != nil &&
				(debt.BlockNumber-eng.liquidableBlockTracker[debt.SessionId].BlockNum) >= 20 {
				eng.repo.RecentMsgf(log.RiskHeader{
					BlockNumber: debt.BlockNumber,
					EventCode:   "AMQP",
				}, `HealthFactor safe again: 
				SessionId:%s
				HF: %s@(block:%d) -> %s@(block:%d)`,
					debt.SessionId,
					lastDebt.CalHealthFactor, lastDebt.BlockNumber, debt.CalHealthFactor, debt.BlockNumber)
			}
			eng.addLiquidableAccount(debt.SessionId, 0)
			eng.notifiedIfLiquidable(debt.SessionId, false)
		} else if core.IntGreaterThanEqualTo(lastDebt.CalHealthFactor, 10000) &&
			!core.IntGreaterThanEqualTo(debt.CalHealthFactor, 10000) {
			eng.addLiquidableAccount(debt.SessionId, debt.BlockNumber)
		}
	}
	var sendMsgAfterXBlocks int64 = 20
	// sent the account is liquidable notification after 20 blocks
	if !core.IntGreaterThanEqualTo(debt.CalHealthFactor, 10000) &&
		eng.liquidableBlockTracker[debt.SessionId] != nil &&
		(debt.BlockNumber-eng.liquidableBlockTracker[debt.SessionId].BlockNum) >= sendMsgAfterXBlocks &&
		!eng.liquidableBlockTracker[debt.SessionId].NotifiedIfLiquidable {
		urls := core.NetworkUIUrl(core.GetChainId(eng.client))
		eng.notifiedIfLiquidable(debt.SessionId, true)
		eng.repo.RecentMsgf(log.RiskHeader{
			BlockNumber: debt.BlockNumber,
			EventCode:   "INV-LOW-HF-NOT-LIQUIDATED",
		}, `After %d blocks:
				Session: %s
				HF: %s
				CreditManager: %s
				Borrower: %s Debt:%f %s
				web: %s/accounts/%s`,
			sendMsgAfterXBlocks,
			debt.SessionId, debt.CalHealthFactor,
			urls.ExplorerAddrUrl(cmAddr),
			borrower,
			utils.GetFloat64Decimal(debt.CalDebtBI.Convert(), token.Decimals), token.Symbol,
			urls.ChartUrl, debt.SessionId,
		)
	}
}

func (eng *DebtEngine) addCurrentDebt(debt *schemas.Debt, decimals int8) {
	curDebt := schemas.CurrentDebt{
		SessionId: debt.SessionId,
		CommonDebtFields: schemas.CommonDebtFields{
			BlockNumber:     debt.BlockNumber,
			CalHealthFactor: debt.CalHealthFactor,
			CalTotalValueBI: core.NewBigInt(debt.CalTotalValueBI),
			// it has fees for v2
			CalDebtBI:              core.NewBigInt(debt.CalDebtBI),
			CalThresholdValueBI:    core.NewBigInt(debt.CalThresholdValueBI),
			ProfitInUSD:            debt.ProfitInUSD,
			ProfitInUnderlying:     debt.ProfitInUnderlying,
			CollateralInUnderlying: debt.CollateralInUnderlying,
			CollateralInUSD:        debt.CollateralInUSD,
		},
		CalTotalValue:     utils.GetFloat64Decimal(debt.CalTotalValueBI.Convert(), decimals),
		CalDebt:           utils.GetFloat64Decimal((debt.CalDebtBI).Convert(), decimals),
		CalThresholdValue: utils.GetFloat64Decimal((debt.CalThresholdValueBI).Convert(), decimals),
		RepayAmountBI:     debt.RepayAmountBI,
		AmountToPool:      utils.GetFloat64Decimal(debt.AmountToPoolBI.Convert(), decimals),
		RepayAmount:       utils.GetFloat64Decimal(debt.RepayAmountBI.Convert(), decimals),
		//
		TotalValueInUSD: debt.TotalValueInUSD,
		TFIndex:         debt.FarmingValUSD / debt.TotalValueInUSD,
	}
	if curDebt.TFIndex >= 1 {
		curDebt.TFIndex = 1
	}
	eng.currentDebts = append(eng.currentDebts, &curDebt)
}

func (eng *DebtEngine) AddDebt(debt *schemas.Debt, forceAdd bool) {
	lastDebt := eng.lastDebts[debt.SessionId]
	if eng.config.ThrottleDebtCal {
		// add debt if throttle is enabled and (last debt is missing or forced add is set)
		if lastDebt == nil || forceAdd {
			eng.addDebt(debt)
		} else if (debt.BlockNumber-lastDebt.BlockNumber) >= core.NoOfBlocksPerHr*eng.config.ThrottleByHrs ||
			core.DiffMoreThanFraction(lastDebt.CalTotalValueBI, debt.CalTotalValueBI, big.NewFloat(0.05)) ||
			core.DiffMoreThanFraction(lastDebt.CalDebtBI, debt.CalDebtBI, big.NewFloat(0.05)) ||
			// add debt when the health factor is on different side of 10000 from the lastdebt
			core.ValueDifferSideOf10000(debt.CalHealthFactor, lastDebt.CalHealthFactor) {
			eng.addDebt(debt)
		}
	} else {
		eng.addDebt(debt)
	}
}

func (eng *DebtEngine) addDebt(debt *schemas.Debt) {
	eng.addLastDebt(debt)
	eng.debts = append(eng.debts, debt)
}

func (eng *DebtEngine) loadLastDebts(lastDebtSync int64) {
	defer utils.Elapsed("Debt(loadLastDebts)")()
	data := []*schemas.Debt{}
	// from debts
	// query := `SELECT debts.* FROM
	// 		(SELECT max(block_num), session_id FROM debts GROUP BY session_id) debt_max_block
	// 		JOIN debts ON debt_max_block.max = debts.block_num AND debt_max_block.session_id = debts.session_id`
	// from current_debts
	query := `SELECT * FROM current_debts WHERE block_num <= ?;`
	err := eng.db.Raw(query, lastDebtSync).Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, debt := range data {
		eng.addLastDebt(debt)
	}
}

func (eng *DebtEngine) addLastDebt(debt *schemas.Debt) {
	eng.lastDebts[debt.SessionId] = debt
}

func (eng *DebtEngine) flushDebt(newDebtSyncTill int64) {
	debtLen := len(eng.debts)
	if debtLen == 0 {
		return
	}
	log.Infof("Flushing %d till block:%d", debtLen, newDebtSyncTill)
	tx := eng.db.Begin()
	err := tx.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&schemas.DebtSync{LastCalculatedAt: newDebtSyncTill, FieldSet: true}).Error
	log.CheckFatal(err)
	liquidableAccounts := []*schemas.LiquidableAccount{}
	for _, la := range eng.liquidableBlockTracker {
		if la.Updated {
			liquidableAccounts = append(liquidableAccounts, la)
			la.Updated = false
		}
	}
	if len(liquidableAccounts) > 0 {
		err = tx.Clauses(clause.OnConflict{
			// err := repo.db.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).CreateInBatches(liquidableAccounts, 50).Error
		log.CheckFatal(err)
	}
	err = tx.CreateInBatches(eng.debts, 50).Error
	log.CheckFatal(err)
	err = tx.CreateInBatches(eng.tvlSnapshots, 50).Error
	log.CheckFatal(err)
	info := tx.Commit()
	if info.Error != nil {
		log.Fatal(info.Error)
	}
	eng.debts = []*schemas.Debt{}
	eng.tvlSnapshots = []*schemas.TvlSnapshots{}
}
