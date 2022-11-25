package credit_manager

import (
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

// multicall
func (mdl *CreditManager) addMulticallToMainEvent(mainEvent *schemas.AccountOperation, allMulticalls []*schemas.AccountOperation) {
	txHash := mainEvent.TxHash
	//
	var eventsMulticalls []*schemas.AccountOperation
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
