package repository

import (
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

func (repo *Repository) loadPool() {
	defer utils.Elapsed("loadPool")()
	data := []*schemas.PoolState{}
	err := repo.db.Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, pool := range data {
		adapter := repo.GetAdapter(pool.Address)
		adapter.SetUnderlyingState(pool)
		repo.dieselTokens[pool.DieselToken] = &schemas.UTokenAndPool{
			Pool:   pool.Address,
			UToken: pool.UnderlyingToken,
		}
	}
}

func (repo *Repository) IsDieselToken(token string) bool {
	return repo.dieselTokens[token] != nil
}

func (repo *Repository) AddPoolStat(ps *schemas.PoolStat) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.setAndGetBlock(ps.BlockNum).AddPoolStat(ps)
}

func (repo *Repository) AddPoolLedger(pl *schemas.PoolLedger) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if "AddLiquidity" == pl.Event {
		repo.AddPoolUniqueUser(pl.Pool, pl.User)
	}
	repo.setAndGetBlock(pl.BlockNumber).AddPoolLedger(pl)
}

func (repo *Repository) AddDieselToken(dieselToken, underlyingToken, pool string) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.dieselTokens[dieselToken] = &schemas.UTokenAndPool{
		UToken: underlyingToken,
		Pool:   pool,
	}
	repo.addToken(dieselToken)
}
