package credit_manager

import (
	"math/big"
	"strings"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/credit_filter"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (cm *CreditManager) addCreditConfigurator(creditConfigurator string) {
	// this is need for mask only
	// cm.Repo.AddCreditManagerToFilter(cm.Address, creditConfigurator)
	cf := credit_filter.NewCreditFilter(creditConfigurator, ds.CreditConfigurator, cm.Address, cm.DiscoveredAt, cm.Client, cm.Repo)
	cm.Repo.AddSyncAdapter(cf)
	cm.Details["configurator"] = creditConfigurator
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
	case core.Topic("LiquidateCreditAccount(address,address,address,uint256)"):
		liquidateCreditAccountEvent, err := mdl.facadeContractV2.ParseLiquidateCreditAccount(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack LiquidateCreditAccount event", err)
		}

		mdl.onLiquidateCreditAccountV2(&txLog,
			liquidateCreditAccountEvent.Borrower.Hex(),
			liquidateCreditAccountEvent.Liquidator.Hex(),
			liquidateCreditAccountEvent.RemainingFunds)
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
			mdl.Repo.GetKit().GetAdapter(oldConfigurator).SetBlockToDisableOn(int64(txLog.BlockNumber))
			mdl.addCreditConfigurator(newConfigurator)
		}
	}
}

func (mdl *CreditManager) onNewTxHashV2() {
	mdl.processRemainingMultiCalls()
	mdl.processNonMultiCalls()
}

// opencreditaccount
// addcollateral
// increase/decase borrow amount
// executeorder
// are added to multicall manager
//
func (mdl *CreditManager) processRemainingMultiCalls() {
	// opencreditaccount
	mainAction := mdl.multicall.OpenEvent
	// if not present use multicall
	if mainAction == nil { // other multicall operations
		mainAction = mdl.multicall.MultiCallStartEvent
		// open credit account without multicall
	} else if mdl.multicall.lenOfMultiCalls() == 0 {
		mdl.setUpdateSession(mainAction.SessionId)
		mdl.Repo.AddAccountOperation(mainAction)
		mdl.openCreditAccountInitialAmount(mainAction.BlockNumber, mainAction)
	}
	//
	if mdl.multicall.lenOfMultiCalls() > 0 {
		mdl.multiCallHandler(mainAction)
	}
	mdl.multicall.OpenEvent = nil
	mdl.multicall.MultiCallStartEvent = nil
}
func (mdl *CreditManager) setUpdateSession(sessionId string) {
	// log.Info(log.DetectFunc(),sessionId, "increased")
	mdl.UpdatedSessions[sessionId]++
}
func (mdl *CreditManager) processNonMultiCalls() {
	events := mdl.multicall.popNonMulticallEventsV2()
	executeEvents := []ds.ExecuteParams{}
	for _, event := range events {
		switch event.Action {
		case "AddCollateral(address,address,uint256)",
			"IncreaseBorrowedAmount(address,uint256)",
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
		}
	}
	if len(executeEvents) > 0 {
		mdl.handleExecuteEvents(executeEvents)
	}
}

func (mdl *CreditManager) getInitialAmount(blockNum int64, mainAction *schemas.AccountOperation) *big.Int {
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
	tokens := make([]string, 0, len(balances))
	for token := range balances {
		tokens = append(tokens, token)
	}
	prices := mdl.Repo.GetPricesInUSD(blockNum, tokens)
	underlyingDecimals := mdl.GetUnderlyingDecimal()
	underlyingToken := mdl.GetUnderlyingToken()
	totalValue := new(big.Float)

	// sigma(tokenAmount(i)*price(i)/exp(tokendecimals- underlyingToken))/price(underlying)
	for token, amount := range balances {
		value := new(big.Float).Mul(new(big.Float).SetInt(amount), big.NewFloat(prices[token]))
		decimals := utils.GetExpFloat(mdl.Repo.GetToken(token).Decimals - underlyingDecimals)
		totalValue = new(big.Float).Add(totalValue, new(big.Float).Quo(value, decimals))
	}
	initialAmount, _ := new(big.Float).Quo(totalValue, big.NewFloat(prices[underlyingToken])).Int(nil)
	return initialAmount
}

func (mdl *CreditManager) openCreditAccountInitialAmount(blockNum int64, mainAction *schemas.AccountOperation) {
	initialAmount := mdl.getInitialAmount(blockNum, mainAction)
	(*mainAction.Args)["amount"] = initialAmount.String()
	mdl.Repo.UpdateCreditSession(mainAction.SessionId, map[string]interface{}{"InitialAmount": initialAmount})
}
