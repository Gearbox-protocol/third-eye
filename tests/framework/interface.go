package framework

import (
	"testing"

	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/test"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/debts"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/engine"
	"github.com/Gearbox-protocol/third-eye/repository"
	"github.com/Gearbox-protocol/third-eye/repository/handlers"
)

func NewEngs(t *testing.T, inputFiles []string) (MockRepo, ds.DebtEngineI) {
	log.SetTestLogging(t)
	client := test.NewTestClient()
	cfg := &config.Config{}
	ep := NewMockExecuteParser()
	repo := repository.GetRepository(nil, client, cfg, handlers.NewExtraRepo(client, ep))
	debtEng := debts.GetDebtEngine(nil, client, cfg, repo, true)
	eng := engine.NewEngine(cfg, client, debtEng, nil, repo)
	r := NewMockRepo(repo, client, t, eng, ep)
	r.Init(inputFiles)
	return r, debtEng
}
