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
		adapter := repo.GetAdapter(cm.Address)
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
		adapter := repo.GetAdapter(cm)
		if adapter != nil && adapter.GetName() == "CreditManager" {
			adapter.SetUnderlyingState(hstore)
		}
	}
}

func (repo *Repository) GetCMState(cmAddr string) *core.CreditManagerState {
	state := repo.GetAdapter(cmAddr).GetUnderlyingState()
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

func (repo *Repository) AddParameters(logID uint, txHash string, params *core.Parameters, token string) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.setAndGetBlock(params.BlockNum).AddParameters(params)
	// cal dao action
	oldCMParams := repo.cmParams[params.CreditManager]
	if oldCMParams == nil {
		oldCMParams = core.NewParameters()
	}
	args := oldCMParams.Diff(params)
	(*args)["token"] = token
	repo.addDAOOperation(&core.DAOOperation{
		BlockNumber: params.BlockNum,
		LogID:       logID,
		TxHash:      txHash,
		Type:        core.EventNewParameters,
		Contract:    params.CreditManager,
		Args:        args,
	})
	//
	repo.cmParams[params.CreditManager] = params
}

func (repo *Repository) paramsDAOV2(logID uint, txHash string, params *core.Parameters, fieldToRemove []string, daoEventType uint) {
	oldCMParams := repo.cmParams[params.CreditManager]
	if oldCMParams == nil {
		oldCMParams = core.NewParameters()
	}
	args := oldCMParams.Diff(params)
	for _, field := range fieldToRemove {
		delete(*args, field)
	}
	repo.addDAOOperation(&core.DAOOperation{
		BlockNumber: params.BlockNum,
		LogID:       logID,
		TxHash:      txHash,
		Type:        daoEventType,
		Contract:    params.CreditManager,
		Args:        args,
	})
	//
	repo.cmParams[params.CreditManager] = params
}

func (repo *Repository) UpdateLimits(logID uint, txHash string, params *core.Parameters) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	// cal dao action
	oldCMParams := repo.cmParams[params.CreditManager]
	if oldCMParams == nil {
		oldCMParams = core.NewParameters()
	}
	repo.paramsDAOV2(logID, txHash, params,
		[]string{"feeLiquidation", "LiquidationDiscount", "feeInterest", "maxLeverage"}, core.LimitsUpdated)
	newParams := &core.Parameters{
		MinAmount:           params.MinAmount,
		MaxAmount:           params.MaxAmount,
		FeeInterest:         oldCMParams.FeeInterest,
		FeeLiquidation:      oldCMParams.FeeInterest,
		LiquidationDiscount: oldCMParams.LiquidationDiscount,
	}
	repo.setAndGetBlock(params.BlockNum).AddParameters(newParams)
	repo.cmParams[params.CreditManager] = newParams
}

func (repo *Repository) UpdateFees(logID uint, txHash string, params *core.Parameters) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	// cal dao action
	oldCMParams := repo.cmParams[params.CreditManager]
	if oldCMParams == nil {
		oldCMParams = core.NewParameters()
	}
	repo.paramsDAOV2(logID, txHash, params,
		[]string{"maxAmount", "maxLeverage", "minAmount"}, core.FeesUpdated)
	newParams := &core.Parameters{
		MinAmount:           oldCMParams.MinAmount,
		MaxAmount:           oldCMParams.MaxAmount,
		FeeInterest:         params.FeeInterest,
		FeeLiquidation:      params.FeeInterest,
		LiquidationDiscount: params.LiquidationDiscount,
	}
	repo.setAndGetBlock(params.BlockNum).AddParameters(newParams)
	repo.cmParams[params.CreditManager] = newParams
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

func (repo *Repository) AddAccountTokenTransfer(tt *core.TokenTransfer) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.setBlock(tt.BlockNum)
	repo.accountManager.AddTokenTransfer(tt)
}
