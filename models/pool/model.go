package pool

import (
	"github.com/Gearbox-protocol/gearscan/artifacts/poolService"
	"github.com/Gearbox-protocol/gearscan/core"
	"github.com/Gearbox-protocol/gearscan/ethclient"
	"github.com/Gearbox-protocol/gearscan/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type Pool struct {
	*core.SyncAdapter
	*core.State
	contractETH *poolService.PoolService
}

func NewPool(addr string, client *ethclient.Client, repo core.RepositoryI, discoveredAt int64) *Pool {
	pool := NewPoolFromAdapter(
		repo,
		core.NewSyncAdapter(addr, "Pool", discoveredAt, client),
	)
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(pool.DiscoveredAt),
	}
	underlyingToken, err := pool.contractETH.UnderlyingToken(opts)
	if err != nil {
		log.Fatal(err)
	}
	repo.AddToken(underlyingToken.Hex())
	dieselToken, err := pool.contractETH.DieselToken(opts)
	if err != nil {
		log.Fatal(err)
	}
	repo.AddToken(dieselToken.Hex())
	repo.AddPool(&core.Pool{
		Address:         pool.Address,
		DieselToken:     dieselToken.Hex(),
		UnderlyingToken: underlyingToken.Hex(),
	})

	return pool
}

func NewPoolFromAdapter(repo core.RepositoryI, adapter *core.SyncAdapter) *Pool {
	cmContract, err := poolService.NewPoolService(common.HexToAddress(adapter.Address), adapter.Client)
	if err != nil {
		log.Fatal(err)
	}
	obj := &Pool{
		SyncAdapter: adapter,
		State:       &core.State{Repo: repo},
		contractETH: cmContract,
	}
	return obj
}

func (mdl *Pool) OnLog(txLog types.Log) {
}
