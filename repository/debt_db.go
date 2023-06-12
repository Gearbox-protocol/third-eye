package repository

import (
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
)

func (repo *Repository) LoadLastDebtSync() int64 {
	data := schemas.DebtSync{}
	query := `SELECT max(b) as last_calculated_at from 
		(select min(firstlog_at) as b from sync_Adapters where type!='PoolLMRewards' 
		union 
		select max(last_calculated_at) as b FROM debt_sync) tmp`
	err := repo.db.Raw(query).Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	// last debt sync starts from the discover at of address provider to the last debt block stored in debt_sync table
	if data.LastCalculatedAt != 0 {
		return data.LastCalculatedAt
	} else {
		return repo.loadDiscoveredAt()
	}
}

func (repo *Repository) LoadLastAdapterSync() int64 {
	data := schemas.DebtSync{}
	query := "SELECT min(last_sync) as last_calculated_at FROM sync_adapters where disabled=false"
	err := repo.db.Raw(query).Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	return data.LastCalculatedAt
}

func (repo *Repository) loadDiscoveredAt() int64 {
	data := schemas.DebtSync{}
	query := "SELECT min(discovered_at) as last_calculated_at FROM sync_adapters"
	err := repo.db.Raw(query).Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	return data.LastCalculatedAt
}
