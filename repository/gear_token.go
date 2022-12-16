package repository

import (
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
)

func (repo *Repository) loadGearBalances() {
	defer utils.Elapsed("loadGearBalances")()
	data := []*schemas.GearBalance{}
	err := repo.db.Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	geartokenAddr := repo.GetAdapterAddressByName(ds.GearToken)
	if len(geartokenAddr) > 0 {
		if adapter := repo.GetAdapter(geartokenAddr[0]); adapter != nil {
			adapter.SetUnderlyingState(data)
		}
	}
}
