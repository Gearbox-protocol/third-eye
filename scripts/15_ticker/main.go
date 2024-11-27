package main

import (
	"strings"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/models/price_oracle"
	"github.com/Gearbox-protocol/third-eye/repository"
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm/clause"
)

type ticker struct {
	Feed   string `gorm:"column:feed;primaryKey"`
	Ticker string `gorm:"column:ticker"`
}

func main() {
	cfg := config.NewConfig()
	db := repository.NewDBClient(cfg)
	client := ethclient.NewEthClient(cfg)
	data := &ds.SyncAdapter{}
	err := db.Raw(`select address, last_sync from sync_adapters where version=300 and type='PriceOracle'`).Find(data).Error
	log.CheckFatal(err)
	//
	node := pkg.Node{Client: client}
	txLogs, err := node.GetLogs(0, data.LastSync, []common.Address{common.HexToAddress(data.Address)}, [][]common.Hash{{
		core.Topic("SetPriceFeed(address,address,uint32,bool,bool)"),
		core.Topic("SetReservePriceFeed(address,address,uint32,bool)"),
	}})
	log.CheckFatal(err)

	ls := []ticker{}
	for _, txLog := range txLogs {
		token := common.BytesToAddress(txLog.Topics[1].Bytes())        // token
		oracle := common.BytesToAddress(txLog.Topics[2].Bytes()).Hex() // priceFeed
		desc := price_oracle.GetDesc(client, token)
		if strings.Contains(desc, "Ticker Token") || oracle == "0x14497e822B70554537dB9950126461C23dC4f237" { // ezETH/ETH and weETH/ETH
			// for arbitrum token 0x144
			log.Info("here")
			ls = append(ls, ticker{
				Feed:   oracle,
				Ticker: token.Hex(),
			})
		}
	}

	err = db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(ls, 10).Error
	log.CheckFatal(err)

}
