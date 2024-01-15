package v3

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/models/pool_lmrewards"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (mdl *LMRewardsv3) SetUnderlyingState(obj interface{}) {
	switch ans := obj.(type) {
	case []*Farmv3:
		farms := map[string]*Farmv3{}
		for _, farm := range ans {
			farms[farm.Farm] = farm
		}
		mdl.farms = farms
	case []*UserLMDetails:
		users := map[string]*UserLMDetails{}
		for _, user := range ans {
			users[user.Account] = user
		}
		mdl.users = users
	default:
		log.Fatalf("Not able to parse underlying state for %T", obj)
	}
}

func (mdl *LMRewardsv3) Save(tx *gorm.DB, currentTs uint64) {
	farms := []*Farmv3{}
	for _, entry := range mdl.farms {
		farms = append(farms, entry)
	}
	err := tx.Clauses(clause.OnConflict{UpdateAll: true}).CreateInBatches(farms, 500).Error
	log.CheckFatal(err)

	//
	users := []*UserLMDetails{}
	for _, entry := range mdl.users {
		if entry.updated {
			users = append(users, entry)
			entry.updated = false
		}
	}
	err = tx.Clauses(clause.OnConflict{UpdateAll: true}).CreateInBatches(users, 500).Error
	log.CheckFatal(err)

	//
	rewards := []*pool_lmrewards.LMReward{}
	for _, user := range mdl.users {
		farm := mdl.farms[user.Farm]
		reward := user.GetPoints(farm, currentTs)
		rewards = append(rewards, &pool_lmrewards.LMReward{
			User:   user.Account,
			Pool:   user.Farm,
			Reward: (*core.BigInt)(reward),
		})
	}
	err = tx.Clauses(clause.OnConflict{UpdateAll: true}).CreateInBatches(rewards, 500).Error
	log.CheckFatal(err)
}
