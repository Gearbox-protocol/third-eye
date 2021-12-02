package repository

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
)

func (repo *Repository) loadPool() {
	data := []*core.Pool{}
	err := repo.db.Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, pool := range data {
		adapter := repo.syncAdapters[pool.Address]
		adapter.SetState(pool)
	}
}

func (repo *Repository) loadPoolUniqueUsers() {
	query := "select distinct pool,address from pool_ledger;"
	data := []*core.PoolLedger{}
	err := repo.db.Raw(query).Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range data {
		repo.AddPoolUniqueUser(entry.Pool, entry.User)
	}
}

func (repo *Repository) AddPoolUniqueUser(pool, user string) {
	if repo.poolUniqueUsers[pool] == nil {
		repo.poolUniqueUsers[pool] = make(map[string]bool)
	}
	repo.poolUniqueUsers[pool][user] = true
}

func (repo *Repository) AddPoolStat(ps *core.PoolStat) {
	repo.blocks[ps.BlockNum].AddPoolStat(ps)
}

func (repo *Repository) AddPoolLedger(pl *core.PoolLedger) {
	repo.AddPoolUniqueUser(pl.Pool, pl.User)
	repo.blocks[pl.BlockNumber].AddPoolLedger(pl)
}

func (repo *Repository) GetPoolUniqueUserLen(pool string) int {
	if repo.poolUniqueUsers[pool] == nil {
		log.Fatal("pool unique user map is nil")
	}
	return len(repo.poolUniqueUsers[pool])
}
