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
			adapter.SetUnderlyingState(cm)
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
	repo.GetBlock(cms.BlockNum).AddCreditManagerStats(cms)
}

func (repo *Repository) AddRepayOnCM(blockNum int64, cmAddr string, pnlOnRepay core.PnlOnRepay) {
	repo.GetBlock(blockNum).AddRepayOnCM(cmAddr, &pnlOnRepay)
}

func (repo *Repository) GetRepayOnCM(blockNum int64, cmAddr string) *core.PnlOnRepay {
	return repo.GetBlock(blockNum).GetRepayOnCM(cmAddr)
}