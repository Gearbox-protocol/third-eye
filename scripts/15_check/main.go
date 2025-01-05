package main

import (
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/repository"
)

func main() {
	db1 := repository.NewDBClient(&config.Config{
		DatabaseUrl: "postgres://postgres:complex_password_123@144.91.114.166:5432,37.60.225.176:5432/gearbox?connect_timeout=5",
	})

	db2 := repository.NewDBClient(&config.Config{
		DatabaseUrl: "host=84.247.174.187 user=debian password=123Sample dbname=tmp_anvil",
	})

	var b int64 = 13810899
	for {
		start := b
		end := start + 500_000
		data := []*schemas.Debt{}
		err := db1.Raw(`select * from debts where block_num > ? and block_num <=?`, start, end).Find(&data).Error
		log.CheckFatal(err)

		log.Info("loaded", len(data))
		for ind, debt := range data {
			second := &schemas.Debt{}
			err := db2.Raw(`select * from debts where block_num =? and session_id=?`, debt.BlockNumber, debt.SessionId).Find(second).Error
			log.CheckFatal(err)
			if second.CalTotalValueBI.Cmp(debt.CalTotalValueBI) != 0 || second.CalHealthFactor.Cmp(debt.CalHealthFactor) != 0 {
				log.Fatal(utils.ToJson(debt))
				log.Fatal(utils.ToJson(second))
			}
			if ind%1000 == 0 {
				log.Info("checked till", ind)
			}
		}
		log.Info("checked till blockNum", end)
		b += 500_000
	}

}
