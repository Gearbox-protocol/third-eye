package v3

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"

	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/ds/dc_wrapper"
	"github.com/Gearbox-protocol/third-eye/models/pool_lmrewards"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// lastsync can  only be set in the init method.
func (mdl *LMRewardsv3) getFarmsAndPoolsv3(blockNum int64) {
	pools := mdl.Repo.GetDCWrapper().GetZapperInfo(blockNum)
	mdl.SetFarm(pools)
	for _, pool := range mdl.Repo.GetAdapterAddressByName(ds.Pool) {
		discoveredAt := mdl.Repo.GetAdapter(pool).GetDiscoveredAt()
		mdl.setPoolSyncedTill(common.HexToAddress(pool), discoveredAt)
		mdl.LastSync = utils.Min(mdl.LastSync, mdl.poolsToSyncedTill[common.HexToAddress(pool)])
	}
}

func (mdl *LMRewardsv3) SetFarm(pools []dc_wrapper.PoolZapperInfo) {
	// farmingPools := core.GetFarmingPoolsToSymbolByChainId(core.GetChainId(mdl.Client))
	for _, pool := range pools {
		newfarmsForPool := []common.Address{}
		oldfarmsForPool := []common.Address{}
		for _, zapper := range pool.Zappers {
			if zapper.TokenOut.Hex() == "0x580e39ADb33E106fFc2712cBD57B9cE954dcfE75" { // GHO
				zapper.TokenOut = common.HexToAddress("0xE2037090f896A858E3168B978668F22026AC52e7")
			}
			if zapper.TokenOut.Hex() == "0x7aB44F17EE21A3D6Bb2aeb1c6cA8B875041608C4" { // DAI
				zapper.TokenOut = common.HexToAddress("0xC853E4DA38d9Bd1d01675355b8c8f3BBC1451973")
			}
			// can be diselToken zapperOut -- https://etherscan.io/address/0xcaa199f91294e6ee95f9ea90fe716cbd2f9f2900#code
			if zapper.TokenIn == pool.Underlying && zapper.TokenOut != pool.DieselToken {
				_, err := core.CallFuncGetSingleValue(mdl.Client, "bfe10928", zapper.TokenOut, 0, nil) // distributor on the farm
				if err != nil {
					continue
				}
				farm := &Farmv3{
					Farm:        zapper.TokenOut.Hex(),
					Pool:        pool.Addr.Hex(),
					DieselToken: pool.DieselToken.Hex(),
					// initial
					Fpt:            (*core.BigInt)(new(big.Int)),
					TotalSupply:    (*core.BigInt)(new(big.Int)),
					Reward:         (*core.BigInt)(new(big.Int)),
					FarmSyncedTill: mdl.Repo.GetAdapter(pool.Addr.Hex()).GetDiscoveredAt(),
				}
				if mdl.farms[farm.Farm] == nil {
					farm.setRewardToken(mdl.Client)
					mdl.farms[farm.Farm] = farm
					// poolAndFarms = append(poolAndFarms, farm)
					newfarmsForPool = append(newfarmsForPool, common.HexToAddress(farm.Farm))
					mdl.LastSync = utils.Min(mdl.LastSync, farm.FarmSyncedTill)
				} else {
					farm = mdl.farms[farm.Farm]
					oldfarmsForPool = append(oldfarmsForPool, common.HexToAddress(farm.Farm))
				}
				// poolsyncedTill is zero  then rejected. in the case of new pool.
				// poolSyncedTill is zero in db and then it is updated to discoveredat of pool. and used to calculate min poolsyncedTill
			}
		}
		if len(newfarmsForPool) != 0 && len(oldfarmsForPool) != 0 {
			log.Warn("pool tracking for disel balance will be affected as new farms are added", pool.Addr)
		}
		if len(newfarmsForPool) != 0 {
			log.Warnf("farms for pool(%s): new:%v: old: %v", pool.Addr, newfarmsForPool, oldfarmsForPool)
		}
	}
	// mdl.SetUnderlyingState(poolAndFarms)
}

// from contractRegister, Is that opposed to saying "Tell will be sad to discover that"?
// so use that blocknumber
func (mdl *LMRewardsv3) AddPoolv3(discoveredAt int64, pool string) {
	data := mdl.Repo.GetDCWrapper().GetZapperInfo(discoveredAt, common.HexToAddress(pool))
	mdl.SetFarm(data)
	mdl.setPoolSyncedTill(common.HexToAddress(pool), discoveredAt)
}
func (mdl *LMRewardsv3) setPoolSyncedTill(pool common.Address, discoveredAt int64) {
	if mdl.poolsToSyncedTill[pool] < discoveredAt {
		mdl.poolsToSyncedTill[pool] = discoveredAt
	}
}

func (mdl *LMRewardsv3) SetUnderlyingState(obj interface{}) {
	switch ans := obj.(type) {
	case []*Farmv3:
		if mdl.farms == nil {
			mdl.farms = map[string]*Farmv3{}
		}
		for _, farm := range ans {
			if mdl.farms[farm.Farm] == nil {
				farm.setRewardToken(mdl.Client)
				mdl.farms[farm.Farm] = farm
			}
		}
	case []*DieselSync:
		for _, sync := range ans {
			mdl.poolsToSyncedTill[common.HexToAddress(sync.Pool)] = sync.PoolSyncedTill
		}
	case []*UserLMDetails:
		users := map[common.Address]map[string]*UserLMDetails{}
		for _, user := range ans {
			if users[common.HexToAddress(user.Farm)] == nil {
				users[common.HexToAddress(user.Farm)] = map[string]*UserLMDetails{}
			}
			users[common.HexToAddress(user.Farm)][user.Account] = user
		}
		mdl.farmUserRewards = users
	case []*ds.DieselBalance:
		users := map[common.Address]map[string]*ds.DieselBalance{}
		for _, user := range ans {
			pool := common.HexToAddress(user.Pool)
			if users[pool] == nil {
				users[pool] = map[string]*ds.DieselBalance{}
			}
			users[pool][user.User] = user
		}
		mdl.dieselBalances = users
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
	for _, farmAndItsUsers := range mdl.farmUserRewards {
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
	dieselBalance := []*ds.DieselBalance{}
	for _, userAndBalances := range mdl.dieselBalances {
		for _, entry := range userAndBalances {
			if entry.Updated {
				dieselBalance = append(dieselBalance, entry)
				entry.Updated = false
			}
		}
	}
	err = tx.Clauses(clause.OnConflict{UpdateAll: true}).CreateInBatches(dieselBalance, 500).Error
	log.CheckFatal(err)
	//
	rewards := map[string]*pool_lmrewards.LMReward{}
	for _, farmAndItsUsers := range mdl.farmUserRewards {
		for _, user := range farmAndItsUsers {
			farm := mdl.farms[user.Farm]
			reward := user.GetPoints(farm, currentTs)
			key := user.Account + farm.Pool + farm.RewardToken
			if rewards[key] == nil {
				rewards[key] = &pool_lmrewards.LMReward{
					User:        user.Account,
					Pool:        farm.Pool,
					RewardToken: farm.RewardToken,
					// Farm:   farm.Farm,
					Reward: core.NewBigInt(nil),
				}
			}
			rewards[key].Reward = (*core.BigInt)(new(big.Int).Add(rewards[key].Reward.Convert(), reward))
		}
	}
	dataToSave := []*pool_lmrewards.LMReward{}
	for _, entry := range rewards {
		dataToSave = append(dataToSave, entry)
	}
	err = tx.Clauses(clause.OnConflict{UpdateAll: true}).CreateInBatches(dataToSave, 500).Error
	log.CheckFatal(err)

	dieselSYncBLock := []DieselSync{}
	for pool, syncedTill := range mdl.poolsToSyncedTill {
		dieselSYncBLock = append(dieselSYncBLock, DieselSync{
			Pool:           pool.Hex(),
			PoolSyncedTill: syncedTill,
		})
	}
	err = tx.Clauses(clause.OnConflict{UpdateAll: true}).CreateInBatches(dieselSYncBLock, 500).Error
	log.CheckFatal(err)
}
