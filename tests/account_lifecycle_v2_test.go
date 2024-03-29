package tests

import (
	"testing"

	"github.com/Gearbox-protocol/third-eye/tests/framework"
)

func TestLifecycleCreditAccountV2(t *testing.T) {
	r, debtEng := framework.NewEngs(t, []string{"mocks/initv2.json", "account_lifecycle_v2/input.json"})
	r.Eng.Sync(10)
	outputBlocks := r.Repo.GetBlocks()
	debtEng.CalculateDebt()
	debts := debtEng.GetDebts()
	delete(outputBlocks, 2)
	delete(outputBlocks, 3)
	r.Check(outputBlocks, "account_lifecycle_v2/blocks.json")
	r.Check(debts, "account_lifecycle_v2/debts.json")
}
