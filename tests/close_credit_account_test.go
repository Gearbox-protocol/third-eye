package tests

import (
	"github.com/Gearbox-protocol/third-eye/tests/framework"
	"testing"
)

func TestCloseCreditAccount(t *testing.T) {
	r, debtEng := framework.NewEngs(t,
		[]string{"account_lifecycle/input.json", "close_credit_account/input.json"})
	r.Eng.Sync(10)
	blocks := r.Repo.GetBlocks()
	r.Check(map[int64]interface{}{7: blocks[7], 8: blocks[8]}, "close_credit_account/blocks.json")
	debtEng.CalculateDebt()
	debts := debtEng.GetDebts()
	filterDebts(debts, t, 4)
	r.Check(debts, "close_credit_account/debts.json")
}
