package credit_manager

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/creditFacade"
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditManager"
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditManagerv2"
	"github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressor/mainnet"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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
	AccountOperation *schemas.AccountOperation
}

type CreditManager struct {
	*ds.SyncAdapter
	contractETHV1    *creditManager.CreditManager
	contractETHV2    *creditManagerv2.CreditManagerv2
	facadeContractV2 *creditFacade.CreditFacade
	LastTxHash       string
	executeParams    []ds.ExecuteParams
	State            *schemas.CreditManagerState
	lastEventBlock   int64
	UpdatedSessions  map[string]int
	ClosedSessions   map[string]*SessionCloseDetails
	// borrower to events, these events have same txHash
	multicall MultiCallProcessor
}

func (CreditManager) TableName() string {
	return "sync_adapters"
}

func NewCreditManager(addr string, client core.ClientI, repo ds.RepositoryI, discoveredAt int64) *CreditManager {
	// credit manager
	cm := NewCreditManagerFromAdapter(
		ds.NewSyncAdapter(addr, ds.CreditManager, discoveredAt, client, repo),
	)
	cm.CommonInit(cm.GetVersion())
	switch cm.GetVersion() {
	case 1:
		cm.addCreditFilter(discoveredAt)
	case 2:
		creditConfigurator, err := cm.contractETHV2.CreditConfigurator(&bind.CallOpts{BlockNumber: big.NewInt(discoveredAt)})
		if err != nil {
			log.Fatal(err)
		}
		cm.addCreditConfigurator(creditConfigurator.Hex())
	}
	return cm
}
func (mdl *CreditManager) GetAbi() {
	switch mdl.GetVersion() {
	case 1:
		mdl.ABI = schemas.GetAbi(mdl.ContractName)
	case 2:
		mdl.ABI = schemas.GetAbi("CreditFacade")
	}
}

func NewCreditManagerFromAdapter(adapter *ds.SyncAdapter) *CreditManager {
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
		if obj.Details != nil && obj.Details["facade"] != nil {
			creditFacadeAddr = common.HexToAddress(obj.Details["facade"].(string))
		} else {
			opts := &bind.CallOpts{BlockNumber: big.NewInt(adapter.DiscoveredAt)}
			creditFacadeAddr, err = cmContract.CreditFacade(opts)
			log.CheckFatal(err)
		}
		obj.SetCreditFacadeContract(creditFacadeAddr)
	}
	return obj
}

func (mdl *CreditManager) SetCreditFacadeContract(creditFacadeAddr common.Address) {
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
	// ON NEW TXHASH
	mdl.onNewTxHashV2("")
	// generate remaining accountoperations and operation state
	mdl.processExecuteEvents()
	// ON NEW BLOCKNUM
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
	var err error
	var data mainnet.DataTypesCreditAccountDataExtended
	// TODO: later detect if the test adapter is used
	// check is added as hack func is called in kovan https://kovan.etherscan.io/tx/0x2e9c3c8c55cd9817c996ffb3d8afeff59754e7370ce4df152b51e1124b741cb7
	// for addressProvider: 0xA526311C39523F60b184709227875b5f34793bD4
	if borrower == "0xeE5998268707e9d57Ab1156b3A87cD7476274362" {
		data, err = cm.Repo.GetDCWrapper().GetCreditAccountDataExtendedForHack(opts,
			common.HexToAddress(cm.GetAddress()),
			common.HexToAddress(borrower),
		)
	} else {
		data, err = cm.Repo.GetDCWrapper().GetCreditAccountDataExtended(opts,
			common.HexToAddress(cm.GetAddress()),
			common.HexToAddress(borrower),
		)
	}
	if err != nil {
		log.Fatalf("For blockNum %d CM:%s Borrower:%s %s", blockNum, cm.GetAddress(), borrower, err)
	}
	return &data
}
