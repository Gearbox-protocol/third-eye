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

	tx := repo.db.Begin()
	for _, adapter := range repo.syncAdapters {
		tx.Clauses(clause.OnConflict{
			// err := repo.db.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(adapter.GetState())
		// if err.Error != nil {
		// 	log.Fatal(err.Error)
		// }
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
	for _, pool := range repo.pools {
		tx.Clauses(clause.OnConflict{
			// err := repo.db.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(pool)
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
	if len(repo.allowedTokens) != 0 {
		tx.Clauses(clause.OnConflict{
			// err1 := repo.db.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(repo.allowedTokens)
		// if err1.Error != nil {
		// 	log.Fatal(err1.Error)
		// }
	}

	info := tx.Commit()
	if info.Error != nil {
		log.Fatal(info.Error, *info.Statement)
	}
	repo.allowedTokens = []*core.AllowedToken{}
	repo.blocks = map[int64]*core.Block{}
	return nil
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
