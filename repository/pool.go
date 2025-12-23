package repository

import (
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/core/schemas/schemas_v3"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
)

func (repo *Repository) loadPool() {
	defer utils.Elapsed("loadPool")()
	data := []*schemas.PoolState{}
	err := repo.db.Raw("select * from pools where address in (select address from sync_adapters where type='Pool' and disabled='f')").Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, pool := range data {
		adapter := repo.GetAdapter(pool.Address)
		adapter.SetUnderlyingState(pool)
	}
}
func (repo *Repository) loadDieselToken() {
	defer utils.Elapsed("loadPool")()
	data := []*schemas.PoolState{}
	err := repo.db.Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, pool := range data {
		repo.AddDieselToken(pool.DieselToken, pool.UnderlyingToken, pool.Address, pool.Version)
	}
}

func (repo *Repository) AddPoolLedger(pl *schemas.PoolLedger) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if pl.Event == "AddLiquidity" {
		repo.AddPoolUniqueUser(pl.Pool, pl.User)
	}
	repo.SetAndGetBlock(pl.BlockNumber).AddPoolLedger(pl)
}

func (repo *Repository) loadQuotaDetails() {
	defer utils.Elapsed("loadQuotaDetails")()
	data := []*schemas_v3.QuotaDetails{}
	err := repo.db.Raw("SELECT DISTINCT ON (pool_quota_keeper, token) * FROM quota_details  ORDER BY pool_quota_keeper, token, block_num DESC").Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, pool := range data {
		adapter := repo.GetAdapter(pool.PoolQuotaKeeper)
		adapter.SetUnderlyingState(pool)
	}
}

func (mdl Repository) GetAccountQuotaMgr() *ds.AccountQuotaMgr {
	return mdl.AccountQuotaMgr
}
