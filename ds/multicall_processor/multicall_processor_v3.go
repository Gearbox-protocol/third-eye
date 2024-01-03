package multicall_processor

import (
	"sort"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/common"
)

type MultiCallProcessorv3 struct {
	// borrower            string
	running            bool // is the multicall running
	nonMultiCallEvents []*schemas.AccountOperation
	noOfOpens          int
	facadeActions      []*FacadeAccountAction
}

// edge case it adds non multicall addCollateral for open credit account
func (p *MultiCallProcessorv3) AddMulticallEvent(operation *schemas.AccountOperation) {
	lastMainAction := p.lastMainAction()
	//
	if !p.running { // non multicall
		p.nonMultiCallEvents = append(p.nonMultiCallEvents, operation)
	} else { // multicall
		lastMainAction.Data.MultiCall = append(lastMainAction.Data.MultiCall, operation)
	}
}

func (p *MultiCallProcessorv3) AddOpenEvent(openEvent *schemas.AccountOperation) {
	if p.noOfOpens > 0 {
		log.Fatal("2 opencreditaccount event are in same txhash", utils.ToJson(p.facadeActions), utils.ToJson(openEvent))
	}
	p.facadeActions = append(p.facadeActions, &FacadeAccountAction{
		Data: openEvent,
		Type: GBFacadeOpenEvent,
	})
	p.noOfOpens++
}

func (p *MultiCallProcessorv3) lastMainAction() *FacadeAccountAction {
	if len(p.facadeActions) > 0 {
		return p.facadeActions[len(p.facadeActions)-1]
	}
	return nil
}

func (p *MultiCallProcessorv3) Start(txHash string, startEvent *schemas.AccountOperation) {
	lastMainAction := p.lastMainAction()
	if p.running {
		log.Fatalf("Previously started multicall(%s) is not ended for txHash(%s)",
			utils.ToJson(lastMainAction), txHash)
	}
	if lastMainAction == nil || lastMainAction.ended { // since open is before the multicall.
		p.facadeActions = append(p.facadeActions, &FacadeAccountAction{
			Data: startEvent,
			Type: GBFacadeMulticallEvent,
		})
	}
	p.running = true
}

func (p *MultiCallProcessorv3) AddCloseEvent(event *schemas.AccountOperation) {
	p.facadeActions = append(p.facadeActions, &FacadeAccountAction{
		Data: event,
		Type: GBFacadeCloseEvent,
	})
}
func (p *MultiCallProcessorv3) AddLiquidateEvent(event *schemas.AccountOperation) {
	p.facadeActions = append(p.facadeActions, &FacadeAccountAction{
		Data: event,
		Type: GBFacadev3LiqUpdateEvent,
	})
}

func (p *MultiCallProcessorv3) End(logId uint) {
	if !p.running {
		log.Fatal("Multicall end called though multicall not running")
	}
	lastMainAction := p.lastMainAction()
	if lastMainAction != nil {
		lastMainAction.ended = true
		lastMainAction.logId = logId
	}
	p.running = false
}

// pops
// - facadeActions are openWithMulticall, closed, liquidated and multicall actions
// - open call without multicalls
// open call have the multicalls in them
// liquidated, closed and directly multicalls are separated entries
func (p *MultiCallProcessorv3) PopMainActions(txHash string, mgr *ds.AccountQuotaMgr) (facadeActions, openEventWithoutMulticall []*FacadeAccountAction) {
	defer func() { p.facadeActions = nil }()

	p.noOfOpens = 0
	for _, entry := range p.facadeActions {
		facadeActions = append(facadeActions, entry)
	}

	{
		quotas := mgr.GetUpdateQuotaEventForAccount(common.HexToHash(txHash))
		// update quota with session based on accountquotamgr
		for i := len(facadeActions) - 1; i >= 0; i-- {
			facade := facadeActions[i]
			if utils.Contains([]string{
				"StartMultiCall(address,address)",
				"OpenCreditAccount(address,address,address,uint256)", // if open is added then startMulticall event is not present.
			}, facade.Data.Action) {
				splitInd := sort.Search(len(quotas), func(n int) bool {
					return quotas[n].Index > facade.Data.LogId
				})
				addQuotasToFacade(facade, quotas[splitInd:])
				quotas = quotas[:splitInd]
			}
		}
		if len(quotas) != 0 {
			log.Fatal(utils.ToJson(quotas), utils.ToJson(facadeActions))
		}
	}
	return
}

func (p *MultiCallProcessorv3) PopNonMulticallEvents() []*schemas.AccountOperation {
	calls := p.nonMultiCallEvents
	p.nonMultiCallEvents = nil
	return calls
}

func addQuotasToFacade(action *FacadeAccountAction, quotaEvents []*ds.UpdateQuotaEvent) {
	for _, quotaEvent := range quotaEvents {
		if action.logId == 0 || action.logId > quotaEvent.Index { // if there is a liquidation event then at closure there is updatequota emitted outside of the start/end multicall events on facade. So, we need to ignore that updateQuota event.
			action.Data.MultiCall = append(action.Data.MultiCall, &schemas.AccountOperation{
				TxHash:      quotaEvent.TxHash,
				BlockNumber: quotaEvent.BlockNumber,
				LogId:       quotaEvent.Index,
				Borrower:    action.Data.Borrower,
				SessionId:   action.Data.SessionId,
				Dapp:        action.Data.Dapp,
				Action:      "UpdateQuota",
				Args:        &core.Json{"token": quotaEvent.Token, "change": (*core.BigInt)(quotaEvent.QuotaChange)},
				AdapterCall: false,
				// Transfers:   &core.Transfers{tx.Token: amount},
			})
		}
	}
	sort.Slice(action.Data.MultiCall, func(i, j int) bool {
		return action.Data.MultiCall[i].LogId < action.Data.MultiCall[j].LogId
	})
}
