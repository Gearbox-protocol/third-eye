package v3

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressorv3"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"

	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/pool_lmrewards"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (mdl *LMRewardsv3) getFarmsAndPoolsv3() {
	// if len(mdl.farms) != 0 { // already set
	// 	return
	// }
	pools, found := mdl.Repo.GetDCWrapper().GetPoolListv3()
	if found {
		mdl.SetFarm(pools)
	}
}

func (mdl *LMRewardsv3) setMinPoolSyncedTill(pool common.Address, syncedTill int64, farm string) {
	if mdl.poolsToSyncedTill[pool] == 0 {
		mdl.poolsToSyncedTill[pool] = syncedTill
	}
	if mdl.poolsToSyncedTill[pool] != syncedTill {
		log.Fatal("in farm_v3 the pool_synced_till is different from different farms of that pool", syncedTill, mdl.poolsToSyncedTill[pool], farm)
	}
}

func (mdl *LMRewardsv3) SetFarm(pools []dataCompressorv3.PoolData) {
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
				_, err := core.CallFuncWithExtraBytes(mdl.Client, "bfe10928", zapper.TokenOut, 0, nil) // distributor on the farm
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
					PoolSyncedTill: mdl.Repo.GetAdapter(pool.Addr.Hex()).GetDiscoveredAt(),
				}
				if mdl.farms[farm.Farm] == nil {
					farm.setRewardToken(mdl.Client)
					mdl.farms[farm.Farm] = farm
					// poolAndFarms = append(poolAndFarms, farm)
					newfarmsForPool = append(newfarmsForPool, common.HexToAddress(farm.Farm))
					mdl.LastSync = utils.Min(mdl.LastSync, farm.GetMinSyncedTill())
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

func (mdl *LMRewardsv3) AddPoolv3(blockNum int64, pool string) {
	dcAddr, found := mdl.Repo.GetDCWrapper().GetLatestv3DC()
	if !found {
		if core.GetBaseChainId(mdl.Client) == 146 && blockNum < 9790594 {
			return
		}
		log.Fatalf("DC not found for for %s at latest, blockNum %d ", pool, blockNum)
	}
	con, err := dataCompressorv3.NewDataCompressorv3(dcAddr, mdl.Client)
	log.CheckFatal(err)
	data, err := con.GetPoolData(nil, common.HexToAddress(pool))
	if err != nil {
		log.Fatal(err, "latest", pool)
	}
	mdl.SetFarm([]dataCompressorv3.PoolData{data})
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
				mdl.setMinPoolSyncedTill(common.HexToAddress(farm.Pool), farm.PoolSyncedTill, farm.Farm)
			}
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
}
