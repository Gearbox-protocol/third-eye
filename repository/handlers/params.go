package handlers

import (
	"sync"

	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"gorm.io/gorm"
)

type ParamsRepo struct {
	// for params diff calculation
	cmParams          map[string]*schemas.Parameters
	cmFastCheckParams map[string]*schemas.FastCheckParams
	mu                *sync.Mutex
}

func NewParamsRepo() *ParamsRepo {
	return &ParamsRepo{
		// for dao events to get diff
		cmParams:          make(map[string]*schemas.Parameters),
		cmFastCheckParams: make(map[string]*schemas.FastCheckParams),
		mu:                &sync.Mutex{},
	}
}

func (repo *ParamsRepo) LoadAllParams(db *gorm.DB) {
	defer utils.Elapsed("loadAllParams")()
	// parameters
	data := []*schemas.Parameters{}
	err := db.Raw(`SELECT distinct on (credit_manager) * FROM parameters 
		ORDER BY credit_manager, block_num desc`).Find(&data).Error
	log.CheckFatal(err)
	for _, entry := range data {
		repo.cmParams[entry.CreditManager] = entry
	}

	// fast check parameters
	fcparams := []*schemas.FastCheckParams{}
	err = db.Raw(`SELECT distinct on (credit_manager) * FROM fast_check_params 
		ORDER BY credit_manager, block_num desc`).Find(&fcparams).Error
	log.CheckFatal(err)
	for _, entry := range fcparams {
		repo.cmFastCheckParams[entry.CreditManager] = entry
	}
}

func (repo *ParamsRepo) AddFastCheckParams(logID uint, txHash, cm, creditFilter string, fcParams *schemas.FastCheckParams) *schemas.DAOOperation {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	// set the dao action
	oldFCParams := repo.cmFastCheckParams[fcParams.CreditManager]
	if oldFCParams == nil {
		oldFCParams = schemas.NewFastCheckParams()
	}
	args := oldFCParams.Diff(fcParams)
	(*args)["creditManager"] = cm

	//
	repo.cmFastCheckParams[fcParams.CreditManager] = fcParams
	return &schemas.DAOOperation{
		BlockNumber: fcParams.BlockNum,
		LogID:       logID,
		TxHash:      txHash,
		Contract:    creditFilter,
		Args:        args,
		Type:        schemas.NewFastCheckParameters,
	}
}

// params on credit filter
func (repo *ParamsRepo) AddParameters(logID uint, txHash string, params *schemas.Parameters, token string) *schemas.DAOOperation {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	// repo.setAndGetBlock(params.BlockNum).AddParameters(params)
	// cal dao action
	oldCMParams := repo.cmParams[params.CreditManager]
	if oldCMParams == nil {
		oldCMParams = schemas.NewParameters()
	}
	args := oldCMParams.Diff(params)
	(*args)["token"] = token
	repo.cmParams[params.CreditManager] = params
	return &schemas.DAOOperation{
		BlockNumber: params.BlockNum,
		LogID:       logID,
		TxHash:      txHash,
		Type:        schemas.EventNewParameters,
		Contract:    params.CreditManager,
		Args:        args,
	}
}

// params on credit configurator
func (repo *ParamsRepo) paramsDAOV2(logID uint, txHash, creditConfigurator string, params *schemas.Parameters, fieldToRemove []string, daoEventType uint) *schemas.DAOOperation {
	oldCMParams := repo.cmParams[params.CreditManager]
	if oldCMParams == nil {
		oldCMParams = schemas.NewParameters()
	}
	args := oldCMParams.Diff(params)
	for _, field := range fieldToRemove {
		delete(*args, field)
	}
	(*args)["creditManager"] = params.CreditManager
	return &schemas.DAOOperation{
		BlockNumber: params.BlockNum,
		LogID:       logID,
		TxHash:      txHash,
		Type:        daoEventType,
		Contract:    creditConfigurator,
		Args:        args,
	}
}

func (repo *ParamsRepo) UpdateLimits(logID uint, txHash, creditConfigurator string, params *schemas.Parameters) (*schemas.Parameters, *schemas.DAOOperation) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	// cal dao action
	oldCMParams := repo.cmParams[params.CreditManager]
	if oldCMParams == nil {
		oldCMParams = schemas.NewParameters()
	}
	daoOperation := repo.paramsDAOV2(logID, txHash, creditConfigurator, params,
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
	repo.cmParams[params.CreditManager] = newParams
	return newParams, daoOperation
}

func (repo *ParamsRepo) UpdateFees(logID uint, txHash, creditConfigurator string, params *schemas.Parameters) (*schemas.Parameters, *schemas.DAOOperation) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	// cal dao action
	oldCMParams := repo.cmParams[params.CreditManager]
	if oldCMParams == nil {
		oldCMParams = schemas.NewParameters()
	}
	daoOperation := repo.paramsDAOV2(logID, txHash, creditConfigurator, params,
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
	repo.cmParams[params.CreditManager] = newParams
	return newParams, daoOperation
}
