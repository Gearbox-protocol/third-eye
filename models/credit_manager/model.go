package credit_manager

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/creditFacade"
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditManager"
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditManagerv2"
	"github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressor/dataCompressorv2"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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
	obj := &CreditManager{
		CMv2Fields:      CMv2Fields{},
		SyncAdapter:     adapter,
		UpdatedSessions: make(map[string]int),
		ClosedSessions:  make(map[string]*SessionCloseDetails),
		pnlOnCM:         NewPnlCM(),
	}
	obj.addAdaptersAndReturnDCData(obj.LastSync)
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
		obj.setCreditFacadeSyncer(obj.GetDetailsByKey("facade"), 0)
		// set credit cofigurator syncer
		obj.setConfiguratorSyncer(obj.GetDetailsByKey("configurator"), 0)
	}
	return obj
}

func (mdl *CreditManager) GetUnderlyingDecimal() int8 {
	decimals := mdl.Repo.GetToken(mdl.GetUnderlyingToken()).Decimals
	return decimals
}

func (mdl *CreditManager) AfterSyncHook(syncTill int64) {
	// process remaining v2 events
	for _, txLog := range mdl.getv2ExtraLogs(types.Log{BlockNumber: uint64(syncTill), Index: 10_000_000}) {
		mdl.logHandler(txLog)
	}
	// ON NEW TXHASH
	mdl.onTxHash("") // handles for v1(for multicalls) and v1 (for executeorder)

	// no logs where detected for current sync
	// no need to explicitly call ProcessDirectToken, it will be called by onBlockChange
	// if mdl.lastEventBlock == 0 {
	// 	mdl.ProcessDirectTokenTransfer(mdl.GetLastSync()+1, syncTill+1)
	// }
	//
	//
	// ON NEW BLOCKNUM
	//
	// try with blocknum greater than syncTill
	// so that if there is direct transfer and some credit manager event
	// at synctill == mdl.LasteventBlock it is processed
	mdl.onBlockChange(syncTill + 1)
	mdl.SyncAdapter.AfterSyncHook(syncTill)
}

func (cm *CreditManager) GetCreditSessionData(blockNum int64, borrower string) *dataCompressorv2.CreditAccountData {
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
	}
	data, err := cm.Repo.GetDCWrapper().GetCreditAccountData(opts,
		common.HexToAddress(cm.GetAddress()),
		common.HexToAddress(borrower),
	)
	// TODO: later detect if the test adapter is used
	// check is added as hack func is called in kovan https://kovan.etherscan.io/tx/0x2e9c3c8c55cd9817c996ffb3d8afeff59754e7370ce4df152b51e1124b741cb7
	// for addressProvider: 0xA526311C39523F60b184709227875b5f34793bD4
	// 0xeE5998268707e9d57Ab1156b3A87cD7476274362 is a test account
	// check if the dc call is failing due to totalvalue being zero
	//
	// if err != nil && err.Error() == "VM execution error." {
	// 	// variables are shadowed on purpose
	// 	// so that outer error is preserved
	// 	data, err := cm.Repo.GetDCWrapper().GetCreditAccountDataExtendedForHack(opts,
	// 		common.HexToAddress(cm.GetAddress()),
	// 		common.HexToAddress(borrower),
	// 	)
	// 	log.CheckFatal(err)
	// 	if data.TotalValue.Cmp(big.NewInt(0)) == 0 {
	// 		return data
	// 	}
	// }
	if err != nil {
		log.Fatalf("For blockNum %d CM:%s Borrower:%s %s", blockNum, cm.GetAddress(), borrower, err)
	}
	return &data
}
