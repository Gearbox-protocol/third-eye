package tests

import (
	"testing"

	"github.com/Gearbox-protocol/third-eye/tests/framework"
)

func TestDAOV2Operations(t *testing.T) {
	r, _ := framework.NewEngs(t, []string{"dao_operations_v2/input.json"})
	r.Eng.Sync(10)

	data := map[string]interface{}{
		"3": map[string]interface{}{
			"daoOperations": r.Repo.GetBlocks()[3].DAOOperations,
			"allowedTokens": r.Repo.GetBlocks()[3].AllowedTokens,
		},
		"4": map[string]interface{}{
			"daoOperations": r.Repo.GetBlocks()[4].DAOOperations,
			"allowedTokens": r.Repo.GetBlocks()[4].AllowedTokens,
		},
		"5": map[string]interface{}{
			"daoOperations": r.Repo.GetBlocks()[5].DAOOperations,
			"allowedTokens": r.Repo.GetBlocks()[5].AllowedTokens,
		},
		"6": map[string]interface{}{
			"daoOperations": r.Repo.GetBlocks()[6].DAOOperations,
			"allowedTokens": r.Repo.GetBlocks()[6].AllowedTokens,
		},
	}
	r.Check(data, "dao_operations_v2/blocks.json")
	r.Check(map[string]interface{}{"data": r.Repo.GetDisabledTokens()}, "dao_operations_v2/disabled_tokens.json")
}
