package framework

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/test"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/stretchr/testify/require"

	"github.com/ethereum/go-ethereum/core/types"
)

type SyncAdapterMock struct {
	Adapters  []*ds.SyncAdapter             `json:"adapters"`
	CMState   []*schemas.CreditManagerState `json:"cmState"`
	PoolState []*schemas.PoolState          `json:"poolState"`
	Tokens    []*schemas.Token              `json:"tokens"`
}

type MockRepo struct {
	Repo         ds.RepositoryI
	client       *test.TestClient
	InputFile    *TestInput3Eye
	AddressMap   core.AddressMap
	SyncAdapters []*ds.SyncAdapter
	t            *testing.T
	Eng          ds.EngineI
	//oracle to token
	feedToToken   map[string]string
	executeParser *MockExecuteParser
}

func NewMockRepo(repo ds.RepositoryI, client *test.TestClient,
	t *testing.T, eng ds.EngineI, ep *MockExecuteParser) MockRepo {
	return MockRepo{
		Repo:          repo,
		client:        client,
		t:             t,
		Eng:           eng,
		feedToToken:   make(map[string]string),
		executeParser: ep,
	}
}

func (m *MockRepo) Init(fileNames []string) {
	inputFile, addressMap := ParseMockClientInput(m.t, m.client, fileNames)
	for _, token := range m.client.GetToken() {
		m.Repo.AddTokenObj(token)
	}
	m.AddressMap = addressMap
	m.processInputTestFile(inputFile)
}

func ParseMockClientInput(t *testing.T, client *test.TestClient, fileNames []string) (*TestInput3Eye, core.AddressMap) {
	filePaths := make([]string, len(fileNames))
	for i, fileName := range fileNames {
		filePaths[i] = fmt.Sprintf("../inputs/%s", fileName)
	}
	files := make([]test.TestInputI, len(fileNames))
	for i := range fileNames {
		files[i] = NewTestInput3Eye()
	}
	testInput3Eye, addressMap := test.LoadTestFiles(filePaths, files, t)
	//
	inputFile := testInput3Eye.(*TestInput3Eye)
	inputFile.AddToClient(client, addressMap)
	return inputFile, addressMap
}

func (m *MockRepo) processInputTestFile(inputFile *TestInput3Eye) {
	syncAdapterObj := inputFile.GetSyncAdapter(m.AddressMap, m.t)
	// set feed to token
	if syncAdapterObj != nil {
		for _, adapter := range syncAdapterObj.Adapters {
			if adapter.GetName() == ds.ChainlinkPriceFeed {
				m.feedToToken[adapter.GetAddress()] = adapter.GetDetailsByKey("token")
			}
		}
	}
	for _, executeLogs := range inputFile.ExecuteParser {
		m.executeParser.setCalls(executeLogs.ExecuteOnCM)
		m.executeParser.setMainEvents(executeLogs.MainEventLogs)
		m.executeParser.setTransfers(executeLogs.ExecuteTransfers)
	}

	m.processDCCalls(inputFile)
	m.processPrices(inputFile)
	m.setSyncAdapters(syncAdapterObj)
}

func (m *MockRepo) setSyncAdapters(obj *SyncAdapterMock) {
	if obj == nil {
		return
	}
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
			m.Repo.DirectlyAddTokenOracle(&schemas.TokenOracle{
				Token:       token,
				Oracle:      oracle,
				Feed:        actualAdapter.GetAddress(),
				FeedType:    ds.ChainlinkPriceFeed,
				BlockNumber: actualAdapter.GetDiscoveredAt(),
				Version:     actualAdapter.GetVersion(),
			})
		case ds.CreditManager:
			for _, state := range obj.CMState {
				if state.Address == actualAdapter.GetAddress() {
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
		m.Repo.AddSyncAdapter(actualAdapter)
	}
}

func (m *MockRepo) processPrices(inputFile *TestInput3Eye) {
	events := map[int64]map[string][]types.Log{}
	prices := map[string]map[int64]*big.Int{}
	for blockNum, block := range inputFile.Blocks {
		if events[blockNum] == nil {
			events[blockNum] = make(map[string][]types.Log)
		}
		for _, event := range block.Events {
			if event.Topics[0] == core.Topic("AnswerUpdated(int256,uint256,uint256)").Hex() {
				price, ok := new(big.Int).SetString(event.Topics[1][7:], 10)
				if !ok {
					log.Fatal("Failed in parsing price in answerupdated")
				}
				token := m.feedToToken[event.Address]
				if prices[token] == nil {
					prices[token] = make(map[int64]*big.Int)
				}
				prices[token][blockNum] = price
			}
		}
	}
	m.client.SetPrices(prices)
}

func (m *MockRepo) processDCCalls(inputFile *TestInput3Eye) {
	wrapper := m.Repo.GetDCWrapper()
	for blockNum, block := range inputFile.Blocks {
		calls := test.NewDCCalls()
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
		wrapper.SetCalls(blockNum, calls)
	}
}

func (m *MockRepo) Check(value interface{}, fileName string) {
	outputJson := test.ReplaceWithVariable(value, m.AddressMap)
	fileName = fmt.Sprintf("../inputs/%s", fileName)
	require.JSONEq(m.t, string(utils.ReadFile(fileName)), utils.ToJson(outputJson))
}
func Check(t *testing.T, addressMap map[string]string, value interface{}, fileName string) {
	outputJson := test.ReplaceWithVariable(value, addressMap)
	fileName = fmt.Sprintf("../inputs/%s", fileName)
	require.JSONEq(t, string(utils.ReadFile(fileName)), utils.ToJson(outputJson))
}

func (m *MockRepo) Print(value interface{}) {
	Print(m.t, m.AddressMap, value)
}
func Print(t *testing.T, addressMap map[string]string, value interface{}) {
	outputJson := test.ReplaceWithVariable(value, addressMap)
	t.Fatal(utils.ToJson(outputJson))
}
