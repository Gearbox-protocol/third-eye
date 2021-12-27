package debts

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
)

func (eng *DebtEngine) loadParameters(blockNum int64) {
	data := []*core.Parameters{}
	err := eng.db.Raw("SELECT distinct on (credit_manager) * FROM parameters ORDER BY credit_manager, block_num DESC;").Find(&data).Error
	log.CheckFatal(err)
	for _, entry := range data {
		eng.addLastParameters(entry)
	}
}

func (eng *DebtEngine) addLastParameters(params *core.Parameters) {
	eng.lastParameters[params.CreditManager] = params
}
