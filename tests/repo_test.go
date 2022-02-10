package tests

import (
	"encoding/json"
	"fmt"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/debts"
	"github.com/Gearbox-protocol/third-eye/engine"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/repository"
	"github.com/Gearbox-protocol/third-eye/utils"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
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

type SyncAdapterMock struct {
	Data      []*core.SyncAdapter        `json:"adapters"`
	CMState   []*core.CreditManagerState `json:"cmState"`
	PoolState []*core.PoolState          `json:"poolState"`
	Tokens    []*core.Token              `json:"tokens"`
}

type MockRepo struct {
	file          string
	repo          core.RepositoryI
	client        *ethclient.TestClient
	InputFile     *TestInput
	AddressMap    core.AddressMap
	SyncAdapters  []*core.SyncAdapter
	t             *testing.T
	eng           core.EngineI
	addressToType map[string]string
}

func (m *MockRepo) init() {
	m.handleMocks()
	m.ProcessEvents()
	m.ProcessCalls()
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
			m.setSyncAdapters(mockFilePath)
		}
	}
	//
	m.addAddressSetJson(filePath, m.InputFile)

}

func (m *MockRepo) setSyncAdapters(mockFilePath string) {
	obj := &SyncAdapterMock{}
	kit := m.repo.GetKit()
	m.addAddressSetJson(mockFilePath, obj)
	for _, adapter := range obj.Data {
		if adapter.DiscoveredAt == 0 {
			adapter.DiscoveredAt = adapter.LastSync
			adapter.FirstLogAt = adapter.LastSync + 1
		}
		switch adapter.GetName() {
		case core.ChainlinkPriceFeed:
			oracle := adapter.GetDetails("oracle")
			token := adapter.GetDetails("token")
			m.repo.AddTokenOracle(token, oracle, adapter.GetAddress(), adapter.DiscoveredAt)
		case core.CreditManager:
			for _, state := range obj.CMState {
				if state.Address == adapter.GetAddress() {
					adapter.SetUnderlyingState(state)
				}
			}
		case core.Pool:
			for _, state := range obj.PoolState {
				if state.Address == adapter.GetAddress() {
					adapter.SetUnderlyingState(state)
				}
			}
		}
		kit.Add(m.repo.PrepareSyncAdapter(adapter))
	}
	for _, tokenObj := range obj.Tokens {
		m.repo.AddTokenObj(tokenObj)
	}
	m.SyncAdapters = obj.Data
	for key, value := range m.AddressMap {
		splits := strings.Split(key, "_")
		if len(splits) == 2 {
			m.addressToType[value] = splits[0]
		} else {
			m.t.Fatalf("Not properly formatted key: %s", key)
		}
	}
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
	client := ethclient.NewTestClient()
	cfg := &config.Config{}
	repo := repository.GetRepository(nil, client, cfg, nil)
	debtEng := debts.NewDebtEngine(nil, client, cfg, repo)
	eng := engine.NewEngine(cfg, client, debtEng, repo)
	r := MockRepo{
		repo:   repo,
		client: client,
		file:   "test1.json",
		t:      t,
		eng:    eng,
	}
	r.init()
	eng.Sync(10)
	debtEng.CalculateDebt()
}

func (m *MockRepo) ProcessEvents() {
	events := map[int64]map[string][]types.Log{}
	for blockNum, block := range m.InputFile.Blocks {
		if events[blockNum] == nil {
			events[blockNum] = make(map[string][]types.Log)
		}
		for ind, event := range block.Events {
			txLog := event.Process(m.addressToType[event.Address])
			txLog.Index = uint(ind)
			txLog.BlockNumber = uint64(blockNum)
			events[blockNum][event.Address] = append(events[blockNum][event.Address], txLog)
		}
	}
	m.client.SetEvents(events)
}
func (m *MockRepo) ProcessCalls() {
	wrapper := m.repo.GetDCWrapper()
	for blockNum, block := range m.InputFile.Blocks {
		calls := core.NewDCCalls()
		for _, poolCall := range block.Calls.Pools {
			calls.Pools[poolCall.Addr.Hex()] = poolCall
		}
		for _, accountCall := range block.Calls.Accounts {
			key := fmt.Sprintf("%s_%s", accountCall.CreditManager, accountCall.Borrower)
			calls.Accounts[key] = accountCall
		}
		for _, cmCall := range block.Calls.CMs {
			calls.CMs[cmCall.Addr.Hex()] = cmCall
		}
		wrapper.SetCalls(blockNum, calls)
	}
}
