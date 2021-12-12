package repository

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
)

func (repo *Repository) loadPool() {
	data := []*core.PoolState{}
	err := repo.db.Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, pool := range data {
		adapter := repo.kit.GetAdapter(pool.Address)
		adapter.SetUnderlyingState(pool)
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
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if repo.poolUniqueUsers[pool] == nil {
		repo.poolUniqueUsers[pool] = make(map[string]bool)
	}
	repo.poolUniqueUsers[pool][user] = true
}

func (repo *Repository) AddPoolStat(ps *core.PoolStat) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.blocks[ps.BlockNum].AddPoolStat(ps)
}

func (repo *Repository) AddPoolLedger(pl *core.PoolLedger) {
	repo.AddPoolUniqueUser(pl.Pool, pl.User)
	repo.blocks[pl.BlockNumber].AddPoolLedger(pl)
}

func (repo *Repository) GetPoolUniqueUserLen(pool string) int {
	return len(repo.poolUniqueUsers[pool])
}

// pool interest state fetch
func (repo *Repository) loadPoolLastInterestData(lastDebtSync int64) {
	data := []*core.PoolInterestData{}
	query := `SELECT * FROM pool_stats 
	JOIN (SELECT max(block_num) as bn, pool FROM pool_stats WHERE block_num <= ? group by pool) as p
	JOIN blocks ON p.bn = blocks.id
	ON p.bn = pool_stats.block_num
	AND p.pool = pool_stats.pool;`
	err := repo.db.Raw(query, lastDebtSync).Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, pd := range data {
		repo.AddPoolLastInterestData(pd)
	}
}

func (repo *Repository) AddPoolLastInterestData(pd *core.PoolInterestData) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.poolLastInterestData[pd.Address] = pd
}
