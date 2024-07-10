package pool_v3

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/models/pool/pool_common"
	"github.com/Gearbox-protocol/third-eye/models/pool_quota_keeper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// called on lend,repay,add,remove and NewInterestRateModel
func (mdl *Poolv3) updateBorrowRate(blockNum int64) {
	mdl.lastEventBlock = blockNum
}

func (mdl Poolv3) getDecimals() int8 {
	return mdl.Repo.GetToken(mdl.State.UnderlyingToken).Decimals
}

func (mdl *Poolv3) OnLog(txLog types.Log) {
	blockNum := int64(txLog.BlockNumber)
	switch txLog.Topics[0] {
	case core.Topic("Transfer(address,address,uint256)"):
		if utils.Contains(mdl.zappers.GetFarm(), txLog.Address.Hex()) { //if farm
			mdl.updateFarmedv3(txLog)
		}
	case core.Topic("UpdateTokenQuotaRate(address,uint256)"):
		mdl.updateBorrowRate(blockNum)
	case core.Topic("SetInterestRateModel(address)"):
		interestRateModel := common.BytesToAddress(txLog.Topics[1][:])
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: blockNum,
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    mdl.Address,
			Type:        schemas.NewInterestRateModel,
			Args:        &core.Json{"newInterestRateModel": interestRateModel.Hex()},
		})
		mdl.updateBorrowRate(blockNum)
	case core.Topic("SetCreditManagerDebtLimit(address,uint256)"):
		cm := common.BytesToAddress(txLog.Topics[1][:])
		total := new(big.Int).SetBytes(txLog.Data[:32])
		mdl.Repo.GetAdapter(cm.Hex()).SetUnderlyingState(total)
		// while processing deposit event, sub from user and add to receiver
	case core.Topic("Deposit(address,address,uint256,uint256)"):
		deposit, err := mdl.contract.ParseDeposit(txLog)
		log.CheckFatal(err)
		event := &schemas.PoolLedger{
			LogId:       txLog.Index,
			BlockNumber: blockNum,
			TxHash:      txLog.TxHash.Hex(),
			Pool:        mdl.Address,
			Event:       "AddLiquidity",
			Executor:    deposit.Sender.Hex(),
			User:        deposit.Sender.Hex(),
			Receiver:    deposit.Owner.Hex(),
			SharesBI:    (*core.BigInt)(deposit.Shares),
			Shares:      utils.GetFloat64Decimal(deposit.Shares, mdl.getDecimals()),
			AmountBI:    (*core.BigInt)(deposit.Assets),
			Amount:      utils.GetFloat64Decimal(deposit.Assets, mdl.getDecimals()),
		}
		if mdl.checkIfZapAddr(deposit.Sender.Hex()) {
			mdl.changeAddressOnAddLiq(event)
		} else {
			mdl.Repo.AddPoolLedger(event)
		}
		pool_common.CheckIfAmountMoreThan1Mil(mdl.Client, mdl.Repo, mdl.State, deposit.Assets, blockNum, txLog.TxHash.Hex(), "deposit")
		mdl.updateBorrowRate(blockNum)
		// while processing withdrawal event, add to receiver and sub from User
	case core.Topic("Withdraw(address,address,address,uint256,uint256)"):
		withdrawal, err := mdl.contract.ParseWithdraw(txLog)
		log.CheckFatal(err)
		event := &schemas.PoolLedger{
			LogId:       txLog.Index,
			BlockNumber: blockNum,
			TxHash:      txLog.TxHash.Hex(),
			Pool:        mdl.Address,
			Event:       "RemoveLiquidity",
			//
			Executor: withdrawal.Sender.Hex(),
			User:     withdrawal.Owner.Hex(),
			Receiver: withdrawal.Receiver.Hex(),
			//
			SharesBI: (*core.BigInt)(withdrawal.Shares),
			Shares:   utils.GetFloat64Decimal(withdrawal.Shares, mdl.getDecimals()),
			AmountBI: (*core.BigInt)(withdrawal.Assets),
			Amount:   utils.GetFloat64Decimal(withdrawal.Assets, mdl.getDecimals()),
		}
		if mdl.checkIfZapAddr(withdrawal.Sender.Hex()) {
			mdl.changeAddressOnRemoveLiq(event)
		} else {
			mdl.Repo.AddPoolLedger(event)
		}
		//
		pool_common.CheckIfAmountMoreThan1Mil(mdl.Client, mdl.Repo, mdl.State, withdrawal.Assets, blockNum, txLog.TxHash.Hex(), "withdraw")
		mdl.updateBorrowRate(blockNum)
	case core.Topic("Borrow(address,address,uint256)"):
		mdl.updateBorrowRate(blockNum)
	case core.Topic("Repay(address,uint256,uint256,uint256)"):
		repayEvent, err := mdl.contract.ParseRepay(txLog)
		if err != nil {
			log.Fatal("[PoolServiceModel]: Cant unpack RemoveLiquidity event", err)
		}
		mdl.Repo.AddRepayOnCM(repayEvent.CreditManager.Hex(), schemas.PnlOnRepay{
			BlockNum:       blockNum,
			BorrowedAmount: repayEvent.BorrowedAmount,
			Profit:         repayEvent.Profit,
			Loss:           repayEvent.Loss,
		})
		amount := new(big.Int).Sub(new(big.Int).Add(repayEvent.BorrowedAmount, repayEvent.Profit), repayEvent.Loss)
		mdl.repayEvents = append(mdl.repayEvents, &schemas.PoolLedger{
			BlockNumber: blockNum,
			Pool:        mdl.Address,
			//
			LogId:    txLog.Index,
			Executor: repayEvent.CreditManager.Hex(),
			TxHash:   txLog.TxHash.Hex(),
			Event:    "Repay",
			//
			AmountBI: (*core.BigInt)(amount),
			Amount:   utils.GetFloat64Decimal(amount, mdl.getDecimals()),
		})
		mdl.updateBorrowRate(blockNum)
	case core.Topic("SetPoolQuotaKeeper(address)"):
		poolQuotaKeeper := common.BytesToAddress(txLog.Topics[1][:]).Hex()
		mdl.setPoolKeeperAdapter(poolQuotaKeeper, blockNum)
	case core.Topic("AddCreditManager(address)"):
		newCreditManager := common.BytesToAddress(txLog.Topics[1][:])
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: blockNum,
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    mdl.Address,
			Type:        schemas.NewCreditManagerConnected,
			Args:        &core.Json{"creditManager": newCreditManager.Hex()},
		})
	case core.Topic("SetWithdrawFee(uint256)"):
		oldFee := (*core.BigInt)(mdl.State.WithdrawFee)
		if oldFee == nil {
			oldFee = (*core.BigInt)(new(big.Int))
		}
		withdrawFee := new(big.Int).SetBytes(txLog.Data)
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: blockNum,
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    mdl.Address,
			Type:        schemas.NewWithdrawFee,
			Args: &core.Json{
				"token":  mdl.State.UnderlyingToken,
				"oldFee": oldFee,
				"newFee": (*core.BigInt)(withdrawFee),
			},
		})
		mdl.State.WithdrawFee = (*core.BigInt)(withdrawFee)
	case core.Topic("AddCreditManager(address)"):
		newCreditManager := common.BytesToAddress(txLog.Topics[1][:])
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: blockNum,
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    mdl.Address,
			Type:        schemas.NewCreditManagerConnected,
			Args:        &core.Json{"creditManager": newCreditManager.Hex()},
		})
	}
}

func (mdl Poolv3) setPoolKeeperAdapter(poolQuotaKeeper string, blockNum int64) {
	pqk := pool_quota_keeper.NewPoolQuotaKeeper(poolQuotaKeeper, mdl.Address, blockNum, mdl.Client, mdl.Repo)
	mdl.setDetailsByKey("PoolKeeper", poolQuotaKeeper)
	mdl.Repo.AddSyncAdapter(pqk)
}
