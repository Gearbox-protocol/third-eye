package tests

import (
	"github.com/Gearbox-protocol/third-eye/tests/framework"
	"testing"
)

func TestOpenCreditAccount(t *testing.T) {
	r, debtEng := framework.NewEngs(t, []string{"open_credit_account/input.json"})
	r.Eng.Sync(10)
	r.Check(r.Repo.GetBlocks()[3], "open_credit_account/blocks.json")
	debtEng.CalculateDebt()
	r.Check(debtEng.GetDebts(), "open_credit_account/debts.json")
}
