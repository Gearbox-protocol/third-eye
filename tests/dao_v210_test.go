package tests

import (
	"fmt"
	"testing"

	"github.com/Gearbox-protocol/third-eye/tests/framework"
)

func TestDAOV210Operations(t *testing.T) {
	r, _ := framework.NewEngs(t, []string{"dao_operations_v210/input.json"})
	r.Eng.Sync(10)
	fmt.Println(r.Repo.GetBlocks())
	data := map[string]interface{}{
		"3": map[string]interface{}{
			"daoOperations": r.Repo.GetBlocks()[3].DAOOperations,
		},
		"4": map[string]interface{}{
			"daoOperations": r.Repo.GetBlocks()[4].DAOOperations,
		},
	}
	r.Check(data, "dao_operations_v210/blocks.json")
}
