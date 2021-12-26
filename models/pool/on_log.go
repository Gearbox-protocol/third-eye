package pool

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/utils"
	"github.com/ethereum/go-ethereum/core/types"
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
			LogId:       txLog.Index,
			BlockNumber: blockNum,
			TxHash:      txLog.TxHash.Hex(),
			Pool:        mdl.Address,
			Event:       "AddLiquidity",
			User:        addLiquidityEvent.OnBehalfOf.Hex(),
			AmountBI:    (*core.BigInt)(addLiquidityEvent.Amount),
			Amount:      utils.GetFloat64Decimal(addLiquidityEvent.Amount, mdl.Repo.GetToken(mdl.State.UnderlyingToken).Decimals),
		})
		mdl.lastEventBlock = blockNum
	case core.Topic("RemoveLiquidity(address,address,uint256)"):
		removeLiquidityEvent, err := mdl.contractETH.ParseRemoveLiquidity(txLog)
		if err != nil {
			log.Fatal("[PoolServiceModel]: Cant unpack RemoveLiquidity event", err)
		}
		mdl.Repo.AddPoolLedger(&core.PoolLedger{
			LogId:       txLog.Index,
			BlockNumber: blockNum,
			TxHash:      txLog.TxHash.Hex(),
			Pool:        mdl.Address,
			Event:       "RemoveLiquidity",
			User:        removeLiquidityEvent.Sender.Hex(),
			AmountBI:    (*core.BigInt)(removeLiquidityEvent.Amount),
			Amount:      utils.GetFloat64Decimal(removeLiquidityEvent.Amount, mdl.Repo.GetToken(mdl.State.UnderlyingToken).Decimals),
		})
		mdl.lastEventBlock = blockNum
	case core.Topic("Borrow(address,address,uint256)"):
		mdl.lastEventBlock = blockNum
	case core.Topic("Repay(address,uint256,uint256,uint256)"):
		repayEvent, err := mdl.contractETH.ParseRepay(txLog)
		if err != nil {
			log.Fatal("[PoolServiceModel]: Cant unpack RemoveLiquidity event", err)
		}
		mdl.Repo.AddRepayOnCM(blockNum, repayEvent.CreditManager.Hex(), core.PnlOnRepay{
			BorrowedAmount: repayEvent.BorrowedAmount,
			Profit:         repayEvent.Profit,
			Loss:           repayEvent.Loss,
		})

		mdl.lastEventBlock = blockNum
		// case core.Topic("NewInterestRateModel(address)"):
		// 	mdl.lastEventBlock = blockNum
	}
}
