package framework

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/test"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
)

type TestExecuteParser struct {
	ExecuteOnCM   map[string][]*ds.KnownCall               `json:"executeOnCM"`
	MainEventLogs map[string][]*ds.MainactionWithMulticall `json:"mainEventLogs"`
	// txHash to transfers
	ExecuteTransfers map[string]core.Transfers `json:"executeTransfers"`
}

type TestInput3Eye struct {
	*test.TestInput
	ExecuteParser map[int64]*TestExecuteParser `json:"executeParser"`
}

func NewTestInput3Eye() *TestInput3Eye {
	return &TestInput3Eye{
		TestInput:     test.NewTestInput(),
		ExecuteParser: make(map[int64]*TestExecuteParser),
	}
}

func (file *TestInput3Eye) Merge(other test.TestInputI) {
	otherFile := other.(*TestInput3Eye)
	file.TestInput.Merge(otherFile.TestInput)
	for block, executeLogs := range otherFile.ExecuteParser {
		if file.ExecuteParser[block] != nil {
			log.Fatal("execute parser already present for block", block)
		}
		file.ExecuteParser[block] = executeLogs
	}
}
func (testInput *TestInput3Eye) GetSyncAdapter(addressMap core.AddressMap, t *testing.T) *SyncAdapterMock {
	for key, fileName := range testInput.MockFiles {
		mockFilePath := fmt.Sprintf("../inputs/%s", fileName)
		if key == "syncAdapters" {
			syncAdapterObj := &SyncAdapterMock{}
			addAddressSetJson(mockFilePath, syncAdapterObj, addressMap, t)
			return syncAdapterObj
		}
	}
	return nil
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
