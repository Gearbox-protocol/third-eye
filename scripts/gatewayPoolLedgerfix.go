package main

import (
	"context"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
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
	client, err := ethclient.Dial(utils.GetEnvOrDefault("REACT_APP_ADDRESS_PROVIDER", ""))
	log.CheckFatal(err)
	cfg := config.Config{DatabaseUrl: utils.GetEnvOrDefault("DATABASE_URL", "")}
	db := repository.NewDBClient(&cfg)
	chainId, err := client.ChainID(context.Background())
	log.CheckFatal(err)

	store := core.GetSymToAddrStore(log.GetNetworkName(chainId.Int64()) + ".jsonnet")
	for _, syms := range [][]string{
		{"GEARBOX_WETH_POOL", "WETH_GATEWAY"},
		{"GEARBOX_WSTETH_POOL", "WSTETH_GATEWAY"},
	} {
		poolSym := syms[0]
		gatewaySym := syms[1]
		gateway := store.Exchanges[gatewaySym]
		pool := store.Exchanges[poolSym]
		processGateway(gateway, pool, gatewaySym, client, db)
	}
}

func checkGatewayInDB(gateway, poolOriginal common.Address, txLogs []types.Log, db *gorm.DB) (count int) {
	for _, txLog := range txLogs {
		indToSearch := (txLog.Index - 2)
		pool := common.BytesToAddress(txLog.Topics[1][:]).Hex()
		// user := common.BytesToAddress(txLog.Topics[2][:]).Hex()
		ans := schemas.PoolLedger{}
		if obj := db.Model(ans).Where("log_id=? and block_num =? and user_address=? and pool=?", indToSearch, txLog.BlockNumber, gateway.Hex(), pool).First(&ans); obj.Error != nil {
			log.Fatal(indToSearch, txLog.BlockNumber, gateway.Hex(), pool, obj.Error)
		} else {
			count += int(obj.RowsAffected)
		}
	}
	//
	type Allgateway struct {
		count int `gorm:"column:count"`
	}
	allgatewayEntries := &Allgateway{}
	if err := db.Raw(`SELECT count(*) FROM pool_ledger WHERE gateway=? and pool=? and event='RemoveLiquidity'`, gateway.Hex(), poolOriginal.Hex()).Find(allgatewayEntries).Error; err != nil {
		log.Fatal(err)
	}
	log.Info("Number of gateway records in pool_ledger: ", allgatewayEntries.count)
	if allgatewayEntries.count != count {
		log.Fatal("Number of gateway records in pool_ledger is not equal to number of records in pool_ledger to update")
	}
	return
}

func updateGatewayInDB(gateway, poolOriginal common.Address, txLogs []types.Log, db *gorm.DB) (count int) {
	for _, txLog := range txLogs {
		indToSearch := (txLog.Index - 2)
		pool := common.BytesToAddress(txLog.Topics[1][:]).Hex()
		user := common.BytesToAddress(txLog.Topics[2][:]).Hex()
		if obj := db.Exec("update pool_ledger set user_address=? WHERE log_id=? and block_num =? and user_address=? and pool=?", user, indToSearch, txLog.BlockNumber, gateway.Hex(), pool); obj.Error != nil {
			log.Fatal(indToSearch, txLog.BlockNumber, gateway.Hex(), pool, obj.Error)
		} else {
			count += int(obj.RowsAffected)
		}
	}
	return
}

func processGateway(gateway, pool common.Address, gatewaySym string, client core.ClientI, db *gorm.DB) {
	maxBlock := schemas.Block{}
	if err := db.Raw("SELECT max(id) id from blocks").Find(&maxBlock).Error; err != nil {
		log.Fatal(err)
	}
	txLogs, err := core.Node{Client: client}.GetLogs(0, maxBlock.BlockNumber, []common.Address{gateway}, [][]common.Hash{{
		core.Topic("WithdrawETH(address,address)"),
	}})
	log.CheckFatal(err)

	countCheck := checkGatewayInDB(gateway, pool, txLogs, db)
	log.Infof("Number of records(%s) to update: %d", gatewaySym, countCheck)
	// countUpdate := updateGatewayInDB(gateway, txLogs, db)
	// log.Info("Numbers of records  updated: ", countUpdate)
}
