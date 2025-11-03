package main

import (
	"math/big"
	"os"
	"strconv"

	"github.com/Gearbox-protocol/sdk-go/artifacts/poolv3"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/repository"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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
	Pool            string       `gorm:"column:pool_address"`
}

func main() {
	cfg := config.NewConfig()
	// client := ethclient.NewEthClient(cfg)
	db := repository.NewDBClient(cfg)
	allstats := []*cMStats{}
	blockNum, err := strconv.ParseInt(os.Args[1], 10, 64)
	log.CheckFatal(err)

	err = db.Raw(`select cms.credit_manager, cms.block_num, cms.total_borrowed, cm.pool_address, t.decimals, cms.total_borrowed_bi from credit_manager_stats cms join credit_managers cm on cm.address=cms.credit_manager join tokens t on t.address=cm.underlying_token  where _version=300 and block_num > ? order by block_num`, blockNum).Find(&allstats).Error
	log.CheckFatal(err)
	log.Info(len(allstats))

	// query := `select a.args, a.block_num, cs.credit_manager, a.action from (select args, session_id, block_num, action from account_operations where action like 'Increase%' or action like 'Decrease%') a join credit_sessions cs on cs.id=a.session_id where cs.credit_manager in ('0xec76E7652E1B94bd2A4E7A9372F8359Dd571eF4C','0x9A0fDF7CdAb4604FC27ebeab4b3D57BD825e8ebe','0x732F28d627f3F5cfb599A539F58Fa7CBA6698297','0x63Ae843b332DE97c55a007e27c5697c2B8B81627','0xfe83807DeC8C6a4f4D93b7DBD6340771753e2Cd8','0x0Fafa30Cd35bc6a48ff2B40694d4A73d4F4BcC92','0x0aF1324369e3fD78325Fab0CB62EeA19F3e4ebf0','0x79C6C1ce5B12abCC3E407ce8C160eE1160250921','0x8c118E8C20CEbbaa2467b735BBB8B13d614e6608','0xAf5A052BA444Ed90F887D40088548285df33A603','0xbCd2fFaC58189E57334Bb63253AcbF34D776DE53','0xb79d6544839d169869476589d2e54014A074317b') order by block_num ;`

	borrowedMap := map[string]*big.Int{}
	client := ethclient.NewEthClient(cfg)
	for ind, stat := range allstats {
		pool, err := poolv3.NewPoolv3(common.HexToAddress(stat.Pool), client)
		log.CheckFatal(err)
		data, err := pool.CreditManagerBorrowed(&bind.CallOpts{BlockNumber: big.NewInt(stat.BlockNum)}, common.HexToAddress(stat.CreditManager))
		log.CheckFatal(err)
		borrowedBI := (*core.BigInt)(data)
		borrowedMap[stat.CreditManager] = data
		err = db.Exec(`update credit_manager_stats set total_borrowed=?, total_borrowed_bi=? where credit_manager=? and block_num=?`, utils.GetFloat64Decimal(borrowedBI, stat.Decimals), borrowedBI, stat.CreditManager, stat.BlockNum).Error
		log.CheckFatal(err)
		log.Info(stat.BlockNum, ind, ":", len(allstats))
	}

	for cm, amount := range borrowedMap {
		decimals := &cMStats{}
		err = db.Raw(`select t.decimals from credit_managers cm join tokens t on t.address=cm.underlying_token where cm.address=?`, cm).Find(decimals).Error
		log.CheckFatal(err)
		err = db.Exec(`update credit_managers set total_borrowed_bi=?, total_borrowed=? where address=?`, (*core.BigInt)(amount), utils.GetFloat64Decimal(amount, decimals.Decimals), cm).Error
		log.CheckFatal(err)
	}
}
