package core

import (
	"github.com/Gearbox-protocol/gearscan/utils"	
)

type EngineI interface {
	Sync()
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
}
