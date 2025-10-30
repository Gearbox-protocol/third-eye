package cm_v3

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds/multicall_processor"
	"github.com/Gearbox-protocol/third-eye/models/pool/pool_v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// TODO:
// - on updateQuota/increasedebt/decreasedebt, update creditsessionsnapshot with accrued fees/interest
// - use accountQuotaInfo for calculating new accrued interest
// - add use quota for calculating the tvw.

// SetEnabledTokensMask(address,uint256) // same as enabledToken/disabledToken
func (mdl *CMv3) checkLogV3(txLog types.Log) {
	switch txLog.Topics[0] {
	case core.Topic("AddCollateral(address,address,uint256)"):
		addCollateralEvent, err := mdl.facadeContractv3.ParseAddCollateral(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack AddCollateral event", err)
		}

		mdl.onAddCollateralV3(&txLog, addCollateralEvent.CreditAccount.Hex(),
			addCollateralEvent.Token.Hex(),
			addCollateralEvent.Amount)
	case core.Topic("WithdrawCollateral(address,address,uint256,address)"):
		addCollateralEvent, err := mdl.facadeContractv3.ParseWithdrawCollateral(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack AddCollateral event", err)
		}

		mdl.onWithdrawCollateralV3(&txLog, addCollateralEvent.CreditAccount.Hex(),
			addCollateralEvent.Token.Hex(),
			addCollateralEvent.Amount,
			addCollateralEvent.To.Hex())
	case core.Topic("OpenCreditAccount(address,address,address,uint256)"):
		openCreditAccountEvent, err := mdl.facadeContractv3.ParseOpenCreditAccount(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack OpenCreditAccount event", err)
		}
		mdl.Repo.AddAccountAddr(openCreditAccountEvent.CreditAccount.Hex())
		mdl.onOpenCreditAccountV3(&txLog,
			openCreditAccountEvent.OnBehalfOf.Hex(),
			openCreditAccountEvent.CreditAccount.Hex(),
			openCreditAccountEvent.ReferralCode)
	case core.Topic("CloseCreditAccount(address,address)"):
		closeCreditAccountEvent, err := mdl.facadeContractv3.ParseCloseCreditAccount(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack CloseCreditAccount event", err)
		}
		mdl.onCloseCreditAccountV3(&txLog,
			closeCreditAccountEvent.CreditAccount.Hex(), // borrower not used
			closeCreditAccountEvent.Borrower.Hex())
	// for getting correct liquidation status
	case core.Topic("Paused(address)"):
		if txLog.Address.Hex() == mdl.Address { // set pause on cm, if Paused event is emitted only on cm address
			mdl.State.Paused = true
		}
	case core.Topic("Unpaused(address)"):
		if txLog.Address.Hex() == mdl.Address { // unset pause on cm, if Unpaused event is emitted only on cm address
			mdl.State.Paused = false
		}
	case core.Topic("LiquidateCreditAccount(address,address,address,uint256)"):
		liquidateCreditAccountEvent, err := mdl.facadeContractv3.ParseLiquidateCreditAccount(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack LiquidateCreditAccount event", err)
		}
		mdl.onLiquidateCreditAccountV3(&txLog,
			liquidateCreditAccountEvent.CreditAccount.Hex(), // borrower not used
			liquidateCreditAccountEvent.Liquidator.Hex(),
			liquidateCreditAccountEvent.To.Hex(),
			liquidateCreditAccountEvent.RemainingFunds,
		)
	case core.Topic("StartMultiCall(address,address)"):
		creditAccount := common.BytesToAddress(txLog.Topics[1][:]).Hex()
		sessionId, borrower := mdl.GetSessionIdAndBorrower(creditAccount)
		mdl.SetSessionIsUpdated(sessionId) // it is needed for updateQuota, as we don't update session on Updatequota and updateQuota can't be emitted without StartMulticall
		mdl.MulticallMgr.Start(txLog.TxHash.Hex(), &schemas.AccountOperation{
			TxHash:      txLog.TxHash.Hex(),
			BlockNumber: int64(txLog.BlockNumber),
			SessionId:   sessionId,
			Borrower:    borrower,
			Dapp:        txLog.Address.Hex(),
			LogId:       txLog.Index,
			Args:        &core.Json{},
			Action:      "StartMultiCall(address,address)",
		}, mdl.Address)
	case core.Topic("FinishMultiCall()"): // debts are now always in multicalls
		poolv3 := mdl.Repo.GetAdapter(mdl.State.PoolAddress).(*pool_v3.Poolv3)
		debts := poolv3.GetDebt(txLog.TxHash, mdl.Address, txLog.Index)
		// log.Info(debts)
		for _, debt := range debts {
			sessionId, borrower := mdl.GetSessionIdAndBorrower(debt.Account)
			mdl.PoolBorrow(&txLog, sessionId, borrower, debt.Amount)
			// mdl.onIncreaseBorrowedAmountV3(&txLog, debt.Account,
			// 	debt.Amount, "IncreaseDebt")
		}
		mdl.MulticallMgr.End(txLog.Index, debts, mdl.GetUnderlyingToken(), mdl.Address)
	case core.Topic("IncreaseDebt(address,uint256)"): // works for v300 for V3 10: The increased or decreased debt event is not emitted. We can check that using the debt that is emitted from the pool. In the finish_multicall we are emitting that call. So this function is just above you.
		increaseBorrowEvent, err := mdl.facadeContractv3.ParseIncreaseDebt(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack IncreaseBorrowedAmount event", err)
		}
		mdl.onIncreaseBorrowedAmountV3(&txLog, increaseBorrowEvent.CreditAccount.Hex(),
			increaseBorrowEvent.Amount, "IncreaseDebt")
	case core.Topic("PartiallyLiquidateCreditAccount(address,address,address,uint256,uint256,uint256)"):
		creditAccount := common.BytesToAddress(txLog.Topics[1][:]) // CA
		token := common.BytesToAddress(txLog.Topics[2][:])         // CA
		repaidDebt := new(big.Int).SetBytes(txLog.Data[:32])
		seizedCol := new(big.Int).SetBytes(txLog.Data[32:64])
		fee := new(big.Int).SetBytes(txLog.Data[64:96])
		sessionId, borrower := mdl.GetSessionIdAndBorrower(creditAccount.Hex())
		accountOp := &schemas.AccountOperation{
			TxHash:      txLog.TxHash.Hex(),
			BlockNumber: int64(txLog.BlockNumber),
			LogId:       txLog.Index,
			Borrower:    borrower,
			SessionId:   sessionId,
			Dapp:        mdl.GetCreditFacadeAddr(),
			AdapterCall: false,
			Action:      "PartialLiquidation",
			Args: &core.Json{
				"creditAccount": creditAccount,
				"token":         token,
				"repaidDebt":    repaidDebt,
				"seizedCol":     seizedCol,
				"fee":           fee,
			},
			Transfers: &core.Transfers{},
		}
		mdl.MulticallMgr.SetPartialLiq(creditAccount)
		poolv3 := mdl.Repo.GetAdapter(mdl.State.PoolAddress).(*pool_v3.Poolv3)
		debts := poolv3.GetDebt(txLog.TxHash, mdl.Address, txLog.Index)
		multicall_processor.AddManageDebtsToMain(accountOp, debts, mdl.GetUnderlyingToken())
		mdl.Repo.AddAccountOperation(accountOp)
		//
		mdl.poolRepayv3(txLog.TxHash.Hex(), sessionId, borrower)
	case core.Topic("DecreaseDebt(address,uint256)"):
		decreaseBorrowEvent, err := mdl.facadeContractv3.ParseDecreaseDebt(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack DecreaseBorrowedAmount event", err)
		}
		mdl.onIncreaseBorrowedAmountV3(&txLog, decreaseBorrowEvent.CreditAccount.Hex(),
			new(big.Int).Neg(decreaseBorrowEvent.Amount), "DecreaseDebt")
	case core.Topic("SetCreditConfigurator(address)"): // on credit manager
		newConfigurator := utils.ChecksumAddr(txLog.Topics[1].Hex())
		configuratorAtBlock, err := core.CallFuncGetSingleValue(mdl.Client, "2f7a1881", common.HexToAddress(mdl.Address), int64(txLog.BlockNumber), nil)
		log.CheckFatal(err)
		if common.BytesToAddress(configuratorAtBlock).Hex() != newConfigurator { // https://etherscan.io/tx/0x6d46dc6dcc2045b4f282134f106d8ad6f59e904b1aa3a6ad78d6ded72e02f7d9You due to two cm at the same block number.
			return
		}
		oldConfigurator := mdl.GetDetailsByKey("configurator")
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: int64(txLog.BlockNumber),
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    mdl.Address,
			Args:        &core.Json{"oldConfigurator": oldConfigurator, "configurator": newConfigurator},
			Type:        schemas.NewConfigurator,
		})
		if oldConfigurator != newConfigurator {
			mdl.Repo.GetAdapter(oldConfigurator).SetBlockToDisableOn(int64(txLog.BlockNumber))
			mdl.addCreditConfiguratorAdapter(newConfigurator)
			mdl.setConfiguratorSyncer(newConfigurator)
		}
		// on credit manager
	case core.Topic("Execute(address,address)"):
		execute, err := mdl.facadeContractv3.ParseExecute(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack ExecuteOrder event", err)
		}
		mdl.AddExecuteParamsV3(&txLog, execute.CreditAccount, execute.TargetContract)
	}
}
