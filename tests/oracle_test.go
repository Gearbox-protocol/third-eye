package tests

import (
	"github.com/Gearbox-protocol/third-eye/tests/framework"
	"testing"
)

func TestPriceOracle(t *testing.T) {
	r, _ := framework.NewEngs(t, []string{"oracle/input.json"})
	r.Eng.Sync(10)
	blocks := r.Repo.GetBlocks()
	delete(blocks, 2)
	r.Check(blocks, "oracle/blocks.json")
	r.Check(r.Repo.GetTokenOracles(), "oracle/token_oracle.json")
}
