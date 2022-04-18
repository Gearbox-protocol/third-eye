package repository

import (
	"math/big"
	"reflect"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
)

func (repo *Repository) loadCreditSessions(lastDebtSync int64) {
	defer utils.Elapsed("loadCreditSessions")()
	data := []*schemas.CreditSession{}
	err := repo.db.Raw(`SELECT * FROM credit_sessions cs 
	JOIN (SELECT distinct on (session_id) collateral_usd, collateral_underlying, session_id FROM credit_session_snapshots ORDER BY session_id, block_num DESC) css
	ON css.session_id = cs.id
	WHERE status = ? OR (status <> ? AND closed_at > ?)`,
		schemas.Active, schemas.Active, lastDebtSync).Find(&data).Error
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

func (repo *Repository) addCreditSession(session *schemas.CreditSession, loadedFromDB bool) {
	if repo.sessions[session.ID] == nil {
		if !loadedFromDB {
			log.Infof("Add session %s", session.ID)
		}
		repo.sessions[session.ID] = session
	} else {
		log.Fatalf("Credit session already present %s", session.ID)
	}
}

func (repo *Repository) AddCreditSession(session *schemas.CreditSession, loadedFromDB bool, txHash string, logID uint) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.addCreditSession(session, loadedFromDB)
	repo.accountManager.AddAccountDetails(&ds.SessionData{
		Since:         session.Since,
		Account:       session.Account,
		CreditManager: session.CreditManager,
		SessionID:     session.ID,
		OpenTxHash:    txHash,
		OpenLogId:     logID,
	})
}

func (repo *Repository) GetCreditSession(sessionId string) *schemas.CreditSession {
	return repo.sessions[sessionId]
}
func (repo *Repository) UpdateCreditSession(sessionId string, values map[string]interface{}) *schemas.CreditSession {
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

func (repo *Repository) GetSessions() map[string]*schemas.CreditSession {
	return repo.sessions
}

// for account manager
func (repo *Repository) loadAccountLastSession() {
	defer utils.Elapsed("loadAccountLastSession")()
	data := []*ds.SessionData{}
	err := repo.db.Raw(`SELECT t1.*,t2.*, t3.closed_tx_hash, t3.closed_log_id 
		FROM (SELECT DISTINCT ON (account) credit_manager, since, id, closed_at, account 
				FROM credit_sessions ORDER BY account, since DESC) t1 JOIN 
			(SELECT session_id, tx_hash open_tx_hash, log_id open_log_id
				FROM account_operations WHERE action like 'OpenCreditAccount%') t2 ON t1.id = t2.session_id
		LEFT JOIN (SELECT session_id, tx_hash closed_tx_hash, log_id closed_log_id
			FROM account_operations 
			WHERE (action like 'CloseCreditAccount%' 
			or action like 'LiquidateCreditAccount%' 
			or action like 'RepayCreditAccount%')) t3 ON t1.id = t3.session_id`).Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range data {
		repo.accountManager.AddAccountDetails(entry)
	}
}

func (repo *Repository) GetAccountManager() *ds.AccountTokenManager {
	return repo.accountManager
}

func (repo *Repository) AddAccountAddr(account string) {
	addrs := repo.kit.GetAdapterAddressByName(ds.AccountManager)
	if len(addrs) == 1 {
		adapter := repo.GetAdapter(addrs[0])
		adapter.SetDetails(account)
	} else {
		log.Fatalf("%d account manager model available in adapter kit", len(addrs))
	}
}
