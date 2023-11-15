package cm_common

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"sort"
	"strings"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	mpi "github.com/Gearbox-protocol/third-eye/ds/multicall_processor"
	"github.com/ethereum/go-ethereum/common"
)

// multicalls and liquidate/close/openwithmulticalls are separate data points,
// this function adds multicall to mainFacadeActions
// if that is the correct structure of operation
func (mdl *CommonCMAdapter) fixFacadeActionStructureViaTenderlyCalls(mainCalls []*ds.FacadeCallNameWithMulticall,
	facadeActions []*mpi.FacadeAccountAction) (result []*mpi.FacadeAccountAction) { // facadeEvents from rpc, mainCalls from tenderly
	if len(mainCalls) > len(facadeActions) {
		log.Fatalf("Len of calls(%d) can't be more than separated close/liquidate and multicall(%d).",
			len(mainCalls), len(facadeActions),
		)
	}
	//
	var ind int
	for _, mainCall := range mainCalls {
		action := facadeActions[ind]
		switch mainCall.Name {
		case ds.FacadeOpenMulticallCall:
			if !action.IsOpen() {
				log.Fatal()
			}
			result = append(result, action)
		case ds.FacadeMulticallCall:
			result = append(result, action)
		case ds.FacadeLiquidateCall, ds.FacadeLiquidateExpiredCall, ds.FacadeCloseAccountCall:
			if mainCall.LenOfMulticalls() != 0 && len(facadeActions) > ind+1 { // combine next facadeAccountAction with current,
				// if number of multicall reported by tenderly are more than 0 for close,expiredliquidate or liquidate calls.
				// this first action is multicall so just take the executeOrders from it.
				multicallToAttach := action.GetMulticallsFromEvent()
				action = facadeActions[ind+1]
				action.SetMulticalls(multicallToAttach)
				ind++
			}
			result = append(result, action)
		}
		ind++
	}
	//
	if ind != len(facadeActions) {
		log.Fatalf(`Not able to completely process facade action in tx, 
		mismatch with facade calls we got from tenderly. 
		Len: %d, processed: %d`, len(facadeActions), ind)
	}
	return
}

func (mdl CommonCMAdapter) updateQuotasWithSessionId(sessionId string, mainCall *ds.FacadeCallNameWithMulticall) {
	for _, multicall := range mainCall.GetMulticalls() {
		sig := hex.EncodeToString(multicall.CallData[:4])
		if sig == "712c10ad" { // updateQuota on v3
			quotaEvent := mdl.Repo.GetAccountQuotaMgr().GetUpdateQuotaEventForAccount(strings.Split(sessionId, "_")[0])
			//
			//
			mdl.Repo.AddAccountOperation(&schemas.AccountOperation{
				TxHash:      quotaEvent.TxHash,
				BlockNumber: quotaEvent.BlockNumber,
				LogId:       quotaEvent.Index,
				// Borrower:    session.Borrower,
				SessionId:   sessionId,
				Dapp:        mdl.GetCreditFacadeAddr(),
				Action:      "UpdateQuota",
				Args:        &core.Json{"token": quotaEvent.Token, "change": quotaEvent.QuotaChange},
				AdapterCall: false,
				// Transfers:   &core.Transfers{tx.Token: amount},
			})
			mdl.SetSessionIsUpdated(sessionId)
		}
	}
}

// check name
// check multicall for facade action vs tenderly response
// add to db
func (mdl *CommonCMAdapter) validateAndSaveFacadeActions(version core.VersionType, txHash string,
	facadeActions []*mpi.FacadeAccountAction,
	mainCalls []*ds.FacadeCallNameWithMulticall,
	nonMultiCallExecuteEvents []ds.ExecuteParams) {

	executeParams := []ds.ExecuteParams{} // non multicall and multicall execute orders for a tx to be compared with call trace
	for ind, _mainAction := range facadeActions {
		mainEvent := _mainAction.Data

		mainCall := mainCalls[ind]
		//
		mainEventFromCall := mdl.getEventNameFromCall(version, mainCall.Name, mainEvent.SessionId)

		if mainEventFromCall != mainEvent.Action { // if the mainaction name is different for events(parsed with eth rpc) and calls (received from tenderly)
			msg := fmt.Sprintf("Tenderly call(%s)is different from facade event(%s)", mainCall.Name, mainEvent.Action)
			log.Fatal(msg)
		}
		//
		eventMulticalls := mainEvent.MultiCall
		if !mainCall.SameMulticallLenAsEvents(version, eventMulticalls) {
			log.Fatalf("%s expected %d multicalls, but third-eye detected %d. Events: %s. Calls: %s. txhash: %s",
				mainCall.Name, mainCall.LenOfMulticalls(), len(eventMulticalls),
				utils.ToJson(eventMulticalls), mainCall.String(), mainEvent.TxHash)
		}
		// update quota with session based on calls from tenderly and also mark sesssion as updated
		mdl.updateQuotasWithSessionId(mainEvent.SessionId, mainCall)
		//
		account := strings.Split(mainEvent.SessionId, "_")[0]
		for _, event := range eventMulticalls {
			switch event.Action {
			case "ExecuteOrder":
				executeParams = append(executeParams, ds.ExecuteParams{
					SessionId:     mainEvent.SessionId,
					CreditAccount: common.HexToAddress(account),
					Protocol:      common.HexToAddress(event.Dapp),
					Borrower:      common.HexToAddress(mainEvent.Borrower),
					Index:         event.LogId,
					BlockNumber:   event.BlockNumber,
				})
			case "WithdrawCollateral(address,address,uint256,address)":
				if mainEvent.Action == "LiquidateCreditAccount(address,address,address,address,uint256)" { // REV_COL_LIQ_V3: v3 liquidate reverse the collateral
					// since liquidation the withdraw collateral is not to the account owner.
					mdl.AddCollateralToSession(event.BlockNumber, event.SessionId,
						(*event.Args)["token"].(common.Address).Hex(),
						(*event.Args)["amount"].(*big.Int),
					)
				}
			}
		}
	}

	mdl.executeOperations(txHash, facadeActions, executeParams, nonMultiCallExecuteEvents)
}

// process non multicall execute operations.
// attach multicall execute operations to facade main actions
func (mdl *CommonCMAdapter) executeOperations(txHash string, facadeActions []*mpi.FacadeAccountAction,
	executeParams, nonMultiCallExecuteEvents []ds.ExecuteParams) {
	sort.Slice(executeParams, func(i, j int) bool { return executeParams[i].Index < executeParams[j].Index })
	tenderlyExecOperations := mdl.GetExecuteOrderAccountOperationFromParams(txHash, executeParams)

	// process non multicall execute order operations
	remainingExecOperations := []*schemas.AccountOperation{}
	var nonMulticallInd int
	for _, accountOperation := range tenderlyExecOperations {
		if nonMulticallInd < len(nonMultiCallExecuteEvents) &&
			accountOperation.LogId == nonMultiCallExecuteEvents[nonMulticallInd].Index { // add non multicall execute order to db
			mdl.Repo.AddAccountOperation(accountOperation)
			nonMulticallInd++
		} else {
			remainingExecOperations = append(remainingExecOperations, accountOperation)
		}
	}

	// called for  open_with_multicall, multicall, liquidate, close
	var indTenderlyCall int
	for _, mainAction := range facadeActions {
		multicalls := mainAction.GetMulticallsFromEvent()
		for multicallInd, innerEvent := range multicalls {
			if innerEvent.Action == "ExecuteOrder" {
				if innerEvent.LogId == remainingExecOperations[indTenderlyCall].LogId { // add multicall execute order to main event
					multicalls[multicallInd] = remainingExecOperations[indTenderlyCall]
				} else {
					log.Fatalf("execute order index mismatch: events: %s, calls: %s", utils.ToJson(innerEvent), utils.ToJson(remainingExecOperations[indTenderlyCall]))
				}
				indTenderlyCall++
			}
		}
		mainEvent := mainAction.Data
		mdl.addMulticallToMainEvent(mainEvent, multicalls)
		mdl.Repo.AddAccountOperation(mainEvent)
	}
}

func (mdl CommonCMAdapter) GetExecuteOrderAccountOperationFromParams(txHash string, executeParams []ds.ExecuteParams) (multiCalls []*schemas.AccountOperation) {
	// credit manager has the execute event
	calls := mdl.Repo.GetExecuteParser().GetExecuteCalls(txHash,
		mdl.Address,
		executeParams,
	)
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
		}
		multiCalls = append(multiCalls, accountOperation)
	}
	return
}

// multicall
func (mdl *CommonCMAdapter) addMulticallToMainEvent(mainEvent *schemas.AccountOperation, allMulticalls []*schemas.AccountOperation) {
	txHash := mainEvent.TxHash
	//
	eventsMulticalls := make([]*schemas.AccountOperation, 0, len(allMulticalls))
	for _, event := range allMulticalls {
		if event.BlockNumber != mainEvent.BlockNumber || event.TxHash != txHash {
			log.Fatalf("%s has different blockNumber or txhash from opencreditaccount(%d, %s)",
				utils.ToJson(event), mainEvent.BlockNumber, txHash)
		}

		switch event.Action {
		case "AddCollateral(address,address,uint256)":
			if event.Borrower == mainEvent.Borrower {
				eventsMulticalls = append(eventsMulticalls, event)
				// add collateral can have different borrower then the mainaction user/borrower.
				// related to issue #37.
			} else {
				mdl.Repo.AddAccountOperation(event)
			}
		case "TokenEnabled(address,address)",
			"TokenDisabled(address,address)",
			// for v2
			"IncreaseBorrowedAmount(address,uint256)",
			"DecreaseBorrowedAmount(address,uint256)",
			// for v3
			"IncreaseDebt(address,uint256)",
			"DecreaseDebt(address,uint256)":
			eventsMulticalls = append(eventsMulticalls, event)
		default: // for all the ExecuteOrder
			if event.AdapterCall {
				eventsMulticalls = append(eventsMulticalls, event)
			} else {
				log.Fatal(utils.ToJson(event))
			}
		}
	}
	//
	mainEvent.MultiCall = eventsMulticalls
	// calculate initialAmount on open new credit creditaccount
	if mainEvent.Action == "OpenCreditAccount(address,address,uint256,uint16)" {
		mdl.AddCollateralForOpenCreditAccount(mainEvent.BlockNumber, mainEvent)
	}
}
