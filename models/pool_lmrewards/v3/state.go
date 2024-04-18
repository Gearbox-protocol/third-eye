package v3

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"

	"github.com/Gearbox-protocol/third-eye/models/pool_lmrewards"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (mdl *LMRewardsv3) SetUnderlyingState(obj interface{}) {
	switch ans := obj.(type) {
	case []*Farmv3:
		farms := map[string]*Farmv3{}
		pools := map[common.Address]string{}
		for _, farm := range ans {
			farms[farm.Farm] = farm
			pools[common.HexToAddress(farm.Pool)] = farm.Farm
		}
		mdl.farms = farms
		mdl.pools = pools
	case []*UserLMDetails:
		users := map[common.Address]map[string]*UserLMDetails{}
		for _, user := range ans {
			if users[common.HexToAddress(user.Farm)] == nil {
				users[common.HexToAddress(user.Farm)] = map[string]*UserLMDetails{}
			}
			users[common.HexToAddress(user.Farm)][user.Account] = user
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
	for _, farmAndItsUsers := range mdl.users {
		for _, entry := range farmAndItsUsers {
			if entry.updated {
				users = append(users, entry)
				entry.updated = false
			}
		}
	}
	err = tx.Clauses(clause.OnConflict{UpdateAll: true}).CreateInBatches(users, 500).Error
	log.CheckFatal(err)

	//
	rewards := []*pool_lmrewards.LMReward{}
	for _, farmAndItsUsers := range mdl.users {
		for _, user := range farmAndItsUsers {
			farm := mdl.farms[user.Farm]
			reward := user.GetPoints(farm, currentTs)
			rewards = append(rewards, &pool_lmrewards.LMReward{
				User:   user.Account,
				Pool:   farm.Pool,
				Farm:   farm.Farm,
				Reward: (*core.BigInt)(reward),
			})
		}
	}
	err = tx.Clauses(clause.OnConflict{UpdateAll: true}).CreateInBatches(rewards, 500).Error
	for i, a := range rewards[:utils.Max(len(rewards)-1, 0)] {
		for _, b := range rewards[i+1:] {
			if a.User == b.User && a.Pool == b.Pool {
				log.Fatalf("Duplicate entries for %s %s %s %s", a.User, a.Pool, a.Farm, b.Farm)
			}
		}
	}
	log.CheckFatal(err)
}
