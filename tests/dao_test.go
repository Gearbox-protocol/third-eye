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

func TestDAOOperations(t *testing.T) {
	client := NewTestClient()
	cfg := &config.Config{}
	repo := repository.GetRepository(nil, client, cfg, nil)
	debtEng := debts.NewDebtEngine(nil, client, cfg, repo)
	eng := engine.NewEngine(cfg, client, debtEng, repo)
	r := MockRepo{
		repo:          repo,
		client:        client,
		file:          "dao_operations_input.json",
		t:             t,
		eng:           eng,
		addressToType: make(map[string]string),
		feedToToken:   make(map[string]string),
	}
	r.init()
	log.Info(utils.ToJson(r.AddressMap))
	eng.Sync(10)

	r.check(t, map[string]interface{}{"data": repo.GetBlocks()[3].DAOOperations}, "dao_operations_blocks.json")
}
