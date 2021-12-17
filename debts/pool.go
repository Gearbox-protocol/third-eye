package debts

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
)

// pool interest state fetch
func (eng *DebtEngine) loadPoolLastInterestData(lastDebtSync int64) {
	data := []*core.PoolInterestData{}
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

func (eng *DebtEngine) AddPoolLastInterestData(pd *core.PoolInterestData) {
	eng.poolLastInterestData[pd.Address] = pd
}
