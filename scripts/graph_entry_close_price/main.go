package main

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg/priceFetcher"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/repository"
)

type SessionLoad struct {
	SessionId  string                `gorm:"column:id"`
	Since      int64                 `gorm:"column:since"`
	ClosedAt   int64                 `gorm:"column:closed_at"`
	OpenBal    *core.DBBalanceFormat `gorm:"column:open_bal"`
	CloseBal   *core.DBBalanceFormat `gorm:"column:close_bal"`
	Underlying string                `gorm:"column:underlying_token"`
}

type SessionSave struct {
	SessionId  string  `gorm:"column:id;primaryKey"`
	EntryPrice float64 `gorm:"column:entry_price"`
	ClosePrice float64 `gorm:"column:close_price"`
}

func (SessionSave) TableName() string {
	return "credit_sessions"
}

func main() {
	cfg := config.NewConfig()
	db := repository.NewDBClient(cfg)
	client := ethclient.NewEthClient(cfg)

	tStore := priceFetcher.NewTokensStore(client)
	oracle := ds.SetOneInchUpdater(client, tStore)

	sessions := []*SessionLoad{}
	err := db.Raw(`SELECT cs.*,cm.underlying_token, close.balances close_bal, ao.balances open_bal FROM (SELECT id, since, closed_at, credit_manager  from credit_sessions WHERE version>=300) cs 
	JOIN credit_managers cm ON cs.credit_manager = cm.address 
	JOIN credit_session_snapshots ao ON ao.block_num = cs.since and ao.session_id=cs.id
	LEFT JOIN credit_session_snapshots close ON close.block_num = cs.closed_at-1 and close.session_id=cs.id`).Find(&sessions).Error
	log.CheckFatal(err)
	ans := []*SessionSave{}
	//
	for _, session := range sessions {
		entry := &SessionSave{
			SessionId:  session.SessionId,
			EntryPrice: oracle.GetCurrentPriceAtBlockNum(session.Since, *session.OpenBal, session.Underlying),
		}
		if session.ClosedAt != 0 {
			entry.ClosePrice = oracle.GetCurrentPriceAtBlockNum(session.ClosedAt, *session.CloseBal, session.Underlying)
		}
		ans = append(ans, entry)
		log.Infof("fetched for %s", session.SessionId)
	}

	log.Info(utils.ToJson(ans))

	for _, session := range ans {
		err := db.Model(session).Select("entry_price", "close_price").Updates(session).Error
		log.CheckFatal(err)
		log.Infof("updated for %s", session.SessionId)
	}
}
