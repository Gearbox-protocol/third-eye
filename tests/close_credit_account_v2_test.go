package tests

import (
	"testing"

	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/third-eye/tests/framework"
)

func TestCloseCreditAccountV2(t *testing.T) {
	r, debtEng := framework.NewEngs(t,
		[]string{"mocks/initv2.json", "account_lifecycle_v2/input.json", "close_credit_account_v2/input.json"})
	r.Eng.Sync(10)
	blocks := filterBlocks(r.Repo.GetBlocks(), []int64{8, 9})
	r.Check(blocks, "close_credit_account_v2/blocks.json")
	debtEng.CalculateDebt()
	debts := debtEng.GetDebts()
	filterDebts(debts, t, 4)
	r.Check(debts, "close_credit_account_v2/debts.json")
	for _, session := range r.Repo.GetSessions() {
		r.Check(session, "close_credit_account_v2/sessions.json")
	}
}

func filterBlocks(blocks map[int64]*schemas.Block, filter []int64) map[int64]*schemas.Block {
	newBlocks := map[int64]*schemas.Block{}
	for _, blockNum := range filter {
		newBlocks[blockNum] = blocks[blockNum]
	}
	return newBlocks
}
