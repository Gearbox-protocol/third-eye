package v3

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressorv3"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/Gearbox-protocol/third-eye/models/pool_lmrewards"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (mdl *LMRewardsv3) getFarmsAndPoolsv3() {
	if len(mdl.farms) != 0 { // already set
		return
	}
	pools, found := mdl.Repo.GetDCWrapper().GetPoolListv3()
	if found && len(mdl.farms) == 0 {
		mdl.SetData(pools)
	}
}

func (mdl *LMRewardsv3) SetData(pools []dataCompressorv3.PoolData) {
	farmingPools := core.GetFarmingPoolsToSymbolByChainId(core.GetChainId(mdl.Client))
	poolAndFarms := []*Farmv3{}
	for _, pool := range pools {
		for _, zapper := range pool.Zappers {
			if zapper.TokenOut.Hex() == "0x580e39ADb33E106fFc2712cBD57B9cE954dcfE75" { // GHO
				zapper.TokenOut = common.HexToAddress("0xE2037090f896A858E3168B978668F22026AC52e7")
			}
			if zapper.TokenOut.Hex() == "0x7aB44F17EE21A3D6Bb2aeb1c6cA8B875041608C4" { // DAI
				zapper.TokenOut = common.HexToAddress("0xC853E4DA38d9Bd1d01675355b8c8f3BBC1451973")
			}
			// can be diselToken zapperOut -- https://etherscan.io/address/0xcaa199f91294e6ee95f9ea90fe716cbd2f9f2900#code
			if _, ok := farmingPools[zapper.TokenOut]; ok && zapper.TokenIn == pool.Underlying && zapper.TokenOut != pool.DieselToken {
				poolAndFarms = append(poolAndFarms, &Farmv3{
					Farm:        zapper.TokenOut.Hex(),
					Pool:        pool.Addr.Hex(),
					DieselToken: pool.DieselToken.Hex(),
					// initial
					Fpt:         (*core.BigInt)(new(big.Int)),
					TotalSupply: (*core.BigInt)(new(big.Int)),
					Reward:      (*core.BigInt)(new(big.Int)),
				})
			}
		}
	}
	mdl.SetUnderlyingState(poolAndFarms)
}

func (mdl *LMRewardsv3) AddPoolv3(blockNum int64, pool string) {
	dcAddr, found := mdl.Repo.GetDCWrapper().GetLatestv3DC()
	if !found {
		log.Fatalf("DC not found for for %s at %d", pool, blockNum)
	}
	con, err := dataCompressorv3.NewDataCompressorv3(dcAddr, mdl.Client)
	log.CheckFatal(err)
	data, err := con.GetPoolData(&bind.CallOpts{BlockNumber: big.NewInt(blockNum)}, common.HexToAddress(pool))
	if err != nil {
		log.Fatal(err, blockNum, pool)
	}
	mdl.SetData([]dataCompressorv3.PoolData{data})
}

func (mdl *LMRewardsv3) SetUnderlyingState(obj interface{}) {
	switch ans := obj.(type) {
	case []*Farmv3:
		if mdl.farms == nil {
			mdl.farms = map[string]*Farmv3{}
		}
		if mdl.pools == nil {
			mdl.pools = map[common.Address]string{}
		}
		for _, farm := range ans {
			if mdl.farms[farm.Farm] == nil {
				mdl.farms[farm.Farm] = farm
				mdl.pools[common.HexToAddress(farm.Pool)] = farm.Farm
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
