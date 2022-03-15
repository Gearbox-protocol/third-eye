package credit_manager

import (
	"github.com/Gearbox-protocol/third-eye/artifacts/creditFacade"
	"github.com/Gearbox-protocol/third-eye/artifacts/creditManager"
	"github.com/Gearbox-protocol/third-eye/artifacts/creditManagerv2"
	"github.com/Gearbox-protocol/third-eye/artifacts/dataCompressor/mainnet"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Collateral struct {
	Amount *big.Int
	Token  string
}
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
	contractETHV1    *creditManager.CreditManager
	contractETHV2    *creditManagerv2.CreditManagerv2
	facadeContractV2 *creditFacade.CreditFacade
	LastTxHash       string
	executeParams    []core.ExecuteParams
	State            *core.CreditManagerState
	lastEventBlock   int64
	UpdatedSessions  map[string]int
	ClosedSessions   map[string]*SessionCloseDetails
	// borrower to events, these events have same txHash
	multicall MultiCallProcessor
}

func (CreditManager) TableName() string {
	return "sync_adapters"
}

func NewCreditManager(addr string, client ethclient.ClientI, repo core.RepositoryI, discoveredAt int64) *CreditManager {
	// credit manager
	adapter := core.NewSyncAdapter(addr, core.CreditManager, discoveredAt, client, repo)
	version := adapter.FetchVersion(0)
	adapter.SetVersion(version)
	cm := NewCreditManagerFromAdapter(
		adapter,
	)
	cm.CommonInit()
	switch version {
	case 1:
		cm.addCreditFilter(discoveredAt)
	case 2:
		cm.addCreditConfigurator()
	}
	return cm
}
func (mdl *CreditManager) GetAbi() {
	switch mdl.GetVersion() {
	case 1:
		mdl.ABI = core.GetAbi(mdl.ContractName)
	case 2:
		mdl.ABI = core.GetAbi("CreditFacade")
	}
}

func NewCreditManagerFromAdapter(adapter *core.SyncAdapter) *CreditManager {
	obj := &CreditManager{
		SyncAdapter:     adapter,
		UpdatedSessions: make(map[string]int),
		ClosedSessions:  make(map[string]*SessionCloseDetails),
		multicall:       MultiCallProcessor{},
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
			// should only be called on discovered, not when loading form db.
			opts := &bind.CallOpts{BlockNumber: big.NewInt(adapter.DiscoveredAt)}
			creditFacadeAddr, err = cmContract.CreditFacade(opts)
			log.CheckFatal(err)
			obj.SetCreditFacade(creditFacadeAddr)
			var creditConfigurator common.Address
			creditConfigurator, err = cmContract.CreditConfigurator(opts)
			log.CheckFatal(err)
			obj.Details["configurator"] = creditConfigurator.Hex()
		}

	}
	return obj
}

func (mdl *CreditManager) SetCreditFacade(creditFacadeAddr common.Address) {
	var err error
	mdl.facadeContractV2, err = creditFacade.NewCreditFacade(creditFacadeAddr, mdl.Client)
	log.CheckFatal(err)
	if mdl.Details == nil {
		mdl.Details = map[string]interface{}{}
	}
	mdl.Details["facade"] = creditFacadeAddr.Hex()
}

func (mdl *CreditManager) GetCreditFacadeAddr() string {
	return mdl.GetDetailsByKey("facade")
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
