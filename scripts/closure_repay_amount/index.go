package main

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/repository"
)

type data struct {
	Transfers core.Hstore  `gorm:"column:transfers"`
	AmountBi  *core.BigInt `gorm:"column:amount_bi"`
	LogId     int64        `gorm:"column:log_id"`
	BlockNum  int64        `gorm:"column:block_num"`
}

func main() {
	cfg := config.NewConfig()
	db := repository.NewDBClient(cfg)
	query := `select a.log_id , a.block_num, a.tx_hash, a.transfers, b.amount_bi from (select * from account_operations where action like 'Decrease%' and args->>'amount' = '115792089237316195423570985008687907853269984665640564039457584007913129639935' ) a join (select * from pool_ledger where event='Repay' ) b on b.tx_hash = a.tx_hash;`
	d := []data{}
	err := db.Raw(query).Find(&d).Error
	log.CheckFatal(err)
	for _, entry := range d {
		m := entry.Transfers.GetMap()
		for key := range m {
			m[key] = new(big.Int).Neg(entry.AmountBi.Convert()).String()
		}
		log.Info(m)
		err = db.Exec(`update account_operations set transfers =? where log_id =? and block_num = ?`, entry.Transfers, entry.LogId, entry.BlockNum).Error
		log.CheckFatal(err)
	}
}
