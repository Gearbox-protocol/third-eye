package tests

import (
	"github.com/Gearbox-protocol/third-eye/tests/framework"
	"testing"
)

func TestLifecycleCreditAccount(t *testing.T) {
	r, debtEng := framework.NewEngs(t, []string{"account_lifecycle/input.json"})
	r.Eng.Sync(10)

	outputBlocks := r.Repo.GetBlocks()
	delete(outputBlocks, 2)
	r.Check(outputBlocks, "account_lifecycle/blocks.json")
	debtEng.CalculateDebt()
	r.Check(debtEng.GetDebts(), "account_lifecycle/debts.json")
}
