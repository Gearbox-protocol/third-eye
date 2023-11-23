package multicall_processor

import (
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
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

func (p *MultiCallProcessorv3) End() {
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
// open call have the multicalls in them
// liquidated, closed and directly multicalls are separated entries
func (p *MultiCallProcessorv3) PopMainActions() (facadeActions, openEventWithoutMulticall []*FacadeAccountAction) {
	defer func() { p.facadeActions = nil }()
	p.noOfOpens = 0
	for _, entry := range p.facadeActions {
		facadeActions = append(facadeActions, entry)
		if entry.Type == GBFacadeOpenEvent && entry.Data != nil && // only for open credit accounts without multicalls // v2
			entry.withoutMC {
			openEventWithoutMulticall = append(openEventWithoutMulticall, entry)
		} else {
			facadeActions = append(facadeActions, entry)
		}
	}
	return
}

func (p *MultiCallProcessorv3) PopNonMulticallEvents() []*schemas.AccountOperation {
	calls := p.nonMultiCallEvents
	p.nonMultiCallEvents = nil
	return calls
}
