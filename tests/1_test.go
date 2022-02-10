package tests

import (
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/debts"
	"github.com/Gearbox-protocol/third-eye/engine"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/repository"
	// "github.com/Gearbox-protocol/third-eye/log"
	// "github.com/Gearbox-protocol/third-eye/utils"
	"testing"
)

func TestRepo(t *testing.T) {
	client := ethclient.NewTestClient()
	cfg := &config.Config{}
	repo := repository.GetRepository(nil, client, cfg, nil)
	debtEng := debts.NewDebtEngine(nil, client, cfg, repo)
	eng := engine.NewEngine(cfg, client, debtEng, repo)
	r := MockRepo{
		repo:          repo,
		client:        client,
		file:          "test1.json",
		t:             t,
		eng:           eng,
		addressToType: make(map[string]string),
		feedToToken:   make(map[string]string),
	}
	r.init()
	// log.Info(utils.ToJson(r.AddressMap))
	eng.Sync(10)
	debtEng.CalculateDebt()
}
