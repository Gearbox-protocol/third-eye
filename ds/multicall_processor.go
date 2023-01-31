package ds

import (
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

const (
	GBv2FacadeOpenEvent = iota
	GBv2FacadeMulticallEvent
	GBv2FacadeCloseEvent
)

// facade Action
type FacadeAccountActionv2 struct {
	Data      *schemas.AccountOperation
	Type      int
	withoutMC bool
	// end if MulticallFinished is emitted / addcollateral is emitted after openCreditAccountWithoutMulticall
	ended bool
}

func (v FacadeAccountActionv2) LenofMulticalls() int {
	return len(v.Data.MultiCall)
}
func (v FacadeAccountActionv2) GetMulticallsFromFA() []*schemas.AccountOperation {
	return v.Data.MultiCall
}
func (v *FacadeAccountActionv2) SetMulticalls(ops []*schemas.AccountOperation) {
	v.Data.MultiCall = ops
}

type MultiCallProcessor struct {
	// borrower            string
	running            bool // is the multicall running
	nonMultiCallEvents []*schemas.AccountOperation
	noOfOpens          int
	facadeActions      []*FacadeAccountActionv2
}

// edge case it adds non multicall addCollateral for open credit account
func (p *MultiCallProcessor) AddMulticallEvent(operation *schemas.AccountOperation) {
	lastMainAction := p.lastMainAction()
	//
	if !p.running { // non multicall
		// open credit account without multicall (done to calculate initialamount)
		if lastMainAction != nil && lastMainAction.Type == GBv2FacadeOpenEvent &&
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

func (p *MultiCallProcessor) AddOpenEvent(openEvent *schemas.AccountOperation) {
	if p.noOfOpens > 0 {
		log.Fatal("2 opencreditaccount event are in same txhash", utils.ToJson(p.facadeActions), utils.ToJson(openEvent))
	}
	p.facadeActions = append(p.facadeActions, &FacadeAccountActionv2{
		Data: openEvent,
		Type: GBv2FacadeOpenEvent,
	})
	p.noOfOpens++
}

func (p *MultiCallProcessor) lastMainAction() *FacadeAccountActionv2 {
	if len(p.facadeActions) > 0 {
		return p.facadeActions[len(p.facadeActions)-1]
	}
	return nil
}

func (p *MultiCallProcessor) Start(txHash string, startEvent *schemas.AccountOperation) {
	lastMainAction := p.lastMainAction()
	if p.running {
		log.Fatalf("Previously started multicall(%s) is not ended for txHash(%s)",
			utils.ToJson(lastMainAction), txHash)
	}
	if lastMainAction == nil || lastMainAction.ended {
		p.facadeActions = append(p.facadeActions, &FacadeAccountActionv2{
			Data: startEvent,
			Type: GBv2FacadeMulticallEvent,
		})
	}
	p.running = true
}

func (p *MultiCallProcessor) AddCloseOrLiquidateEvent(event *schemas.AccountOperation) {
	p.facadeActions = append(p.facadeActions, &FacadeAccountActionv2{
		Data: event,
		Type: GBv2FacadeCloseEvent,
	})
}

func (p *MultiCallProcessor) End() {
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
func (p *MultiCallProcessor) PopMainActionsv2() (facadeActions, openEventWithoutMulticall []*FacadeAccountActionv2) {
	defer func() { p.facadeActions = nil }()
	p.noOfOpens = 0
	for _, entry := range p.facadeActions {
		if entry.Type == GBv2FacadeOpenEvent && entry.Data != nil && // only for open credit accounts without multicalls
			entry.withoutMC {
			openEventWithoutMulticall = append(openEventWithoutMulticall, entry)
		} else {
			facadeActions = append(facadeActions, entry)
		}
	}
	return
}

func (p *MultiCallProcessor) PopNonMulticallEventsV2() []*schemas.AccountOperation {
	calls := p.nonMultiCallEvents
	p.nonMultiCallEvents = nil
	return calls
}
