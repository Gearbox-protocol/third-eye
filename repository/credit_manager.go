package repository

import (
	"reflect"

	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

func (repo *Repository) loadCreditManagers() {
	defer utils.Elapsed("loadCreditManagers")()
	data := []*schemas.CreditManagerState{}
	err := repo.db.Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, cm := range data {
		adapter := repo.GetAdapter(cm.Address)
		if adapter != nil && adapter.GetName() == "CreditManager" {
			cm.Sessions = map[string]string{}
			adapter.SetUnderlyingState(cm)
		}
	}
	repo.loadSessionIdToBorrower()
	repo.loadParametersToCM()
}

func (repo *Repository) loadParametersToCM() {
	defer utils.Elapsed("loadCreditManagers")()
	data := []*schemas.Parameters{}
	err := repo.db.Raw("SELECT distinct on(credit_manager) * FROM parameters ORDER BY credit_manager, block_num DESC").Error
	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range data {
		adapter := repo.GetAdapter(entry.CreditManager)
		if adapter != nil && adapter.GetName() == "CreditManager" {
			adapter.SetUnderlyingState(entry)
		}
	}
}

func (repo *Repository) loadSessionIdToBorrower() {
	data := []*schemas.CreditSession{}
	err := repo.db.Raw(`SELECT credit_manager, id, borrower FROM credit_sessions where status=0;`).Find(&data).Error
	log.CheckFatal(err)
	borrowerToSession := map[string]map[string]string{}
	for _, cs := range data {
		hstore := borrowerToSession[cs.CreditManager]
		if hstore == nil {
			borrowerToSession[cs.CreditManager] = map[string]string{}
			hstore = borrowerToSession[cs.CreditManager]
		}
		hstore[cs.Borrower] = cs.ID
	}
	for cm, hstore := range borrowerToSession {
		adapter := repo.GetAdapter(cm)
		if adapter != nil && adapter.GetName() == "CreditManager" {
			adapter.SetUnderlyingState(hstore)
		}
	}
}

// safe
func (repo *Repository) GetCMState(cmAddr string) *schemas.CreditManagerState {
	adapter := repo.GetAdapter(cmAddr)
	// if cm doesn't exist return nil, it is used by debt engine
	// if the block_num is before cm exist don't error
	// adapter not equal to nil is not used as underlying type of adapter is not nil
	if reflect.ValueOf(adapter).IsZero() {
		return nil
	}
	state := adapter.GetUnderlyingState()
	cm, ok := state.(*schemas.CreditManagerState)
	if !ok {
		log.Fatal("Type assertion for credit manager state failed")
	}
	return cm
}

func (repo *Repository) GetUnderlyingDecimal(cmAddr string) int8 {
	cm := repo.GetCMState(cmAddr)
	return repo.GetToken(cm.UnderlyingToken).Decimals
}

func (repo *Repository) AddAccountTokenTransfer(tt *schemas.TokenTransfer) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.SetBlock(tt.BlockNum)
	repo.accountManager.AddTokenTransfer(tt)
}

func (repo *Repository) AddRepayOnCM(cmAddr string, pnlOnRepay schemas.PnlOnRepay) {
	adapter := repo.GetAdapter(cmAddr)
	adapter.SetUnderlyingState(&pnlOnRepay)
}
