package tests

import (
	"encoding/json"
	"testing"

	"fmt"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/debts"
	"github.com/Gearbox-protocol/third-eye/engine"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/repository"
	"github.com/Gearbox-protocol/third-eye/utils"
	"github.com/stretchr/testify/require"
)

func Test1(t *testing.T) {
	client := NewTestClient()
	cfg := &config.Config{}
	repo := repository.GetRepository(nil, client, cfg, nil)
	debtEng := debts.NewDebtEngine(nil, client, cfg, repo)
	eng := engine.NewEngine(cfg, client, debtEng, repo)
	r := MockRepo{
		repo:          repo,
		client:        client,
		file:          "test1_input.json",
		t:             t,
		eng:           eng,
		addressToType: make(map[string]string),
		feedToToken:   make(map[string]string),
	}
	r.init()
	log.Info(utils.ToJson(r.AddressMap))
	eng.Sync(10)

	r.check(t, repo.GetBlocks()[3], "test1_blocks.json")
	debtEng.CalculateDebt()
	r.check(t, debtEng.GetDebts(), "test1_debts.json")
}

func (m *MockRepo) replaceWithVariable(obj interface{}) core.Json {
	bytes, err := json.Marshal(obj)
	log.CheckFatal(err)
	addrToVariable := core.AddressMap{}
	for variable, addr := range m.AddressMap {
		addrToVariable[addr] = "#" + variable
	}
	outputJson := core.Json{}
	err = json.Unmarshal(bytes, &outputJson)
	log.CheckFatal(err)
	outputJson.ReplaceWithVariable(addrToVariable)
	return outputJson
}

func (m *MockRepo) check(t *testing.T, value interface{}, fileName string) {
	outputJson := m.replaceWithVariable(value)
	fileName = fmt.Sprintf("../inputs/%s", fileName)
	require.JSONEq(t, string(utils.ReadFile(fileName)), utils.ToJson(outputJson))
}

func (m *MockRepo) print(t *testing.T, value interface{}) {
	outputJson := m.replaceWithVariable(value)
	log.Fatal(utils.ToJson(outputJson))
}
