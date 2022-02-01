package repository

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/utils"
	"gorm.io/gorm/clause"
	"time"
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

	adapters := make([]*core.SyncAdapter, 0, repo.kit.Len())
	for lvlIndex := 0; lvlIndex < repo.kit.Len(); lvlIndex++ {
		for repo.kit.Next(lvlIndex) {
			adapter := repo.kit.Get(lvlIndex)
			if adapter.GetName() != core.AggregatedBlockFeed {
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
	// save yearnPriceFeeds
	for _, adapter := range repo.aggregatedFeed.GetYearnFeeds() {
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

	if len(repo.relations) > 0 {
		err := tx.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).CreateInBatches(repo.relations, 50).Error
		log.CheckFatal(err)
		repo.relations = []*core.UniPriceAndChainlink{}
	}
	log.Infof("created sync adapters sql update in %f sec", time.Now().Sub(now).Seconds())
	now = time.Now()

	tokens := make([]*core.Token, 0, len(repo.tokens))
	for _, token := range repo.tokens {
		tokens = append(tokens, token)
	}
	err = tx.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(tokens, 50).Error
	log.CheckFatal(err)

	log.Infof("created tokens sql statements in %f sec", time.Now().Sub(now).Seconds())
	now = time.Now()

	for _, session := range repo.sessions {
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

	blocksToSync := make([]*core.Block, 0, len(repo.blocks))
	for _, block := range repo.blocks {
		blocksToSync = append(blocksToSync, block)
	}
	err = tx.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(blocksToSync, 100).Error
	log.CheckFatal(err)

	// add disabled tokens after the block num is synced to db
	if len(repo.disabledTokens) > 0 {
		err = tx.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).CreateInBatches(repo.disabledTokens, 50).Error
		log.CheckFatal(err)
		repo.disabledTokens = []*core.AllowedToken{}
	}

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
	for num := range repo.blocks {
		maxBlockNum = utils.Max(maxBlockNum, num)
	}
	for _, session := range repo.sessions {
		if session.ClosedAt != 0 && maxBlockNum >= session.ClosedAt {
			delete(repo.sessions, session.ID)
		}
	}
	repo.blocks = map[int64]*core.Block{}
}
