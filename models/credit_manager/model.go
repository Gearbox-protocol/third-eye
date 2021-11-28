package credit_manager

import (
	"github.com/Gearbox-protocol/gearscan/core"
	"github.com/Gearbox-protocol/gearscan/ethclient"
	"github.com/Gearbox-protocol/gearscan/log"
	"github.com/Gearbox-protocol/gearscan/models/credit_filter"
	"github.com/ethereum/go-ethereum/common"
	"github.com/Gearbox-protocol/gearscan/artifacts/creditManager"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
)

type CreditManager struct {
	*core.SyncAdapter
	*core.State
	contractETH *creditManager.CreditManager
	LastTxHash string
	UToken string
	UDecimals int64
}

func NewCreditManager(addr string, client *ethclient.Client, repo core.RepositoryI, discoveredAt int64) *CreditManager {
	cmContract, err:=creditManager.NewCreditManager(common.HexToAddress(addr), client)
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
		Address:addr, 
		PoolAddress:poolAddr.Hex(), 
		UnderlyingToken: underlyingToken.Hex(), 
		Sessions: core.NewHstore(),
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
	cmContract, err:=creditManager.NewCreditManager(common.HexToAddress(adapter.Address), adapter.Client)
	if err != nil {
		log.Fatal(err)
	}
	obj := &CreditManager{
		SyncAdapter: adapter,
		State: &core.State{Repo: repo},
		contractETH: cmContract,
	}
	obj.GetAbi()
	obj.SetUToken()
	return obj
}

func (mdl *CreditManager) SetUToken() {
	if mdl.UToken == "" {
		mdl.UToken = mdl.Repo.GetUnderlyingToken(mdl.Address)
		mdl.UDecimals = int64(mdl.Repo.GetToken(mdl.UToken).Decimals)
	}
}