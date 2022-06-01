package repository

import (
	"time"

	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
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
	repo.SyncAdaptersRepo.Save(tx)

	repo.TokensRepo.Save(tx)

	repo.SessionRepo.Save(tx)

	repo.BlocksRepo.Save(tx)

	if len(repo.relations) > 0 {
		err := tx.CreateInBatches(repo.relations, 3000).Error
		log.CheckFatal(err)
		repo.relations = []*schemas.UniPriceAndChainlink{}
	}

	repo.AllowedTokenRepo.Save(tx)

	// save current treasury snapshot
	now := time.Now()
	if repo.treasurySnapshot.Date != "" {
		err := tx.Clauses(clause.OnConflict{
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
