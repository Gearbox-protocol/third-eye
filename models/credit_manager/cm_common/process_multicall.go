package cm_common

import (
	"strings"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/common"
)

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
func (mdl CommonCMAdapter) ProcessRemainingMultiCalls(version core.VersionType, lastTxHash string, nonMultiCallExecuteEvents []ds.ExecuteParams) {

	accountQuotaMgr := mdl.Repo.GetAccountQuotaMgr()
	facadeActions, openEventWithoutMulticall, partialLiqAccount := mdl.MulticallMgr.PopMainActions(lastTxHash, accountQuotaMgr)

	// only for v2/v2.10
	for _, entry := range openEventWithoutMulticall {
		// opencreditaccount without mulitcall
		openWithoutMC := entry.Data
		mdl.SetSessionIsUpdated(openWithoutMC.SessionId)
		mdl.Repo.AddAccountOperation(openWithoutMC)
		mdl.AddCollateralForOpenCreditAccount(openWithoutMC.BlockNumber, core.NewVersion(2), openWithoutMC)
	}
	// TRACE-LOGIC
	if !ds.CallTraceAllowed(mdl.Client) {
		return
	}
	if len(facadeActions) > 0 { // account operation will only exist if there are one or more facade actions
		mainCalls := mdl.Repo.GetExecuteParser().GetMainCalls(lastTxHash, mdl.GetCreditFacadeAddr())
		fixedFacadeActions := mdl.fixFacadeActionStructureViaTenderlyCalls(&mainCalls, facadeActions, partialLiqAccount)
		mdl.validateAndSaveFacadeActions(version, lastTxHash, fixedFacadeActions, mainCalls, nonMultiCallExecuteEvents)
	} else if len(nonMultiCallExecuteEvents) > 0 {
		mdl.SaveExecuteEvents(lastTxHash, nonMultiCallExecuteEvents)
	}
}

func (mdl CommonCMAdapter) ProcessNonMultiCalls() (executeEvents []ds.ExecuteParams) {
	events := mdl.MulticallMgr.PopNonMulticallEvents()

	for _, event := range events {
		switch event.Action {

		case "AddCollateral(address,address,uint256)",
			// v2
			"IncreaseBorrowedAmount(address,uint256)",
			"TokenEnabled(address,address)",
			"TokenDisabled(address,address)",
			"DecreaseBorrowedAmount(address,uint256)",
			// v3
			"IncreaseDebt(address,uint256)",
			"DecreaseDebt(address,uint256)",
			"UpdateQuota",
			"WithdrawCollateral(address,address,uint256,address)":
			mdl.SetSessionIsUpdated(event.SessionId)
			mdl.Repo.AddAccountOperation(event)
		case "ExecuteOrder":
			account := strings.Split(event.SessionId, "_")[0]
			mdl.SetSessionIsUpdated(event.SessionId)
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
