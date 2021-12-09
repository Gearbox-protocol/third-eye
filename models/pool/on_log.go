package pool

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

func (mdl *Pool) createPoolStat() {
	// datacompressor works for pool address only after the address is registered with contractregister
	// i.e. discoveredAt
	if mdl.lastEventBlock != 0 && mdl.lastEventBlock >= mdl.DiscoveredAt {
		mdl.calculatePoolStat(mdl.lastEventBlock)
		mdl.lastEventBlock = 0
	}
}

func (mdl *Pool) OnLog(txLog types.Log) {
	blockNum := int64(txLog.BlockNumber)
	if mdl.lastEventBlock != blockNum {
		mdl.createPoolStat()
	}
	switch txLog.Topics[0] {
	case core.Topic("AddLiquidity(address,address,uint256,uint256)"):
		addLiquidityEvent, err := mdl.contractETH.ParseAddLiquidity(txLog)
		if err != nil {
			log.Fatal("[PoolServiceModel]: Cant unpack AddLiquidity event", err)
		}
		mdl.Repo.AddPoolLedger(&core.PoolLedger{
			LogId:       int64(txLog.Index),
			BlockNumber: blockNum,
			Pool:        mdl.Address,
			Event:       "AddLiquidity",
			User:        addLiquidityEvent.OnBehalfOf.Hex(),
			Liquidity:   (*core.BigInt)(addLiquidityEvent.Amount),
		})
		mdl.lastEventBlock = blockNum
	case core.Topic("RemoveLiquidity(address,address,uint256)"):
		removeLiquidityEvent, err := mdl.contractETH.ParseRemoveLiquidity(txLog)
		if err != nil {
			log.Fatal("[PoolServiceModel]: Cant unpack RemoveLiquidity event", err)
		}
		mdl.Repo.AddPoolLedger(&core.PoolLedger{
			LogId:       int64(txLog.Index),
			BlockNumber: blockNum,
			Pool:        mdl.Address,
			Event:       "RemoveLiquidity",
			User:        removeLiquidityEvent.Sender.Hex(),
			Liquidity:   (*core.BigInt)(new(big.Int).Neg(removeLiquidityEvent.Amount)),
		})
		mdl.lastEventBlock = blockNum
	case core.Topic("Borrow(address,address,uint256)"):
		mdl.lastEventBlock = blockNum
	case core.Topic("Repay(address,uint256,uint256,uint256)"):
		mdl.lastEventBlock = blockNum
		// case core.Topic("NewInterestRateModel(address)"):
		// 	mdl.lastEventBlock = blockNum
	}
}
