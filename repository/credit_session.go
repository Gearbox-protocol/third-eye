package repository

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
)

func (repo *Repository) loadCreditSessions(lastDebtSync int64) {
	data := []*core.CreditSession{}
	err := repo.db.Find(&data, "status = ? OR (status <> ? AND closed_at > ?)",
		core.Active, core.Active, lastDebtSync).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, session := range data {
		repo.AddCreditSession(session, true)
	}
}

func (repo *Repository) AddDataCompressor(blockNum int64, addr string) {
	repo.dcWrapper.AddDataCompressor(blockNum, addr)
}

func (repo *Repository) AddCreditSession(session *core.CreditSession, loadedFromDB bool) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if repo.sessions[session.ID] == nil {
		if !loadedFromDB {
			log.Infof("Add creditAccount(%s) with sessionId %s", session.Account, session.ID)
		}
		repo.sessions[session.ID] = session
	} else {
		log.Fatalf("Credit session already present %s", session.ID)
	}
}

func (repo *Repository) GetCreditSession(sessionId string) *core.CreditSession {
	return repo.sessions[sessionId]
}

func (repo *Repository) GetSessions() map[string]*core.CreditSession {
	return repo.sessions
}
