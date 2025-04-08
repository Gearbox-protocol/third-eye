package main

import (
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/repository"
)

func main() {
	cfg := config.NewConfig()
	// log.Info(cfg.EthProvider)
	// log.Fatal(cfg.DatabaseUrl)
	db := repository.NewDBClient(cfg)
	b := &struct {
		BlockNum int64 `gorm:"column:block_num"`
	}{}
	err := db.Raw("select min (block_num) block_num from price_feeds").Find(b).Error
	log.CheckFatal(err)
	b.BlockNum = b.BlockNum - 1
	for {
		start := b.BlockNum + 1
		end := b.BlockNum + 100_00_000
		// log.Info(start, end)
		tx := db.Begin()
		//
		err := tx.Exec(`create table jj as select distinct on (feed, block_num) (block_num)  block_num, feed, price_bi, 
		price, round_id from price_feeds where block_num >=? and block_num <=? order by block_num ,feed`, start, end)
		log.CheckFatal(err.Error)
		log.Info("create", err.RowsAffected)
		if err.RowsAffected == 0 {
			err = tx.Exec(`drop table jj`)
			log.CheckFatal(err.Error)
			log.Info("drop")
			tx.Commit()
			break
		}
		err = tx.Exec(`delete from price_feeds where block_num >=? and block_num <=?`, start, end)
		log.CheckFatal(err.Error)
		log.Info("delete", err.RowsAffected)
		//
		err = tx.Exec(`insert into price_feeds(block_num , feed, price_bi, price , round_id) select * from jj`)
		log.CheckFatal(err.Error)
		log.Info("insert", err.RowsAffected)
		//
		//
		err = tx.Exec(`drop table jj`)
		log.CheckFatal(err.Error)
		log.Info("drop")
		tx.Commit()
		//
		log.Info("saved till", end)
		b.BlockNum += 500_000
	}

	err = db.Exec(`alter table price_feeds add PRIMARY KEY (block_num, feed)`).Error
	log.CheckFatal(err)
}
