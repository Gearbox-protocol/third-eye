package credit_manager

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/utils"
)

type MultiCallProcessor struct {
	borrower            string
	txHash              string
	OpenEvent           *core.AccountOperation
	events              []*core.AccountOperation
	nonMultiCallEvents  []*core.AccountOperation
	MultiCallStartEvent *core.AccountOperation
	running             bool
}

func (p *MultiCallProcessor) AddMulticallEvent(operation *core.AccountOperation) {
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
func (p *MultiCallProcessor) AddOpenEvent(openEvent *core.AccountOperation) {
	if len(p.events) > 0 {
		log.Info("Previous multicall events not processed", utils.ToJson(p.events))
	}
	if len(p.nonMultiCallEvents) > 0 {
		log.Fatal("There can't be non multicall events while multicall is running",
			utils.ToJson(p.nonMultiCallEvents))
	}
	p.OpenEvent = openEvent
}
func (p *MultiCallProcessor) Start(borrower, txHash string, startEvent *core.AccountOperation) {
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
func (p *MultiCallProcessor) PopMulticallEventsV2() []*core.AccountOperation {
	collaterals := p.events
	p.events = []*core.AccountOperation{}
	return collaterals
}

func (p *MultiCallProcessor) popNonMulticallEventsV2() []*core.AccountOperation {
	calls := p.nonMultiCallEvents
	p.nonMultiCallEvents = []*core.AccountOperation{}
	return calls
}

func (p *MultiCallProcessor) GetTxHash() string {
	return p.txHash
}
