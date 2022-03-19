package framework

import (
	"github.com/Gearbox-protocol/third-eye/core"
)

type MockExecuteParser struct {
	executeCalls map[string][]*core.KnownCall
	mainEventLogs map[string][]*core.FuncWithMultiCall
}

func (m *MockExecuteParser) setCalls(obj map[string][]*core.KnownCall) {
	if obj == nil {
		return
	}
	for txHash, calls := range obj {
		m.executeCalls[txHash] = calls
	}
}
func (m *MockExecuteParser) setMainEvents(obj map[string][]*core.FuncWithMultiCall) {
	if obj == nil {
		return
	}
	for txHash, mainEvents := range obj {
		m.mainEventLogs[txHash] = mainEvents
	}
}

func NewMockExecuteParser() *MockExecuteParser {
	return &MockExecuteParser{
		executeCalls: map[string][]*core.KnownCall{},
		mainEventLogs: map[string][]*core.FuncWithMultiCall{},
	}
}

func (m *MockExecuteParser) GetExecuteCalls(txHash, creditManagerAddr string, paramsList []core.ExecuteParams) []*core.KnownCall {
	return m.executeCalls[txHash]
}
func (m *MockExecuteParser) GetMainEventLogs(txHash, creditFacade string) []*core.FuncWithMultiCall {
	return m.mainEventLogs[txHash]
}
func (m *MockExecuteParser) GetTransfers(txHash string, owner []string) core.Transfers {
	return nil
}
