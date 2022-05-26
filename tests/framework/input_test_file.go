package framework

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/test"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/ds/dc_wrapper"
)

type TestMask struct {
	Mask    *core.BigInt `json:"mask"`
	Account string       `json:"account"`
}

type TestCall struct {
	Pools         []dc_wrapper.TestPoolCallData      `json:"pools"`
	CMs           []dc_wrapper.TestCMCallData        `json:"cms"`
	Accounts      []dc_wrapper.TestAccountCallData   `json:"accounts"`
	Masks         []TestMask                         `json:"masks"`
	ExecuteOnCM   map[string][]*ds.KnownCall         `json:"executeOnCM"`
	MainEventLogs map[string][]*ds.FuncWithMultiCall `json:"mainEventLogs"`
	OtherCalls    map[string][]string                `json:"others"`
	// txHash to transfers
	ExecuteTransfers map[string]core.Transfers `json:"executeTransfers"`
}

func (c *TestCall) Process() {
	return
}

type BlockInput struct {
	Events []test.TestEvent `json:"events"`
	Calls  TestCall         `json:"calls"`
}
type TestInput struct {
	Blocks    map[int64]BlockInput `json:"blocks"`
	MockFiles map[string]string    `json:"mocks"`
	States    test.TestState       `json:"states"`
}

func (testInput *TestInput) Get(file string, addressMap core.AddressMap, t *testing.T) *SyncAdapterMock {
	filePath := fmt.Sprintf("../inputs/%s", file)
	tmpObj := TestInput{}
	utils.ReadJsonAndSetInterface(filePath, &tmpObj)
	var syncAdapterObj *SyncAdapterMock

	for key, fileName := range tmpObj.MockFiles {
		mockFilePath := fmt.Sprintf("../inputs/%s", fileName)
		if key == "syncAdapters" {
			syncAdapterObj = &SyncAdapterMock{}
			addAddressSetJson(mockFilePath, syncAdapterObj, addressMap, t)
		}
	}
	addAddressSetJson(filePath, testInput, addressMap, t)
	return syncAdapterObj
}

func addAddressSetJson(filePath string, obj interface{}, addressMap core.AddressMap, t *testing.T) {
	var mock core.Json = utils.ReadJson(filePath)
	mock.ParseAddress(t, addressMap)
	// log.Info(utils.ToJson(mock))
	b, err := json.Marshal(mock)
	if err != nil {
		t.Error(err)
	}
	utils.SetJson(b, obj)
}
