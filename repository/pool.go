package repository

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/utils"
)

func (repo *Repository) loadPool() {
	defer utils.Elapsed("loadPool")()
	data := []*core.PoolState{}
	err := repo.db.Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, pool := range data {
		adapter := repo.GetAdapter(pool.Address)
		adapter.SetUnderlyingState(pool)
		repo.dieselTokens[pool.DieselToken] = &core.UTokenAndPool{
			Pool:   pool.Address,
			UToken: pool.UnderlyingToken,
		}
	}
}

func (repo *Repository) IsDieselToken(token string) bool {
	return repo.dieselTokens[token] != nil
}

func (repo *Repository) loadPoolUniqueUsers() {
	defer utils.Elapsed("loadPoolUniqueUsers")()
	query := "select distinct pool, user_address from pool_ledger WHERE event = 'AddLiquidity';"
	data := []*core.PoolLedger{}
	err := repo.db.Raw(query).Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range data {
		repo.addPoolUniqueUser(entry.Pool, entry.User)
	}
}

func (repo *Repository) addPoolUniqueUser(pool, user string) {
	if repo.poolUniqueUsers[pool] == nil {
		repo.poolUniqueUsers[pool] = make(map[string]bool)
	}
	repo.poolUniqueUsers[pool][user] = true
}

func (repo *Repository) AddPoolStat(ps *core.PoolStat) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.setAndGetBlock(ps.BlockNum).AddPoolStat(ps)
}

func (repo *Repository) AddPoolLedger(pl *core.PoolLedger) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if "AddLiquidity" == pl.Event {
		repo.addPoolUniqueUser(pl.Pool, pl.User)
	}
	repo.setAndGetBlock(pl.BlockNumber).AddPoolLedger(pl)
}

func (repo *Repository) GetPoolUniqueUserLen(pool string) int {
	return len(repo.poolUniqueUsers[pool])
}

func (repo *Repository) AddDieselToken(dieselToken, underlyingToken, pool string) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.dieselTokens[dieselToken] = &core.UTokenAndPool{
		UToken: underlyingToken,
		Pool:   pool,
	}
	repo.addToken(dieselToken)
}
