package repository

import (
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
)

func (repo *Repository) AddCreditSession(session *schemas.CreditSession, loadedFromDB bool, txHash string, logID uint) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.SessionRepo.AddCreditSession(session, loadedFromDB)
	repo.accountManager.AddAccountDetails(&ds.SessionData{
		Since:         session.Since,
		Account:       session.Account,
		CreditManager: session.CreditManager,
		SessionID:     session.ID,
		OpenTxHash:    txHash,
		OpenLogId:     logID,
	})
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

func (repo *Repository) GetAccountManager() *ds.DirectTransferManager {
	return repo.accountManager
}

// add account addr with account manager
// this func is currently used by account factory
func (repo *Repository) AddAccountAddr(account string) {
	addrs := repo.GetAdapterAddressByName(ds.AccountManager)
	if len(addrs) == 1 {
		adapter := repo.GetAdapter(addrs[0])
		adapter.SetDetails(account)
	} else {
		log.Fatalf("%d account manager model available in adapter kit", len(addrs))
	}
}
