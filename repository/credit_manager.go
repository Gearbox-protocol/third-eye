package repository

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
)

func (repo *Repository) loadCreditManagers() {
	data := []*core.CreditManagerState{}
	err := repo.db.Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, cm := range data {
		adapter := repo.kit.GetAdapter(cm.Address)
		if adapter != nil && adapter.GetName() == "CreditManager" {
			cm.Sessions = map[string]string{}
			adapter.SetUnderlyingState(cm)
		}
	}
	repo.loadSessionIdToBorrower()
}

func (repo *Repository) loadSessionIdToBorrower() {
	data := []*core.CreditSession{}
	err := repo.db.Raw(`SELECT credit_manager, id, borrower FROM credit_sessions where status=0;`).Find(&data).Error
	log.CheckFatal(err)
	borrowerToSession := map[string]map[string]string{}
	for _, cs := range data {
		hstore := borrowerToSession[cs.CreditManager]
		if hstore == nil {
			borrowerToSession[cs.CreditManager] = map[string]string{}
			hstore = borrowerToSession[cs.CreditManager]
		}
		hstore[cs.Borrower] = cs.ID
	}
	for cm, hstore := range borrowerToSession {
		adapter := repo.kit.GetAdapter(cm)
		if adapter != nil && adapter.GetName() == "CreditManager" {
			adapter.SetUnderlyingState(hstore)
		}
	}
}

func (repo *Repository) GetCMState(cmAddr string) *core.CreditManagerState {
	state := repo.kit.GetAdapter(cmAddr).GetUnderlyingState()
	cm, ok := state.(*core.CreditManagerState)
	if !ok {
		log.Fatal("Type assertion for credit manager state failed")
	}
	return cm
}

func (repo *Repository) GetUnderlyingDecimal(cmAddr string) int8 {
	cm := repo.GetCMState(cmAddr)
	return repo.GetToken(cm.UnderlyingToken).Decimals
}

func (repo *Repository) AddCreditManagerStats(cms *core.CreditManagerStat) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.GetBlock(cms.BlockNum).AddCreditManagerStats(cms)
}

func (repo *Repository) AddRepayOnCM(blockNum int64, cmAddr string, pnlOnRepay core.PnlOnRepay) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.GetBlock(blockNum).AddRepayOnCM(cmAddr, &pnlOnRepay)
}

func (repo *Repository) GetRepayOnCM(blockNum int64, cmAddr string) *core.PnlOnRepay {
	return repo.GetBlock(blockNum).GetRepayOnCM(cmAddr)
}

func (repo *Repository) AddParameters(params *core.Parameters) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.GetBlock(params.BlockNum).AddParameters(params)
}
