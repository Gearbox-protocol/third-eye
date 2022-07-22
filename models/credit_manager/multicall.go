package credit_manager

import (
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

type MultiCallProcessor struct {
	borrower            string
	txHash              string
	OpenEvent           *schemas.AccountOperation
	events              []*schemas.AccountOperation
	nonMultiCallEvents  []*schemas.AccountOperation
	MultiCallStartEvent *schemas.AccountOperation
	running             bool
}

func (p *MultiCallProcessor) AddMulticallEvent(operation *schemas.AccountOperation) {
	if !p.running { // non multicall
		// open credit account without multicall (done to calculate initialamount)
		if p.OpenEvent != nil && operation.Action == "AddCollateral(address,address,uint256)" {
			if len(p.OpenEvent.MultiCall) != 0 {
				log.Info("previous addcollateral for openevent found", utils.ToJson(operation))
			}
			p.OpenEvent.MultiCall = make([]*schemas.AccountOperation, 0, 1)
			p.OpenEvent.MultiCall = append(p.OpenEvent.MultiCall, operation)
		} else {
			p.nonMultiCallEvents = append(p.nonMultiCallEvents, operation)
		}
	} else { // multicall
		if operation.TxHash != p.txHash {
			log.Info("While multicall is running, event(%s) has different txhash %s", utils.ToJson(p.events), operation.TxHash)
		}
		operation.Borrower = p.borrower
		p.events = append(p.events, operation)
	}
}
func (p *MultiCallProcessor) AddOpenEvent(openEvent *schemas.AccountOperation) {
	if len(p.events) > 0 {
		log.Info("Previous multicall events not processed", utils.ToJson(p.events))
	}
	if len(p.nonMultiCallEvents) > 0 {
		log.Fatal("There can't be non multicall events while multicall is running",
			utils.ToJson(p.nonMultiCallEvents))
	}
	if p.OpenEvent != nil {
		log.Fatal("2 opencreditaccount event are in same txhash", utils.ToJson(p.OpenEvent), utils.ToJson(openEvent))
	}
	p.OpenEvent = openEvent
}
func (p *MultiCallProcessor) Start(borrower, txHash string, startEvent *schemas.AccountOperation) {
	if len(p.events) > 0 {
		log.Infof("Previous multicall events not processed %s", utils.ToJson(p.events))
	}
	if len(p.nonMultiCallEvents) > 0 {
		log.Fatal("There can't be non multicall events while multicall is running",
			utils.ToJson(p.nonMultiCallEvents))
	}
	p.txHash = txHash
	p.borrower = borrower
	p.MultiCallStartEvent = startEvent
	p.running = true
}
func (p *MultiCallProcessor) lenOfMultiCalls() int {
	return len(p.events)
}
func (p *MultiCallProcessor) End() {
	if !p.running {
		log.Fatal("Multicall end called though multicall not running")
	}
	if len(p.nonMultiCallEvents) > 0 {
		log.Fatal("There can't be non multicall events while multicall is running")
	}
	p.running = false
}
func (p *MultiCallProcessor) PopMulticallEventsV2() []*schemas.AccountOperation {
	collaterals := p.events
	p.events = []*schemas.AccountOperation{}
	return collaterals
}

func (p *MultiCallProcessor) popNonMulticallEventsV2() []*schemas.AccountOperation {
	calls := p.nonMultiCallEvents
	p.nonMultiCallEvents = []*schemas.AccountOperation{}
	return calls
}
