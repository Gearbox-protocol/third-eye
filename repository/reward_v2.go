package repository

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/pool_lmrewards"
	lmrewardsv2 "github.com/Gearbox-protocol/third-eye/models/pool_lmrewards/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (repo *Repository) loadLMRewardDetailsv2() {
	if log.GetBaseNet(core.GetChainId(repo.client)) != "MAINNET" {
		return
	}
	defer utils.Elapsed("loadLMRewardDetailsv2")()
	adapterAddrs := repo.GetAdapterAddressByName(ds.LMRewardsv2)
	if len(adapterAddrs) == 0 {
		return
	}
	// load poolLMRewardadapter
	adapterAddr := adapterAddrs[0]
	adapter := repo.GetAdapter(adapterAddr).(*lmrewardsv2.LMRewardsv2)
	// lm rewards
	rewardData := []pool_lmrewards.LMReward{}
	err := repo.db.Raw(`SELECT * FROM lm_rewards where pool in (SELECT address FROM pools where _version!=300)`).Find(&rewardData).Error
	if err != nil {
		log.Fatal(err)
	}
	adapter.LoadLMRewards(rewardData)
	//
	dBalanceData := []lmrewardsv2.DieselBalance{}
	err = repo.db.Raw(`SELECT * FROM diesel_balances`).Find(&dBalanceData).Error
	if err != nil {
		log.Fatal(err)
	}
	adapter.LoadDieselBalances(dBalanceData)
}

func (repo Repository) saveLMRewardDetailsv2(tx *gorm.DB) {
	if log.GetBaseNet(core.GetChainId(repo.client)) != "MAINNET" {
		return
	}
	adapterAddr := repo.GetAdapterAddressByName(ds.LMRewardsv2)[0]
	adapter := repo.GetAdapter(adapterAddr).(*lmrewardsv2.LMRewardsv2)
	//

	if rewards := adapter.GetLMRewards(); len(rewards) != 0 {
		err := tx.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).CreateInBatches(rewards, 500).Error
		log.CheckFatal(err)
	}

	//

	if balances := adapter.GetDieselBalances(); len(balances) != 0 {
		err := tx.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).CreateInBatches(balances, 500).Error
		log.CheckFatal(err)
	}

	adapter.SyncComplete()
}
