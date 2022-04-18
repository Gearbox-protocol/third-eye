package repository

import (
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

func (repo *Repository) GetCMState(cmAddr string) *schemas.CreditManagerState {
	state := repo.GetAdapter(cmAddr).GetUnderlyingState()
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

func (repo *Repository) AddCreditManagerStats(cms *schemas.CreditManagerStat) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.setAndGetBlock(cms.BlockNum).AddCreditManagerStats(cms)
}

func (repo *Repository) AddRepayOnCM(blockNum int64, cmAddr string, pnlOnRepay schemas.PnlOnRepay) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.setAndGetBlock(blockNum).AddRepayOnCM(cmAddr, &pnlOnRepay)
}

func (repo *Repository) GetRepayOnCM(blockNum int64, cmAddr string) *schemas.PnlOnRepay {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	return repo.blocks[blockNum].GetRepayOnCM(cmAddr)
}

func (repo *Repository) AddParameters(logID uint, txHash string, params *schemas.Parameters, token string) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.setAndGetBlock(params.BlockNum).AddParameters(params)
	// cal dao action
	oldCMParams := repo.cmParams[params.CreditManager]
	if oldCMParams == nil {
		oldCMParams = schemas.NewParameters()
	}
	args := oldCMParams.Diff(params)
	(*args)["token"] = token
	repo.addDAOOperation(&schemas.DAOOperation{
		BlockNumber: params.BlockNum,
		LogID:       logID,
		TxHash:      txHash,
		Type:        schemas.EventNewParameters,
		Contract:    params.CreditManager,
		Args:        args,
	})
	//
	repo.cmParams[params.CreditManager] = params
}

func (repo *Repository) paramsDAOV2(logID uint, txHash, creditConfigurator string, params *schemas.Parameters, fieldToRemove []string, daoEventType uint) {
	oldCMParams := repo.cmParams[params.CreditManager]
	if oldCMParams == nil {
		oldCMParams = schemas.NewParameters()
	}
	args := oldCMParams.Diff(params)
	for _, field := range fieldToRemove {
		delete(*args, field)
	}
	(*args)["creditManager"] = params.CreditManager
	repo.addDAOOperation(&schemas.DAOOperation{
		BlockNumber: params.BlockNum,
		LogID:       logID,
		TxHash:      txHash,
		Type:        daoEventType,
		Contract:    creditConfigurator,
		Args:        args,
	})
	//
	repo.cmParams[params.CreditManager] = params
}

func (repo *Repository) UpdateLimits(logID uint, txHash, creditConfigurator string, params *schemas.Parameters) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	// cal dao action
	oldCMParams := repo.cmParams[params.CreditManager]
	if oldCMParams == nil {
		oldCMParams = schemas.NewParameters()
	}
	repo.paramsDAOV2(logID, txHash, creditConfigurator, params,
		[]string{"feeLiquidation", "LiquidationDiscount", "feeInterest", "maxLeverage"}, schemas.LimitsUpdated)
	newParams := &schemas.Parameters{
		MinAmount:           params.MinAmount,
		MaxAmount:           params.MaxAmount,
		FeeInterest:         oldCMParams.FeeInterest,
		FeeLiquidation:      oldCMParams.FeeInterest,
		LiquidationDiscount: oldCMParams.LiquidationDiscount,
		BlockNum:            params.BlockNum,
		CreditManager:       params.CreditManager,
	}
	repo.setAndGetBlock(params.BlockNum).AddParameters(newParams)
	repo.cmParams[params.CreditManager] = newParams
}

func (repo *Repository) UpdateFees(logID uint, txHash, creditConfigurator string, params *schemas.Parameters) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	// cal dao action
	oldCMParams := repo.cmParams[params.CreditManager]
	if oldCMParams == nil {
		oldCMParams = schemas.NewParameters()
	}
	repo.paramsDAOV2(logID, txHash, creditConfigurator, params,
		[]string{"maxAmount", "maxLeverage", "minAmount"}, schemas.FeesUpdated)
	newParams := &schemas.Parameters{
		MinAmount:           oldCMParams.MinAmount,
		MaxAmount:           oldCMParams.MaxAmount,
		FeeInterest:         params.FeeInterest,
		FeeLiquidation:      params.FeeInterest,
		LiquidationDiscount: params.LiquidationDiscount,
		BlockNum:            params.BlockNum,
		CreditManager:       params.CreditManager,
	}
	repo.setAndGetBlock(params.BlockNum).AddParameters(newParams)
	repo.cmParams[params.CreditManager] = newParams
}

func (repo *Repository) loadAllParams() {
	defer utils.Elapsed("loadAllParams")()
	// parameters
	data := []*schemas.Parameters{}
	err := repo.db.Raw(`SELECT distinct on (credit_manager) * FROM parameters 
		ORDER BY credit_manager, block_num desc`).Find(&data).Error
	log.CheckFatal(err)
	for _, entry := range data {
		repo.cmParams[entry.CreditManager] = entry
	}

	// fast check parameters
	fcparams := []*schemas.FastCheckParams{}
	err = repo.db.Raw(`SELECT distinct on (credit_manager) * FROM fast_check_params 
		ORDER BY credit_manager, block_num desc`).Find(&fcparams).Error
	log.CheckFatal(err)
	for _, entry := range fcparams {
		repo.cmFastCheckParams[entry.CreditManager] = entry
	}
}

func (repo *Repository) AddAccountTokenTransfer(tt *schemas.TokenTransfer) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.setBlock(tt.BlockNum)
	repo.accountManager.AddTokenTransfer(tt)
}
