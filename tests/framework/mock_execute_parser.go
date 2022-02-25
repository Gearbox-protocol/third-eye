package framework

import "github.com/Gearbox-protocol/third-eye/core"

type MockExecuteParser struct {
	executeCalls map[string][]*core.KnownCall
}

func (m *MockExecuteParser) GetExecuteCalls(txHash, creditManagerAddr string, paramsList []core.ExecuteParams) []*core.KnownCall {
	return m.executeCalls[txHash]
}
func (m *MockExecuteParser) setCalls(obj map[string][]*core.KnownCall) {
	if obj == nil {
		return
	}
	for txHash, calls := range obj {
		m.executeCalls[txHash] = calls
	}
}

func NewMockExecuteParser() *MockExecuteParser {
	return &MockExecuteParser{
		executeCalls: map[string][]*core.KnownCall{},
	}
}
