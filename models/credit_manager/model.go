package credit_manager

import (
	"github.com/Gearbox-protocol/third-eye/artifacts/creditManager"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/models/credit_filter"
	"github.com/Gearbox-protocol/third-eye/services"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"sort"
)

type CreditManager struct {
	*core.SyncAdapter
	*core.State
	contractETH   *creditManager.CreditManager
	LastTxHash    string
	UToken        string
	UDecimals     uint8
	executeParams []services.ExecuteParams
	eventBalances SortedEventbalances
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
	// create credit manager for db
	repo.AddCreditManager(&core.CreditManager{
		Address:         addr,
		PoolAddress:     poolAddr.Hex(),
		UnderlyingToken: underlyingToken.Hex(),
		Sessions:        core.NewHstore(),
	})

	// create creditFilter syncadapter
	creditFilter, err := cmContract.CreditFilter(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	cf := credit_filter.NewCreditFilter(creditFilter.Hex(), addr, discoveredAt, client, repo)
	repo.AddSyncAdapter(cf)
	//

	cm := NewCreditManagerFromAdapter(
		repo,
		core.NewSyncAdapter(addr, "CreditManager", discoveredAt, client),
	)
	return cm
}

func NewCreditManagerFromAdapter(repo core.RepositoryI, adapter *core.SyncAdapter) *CreditManager {
	cmContract, err := creditManager.NewCreditManager(common.HexToAddress(adapter.Address), adapter.Client)
	if err != nil {
		log.Fatal(err)
	}
	obj := &CreditManager{
		SyncAdapter: adapter,
		State:       &core.State{Repo: repo},
		contractETH: cmContract,
	}
	obj.GetAbi()
	obj.SetUToken()
	return obj
}

func (mdl *CreditManager) SetUToken() {
	if mdl.UToken == "" {
		mdl.UToken = mdl.Repo.GetUnderlyingToken(mdl.Address)
		mdl.UDecimals = mdl.Repo.GetToken(mdl.UToken).Decimals
	}
}

func (mdl *CreditManager) AfterSyncHook(syncTill int64) {
	mdl.processExecuteEvents()
	sort.Sort(mdl.eventBalances)
	for _, eventBalance := range mdl.eventBalances {
		mdl.updateBalance(eventBalance)
	}
	mdl.eventBalances = SortedEventbalances{}
	mdl.SetLastSync(syncTill)
}
