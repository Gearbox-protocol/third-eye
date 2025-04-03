package repository

import (
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

func (repo *Repository) Flush(syncTill int64) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	// log.Fatal("")
	// time.Sleep(time.Hour)
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
	repo.saveLMRewardDetailsv2(tx) // save LM reward  and diesel token balances of users
	repo.saveLMRewardDetailsv3(tx, syncTill)

	repo.SessionRepo.Save(tx)
	repo.BlocksRepo.Save(tx)

	repo.AllowedTokenRepo.Save(tx)
	repo.TokenOracleRepo.Save(tx, syncTill)

	// save current treasury snapshot
	repo.TreasuryRepo.Save(tx)
	repo.saveTicker(tx)

	// fix pool v2 ledger
	repo.GetPoolWrapper().UpdatePoolv2Ledger(tx)
	//
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
		maxBlockNum = utils.Max[int64](maxBlockNum, num)
	}
	repo.SessionRepo.Clear(maxBlockNum)
	repo.BlocksRepo.Clear()
}
