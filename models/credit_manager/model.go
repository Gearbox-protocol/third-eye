package credit_manager

import (
	"github.com/Gearbox-protocol/third-eye/artifacts/creditManager"
	"github.com/Gearbox-protocol/third-eye/artifacts/dataCompressor"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/models/credit_filter"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type SessionCloseDetails struct {
	RemainingFunds *big.Int
	Status         int
	LogId          uint
}
type CreditManager struct {
	*core.SyncAdapter
	contractETH     *creditManager.CreditManager
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

func NewCreditManager(addr string, client *ethclient.Client, repo core.RepositoryI, discoveredAt int64) *CreditManager {
	cmContract, err := creditManager.NewCreditManager(common.HexToAddress(addr), client)
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(discoveredAt),
	}
	// create underlying token
	underlyingToken, err := cmContract.UnderlyingToken(opts)
	if err != nil {
		log.Fatal(err)
	}
	repo.AddToken(underlyingToken.Hex())
	//
	poolAddr, err := cmContract.PoolService(opts)
	if err != nil {
		log.Fatal(err)
	}

	// create creditFilter syncadapter
	creditFilter, err := cmContract.CreditFilter(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	cf := credit_filter.NewCreditFilter(creditFilter.Hex(), addr, discoveredAt, client, repo)
	repo.AddSyncAdapter(cf)
	//

	cm := NewCreditManagerFromAdapter(
		core.NewSyncAdapter(addr, core.CreditManager, discoveredAt, client, repo),
	)
	// create credit manager state
	cm.SetUnderlyingState(&core.CreditManagerState{
		Address:         addr,
		PoolAddress:     poolAddr.Hex(),
		UnderlyingToken: underlyingToken.Hex(),
		Sessions:        core.NewHstore(),
	})
	return cm
}

func NewCreditManagerFromAdapter(adapter *core.SyncAdapter) *CreditManager {
	cmContract, err := creditManager.NewCreditManager(common.HexToAddress(adapter.Address), adapter.Client)
	if err != nil {
		log.Fatal(err)
	}
	obj := &CreditManager{
		SyncAdapter:     adapter,
		contractETH:     cmContract,
		UpdatedSessions: make(map[string]int),
		ClosedSessions:  make(map[string]*SessionCloseDetails),
	}
	obj.GetAbi()
	return obj
}

func (mdl *CreditManager) GetUnderlyingDecimal() int8 {
	decimals := mdl.Repo.GetToken(mdl.GetUnderlyingToken()).Decimals
	return decimals
}

func (mdl *CreditManager) AfterSyncHook(syncTill int64) {
	// generate remaining accountoperations and operation state
	mdl.processExecuteEvents()
	mdl.onBlockChange()
	mdl.SetLastSync(syncTill)
}

func (cm *CreditManager) GetCreditSessionData(blockNum int64, borrower string) *dataCompressor.DataTypesCreditAccountDataExtended {
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
	}
	data, err := cm.Repo.GetDataCompressor(blockNum).GetCreditAccountDataExtended(opts,
		common.HexToAddress(cm.GetAddress()),
		common.HexToAddress(borrower),
	)
	if err != nil {
		log.Fatalf("CM:%s Borrower:%s %s", cm.GetAddress(), borrower, err)
	}
	return &data
}
