package debts

import (
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

// pool interest state fetch
func (eng *DebtEngine) loadPoolLastInterestData(lastDebtSync int64) {
	defer utils.Elapsed("Debt(loadPoolLastInterestData)")()
	data := []*schemas.PoolInterestData{}
	query := `SELECT * FROM pool_stats 
	JOIN (SELECT max(block_num) as bn, pool FROM pool_stats WHERE block_num <= ? group by pool) as p
	JOIN blocks ON p.bn = blocks.id
	ON p.bn = pool_stats.block_num
	AND p.pool = pool_stats.pool;`
	err := eng.db.Raw(query, lastDebtSync).Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, pd := range data {
		eng.AddPoolLastInterestData(pd)
	}
}

func (eng *DebtEngine) AddPoolLastInterestData(pd *schemas.PoolInterestData) {
	eng.poolLastInterestData[pd.Address] = pd
}
