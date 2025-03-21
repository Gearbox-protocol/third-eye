package main

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/repository"
	"gorm.io/gorm/clause"
)

func main() {
	cfg := config.NewConfig()
	client := ethclient.NewEthClient(cfg)
	db := repository.NewDBClient(cfg)
	tvlBlock := schemas.LastSync{}
	err := db.Raw(`select * from debt_sync`).Find(&tvlBlock).Error
	log.CheckFatal(err)
	if tvlBlock.Tvl == 0 {
		data := struct {
			Tvl int64 `gorm:"column:tvl_block"`
		}{}
		query := `select min(discovered_at) tvl_block from sync_adapters where type not in ('RebaseToken','Treasury','LMRewardsv2','LMRewardsv3','GearToken')`
		err := db.Raw(query).Find(&data).Error
		log.CheckFatal(err)
		tvlBlock.Tvl = data.Tvl
	}
	if core.GetChainId(client) == 1 {
		tvlBlock.Tvl = utils.Max(15818887+10_000, tvlBlock.Tvl)
	}
	//
	type Entry = struct {
		Pool        string  `gorm:"column:pool"`
		ExpectedLiq float64 `gorm:"column:expected_liquidity"`
		Price       float64 `gorm:"column:price"`
		BlockNum    int64   `gorm:"column:block_num"`
	}
	entries := []Entry{}
	qyery := `select pool, ps.block_num, expected_liquidity, underlying_token, 
		(select price from price_feeds where 
			token=underlying_token and 
			block_num<=ps.block_num and 
			merged_pf_version>=2 
			order by block_num desc limit 1) price  from 
			(select distinct on (pool) pool, expected_liquidity, block_num from pool_stats  where block_num < ? order by pool, block_num desc) ps join pools p on p.address= ps.pool `
	err = db.Raw(qyery, tvlBlock.Tvl).Find(&entries).Error
	log.CheckFatal(err)
	expected := map[string]float64{}
	for _, v := range entries {
		expected[v.Pool] = v.ExpectedLiq * v.Price
	}
	log.Info("init", tvlBlock, utils.ToJson(entries), "chainid", core.GetChainId(client))

	for start := tvlBlock.Tvl; start < tvlBlock.Debt; {
		end := start + 100_000
		entries = entries[:0]
		query := `select pool, ps.block_num, expected_liquidity, underlying_token, 
		(select price from price_feeds where 
			token=underlying_token and 
			block_num<=ps.block_num and 
			merged_pf_version>=2 
			order by block_num desc limit 1) price 
			from pool_stats ps join pools p on p.address= ps.pool where ps.block_num > ? and ps.block_num <= ? order by block_num`
		err = db.Raw(query, start, end).Find(&entries).Error
		log.CheckFatal(err)

		snaps := []schemas.TvlSnapshots{}
		err = db.Raw(`select * from tvl_snapshots where block_num > ? and block_num <= ? order by block_num`, start, end).Find(&snaps).Error
		log.CheckFatal(err)
		var ind = 0
		var ans []schemas.TvlSnapshots
		for _, snap := range snaps {
			for ind < len(entries) && entries[ind].BlockNum <= snap.BlockNum {
				neww := entries[ind]
				expected[neww.Pool] = neww.ExpectedLiq * neww.Price
				ind++
			}
			snap.ExpectedLiq = summ(expected)
			if snap.ExpectedLiq+100 < snap.AvailableLiquidity {
				// log.Fatal("Expected liquidity is less than available liquidity", utils.ToJson(snap))
			}
			ans = append(ans, snap)
		}
		for _, neww := range entries[ind:] {
			expected[neww.Pool] = neww.ExpectedLiq * neww.Price
		}

		// log.Fatal(end, utils.ToJson(ans))
		tx := db.Begin()
		err = tx.Clauses(clause.OnConflict{UpdateAll: true}).CreateInBatches(ans, 50).Error
		log.CheckFatal(err)
		tx.Exec(`UPDATE debt_sync set tvl_block=?, field_set='t'`, end)
		err := tx.Commit().Error
		log.CheckFatal(err)
		sample := schemas.TvlSnapshots{}
		if len(ans) != 0 {
			sample = ans[len(ans)-1]
		}
		log.Infof("synced till %d: %d exp liq: %f avai: %f", end, len(ans), sample.ExpectedLiq, sample.AvailableLiquidity)
		//
		start = end
	}
}

func summ(expected map[string]float64) float64 {
	var sum float64
	for pool, v := range expected {
		// log.Info(pool, v)
		_ = pool
		sum += v
	}
	return sum
}
