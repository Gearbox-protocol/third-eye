package tests

import (
	"testing"

	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/debts"
	"github.com/Gearbox-protocol/third-eye/engine"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/repository"
	"github.com/Gearbox-protocol/third-eye/tests/framework"
	"github.com/Gearbox-protocol/third-eye/utils"
)

func TestLiquidateCreditAccount(t *testing.T) {
	log.SetTestLogging(t)
	client := framework.NewTestClient()
	cfg := &config.Config{}
	ep := framework.NewMockExecuteParser()
	repo := repository.GetRepository(nil, client, cfg, ep)
	debtEng := debts.GetDebtEngine(nil, client, cfg, repo, true)
	eng := engine.NewEngine(cfg, client, debtEng, repo)
	r := framework.NewMockRepo(repo, client, t, eng, ep)
	r.Init([]string{"account_lifecycle/input.json", "liquidate_credit_account/input.json"})
	log.Info(utils.ToJson(r.AddressMap))
	eng.Sync(10)
	data := map[string]interface{}{}
	blocks := repo.GetBlocks()
	data["7"] = blocks[7]
	data["8"] = blocks[8]
	r.Check(data, "liquidate_credit_account/blocks.json")
	debtEng.CalculateDebt()
	debtsAndCurrentDebts := debtEng.GetDebts()
	filterDebts(debtsAndCurrentDebts, t, 4)
	r.Check(debtsAndCurrentDebts, "liquidate_credit_account/debts.json")
}

func filterDebts(debtsAndCurrentDebts core.Json, t *testing.T, indexes ...int) {
	debts := debtsAndCurrentDebts["debts"]
	parsedDebts, ok := debts.([]*core.Debt)
	if !ok {
		t.Errorf("parsing debts from engine failed: %s", utils.ToJson(debts))
	}
	filteredDebts := []*core.Debt{}
	for _, index := range indexes {
		filteredDebts = append(filteredDebts, parsedDebts[index])
	}
	debtsAndCurrentDebts["debts"] = filteredDebts
}
