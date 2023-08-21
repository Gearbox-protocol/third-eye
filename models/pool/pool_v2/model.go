package pool_v2

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/poolService"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/pool/pool_common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type Poolv2 struct {
	*ds.SyncAdapter
	contractETH *poolService.PoolService
	// used for when to take a snapshot of pool state, these can only be taken for 5 events, new interest rate, add/remove liquidity and borrow/repay pool owed amount
	lastEventBlock int64
	State          *schemas.PoolState
	dieselRate     *big.Int
	gatewayHandler pool_common.GatewayHandler
}

func NewPool(addr string, client core.ClientI, repo ds.RepositoryI, discoveredAt int64) *Poolv2 {
	syncAdapter := ds.NewSyncAdapter(addr, ds.Pool, discoveredAt, client, repo)
	// syncAdapter.V = syncAdapter.FetchVersion(discoveredAt)
	pool := NewPoolFromAdapter(
		syncAdapter,
	)
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(pool.DiscoveredAt),
	}
	//
	underlyingToken, err := pool.contractETH.UnderlyingToken(opts)
	log.CheckFatal(err)
	repo.GetToken(underlyingToken.Hex())
	//
	dieselToken, err := pool.contractETH.DieselToken(opts)
	log.CheckFatal(err)
	//
	pool.SetUnderlyingState(&schemas.PoolState{
		Address:         pool.Address,
		DieselToken:     dieselToken.Hex(),
		UnderlyingToken: underlyingToken.Hex(),
	})
	// create a pool stat snapshot at first log of the pool
	pool.onBlockChangeInternally(pool.DiscoveredAt)

	return pool
}

// REVERT_POOL_WRAPPER
// func (p *Pool) AfterSyncHook(syncedTill int64) {
// 	if p.lastEventBlock != 0 {
// 		p.onBlockChangeInternally(p.lastEventBlock)
// 		p.lastEventBlock = 0
// 	}
// 	p.SyncAdapter.AfterSyncHook(syncedTill)
// }

func NewPoolFromAdapter(adapter *ds.SyncAdapter) *Poolv2 {
	poolAddr := common.HexToAddress(adapter.Address)
	cmContract, err := poolService.NewPoolService(poolAddr, adapter.Client)
	log.CheckFatal(err)
	gateway := pool_common.GetPoolGateways(adapter.Client)[poolAddr]
	obj := &Poolv2{
		SyncAdapter:    adapter,
		contractETH:    cmContract,
		gatewayHandler: pool_common.NewGatewayHandler(gateway),
	}
	return obj
}

func (mdl Poolv2) Topics() [][]common.Hash {
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

func (mdl Poolv2) GetAllAddrsForLogs() (addrs []common.Address) {
	addrs = append(addrs, mdl.SyncAdapter.GetAllAddrsForLogs()...)
	if mdl.gatewayHandler.Gateway == core.NULL_ADDR {
		return
	}
	addrs = append(addrs, mdl.gatewayHandler.Gateway, mdl.gatewayHandler.Token)
	return
}
