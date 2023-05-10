package credit_filter

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (mdl *CreditFilter) OnLogv2(txLog types.Log) {
	blockNum := int64(txLog.BlockNumber)
	creditManager := mdl.GetCM()
	switch txLog.Topics[0] {
	/////////////////////////////
	// credit configurator events
	/////////////////////////////
	case core.Topic("TokenAllowed(address)"):
		tokenEvent, err := mdl.cfgContract.ParseTokenAllowed(txLog)
		log.CheckFatal(err)
		mdl.Repo.AddAllowedTokenV2(txLog.Index, txLog.TxHash.Hex(), mdl.Address, &schemas.AllowedToken{
			BlockNumber:        blockNum,
			CreditManager:      creditManager,
			Token:              tokenEvent.Token.Hex(),
			LiquidityThreshold: nil,
			Configurator:       mdl.Address,
		})
	case core.Topic("TokenLiquidationThresholdUpdated(address,uint16)"):
		tokenEvent, err := mdl.cfgContract.ParseTokenLiquidationThresholdUpdated(txLog)
		log.CheckFatal(err)
		mdl.Repo.AddAllowedTokenV2(txLog.Index, txLog.TxHash.Hex(), mdl.Address, &schemas.AllowedToken{
			BlockNumber:        blockNum,
			CreditManager:      creditManager,
			Token:              tokenEvent.Token.Hex(),
			LiquidityThreshold: (*core.BigInt)(big.NewInt(int64(tokenEvent.LiquidityThreshold))),
			Configurator:       mdl.Address,
		})
	case core.Topic("LimitsUpdated(uint256,uint256)"):
		limitEvent, err := mdl.cfgContract.ParseLimitsUpdated(txLog)
		log.CheckFatal(err)
		mdl.Repo.UpdateLimits(txLog.Index, txLog.TxHash.Hex(), mdl.GetAddress(), &schemas.Parameters{
			BlockNum:      int64(txLog.BlockNumber),
			CreditManager: creditManager,
			MinAmount:     (*core.BigInt)(limitEvent.MinBorrowedAmount),
			MaxAmount:     (*core.BigInt)(limitEvent.MaxBorrowedAmount),
		})
	case core.Topic("FeesUpdated(uint16,uint16,uint16,uint16,uint16)"):
		feesEvent, err := mdl.cfgContract.ParseFeesUpdated(txLog)
		log.CheckFatal(err)
		params := &schemas.Parameters{
			BlockNum:                   blockNum,
			CreditManager:              creditManager,
			FeeInterest:                feesEvent.FeeInterest,
			FeeLiquidation:             feesEvent.FeeLiquidation,
			LiquidationDiscount:        10000 - feesEvent.LiquidationPremium,
			FeeLiquidationExpired:      feesEvent.FeeLiquidationExpired,
			LiquidationDiscountExpired: 10000 - feesEvent.LiquidationPremiumExpired,
		}
		mdl.Repo.UpdateFees(txLog.Index, txLog.TxHash.Hex(), mdl.GetAddress(), params)
	//
	// Previous fastcheck has some security issues, we change it for better security
	// case core.Topic("FastCheckParametersUpdated(uint256,uint256)"):
	// 	fcParams, err := mdl.cfgContract.ParseFastCheckParametersUpdated(txLog)
	// 	log.CheckFatal(err)
	// 	mdl.Repo.AddFastCheckParams(txLog.Index, txLog.TxHash.Hex(), creditManager, mdl.GetAddress(), &schemas.FastCheckParams{
	// 		BlockNum:        blockNum,
	// 		CreditManager:   creditManager,
	// 		ChiThreshold:    (*core.BigInt)(fcParams.ChiThreshold),
	// 		HFCheckInterval: (*core.BigInt)(fcParams.FastCheckDelay),
	// 	})
	//
	//
	// we are add dao event on here instead of in logs of creditmanager
	// as it might happen that this event is emitted before the first event on credit manager
	// in that case, it won't have added to db bcz we get logs from the first event blocknum on model
	// for credit manager that will lead to ignoring this dao event
	case core.Topic("CreditFacadeUpgraded(address)"):
		newFacade := common.BytesToAddress(txLog.Topics[1][:]).Hex()
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: int64(txLog.BlockNumber),
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    txLog.Address.Hex(),
			Args:        &core.Json{"facade": newFacade, "creditManager": creditManager},
			Type:        schemas.CreditFacadeUpgraded,
		})
	case core.Topic("AdapterForbidden(address)"):
		adapterForbidden := common.BytesToAddress(txLog.Topics[1][:]).Hex()
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: int64(txLog.BlockNumber),
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    txLog.Address.Hex(),
			Args: &core.Json{
				"adapter":       adapterForbidden,
				"creditManager": creditManager,
			},
			Type: schemas.AdapterForbidden,
		})
		// new events
	case core.Topic("PriceOracleUpgraded(address)"):
		priceOracle := common.BytesToAddress(txLog.Topics[1][:]).Hex()
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: int64(txLog.BlockNumber),
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    txLog.Address.Hex(),
			Args:        &core.Json{"newPriceOracle": priceOracle, "creditManager": creditManager},
			Type:        schemas.PriceOracleUpdated,
		})
	case core.Topic("IncreaseDebtForbiddenModeChanged(bool)"):
		increaseDebtForbiddenMode := new(big.Int).SetBytes(txLog.Data).Int64()
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: int64(txLog.BlockNumber),
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    txLog.Address.Hex(),
			Args: &core.Json{
				"creditManager":             creditManager,
				"increaseDebtForbiddenMode": increaseDebtForbiddenMode,
			},
			Type: schemas.IncreaseDebtForbiddenModeChanged,
		})
	case core.Topic("ExpirationDateUpdated(uint40)"):
		date := new(big.Int).SetBytes(txLog.Data).Int64()
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: int64(txLog.BlockNumber),
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    txLog.Address.Hex(),
			Args:        &core.Json{"creditManager": creditManager, "date": date},
			Type:        schemas.ExpirationDateUpdated,
		})
	case core.Topic("MaxEnabledTokensUpdated(uint8)"):
		maxEnabledTokens := new(big.Int).SetBytes(txLog.Data).Int64()
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: int64(txLog.BlockNumber),
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    txLog.Address.Hex(),
			Args:        &core.Json{"creditManager": creditManager, "maxEnabledTokens": maxEnabledTokens},
			Type:        schemas.MaxEnabledTokensUpdated,
		})
	case core.Topic("LimitPerBlockUpdated(uint128)"):
		limit := new(big.Int).SetBytes(txLog.Data).Int64()
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: int64(txLog.BlockNumber),
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    txLog.Address.Hex(),
			Args:        &core.Json{"creditManager": creditManager, "limitPerBlock": limit},
			Type:        schemas.LimitPerBlockUpdated,
		})
	case core.Topic("AddedToUpgradeable(address)"):
		contractAddr := common.BytesToAddress(txLog.Topics[1][:]).Hex()
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: int64(txLog.BlockNumber),
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    txLog.Address.Hex(),
			Args:        &core.Json{"creditManager": creditManager, "contract": contractAddr},
			Type:        schemas.AddedToUpgradeable,
		})
	case core.Topic("RemovedFromUpgradeable(address)"):
		contractAddr := common.BytesToAddress(txLog.Topics[1][:]).Hex()
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: int64(txLog.BlockNumber),
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    txLog.Address.Hex(),
			Args:        &core.Json{"creditManager": creditManager, "contract": contractAddr},
			Type:        schemas.RemovedFromUpgradeable,
		})
	case core.Topic("EmergencyLiquidatorAdded(address)"):
		emergencyLiquidator := common.BytesToAddress(txLog.Topics[1][:]).Hex()
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: int64(txLog.BlockNumber),
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    txLog.Address.Hex(),
			Args:        &core.Json{"creditManager": creditManager, "emergencyLiquidator": emergencyLiquidator},
			Type:        schemas.EmergencyLiquidatorAdded,
		})
	case core.Topic("EmergencyLiquidatorRemoved(address)"):
		emergencyLiquidator := common.BytesToAddress(txLog.Topics[1][:]).Hex()
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: int64(txLog.BlockNumber),
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    txLog.Address.Hex(),
			Args:        &core.Json{"creditManager": creditManager, "emergencyLiquidator": emergencyLiquidator},
			Type:        schemas.EmergencyLiquidatorRemoved,
		})
		// version 2_10
		// https://github.com/Gearbox-protocol/core-v2/commit/e5db57f447773d992b2505c344aa004a25b9e74e#diff-9a469bfc5a4690f063eb1df38a153a2764e81322e3ceed74815bfe3e121fcde1R795
	case core.Topic("NewMaxCumulativeLoss(uint256)"):
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: int64(txLog.BlockNumber),
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    txLog.Address.Hex(),
			Args:        &core.Json{"maxLoss": new(big.Int).SetBytes(txLog.Data[:]).String()},
			Type:        schemas.NewMaxCumulativeLoss,
		})
	case core.Topic("CumulativeLossReset()"):
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: int64(txLog.BlockNumber),
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    txLog.Address.Hex(),
			Args:        &core.Json{},
			Type:        schemas.CumulativeLossReset,
		})
	// https://github.com/Gearbox-protocol/core-v2/commit/6e22f2e66e50e42355aece9bca8dca25b8fc47cc#diff-9a469bfc5a4690f063eb1df38a153a2764e81322e3ceed74815bfe3e121fcde1R829
	case core.Topic("NewEmergencyLiquidationDiscount(uint16)"):
		liqDiscount := new(big.Int).SetBytes(txLog.Data).Int64()
		mdl.Repo.UpdateEmergencyLiqDiscount(txLog.Index, txLog.TxHash.Hex(), mdl.GetAddress(), &schemas.Parameters{
			BlockNum:             int64(txLog.BlockNumber),
			CreditManager:        creditManager,
			EmergencyLiqDiscount: uint16(liqDiscount),
		})
	}
}
