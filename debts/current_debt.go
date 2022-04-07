package debts

import (
	"context"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
	"gorm.io/gorm/clause"
	"math/big"
)

type BlockTmp struct {
	ID        int64  `gorm:"primaryKey;column:id"`
	Timestamp uint64 `gorm:"column:timestamp"`
}

func (BlockTmp) TableName() string {
	return "blocks"
}

func (eng *DebtEngine) CalCurrentDebts(to int64) {
	// get timestamp of the block
	b, err := eng.client.BlockByNumber(context.Background(), big.NewInt(to))
	if err != nil {
		log.Fatal(err)
	}
	var cmAddrToCumIndex map[string]*ds.CumIndexAndUToken
	if len(eng.lastCSS) > 0 {
		cmAddrToCumIndex = eng.GetCumulativeIndexAndDecimalForCMs(to, b.Time())
	}
	for _, session := range eng.repo.GetSessions() {
		// calculate the current debt for accounts that are only closed after `to` + 1
		// for account closed till `to` + 1 current_debt is calculated while calculating debt
		//
		// for closed accounts the debts are calculated in the debt engine debts/engine.go#L209
		if (session.ClosedAt == 0 || session.ClosedAt > to+1) && session.Since <= to {
			cumIndex := cmAddrToCumIndex[session.CreditManager]
			debt, profile := eng.CalculateSessionDebt(to, session, cumIndex)
			if profile != nil {
				log.Info(profile)
			}
			eng.addCurrentDebt(debt, cumIndex.Decimals)
		}
	}
}
func (eng *DebtEngine) flushCurrentDebts(to int64) {
	// get timestamp of the block
	b, err := eng.client.BlockByNumber(context.Background(), big.NewInt(to))
	if err != nil {
		log.Fatal(err)
	}
	// sync the current debts
	tx := eng.db.Begin()
	if err := tx.Clauses(clause.OnConflict{
		// err := repo.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&BlockTmp{ID: to, Timestamp: b.Time()}).Error; err != nil {
		log.Fatal(err)
	}
	if err := tx.Clauses(clause.OnConflict{
		// err := repo.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(eng.currentDebts, 50).Error; err != nil {
		log.Fatal(err)
	}
	if err := tx.Commit().Error; err != nil {
		log.Fatal(err)
	}
	eng.currentDebts = []*schemas.CurrentDebt{}
}
