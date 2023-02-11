package main

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/models/pool"
	"github.com/Gearbox-protocol/third-eye/repository"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	log.CheckFatal(err)
	client, err := ethclient.Dial(utils.GetEnvOrDefault("ETH_PROVIDER", ""))
	log.CheckFatal(err)
	cfg := config.Config{DatabaseUrl: utils.GetEnvOrDefault("DATABASE_URL", "")}
	db := repository.NewDBClient(&cfg)
	for pool, details := range pool.GetPoolGateways(client) {
		log.Info(pool, details.Gateway, details.Sym)
		processGateway(pool, details, client, db)
	}
}

// ind, user, ignore
func getIndUser(txLog types.Log, details pool.GatewayDetails, pool common.Address) (indToSearch []int64, user common.Address, ignore bool) {
	if details.Sym == "WETH" {
		indToSearch = []int64{int64(txLog.Index - 2)}
		user = common.BytesToAddress(txLog.Topics[2][:])
	}
	if details.Sym == "WSTETH" {
		from := common.BytesToAddress(txLog.Topics[1][:]).Hex()
		user = common.BytesToAddress(txLog.Topics[2][:])
		if !(from == details.Gateway.Hex() && user != details.UserCantBe) {
			return nil, common.Address{}, true
		}
		indToSearch = []int64{int64(txLog.Index - 3), int64(txLog.Index - 4)}
	}
	return indToSearch, user, false
}

func checkGatewayInDB(poolOriginal common.Address, details pool.GatewayDetails, txLogs []types.Log, db *gorm.DB) {
	type Allgateway struct {
		Count int `gorm:"column:count"`
	}
	allgatewayEntries := &Allgateway{}
	if err := db.Raw(`SELECT count(*) FROM pool_ledger WHERE user_address=? and pool=? and event='RemoveLiquidity'`, details.Gateway.Hex(), poolOriginal.Hex()).Find(allgatewayEntries).Error; err != nil {
		log.Fatal(err)
	}
	log.Info("Number of gateway records in pool_ledger: ", allgatewayEntries.Count)
	//
	updateCount := 0
	for _, txLog := range txLogs {
		pool := common.BytesToAddress(txLog.Topics[1][:])
		if details.Sym == "WETH" && pool != poolOriginal {
			log.Fatal("Pool in gateway's WithdrawETH is not equal to pool WETH")
		}
		indToSearch, _, ignore := getIndUser(txLog, details, poolOriginal)
		if ignore {
			continue
		}
		//
		ans := schemas.PoolLedger{}
		if obj := db.Model(ans).Where("log_id in ? and block_num =? and user_address=? and pool=?", indToSearch, txLog.BlockNumber, details.Gateway.Hex(), poolOriginal.Hex()).First(&ans); obj.Error != nil {
			log.Fatal(indToSearch, txLog.BlockNumber, txLog.TxHash, details.Gateway.Hex(), pool)
		} else {
			updateCount += int(obj.RowsAffected)
		}
	}
	log.Infof("Number of records(%s) to update: %d", details.Sym, updateCount)
	//

	if allgatewayEntries.Count != updateCount {
		log.Fatalf("Number of gateway records in pool_ledger(%d) is not equal to number of records in pool_ledger to update(%d)", allgatewayEntries.Count, updateCount)
	}
	return
}

func updateGatewayInDB(pool common.Address, details pool.GatewayDetails, txLogs []types.Log, db *gorm.DB) (count int) {
	for _, txLog := range txLogs {
		indToSearch, user, ignore := getIndUser(txLog, details, pool)
		if ignore {
			continue
		}
		if obj := db.Exec("update pool_ledger set user_address=? WHERE log_id in ? and block_num =? and user_address=? and pool=?",
			user.Hex(), indToSearch, txLog.BlockNumber, details.Gateway.Hex(), pool.Hex()); obj.Error != nil {
			log.Fatal(indToSearch, txLog.BlockNumber, details.Gateway.Hex(), pool, obj.Error)
		} else {
			count += int(obj.RowsAffected)
		}
	}
	return
}

func processGateway(pool common.Address, details pool.GatewayDetails, client core.ClientI, db *gorm.DB) {
	maxBlock := schemas.Block{}
	if err := db.Raw("SELECT max(id) id from blocks").Find(&maxBlock).Error; err != nil {
		log.Fatal(err)
	}
	txLogs, err := core.Node{Client: client}.GetLogs(0, maxBlock.BlockNumber, []common.Address{details.Gateway, details.Token}, [][]common.Hash{{
		core.Topic("WithdrawETH(address,address)"),
		core.Topic("Transfer(address,address,uint256)"),
	}})
	log.CheckFatal(err)
	checkGatewayInDB(pool, details, txLogs, db)
	countUpdate := updateGatewayInDB(pool, details, txLogs, db)
	log.Info("Numbers of records  updated: ", countUpdate)
}
