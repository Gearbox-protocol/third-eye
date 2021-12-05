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
