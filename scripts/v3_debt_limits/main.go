package main

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/repository"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	cfg := config.NewConfig()
	client := ethclient.NewEthClient(cfg)

	db := repository.NewDBClient(cfg)

	a := []struct {
		Address      string `gorm:"column:address"`
		Configurator string `gorm:"column:configurator"`
	}{}
	err := db.Raw(`select address, details->>'configurator' configurator from sync_adapters where version=300 and type='CreditManager'`).Find(&a).Error
	log.CheckFatal(err)

	node := pkg.Node{Client: client}
	block := node.GetLatestBlockNumber()
	for _, entry := range a {
		log.Infof("%+v", entry)
		txLogs, err := node.GetLogs(0, block, []common.Address{common.HexToAddress(entry.Configurator)}, [][]common.Hash{
			{core.Topic("SetBorrowingLimits(uint256,uint256)")},
		})
		log.CheckFatal(err)
		txLog := txLogs[len(txLogs)-1]
		minDebt := new(big.Int).SetBytes(txLog.Data[:32])
		maxDebt := new(big.Int).SetBytes(txLog.Data[32:])
		log.Info(minDebt, maxDebt)
		err = db.Exec(`update credit_managers set min_amount=? , max_amount=? where address=?`, (*core.BigInt)(minDebt), (*core.BigInt)(maxDebt), entry.Address).Error
		log.CheckFatal(err)
	}
}
