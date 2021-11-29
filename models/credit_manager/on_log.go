package credit_manager

import (
	"github.com/Gearbox-protocol/gearscan/core"
	"github.com/Gearbox-protocol/gearscan/log"
	"github.com/Gearbox-protocol/gearscan/services"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (mdl *CreditManager) processExecuteEvents() {
	if len(mdl.executeParams) > 0 {
		mdl.handleExecuteEvents()
		mdl.executeParams = []services.ExecuteParams{}
	}
}
func (mdl *CreditManager) OnLog(txLog types.Log) {
	// storing execute order in a single tx and processing them  single go on next tx
	if mdl.LastTxHash != txLog.TxHash.Hex() {
		mdl.processExecuteEvents()
		mdl.LastTxHash = txLog.TxHash.Hex()
	}
	switch txLog.Topics[0] {
	case core.Topic("OpenCreditAccount(address,address,address,uint256,uint256,uint256)"):
		openCreditAccountEvent, err := mdl.contractETH.ParseOpenCreditAccount(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack OpenCreditAccount event", err)
		}

		mdl.onOpenCreditAccount(&txLog, openCreditAccountEvent.Sender.Hex(),
			openCreditAccountEvent.OnBehalfOf.Hex(),
			openCreditAccountEvent.CreditAccount.Hex(),
			openCreditAccountEvent.Amount,
			openCreditAccountEvent.BorrowAmount,
			openCreditAccountEvent.ReferralCode)
	case core.Topic("CloseCreditAccount(address,address,uint256)"):
		closeCreditAccountEvent, err := mdl.contractETH.ParseCloseCreditAccount(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack CloseCreditAccount event", err)
		}

		mdl.onCloseCreditAccount(&txLog,
			closeCreditAccountEvent.Owner.Hex(),
			closeCreditAccountEvent.To.Hex(),
			closeCreditAccountEvent.RemainingFunds)
	case core.Topic("LiquidateCreditAccount(address,address,uint256)"):
		liquidateCreditAccountEvent, err := mdl.contractETH.ParseLiquidateCreditAccount(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack LiquidateCreditAccount event", err)
		}

		mdl.onLiquidateCreditAccount(&txLog,
			liquidateCreditAccountEvent.Owner.Hex(),
			liquidateCreditAccountEvent.Liquidator.Hex(),
			liquidateCreditAccountEvent.RemainingFunds)
	case core.Topic("RepayCreditAccount(address,address)"):
		repayCreditAccountEvent, err := mdl.contractETH.ParseRepayCreditAccount(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack RepayCreditAccount event", err)
		}

		mdl.onRepayCreditAccount(&txLog,
			repayCreditAccountEvent.Owner.Hex(),
			repayCreditAccountEvent.To.Hex())
	case core.Topic("IncreaseBorrowedAmount(address,uint256)"):
		increaseBorrowEvent, err := mdl.contractETH.ParseIncreaseBorrowedAmount(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack IncreaseBorrowedAmount event", err)
		}

		mdl.onIncreaseBorrowedAmount(&txLog,
			increaseBorrowEvent.Borrower.Hex(), increaseBorrowEvent.Amount)
	case core.Topic("AddCollateral(address,address,uint256)"):
		addCollateralEvent, err := mdl.contractETH.ParseAddCollateral(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack AddCollateral event", err)
		}

		mdl.onAddCollateral(&txLog, addCollateralEvent.OnBehalfOf.Hex(),
			addCollateralEvent.Token.Hex(),
			addCollateralEvent.Value)
	case core.Topic("ExecuteOrder(address,address)"):
		execute, err := mdl.contractETH.ParseExecuteOrder(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack ExecuteOrder event", err)
		}
		mdl.AddParams(&txLog, execute.Borrower, execute.Target)
	case core.Topic("TransferAccount(address,address)"):
		if len(txLog.Data) == 0 { // oldowner and newowner are indexed
			transferAccount, err := mdl.contractETH.ParseTransferAccount(txLog)
			if err != nil {
				log.Fatal("[CreditManagerModel]: Cant unpack TransferAccount event", err)
			}
			mdl.onTransferAccount(&txLog, transferAccount.OldOwner.Hex(), transferAccount.NewOwner.Hex())
		} else {
			oldOwner := common.BytesToAddress(txLog.Data[:32])
			newOwner := common.BytesToAddress(txLog.Data[32:])
			mdl.onTransferAccount(&txLog, oldOwner.Hex(), newOwner.Hex())
		}
	}
}
