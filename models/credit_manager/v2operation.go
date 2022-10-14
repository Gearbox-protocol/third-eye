package credit_manager

import (
	"fmt"
	"math/big"
	"sort"
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
	mdl.addBorrowAmountForBlock(borrowAmount)
}
func (mdl *CreditManager) addBorrowAmountForBlock(borrowAmount *big.Int) {
	if mdl.borrowedAmountForBlock == nil {
		mdl.borrowedAmountForBlock = new(big.Int)
	}
	mdl.borrowedAmountForBlock = new(big.Int).Add(mdl.borrowedAmountForBlock, borrowAmount)
}
func (mdl *CreditManager) getBorrowAmountForBlockAndClear() *big.Int {
	if mdl.borrowedAmountForBlock == nil {
		return new(big.Int)
	}
	lastValue := mdl.borrowedAmountForBlock
	mdl.borrowedAmountForBlock = new(big.Int)
	return lastValue
}

// multicall
func (mdl *CreditManager) multiCallHandler(mainEvent *schemas.AccountOperation) {
	account := strings.Split(mainEvent.SessionId, "_")[0]
	txHash := mainEvent.TxHash
	mainactionWithMulticall := mdl.Repo.GetExecuteParser().GetMainEventLogs(txHash, mdl.GetCreditFacadeAddr())
	if len(mainactionWithMulticall) != 1 { // if more than oneAction per tx
		log.Fatal(utils.ToJson(mainactionWithMulticall), utils.ToJson(mainEvent))
	}
	mainCall := mainactionWithMulticall[0] // assuming only one action per tx
	var mainEventFromCall string
	switch mainCall.Name {
	case "multicall":
		mdl.setUpdateSession(mainEvent.SessionId)
		mainEventFromCall = "MultiCallStarted(address)"
	case "openCreditAccountMulticall":
		mdl.setUpdateSession(mainEvent.SessionId)
		mainEventFromCall = "OpenCreditAccount(address,address,uint256,uint16)"
	case "liquidateCreditAccount", "liquidateExpiredCreditAccount":
		mdl.setLiquidateStatus(mainEvent.SessionId, mainCall.Name == "liquidateExpiredCreditAccount")
		mainEventFromCall = "LiquidateCreditAccount(address,address,address,uint256)"
	case "closeCreditAccount":
		mainEventFromCall = "CloseCreditAccount(address,address)"
	}
	if mainEventFromCall != mainEvent.Action { // if the mainaction name is different for events(parsed with eth rpc) and calls (received from tenderly)
		msg := fmt.Sprintf("Tenderly event %s is different from %s", mainCall.Name, mainEvent.Action)
		log.Fatal(msg)
	}
	events := mdl.multicall.PopMulticallEventsV2()
	//
	if !mainCall.SameLenAsEvents(events) {
		log.Fatalf("%s expected %d of multi calls, but third-eye detected %d. Events: %s. Calls: %s. txhash: %s",
			mainCall.Name, mainCall.Len(), len(events),
			utils.ToJson(events), mainCall.String(), mainEvent.TxHash)
	}
	//
	executeEvents := []ds.ExecuteParams{}
	var multicalls []*schemas.AccountOperation
	for _, event := range events {
		if event.BlockNumber != mainEvent.BlockNumber || event.TxHash != txHash {
			log.Fatal("%s has different blockNumber or txhash from opencreditaccount(%d, %s)",
				utils.ToJson(event), mainEvent.BlockNumber, txHash)
		}

		switch event.Action {
		case "AddCollateral(address,address,uint256)":
			if event.Borrower == mainEvent.Borrower {
				multicalls = append(multicalls, event)
				// add collateral can have different borrower then the mainaction user/borrower.
				// related to issue #37.
			} else {
				mdl.Repo.AddAccountOperation(event)
			}
		case "TokenEnabled(address,address)",
			"DisableToken(address,address)",
			"IncreaseBorrowedAmount(address,uint256)",
			"DecreaseBorrowedAmount(address,uint256)":
			multicalls = append(multicalls, event)
		case "ExecuteOrder":
			executeEvents = append(executeEvents, ds.ExecuteParams{
				SessionId:     mainEvent.SessionId,
				CreditAccount: common.HexToAddress(account),
				Protocol:      common.HexToAddress(event.Dapp),
				Borrower:      common.HexToAddress(mainEvent.Borrower),
				Index:         event.LogId,
				BlockNumber:   event.BlockNumber,
			})
		default:
			log.Fatal(utils.ToJson(event))
		}
	}
	//
	multicalls = append(multicalls, mdl.getProcessedExecuteEvents(txHash, executeEvents)...)
	sort.Slice(multicalls, func(i, j int) bool { return multicalls[i].LogId < multicalls[j].LogId })
	//
	mainEvent.MultiCall = multicalls
	// calculate initialAmount on open new credit creditaccount
	if mainEvent.Action == "OpenCreditAccount(address,address,uint256,uint16)" {
		mdl.addCollteralForOpenCreditAccount(mainEvent.BlockNumber, mainEvent)
	}
	mdl.Repo.AddAccountOperation(mainEvent)
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
	borrowAmount *big.Int,
	referralCode uint16) error {
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

// while closing funds can be transferred from the owner account too
// https://github.com/Gearbox-protocol/contracts-v2/blob/main/contracts/credit/CreditManager.sol#L286-L291
func (mdl *CreditManager) onCloseCreditAccountV2(txLog *types.Log, owner, to string) error {
	mdl.State.TotalClosedAccounts++ // update totalclosedStats

	sessionId := mdl.GetCreditOwnerSession(owner)
	account := strings.Split(sessionId, "_")[0]
	cmAddr := txLog.Address.Hex()
	blockNum := int64(txLog.BlockNumber)

	//////////
	// get token transfer when account was closed
	transfers := mdl.Repo.GetExecuteParser().GetTransfers(txLog.TxHash.Hex(), owner, account, mdl.GetUnderlyingToken(), []string{owner, to})
	balances := mdl.toJsonBalance(transfers)

	//////////
	// calculate remainingFunds
	var tokens []string
	for token := range *balances {
		tokens = append(tokens, token)
	}
	tokens = append(tokens, mdl.GetUnderlyingToken())
	prices := mdl.Repo.GetPricesInUSD(blockNum, tokens)
	remainingFunds := (balances.ValueInUnderlying(
		mdl.GetUnderlyingToken(), mdl.GetUnderlyingDecimal(), prices))

	//////////
	// use remainingFunds
	action, args := mdl.ParseEvent("CloseCreditAccount", txLog)
	(*args)["remainingFunds"] = (*core.BigInt)(remainingFunds)
	accountOperation := &schemas.AccountOperation{ // add account operation
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

	////////////////////
	//// update from there
	////////////////////

	session := mdl.Repo.UpdateCreditSession(sessionId, nil) // update session
	session.Balances = balances

	mdl.multicall.AddCloseOrLiquidateEvent(accountOperation) // add event to multicall processor
	session.RemainingFunds = (*core.BigInt)(remainingFunds)

	mdl.ClosedSessions[sessionId] = &SessionCloseDetails{ // update closeSession map with session details
		LogId:          txLog.Index,
		RemainingFunds: remainingFunds,
		Status:         schemas.Closed,
		TxHash:         txLog.TxHash.Hex(),
		Borrower:       owner,
	}

	mdl.RemoveCreditOwnerSession(owner) // remove session to manager object
	mdl.closeAccount(sessionId, blockNum, txLog.TxHash.Hex(), txLog.Index)
	return nil
}

func (mdl *CreditManager) toJsonBalance(z core.Transfers) *core.DBBalanceFormat {
	dbFormat := core.DBBalanceFormat{}
	for token, amt := range z {
		dbFormat[token] = core.CoreIntBalance{
			BI:        (*core.BigInt)(amt),
			F:         utils.GetFloat64Decimal(amt, mdl.Repo.GetToken(token).Decimals),
			IsAllowed: true,
			IsEnabled: true,
		}
	}
	return &dbFormat
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
	// add event to multicall processor
	mdl.multicall.AddCloseOrLiquidateEvent(accountOperation)
	mdl.ClosedSessions[sessionId] = &SessionCloseDetails{
		LogId:          txLog.Index,
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

func (mdl *CreditManager) setLiquidateStatus(sessionId string, isExpired bool) {
	status := schemas.Liquidated
	if mdl.State.Paused {
		status = schemas.LiquidatePaused
	} else if isExpired {
		status = schemas.LiquidateExpired
	}
	mdl.ClosedSessions[sessionId].Status = status
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
	mdl.addBorrowAmountForBlock(amount)
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
	sessionId := mdl.GetCreditOwnerSession(borrower.Hex(), true) // for borrower = creditfacade, session id is ""
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

func (mdl *CreditManager) enableOrDisableToken(txLog types.Log, action string) {
	borrower := common.BytesToAddress(txLog.Topics[1][:]).Hex()
	token := common.BytesToAddress(txLog.Topics[2][:]).Hex()
	//
	sessionId := mdl.GetCreditOwnerSession(borrower)
	//
	accountOperation := &schemas.AccountOperation{
		TxHash:      txLog.TxHash.Hex(),
		BlockNumber: int64(txLog.BlockNumber),
		LogId:       txLog.Index,
		Borrower:    borrower,
		SessionId:   sessionId,
		AdapterCall: false,
		Action:      action,
		Args:        &core.Json{"token": token},
		Dapp:        txLog.Address.Hex(),
	}
	mdl.multicall.AddMulticallEvent(accountOperation)
}
