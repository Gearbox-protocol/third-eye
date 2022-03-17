package tests

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/tests/framework"
	"github.com/Gearbox-protocol/third-eye/utils"
	"testing"
)

func TestLiquidateCreditAccount(t *testing.T) {
	r, debtEng := framework.NewEngs(t, []string{"account_lifecycle/input.json", "liquidate_credit_account/input.json"})
	r.Eng.Sync(10)
	data := map[string]interface{}{}
	blocks := r.Repo.GetBlocks()
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
