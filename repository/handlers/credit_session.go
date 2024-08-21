package handlers

import (
	"math/big"
	"reflect"
	"sync"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SessionRepo struct {
	// changed during syncing
	sessions map[string]*schemas.CreditSession
	mu       *sync.Mutex
}

func NewSessionRepo() *SessionRepo {
	return &SessionRepo{
		sessions: map[string]*schemas.CreditSession{},
		mu:       &sync.Mutex{},
	}
}

// load/save

// where clause is for debts
// so that all the credit sessions that were present at lastDebtsync +1 can be loaded from db.
//
// join with css is not dependent on lastdebtsync block
// these values are for use by sync engine
func (repo *SessionRepo) LoadCreditSessions(db *gorm.DB, lastDebtSync int64) {
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

func (repo *SessionRepo) Save(tx *gorm.DB) {
	utils.Elapsed("session sql statements")()
	for _, session := range repo.GetSessions() {
		if session.IsDirty {
			err := tx.Clauses(clause.OnConflict{
				UpdateAll: true,
			}).Create(session).Error
			log.CheckFatal(err)
			session.IsDirty = false
		}
	}
}

// external funcs
func (repo *SessionRepo) AddCreditSession(session *schemas.CreditSession, loadedFromDB bool) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
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
	repo.mu.Lock()
	defer repo.mu.Unlock()
	return repo.sessions[sessionId]
}

func (repo *SessionRepo) GetSessions() map[string]*schemas.CreditSession {
	return repo.sessions
}

func (repo *SessionRepo) UpdateCreditSession(sessionId string, values map[string]interface{}) *schemas.CreditSession {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	session := repo.sessions[sessionId]
	session.IsDirty = true
	ref := reflect.ValueOf(session).Elem()
	for k, v := range values {
		switch typedV := v.(type) {
		case string:
			ref.FieldByName(k).SetString(typedV)
		case int64:
			ref.FieldByName(k).SetInt(typedV)
		case int:
			ref.FieldByName(k).SetInt(int64(typedV))
		case *big.Int:
			val := (*core.BigInt)(typedV)
			pointer := reflect.ValueOf(val)
			ref.FieldByName(k).Set(pointer)
		default:
			log.Fatalf("Not able to set %s %v", k, v)
		}
	}
	return session
}

func (repo *SessionRepo) Clear(closedBefore int64) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	for _, session := range repo.sessions {
		if session.ClosedAt != 0 && closedBefore >= session.ClosedAt {
			delete(repo.sessions, session.ID)
		}
	}
}