package credit_manager

import (
	"fmt"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/creditFacade"
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditManager"
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditManagerv2"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
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

func (x SessionCloseDetails) String() string {
	return fmt.Sprintf("ClosingDetails(Status: %d LogId %d TxHash %s Borrower %s RemainingFunds %s)",
		x.Status, x.LogId, x.TxHash, x.Borrower, x.RemainingFunds)
}

type CreditManager struct {
	*ds.SyncAdapter
	CMv2Fields
	//
	contractETHV1 *creditManager.CreditManager
	//
	State *schemas.CreditManagerState
	//
	LastTxHash      string
	lastEventBlock  int64
	executeParams   []ds.ExecuteParams
	pnlOnCM         *PnlCM
	UpdatedSessions map[string]int
	ClosedSessions  map[string]*SessionCloseDetails
	//
	// tmp storage
	borrowedAmountForBlock *big.Int
	params                 *schemas.Parameters
	//
	allowedProtocols map[string]bool
	// only used for testing, in reward_claimed_test.go
	dontGetSessionFromDC bool
}

func (CreditManager) TableName() string {
	return "sync_adapters"
}

func NewCreditManager(addr string, client core.ClientI, repo ds.RepositoryI, discoveredAt int64) *CreditManager {
	// credit manager
	mdl := NewCreditManagerFromAdapter(
		ds.NewSyncAdapter(addr, ds.CreditManager, discoveredAt, client, repo),
	)
	mdl.CommonInit(mdl.GetVersion())
	switch mdl.GetVersion() {
	case 1:
		mdl.addCreditFilterAdapter(discoveredAt)
	case 2:
		mdl.addCreditConfiguratorAdapter(mdl.GetDetailsByKey("configurator"))
		// if params was updated before credit manager was added to addressprovider
		// sync credit configurator to get params
		// mdl.configuratorSyncer.FetchLogs(0, mdl.LastSync)
	}
	return mdl
}
func (mdl *CreditManager) GetAbi() {
	switch mdl.GetVersion() {
	case 1:
		mdl.ABI = core.GetAbi(mdl.ContractName)
	case 2:
		mdl.ABI = core.GetAbi("CreditFacade")
	}
}

func NewCreditManagerFromAdapter(adapter *ds.SyncAdapter) *CreditManager {
	//
	obj := &CreditManager{
		CMv2Fields:      CMv2Fields{},
		SyncAdapter:     adapter,
		UpdatedSessions: make(map[string]int),
		ClosedSessions:  make(map[string]*SessionCloseDetails),
		pnlOnCM:         NewPnlCM(),
	}
	// cm is registered with dataCompressor after discoveredAt, so we can get adapters for blockNum more than discoveredAt
	blockToFetchCMData := obj.DiscoveredAt
	if blockToFetchCMData < obj.LastSync {
		blockToFetchCMData = obj.LastSync
	}
	obj.addProtocolAdaptersLocally(blockToFetchCMData)
	obj.GetAbi()
	switch obj.GetVersion() {
	case 1:
		cmContract, err := creditManager.NewCreditManager(common.HexToAddress(adapter.Address), adapter.Client)
		if err != nil {
			log.Fatal(err)
		}
		obj.contractETHV1 = cmContract
	case 2:
		// set credit manager
		cmContract, err := creditManagerv2.NewCreditManagerv2(common.HexToAddress(adapter.Address), adapter.Client)
		if err != nil {
			log.Fatal(err)
		}
		obj.contractETHV2 = cmContract
		// set credit facade contract
		obj.facadeContractV2, err = creditFacade.NewCreditFacade(core.NULL_ADDR, nil)
		log.CheckFatal(err)

		// set facade and configurator in map
		obj.setv2AddrIfNotPresent()
		// credit facade syncer
		obj.setCreditFacadeSyncer(obj.GetDetailsByKey("facade"))
		// set credit cofigurator syncer
		obj.setConfiguratorSyncer(obj.GetDetailsByKey("configurator"))
	}
	return obj
}

func (mdl *CreditManager) GetUnderlyingDecimal() int8 {
	decimals := mdl.Repo.GetToken(mdl.GetUnderlyingToken()).Decimals
	return decimals
}

func (mdl *CreditManager) AfterSyncHook(syncTill int64) {
	mdl.SyncAdapter.AfterSyncHook(syncTill)
}
