package tests

import (
	"testing"

	"github.com/Gearbox-protocol/third-eye/tests/framework"
)

func TestPriceOracle(t *testing.T) {
	r, _ := framework.NewEngs(t, []string{"oracle/input.json"})
	r.Eng.Sync(10)
	r.Eng.Sync(10) // bcz one of the chainlink is created so should be added again
	blocks := r.Repo.GetBlocks()
	delete(blocks, 2)
	r.Check(blocks, "oracle/blocks.json")
	r.Check(r.Repo.GetTokenOracles(), "oracle/token_oracle.json")
}
