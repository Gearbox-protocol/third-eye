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
	blocks            *BlocksRepo
}

func NewParamsRepo(blocks *BlocksRepo) *ParamsRepo {
	return &ParamsRepo{
		// for dao events to get diff
		cmParams:          make(map[string]*schemas.Parameters),
		cmFastCheckParams: make(map[string]*schemas.FastCheckParams),
		mu:                &sync.Mutex{},
		blocks:            blocks,
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

func (repo *ParamsRepo) AddFastCheckParams(logID uint, txHash, cm, creditFilter string, fcParams *schemas.FastCheckParams) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.blocks.SetAndGetBlock(fcParams.BlockNum).AddFastCheckParams(fcParams)
	// set the dao action
	oldFCParams := repo.cmFastCheckParams[fcParams.CreditManager]
	if oldFCParams == nil {
		oldFCParams = schemas.NewFastCheckParams()
	}
	args := oldFCParams.Diff(fcParams)
	(*args)["creditManager"] = cm

	//
	repo.cmFastCheckParams[fcParams.CreditManager] = fcParams
	repo.blocks.AddDAOOperation(&schemas.DAOOperation{
		BlockNumber: fcParams.BlockNum,
		LogID:       logID,
		TxHash:      txHash,
		Contract:    creditFilter,
		Args:        args,
		Type:        schemas.NewFastCheckParameters,
	})
}

// params on credit filter
func (repo *ParamsRepo) AddParameters(logID uint, txHash string, params *schemas.Parameters, token string) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.blocks.SetAndGetBlock(params.BlockNum).AddParameters(params)
	// cal dao action
	oldCMParams := repo.cmParams[params.CreditManager]
	if oldCMParams == nil {
		oldCMParams = schemas.NewParameters()
	}
	args := oldCMParams.Diffv1(params)
	delete(*args, "feeLiquidationExpired")
	delete(*args, "liquidationDiscountExpired")
	//
	(*args)["token"] = token
	repo.cmParams[params.CreditManager] = params
	repo.blocks.AddDAOOperation(&schemas.DAOOperation{
		BlockNumber: params.BlockNum,
		LogID:       logID,
		TxHash:      txHash,
		Type:        schemas.EventNewParameters,
		Contract:    params.CreditManager,
		Args:        args,
	})
}

// for creating DAO event entry
func (repo *ParamsRepo) paramsDAOV2(logID uint, txHash, creditConfigurator string, params *schemas.Parameters, fieldsToKeep []string, daoEventType uint) {
	oldCMParams := repo.cmParams[params.CreditManager]
	if oldCMParams == nil {
		oldCMParams = schemas.NewParameters()
	}
	args := oldCMParams.Diffv2(params)
	for _, field := range fieldsToKeep {
		if (*args)[field] == nil { // check if misspelled field name is not entered
			log.Fatal("Wrong parameter field name", field)
		}
		if !utils.Contains(fieldsToKeep, field) {
			delete(*args, field)
		}
	}
	(*args)["creditManager"] = params.CreditManager
	repo.blocks.AddDAOOperation(&schemas.DAOOperation{
		BlockNumber: params.BlockNum,
		LogID:       logID,
		TxHash:      txHash,
		Type:        daoEventType,
		Contract:    creditConfigurator,
		Args:        args,
	})
}

func (repo *ParamsRepo) UpdateLimits(logID uint, txHash, creditConfigurator string, params *schemas.Parameters) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	// cal dao action
	oldCMParams := repo.cmParams[params.CreditManager]
	if oldCMParams == nil {
		oldCMParams = schemas.NewParameters()
	}
	//
	repo.paramsDAOV2(logID, txHash, creditConfigurator, params,
		[]string{"minAmount", "maxAmount"}, schemas.LimitsUpdated)
	newParams := oldCMParams
	newParams.MinAmount = params.MinAmount
	newParams.MaxAmount = params.MaxAmount
	newParams.BlockNum = params.BlockNum
	newParams.CreditManager = params.CreditManager
	//
	repo.cmParams[params.CreditManager] = newParams
	repo.blocks.SetAndGetBlock(params.BlockNum).AddParameters(newParams)
}

func (repo *ParamsRepo) UpdateEmergencyLiqPremium(logID uint, txHash, creditConfigurator string, params *schemas.Parameters) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	// cal dao action
	oldCMParams := repo.cmParams[params.CreditManager]
	if oldCMParams == nil {
		oldCMParams = schemas.NewParameters()
	}
	//
	repo.paramsDAOV2(logID, txHash, creditConfigurator, params,
		[]string{"emergencyLiqDiscount"}, schemas.NewEmergencyLiquidationPremium)
	newParams := oldCMParams
	newParams.EmergencyLiqDiscount = params.EmergencyLiqDiscount
	newParams.BlockNum = params.BlockNum
	newParams.CreditManager = params.CreditManager
	//
	repo.cmParams[params.CreditManager] = newParams
	repo.blocks.SetAndGetBlock(params.BlockNum).AddParameters(newParams)
}

func (repo *ParamsRepo) UpdateFees(logID uint, txHash, creditConfigurator string, params *schemas.Parameters) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	// cal dao action
	oldCMParams := repo.cmParams[params.CreditManager]
	if oldCMParams == nil {
		oldCMParams = schemas.NewParameters()
	}
	repo.paramsDAOV2(logID, txHash, creditConfigurator, params,
		[]string{
			"feeInterest",
			"feeInterest",
			"liquidationDiscount",
			"liquidationDiscountExpired",
			"feeLiquidationExpired",
		}, schemas.FeesUpdated)
	//
	newParams := oldCMParams
	newParams.FeeInterest = params.FeeInterest
	newParams.FeeLiquidation = params.FeeLiquidation
	newParams.LiquidationDiscount = params.LiquidationDiscount
	newParams.LiquidationDiscountExpired = params.LiquidationDiscountExpired
	newParams.FeeLiquidationExpired = params.FeeLiquidationExpired
	newParams.BlockNum = params.BlockNum
	newParams.CreditManager = params.CreditManager
	//
	repo.cmParams[params.CreditManager] = newParams
	repo.blocks.SetAndGetBlock(params.BlockNum).AddParameters(newParams)
}
