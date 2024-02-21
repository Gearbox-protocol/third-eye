package main

import (
	"sort"
	"strings"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/repository"
	"github.com/Gearbox-protocol/third-eye/services"
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"
)

func main() {
	cfg := config.NewConfig()
	client := ethclient.NewEthClient(cfg)
	//
	parser := services.NewExecuteParser(cfg, client)
	db := repository.NewDBClient(cfg)
	//

	sessions := getExecuteTxHash(db)
	for _, session := range sessions {
		log.Info(session.BlockNumber, session.TxHash)
		ops := []*schemas.AccountOperation{}
		err := db.Raw(`select * from account_operations where dapp !=? and tx_hash = ? and main_action is not null order by log_id `, session.CreditFacade, session.TxHash).Find(&ops).Error
		log.CheckFatal(err)
		//
		params := []ds.ExecuteParams{}
		for _, entry := range ops {
			params = append(params, ds.ExecuteParams{
				SessionId:     entry.SessionId,
				Protocol:      common.HexToAddress(entry.Dapp),
				CreditAccount: common.HexToAddress(strings.Split(entry.SessionId, "_")[0]),
				Borrower:      common.HexToAddress(entry.Borrower),
				Index:         entry.LogId,
				BlockNumber:   entry.BlockNumber,
			})
		}
		x := parser.GetExecuteCalls(core.NewVersion(300), session.TxHash, session.CreditManager, params)
		for ind, execTransfer := range x {
			_ = execTransfer
			log.Info(ops[ind].ID)
			err = db.Exec(`update account_operations set transfers = ? where id=?`, execTransfer.Transfers.String(), ops[ind].ID).Error
			log.CheckFatal(err)
		}
	}
	//
}

type sessAndTx struct {
	SessionId     string `gorm:"column:session_id"`
	TxHash        string `gorm:"column:tx_hash"`
	CreditManager string `gorm:"column:credit_manager"`
	CreditFacade  string `gorm:"column:credit_facade"`
	BlockNumber   int64  `gorm:"column:block_num"`
}

func getExecuteTxHash(db *gorm.DB) []sessAndTx {
	accounts := []sessAndTx{}
	err := db.Raw(`select a.*, c.* 
	from (select  block_num, session_id,   tx_hash from account_operations where dapp not in (select details->>'facade' from sync_adapters where type ='CreditManager' and version!=1)) a 
	join (select id, credit_manager from credit_sessions where version=300) b on b.id = a.session_id 
	join (select address as credit_manager, details->>'facade' as credit_facade from sync_adapters where type ='CreditManager') c on c.credit_manager =b.credit_manager where block_num > 0
	 `).Find(&accounts).Error
	log.CheckFatal(err)
	m := map[string]sessAndTx{}
	for _, account := range accounts {
		m[account.TxHash] = account
	}
	var ss []sessAndTx
	for _, s := range m {
		ss = append(ss, s)
	}
	sort.Slice(ss, func(i, j int) bool { return ss[i].BlockNumber < ss[j].BlockNumber })
	return ss
}
