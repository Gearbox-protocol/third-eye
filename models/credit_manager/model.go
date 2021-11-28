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
}

func NewCreditManager(addr string, client *ethclient.Client, repo core.RepositoryI, discoveredAt int64) *CreditManager {
	cm := NewCreditManagerFromAdapter(
		repo,
		core.NewSyncAdapter(addr, "CreditManager", discoveredAt, client),
	)
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(cm.DiscoveredAt),
	}
	poolAddr, err := cm.contractETH.PoolService(opts)
	if err != nil {
		log.Fatal(err)
	}
	underlyingToken, err := cm.contractETH.UnderlyingToken(opts)
	if err != nil {
		log.Fatal(err)
	}
	repo.AddToken(underlyingToken.Hex())
	cm.Repo.AddCreditManager(&core.CreditManager{Address:cm.Address, PoolAddress:poolAddr.Hex(), UnderlyingToken: underlyingToken.Hex(), Sessions: core.NewHstore()})

	creditFilter, err := cm.contractETH.CreditFilter(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	cf := credit_filter.NewCreditFilter(creditFilter.Hex(), addr, discoveredAt, client, repo)
	repo.AddSyncAdapter(cf)

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
	return obj
}