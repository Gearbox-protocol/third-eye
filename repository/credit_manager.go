package repository

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
)

func (repo *Repository) loadCreditManagers() {
	data := []*core.CreditManagerState{}
	err := repo.db.Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, cm := range data {
		adapter := repo.kit.GetAdapter(cm.Address)
		if adapter != nil && adapter.GetName() == "CreditManager" {
			cm.Sessions = map[string]string{}
			adapter.SetUnderlyingState(cm)
		}
	}
	repo.loadSessionIdToBorrower()
}

func (repo *Repository) loadSessionIdToBorrower() {
	data := []*core.CreditSession{}
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
		adapter := repo.kit.GetAdapter(cm)
		if adapter != nil && adapter.GetName() == "CreditManager" {
			adapter.SetUnderlyingState(hstore)
		}
	}
}

func (repo *Repository) GetCMState(cmAddr string) *core.CreditManagerState {
	state := repo.kit.GetAdapter(cmAddr).GetUnderlyingState()
	cm, ok := state.(*core.CreditManagerState)
	if !ok {
		log.Fatal("Type assertion for credit manager state failed")
	}
	return cm
}

func (repo *Repository) GetUnderlyingDecimal(cmAddr string) int8 {
	cm := repo.GetCMState(cmAddr)
	return repo.GetToken(cm.UnderlyingToken).Decimals
}

func (repo *Repository) AddCreditManagerStats(cms *core.CreditManagerStat) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.setAndGetBlock(cms.BlockNum).AddCreditManagerStats(cms)
}

func (repo *Repository) AddRepayOnCM(blockNum int64, cmAddr string, pnlOnRepay core.PnlOnRepay) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.setAndGetBlock(blockNum).AddRepayOnCM(cmAddr, &pnlOnRepay)
}

func (repo *Repository) GetRepayOnCM(blockNum int64, cmAddr string) *core.PnlOnRepay {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	return repo.blocks[blockNum].GetRepayOnCM(cmAddr)
}

func (repo *Repository) AddParameters(logID uint, txHash string, params *core.Parameters) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.setAndGetBlock(params.BlockNum).AddParameters(params)
	// cal dao action
	oldCMParams := repo.cmParams[params.CreditManager]
	if oldCMParams == nil {
		oldCMParams = core.NewParameters()
	}

	repo.addDAOOperation(&core.DAOOperation{
		BlockNumber: params.BlockNum,
		LogID:       logID,
		TxHash:      txHash,
		Contract:    params.CreditManager,
		Args:        oldCMParams.Diff(params),
	})
	//
	repo.cmParams[params.CreditManager] = params
}

func (repo *Repository) loadAllParams() {
	// parameters
	data := []*core.Parameters{}
	err := repo.db.Raw(`SELECT distinct on (credit_manager) * FROM parameters 
		ORDER BY credit_manager, block_num desc`).Find(&data).Error
	log.CheckFatal(err)
	for _, entry := range data {
		repo.cmParams[entry.CreditManager] = entry
	}

	// fast check parameters
	fcparams := []*core.FastCheckParams{}
	err = repo.db.Raw(`SELECT distinct on (credit_manager) * FROM fast_check_params 
		ORDER BY credit_manager, block_num desc`).Find(&fcparams).Error
	log.CheckFatal(err)
	for _, entry := range fcparams {
		repo.cmFastCheckParams[entry.CreditManager] = entry
	}
}
