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
	"gorm.io/gorm"
)

func main() {
	cfg := config.NewConfig()
	client := ethclient.NewEthClient(cfg)

	db := repository.NewDBClient(cfg)

	node := pkg.Node{Client: client}
	block := node.GetLatestBlockNumber()

	// setlimits(db, block, node)
	setTotalDebtLimit(db, block, node)
}

func setlimits(db *gorm.DB, block int64, node pkg.Node) {
	a := []struct {
		Address      string `gorm:"column:address"`
		Configurator string `gorm:"column:configurator"`
	}{}
	err := db.Raw(`select address, details->>'configurator' configurator from sync_adapters where version=300 and type='CreditManager'`).Find(&a).Error
	log.CheckFatal(err)
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
		//
		{
			type _d struct {
				MinAmount *core.BigInt `gorm:"column:min_amount"`
				MaxAmount *core.BigInt `gorm:"column:max_amount"`
			}
			d := &_d{}
			err = db.Raw(`select min_amount, max_amount from credit_managers where address=?`, entry.Address).Find(d).Error
			log.CheckFatal(err)
			log.Info(d.MinAmount, d.MaxAmount)
		}
		err = db.Exec(`update credit_managers set min_amount=? , max_amount=? where address=?`, (*core.BigInt)(minDebt), (*core.BigInt)(maxDebt), entry.Address).Error
		log.CheckFatal(err)
		err = db.Exec(`update parameters set min_amount=? , max_amount=? where credit_manager=?`, (*core.BigInt)(minDebt), (*core.BigInt)(maxDebt), entry.Address).Error
		log.CheckFatal(err)
	}
}
func setTotalDebtLimit(db *gorm.DB, block int64, node pkg.Node) {
	a := []struct {
		Pool string `gorm:"column:address"`
	}{}
	err := db.Raw(`select address from pools where _version=300`).Find(&a).Error
	log.CheckFatal(err)
	for _, entry := range a {
		log.Infof("%+v", entry)
		txLogs, err := node.GetLogs(0, block, []common.Address{common.HexToAddress(entry.Pool)}, [][]common.Hash{
			{core.Topic("SetCreditManagerDebtLimit(address,uint256)")},
		})
		log.CheckFatal(err)
		for _, txLog := range txLogs {
			cm := common.BytesToAddress(txLog.Topics[1][:])
			total := new(big.Int).SetBytes(txLog.Data[:32])
			err = db.Exec(`update credit_managers set total_debt_limit=?  where address=?`, (*core.BigInt)(total), cm.Hex()).Error
			log.CheckFatal(err)
		}
	}
}
