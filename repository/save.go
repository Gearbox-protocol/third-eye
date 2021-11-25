package repository

import (
	"github.com/Gearbox-protocol/gearscan/log"
	"gorm.io/gorm/clause"
)

func (repo *Repository) Flush() (err error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	tx := repo.db.Begin()
	for _, adapter := range repo.syncAdapters {
		tx.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(adapter.GetAdapterState())
	}
	for _, block := range repo.blocks {
		tx.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(block)
	}
	for _, cm := range repo.creditManagers {
		tx.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(cm)
	}
	info := tx.Commit()
	if info.Error != nil {
		log.Fatal(info.Error, *info.Statement)
	}
	return nil
}
