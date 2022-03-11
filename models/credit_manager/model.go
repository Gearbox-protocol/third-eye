package credit_manager

import (
	"github.com/Gearbox-protocol/third-eye/artifacts/creditManager"
	"github.com/Gearbox-protocol/third-eye/artifacts/creditManagerv2"
	"github.com/Gearbox-protocol/third-eye/artifacts/creditFacade"
	"github.com/Gearbox-protocol/third-eye/artifacts/dataCompressor/mainnet"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type SessionCloseDetails struct {
	RemainingFunds   *big.Int
	Status           int
	LogId            uint
	TxHash           string
	Borrower         string
	AccountOperation *core.AccountOperation
}

type CreditManager struct {
	*core.SyncAdapter
	contractETHV1     *creditManager.CreditManager
	contractETHV2     *creditManagerv2.CreditManagerv2
	facadeContractV2     *creditFacade.CreditFacade
	LastTxHash      string
	executeParams   []core.ExecuteParams
	State           *core.CreditManagerState
	lastEventBlock  int64
	UpdatedSessions map[string]int
	ClosedSessions  map[string]*SessionCloseDetails
}

func (CreditManager) TableName() string {
	return "sync_adapters"
}

func NewCreditManager(addr string, client ethclient.ClientI, repo core.RepositoryI, discoveredAt int64) *CreditManager {
	// get version
	cmContract, err := creditManager.NewCreditManager(common.HexToAddress(addr), client)
	log.CheckFatal(err)
	version, err := cmContract.Version(&bind.CallOpts{})
	if err != nil {
		version = big.NewInt(1)
	}

	// credit manager
	adapter := core.NewSyncAdapter(addr, core.CreditManager, discoveredAt, client, repo)
	adapter.SetVersion(version.Int64())
	cm := NewCreditManagerFromAdapter(
		adapter,
	)
	cm.CommonInit()
	switch version.Int64() {
	case 1:
		cm.addCreditFilter()
	case 2:
		cm.addCreditConfigurator()
	}
	return cm
}

func NewCreditManagerFromAdapter(adapter *core.SyncAdapter) *CreditManager {
	obj := &CreditManager{
		SyncAdapter:     adapter,
		UpdatedSessions: make(map[string]int),
		ClosedSessions:  make(map[string]*SessionCloseDetails),
	}
	obj.GetAbi()
	switch obj.GetVersion() {
	case 1:
		cmContract, err := creditManager.NewCreditManager(common.HexToAddress(adapter.Address), adapter.Client)
		if err != nil {
			log.Fatal(err)
		}
		obj.contractETHV1 = cmContract
	case 2:
		// set credit manager and credit facade contracts
		cmContract, err := creditManagerv2.NewCreditManagerv2(common.HexToAddress(adapter.Address), adapter.Client)
		if err != nil {
			log.Fatal(err)
		}
		obj.contractETHV2 = cmContract
		var creditFacadeAddr common.Address
		if obj.Details != nil && obj.Details["creditFacade"] != nil {
			creditFacadeAddr = common.HexToAddress(obj.Details["creditFacade"].(string))
		} else {
			creditFacadeAddr, err = cmContract.CreditFacade(&bind.CallOpts{})
			log.CheckFatal(err)
		}
		obj.facadeContractV2, err = creditFacade.NewCreditFacade(creditFacadeAddr, adapter.Client)
		log.CheckFatal(err)
		if obj.Details == nil {
			obj.Details = map[string]interface{}{}
		}
		obj.Details["creditFacade"] = creditFacadeAddr.Hex()
	}
	return obj
}

func (mdl *CreditManager) GetUnderlyingDecimal() int8 {
	decimals := mdl.Repo.GetToken(mdl.GetUnderlyingToken()).Decimals
	return decimals
}

func (mdl *CreditManager) AfterSyncHook(syncTill int64) {
	// generate remaining accountoperations and operation state
	mdl.processExecuteEvents()
	// no logs where detected for current sync
	if mdl.lastEventBlock == 0 {
		mdl.ProcessDirectTokenTransfer(mdl.GetLastSync()+1, syncTill+1)
	}
	// try with blocknum greater than syncTill
	// so that if there is direct transfer and some credit manager event
	// at synctill == mdl.LasteventBlock it is processed
	mdl.onBlockChange(syncTill + 1)
	mdl.SyncAdapter.AfterSyncHook(syncTill)
}

func (cm *CreditManager) GetCreditSessionData(blockNum int64, borrower string) *mainnet.DataTypesCreditAccountDataExtended {
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
	}
	data, err := cm.Repo.GetDCWrapper().GetCreditAccountDataExtended(opts,
		common.HexToAddress(cm.GetAddress()),
		common.HexToAddress(borrower),
	)
	if err != nil {
		log.Fatalf("CM:%s Borrower:%s %s", cm.GetAddress(), borrower, err)
	}
	return &data
}
