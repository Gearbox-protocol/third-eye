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
			tx.Clauses(clause.OnConflict{
				// err := repo.db.Clauses(clause.OnConflict{
				UpdateAll: true,
			}).Create(adapter.GetAdapterState())
			// if err.Error != nil {
			// 	log.Fatal(err.Error)
			// }
			if adapter.HasUnderlyingState() {
				tx.Clauses(clause.OnConflict{
					// err := repo.db.Clauses(clause.OnConflict{
					UpdateAll: true,
				}).Create(adapter.GetUnderlyingState())
			}
		}
		repo.kit.Reset(lvlIndex)
	}
	for _, token := range repo.tokens {
		tx.Clauses(clause.OnConflict{
			// err := repo.db.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(token)
		// if err.Error != nil {
		// 	log.Fatal(err.Error)
		// }
	}
	for _, session := range repo.sessions {
		tx.Clauses(clause.OnConflict{
			// err := repo.db.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(session)
		// if err.Error != nil {
		// 	log.Fatal(err.Error)
		// }
	}
	for _, block := range repo.blocks {
		tx.Clauses(clause.OnConflict{
			// err := repo.db.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(block)
		// if err.Error != nil {
		// 	log.Fatal(err.Error)
		// }
	}

	info := tx.Commit()
	if info.Error != nil {
		log.Fatal(info.Error, *info.Statement)
	}
	return nil
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (repo *Repository) flushDebt(newDebtSync int64) {
	err := repo.db.Create(DebtSync{LastCalculatedAt: newDebtSync}).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, session := range repo.sessions {
		if session.ClosedAt != 0 {
			delete(repo.sessions, session.ID)
		}
	}
	repo.blocks = map[int64]*core.Block{}
}
