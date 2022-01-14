package repository

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
)

func (repo *Repository) loadCreditSessions(lastDebtSync int64) {
	data := []*core.CreditSession{}
	err := repo.db.Raw(`SELECT * FROM credit_sessions cs 
	JOIN (SELECT distinct on (session_id) collateral_in_usd, session_id FROM credit_session_snapshots ORDER BY session_id, block_num DESC) css
	ON css.session_id = cs.id
	WHERE status = ? OR (status <> ? AND closed_at > ?)`,
		core.Active, core.Active, lastDebtSync).Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, session := range data {
		repo.addCreditSession(session, true)
	}
}

func (repo *Repository) AddDataCompressor(blockNum int64, addr string) {
	repo.dcWrapper.AddDataCompressor(blockNum, addr)
}

func (repo *Repository) addCreditSession(session *core.CreditSession, loadedFromDB bool) {
	if repo.sessions[session.ID] == nil {
		if !loadedFromDB {
			log.Infof("Add session %s", session.ID)
		}
		repo.sessions[session.ID] = session
	} else {
		log.Fatalf("Credit session already present %s", session.ID)
	}
}

func (repo *Repository) AddCreditSession(session *core.CreditSession, loadedFromDB bool) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.addCreditSession(session, loadedFromDB)
}

func (repo *Repository) GetCreditSession(sessionId string) *core.CreditSession {
	return repo.sessions[sessionId]
}

func (repo *Repository) GetSessions() map[string]*core.CreditSession {
	return repo.sessions
}
