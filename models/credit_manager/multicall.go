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
	if !p.running {
		p.nonMultiCallEvents = append(p.nonMultiCallEvents, operation)
		return
	}
	if operation.TxHash != p.txHash {
		log.Info("While multicall is running, event(%s) has different txhash %s", utils.ToJson(p.events), p.txHash)
	}
	operation.Borrower = p.borrower
	p.events = append(p.events, operation)
}
func (p *MultiCallProcessor) AddOpenEvent(openEvent *schemas.AccountOperation) {
	if len(p.events) > 0 {
		log.Info("Previous multicall events not processed", utils.ToJson(p.events))
	}
	if len(p.nonMultiCallEvents) > 0 {
		log.Fatal("There can't be non multicall events while multicall is running",
			utils.ToJson(p.nonMultiCallEvents))
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

func (p *MultiCallProcessor) GetTxHash() string {
	return p.txHash
}
