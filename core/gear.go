package core

import (
	"github.com/Gearbox-protocol/third-eye/artifacts/dataCompressor"
	"github.com/Gearbox-protocol/third-eye/services"
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
	// getting all the adapters for syncing in the engine
	GetSyncAdapters() map[string]SyncAdapterI
	// adding the adapters for syncing
	AddSyncAdapter(adapterI SyncAdapterI)
	// saving to the db
	Flush() error
	// adding block/timestamp
	SetBlock(blockNum int64)
	// credit account operations
	AddAccountOperation(accountOperation *AccountOperation)
	// for getting executeparser
	GetExecuteParser() *services.ExecuteParser
	// price feed/oracle funcs
	AddTokenOracle(token, oracle string, blockNum int64)
	AddPriceFeed(blockNum int64, pf *PriceFeed)
	// token funcs
	AddAllowedProtocol(p *Protocol)
	AddToken(token string)
	AddAllowedToken(atoken *AllowedToken)
	AddTokenObj(token *Token)
	AddDataCompressor(blockNum int64, addr string)
	GetToken(addr string) *Token
	// credit session funcs
	AddCreditSession(session *CreditSession)
	GetCreditSession(sessionId string) *CreditSession
	// credit session snapshots funcs
	AddCreditSessionSnapshot(css *CreditSessionSnapshot)
	AddLastCSS(css *CreditSessionSnapshot)
	GetLastCSS(sessionId string) *CreditSessionSnapshot
	GetDataCompressor(blockNum int64) *dataCompressor.DataCompressor
	// pools
	AddPoolStat(ps *PoolStat)
	AddPoolLedger(pl *PoolLedger)
	GetPoolUniqueUserLen(pool string) int
	// credit manager
	AddCreditManagerStats(cms *CreditManagerStat)
}
