package repository

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"gorm.io/gorm/clause"
)

func (repo *Repository) Flush() (err error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	// preferred order adapter => token => pools => cm => credit session => blocks => allowedTokens

	// credit manager depends on pools
	// allowed token depends on credit managers
	// credit sesion depends on credit manager
	// credit session snapshot on credit session

	// will be depended in future
	// block->pricefeed on token
	// block->protocols on creditManager
	// block->AccountOperation on session
	// block->AllowedTOken on session

	tx := repo.db.Begin()
	for lvlIndex := 0; lvlIndex < repo.kit.Len(); lvlIndex++ {
		for repo.kit.Next(lvlIndex) {
			adapter := repo.kit.Get(lvlIndex)
			err := tx.Clauses(clause.OnConflict{
				// err := repo.db.Clauses(clause.OnConflict{
				UpdateAll: true,
			}).Create(adapter.GetAdapterState()).Error
			log.CheckFatal(err)
			if adapter.HasUnderlyingState() {
				err := tx.Clauses(clause.OnConflict{
					// err := repo.db.Clauses(clause.OnConflict{
					UpdateAll: true,
				}).Create(adapter.GetUnderlyingState()).Error
				log.CheckFatal(err)
			}
		}
		repo.kit.Reset(lvlIndex)
	}
	for _, token := range repo.tokens {
		err := tx.Clauses(clause.OnConflict{
			// err := repo.db.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(token).Error
		log.CheckFatal(err)
	}
	for _, session := range repo.sessions {
		err := tx.Clauses(clause.OnConflict{
			// err := repo.db.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(session).Error
		log.CheckFatal(err)
	}
	for _, block := range repo.blocks {
		err := tx.Clauses(clause.OnConflict{
			// err := repo.db.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(block).Error
		log.CheckFatal(err)
	}

	info := tx.Commit()
	log.CheckFatal(info.Error)
	return nil
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (repo *Repository) flushDebt(newDebtSyncTill int64) {
	debtLen := len(repo.debts)
	if debtLen == 0 {
		return
	}
	log.Infof("Flushing %d for block:%d", debtLen, newDebtSyncTill)
	tx := repo.db.Begin()
	err := tx.Create(core.DebtSync{LastCalculatedAt: newDebtSyncTill}).Error
	log.CheckFatal(err)
	err = tx.Create(repo.debts).Error
	log.CheckFatal(err)
	info := tx.Commit()
	if info.Error != nil {
		log.Fatal(info.Error, *info.Statement)
	}
	repo.debts = []*core.Debt{}
}

func (repo *Repository) clear() {
	for _, session := range repo.sessions {
		if session.ClosedAt != 0 {
			delete(repo.sessions, session.ID)
		}
	}
	repo.blocks = map[int64]*core.Block{}
}
