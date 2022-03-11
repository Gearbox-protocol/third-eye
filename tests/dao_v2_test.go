package tests

import (
	"testing"

	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/debts"
	"github.com/Gearbox-protocol/third-eye/engine"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/repository"
	"github.com/Gearbox-protocol/third-eye/tests/framework"
	"github.com/Gearbox-protocol/third-eye/utils"
)

func TestDAOV2Operations(t *testing.T) {
	log.SetTestLogging(t)
	client := framework.NewTestClient()
	cfg := &config.Config{}
	ep := framework.NewMockExecuteParser()
	repo := repository.GetRepository(nil, client, cfg, ep)
	debtEng := debts.GetDebtEngine(nil, client, cfg, repo, true)
	eng := engine.NewEngine(cfg, client, debtEng, repo)
	r := framework.NewMockRepo(repo, client, t, eng, ep)
	r.Init([]string{"dao_operations_v2/input.json"})
	log.Info(utils.ToJson(r.AddressMap))
	eng.Sync(10)

	data := map[string]interface{}{
		"3": map[string]interface{}{
			"daoOperations": repo.GetBlocks()[3].DAOOperations,
			"allowedTokens": repo.GetBlocks()[3].AllowedTokens,
		},
		"4": map[string]interface{}{
			"daoOperations": repo.GetBlocks()[4].DAOOperations,
			"allowedTokens": repo.GetBlocks()[4].AllowedTokens,
		},
		"5": map[string]interface{}{
			"daoOperations": repo.GetBlocks()[5].DAOOperations,
			"allowedTokens": repo.GetBlocks()[5].AllowedTokens,
		},
	}
	r.Check(data, "dao_operations_v2/blocks.json")
	r.Check(map[string]interface{}{"data":repo.GetDisabledTokens()}, "dao_operations_v2/disabled_tokens.json")
}
