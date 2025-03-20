package repository

import (
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
)

type _x struct {
	LastCalculatedAt int64 `json:"last_calculated_at"`
}

func (repo *Repository) LoadLastDebtSync() schemas.LastSync {
	data := schemas.LastSync{}
	// query := `SELECT max(b) as last_calculated_at from
	// 	(select min(firstlog_at) as b from sync_Adapters
	// 	WHERE type NOT IN ('RebaseToken','Treasury','LMRewardsv2','LMRewardsv3','GearToken')
	// 	union
	// 	select max(last_calculated_at) as b FROM debt_sync) tmp`
	query := `select * from debt_sync`
	err := repo.db.Raw(query).Find(&data).Error
	log.CheckFatal(err)
	// last debt sync starts from the discover at of address provider to the last debt block stored in debt_sync table
	defaultVal := repo.loadDiscoveredAt()
	if data.Debt == 0 {
		data.Debt = defaultVal
	}
	if data.Tvl == 0 {
		data.Tvl = defaultVal
	}
	return data
}

func (repo *Repository) LoadLastAdapterSync() int64 {
	data := _x{}
	query := "SELECT min(last_sync) as last_calculated_at FROM sync_adapters where disabled=false"
	err := repo.db.Raw(query).Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	return data.LastCalculatedAt
}

func (repo *Repository) loadDiscoveredAt() int64 {
	data := _x{}
	query := "SELECT min(discovered_at) as last_calculated_at FROM sync_adapters WHERE type NOT IN ('RebaseToken','Treasury','LMRewardsv2','LMRewardsv3','GearToken')"
	err := repo.db.Raw(query).Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	return data.LastCalculatedAt
}
