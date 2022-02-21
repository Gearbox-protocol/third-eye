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

func TestOpenCreditAccount(t *testing.T) {
	log.SetTestLogging(t)
	client := framework.NewTestClient()
	cfg := &config.Config{}
	repo := repository.GetRepository(nil, client, cfg, nil)
	debtEng := debts.NewDebtEngine(nil, client, cfg, repo)
	eng := engine.NewEngine(cfg, client, debtEng, repo)
	 r := framework.NewMockRepo(repo, client,  "open_credit_account/input.json", t, eng)
	r.Init()
	log.Info(utils.ToJson(r.AddressMap))
	eng.Sync(10)
	r.Check(t, repo.GetBlocks()[3], "open_credit_account/blocks.json")
	debtEng.CalculateDebt()
	r.Check(t, debtEng.GetDebts(), "open_credit_account/debts.json")
}

