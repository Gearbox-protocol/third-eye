package debts

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"math/big"
)

func (eng *DebtEngine) AddDebt(debt *core.Debt, forceAdd bool) {
	if eng.config.ThrottleDebtCal {
		lastDebt := eng.lastDebts[debt.SessionId]
		// add debt if throttle is enabled and (last debt is missing or forced add is set)
		if lastDebt == nil || forceAdd {
			eng.addLastDebt(debt)
			eng.debts = append(eng.debts, debt)
		} else if (debt.BlockNumber-lastDebt.BlockNumber) >= core.NoOfBlocksPerHr ||
			core.DiffMoreThanFraction(lastDebt.CalTotalValueBI, debt.CalTotalValueBI, big.NewFloat(0.05)) ||
			core.DiffMoreThanFraction(lastDebt.CalBorrowedAmountPlusInterestBI, debt.CalBorrowedAmountPlusInterestBI, big.NewFloat(0.05)) ||
			// add debt when the health factor is on different side of 10000 from the lastdebt
			(debt.CalHealthFactor >= 10000) != (lastDebt.CalHealthFactor >= 10000) {
			eng.addLastDebt(debt)
			eng.debts = append(eng.debts, debt)
		}
	} else {
		eng.debts = append(eng.debts, debt)
	}
}

func (eng *DebtEngine) loadLastDebts() {
	data := []*core.Debt{}
	query := `SELECT debts.* FROM 
			(SELECT max(block_num), session_id FROM debts GROUP BY session_id) debt_max_block
			JOIN debts ON debt_max_block.max = debts.block_num AND debt_max_block.session_id = debts.session_id`
	err := eng.db.Raw(query).Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, debt := range data {
		eng.addLastDebt(debt)
	}
}

func (eng *DebtEngine) addLastDebt(debt *core.Debt) {
	eng.lastDebts[debt.SessionId] = debt
}

func (eng *DebtEngine) flushDebt(newDebtSyncTill int64) {
	debtLen := len(eng.debts)
	if debtLen == 0 {
		return
	}
	log.Infof("Flushing %d for block:%d", debtLen, newDebtSyncTill)
	tx := eng.db.Begin()
	err := tx.Create(core.DebtSync{LastCalculatedAt: newDebtSyncTill}).Error
	log.CheckFatal(err)
	err = tx.CreateInBatches(eng.debts, 50).Error
	log.CheckFatal(err)
	info := tx.Commit()
	if info.Error != nil {
		log.Fatal(info.Error, *info.Statement)
	}
	eng.debts = []*core.Debt{}
}
