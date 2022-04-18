package framework

import (
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/debts"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/engine"
	"github.com/Gearbox-protocol/third-eye/repository"
	"testing"
)

func NewEngs(t *testing.T, inputFiles []string) (MockRepo, ds.DebtEngineI) {
	log.SetTestLogging(t)
	client := NewTestClient()
	cfg := &config.Config{}
	ep := NewMockExecuteParser()
	repo := repository.GetRepository(nil, client, cfg, ep)
	debtEng := debts.GetDebtEngine(nil, client, cfg, repo, true)
	eng := engine.NewEngine(cfg, client, debtEng, repo)
	r := NewMockRepo(repo, client, t, eng, ep)
	r.Init(inputFiles)
	log.Info(utils.ToJson(r.AddressMap))
	return r, debtEng
}
