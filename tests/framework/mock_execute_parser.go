package framework

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/third-eye/ds"
)

type MockExecuteParser struct {
	executeCalls     map[string][]*ds.KnownCall
	mainEventLogs    map[string][]*ds.FacadeCallNameWithMulticall
	executeTransfers map[string]core.Transfers
}

func (m *MockExecuteParser) setCalls(obj map[string][]*ds.KnownCall) {
	if obj == nil {
		return
	}
	for txHash, calls := range obj {
		m.executeCalls[txHash] = calls
	}
}
func (m *MockExecuteParser) setMainEvents(obj map[string][]*ds.FacadeCallNameWithMulticall) {
	if obj == nil {
		return
	}
	for txHash, mainEvents := range obj {
		m.mainEventLogs[txHash] = mainEvents
	}
}

func (m *MockExecuteParser) setTransfers(obj map[string]core.Transfers) {
	if obj == nil {
		return
	}
	for txHash, transfers := range obj {
		m.executeTransfers[txHash] = transfers
	}
}

func NewMockExecuteParser() *MockExecuteParser {
	return &MockExecuteParser{
		executeCalls:     map[string][]*ds.KnownCall{},
		mainEventLogs:    map[string][]*ds.FacadeCallNameWithMulticall{},
		executeTransfers: map[string]core.Transfers{},
	}
}

func (m *MockExecuteParser) GetExecuteCalls(txHash, creditManagerAddr string, paramsList []ds.ExecuteParams) []*ds.KnownCall {
	return m.executeCalls[txHash]
}
func (m *MockExecuteParser) GetMainCalls(txHash, creditFacade string) (mainCalls []*ds.FacadeCallNameWithMulticall) {
	for _, entry := range m.mainEventLogs[txHash] {
		entry.Name = ds.FacadeAccountMethodSigToCallName(entry.Name)
		mainCalls = append(mainCalls, entry)
	}

	return mainCalls
}
func (m *MockExecuteParser) GetTransfersAtClosev2(txHash, account, underlyingToken string, users ds.BorrowerAndTo) core.Transfers {
	return m.executeTransfers[txHash]
}
