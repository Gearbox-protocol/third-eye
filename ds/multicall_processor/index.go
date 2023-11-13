package multicall_processor

import "github.com/Gearbox-protocol/sdk-go/core/schemas"

const (
	// v2
	NotSet = iota
	GBFacadeOpenEvent
	GBFacadeMulticallEvent
	GBFacadeCloseEvent // v2 for close/liquidate and for v3 only close
	// v3
	GBFacadev3LiqUpdateEvent
)

// facade Action
type FacadeAccountAction struct {
	Data      *schemas.AccountOperation
	Type      int
	withoutMC bool
	// end if MulticallFinished is emitted / addcollateral is emitted after openCreditAccountWithoutMulticall
	ended bool
}

type (
	v FacadeAccountAction
)

func (v FacadeAccountAction) IsOpen() bool {
	return v.Type == GBFacadeOpenEvent
}
func (v FacadeAccountAction) LenofMulticalls() int {
	return len(v.Data.MultiCall)
}
func (v FacadeAccountAction) GetMulticallsFromEvent() []*schemas.AccountOperation {
	return v.Data.MultiCall
}
func (v *FacadeAccountAction) SetMulticalls(ops []*schemas.AccountOperation) {
	v.Data.MultiCall = ops
}

type MulticallProcessorI interface {
	AddMulticallEvent(operation *schemas.AccountOperation)
	Start(txHash string, startEvent *schemas.AccountOperation)
	AddOpenEvent(openEvent *schemas.AccountOperation)
	AddCloseEvent(event *schemas.AccountOperation)
	AddLiquidateEvent(event *schemas.AccountOperation)
	PopNonMulticallEvents() []*schemas.AccountOperation
	PopMainActions() (facadeActions, openEventWithoutMulticall []*FacadeAccountAction)
	End()
}
