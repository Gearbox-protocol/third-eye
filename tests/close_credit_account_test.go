package tests

import (
	"github.com/Gearbox-protocol/third-eye/tests/framework"
	"testing"
)

func TestCloseCreditAccount(t *testing.T) {
	r, _ := framework.NewEngs(t,
		[]string{"account_lifecycle/input.json", "close_credit_account/input.json"})
	r.Eng.Sync(10)
	r.Check(r.Repo.GetBlocks()[7], "close_credit_account/blocks.json")
}
