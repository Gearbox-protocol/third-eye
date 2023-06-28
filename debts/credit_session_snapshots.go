package debts

import (
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

func (eng *DebtEngine) loadLastCSS(lastDebtSync int64) {
	defer utils.Elapsed("Debt(loadLastCSS)")()
	data := []*schemas.CreditSessionSnapshot{}
	query := `SELECT css_2.* FROM credit_session_snapshots as css_2 JOIN
		(SELECT session_id, max(block_num) AS block_num FROM credit_session_snapshots WHERE block_num <= ? GROUP BY session_id) AS css
		ON css_2.block_num = css.block_num AND css_2.session_id = css.session_id`
	err := eng.db.Raw(query, lastDebtSync).Find(&data).Error
	log.CheckFatal(err)
	for _, sessionSnapshot := range data {
		eng.AddLastCSS(sessionSnapshot)
	}
}

func (eng *DebtEngine) AddLastCSS(css *schemas.CreditSessionSnapshot) {
	eng.lastCSS[css.SessionId] = css
}

func (eng *DebtEngine) loadLastRebaseDetails(lastDebtSync int64) {
	obj := &schemas.RebaseDetailsForDB{}
	query := `SELECT * FROM rebase_details WHERE block_num <= ? order by block_num DESC Limit 1`
	err := eng.db.Raw(query, lastDebtSync).Find(obj).Error
	log.CheckFatal(err)
	if obj.BlockNum != 0 {
		eng.lastRebaseDetails = obj
	}
}
