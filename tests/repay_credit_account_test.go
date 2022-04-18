package tests

import (
	"github.com/Gearbox-protocol/third-eye/tests/framework"
	"testing"
)

func TestRepayCreditAccount(t *testing.T) {
	r, debtEng := framework.NewEngs(t, []string{"account_lifecycle/input.json", "repay_credit_account/input.json"})
	r.Eng.Sync(10)
	r.Check(r.Repo.GetBlocks()[7], "repay_credit_account/blocks.json")
	debtEng.CalculateDebt()
	debtsAndCurrentDebts := debtEng.GetDebts()
	filterDebts(debtsAndCurrentDebts, t)
	r.Check(debtsAndCurrentDebts, "repay_credit_account/debts.json")
}
