package tests

import (
	"testing"

	"github.com/Gearbox-protocol/third-eye/tests/framework"
)

func TestDAOOperations(t *testing.T) {
	r, _ := framework.NewEngs(t, []string{"dao_operations/input.json"})
	r.Eng.Sync(10)
	r.Check(map[string]interface{}{"data": r.Repo.GetBlocks()[3].DAOOperations},
		"dao_operations/blocks.json")
}
