package cm_v3

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
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
		})
	case core.Topic("FinishMultiCall()"):
		mdl.MulticallMgr.End(txLog.Index)
	case core.Topic("IncreaseDebt(address,uint256)"):
		increaseBorrowEvent, err := mdl.facadeContractv3.ParseIncreaseDebt(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack IncreaseBorrowedAmount event", err)
		}
		mdl.onIncreaseBorrowedAmountV3(&txLog, increaseBorrowEvent.CreditAccount.Hex(),
			increaseBorrowEvent.Amount, "IncreaseDebt")
	case core.Topic("DecreaseDebt(address,uint256)"):
		decreaseBorrowEvent, err := mdl.facadeContractv3.ParseDecreaseDebt(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack DecreaseBorrowedAmount event", err)
		}
		mdl.onIncreaseBorrowedAmountV3(&txLog, decreaseBorrowEvent.CreditAccount.Hex(),
			new(big.Int).Neg(decreaseBorrowEvent.Amount), "DecreaseDebt")
	case core.Topic("SetCreditConfigurator(address)"): // on credit manager
		newConfigurator := utils.ChecksumAddr(txLog.Topics[1].Hex())
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
