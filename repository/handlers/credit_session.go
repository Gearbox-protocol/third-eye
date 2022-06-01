package handlers

import (
	"math/big"
	"reflect"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"gorm.io/gorm"
)

type SessionRepo struct {
	// changed during syncing
	sessions map[string]*schemas.CreditSession
}

func NewSessionRepo() *SessionRepo {
	return &SessionRepo{
		sessions: map[string]*schemas.CreditSession{},
	}
}

// where clause is for debts
// so that all the credit sessions that were present at lastDebtsync +1 can be loaded from db.
//
// join with css is not dependent on lastdebtsync block
// these values are for use by sync engine
func (repo *SessionRepo) loadCreditSessions(db *gorm.DB, lastDebtSync int64) {
	defer utils.Elapsed("loadCreditSessions")()
	data := []*schemas.CreditSession{}
	err := db.Raw(`SELECT * FROM credit_sessions cs 
	JOIN (SELECT distinct on (session_id) collateral_usd, collateral_underlying, session_id FROM credit_session_snapshots ORDER BY session_id, block_num DESC) css
	ON css.session_id = cs.id
	WHERE status = ? OR (status <> ? AND closed_at > ?)`,
		schemas.Active, schemas.Active, lastDebtSync).Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, session := range data {
		repo.AddCreditSession(session, true)
	}
}

func (repo *SessionRepo) AddCreditSession(session *schemas.CreditSession, loadedFromDB bool) {
	if repo.sessions[session.ID] == nil {
		if !loadedFromDB {
			log.Infof("Add session %s", session.ID)
		}
		repo.sessions[session.ID] = session
	} else {
		log.Fatalf("Credit session already present %s", session.ID)
	}
}

func (repo *SessionRepo) GetCreditSession(sessionId string) *schemas.CreditSession {
	return repo.sessions[sessionId]
}

func (repo *SessionRepo) GetSessions() map[string]*schemas.CreditSession {
	return repo.sessions
}

func (repo *SessionRepo) UpdateCreditSession(sessionId string, values map[string]interface{}) *schemas.CreditSession {
	session := repo.sessions[sessionId]
	session.IsDirty = true
	ref := reflect.ValueOf(session).Elem()
	for k, v := range values {
		switch v.(type) {
		case string:
			ref.FieldByName(k).SetString(v.(string))
		case int64:
			ref.FieldByName(k).SetInt(v.(int64))
		case int:
			ref.FieldByName(k).SetInt(int64(v.(int)))
		case *big.Int:
			val := (*core.BigInt)(v.(*big.Int))
			pointer := reflect.ValueOf(val)
			ref.FieldByName(k).Set(pointer)
		default:
			log.Fatal("Not able to set %s %v", k, v)
		}
	}
	return session
}

func (repo *SessionRepo) Clear(closedBefore int64) {
	for _, session := range repo.sessions {
		if session.ClosedAt != 0 && closedBefore >= session.ClosedAt {
			delete(repo.sessions, session.ID)
		}
	}
}
