package repository

import (
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	lmrewardsv3 "github.com/Gearbox-protocol/third-eye/models/pool_lmrewards/v3"
	v3 "github.com/Gearbox-protocol/third-eye/models/pool_lmrewards/v3"
	"gorm.io/gorm"
)

func (repo *Repository) loadLMRewardDetailsv3() {
	defer utils.Elapsed("loadLMRewardDetailsv3")()
	// load poolLMRewardadapter
	adapterAddrs := repo.GetAdapterAddressByName(ds.LMRewardsv3)
	if len(adapterAddrs) == 0 {
		return
	}
	adapterAddr := adapterAddrs[0]
	adapter := repo.GetAdapter(adapterAddr).(*lmrewardsv3.LMRewardsv3)
	//
	farms := []*lmrewardsv3.Farmv3{}
	err := repo.db.Raw(`SELECT * FROM farm_v3`).Find(&farms).Error
	log.CheckFatal(err)
	adapter.SetUnderlyingState(farms)

	//
	dieselsync := []*v3.DieselSync{}
	err = repo.db.Raw(`SELECT * FROM diesel_sync`).Find(&dieselsync).Error
	log.CheckFatal(err)
	adapter.SetUnderlyingState(dieselsync)
	//
	details := []*lmrewardsv3.UserLMDetails{}
	err = repo.db.Raw(`SELECT * FROM user_lmdetails_v3`).Find(&details).Error
	log.CheckFatal(err)
	adapter.SetUnderlyingState(details)
	//
	dBalances := []*ds.DieselBalance{}
	err = repo.db.Raw(`SELECT * FROM diesel_balances where pool in (select address from pools where _version=300)`).Find(&dBalances).Error
	log.CheckFatal(err)
	adapter.SetUnderlyingState(dBalances)
}

func (repo Repository) saveLMRewardDetailsv3(tx *gorm.DB, syncTill int64) {
	adapters := repo.GetAdapterAddressByName(ds.LMRewardsv3)
	if len(adapters) == 0 {
		return
	}
	adapterAddr := adapters[0]
	adapter := repo.GetAdapter(adapterAddr).(*lmrewardsv3.LMRewardsv3)
	//
	currentTs := repo.SetAndGetBlock(syncTill).Timestamp
	adapter.Save(tx, currentTs)

}
