package repository

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/utils"
)

func getTestAdapter(name string, lastSync int64, details core.Json) *core.SyncAdapter {
	return &core.SyncAdapter{
		LastSync: lastSync,
		Contract: &core.Contract{
			ContractName: name,
			Address:      utils.RandomAddr(),
			DiscoveredAt: lastSync,
			FirstLogAt:   lastSync + 1,
		},
		Details: details,
	}
}

type call struct {
	Func string   `json:"func"`
	Data []string `json:"data"`
	Args []string `json:"args"`
}
type event struct {
	Address string   `json:"address"`
	Data    []string `json:"data"`
	Topics  []string `json:"topics"`
}
type BlockInput struct {
	Events []event `json:"events"`
	Calls  []call  `json:"calls"`
}
type TestInput struct {
	Blocks    map[int64]BlockInput `json:"blocks"`
	MockFiles map[string]string    `json:"mocks"`
}

type SyncAdapterMock struct {
	Data []*core.SyncAdapter `json:"adapters"`
}

type MockRepo struct {
	file         string
	repo         core.RepositoryI
	client       ethclient.ClientI
	InputFile    *TestInput
	AddressMap   core.AddressMap
	SyncAdapters []*core.SyncAdapter
	t            *testing.T
}

func (m *MockRepo) init() {
	m.handleMocks()
}

func (m *MockRepo) handleMocks() {
	m.InputFile = &TestInput{}
	m.AddressMap = core.AddressMap{}
	filePath := fmt.Sprintf("../tests/%s", m.file)
	//
	tmpObj := TestInput{}
	utils.ReadJsonAndSetInterface(filePath, &tmpObj)
	for key, fileName := range tmpObj.MockFiles {
		mockFilePath := fmt.Sprintf("../tests/%s", fileName)
		if key == "syncAdapters" {
			obj := &SyncAdapterMock{}
			m.addAddressSetJson(mockFilePath, obj)
			for _, adapter := range obj.Data {
				if adapter.DiscoveredAt == 0 {
					adapter.DiscoveredAt = adapter.LastSync
					adapter.FirstLogAt = adapter.LastSync + 1
				}
			}
			m.SyncAdapters = obj.Data
		}
	}
	//
	m.addAddressSetJson(filePath, m.InputFile)
	m.t.Fatal(utils.ToJson(m))
}

func (m *MockRepo) addAddressSetJson(filePath string, obj interface{}) {
	var mock core.Json = utils.ReadJson(filePath)
	mock.ParseAddress(m.t, m.AddressMap)
	b, err := json.Marshal(mock)
	if err != nil {
		m.t.Error(err)
	}
	utils.SetJson(b, obj)
}

func TestRepo(t *testing.T) {
	client := &ethclient.TestClient{}
	repo := newRepository(nil, client, &config.Config{}, nil)
	r := MockRepo{
		repo:   repo,
		client: client,
		file:   "test1.json",
		t:      t,
	}
	r.init()
}
