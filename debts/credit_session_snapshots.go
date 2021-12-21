package debts

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
)

func (eng *DebtEngine) loadLastCSS(lastDebtSync int64) {
	data := []*core.CreditSessionSnapshot{}
	query := `SELECT css_2.* FROM credit_session_snapshots as css_2 JOIN
		(SELECT session_id, max(block_num) AS block_num FROM credit_session_snapshots WHERE block_num <= ? GROUP BY session_id) AS css
		ON css_2.block_num = css.block_num AND css_2.session_id = css.session_id`
	err := eng.db.Raw(query, lastDebtSync).Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, sessionSnapshot := range data {
		eng.AddLastCSS(sessionSnapshot)
	}
}

func (eng *DebtEngine) AddLastCSS(css *core.CreditSessionSnapshot) {
	eng.lastCSS[css.SessionId] = css
}

func (eng *DebtEngine) GetLastCSS(sessionId string) *core.CreditSessionSnapshot {
	css := eng.lastCSS[sessionId]
	if css == nil {
		log.Infof("Last Credit session snapshot not found: %s", sessionId)
		eng.lastCSS[sessionId] = &core.CreditSessionSnapshot{SessionId: sessionId, Balances: &core.JsonBalance{}}
		css = eng.lastCSS[sessionId]
	}
	return css
}
