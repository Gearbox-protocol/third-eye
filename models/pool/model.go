package pool

import (
	"github.com/Gearbox-protocol/third-eye/artifacts/poolService"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Pool struct {
	*core.SyncAdapter
	contractETH    *poolService.PoolService
	lastEventBlock int64
	State          *core.PoolState
}

func (Pool) TableName() string {
	return "sync_adapters"
}

func NewPool(addr string, client ethclient.ClientI, repo core.RepositoryI, discoveredAt int64) *Pool {
	pool := NewPoolFromAdapter(
		core.NewSyncAdapter(addr, core.Pool, discoveredAt, client, repo),
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
	repo.AddDieselToken(dieselToken.Hex(), underlyingToken.Hex(), addr)
	pool.SetUnderlyingState(&core.PoolState{
		Address:         pool.Address,
		DieselToken:     dieselToken.Hex(),
		UnderlyingToken: underlyingToken.Hex(),
	})
	// create a pool stat snapshot at first log of the pool
	pool.lastEventBlock = pool.DiscoveredAt
	pool.createPoolStat()

	return pool
}

func NewPoolFromAdapter(adapter *core.SyncAdapter) *Pool {
	cmContract, err := poolService.NewPoolService(common.HexToAddress(adapter.Address), adapter.Client)
	if err != nil {
		log.Fatal(err)
	}
	obj := &Pool{
		SyncAdapter: adapter,
		contractETH: cmContract,
	}
	return obj
}

func (mdl *Pool) AfterSyncHook(syncTill int64) {
	mdl.createPoolStat()
	mdl.SyncAdapter.AfterSyncHook(syncTill)
}
