package repository

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
)

func (repo *Repository) loadLastCSS(lastDebtSync int64) {
	data := []*core.CreditSessionSnapshot{}
	query := `SELECT css_2.* FROM credit_session_snapshots as css_2 JOIN
		(SELECT session_id, max(block_num) AS block_num FROM credit_session_snapshots WHERE block_num <= ? GROUP BY session_id) AS css
		ON css_2.block_num = css.block_num AND css_2.session_id = css.session_id`
	err := repo.db.Raw(query, lastDebtSync).Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, sessionSnapshot := range data {
		repo.AddLastCSS(sessionSnapshot)
	}
}

func (repo *Repository) AddLastCSS(css *core.CreditSessionSnapshot) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.lastCSS[css.SessionId] = css
}

func (repo *Repository) GetLastCSS(sessionId string) *core.CreditSessionSnapshot {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	css := repo.lastCSS[sessionId]
	if css == nil {
		log.Infof("Last Credit session snapshot not found: %s", sessionId)
		repo.lastCSS[sessionId] = &core.CreditSessionSnapshot{SessionId: sessionId, Balances: &core.JsonBalance{}}
		css = repo.lastCSS[sessionId]
	}
	return css
}

func (repo *Repository) AddCreditSessionSnapshot(css *core.CreditSessionSnapshot) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	css.ID = 0
	repo.blocks[css.BlockNum].AddCreditSessionSnapshot(css)
}
