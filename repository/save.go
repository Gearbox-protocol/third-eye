package repository

import (
	"gorm.io/gorm/clause"
	"github.com/Gearbox-protocol/gearscan/log"
)

func (repo *Repository) Flush() (err error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	tx := repo.db.Begin()
	for _, adapter:= range repo.syncAdapters{
		tx.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(adapter.GetAdapterState())
	}
	for _, block:= range repo.blocks {
		tx.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(block)
	}
	info:=tx.Commit()
	if info.Error != nil {
		log.Fatal(info.Error,*info.Statement)
	}
	return nil
}