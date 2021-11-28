package repository

import (
	"github.com/Gearbox-protocol/gearscan/log"
	"github.com/Gearbox-protocol/gearscan/core"
	"gorm.io/gorm/clause"
	"gorm.io/gorm"
)

func (repo *Repository) Flush() (err error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	// tx := repo.db.Begin()
	for _, adapter := range repo.syncAdapters {
		// tx.Clauses(clause.OnConflict{
		err :=repo.db.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(adapter.GetAdapterState())
		if err.Error!=nil {
			log.Fatal(err.Error)
		}
	}
	for _, block := range repo.blocks {
		// tx.Clauses(clause.OnConflict{
		err :=repo.db.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(block)
		if err.Error!=nil {
			err = repo.db.Session(&gorm.Session{DryRun: true}).Clauses(clause.OnConflict{
				UpdateAll: true,
			}).Create(block)
			log.Fatal(err.Error, err.Statement.SQL.String(), err.Statement.Vars)
		}
	}
	for _, cm := range repo.creditManagers {
		// tx.Clauses(clause.OnConflict{
		err :=repo.db.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(cm)
		if err.Error!=nil {
			log.Fatal(err.Error)
		}
	}
	for _, token := range repo.tokens {
		// tx.Clauses(clause.OnConflict{
		err :=repo.db.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(token)
		if err.Error!=nil {
			log.Fatal(err.Error)
		}
	}
	for _, pool := range repo.pools {
		// tx.Clauses(clause.OnConflict{
		err :=repo.db.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(pool)
		if err.Error!=nil {
			log.Fatal(err.Error)
		}
	}
	for _, session := range repo.sessions {
		// tx.Clauses(clause.OnConflict{
		err :=repo.db.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(session)
		if err.Error!=nil {
			log.Fatal(err.Error)
		}
	}
	// tx.Clauses(clause.OnConflict{
	if len(repo.allowedTokens) != 0 {
		err1 := repo.db.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(repo.allowedTokens)
		if err1.Error!=nil {
			log.Fatal(err1.Error)
		}
	}

	// info := tx.Commit()
	// if info.Error != nil {
	// 	log.Fatal(info.Error, *info.Statement)
	// }
	repo.allowedTokens = []*core.AllowedToken{}
	repo.blocks = map[int64]*core.Block{}
	return nil
}

func check(err error) {
	if err!=nil {
		log.Fatal(err)
	}
}
