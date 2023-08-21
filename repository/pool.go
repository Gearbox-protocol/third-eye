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
	err := repo.db.Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, pool := range data {
		adapter := repo.GetAdapter(pool.Address)
		adapter.SetUnderlyingState(pool)
	}
}

func (repo *Repository) AddPoolLedger(pl *schemas.PoolLedger) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if "AddLiquidity" == pl.Event {
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

func (repo *Repository) loadAccountQuotaInfo() {
	defer utils.Elapsed("loadAccountQuotaInfo")()
	data := []*schemas_v3.AccountQuotaInfo{}
	err := repo.db.Raw(`WITH all_data as (SELECT DISTINCT ON (cs.account, token) aqi.* 
	FROM account_quota_info aqi JOIN credit_sessions cs 
	ON cs.id = aqi.session_id ORDER BY cs.account, token, block_num DESC)
	SELECT * from all_data WHERE quota::float> 1`).Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	repo.AccountQuotaMgr.InitQuotas(data)
}

func (mdl Repository) GetAccountQuotaMgr() *ds.AccountQuotaMgr {
	return mdl.AccountQuotaMgr
}
