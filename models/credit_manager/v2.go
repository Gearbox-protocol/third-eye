package credit_manager

import (
	"math/big"
	"strings"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

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
	case core.Topic("OpenCreditAccount(address,address,uint256,uint16)"):
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
			closeCreditAccountEvent.Borrower.Hex(),
			closeCreditAccountEvent.To.Hex())
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
		liquidateCreditAccountEvent, err := mdl.facadeContractV2.ParseLiquidateCreditAccount(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack LiquidateCreditAccount event", err)
		}

		mdl.onLiquidateCreditAccountV2(&txLog,
			liquidateCreditAccountEvent.Borrower.Hex(),
			liquidateCreditAccountEvent.Liquidator.Hex(),
			liquidateCreditAccountEvent.RemainingFunds)
	case core.Topic("TokenEnabled(address,address)"):
		mdl.enableOrDisableToken(txLog, "TokenEnabled(address,address)")
	case core.Topic("TokenDisabled(address,address)"):
		mdl.enableOrDisableToken(txLog, "TokenDisabled(address,address)")
	case core.Topic("MultiCallStarted(address)"):
		borrower := common.HexToAddress(txLog.Topics[1].Hex()).Hex()
		sessionId := mdl.GetCreditOwnerSession(borrower)
		mdl.multicall.Start(txLog.TxHash.Hex(), &schemas.AccountOperation{
			TxHash:      txLog.TxHash.Hex(),
			BlockNumber: int64(txLog.BlockNumber),
			SessionId:   sessionId,
			Borrower:    borrower,
			Dapp:        txLog.Address.Hex(),
			Action:      "MultiCallStarted(address)",
		})
	case core.Topic("MultiCallFinished()"):
		mdl.multicall.End()
	case core.Topic("IncreaseBorrowedAmount(address,uint256)"):
		increaseBorrowEvent, err := mdl.facadeContractV2.ParseIncreaseBorrowedAmount(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack IncreaseBorrowedAmount event", err)
		}
		mdl.onIncreaseBorrowedAmountV2(&txLog, increaseBorrowEvent.Borrower.Hex(),
			increaseBorrowEvent.Amount, "IncreaseBorrowedAmount")
	case core.Topic("DecreaseBorrowedAmount(address,uint256)"):
		decreaseBorrowEvent, err := mdl.facadeContractV2.ParseDecreaseBorrowedAmount(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack DecreaseBorrowedAmount event", err)
		}
		mdl.onIncreaseBorrowedAmountV2(&txLog, decreaseBorrowEvent.Borrower.Hex(),
			new(big.Int).Neg(decreaseBorrowEvent.Amount), "DecreaseBorrowedAmount")
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
	case core.Topic("TransferAccountAllowed(address,address,bool)"):
		transferAccount, err := mdl.facadeContractV2.ParseTransferAccountAllowed(txLog)
		log.CheckFatal(err)
		mdl.Repo.TransferAccountAllowed(&schemas.TransferAccountAllowed{
			From:        transferAccount.From.Hex(),
			To:          transferAccount.To.Hex(),
			Allowed:     transferAccount.State,
			LogId:       int64(txLog.Index),
			BlockNumber: int64(txLog.BlockNumber),
		})
	// on credit manager
	case core.Topic("ExecuteOrder(address,address)"):
		execute, err := mdl.contractETHV2.ParseExecuteOrder(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack ExecuteOrder event", err)
		}
		mdl.AddExecuteParamsV2(&txLog, execute.Borrower, execute.Target)
	case core.Topic("NewConfigurator(address)"):
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
	}
}

func (mdl *CreditManager) onNewTxHashV2() {
	nonMulticallExecuteEvents := mdl.processNonMultiCalls()
	mdl.processRemainingMultiCalls(nonMulticallExecuteEvents)
}

// opencreditaccount
// addcollateral
// increase/decase borrow amount
// executeorder
// are added to multicall manager
//
// #######
// FLOWS ->
// openwithoutmulticall => add collateral
// openwithmulticall => other calls
// multicallstarted => other calls
// other calls => closed/liquidated
func (mdl *CreditManager) processRemainingMultiCalls(nonMultiCallExecuteEvents []ds.ExecuteParams) {

	facadeActions, openEventWithoutMulticall := mdl.multicall.PopMainActionsv2()

	for _, entry := range openEventWithoutMulticall {
		// opencreditaccount without mulitcall
		openWithoutMC := entry.Data
		mdl.setUpdateSession(openWithoutMC.SessionId)
		mdl.Repo.AddAccountOperation(openWithoutMC)
		mdl.addCollteralForOpenCreditAccount(openWithoutMC.BlockNumber, openWithoutMC)
	}
	if len(facadeActions) > 0 { // account operation will only exist if there are one or more facade actions
		mainCalls := mdl.Repo.GetExecuteParser().GetMainCalls(mdl.LastTxHash, mdl.GetCreditFacadeAddr())
		fixedFacadeActions := mdl.fixFacadeActionStructureViaTenderlyCalls(mainCalls, facadeActions)
		mdl.validateAndSaveFacadeActions(mdl.LastTxHash, fixedFacadeActions, mainCalls, nonMultiCallExecuteEvents)
	} else if len(nonMultiCallExecuteEvents) > 0 {
		mdl.saveExecuteEvents(nonMultiCallExecuteEvents)
	}
}

func (mdl *CreditManager) setUpdateSession(sessionId string) {
	// log.Info(log.DetectFunc(),sessionId, "increased")
	mdl.UpdatedSessions[sessionId]++
}

func (mdl *CreditManager) processNonMultiCalls() (executeEvents []ds.ExecuteParams) {
	events := mdl.multicall.PopNonMulticallEventsV2()

	for _, event := range events {
		switch event.Action {
		case "AddCollateral(address,address,uint256)",
			"IncreaseBorrowedAmount(address,uint256)",
			"TokenEnabled(address,address)",
			"TokenDisabled(address,address)",
			"DecreaseBorrowedAmount(address,uint256)":
			mdl.setUpdateSession(event.SessionId)
			mdl.Repo.AddAccountOperation(event)
		case "ExecuteOrder":
			account := strings.Split(event.SessionId, "_")[0]
			mdl.setUpdateSession(event.SessionId)
			executeEvents = append(executeEvents, ds.ExecuteParams{
				SessionId:     event.SessionId,
				CreditAccount: common.HexToAddress(account),
				Protocol:      common.HexToAddress(event.Dapp),
				Borrower:      common.HexToAddress(event.Borrower),
				Index:         event.LogId,
				BlockNumber:   event.BlockNumber,
			})
		default:
			log.Fatal(event.Action)
		}
	}
	return
}

// TO CHECK
func (mdl *CreditManager) getCollateralAmount(blockNum int64, mainAction *schemas.AccountOperation) *big.Int {
	balances := map[string]*big.Int{}
	for _, event := range mainAction.MultiCall {
		if event.Action == "AddCollateral(address,address,uint256)" {
			for token, amount := range *event.Transfers {
				if balances[token] == nil {
					balances[token] = new(big.Int)
				}
				balances[token] = new(big.Int).Add(balances[token], amount)
			}
		}
	}
	tokens := make([]string, 0, len(balances)+1)
	for token := range balances {
		tokens = append(tokens, token)
	}
	underlyingToken := mdl.GetUnderlyingToken()
	if balances[underlyingToken] == nil {
		tokens = append(tokens, underlyingToken)
	}
	//
	prices := mdl.Repo.GetPricesInUSD(blockNum, tokens)
	underlyingDecimals := mdl.GetUnderlyingDecimal()
	//
	totalValue := new(big.Float)
	// sigma(tokenAmount(i)*price(i)/exp(tokendecimals- underlyingToken))/price(underlying)
	for token, amount := range balances {
		if token == underlyingToken { // directly add collateral for underlying token
			continue
		}
		calcValue := utils.GetFloat64(amount, -1*underlyingDecimals)
		nomunerator := new(big.Float).Mul(calcValue, big.NewFloat(prices[token]))
		//
		tokenDecimals := utils.GetExpFloat(mdl.Repo.GetToken(token).Decimals)
		//
		totalValue = new(big.Float).Add(totalValue, new(big.Float).Quo(nomunerator, tokenDecimals))
	}
	initialAmount, _ := new(big.Float).Quo(totalValue, big.NewFloat(prices[underlyingToken])).Int(nil)

	if balances[underlyingToken] != nil { // directly add collateral for underlying token
		initialAmount = new(big.Int).Add(initialAmount, balances[underlyingToken])
	}
	if initialAmount == nil || initialAmount.Cmp(new(big.Int)) == 0 {
		log.Fatal("Collateral for opencreditaccount v2 is zero or nil")
	}
	return initialAmount
}

func (mdl *CreditManager) addCollteralForOpenCreditAccount(blockNum int64, mainAction *schemas.AccountOperation) {
	collateral := mdl.getCollateralAmount(blockNum, mainAction)
	(*mainAction.Args)["amount"] = collateral.String()
	mdl.Repo.UpdateCreditSession(mainAction.SessionId, map[string]interface{}{"InitialAmount": collateral})
}
