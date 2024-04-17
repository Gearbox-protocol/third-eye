package main

import (
	"strings"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg/priceFetcher"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/models/credit_manager/cm_common"
	"github.com/Gearbox-protocol/third-eye/repository"
	"github.com/Gearbox-protocol/third-eye/scripts/helper"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	cfg := config.NewConfig()
	client := ethclient.NewEthClient(cfg)

	db := repository.NewDBClient(cfg)

	a := []struct {
		SessionId string `gorm:"column:session_id"`
		BlockNum  int64  `gorm:"column:block_num"`
	}{}
	err := db.Raw(`select distinct on (session_id) session_id, block_num from credit_session_snapshots where session_id in (select id from credit_sessions where version=300 and status=0) order by session_id , block_num desc`).Find(&a).Error
	log.CheckFatal(err)

	dc := helper.GetDC(client, db)
	//
	store := priceFetcher.NewTokensStore(client)
	tokens := core.GetSymToAddrByChainId(core.GetChainId(client)).Tokens
	for _, entry := range a {
		account := strings.Split(entry.SessionId, "_")[0]
		call, result, err := dc.GetCreditAccountData(core.NewVersion(300), entry.BlockNum, core.NULL_ADDR, core.NULL_ADDR, common.HexToAddress(account))
		log.CheckFatal(err)
		results := core.MakeMultiCall(client, entry.BlockNum, false, []multicall.Multicall2Call{call})
		if results[0].Success {
			data, err := result(results[0].ReturnData)
			if err != nil {
				log.Fatal(err, entry.SessionId, entry.BlockNum, results[0].ReturnData)
			}

			balances := cm_common.AddStETHBalance(data.Addr.Hex(), entry.BlockNum, data.Balances, client, store, tokens["stETH"].Hex())
			log.Info(entry.SessionId, utils.ToJson(balances))
			err = db.Exec(`update credit_session_snapshots set balances=? where session_id=? and block_num=?`, balances, entry.SessionId, entry.BlockNum).Error
			log.CheckFatal(err)
			err = db.Exec(`update credit_sessions set balances=? where id=?`, balances, entry.SessionId).Error
			log.CheckFatal(err)
		}
	}
}
