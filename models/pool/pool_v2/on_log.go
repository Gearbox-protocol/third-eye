package pool_v2

import (
	"fmt"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/models/pool/pool_common"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// addLiquidity , amount is underlying token
// remove Liquidity , amount is in diesel. though after the gatewayHandler all the removeLiquidity amount are changed to underlyingToken in fixPoolLedgerAddrForGateway
// lend amount is in underlying
// repay amount/profit/loss is in underlying

// called on lend,repay,add,remove and NewInterestRateModel
func (mdl *Poolv2) updateBorrowRate(blockNum int64) {
	mdl.lastEventBlock = blockNum
}
func (mdl *Poolv2) OnLog(txLog types.Log) {
	blockNum := int64(txLog.BlockNumber)
	// REVERT_POOL_WRAPPER
	// if mdl.lastEventBlock != 0 && blockNum != mdl.lastEventBlock && mdl.lastEventBlock >= mdl.DiscoveredAt {
	// 	mdl.onBlockChangeInternally(mdl.lastEventBlock)
	// 	mdl.lastEventBlock = 0
	// }
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
			//
			Executor: addLiquidityEvent.Sender.Hex(),
			User:     addLiquidityEvent.Sender.Hex(),
			Receiver: addLiquidityEvent.OnBehalfOf.Hex(),
			//
			AmountBI: (*core.BigInt)(addLiquidityEvent.Amount),
			Amount:   utils.GetFloat64Decimal(addLiquidityEvent.Amount, mdl.Repo.GetToken(mdl.State.UnderlyingToken).Decimals),
		})
		pool_common.CheckIfAmountMoreThan1Mil(mdl.Client, mdl.Repo, mdl.State, addLiquidityEvent.Amount, blockNum, txLog.TxHash.Hex(),
			fmt.Sprintf("%s deposit", mdl.Repo.GetToken(mdl.State.UnderlyingToken).Symbol))
		mdl.updateBorrowRate(blockNum)
	case core.Topic("RemoveLiquidity(address,address,uint256)"):
		removeLiquidityEvent, err := mdl.contractETH.ParseRemoveLiquidity(txLog)
		if err != nil {
			log.Fatal("[PoolServiceModel]: Cant unpack RemoveLiquidity event", err)
		}
		// log.Info(removeLiquidityEvent.Sender.Hex())
		mdl.gatewayHandler.AddRemoveLiqEvent(&schemas.PoolLedger{
			LogId:       txLog.Index,
			BlockNumber: blockNum,
			TxHash:      txLog.TxHash.Hex(),
			Pool:        mdl.Address,
			Event:       "RemoveLiquidity",
			Executor:    removeLiquidityEvent.Sender.Hex(),
			User:        removeLiquidityEvent.Sender.Hex(),
			Receiver:    removeLiquidityEvent.To.Hex(),
			AmountBI:    (*core.BigInt)(removeLiquidityEvent.Amount),
		})
		pool_common.CheckIfAmountMoreThan1Mil(mdl.Client, mdl.Repo, mdl.State, removeLiquidityEvent.Amount, blockNum, txLog.TxHash.Hex(),
			fmt.Sprintf("%s withdrawal", mdl.Repo.GetToken(mdl.State.UnderlyingToken).Symbol))

		mdl.updateBorrowRate(blockNum)
	case core.Topic("Borrow(address,address,uint256)"):
		mdl.updateBorrowRate(blockNum)
	case core.Topic("Repay(address,uint256,uint256,uint256)"):
		repayEvent, err := mdl.contractETH.ParseRepay(txLog)
		if err != nil {
			log.Fatal("[PoolServiceModel]: Cant unpack RemoveLiquidity event", err)
		}
		mdl.Repo.AddRepayOnCM(repayEvent.CreditManager.Hex(), schemas.PnlOnRepay{
			BlockNum:       blockNum,
			BorrowedAmount: repayEvent.BorrowedAmount,
			Profit:         repayEvent.Profit,
			Loss:           repayEvent.Loss,
		})
		mdl.updateBorrowRate(blockNum)
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
		mdl.updateBorrowRate(blockNum)
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
		// for withdrawETh
	case core.Topic("WithdrawETH(address,address)"):
		pool := common.BytesToAddress(txLog.Topics[1][:]).Hex()
		user := common.BytesToAddress(txLog.Topics[2][:]).Hex()
		ind := txLog.Index - 2
		blockNum := int64(txLog.BlockNumber)
		// for weth pool, WithdrawETH is emitted on weth gateway, so we track withdraETH on gateway for getting user
		mdl.gatewayHandler.CheckWithdrawETH(txLog.TxHash.Hex(), blockNum, int64(ind), pool, user)
	case core.Topic("Transfer(address,address,uint256)"):
		from := common.BytesToAddress(txLog.Topics[1][:])
		to := common.BytesToAddress(txLog.Topics[2][:]).Hex()
		// for wsteth pool, there is wsteth wrapper,
		// to get the user address we have to track Transfer on steth where from is gateway and user is not pool
		if !(from == mdl.gatewayHandler.Gateway && to != mdl.gatewayHandler.UserCantBe.Hex()) {
			return
		}
		ind := txLog.Index - 3
		blockNum := int64(txLog.BlockNumber)
		mdl.gatewayHandler.CheckWithdrawETH(txLog.TxHash.Hex(), blockNum, int64(ind), mdl.Address, to)
	}
}
