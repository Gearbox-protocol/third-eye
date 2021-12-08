package core

import (
	"github.com/Gearbox-protocol/third-eye/artifacts/dataCompressor"
)

type EngineI interface {
	SyncHandler()
}

type Protocol struct {
	Id            string `gorm:"primaryKey;column:id;autoincrement:true"`
	Protocol      string `gorm:"column:protocol"`
	Adapter       string `gorm:"column:adapter"`
	BlockNumber   int64  `gorm:"column:block_num"`
	CreditManager string `gorm:"column:credit_manager"`
}

func (Protocol) TableName() string {
	return "allowed_protocols"
}

type RepositoryI interface {
	// sync adapters
	GetKit() *AdapterKit
	AddSyncAdapter(adapterI SyncAdapterI)
	// saving to the db
	Flush() error
	// adding block/timestamp
	SetBlock(blockNum int64)
	// credit account operations
	AddAccountOperation(accountOperation *AccountOperation)
	// for getting executeparser
	GetExecuteParser() ExecuteParserI
	// price feed/oracle funcs
	AddTokenOracle(token, oracle, feed string, blockNum int64)
	AddPriceFeed(blockNum int64, pf *PriceFeed)
	// token funcs
	AddAllowedProtocol(p *Protocol)
	AddToken(token string) *Token
	AddAllowedToken(atoken *AllowedToken)
	AddTokenObj(token *Token)
	AddDataCompressor(blockNum int64, addr string)
	GetToken(addr string) *Token
	// credit session funcs
	AddCreditSession(session *CreditSession)
	GetCreditSession(sessionId string) *CreditSession
	// credit session snapshots funcs
	AddCreditSessionSnapshot(css *CreditSessionSnapshot)
	GetLastCSS(sessionId string) *CreditSessionSnapshot
	GetDataCompressor(blockNum int64) *dataCompressor.DataCompressor
	AddEventBalance(eb EventBalance)
	CalculateDebt()
	// pools
	AddPoolStat(ps *PoolStat)
	AddPoolLedger(pl *PoolLedger)
	GetPoolUniqueUserLen(pool string) int
	// credit manager
	AddCreditManagerStats(cms *CreditManagerStat)
}
