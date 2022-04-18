package pool

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
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
		mdl.Repo.AddPoolLedger(&schemas.PoolLedger{
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
		mdl.Repo.AddPoolLedger(&schemas.PoolLedger{
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
		mdl.Repo.AddRepayOnCM(blockNum, repayEvent.CreditManager.Hex(), schemas.PnlOnRepay{
			BorrowedAmount: repayEvent.BorrowedAmount,
			Profit:         repayEvent.Profit,
			Loss:           repayEvent.Loss,
		})

		mdl.lastEventBlock = blockNum
	case core.Topic("NewInterestRateModel(address)"):
		interestModel, err := mdl.contractETH.ParseNewInterestRateModel(txLog)
		log.CheckFatal(err)
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: blockNum,
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    mdl.Address,
			Type:        schemas.NewInterestRateModel,
			Args:        &core.Json{"newInterestRateModel": interestModel.NewInterestRateModel.Hex()},
		})
	case core.Topic("NewCreditManagerConnected(address)"):
		newCreditManager, err := mdl.contractETH.ParseNewCreditManagerConnected(txLog)
		log.CheckFatal(err)
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: blockNum,
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    mdl.Address,
			Type:        schemas.NewCreditManagerConnected,
			Args:        &core.Json{"creditManager": newCreditManager.CreditManager.Hex()},
		})
	case core.Topic("BorrowForbidden(address)"):
		borrowForbidden, err := mdl.contractETH.ParseBorrowForbidden(txLog)
		log.CheckFatal(err)
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: blockNum,
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    mdl.Address,
			Type:        schemas.BorrowForbidden,
			Args:        &core.Json{"creditManager": borrowForbidden.CreditManager.Hex()},
		})
	case core.Topic("NewWithdrawFee(uint256)"):
		oldFee := (*core.BigInt)(mdl.State.WithdrawFee)
		if oldFee == nil {
			oldFee = (*core.BigInt)(new(big.Int))
		}
		withdrawFee, err := mdl.contractETH.ParseNewWithdrawFee(txLog)
		log.CheckFatal(err)
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: blockNum,
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    mdl.Address,
			Type:        schemas.NewWithdrawFee,
			Args: &core.Json{
				"token":  mdl.State.UnderlyingToken,
				"oldFee": oldFee,
				"newFee": (*core.BigInt)(withdrawFee.Fee),
			},
		})
		mdl.State.WithdrawFee = (*core.BigInt)(withdrawFee.Fee)
	case core.Topic("NewExpectedLiquidityLimit(uint256)"):
		expectedLiq, err := mdl.contractETH.ParseNewExpectedLiquidityLimit(txLog)
		log.CheckFatal(err)
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: blockNum,
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    mdl.Address,
			Type:        schemas.NewExpectedLiquidityLimit,
			Args: &core.Json{
				"oldLimit": mdl.State.ExpectedLiquidityLimit,
				"newLimit": (*core.BigInt)(expectedLiq.NewLimit),
				"token":    mdl.State.UnderlyingToken,
			},
		})
		mdl.State.ExpectedLiquidityLimit = (*core.BigInt)(expectedLiq.NewLimit)
	}
}
