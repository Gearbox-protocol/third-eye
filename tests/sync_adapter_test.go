package tests

import (
	"testing"

	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/debts"
	"github.com/Gearbox-protocol/third-eye/engine"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/repository"
	"github.com/Gearbox-protocol/third-eye/tests/framework"
	"github.com/Gearbox-protocol/third-eye/utils"
)

func TestSyncAdapters(t *testing.T) {
	log.SetTestLogging(t)
	client := framework.NewTestClient()
	cfg := &config.Config{}
	ep := framework.NewMockExecuteParser()
	repo := repository.GetRepository(nil, client, cfg, ep)
	debtEng := debts.GetDebtEngine(nil, client, cfg, repo, true)
	eng := engine.NewEngine(cfg, client, debtEng, repo)
	r := framework.NewMockRepo(repo, client, t, eng, ep)
	r.Init([]string{"sync_adapters/input.json"})
	log.Info(utils.ToJson(r.AddressMap))
	eng.Sync(10)

	adapters := getAdapters(repo.GetKit())
	r.Check(map[string]interface{}{"data": adapters}, "sync_adapters/adapters.json")
}

func getAdapters(kit *core.AdapterKit) (array []*core.SyncAdapter) {
	for lvlIndex := 0; lvlIndex < kit.Len(); lvlIndex++ {
		for kit.Next(lvlIndex) {
			array = append(array, kit.Get(lvlIndex).GetAdapterState())
		}
		kit.Reset(lvlIndex)
	}
	return
}
