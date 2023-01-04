package credit_manager

import (
	"fmt"
	"strings"

	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/common"
)

func (mdl *CreditManager) fixFacadeActionStructureViaTenderlyCalls(mainCalls []*ds.FacadeCallNameWithMulticall,
	facadeActions []*ds.FacadeAccountActionv2) (result []*ds.FacadeAccountActionv2) { // facadeEvents from rpc, mainCalls from tenderly
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
			if action.Type != ds.GBv2FacadeOpenEvent {
				log.Fatal()
			}
			result = append(result, action)
		case ds.FacadeMulticallCall:
			result = append(result, action)
		case ds.FacadeLiquidateCall, ds.FacadeLiquidateExpiredCall, ds.FacadeCloseAccountCall:
			if mainCall.LenOfMulticalls() != 0 && len(facadeActions) > ind+1 { // combine next facadeAccountAction with current,
				// if number of multicall reported by tenderly are more than 0 for close,expiredliquidate or liquidate calls.
				// this first action is multicall so just take the executeOrders from it.
				multicallToAttach := action.GetMulticallsFromFA()
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

// check name
// check multicall for facade action vs tenderly response
// add to db
func (mdl *CreditManager) validateAndSaveFacadeActions(txHash string, facadeActions []*ds.FacadeAccountActionv2, mainCalls []*ds.FacadeCallNameWithMulticall) {
	executeParams := []ds.ExecuteParams{}
	for ind, mainAction := range facadeActions {
		mainEvent := mainAction.Data

		mainCall := mainCalls[ind]
		var mainEventFromCall string
		switch mainCall.Name {
		case ds.FacadeMulticallCall:
			mdl.setUpdateSession(mainEvent.SessionId)
			mainEventFromCall = "MultiCallStarted(address)"
		case ds.FacadeOpenMulticallCall:
			mdl.setUpdateSession(mainEvent.SessionId)
			mainEventFromCall = "OpenCreditAccount(address,address,uint256,uint16)"
		case ds.FacadeLiquidateCall, ds.FacadeLiquidateExpiredCall:
			mdl.setLiquidateStatus(mainEvent.SessionId, mainCall.Name == ds.FacadeLiquidateExpiredCall)
			mainEventFromCall = "LiquidateCreditAccount(address,address,address,uint256)"
		case ds.FacadeCloseAccountCall:
			mainEventFromCall = "CloseCreditAccount(address,address)"
		}
		if mainEventFromCall != mainEvent.Action { // if the mainaction name is different for events(parsed with eth rpc) and calls (received from tenderly)
			msg := fmt.Sprintf("Tenderly call(%s)is different from facade event(%s)", mainCall.Name, mainEvent.Action)
			log.Fatal(msg)
		}
		//
		eventMulticalls := mainAction.GetMulticallsFromFA()
		if !mainCall.SameMulticallLenAsEvents(eventMulticalls) {
			log.Fatalf("%s expected %d multicalls, but third-eye detected %d. Events: %s. Calls: %s. txhash: %s",
				mainCall.Name, mainCall.LenOfMulticalls(), len(eventMulticalls),
				utils.ToJson(eventMulticalls), mainCall.String(), mainEvent.TxHash)
		}
		account := strings.Split(mainEvent.SessionId, "_")[0]
		for _, event := range eventMulticalls {
			if event.Action == "ExecuteOrder" {
				executeParams = append(executeParams, ds.ExecuteParams{
					SessionId:     mainEvent.SessionId,
					CreditAccount: common.HexToAddress(account),
					Protocol:      common.HexToAddress(event.Dapp),
					Borrower:      common.HexToAddress(mainEvent.Borrower),
					Index:         event.LogId,
					BlockNumber:   event.BlockNumber,
				})
			}
		}
	}
	tenderlyExecuteEvents := mdl.getExecuteOrderAccountOperationFromParams(txHash, executeParams)
	// called for  open_with_multicall, multicall, liquidate, close
	var ind int
	for _, mainAction := range facadeActions {
		multicalls := mainAction.GetMulticallsFromFA()
		for multicallInd, innerEvent := range multicalls {
			if innerEvent.Action == "ExecuteOrder" {
				multicalls[multicallInd] = tenderlyExecuteEvents[ind]
				ind++
			}
		}
		mainEvent := mainAction.Data
		mdl.addMulticallToMainEvent(mainEvent, multicalls)
		mdl.Repo.AddAccountOperation(mainEvent)
	}
}

// multicall
func (mdl *CreditManager) addMulticallToMainEvent(mainEvent *schemas.AccountOperation, allMulticalls []*schemas.AccountOperation) {
	txHash := mainEvent.TxHash
	//
	eventsMulticalls := make([]*schemas.AccountOperation, 0, len(allMulticalls))
	for _, event := range allMulticalls {
		if event.BlockNumber != mainEvent.BlockNumber || event.TxHash != txHash {
			log.Fatal("%s has different blockNumber or txhash from opencreditaccount(%d, %s)",
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
			"DisableToken(address,address)",
			"IncreaseBorrowedAmount(address,uint256)",
			"DecreaseBorrowedAmount(address,uint256)":
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
		mdl.addCollteralForOpenCreditAccount(mainEvent.BlockNumber, mainEvent)
	}
}
