package framework

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/test"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/ds/dc_wrapper"
	"github.com/stretchr/testify/require"

	"github.com/ethereum/go-ethereum/core/types"
)

func getTestAdapter(name string, lastSync int64, details core.Json) *ds.SyncAdapter {
	return &ds.SyncAdapter{
		SyncAdapterSchema: &schemas.SyncAdapterSchema{
			LastSync: lastSync,
			Contract: &schemas.Contract{
				ContractName: name,
				Address:      utils.RandomAddr(),
				DiscoveredAt: lastSync,
				FirstLogAt:   lastSync + 1,
			},
			Details: details,
		},
	}
}

type SyncAdapterMock struct {
	Adapters  []*ds.SyncAdapter             `json:"adapters"`
	CMState   []*schemas.CreditManagerState `json:"cmState"`
	PoolState []*schemas.PoolState          `json:"poolState"`
	Tokens    []*schemas.Token              `json:"tokens"`
}

type MockRepo struct {
	Repo         ds.RepositoryI
	client       *test.TestClient
	InputFile    *TestInput
	AddressMap   core.AddressMap
	SyncAdapters []*ds.SyncAdapter
	t            *testing.T
	Eng          ds.EngineI
	//oracle to token
	feedToToken   map[string]string
	addressToType map[string]string
	executeParser *MockExecuteParser
}

func NewMockRepo(repo ds.RepositoryI, client *test.TestClient,
	t *testing.T, eng ds.EngineI, ep *MockExecuteParser) MockRepo {
	return MockRepo{
		Repo:          repo,
		client:        client,
		t:             t,
		Eng:           eng,
		addressToType: make(map[string]string),
		feedToToken:   make(map[string]string),
		executeParser: ep,
	}
}

func (m *MockRepo) Init(files []string) {
	m.AddressMap = core.AddressMap{}
	for _, file := range files {
		m.fetchInputTestFile(file)
	}
}

func (m *MockRepo) fetchInputTestFile(inputFile string) *TestInput {
	testInput := &TestInput{}
	syncAdapterObj := testInput.Get(inputFile, m.AddressMap, m.t)
	// map address to type
	for key, value := range m.AddressMap {
		splits := strings.Split(key, "_")
		if len(splits) == 2 {
			m.addressToType[value] = splits[0]
		} else {
			m.t.Fatalf("Not properly formatted key: %s", key)
		}
	}
	// set feed to token
	if syncAdapterObj != nil {
		for _, adapter := range syncAdapterObj.Adapters {
			if adapter.GetName() == ds.ChainlinkPriceFeed {
				m.feedToToken[adapter.GetAddress()] = adapter.GetDetailsByKey("token")
			}
		}
	}

	m.ProcessState(testInput)
	m.ProcessCalls(testInput)
	m.ProcessEvents(testInput)
	m.setSyncAdapters(syncAdapterObj)
	return testInput

}

func (m *MockRepo) setSyncAdapters(obj *SyncAdapterMock) {
	if obj == nil {
		return
	}
	kit := m.Repo.GetKit()
	for _, adapter := range obj.Adapters {
		if adapter.DiscoveredAt == 0 {
			adapter.DiscoveredAt = adapter.LastSync
			adapter.FirstLogAt = adapter.LastSync + 1
		}
		actualAdapter := m.Repo.PrepareSyncAdapter(adapter)
		switch actualAdapter.GetName() {
		case ds.ChainlinkPriceFeed:
			oracle := actualAdapter.GetDetailsByKey("oracle")
			token := actualAdapter.GetDetailsByKey("token")
			m.Repo.AddTokenOracle(&schemas.TokenOracle{
				Token:       token,
				Oracle:      oracle,
				Feed:        actualAdapter.GetAddress(),
				BlockNumber: actualAdapter.GetDiscoveredAt(),
				Version:     actualAdapter.GetVersion()})
		case ds.CreditManager:
			for _, state := range obj.CMState {
				if state.Address == actualAdapter.GetAddress() {
					state.Sessions = map[string]string{}
					actualAdapter.SetUnderlyingState(state)
				}
			}
		case ds.Pool:
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
		m.Repo.AddTokenObj(tokenObj)
		m.client.AddToken(tokenObj.Address, tokenObj.Decimals)
	}
	// m.SyncAdapters = obj.Adapters
}

func (m *MockRepo) ProcessEvents(inputFile *TestInput) {
	events := map[int64]map[string][]types.Log{}
	prices := map[string]map[int64]*big.Int{}
	for blockNum, block := range inputFile.Blocks {
		if events[blockNum] == nil {
			events[blockNum] = make(map[string][]types.Log)
		}
		for ind, event := range block.Events {
			txLog := event.Process(m.addressToType[event.Address])
			txLog.Index = uint(ind)
			txLog.BlockNumber = uint64(blockNum)
			events[blockNum][event.Address] = append(events[blockNum][event.Address], txLog)
			if event.Topics[0] == core.Topic("AnswerUpdated(int256,uint256,uint256)").Hex() {
				price, ok := new(big.Int).SetString(txLog.Topics[1].Hex()[2:], 16)
				if !ok {
					log.Fatal("Failed in parsing price in answerupdated")
				}
				token := m.feedToToken[txLog.Address.Hex()]
				if prices[token] == nil {
					prices[token] = make(map[int64]*big.Int)
				}
				prices[token][blockNum] = price
			}
		}
	}
	m.client.SetEvents(events)
	// log.Info(utils.ToJson(prices))
	m.client.SetPrices(prices)
}
func (m *MockRepo) ProcessCalls(inputFile *TestInput) {
	accountMask := make(map[int64]map[string]*big.Int)
	wrapper := m.Repo.GetDCWrapper()
	otherCalls := make(map[int64]map[string][]string)
	for blockNum, block := range inputFile.Blocks {
		otherCalls[blockNum] = block.Calls.OtherCalls
		calls := dc_wrapper.NewDCCalls()
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
		m.executeParser.setCalls(block.Calls.ExecuteOnCM)
		m.executeParser.setMainEvents(block.Calls.MainEventLogs)
		m.executeParser.setTransfers(block.Calls.ExecuteTransfers)
		for _, maskDetails := range block.Calls.Masks {
			if accountMask[blockNum] == nil {
				accountMask[blockNum] = make(map[string]*big.Int)
			}
			accountMask[blockNum][maskDetails.Account] = maskDetails.Mask.Convert()

		}
		wrapper.SetCalls(blockNum, calls)
	}
	m.client.SetOtherCalls(otherCalls)
	m.client.SetMasks(accountMask)
}

func (m *MockRepo) ProcessState(inputFile *TestInput) {
	for _, oracle := range inputFile.States.Oracles {
		m.client.SetOracleState(oracle)
	}
}

// for matching state with the expected output
func (m *MockRepo) replaceWithVariable(obj interface{}) core.Json {
	bytes, err := json.Marshal(obj)
	log.CheckFatal(err)
	addrToVariable := core.AddressMap{}
	// TODO: FIX FOR HASH
	for variable, addr := range m.AddressMap {
		addrToVariable[addr] = "#" + variable
	}
	outputJson := core.Json{}
	err = json.Unmarshal(bytes, &outputJson)
	log.CheckFatal(err)
	outputJson.ReplaceWithVariable(addrToVariable)
	return outputJson
}

func (m *MockRepo) Check(value interface{}, fileName string) {
	outputJson := m.replaceWithVariable(value)
	fileName = fmt.Sprintf("../inputs/%s", fileName)
	require.JSONEq(m.t, string(utils.ReadFile(fileName)), utils.ToJson(outputJson))
}

func (m *MockRepo) Print(value interface{}) {
	outputJson := m.replaceWithVariable(value)
	m.t.Fatal(utils.ToJson(outputJson))
}
