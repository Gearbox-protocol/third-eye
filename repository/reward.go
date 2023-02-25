package repository

import (
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/pool_lmrewards"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (repo *Repository) loadLMRewardDetails() {
	defer utils.Elapsed("loadLMRewardDetails")()
	adapterAddrs := repo.GetAdapterAddressByName(ds.PoolLMRewards)
	if len(adapterAddrs) == 0 {
		return
	}
	// load poolLMRewardadapter
	adapterAddr := adapterAddrs[0]
	adapter := repo.GetAdapter(adapterAddr).(*pool_lmrewards.PoolLMRewards)
	// lm rewards
	rewardData := []pool_lmrewards.LMReward{}
	err := repo.db.Raw(`SELECT * FROM lm_rewards`).Find(&rewardData).Error
	if err != nil {
		log.Fatal(err)
	}
	adapter.LoadLMRewards(rewardData)
	//
	dBalanceData := []pool_lmrewards.DieselBalance{}
	err = repo.db.Raw(`SELECT * FROM diesel_balances`).Find(&dBalanceData).Error
	if err != nil {
		log.Fatal(err)
	}
	adapter.LoadDieselBalances(dBalanceData)
}

func (repo Repository) saveLMRewardDetails(tx *gorm.DB) {
	adapterAddr := repo.GetAdapterAddressByName(ds.PoolLMRewards)[0]
	adapter := repo.GetAdapter(adapterAddr).(*pool_lmrewards.PoolLMRewards)
	//
	rewards := adapter.GetLMRewards()
	err := tx.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(rewards, 500).Error
	log.CheckFatal(err)

	//
	balances := adapter.GetDieselBalances()
	err = tx.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(balances, 500).Error
	log.CheckFatal(err)

}
