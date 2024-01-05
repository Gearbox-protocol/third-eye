package main

import (
	"github.com/Gearbox-protocol/sdk-go/calc"
	"github.com/Gearbox-protocol/sdk-go/core/schemas/schemas_v3"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg/priceFetcher"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/repository"
)

func main() {
	cfg := config.NewConfig()
	db := repository.NewDBClient(cfg)
	client := ethclient.NewEthClient(cfg)

	tStore := priceFetcher.NewTokensStore(client)

	sessions := []*schemas_v3.TradingPriceObj{}
	err := db.Raw(`WITH cs as (select id, cm.underlying_token, since, remaining_funds from credit_sessions _cs 
		JOIN credit_managers cm on cm.address = _cs.credit_manager where version=300),
	css AS (select distinct on (session_id) cs.since, _css.block_num, session_id, balances, 
			borrowed_amount_bi, collateral_underlying  FROM credit_session_snapshots _css 
			JOIN cs ON cs.id =_css.session_id 
			WHERE (SELECT count(*) FROM (select * from jsonb_object_keys(_css.balances) union  (select cs.underlying_token)) t)>=2  
			ORDER BY session_id, block_num)
	SELECT * FROM cs JOIN css ON css.session_id = cs.id`).Find(&sessions).Error
	log.CheckFatal(err)
	//
	for _, session := range sessions {
		price := calc.CalcEntryPriceBySession(1, tStore, session)
		log.Info("session", session.SessionId, "entryPrice", price)
	}
	//

	err = db.Raw(`WITH cs as (select id, credit_manager, closed_at, remaining_funds from credit_sessions where version=300), 
	debts as (select distinct on (session_id) session_id, cal_borrowed_amt_with_interest_bi from debts where session_id in (select id from cs) order by session_id, block_num desc)
	SELECT cs.*, debts.cal_borrowed_amt_with_interest_bi, cm.underlying_token, css.borrowed_amount_bi,  css.collateral_underlying, css.balances FROM cs
	JOIN credit_session_snapshots css ON css.session_id = cs.id AND css.block_num = cs.closed_at-1
	JOIN credit_managers cm ON cm.address = cs.credit_manager
	LEFT JOIN debts ON debts.session_id = cs.id`).Find(&sessions).Error
	log.CheckFatal(err)

	for _, session := range sessions {
		price := calc.CalcClosePriceBySession(1, tStore, session)
		log.Info("session", session.SessionId, "closePrice", price)
	}
}
