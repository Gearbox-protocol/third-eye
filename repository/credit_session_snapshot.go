package repository

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
)

func (repo *Repository) loadLastCSS() {
	data := []*core.CreditSessionSnapshot{}
	query := `select distinct on (session_id) borrower, session_id , status, balances, cs.borrowed_amount from
		credit_sessions as cs inner join credit_session_snapshots as css on css.session_id = cs.id
		where status=0 order by session_id, block_num desc`
	err := repo.db.Raw(query).Find(&data).Error
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
	if repo.lastCSS[css.SessionId] == nil {
		repo.lastCSS[css.SessionId] = css
	} else {
		log.Fatalf("Credit session snapshot already present %s", css.SessionId)
	}
}

func (repo *Repository) GetLastCSS(sessionId string) *core.CreditSessionSnapshot {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	css := repo.lastCSS[sessionId]
	if css == nil {
		log.Infof("Last Credit session snapshot not found: %s", sessionId)
		repo.lastCSS[sessionId] = &core.CreditSessionSnapshot{SessionId: sessionId, Balances: make(core.JsonBalance)}
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
