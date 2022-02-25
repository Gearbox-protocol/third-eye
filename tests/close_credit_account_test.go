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

func TestCloseCreditAccount(t *testing.T) {
	log.SetTestLogging(t)
	client := framework.NewTestClient()
	cfg := &config.Config{}
	ep := framework.NewMockExecuteParser()
	repo := repository.GetRepository(nil, client, cfg, ep)
	debtEng := debts.NewDebtEngine(nil, client, cfg, repo)
	eng := engine.NewEngine(cfg, client, debtEng, repo)
	r := framework.NewMockRepo(repo, client, t, eng, ep)
	r.Init([]string{"account_lifecycle/input.json", "close_credit_account/input.json"})
	log.Info(utils.ToJson(r.AddressMap))
	eng.Sync(10)
	r.Check(repo.GetBlocks()[7], "close_credit_account/blocks.json")
}
