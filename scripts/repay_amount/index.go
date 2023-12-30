package main

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/poolv3"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/repository"
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"
)

type cm_block_amount struct {
	BlockNum int64    `gorm:"column:block_num"`
	Amount   *big.Int `gorm:"column:amount"`
}

type stat struct {
	Id            int64  `gorm:"column:id;primaryKey"`
	CreditManager string `gorm:"column:credit_manager"`
	BlockNum      int64  `gorm:"column:block_num"`
	Decimals      int8   `gorm:"column:decimals"`
}

var decimalsStore = map[string]int8{}

func load_cm_stats(db *gorm.DB) []stat {
	stats := []stat{}
	err := db.Raw(`SELECT id, credit_manager, block_num, t.decimals from credit_manager_stats cms 
	join credit_managers cm on cm.address=cms.credit_manager
	join tokens t on t.address= cm.underlying_token
	where cms.credit_manager in (SELECT address from credit_managers where _version=300) order by block_num`).Find(&stats).Error
	log.CheckFatal(err)
	for _, entry := range stats {
		decimalsStore[entry.CreditManager] = int8(entry.Decimals)
	}
	return stats
}
func main() {
	cfg := config.NewConfig()
	client := ethclient.NewEthClient(cfg)
	db := repository.NewDBClient(cfg)

	data := []*schemas.PoolState{}
	err := db.Raw(`SELECT * from pools where _version=300`).Find(&data).Error
	log.CheckFatal(err)

	pools := []common.Address{}
	for _, entry := range data {
		pools = append(pools, common.HexToAddress(entry.Address))
	}

	stats := load_cm_stats(db)
	type _LastSync struct {
		LastSync int64 `gorm:"column:last_sync"`
	}
	last_sync := _LastSync{}
	err = db.Raw(`SELECT last_sync from sync_adapters where disabled='f' group by last_sync`).First(&last_sync).Error
	log.CheckFatal(err)
	log.Info("Fixing till ", last_sync.LastSync)
	var blockNum int64 = last_sync.LastSync

	amounts := map[string][]*cm_block_amount{}
	txLogs, err := pkg.Node{Client: client}.GetLogs(0, blockNum, pools, [][]common.Hash{
		{
			core.Topic("Repay(address,uint256,uint256,uint256)"),
		},
	})
	log.CheckFatal(err)
	poolcon, err := poolv3.NewPoolv3(core.NULL_ADDR, client)
	log.CheckFatal(err)
	log.Info(len(txLogs))

	// UPDATES
	err = db.Exec(`DELETE from pool_ledger where amount_bi='115792089237316195423570985008687907853269984665640564039457584007913129639935'`).Error
	log.CheckFatal(err)
	//
	for _, txLog := range txLogs {
		event, err := poolcon.ParseRepay(txLog)
		log.CheckFatal(err)

		amount := new(big.Int).Sub(new(big.Int).Add(event.BorrowedAmount, event.Profit), event.Loss)
		decimals := getDecimals(event.CreditManager.Hex())
		// UPDATES
		err = db.Exec(`UPDATE pool_ledger set amount_bi=?, amount=? where pool=? and block_num=?`, (*core.BigInt)(amount),
			utils.GetFloat64Decimal(amount, decimals),
			txLog.Address.Hex(),
			int64(txLog.BlockNumber),
		).Error
		log.CheckFatal(err)

		l := len(amounts[event.CreditManager.Hex()])
		if l == 0 || amounts[event.CreditManager.Hex()][l-1].BlockNum != int64(txLog.BlockNumber) {
			previous := new(big.Int)
			if l > 0 {
				previous = amounts[event.CreditManager.Hex()][l-1].Amount
			}
			amounts[event.CreditManager.Hex()] = append(amounts[event.CreditManager.Hex()], &cm_block_amount{
				BlockNum: int64(txLog.BlockNumber),
				Amount:   previous,
			})
		}

		entry := amounts[event.CreditManager.Hex()][len(amounts[event.CreditManager.Hex()])-1]
		entry.Amount = new(big.Int).Add(
			entry.Amount,
			amount,
		)

	}
	log.Info(utils.ToJson(amounts))

	//

	count := map[string]int{}
	for _, stat := range stats {
		c := count[stat.CreditManager]
		for c < len(amounts[stat.CreditManager]) && amounts[stat.CreditManager][c].BlockNum < stat.BlockNum {
			c++
			count[stat.CreditManager]++
		}
		if c != 0 {
			update := cm_update{
				Id:            stat.Id,
				CreditManager: stat.CreditManager,
				BlockNum:      stat.BlockNum,
				TotalRepaid:   utils.GetFloat64Decimal(amounts[stat.CreditManager][c-1].Amount, int8(stat.Decimals)),
				TotalRepaidBI: (*core.BigInt)(amounts[stat.CreditManager][c-1].Amount),
			}
			log.Info(update.CreditManager, update.BlockNum, update.TotalRepaid)

			// UPDATES
			err := db.Model(update).Select("total_repaid_bi", "total_repaid").Updates(update).Error
			log.CheckFatal(err)
		}
	}
	// UDPATES
	for cm, data := range amounts {
		err := db.Exec(`UPDATE credit_managers set total_repaid_bi=?, total_repaid=? WHERE address=?`,
			(*core.BigInt)(data[len(data)-1].Amount),
			utils.GetFloat64Decimal(data[len(data)-1].Amount, getDecimals(cm)),
			cm,
		).Error
		log.CheckFatal(err)
	}
}

type cm_update struct {
	Id            int64        `gorm:"column:id;primaryKey"`
	CreditManager string       `gorm:"column:credit_manager"`
	BlockNum      int64        `gorm:"column:block_num"`
	TotalRepaid   float64      `gorm:"column:total_repaid"`
	TotalRepaidBI *core.BigInt `gorm:"column:total_repaid_bi"`
}

func (cm_update) TableName() string {
	return "credit_manager_stats"
}

func getDecimals(addr string) int8 {
	decimals := decimalsStore[addr]
	if decimals == 0 {
		log.Fatal("No decimals for ", addr)
	}
	return decimals
}
