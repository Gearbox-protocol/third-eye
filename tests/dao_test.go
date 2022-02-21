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

func TestDAOOperations(t *testing.T) {
	log.SetTestLogging(t)
	client := framework.NewTestClient()
	cfg := &config.Config{}
	repo := repository.GetRepository(nil, client, cfg, nil)
	debtEng := debts.NewDebtEngine(nil, client, cfg, repo)
	eng := engine.NewEngine(cfg, client, debtEng, repo)
	r := framework.NewMockRepo(repo, client, "dao_operations/input.json", t, eng)
	r.Init()
	log.Info(utils.ToJson(r.AddressMap))
	eng.Sync(10)

	r.Check(map[string]interface{}{"data": repo.GetBlocks()[3].DAOOperations}, "dao_operations/blocks.json")
}
