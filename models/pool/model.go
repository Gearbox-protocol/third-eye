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
	contractETH *poolService.PoolService
	// used for when to take a snapshot of pool state, these can only be taken for 5 events, new interest rate, add/remove liquidity and borrow/repay pool owed amount
	lastEventBlock int64
	State          *schemas.PoolState
	dieselRate     *big.Int
	gatewayHandler GatewayHandler
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
	poolAddr := common.HexToAddress(adapter.Address)
	cmContract, err := poolService.NewPoolService(poolAddr, adapter.Client)
	log.CheckFatal(err)
	gateway := GetPoolGateways(adapter.Client)[poolAddr]
	obj := &Pool{
		SyncAdapter:    adapter,
		contractETH:    cmContract,
		gatewayHandler: NewGatewayHandler(gateway),
	}
	return obj
}

func (mdl *Pool) AfterSyncHook(syncTill int64) {
	mdl.createPoolStat()
	mdl.SyncAdapter.AfterSyncHook(syncTill)
}

func (mdl Pool) Topics() [][]common.Hash {
	return [][]common.Hash{
		{
			// for pool
			core.Topic("AddLiquidity(address,address,uint256,uint256)"),
			core.Topic("RemoveLiquidity(address,address,uint256)"),
			core.Topic("Borrow(address,address,uint256)"),
			core.Topic("Repay(address,uint256,uint256,uint256)"),
			core.Topic("NewInterestRateModel(address)"),
			core.Topic("NewCreditManagerConnected(address)"),
			core.Topic("BorrowForbidden(address)"),
			core.Topic("NewWithdrawFee(uint256)"),
			core.Topic("NewExpectedLiquidityLimit(uint256)"),
			// for weth gateway
			core.Topic("WithdrawETH(address,address)"),
			// for wsteth gateway, this event is on stETH token
			core.Topic("Transfer(address,address,uint256)"),
		},
	}
}

func (mdl Pool) GetAllAddrsForLogs() (addrs []common.Address) {
	addrs = append(addrs, mdl.SyncAdapter.GetAllAddrsForLogs()...)
	if mdl.gatewayHandler.Gateway == core.NULL_ADDR {
		return
	}
	addrs = append(addrs, mdl.gatewayHandler.Gateway, mdl.gatewayHandler.Token)
	return
}
