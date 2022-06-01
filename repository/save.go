package repository

import (
	"time"

	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"gorm.io/gorm/clause"
)

func (repo *Repository) Flush() error {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	// preferred order adapter => token => pools => cm => credit session => blocks => allowedTokens

	// credit manager depends on pools
	// allowed token depends on credit managers
	// credit sesion depends on credit manager
	// credit session snapshot on credit session

	// will be depended in future
	// block->pricefeed on token
	// block->protocols on creditManager
	// block->AccountOperation on session
	// block->AllowedTOken on session

	tx := repo.db.Begin()
	now := time.Now()

	adapters := make([]*ds.SyncAdapter, 0, repo.kit.Len())
	for lvlIndex := 0; lvlIndex < repo.kit.Len(); lvlIndex++ {
		for repo.kit.Next(lvlIndex) {
			adapter := repo.kit.Get(lvlIndex)
			if adapter.GetName() != ds.AggregatedBlockFeed {
				adapters = append(adapters, adapter.GetAdapterState())
			}
			if adapter.HasUnderlyingState() {
				err := tx.Clauses(clause.OnConflict{
					UpdateAll: true,
				}).Create(adapter.GetUnderlyingState()).Error
				log.CheckFatal(err)
			}
		}
		repo.kit.Reset(lvlIndex)
	}
	// save qyery feeds from aggregatedFeed
	for _, adapter := range repo.aggregatedFeed.GetQueryFeeds() {
		adapters = append(adapters, adapter.GetAdapterState())
	}
	err := tx.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(adapters, 50).Error
	log.CheckFatal(err)

	if uniPools := repo.aggregatedFeed.GetUniswapPools(); len(uniPools) > 0 {
		err := tx.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).CreateInBatches(uniPools, 50).Error
		log.CheckFatal(err)
	}

	log.Infof("created sync adapters sql update in %f sec", time.Now().Sub(now).Seconds())
	now = time.Now()

	repo.TokensRepo.Save(tx)

	log.Infof("created tokens sql statements in %f sec", time.Now().Sub(now).Seconds())
	now = time.Now()

	for _, session := range repo.GetSessions() {
		if session.IsDirty {
			err := tx.Clauses(clause.OnConflict{
				UpdateAll: true,
			}).Create(session).Error
			log.CheckFatal(err)
			session.IsDirty = false
		}
	}

	log.Infof("created session sql update in %f sec", time.Now().Sub(now).Seconds())
	now = time.Now()

	blocksToSync := make([]*schemas.Block, 0, len(repo.GetBlocks()))
	for _, block := range repo.GetBlocks() {
		blocksToSync = append(blocksToSync, block)
	}
	err = tx.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(blocksToSync, 100).Error
	log.CheckFatal(err)

	if len(repo.relations) > 0 {
		err := tx.CreateInBatches(repo.relations, 3000).Error
		log.CheckFatal(err)
		repo.relations = []*schemas.UniPriceAndChainlink{}
	}

	repo.AllowedTokenRepo.Save(tx)

	// save current treasury snapshot
	if repo.treasurySnapshot.Date != "" {
		err = tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "block_num"}},
			DoUpdates: clause.AssignmentColumns([]string{"date_str", "prices_in_usd", "balances", "value_in_usd"}),
		}).Create(repo.treasurySnapshot).Error
		log.CheckFatal(err)
	}

	log.Infof("created blocks sql update in %f sec", time.Now().Sub(now).Seconds())
	info := tx.Commit()
	log.CheckFatal(info.Error)
	return nil
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (repo *Repository) Clear() {
	var maxBlockNum int64
	for num := range repo.GetBlocks() {
		maxBlockNum = utils.Max(maxBlockNum, num)
	}
	repo.SessionRepo.Clear(maxBlockNum)
	repo.BlocksRepo.Clear()
}
