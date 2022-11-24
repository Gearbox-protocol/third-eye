package tests

import (
	"testing"

	"github.com/Gearbox-protocol/third-eye/tests/framework"
)

func TestMultipleFacadeInteractionSingleTx(t *testing.T) {
	r, debtEng := framework.NewEngs(t, []string{"mocks/initv2.json", "multiple_facade_interaction_single_tx/input.json"})
	r.Eng.Sync(10)
	outputBlocks := r.Repo.GetBlocks()
	debtEng.CalculateDebt()
	debtEng.CalCurrentDebts(10)
	debts := debtEng.GetDebts()
	delete(outputBlocks, 2)
	delete(outputBlocks, 3)
	r.Check(outputBlocks, "multiple_facade_interaction_single_tx/blocks.json")
	r.Check(debts, "multiple_facade_interaction_single_tx/debts.json")
}
