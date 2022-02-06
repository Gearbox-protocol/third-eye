package repository

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
)

func (repo *Repository) loadGearBalances() {
	data := []*core.GearBalance{}
	err := repo.db.Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	geartokenAddr := repo.kit.GetAdapterAddressByName(core.GearToken)
	if len(geartokenAddr) > 0 {
		if adapter := repo.GetAdapter(geartokenAddr[0]); adapter != nil {
			adapter.SetUnderlyingState(data)
		}
	}
}
