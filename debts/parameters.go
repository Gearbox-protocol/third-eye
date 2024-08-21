package debts

import (
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

func (eng *DebtEngine) loadParameters(_ int64) {
	defer utils.Elapsed("Debt(loadLastDebts)")()
	data := []*schemas.Parameters{}
	err := eng.db.Raw("SELECT distinct on (credit_manager) * FROM parameters ORDER BY credit_manager, block_num DESC;").Find(&data).Error
	log.CheckFatal(err)
	for _, entry := range data {
		eng.addLastParameters(entry)
	}
}

func (eng *DebtEngine) addLastParameters(params *schemas.Parameters) {
	eng.lastParameters[params.CreditManager] = params
}
