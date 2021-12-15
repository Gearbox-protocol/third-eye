package repository

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
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
			adapters = append(adapters, adapter.GetAdapterState())
			if adapter.HasUnderlyingState() {
				err := tx.Clauses(clause.OnConflict{
					// err := repo.db.Clauses(clause.OnConflict{
					UpdateAll: true,
				}).Create(adapter.GetUnderlyingState()).Error
				log.CheckFatal(err)
			}
		}
		repo.kit.Reset(lvlIndex)
	}
	err := tx.Clauses(clause.OnConflict{
		// err := repo.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(adapters, 50).Error
	log.CheckFatal(err)

	log.Infof("created sync sql update in %f sec", time.Now().Sub(now).Seconds())
	now = time.Now()

	tokens := make([]*core.Token, 0, len(repo.tokens))
	for _, token := range repo.tokens {
		tokens = append(tokens, token)
	}
	err = tx.Clauses(clause.OnConflict{
		// err := repo.db.Clauses(clause.OnConflict{
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

func (repo *Repository) flushDebt(newDebtSyncTill int64) {
	debtLen := len(repo.debts)
	if debtLen == 0 {
		return
	}
	log.Infof("Flushing %d for block:%d", debtLen, newDebtSyncTill)
	tx := repo.db.Begin()
	err := tx.Create(core.DebtSync{LastCalculatedAt: newDebtSyncTill}).Error
	log.CheckFatal(err)
	err = tx.CreateInBatches(repo.debts, 50).Error
	log.CheckFatal(err)
	info := tx.Commit()
	if info.Error != nil {
		log.Fatal(info.Error, *info.Statement)
	}
	repo.debts = []*core.Debt{}
}

func (repo *Repository) clear() {
	for _, session := range repo.sessions {
		if session.ClosedAt != 0 {
			delete(repo.sessions, session.ID)
		}
	}
	repo.blocks = map[int64]*core.Block{}
}
