package credit_manager

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/models/credit_filter"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strings"
)

func (cm *CreditManager) addCreditConfigurator() {
	creditConfigurator, err := cm.contractETHV2.CreditConfigurator(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	cm.Repo.AddCreditManagerToFilter(cm.Address, creditConfigurator.Hex())
	cf := credit_filter.NewCreditFilter(creditConfigurator.Hex(), core.CreditConfigurator, cm.Address, cm.DiscoveredAt, cm.Client, cm.Repo)
	cm.Repo.AddSyncAdapter(cf)
}

func (mdl *CreditManager) checkLogV2(txLog types.Log) {
	//-- for credit manager stats
	switch txLog.Topics[0] {
	case core.Topic("AddCollateral(address,address,uint256)"):
		addCollateralEvent, err := mdl.facadeContractV2.ParseAddCollateral(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack AddCollateral event", err)
		}

		mdl.onAddCollateralV2(&txLog, addCollateralEvent.OnBehalfOf.Hex(),
			addCollateralEvent.Token.Hex(),
			addCollateralEvent.Value)
	case core.Topic("OpenCreditAccount(address,address,uint256,uint256)"):
		openCreditAccountEvent, err := mdl.facadeContractV2.ParseOpenCreditAccount(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack OpenCreditAccount event", err)
		}
		mdl.onOpenCreditAccountV2(&txLog,
			openCreditAccountEvent.OnBehalfOf.Hex(),
			openCreditAccountEvent.CreditAccount.Hex(),
			openCreditAccountEvent.BorrowAmount,
			openCreditAccountEvent.ReferralCode)
	case core.Topic("CloseCreditAccount(address,address)"):
		closeCreditAccountEvent, err := mdl.facadeContractV2.ParseCloseCreditAccount(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack CloseCreditAccount event", err)
		}

		mdl.onCloseCreditAccountV2(&txLog,
			closeCreditAccountEvent.Owner.Hex(),
			closeCreditAccountEvent.To.Hex())
	case core.Topic("LiquidateCreditAccount(address,address,uint256)"):
		liquidateCreditAccountEvent, err := mdl.contractETHV1.ParseLiquidateCreditAccount(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack LiquidateCreditAccount event", err)
		}

		mdl.onLiquidateCreditAccountV2(&txLog,
			liquidateCreditAccountEvent.Owner.Hex(),
			liquidateCreditAccountEvent.Liquidator.Hex(),
			liquidateCreditAccountEvent.RemainingFunds)
	case core.Topic("MultiCall(borrower)"):
		// liquidateCreditAccountEvent, err := mdl.contractETHV1.ParseLiquidateCreditAccount(txLog)
		// if err != nil {
		// 	log.Fatal("[CreditManagerModel]: Cant unpack LiquidateCreditAccount event", err)
		// }
		sessionId := mdl.GetCreditOwnerSession("")
		mdl.multicall.Start("", txLog.TxHash.Hex(), &core.AccountOperation{
			TxHash:      txLog.TxHash.Hex(),
			BlockNumber: int64(txLog.BlockNumber),
			SessionId:   sessionId,
			Dapp:        txLog.Address.Hex(),
		})
	case core.Topic("MultiCallEnd()"):
		mdl.multicall.End()
	case core.Topic("IncreaseBorrowedAmount(address,uint256)"):
		increaseBorrowEvent, err := mdl.facadeContractV2.ParseIncreaseBorrowedAmount(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack IncreaseBorrowedAmount event", err)
		}
		mdl.onIncreaseBorrowedAmountV2(&txLog,
			increaseBorrowEvent.Borrower.Hex(), increaseBorrowEvent.Amount)
	case core.Topic("DecreaseBorrowedAmount(address,uint256)"):
		decreaseBorrowEvent, err := mdl.facadeContractV2.ParseDecreaseBorrowedAmount(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack DecreaseBorrowedAmount event", err)
		}
		mdl.onIncreaseBorrowedAmountV2(&txLog,
			decreaseBorrowEvent.Borrower.Hex(), new(big.Int).Neg(decreaseBorrowEvent.Amount))
	case core.Topic("TransferAccount(address,address)"):
		if len(txLog.Data) == 0 { // oldowner and newowner are indexed
			transferAccount, err := mdl.facadeContractV2.ParseTransferAccount(txLog)
			if err != nil {
				log.Fatal("[CreditManagerModel]: Cant unpack TransferAccount event", err)
			}
			mdl.onTransferAccountV2(&txLog, transferAccount.OldOwner.Hex(), transferAccount.NewOwner.Hex())
		} else {
			oldOwner := common.BytesToAddress(txLog.Data[:32])
			newOwner := common.BytesToAddress(txLog.Data[32:])
			mdl.onTransferAccountV2(&txLog, oldOwner.Hex(), newOwner.Hex())
		}
	case core.Topic("ExecuteOrder(address,address)"):
		execute, err := mdl.contractETHV2.ParseExecuteOrder(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack ExecuteOrder event", err)
		}
		mdl.AddExecuteParams(&txLog, execute.Borrower, execute.Target)
	case core.Topic("TransferAccountAllowed(address,address,bool)"):
		transferAccount, err := mdl.facadeContractV2.ParseTransferAccountAllowed(txLog)
		log.CheckFatal(err)
		mdl.Repo.TransferAccountAllowed(&core.TransferAccountAllowed{
			From:        transferAccount.From.Hex(),
			To:          transferAccount.To.Hex(),
			Allowed:     transferAccount.State,
			LogId:       int64(txLog.Index),
			BlockNumber: int64(txLog.BlockNumber),
		})
	}
}

func (mdl *CreditManager) onNewTxHashV2(newTxHash string) {
	if mdl.multicall.GetTxHash() == newTxHash {
		return
	}
	mdl.processRemainingMultiCalls()
	mdl.processNonMultiCalls()
}
func (mdl *CreditManager) processRemainingMultiCalls() {
	// opencreditaccount
	mainAction := mdl.multicall.OpenEvent
	// if not present use multicall
	if mainAction == nil {
		mainAction = mdl.multicall.MultiCallStartEvent
	}
	if mdl.multicall.lenOfMultiCalls() > 0 {
		mdl.multiCallHandler(mainAction)
	}
	mdl.multicall.OpenEvent = nil
	mdl.multicall.MultiCallStartEvent = nil
}
func (mdl *CreditManager) processNonMultiCalls() {
	events := mdl.multicall.popNonMulticallEventsV2()
	executeEvents := []core.ExecuteParams{}
	for _, event := range events {
		switch event.Action {
		case "AddCollateral(address,address,uint256)",
			"IncreaseBorrowedAmount(address,uint256)",
			"DecreaseBorrowedAmount(address,uint256)":
			mdl.Repo.AddAccountOperation(event)
		case "ExecuteOrder":
			account := strings.Split(event.SessionId, "_")[0]
			executeEvents = append(executeEvents, core.ExecuteParams{
				SessionId:     event.SessionId,
				CreditAccount: common.HexToAddress(account),
				Protocol:      common.HexToAddress(event.Dapp),
				Borrower:      common.HexToAddress(event.Borrower),
				Index:         event.LogId,
				BlockNumber:   event.BlockNumber,
			})
		}
	}
	if len(executeEvents) > 0 {
		mdl.handleExecuteEvents(executeEvents)
	}
}

func (mdl *CreditManager) GetCreditFacadeAddr() string {
	return mdl.GetDetailsByKey("creditFacade")
}
