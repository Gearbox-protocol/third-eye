package main

import (
	"os"
	"strconv"

	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/aggregated_block_feed/query_price_feed"
	"github.com/Gearbox-protocol/third-eye/repository"
	"gorm.io/gorm/clause"
)

func main() {
	dburl := os.Args[1]
	deletedB, err := strconv.ParseInt(os.Args[2], 10, 64)
	log.CheckFatal(err)
	db := repository.NewDBClient(&config.Config{DatabaseUrl: dburl})
	adapters := []ds.SyncAdapter{}
	err = db.Raw(`select * from sync_adapters where type in ('QueryPriceFeed', 'ChainlinkPriceFeed', 'CompositeChainlinkPF')`).Find(&adapters).Error
	log.CheckFatal(err)
	for _, adapter := range adapters {
		if adapter.GetName() == ds.QueryPriceFeed {
			details := query_price_feed.DetailsDS{}
			details.Load(adapter.GetDetails(), adapter.GetVersion())
			for _, d := range details.Tokens {
				for pf, blocks := range d {
					if len(blocks) == 2 && blocks[1] > deletedB {
						d[pf] = []int64{blocks[0]}
					}
					if len(blocks) >= 1 && blocks[0] > deletedB {
						delete(d, pf)
					}
				}
			}
			adapter.Details = details.Save()
		} else {
			mgr := ds.MergedPFManager{}
			mgr.Load(adapter.GetDetails(), adapter.FirstLogAt)
			mgr.DeleteAfter(deletedB)
			mgr.Save(&adapter.Details)
		}
	}
	err = db.Clauses(clause.OnConflict{
		// err := repo.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(adapters, 50).Error
	log.CheckFatal(err)
}
