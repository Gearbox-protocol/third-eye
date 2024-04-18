package main

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressorv3"
	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	v3 "github.com/Gearbox-protocol/third-eye/models/pool_lmrewards/v3"
	"github.com/Gearbox-protocol/third-eye/repository"
	"github.com/Gearbox-protocol/third-eye/scripts/helper"
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"
)

func main() {
	cfg := config.NewConfig()
	client := ethclient.NewEthClient(cfg)

	db := repository.NewDBClient(cfg)

	block := getLatestBlock(db)
	updateDieselBalances(client, block, db)
}

func getLatestBlock(db *gorm.DB) int64 {
	a := struct {
		Id int64 `gorm:"column:id"`
	}{}
	err := db.Raw(`select max(id) id from blocks`).Find(&a).Error
	log.CheckFatal(err)
	return a.Id
}

type x struct {
	Farm      string `gorm:"column:farm"`
	DieselSym string `gorm:"column:diesel_sym"`
	Pool      string `gorm:"column:pool"`
}

func convert(db *gorm.DB) map[common.Address]x {
	farms := []x{}
	//
	err := db.Raw(`select lm.*, t.address pool from (select distinct farm,  diesel_sym from user_lmdetails_v3) lm join tokens t on t.symbol=lm.diesel_sym`).Find(&farms).Error
	log.CheckFatal(err)
	ans := map[common.Address]x{}
	for _, f := range farms {
		ans[common.HexToAddress(f.Pool)] = f
	}
	return ans
}
func updateDieselBalances(client core.ClientI, block int64, db *gorm.DB) {
	type _dieselUser struct {
		User string `gorm:"column:account"`
	}
	data := []_dieselUser{}
	err := db.Raw(`select account from user_lmdetails_v3;`).Find(&data).Error
	log.CheckFatal(err)

	// data = []_dieselUser{{
	// 	User: "0xd620AADaBaA20d2af700853C4504028cba7C3333",
	// }}

	dc := helper.GetDC(client, db)
	pools, ok := dc.GetPoolListv3()
	if !ok {
		return
	}
	poolToFarm := convert(db)
	for _, user := range data {
		log.Info(user)
		diesleBalances := getBalances(user.User, block, pools, client, poolToFarm)
		for _, entry := range diesleBalances {
			log.Info(utils.ToJson(entry))
			// continue
			obj := db.Exec("update user_lmdetails_v3 set diesel_balance=? where account=? and farm=?", entry.DieselBalance, entry.Account, entry.Farm)
			if obj.Error != nil {
				log.Fatal(obj.Error, utils.ToJson(entry))
			}
			if obj.RowsAffected == 0 {
				err := db.Create(entry).Error
				if err != nil {
					log.Fatal(err, utils.ToJson(entry))
				}
			}
		}
	}
}

func getBalances(user string, block int64, poolMap []dataCompressorv3.PoolData, client core.ClientI, poolToFarm map[common.Address]x) (ans []v3.UserLMDetails) {
	calls := []multicall.Multicall2Call{}
	abi := core.GetAbi("Token")
	// log.Info(user, len(poolMap))
	for _, details := range poolMap {
		data, err := abi.Pack("balanceOf", common.HexToAddress(user))
		log.CheckFatal(err)
		calls = append(calls, multicall.Multicall2Call{
			Target:   details.Addr,
			CallData: data,
		})
	}
	result := core.MakeMultiCall(client, block, false, calls)
	for ind, entry := range result {
		pool := poolMap[ind].Addr
		b, ok := core.MulticallAnsBigInt(entry)
		if !ok {
			log.Fatal(pool, "failed to get balance for user", user)
		}
		if b.Cmp(new(big.Int)) != 0 {
			ans = append(ans, v3.UserLMDetails{
				DieselBalance: (*core.BigInt)(b),
				Account:       user,
				Correction:    core.NewBigInt(nil),
				BalancesBI:    core.NewBigInt(nil),
				Farm:          poolToFarm[pool].Farm,
				DieselSym:     poolToFarm[pool].DieselSym,
			})
		}
	}
	return ans
}
