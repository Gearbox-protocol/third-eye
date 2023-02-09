package repository

import (
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

func (repo *Repository) Flush() error {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	// preferred order (adapter | token) => pools => cm => credit session => blocks => allowedTokens

	// credit manager depends on pools
	// allowed token depends on credit managers
	// credit sesion depends on credit manager
	// credit session snapshot on credit session

	// block->pricefeed on token
	// block->protocols on creditManager
	// block->AccountOperation on session
	// block->AllowedTOken on session

	tx := repo.db.Begin()
	repo.TokensRepo.Save(tx)

	repo.SyncAdaptersRepo.Save(tx)
	repo.saveLMRewardDetails(tx) // save LM reward  and diesel token balances of users

	repo.SessionRepo.Save(tx)

	repo.BlocksRepo.Save(tx)

	if len(repo.relations) > 0 {
		err := tx.CreateInBatches(repo.relations, 3000).Error
		log.CheckFatal(err)
		repo.relations = []*schemas.UniPriceAndChainlink{}
	}

	repo.AllowedTokenRepo.Save(tx)

	// save current treasury snapshot
	repo.TreasuryRepo.Save(tx)
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
