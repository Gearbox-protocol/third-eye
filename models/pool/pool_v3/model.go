package pool_v3

import (
	"github.com/Gearbox-protocol/sdk-go/artifacts/poolv3"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/pool/pool_common"
	"github.com/ethereum/go-ethereum/common"
)

// TODO: basisPointsRate is usually zero in pool USDT. Assumption, this is fee charged per transfer.
type Poolv3 struct {
	*ds.SyncAdapter
	// used for when to take a snapshot of pool state, these can only be taken for 5 events, new interest rate, add/remove liquidity and borrow/repay pool owed amount
	contract       *poolv3.Poolv3
	lastEventBlock int64
	State          *schemas.PoolState
	gatewayHandler pool_common.GatewayHandler
	repayEvents    []*schemas.PoolLedger
	//
	poolKeeper string
}

func (pool *Poolv3) GetRepayEvent() *schemas.PoolLedger {
	ans := pool.repayEvents[0]
	pool.repayEvents = pool.repayEvents[1:]
	return ans
}

func NewPool(addr string, client core.ClientI, repo ds.RepositoryI, discoveredAt int64) *Poolv3 {
	syncAdapter := ds.NewSyncAdapter(addr, ds.Pool, discoveredAt, client, repo)
	// syncAdapter.V = syncAdapter.FetchVersion(discoveredAt)
	pool := NewPoolFromAdapter(
		syncAdapter,
	)
	// underlyingToken
	underlyingTokenData, err := core.CallFuncWithExtraBytes(client, "2495a599", common.HexToAddress(addr), 0, nil)
	log.CheckFatal(err)
	underlyingToken := common.BytesToAddress(underlyingTokenData[:])
	repo.GetToken(underlyingToken.Hex())

	// dieselToken
	dieselToken := addr
	repo.GetToken(dieselToken) //
	//
	pool.SetUnderlyingState(&schemas.PoolState{
		Address:         pool.Address,
		DieselToken:     dieselToken,
		UnderlyingToken: underlyingToken.Hex(),
		Version:         core.NewVersion(300),
		Name: func() string {
			con, err := poolv3.NewPoolv3(common.HexToAddress(addr), client)
			log.CheckFatal(err)
			name, err := con.Name(nil)
			log.CheckFatal(err)
			return name
		}(),
	})

	// create a pool stat snapshot at first log of the pool
	pool.onBlockChangeInternally(pool.DiscoveredAt)

	// poolQuotaKeeper, err := core.CallFuncWithExtraBytes(client, "1ab7c7d7", common.HexToAddress(pool.Address), discoveredAt, nil)
	// if err != nil {
	// 	log.Fatalf("can't get pool quota keeper for %s: %s", pool.Address, err)
	// }
	// pool.setPoolQuotaKeeper(string(poolQuotaKeeper), discoveredAt)
	return pool
}

func (mdl *Poolv3) setPoolQuotaKeeper() {
	if mdl.poolKeeper != "" {
		return
	}
	poolKeeper, err := mdl.contract.PoolQuotaKeeper(nil)
	log.CheckFatal(err)
	mdl.poolKeeper = poolKeeper.Hex()
}

func NewPoolFromAdapter(adapter *ds.SyncAdapter) *Poolv3 {
	poolAddr := common.HexToAddress(adapter.Address)
	gateway := pool_common.GetPoolGateways(adapter.Client)[poolAddr]
	obj := &Poolv3{
		SyncAdapter:    adapter,
		gatewayHandler: pool_common.NewGatewayHandler(gateway),
		contract: func() *poolv3.Poolv3 {
			contract, err := poolv3.NewPoolv3(poolAddr, adapter.Client)
			log.CheckFatal(err)
			return contract
		}(),
	}
	obj.setPoolQuotaKeeper()
	//
	return obj
}

func (mdl Poolv3) Topics() [][]common.Hash {
	return [][]common.Hash{
		{
			// for pool
			core.Topic("SetInterestRateModel(address)"),
			core.Topic("Deposit(address,address,uint256,uint256)"),
			core.Topic("Withdraw(address,address,address,uint256,uint256)"),
			core.Topic("Borrow(address,address,uint256)"),
			core.Topic("Repay(address,uint256,uint256,uint256)"),
			core.Topic("SetInterestRateModel(address)"),
			core.Topic("SetPoolQuotaKeeper(address)"),
			core.Topic("AddCreditManager(address)"),
			core.Topic("SetWithdrawFee(uint256)"),
			core.Topic("UpdateTokenQuotaRate(address,uint256)"),
			// for weth gateway
			core.Topic("WithdrawETH(address,address)"),
			// for wsteth gateway, this event is on stETH token
			core.Topic("Transfer(address,address,uint256)"),
		},
	}
}

func (mdl Poolv3) GetAllAddrsForLogs() (addrs []common.Address) {
	addrs = append(addrs, mdl.SyncAdapter.GetAllAddrsForLogs()...)
	if mdl.gatewayHandler.Gateway == core.NULL_ADDR {
		return
	}
	addrs = append(addrs,
		mdl.gatewayHandler.Gateway,
		mdl.gatewayHandler.Token,
		common.HexToAddress(mdl.poolKeeper),
	)
	return
}
