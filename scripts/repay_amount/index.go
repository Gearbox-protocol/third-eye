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
)

type cm_block_amount struct {
	BlockNum int64    `gorm:"column:block_num"`
	Amount   *big.Int `gorm:"column:amount"`
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
	var blockNum int64 = 18900932

	amounts := map[string][]*cm_block_amount{}
	txLogs, err := pkg.Node{Client: client}.GetLogs(0, blockNum, pools, [][]common.Hash{
		{
			core.Topic("Repay(address,uint256,uint256,uint256)"),
		},
	})
	poolcon, err := poolv3.NewPoolv3(core.NULL_ADDR, client)
	log.CheckFatal(err)
	log.Info(len(txLogs))

	// UPDATES
	// err = db.Exec(`DELETE from pool_ledger where amount_bi='115792089237316195423570985008687907853269984665640564039457584007913129639935'`).Error
	// log.CheckFatal(err)
	//
	for _, txLog := range txLogs {
		event, err := poolcon.ParseRepay(txLog)
		log.CheckFatal(err)

		amount := new(big.Int).Sub(new(big.Int).Add(event.BorrowedAmount, event.Profit), event.Loss)
		var decimals int8 = 6
		if txLog.Address.Hex() == "0xda0002859B2d05F66a753d8241fCDE8623f26F4f" {
			decimals = 18
		}
		log.Info(decimals)
		// UPDATES
		// err = db.Exec(`UPDATE pool_ledger set amount_bi=?, amount=? where pool=? and block_num=?`, (*core.BigInt)(amount),
		// 	utils.GetFloat64Decimal(amount, decimals),
		// 	txLog.Address.Hex(),
		// 	int64(txLog.BlockNumber),
		// ).Error
		// log.CheckFatal(err)
		//
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

	type stat struct {
		Id            int64  `gorm:"column:id;primaryKey"`
		CreditManager string `gorm:"column:credit_manager"`
		BlockNum      int64  `gorm:"column:block_num"`
		Decimals      int64  `gorm:"column:decimals"`
	}
	stats := []stat{}
	err = db.Raw(`SELECT id, credit_manager, block_num, t.decimals from credit_manager_stats cms 
	join credit_managers cm on cm.address=cms.credit_manager
	join tokens t on t.address= cm.underlying_token
	where cms.credit_manager in (SELECT address from credit_managers where _version=300) order by block_num`).Find(&stats).Error
	log.CheckFatal(err)
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
			// err := db.Model(update).Select("total_repaid_bi", "total_repaid").Updates(update).Error
			// log.CheckFatal(err)
		}
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
