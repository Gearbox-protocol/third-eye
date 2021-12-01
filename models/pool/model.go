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
	contractETH *poolService.PoolService
	lastEventBlock int64
}

func NewPool(addr string, client *ethclient.Client, repo core.RepositoryI, discoveredAt int64) *Pool {
	pool := NewPoolFromAdapter(
		core.NewSyncAdapter(addr, "Pool", discoveredAt, client, repo),
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
	mdl.SetLastSync(syncTill)
}