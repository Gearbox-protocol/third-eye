package tests

import (
	"encoding/json"
	"fmt"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/utils"
	"math/big"
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
	Adapters  []*core.SyncAdapter        `json:"adapters"`
	CMState   []*core.CreditManagerState `json:"cmState"`
	PoolState []*core.PoolState          `json:"poolState"`
	Tokens    []*core.Token              `json:"tokens"`
}

type MockRepo struct {
	file         string
	repo         core.RepositoryI
	client       *TestClient
	InputFile    *TestInput
	AddressMap   core.AddressMap
	SyncAdapters []*core.SyncAdapter
	t            *testing.T
	eng          core.EngineI
	//oracle to token
	feedToToken   map[string]string
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
	filePath := fmt.Sprintf("../inputs/%s", m.file)
	//
	tmpObj := TestInput{}
	utils.ReadJsonAndSetInterface(filePath, &tmpObj)
	for key, fileName := range tmpObj.MockFiles {
		mockFilePath := fmt.Sprintf("../inputs/%s", fileName)
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
	for _, adapter := range obj.Adapters {
		if adapter.DiscoveredAt == 0 {
			adapter.DiscoveredAt = adapter.LastSync
			adapter.FirstLogAt = adapter.LastSync + 1
		}
		actualAdapter := m.repo.PrepareSyncAdapter(adapter)
		switch actualAdapter.GetName() {
		case core.ChainlinkPriceFeed:
			oracle := actualAdapter.GetDetails("oracle")
			token := actualAdapter.GetDetails("token")
			m.repo.AddTokenOracle(token, oracle, actualAdapter.GetAddress(), actualAdapter.GetDiscoveredAt())
			m.feedToToken[actualAdapter.GetAddress()] = token
		case core.CreditManager:
			for _, state := range obj.CMState {
				if state.Address == actualAdapter.GetAddress() {
					state.Sessions = map[string]string{}
					actualAdapter.SetUnderlyingState(state)
				}
			}
		case core.Pool:
			for _, state := range obj.PoolState {
				if state.Address == actualAdapter.GetAddress() {
					actualAdapter.SetUnderlyingState(state)
				}
			}
		}
		kit.Add(actualAdapter)
	}
	for _, tokenObj := range obj.Tokens {
		switch tokenObj.Symbol {
		case "USDC":
			m.client.SetUSDC(tokenObj.Address)
		case "WETH":
			m.client.SetWETH(tokenObj.Address)
		}
		m.repo.AddTokenObj(tokenObj)
	}
	m.SyncAdapters = obj.Adapters
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
	// log.Info(utils.ToJson(mock))
	b, err := json.Marshal(mock)
	if err != nil {
		m.t.Error(err)
	}
	utils.SetJson(b, obj)
}

func (m *MockRepo) ProcessEvents() {
	events := map[int64]map[string][]types.Log{}
	prices := map[int64]map[string]*big.Int{}
	for blockNum, block := range m.InputFile.Blocks {
		if events[blockNum] == nil {
			events[blockNum] = make(map[string][]types.Log)
		}
		for ind, event := range block.Events {
			txLog := event.Process(m.addressToType[event.Address])
			txLog.Index = uint(ind)
			txLog.BlockNumber = uint64(blockNum)
			events[blockNum][event.Address] = append(events[blockNum][event.Address], txLog)
			if event.Topics[0] == core.Topic("AnswerUpdated(int256,uint256,uint256)").Hex() {
				price, ok := new(big.Int).SetString(txLog.Topics[2].Hex()[2:], 16)
				if !ok {
					log.Fatal("Failed in parsing price in answerupdated")
				}
				if prices[blockNum] == nil {
					prices[blockNum] = make(map[string]*big.Int)
				}
				token := m.feedToToken[txLog.Address.Hex()]
				prices[blockNum][token] = price
			}
		}
	}
	m.client.SetEvents(events)
	// log.Info(utils.ToJson(prices))
	m.client.SetPrices(prices)
}
func (m *MockRepo) ProcessCalls() {
	accountMask := make(map[int64]map[string]*big.Int)
	wrapper := m.repo.GetDCWrapper()
	for blockNum, block := range m.InputFile.Blocks {
		calls := core.NewDCCalls()
		for _, poolCall := range block.Calls.Pools {
			calls.Pools[poolCall.Addr] = poolCall
		}
		for _, accountCall := range block.Calls.Accounts {
			key := fmt.Sprintf("%s_%s", accountCall.CreditManager, accountCall.Borrower)
			calls.Accounts[key] = accountCall
		}
		for _, cmCall := range block.Calls.CMs {
			calls.CMs[cmCall.Addr] = cmCall
		}
		for _, maskDetails := range block.Calls.Masks {
			if accountMask[blockNum] == nil {
				accountMask[blockNum] = make(map[string]*big.Int)
			}
			accountMask[blockNum][maskDetails.Account] = maskDetails.Mask.Convert()

		}
		wrapper.SetCalls(blockNum, calls)
	}
	m.client.SetMasks(accountMask)
}
