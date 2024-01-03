package multicall_processor

import (
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
)

type MultiCallProcessorv2 struct {
	// borrower            string
	running            bool // is the multicall running
	nonMultiCallEvents []*schemas.AccountOperation
	noOfOpens          int
	facadeActions      []*FacadeAccountAction
}

// edge case it adds non multicall addCollateral for open credit account
func (p *MultiCallProcessorv2) AddMulticallEvent(operation *schemas.AccountOperation) {
	lastMainAction := p.lastMainAction()
	//
	if !p.running { // non multicall
		// open credit account without multicall (done to calculate initialamount)
		if lastMainAction != nil && lastMainAction.Type == GBFacadeOpenEvent &&
			operation.Action == "AddCollateral(address,address,uint256)" && !lastMainAction.ended {
			//
			openEventWithoutMulticall := lastMainAction.Data
			if len(openEventWithoutMulticall.MultiCall) != 0 {
				log.Fatal("previous addcollateral for openevent found", utils.ToJson(operation))
			}
			openEventWithoutMulticall.MultiCall = []*schemas.AccountOperation{operation}
			lastMainAction.withoutMC = true
			lastMainAction.ended = true
		} else {
			p.nonMultiCallEvents = append(p.nonMultiCallEvents, operation)
		}
	} else { // multicall
		lastMainAction.Data.MultiCall = append(lastMainAction.Data.MultiCall, operation)
	}
}

func (p *MultiCallProcessorv2) AddOpenEvent(openEvent *schemas.AccountOperation) {
	if p.noOfOpens > 0 {
		log.Fatal("2 opencreditaccount event are in same txhash", utils.ToJson(p.facadeActions), utils.ToJson(openEvent))
	}
	p.facadeActions = append(p.facadeActions, &FacadeAccountAction{
		Data: openEvent,
		Type: GBFacadeOpenEvent,
	})
	p.noOfOpens++
}

func (p *MultiCallProcessorv2) lastMainAction() *FacadeAccountAction {
	if len(p.facadeActions) > 0 {
		return p.facadeActions[len(p.facadeActions)-1]
	}
	return nil
}

func (p *MultiCallProcessorv2) Start(txHash string, startEvent *schemas.AccountOperation) {
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

func (p *MultiCallProcessorv2) AddCloseEvent(event *schemas.AccountOperation) {
	p.facadeActions = append(p.facadeActions, &FacadeAccountAction{
		Data: event,
		Type: GBFacadeCloseEvent,
	})
}
func (p *MultiCallProcessorv2) AddLiquidateEvent(event *schemas.AccountOperation) {
	p.facadeActions = append(p.facadeActions, &FacadeAccountAction{
		Data: event,
		Type: GBFacadeCloseEvent,
	})
}

func (p *MultiCallProcessorv2) End(logId uint) {
	if !p.running {
		log.Fatal("Multicall end called though multicall not running")
	}
	lastMainAction := p.lastMainAction()
	if lastMainAction != nil {
		lastMainAction.ended = true
	}
	p.running = false
}

// pops
// - facadeActions are openWithMulticall, closed, liquidated and multicall actions
// - open call without multicalls
// - open call have the multicalls in them
// liquidated, closed and directly multicalls are separated entries
func (p *MultiCallProcessorv2) PopMainActions(_ string, _ *ds.AccountQuotaMgr) (facadeActions, openEventWithoutMulticall []*FacadeAccountAction) {
	defer func() { p.facadeActions = nil }()
	p.noOfOpens = 0
	for _, entry := range p.facadeActions {
		if entry.Type == GBFacadeOpenEvent && entry.Data != nil && // only for open credit accounts without multicalls // v2
			entry.withoutMC {
			openEventWithoutMulticall = append(openEventWithoutMulticall, entry)
		} else {
			facadeActions = append(facadeActions, entry)
		}
	}
	return
}

func (p *MultiCallProcessorv2) PopNonMulticallEvents() []*schemas.AccountOperation {
	calls := p.nonMultiCallEvents
	p.nonMultiCallEvents = nil
	return calls
}
