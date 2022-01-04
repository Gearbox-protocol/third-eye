package main

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DBhandler struct {
	db *gorm.DB
}

func NewDBhandler(db *gorm.DB) *DBhandler {
	return &DBhandler{db: db}
}

func (d *DBhandler) delete(table, column string, blockNum int64) {
	err := d.db.Exec("DELETE FROM ? WHERE ? > ?", table, blockNum).Error
	log.CheckFatal(err)
}

func (d *DBhandler) deleteOnBlockNum(table string, blockNum int64) {
	d.delete(table, "block_num", blockNum)
}

func (d *DBhandler) PoolRollback(blockNum int64) {
	d.deleteOnBlockNum("pool_stats", blockNum)
	d.deleteOnBlockNum("pool_ledger", blockNum)
	err := d.db.Exec(`delete from pools where address in 
		(SELECT address WHERE type='Pool' AND sa.discovered_at > ?)`, blockNum).Error
	log.CheckFatal(err)
}

func (d *DBhandler) SyncadapterRollback(blockNum int64) {
	err := d.db.Exec("UPDATE sync_adapters set last_sync = ? WHERE last_sync > ? and disabled='t'", blockNum, blockNum).Error
	log.CheckFatal(err)
	err = d.db.Exec("DELETE FROM sync_adapters WHERE discovered_at > ?", blockNum).Error
	log.CheckFatal(err)
	err = d.db.Exec("UPDATE sync_adapters set last_sync = ?, disabled='f' WHERE disabled='t' AND last_sync > ?",
		blockNum, blockNum).Error
	log.CheckFatal(err)
}
func (d *DBhandler) CSBorrowerAtBlockNum(blockNum int64) map[string]string {
	// get the borrower
	data := []*core.AccountOperation{}
	query := `SELECT distinct on (session_id) session_id, borrower FROM credit_session_snapshots 
		ORDER BY session_id, block_num desc, log_id desc WHERE block_num <= ?`
	err := d.db.Raw(query).Find(&data).Error
	log.CheckFatal(err)
	borrowers := map[string]string{}
	for _, entry := range data {
		borrowers[entry.SessionId] = entry.Borrower
	}
	return borrowers
}
func (d *DBhandler) CreditSessionRollback(blockNum int64) {
	// credit session snapshot for closing of account is added for close_at-1
	// while rollback to x we also need to delete the snapshot for account closed at x+1
	// no doing this will not cause any inconsistency and help prevent multiple session_id/block_num same values
	err := d.db.Exec(`DELETE FROM credit_session_snapshots WHERE id in 
		(SELECT max(id) FROM credit_session_snapshots WHERE 
			session_id in (SELECT session_id FROM credit_sessions WHERE closed_at = ?) AND block_num = ?
			GROUP BY session_id)`, blockNum+1, blockNum).Error
	log.CheckFatal(err)
	err = d.db.Exec("DELETE FROM credit_sessions WHERE since > ?", blockNum).Error
	log.CheckFatal(err)
	err = d.db.Exec("UPDATE credit_sessions set status=0, closed_at = 0 WHERE status != 0 , closed_at > ?", blockNum).Error
	log.CheckFatal(err)
	d.UpdateCSValues(blockNum)
}

func (d *DBhandler) UpdateCSValues(blockNum int64) {
	// for updating the credit_sessions
	data := []*core.CreditSessionSnapshot{}
	query := "SELECT distinct on (session_id) * FROM credit_session_snapshots order by session_id, block_num desc WHERE block_num <= ?"
	err := d.db.Raw(query).Find(&data).Error
	log.CheckFatal(err)
	borrowers := d.CSBorrowerAtBlockNum(blockNum)
	csUpdates := []*core.CreditSessionUpdate{}
	for _, entry := range data {
		borrower := borrowers[entry.SessionId]
		csUpdates = append(csUpdates, &core.CreditSessionUpdate{
			SessionId:        entry.SessionId,
			BorrowedAmountBI: entry.BorrowedAmountBI,
			TotalValueBI:     entry.TotalValueBI,
			HealthFactor:     entry.HealthFactor,
			Borrower:         borrower,
		})
	}
	err = d.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(csUpdates, 50).Error
	log.CheckFatal(err)
}

func (d *DBhandler) CreditManagerRollback(blockNum int64) {
	// delete credit_managers that are not found yet
	err := d.db.Exec(`delete from credit_managers where address IN 
		(SELECT address WHERE type='CreditManager' AND sa.discovered_at > ?)`, blockNum).Error
	log.CheckFatal(err)
	// delete credit filter
	err = d.db.Exec(`delete from sync_adapters WHERE type='CreditFilter' AND details->>'creditManager' in 
		(SELECT address WHERE type='CreditManager' AND sa.discovered_at > ?)`, blockNum).Error
	log.CheckFatal(err)

	data := []*core.CreditManagerStat{}
	err = d.db.Exec(`SELECT distinct on (session_id) * FROM credit_manager_stats 
		WHERE block_num <= ? order by session_id, block_num desc`, blockNum).Error
	log.CheckFatal(err)
	cmUpdates := []*core.CreditManagerUpdate{}
	for _, entry := range data {
		cmUpdates = append(cmUpdates, &core.CreditManagerUpdate{
			Address:           entry.Address,
			CreditManagerData: entry.CreditManagerData,
		})
	}
	err = d.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(data, 50).Error
	log.CheckFatal(err)
}
