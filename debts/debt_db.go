package debts

import (
	"math/big"

	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/utils"
	"gorm.io/gorm/clause"
)

type NetworkUI struct {
	ExplorerUrl string
	ChartUrl    string
}

func (eng *DebtEngine) networkUIUrl() NetworkUI {
	switch eng.config.ChainId {
	case 1:
		return NetworkUI{
			ExplorerUrl: "https://etherscan.io",
			ChartUrl:    "https://charts.gearbox.fi",
		}
	case 42:
		return NetworkUI{
			ExplorerUrl: "https://kovan.etherscan.io",
			ChartUrl:    "https://charts.kovan.gearbox.fi",
		}
	}
	return NetworkUI{}
}
func (eng *DebtEngine) liquidationCheck(debt *core.Debt, cmAddr, borrower string, token *core.CumIndexAndUToken) {
	lastDebt := eng.lastDebts[debt.SessionId]
	if lastDebt != nil {
		if !core.IntGreaterThanEqualTo(lastDebt.CalHealthFactor, 10000) &&
			core.IntGreaterThanEqualTo(debt.CalHealthFactor, 10000) {
			eng.addLiquidableAccount(debt.SessionId, 0)
			log.Msgf(`HealthFactor safe again: 
			SessionId:%s
			HF: %s@(block:%d) -> %s@(block:%d)`,
				debt.SessionId,
				lastDebt.CalHealthFactor, lastDebt.BlockNumber, debt.CalHealthFactor, debt.BlockNumber)

		} else if core.IntGreaterThanEqualTo(lastDebt.CalHealthFactor, 10000) &&
			!core.IntGreaterThanEqualTo(debt.CalHealthFactor, 10000) {
			eng.addLiquidableAccount(debt.SessionId, debt.BlockNumber)
			urls := eng.networkUIUrl()
			log.Msgf(`HealthFactor low:
				Session: %s
				HF: %s -> %s
				CreditManager: %s/address/%s
				Borrower: %s RepayAmount:%f %s
				web: %s/accounts/%s/%s`,
				debt.SessionId,
				lastDebt.CalHealthFactor, debt.CalHealthFactor,
				urls.ExplorerUrl, cmAddr,
				borrower,
				utils.GetFloat64Decimal(debt.CalBorrowedAmountPlusInterestBI.Convert(), token.Decimals), token.Symbol,
				urls.ChartUrl, cmAddr, borrower,
			)
		}
	}
}

func (eng *DebtEngine) addCurrentDebt(debt *core.Debt, decimals int8) {
	curDebt := core.CurrentDebt{
		SessionId:                       debt.SessionId,
		BlockNumber:                     debt.BlockNumber,
		CalHealthFactor:                 debt.CalHealthFactor,
		CalTotalValue:                   utils.GetFloat64Decimal(debt.CalTotalValueBI.Convert(), decimals),
		CalTotalValueBI:                 core.NewBigInt(debt.CalTotalValueBI),
		CalBorrowedAmountPlusInterest:   utils.GetFloat64Decimal((debt.CalBorrowedAmountPlusInterestBI).Convert(), decimals),
		CalBorrowedAmountPlusInterestBI: core.NewBigInt(debt.CalBorrowedAmountPlusInterestBI),
		CalThresholdValue:               utils.GetFloat64Decimal((debt.CalThresholdValueBI).Convert(), decimals),
		CalThresholdValueBI:             core.NewBigInt(debt.CalThresholdValueBI),
		AmountToPoolBI:                  debt.AmountToPoolBI,
		AmountToPool:                    utils.GetFloat64Decimal(debt.AmountToPoolBI.Convert(), decimals),
		ProfitInUSDBI:                   debt.ProfitInUSDBI,
		CollateralInUSDBI:               debt.CollateralInUSDBI,
	}
	eng.currentDebts = append(eng.currentDebts, &curDebt)
}

func (eng *DebtEngine) AddDebt(debt *core.Debt, forceAdd bool) {
	lastDebt := eng.lastDebts[debt.SessionId]
	if eng.config.ThrottleDebtCal {
		// add debt if throttle is enabled and (last debt is missing or forced add is set)
		if lastDebt == nil || forceAdd {
			eng.addDebt(debt)
		} else if (debt.BlockNumber-lastDebt.BlockNumber) >= core.NoOfBlocksPerHr ||
			core.DiffMoreThanFraction(lastDebt.CalTotalValueBI, debt.CalTotalValueBI, big.NewFloat(0.05)) ||
			core.DiffMoreThanFraction(lastDebt.CalBorrowedAmountPlusInterestBI, debt.CalBorrowedAmountPlusInterestBI, big.NewFloat(0.05)) ||
			// add debt when the health factor is on different side of 10000 from the lastdebt
			core.ValueDifferSideOf10000(debt.CalHealthFactor, lastDebt.CalHealthFactor) {
			eng.addDebt(debt)
		}
	} else {
		eng.addDebt(debt)
	}
}

func (eng *DebtEngine) addDebt(debt *core.Debt) {
	eng.addLastDebt(debt)
	eng.debts = append(eng.debts, debt)
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
	err := tx.Create(&core.DebtSync{LastCalculatedAt: newDebtSyncTill}).Error
	log.CheckFatal(err)
	liquidableAccounts := []*core.LiquidableAccount{}
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
	info := tx.Commit()
	if info.Error != nil {
		log.Fatal(info.Error, *info.Statement)
	}
	eng.debts = []*core.Debt{}
}
