package credit_manager

import (
	"fmt"
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

func (mdl *CreditManager) CMStatsOnOpenAccount(borrowAmount *big.Int) {
	// manager state
	mdl.State.TotalOpenedAccounts++
	mdl.State.OpenedAccountsCount++
	mdl.State.TotalBorrowedBI = core.AddCoreAndInt(mdl.State.TotalBorrowedBI, borrowAmount)
	mdl.State.TotalBorrowed = utils.GetFloat64Decimal(mdl.State.TotalBorrowedBI.Convert(), mdl.GetUnderlyingDecimal())
}

// multicall
func (mdl *CreditManager) multiCallHandler(mainAction *schemas.AccountOperation) {
	account := strings.Split(mainAction.SessionId, "_")[0]
	txHash := mainAction.TxHash
	mainEvents := mdl.Repo.GetExecuteParser().GetMainEventLogs(txHash, mdl.GetCreditFacadeAddr())
	if len(mainEvents) != 1 {
		log.Fatal(utils.ToJson(mainEvents))
	}
	var tenderlyEventName string
	switch mainEvents[0].Name {
	case "multicall":
		mdl.setUpdateSession(mainAction.SessionId)
		tenderlyEventName = "MultiCallStarted(address)"
	case "openCreditAccountMulticall":
		mdl.setUpdateSession(mainAction.SessionId)
		tenderlyEventName = "OpenCreditAccount(address,address,uint256,uint256)"
	case "liquidateCreditAccount":
		tenderlyEventName = "LiquidateCreditAccount(address,address,address,uint256)"
	case "closeCreditAccount":
		tenderlyEventName = "CloseCreditAccount(address,address)"
	}
	if tenderlyEventName != mainAction.Action {
		log.Fatalf("Tenderly event %s is different from %s", mainEvents[0].Name, mainAction.Action)
	}
	events := mdl.multicall.PopMulticallEventsV2()
	//
	if len(events) != mainEvents[0].MultiCallsLen {
		log.Fatalf("%s expected %d of multi calls, but third-eye detected %d. Events: %s",
			mainEvents[0].Name, mainEvents[0].MultiCallsLen, len(events), utils.ToJson(events))
	}
	//
	executeEvents := []ds.ExecuteParams{}
	var multicalls []*schemas.AccountOperation
	for _, event := range events {
		if event.BlockNumber != mainAction.BlockNumber || event.TxHash != txHash {
			log.Fatal("%s has different blockNumber or txhash from opencreditaccount(%d, %s)",
				utils.ToJson(event), mainAction.BlockNumber, txHash)
		}

		switch event.Action {
		case "AddCollateral(address,address,uint256)",
			"IncreaseBorrowedAmount(address,uint256)",
			"DecreaseBorrowedAmount(address,uint256)":
			multicalls = append(multicalls, event)
		case "ExecuteOrder":
			executeEvents = append(executeEvents, ds.ExecuteParams{
				SessionId:     mainAction.SessionId,
				CreditAccount: common.HexToAddress(account),
				Protocol:      common.HexToAddress(event.Dapp),
				Borrower:      common.HexToAddress(mainAction.Borrower),
				Index:         event.LogId,
				BlockNumber:   event.BlockNumber,
			})
		default:
			log.Fatal(utils.ToJson(event))
		}
	}
	multicalls = append(multicalls, mdl.getProcessedExecuteEvents(txHash, executeEvents)...)
	mainAction.MultiCall = multicalls
	mdl.Repo.AddAccountOperation(mainAction)
}

func (mdl *CreditManager) getProcessedExecuteEvents(txHash string, executeParams []ds.ExecuteParams) (multiCalls []*schemas.AccountOperation) {
	// credit manager has the execute event
	calls := mdl.Repo.GetExecuteParser().GetExecuteCalls(txHash, mdl.Address, executeParams)
	for i, call := range calls {
		params := executeParams[i]
		// add account operation
		accountOperation := &schemas.AccountOperation{
			BlockNumber: params.BlockNumber,
			TxHash:      txHash,
			LogId:       params.Index,
			// owner/account data
			Borrower:  params.Borrower.Hex(),
			SessionId: params.SessionId,
			// dapp
			Dapp: params.Protocol.Hex(),
			// call/events data
			Action:      call.Name,
			Args:        call.Args,
			AdapterCall: true,
			Transfers:   &call.Transfers,
			// extras
			Depth: call.Depth,
		}
		multiCalls = append(multiCalls, accountOperation)
	}
	return
}

///////////////////////
// Main actions
///////////////////////
func (mdl *CreditManager) onOpenCreditAccountV2(txLog *types.Log, onBehalfOf, account string,
	borrowAmount,
	referralCode *big.Int) error {
	mdl.CMStatsOnOpenAccount(borrowAmount)
	// other operations
	cmAddr := txLog.Address.Hex()
	sessionId := fmt.Sprintf("%s_%d_%d", account, txLog.BlockNumber, txLog.Index)
	blockNum := int64(txLog.BlockNumber)

	// add account operation
	action, args := mdl.ParseEvent("OpenCreditAccount", txLog)
	accountOperation := &schemas.AccountOperation{
		TxHash:      txLog.TxHash.Hex(),
		BlockNumber: blockNum,
		LogId:       txLog.Index,
		Borrower:    onBehalfOf,
		SessionId:   sessionId,
		AdapterCall: false,
		Action:      action,
		Args:        args,
		Transfers: &core.Transfers{
			mdl.GetUnderlyingToken(): borrowAmount,
		},
		Dapp: cmAddr,
	}
	mdl.multicall.AddOpenEvent(accountOperation)
	mdl.PoolBorrow(txLog, sessionId, onBehalfOf, borrowAmount)
	// add session to manager object
	mdl.AddCreditOwnerSession(onBehalfOf, sessionId)
	// create credit session
	newSession := &schemas.CreditSession{
		ID:             sessionId,
		Status:         schemas.Active,
		Borrower:       onBehalfOf,
		CreditManager:  mdl.Address,
		Account:        account,
		Since:          blockNum,
		BorrowedAmount: (*core.BigInt)(borrowAmount),
		IsDirty:        true,
		Version:        2,
	}
	mdl.Repo.AddCreditSession(newSession, false, txLog.TxHash.Hex(), txLog.Index)
	return nil
}

func (mdl *CreditManager) onCloseCreditAccountV2(txLog *types.Log, owner, to string) error {
	mdl.State.TotalClosedAccounts++
	sessionId := mdl.GetCreditOwnerSession(owner)
	account := strings.Split(sessionId, "_")[0]
	//
	cmAddr := txLog.Address.Hex()
	blockNum := int64(txLog.BlockNumber)
	action, args := mdl.ParseEvent("CloseCreditAccount", txLog)
	// add account operation
	transfers := mdl.Repo.GetExecuteParser().GetTransfers(txLog.TxHash.Hex(), owner, account, mdl.GetUnderlyingToken(), []string{owner, to})
	accountOperation := &schemas.AccountOperation{
		TxHash:      txLog.TxHash.Hex(),
		BlockNumber: blockNum,
		LogId:       txLog.Index,
		Borrower:    owner,
		SessionId:   sessionId,
		AdapterCall: false,
		Action:      action,
		Args:        args,
		Transfers:   &transfers,
		Dapp:        cmAddr,
	}
	// process multicalls
	mdl.multiCallHandler(accountOperation)
	// update remainingFunds
	session := mdl.Repo.UpdateCreditSession(sessionId, nil)
	session.Balances = mdl.toJsonBalance(transfers)
	var tokens []string
	for token := range *session.Balances {
		tokens = append(tokens, token)
	}
	tokens = append(tokens, mdl.GetUnderlyingToken())
	prices := mdl.Repo.GetPricesInUSD(blockNum, tokens)
	// log.Info(prices)
	// log.Info(utils.ToJson(session.Balances))
	remainingFunds := (session.Balances.ValueInUnderlying(
		mdl.GetUnderlyingToken(), mdl.GetUnderlyingDecimal(), prices))
	session.RemainingFunds = (*core.BigInt)(remainingFunds)
	mdl.ClosedSessions[sessionId] = &SessionCloseDetails{
		RemainingFunds: remainingFunds,
		Status:         schemas.Closed,
		TxHash:         txLog.TxHash.Hex(),
		Borrower:       owner,
	}
	// remove session to manager object
	mdl.RemoveCreditOwnerSession(owner)
	mdl.closeAccount(sessionId, blockNum, txLog.TxHash.Hex(), txLog.Index)
	return nil
}
func (mdl *CreditManager) toJsonBalance(z core.Transfers) *core.JsonBalance {
	bal := core.JsonBalance{}
	for token, amt := range z {
		bal[token] = &core.BalanceType{
			BI: (*core.BigInt)(amt),
			F:  utils.GetFloat64Decimal(amt, mdl.Repo.GetToken(token).Decimals),
		}
	}
	return &bal
}

func (mdl *CreditManager) getRemainingFundsOnClose(blockNum int64, txHash, borrower string) *core.Transfers {
	// data := mdl.GetCreditSessionData(blockNum-1, borrower)
	// for data.Balances {

	// }
	return nil
}
func (mdl *CreditManager) onLiquidateCreditAccountV2(txLog *types.Log, owner, liquidator string, remainingFunds *big.Int) error {
	mdl.State.TotalLiquidatedAccounts++
	sessionId := mdl.GetCreditOwnerSession(owner)

	blockNum := int64(txLog.BlockNumber)
	action, args := mdl.ParseEvent("LiquidateCreditAccount", txLog)
	// add account operation
	accountOperation := &schemas.AccountOperation{
		TxHash:      txLog.TxHash.Hex(),
		BlockNumber: blockNum,
		LogId:       txLog.Index,
		Borrower:    owner,
		SessionId:   sessionId,
		AdapterCall: false,
		Action:      action,
		Args:        args,
		Transfers: &core.Transfers{
			mdl.GetUnderlyingToken(): remainingFunds,
		},
		Dapp: txLog.Address.Hex(),
	}
	// process multicalls
	mdl.multiCallHandler(accountOperation)
	mdl.ClosedSessions[sessionId] = &SessionCloseDetails{
		RemainingFunds: remainingFunds,
		Status:         schemas.Liquidated,
		TxHash:         txLog.TxHash.Hex(),
		Borrower:       owner,
	}
	session := mdl.Repo.GetCreditSession(sessionId)
	session.Liquidator = liquidator
	session.RemainingFunds = (*core.BigInt)(remainingFunds)
	// remove session to manager object
	mdl.RemoveCreditOwnerSession(owner)
	mdl.closeAccount(sessionId, blockNum, txLog.TxHash.Hex(), txLog.Index)
	return nil
}

///////////////////////
// Side actions that can also be used as multicall events
///////////////////////
func (mdl *CreditManager) onAddCollateralV2(txLog *types.Log, onBehalfOf, token string, value *big.Int) {
	sessionId := mdl.GetCreditOwnerSession(onBehalfOf)
	blockNum := int64(txLog.BlockNumber)
	action, args := mdl.ParseEvent("AddCollateral", txLog)
	// add account operation
	accountOperation := &schemas.AccountOperation{
		TxHash:      txLog.TxHash.Hex(),
		BlockNumber: blockNum,
		LogId:       txLog.Index,
		Borrower:    onBehalfOf,
		SessionId:   sessionId,
		AdapterCall: false,
		Action:      action,
		Args:        args,
		Transfers: &core.Transfers{
			token: value,
		},
		Dapp: txLog.Address.Hex(),
	}
	mdl.multicall.AddMulticallEvent(accountOperation)
	mdl.AddCollateralToSession(blockNum, sessionId, token, value)
}

// amount can be negative, if decrease borrowamount, add pool repay event
func (mdl *CreditManager) onIncreaseBorrowedAmountV2(txLog *types.Log, borrower string, amount *big.Int, eventName string) error {
	// manager state
	mdl.State.TotalBorrowedBI = core.AddCoreAndInt(mdl.State.TotalBorrowedBI, amount)
	mdl.State.TotalBorrowed = utils.GetFloat64Decimal(mdl.State.TotalBorrowedBI.Convert(), mdl.GetUnderlyingDecimal())
	// other operations
	sessionId := mdl.GetCreditOwnerSession(borrower)
	blockNum := int64(txLog.BlockNumber)
	action, args := mdl.ParseEvent(eventName, txLog)
	// add account operation
	accountOperation := &schemas.AccountOperation{
		TxHash:      txLog.TxHash.Hex(),
		BlockNumber: blockNum,
		LogId:       txLog.Index,
		Borrower:    borrower,
		SessionId:   sessionId,
		AdapterCall: false,
		Action:      action,
		Args:        args,
		Transfers: &core.Transfers{
			mdl.GetUnderlyingToken(): amount,
		},
		Dapp: txLog.Address.Hex(),
	}
	mdl.multicall.AddMulticallEvent(accountOperation)
	if amount.Sign() == -1 {
		repayAmount := new(big.Int).Neg(amount)
		// manager state
		mdl.State.TotalRepaidBI = core.AddCoreAndInt(mdl.State.TotalRepaidBI, repayAmount)
		mdl.State.TotalRepaid = utils.GetFloat64Decimal(mdl.State.TotalRepaidBI.Convert(), mdl.GetUnderlyingDecimal())
		mdl.PoolRepay(blockNum, txLog.Index, txLog.TxHash.Hex(), sessionId, borrower, repayAmount)
	} else {
		mdl.PoolBorrow(txLog, sessionId, borrower, amount)
	}
	session := mdl.Repo.UpdateCreditSession(sessionId, nil)
	session.BorrowedAmount = (*core.BigInt)(new(big.Int).Add(session.BorrowedAmount.Convert(), amount))
	return nil
}

func (mdl *CreditManager) AddExecuteParamsV2(txLog *types.Log,
	borrower,
	targetContract common.Address) error {
	sessionId := mdl.GetCreditOwnerSession(borrower.Hex(), true)
	mdl.multicall.AddMulticallEvent(&schemas.AccountOperation{
		BlockNumber: int64(txLog.BlockNumber),
		TxHash:      txLog.TxHash.Hex(),
		LogId:       txLog.Index,
		Borrower:    borrower.Hex(),
		Dapp:        targetContract.Hex(),
		AdapterCall: true,
		SessionId:   sessionId,
		Action:      "ExecuteOrder",
	})
	return nil
}

// copied from v1
func (mdl *CreditManager) onTransferAccountV2(txLog *types.Log, owner, newOwner string) error {
	return mdl.onTransferAccount(txLog, owner, newOwner)
}
