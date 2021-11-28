package core

import (
	"github.com/Gearbox-protocol/gearscan/utils"
	"github.com/Gearbox-protocol/gearscan/artifacts/dataCompressor"
)

type EngineI interface {
	Sync()
}

type Protocol struct {
	Protocol string `gorm:"column:protocol"`
	Adapter string `gorm:"column:adapter"`
	BlockNumber int64 `gorm:"primaryKey;column:block_num"`
	CreditManager string `gorm:"primaryKey;column:credit_manager"`
}

func (Protocol) TableName() string {
	return "allowed_protocols"
}

type RepositoryI interface {
	GetSyncAdapters() []SyncAdapterI
	AddSyncAdapter(adapterI SyncAdapterI)
	Flush() error
	SetBlock(blockNum int64)
	AddAccountOperation(accountOperation *AccountOperation)
	AddCreditManager(cm *CreditManager)
	AddCreditOwnerSession(cmAddr, owner, sessionId string)
	RemoveCreditOwnerSession(cmAddr, owner string)
	GetCreditOwnerSession(cmAddr, owner string) string
	GetExecuteParser() *utils.ExecuteParser
	AddTokenOracle(token, oracle string, blockNum int64)
	AddPriceFeed(blockNum int64, pf *PriceFeed)
	AddAllowedProtocol(p *Protocol)
	AddToken(token string)
	AddAllowedToken(atoken *AllowedToken)
	AddTokenObj(token *Token)
	AddPool(pool *Pool)
	AddDataCompressor(addr string)
	AddCreditSession(session *CreditSession)
	GetCreditSession(sessionId string)*CreditSession
	GetCreditSessionData(blockNum int64, sessionId string) dataCompressor.DataTypesCreditAccountDataExtended
}
