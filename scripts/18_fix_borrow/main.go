package main

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/repository"
)

type Event struct {
	CreditManager string     `gorm:"column:credit_manager"`
	BlockNum      int64      `gorm:"column:block_num"`
	Action        string     `gorm:"column:action"`
	Args          *core.Json `gorm:"column:args"`
}

type cMStats struct {
	CreditManager   string       `gorm:"column:credit_manager"`
	Decimals        int8         `gorm:"column:decimals"`
	BlockNum        int64        `gorm:"column:block_num"`
	TotalBorrowed   float64      `gorm:"column:total_borrowed"`
	TotalBorrowedBi *core.BigInt `gorm:"column:total_borrowed_bi"`
}

func main() {
	cfg := config.NewConfig()
	// client := ethclient.NewEthClient(cfg)
	db := repository.NewDBClient(cfg)
	allstats := []*cMStats{}
	err := db.Raw(`select cms.credit_manager, cms.block_num, cms.total_borrowed, t.decimals, cms.total_borrowed_bi from credit_manager_stats cms join credit_managers cm on cm.address=cms.credit_manager join tokens t on t.address=cm.underlying_token where cms.credit_manager in ('0xec76E7652E1B94bd2A4E7A9372F8359Dd571eF4C','0x9A0fDF7CdAb4604FC27ebeab4b3D57BD825e8ebe','0x732F28d627f3F5cfb599A539F58Fa7CBA6698297','0x63Ae843b332DE97c55a007e27c5697c2B8B81627','0xfe83807DeC8C6a4f4D93b7DBD6340771753e2Cd8','0x0Fafa30Cd35bc6a48ff2B40694d4A73d4F4BcC92','0x0aF1324369e3fD78325Fab0CB62EeA19F3e4ebf0','0x79C6C1ce5B12abCC3E407ce8C160eE1160250921','0x8c118E8C20CEbbaa2467b735BBB8B13d614e6608','0xAf5A052BA444Ed90F887D40088548285df33A603','0xbCd2fFaC58189E57334Bb63253AcbF34D776DE53','0xb79d6544839d169869476589d2e54014A074317b') order by block_num`).Find(&allstats).Error
	log.CheckFatal(err)
	log.Info(len(allstats))

	query := `select a.args, a.block_num, cs.credit_manager, a.action from (select args, session_id, block_num, action from account_operations where action like 'Increase%' or action like 'Decrease%') a join credit_sessions cs on cs.id=a.session_id where cs.credit_manager in ('0xec76E7652E1B94bd2A4E7A9372F8359Dd571eF4C','0x9A0fDF7CdAb4604FC27ebeab4b3D57BD825e8ebe','0x732F28d627f3F5cfb599A539F58Fa7CBA6698297','0x63Ae843b332DE97c55a007e27c5697c2B8B81627','0xfe83807DeC8C6a4f4D93b7DBD6340771753e2Cd8','0x0Fafa30Cd35bc6a48ff2B40694d4A73d4F4BcC92','0x0aF1324369e3fD78325Fab0CB62EeA19F3e4ebf0','0x79C6C1ce5B12abCC3E407ce8C160eE1160250921','0x8c118E8C20CEbbaa2467b735BBB8B13d614e6608','0xAf5A052BA444Ed90F887D40088548285df33A603','0xbCd2fFaC58189E57334Bb63253AcbF34D776DE53','0xb79d6544839d169869476589d2e54014A074317b') order by block_num ;`

	events := []Event{}
	err = db.Raw(query).Find(&events).Error
	log.CheckFatal(err)
	eventsI := 0
	borrowed := map[string]*big.Int{}
	for _, cmstats := range allstats {
		if borrowed[cmstats.CreditManager] == nil {
			borrowed[cmstats.CreditManager] = new(big.Int)
		}
		for eventsI < len(events) && events[eventsI].BlockNum <= cmstats.BlockNum {
			data := (*events[eventsI].Args)["amount"]
			a := new(big.Float).SetFloat64(data.(float64))
			amount, _ := a.Int(nil)
			// amount, _ := new(big.Int).SetString(x, 10)
			if events[eventsI].Action[:8] == "Decrease" {
				amount = new(big.Int).Neg(amount)
			}
			borrowed[cmstats.CreditManager] = new(big.Int).Add(borrowed[cmstats.CreditManager], amount)
			eventsI++
		}
		cmstats.TotalBorrowedBi = (*core.BigInt)(borrowed[cmstats.CreditManager])
		cmstats.TotalBorrowed = utils.GetFloat64Decimal(borrowed[cmstats.CreditManager], cmstats.Decimals)
	}

	for ind, stats := range allstats {
		if ind%100 == 0 {
			log.Info(ind, len(allstats))
		}
		log.Info(stats.CreditManager, stats.BlockNum, stats.TotalBorrowed, stats.TotalBorrowedBi)
		err = db.Exec(`update credit_manager_stats set total_borrowed=?, total_borrowed_bi=? where credit_manager=? and block_num=?`, stats.TotalBorrowed, stats.TotalBorrowedBi, stats.CreditManager, stats.BlockNum).Error
		log.CheckFatal(err)
	}
	for cm, amount := range borrowed {
		decimals := &cMStats{}
		err = db.Raw(`select t.decimals from credit_managers cm join tokens t on t.address=cm.underlying_token where cm.address=?`, cm).Find(decimals).Error
		log.CheckFatal(err)
		err = db.Exec(`update credit_managers set total_borrowed_bi=?, total_borrowed=? where address=?`, (*core.BigInt)(amount), utils.GetFloat64Decimal(amount, decimals.Decimals), cm).Error
		log.CheckFatal(err)
	}
}
