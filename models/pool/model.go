package pool

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/poolService"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type Pool struct {
	*ds.SyncAdapter
	contractETH    *poolService.PoolService
	lastEventBlock int64
	State          *schemas.PoolState
	// for calculating the remove liq amount in underlying instead of diesel token amount which is present in the emitted removeLiquidity event
	RemoveLiqEvents []*schemas.PoolLedger
	dieselRate      *big.Int
}

func (Pool) TableName() string {
	return "sync_adapters"
}

func NewPool(addr string, client core.ClientI, repo ds.RepositoryI, discoveredAt int64) *Pool {
	syncAdapter := ds.NewSyncAdapter(addr, ds.Pool, discoveredAt, client, repo)
	// syncAdapter.V = syncAdapter.FetchVersion(discoveredAt)
	pool := NewPoolFromAdapter(
		syncAdapter,
	)
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(pool.DiscoveredAt),
	}
	underlyingToken, err := pool.contractETH.UnderlyingToken(opts)
	if err != nil {
		log.Fatal(err)
	}
	repo.GetToken(underlyingToken.Hex())
	dieselToken, err := pool.contractETH.DieselToken(opts)
	if err != nil {
		log.Fatal(err)
	}
	pool.SetUnderlyingState(&schemas.PoolState{
		Address:         pool.Address,
		DieselToken:     dieselToken.Hex(),
		UnderlyingToken: underlyingToken.Hex(),
	})
	// create a pool stat snapshot at first log of the pool
	pool.lastEventBlock = pool.DiscoveredAt
	pool.createPoolStat()

	return pool
}

func NewPoolFromAdapter(adapter *ds.SyncAdapter) *Pool {
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
