package tests

import (
	"testing"

	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/debts"
	"github.com/Gearbox-protocol/third-eye/engine"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/repository"
	"github.com/Gearbox-protocol/third-eye/utils"
)

func TestLifecycleCreditAccount(t *testing.T) {
	log.SetTestLogging(t)
	client := NewTestClient()
	cfg := &config.Config{}
	repo := repository.GetRepository(nil, client, cfg, nil)
	debtEng := debts.NewDebtEngine(nil, client, cfg, repo)
	eng := engine.NewEngine(cfg, client, debtEng, repo)
	r := MockRepo{
		repo:          repo,
		client:        client,
		file:          "account_lifecycle/input.json",
		t:             t,
		eng:           eng,
		addressToType: make(map[string]string),
		feedToToken:   make(map[string]string),
	}
	r.init()
	log.Info(utils.ToJson(r.AddressMap))
	eng.Sync(10)

	outputBlocks := repo.GetBlocks()
	delete(outputBlocks, 2)
	r.check(t, outputBlocks, "account_lifecycle/blocks.json")
	debtEng.CalculateDebt()
	r.check(t, debtEng.GetDebts(), "account_lifecycle/debts.json")
}