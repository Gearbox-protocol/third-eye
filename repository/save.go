package repository

import (
	"github.com/Gearbox-protocol/gearscan/log"
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
	// info := tx.Commit()
	// if info.Error != nil {
	// 	log.Fatal(info.Error, *info.Statement)
	// }
	return nil
}

func check(err error) {
	if err!=nil {
		log.Fatal(err)
	}
}
