package framework

import (
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/debts"
	"github.com/Gearbox-protocol/third-eye/engine"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/repository"
	"github.com/Gearbox-protocol/third-eye/utils"
	"testing"
)

func NewEngs(t *testing.T, inputFiles []string) (MockRepo, core.DebtEngineI) {
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
