package repository

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
)

func (repo *Repository) LoadLastDebtSync() int64 {
	data := core.DebtSync{}
	query := "SELECT max(last_calculated_at) as last_calculated_at FROM debt_sync"
	err := repo.db.Raw(query).Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	return data.LastCalculatedAt
}

func (repo *Repository) LoadLastAdapterSync() int64 {
	data := core.DebtSync{}
	query := "SELECT max(last_sync) as last_calculated_at FROM sync_adapters"
	err := repo.db.Raw(query).Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	return data.LastCalculatedAt
}
