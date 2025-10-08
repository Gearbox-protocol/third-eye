package configurator_v3

import (
	"math/big"
	"reflect"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/core/schemas/schemas_v3"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// TODO:  QuoteToken, SetBorrowingLimits, SetBotList
// SetTotalDebtLimit, SetMaxDebtPerBlockMultiplier for v3 -- NewTotalDebtLimit, LimitPerBlockUpdated for v2
func (mdl *Configuratorv3) OnLog(txLog types.Log) {
	blockNum := int64(txLog.BlockNumber)
	creditManager := mdl.GetCM()
	switch txLog.Topics[0] {
	// common
	case core.Topic("AllowAdapter(address,address)"):
		contractAllowedEvent, err := mdl.cfgContract.ParseAllowAdapter(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack contract allowed event", err)
		}
		mdl.Repo.AddAllowedProtocol(txLog.Index, txLog.TxHash.Hex(), mdl.Address, &schemas.Protocol{
			BlockNumber:   blockNum,
			CreditManager: creditManager,
			Protocol:      contractAllowedEvent.TargetContract.Hex(),
			Adapter:       contractAllowedEvent.Adapter.Hex(),
			Configurator:  mdl.Address,
		})
	case core.Topic("ForbidToken(address)"):
		token := common.HexToAddress(txLog.Topics[1].Hex())
		mdl.Repo.DisableAllowedToken(blockNum, txLog.Index, txLog.TxHash.Hex(), creditManager, mdl.Address, token.Hex())
	case core.Topic("ForbidAdapter(address)"):
		contractDisabledEvent, err := mdl.cfgContract.ParseForbidAdapter(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack contract forbidden event", err)
		}
		mdl.Repo.DisableProtocol(blockNum, txLog.Index, txLog.TxHash.Hex(), creditManager, mdl.Address, contractDisabledEvent.TargetContract.Hex())
	/////////////////////////////
	// credit configurator events
	/////////////////////////////
	case core.Topic("AllowToken(address)"):
		// tokenEvent, err := mdl.cfgContract.ParseAllowToken(txLog)
		// log.CheckFatal(err)
		// mdl.Repo.AddAllowedTokenV2(txLog.Index, txLog.TxHash.Hex(), mdl.Address, &schemas.AllowedToken{
		// 	BlockNumber:        blockNum,
		// 	CreditManager:      creditManager,
		// 	Token:              tokenEvent.Token.Hex(),
		// 	LiquidityThreshold: nil,
		// 	Configurator:       mdl.Address,
		// })
	case core.Topic("SetTokenLiquidationThreshold(address,uint16)"):
		tokenEvent, err := mdl.cfgContract.ParseSetTokenLiquidationThreshold(txLog)
		log.CheckFatal(err)
		mdl.Repo.AddAllowedTokenV2(txLog.Index, txLog.TxHash.Hex(), mdl.Address, &schemas.AllowedToken{
			BlockNumber:        blockNum,
			CreditManager:      creditManager,
			LogID:              txLog.Index,
			Token:              tokenEvent.Token.Hex(),
			LiquidityThreshold: (*core.BigInt)(big.NewInt(int64(tokenEvent.LiquidationThreshold))),
			Configurator:       mdl.Address,
		})
	case core.Topic("UpdateFees(uint16,uint16,uint16,uint16,uint16)"):
		feesEvent, err := mdl.cfgContract.ParseUpdateFees(txLog)
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
	case core.Topic("UpdateFees(uint16,uint16,uint16,uint16)"):
		// feesEvent, err := mdl.cfgContractv310.ParseUpdateFees(txLog)
		// log.CheckFatal(err)
		cmabi := core.GetAbi("CreditManagerv3")
		data, err := core.CallFuncGetAllData(mdl.Client, "9af1d35a", common.HexToAddress(mdl.GetCM()), blockNum, nil)
		log.CheckFatal(err)
		values, err := cmabi.Unpack("fees", data)
		log.CheckFatal(err)

		var feesEvent *feesDS
		log.Info(reflect.TypeOf(values[0]))
		_feesEvent := getF(values)
		if _feesEvent == nil {
			feesEvent = &feesDS{
				FeeInterest:                uint16(values[0].(uint16)),
				FeeLiquidation:             uint16(values[1].(uint16)),
				LiquidationDiscount:        uint16(values[2].(uint16)),
				FeeLiquidationExpired:      uint16(values[3].(uint16)),
				LiquidationDiscountExpired: uint16(values[4].(uint16)),
			}
		}
		params := &schemas.Parameters{
			BlockNum:                   blockNum,
			CreditManager:              creditManager,
			FeeInterest:                feesEvent.FeeInterest,
			FeeLiquidation:             feesEvent.FeeLiquidation,
			LiquidationDiscount:        feesEvent.LiquidationDiscount,
			FeeLiquidationExpired:      feesEvent.FeeLiquidationExpired,
			LiquidationDiscountExpired: feesEvent.LiquidationDiscountExpired,
		}
		mdl.Repo.UpdateFees(txLog.Index, txLog.TxHash.Hex(), mdl.GetAddress(), params)
	case core.Topic("SetCreditFacade(address)"):
		newFacade := common.BytesToAddress(txLog.Topics[1][:]).Hex()
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: blockNum,
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    txLog.Address.Hex(),
			Args:        &core.Json{"facade": newFacade, "creditManager": creditManager},
			Type:        schemas.CreditFacadeUpgraded,
		})
	case core.Topic("SetPriceOracle(address)"):
		priceOracle := common.BytesToAddress(txLog.Topics[1][:]).Hex()
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: blockNum,
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    txLog.Address.Hex(),
			Args:        &core.Json{"newPriceOracle": priceOracle, "creditManager": creditManager},
			Type:        schemas.PriceOracleUpdated,
		})
	case core.Topic("SetExpirationDate(uint40)"): // ExpirationDateUpdated(uint40)
		date := new(big.Int).SetBytes(txLog.Data).Int64()
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: int64(txLog.BlockNumber),
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    txLog.Address.Hex(),
			Args:        &core.Json{"creditManager": creditManager, "date": date},
			Type:        schemas.ExpirationDateUpdated,
		})
	case core.Topic("SetMaxEnabledTokens(uint8)"): // MaxEnabledTokensUpdated(uint8)
		maxEnabledTokens := new(big.Int).SetBytes(txLog.Data).Int64()
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: int64(txLog.BlockNumber),
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    txLog.Address.Hex(),
			Args:        &core.Json{"creditManager": creditManager, "maxEnabledTokens": maxEnabledTokens},
			Type:        schemas.MaxEnabledTokensUpdated,
		})
	case core.Topic("SetMaxDebtPerBlockMultiplier(uint128)"): // LimitPerBlockUpdated(uint128)
		limit := new(big.Int).SetBytes(txLog.Data).Int64()
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: int64(txLog.BlockNumber),
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    txLog.Address.Hex(),
			Args:        &core.Json{"creditManager": creditManager, "limitPerBlock": limit},
			Type:        schemas.LimitPerBlockUpdated,
		})
	case core.Topic("SetTotalDebtLimit(uint128)"):
		limit := new(big.Int).SetBytes(txLog.Data).Int64()
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: int64(txLog.BlockNumber),
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    txLog.Address.Hex(),
			Args:        &core.Json{"creditManager": creditManager, "totalDebtLimit": limit},
			Type:        schemas.NewTotalDebtLimit,
		})
	case core.Topic("AddEmergencyLiquidator(address)"):
		emergencyLiquidator := common.BytesToAddress(txLog.Topics[1][:]).Hex()
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: int64(txLog.BlockNumber),
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    txLog.Address.Hex(),
			Args:        &core.Json{"creditManager": creditManager, "emergencyLiquidator": emergencyLiquidator},
			Type:        schemas.EmergencyLiquidatorAdded,
		})
	case core.Topic("RemoveEmergencyLiquidator(address)"):
		emergencyLiquidator := common.BytesToAddress(txLog.Topics[1][:]).Hex()
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: int64(txLog.BlockNumber),
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    txLog.Address.Hex(),
			Args:        &core.Json{"creditManager": creditManager, "emergencyLiquidator": emergencyLiquidator},
			Type:        schemas.EmergencyLiquidatorRemoved,
		})
		// v_210 compartable code
	case core.Topic("SetMaxCumulativeLoss(uint256)"):
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: int64(txLog.BlockNumber),
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    txLog.Address.Hex(),
			Args:        &core.Json{"maxLoss": new(big.Int).SetBytes(txLog.Data[:]).String()},
			Type:        schemas.NewMaxCumulativeLoss,
		})
	case core.Topic("ResetCumulativeLoss()"):
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: int64(txLog.BlockNumber),
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    txLog.Address.Hex(),
			Args:        &core.Json{},
			Type:        schemas.CumulativeLossReset,
		})
	// extra v3 eents

	case core.Topic("ScheduleTokenLiquidationThresholdRamp(address,uint16,uint16,uint40,uint40)"):
		rampDetails, err := mdl.cfgContract.ParseScheduleTokenLiquidationThresholdRamp(txLog)
		log.CheckFatal(err)
		mdl.Repo.AddTokenLTRamp(
			&schemas_v3.TokenLTRamp{
				BlockNum:      blockNum,
				LogID:         txLog.Index,
				CreditManager: mdl.GetCM(),
				Token:         rampDetails.Token.Hex(),
				LtInitial:     rampDetails.LiquidationThresholdInitial,
				LtFinal:       rampDetails.LiquidationThresholdFinal,
				RampStart:     uint64(rampDetails.TimestampRampStart.Int64()),
				RampEnd:       uint64(rampDetails.TimestampRampEnd.Int64()),
			},
		)
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: int64(txLog.BlockNumber),
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    txLog.Address.Hex(),
			Args: &core.Json{
				"creditManager": mdl.GetCM(),
				"token":         rampDetails.Token.Hex(),
				"lt_initial":    rampDetails.LiquidationThresholdInitial,
				"lt_final":      rampDetails.LiquidationThresholdFinal,
				"ramp_start":    rampDetails.TimestampRampStart.Int64(),
				"ramp_end":      rampDetails.TimestampRampEnd.Int64(),
			},
			Type: schemas.SetTokenLiquidationThresholdRamp,
		})
	case core.Topic("SetBorrowingLimits(uint128,uint128)"):
		limitEvent, err := mdl.cfgContract.ParseSetBorrowingLimits(txLog)
		log.CheckFatal(err)
		mdl.Repo.UpdateLimits(txLog.Index, txLog.TxHash.Hex(), mdl.GetAddress(), &schemas.Parameters{
			BlockNum:      int64(txLog.BlockNumber),
			CreditManager: creditManager,
			MinAmount:     (*core.BigInt)(limitEvent.MinDebt),
			MaxAmount:     (*core.BigInt)(limitEvent.MaxDebt),
		})
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: int64(txLog.BlockNumber),
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    txLog.Address.Hex(),
			Args: &core.Json{
				"creditManager": mdl.GetCM(),
				"minDebt":       (*core.BigInt)(new(big.Int).SetBytes(txLog.Data[:32])),
				"maxDebt":       (*core.BigInt)(new(big.Int).SetBytes(txLog.Data[32:])),
			},
			Type: schemas.CumulativeLossReset,
		})
	case core.Topic("SetBotList(address)"):
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: int64(txLog.BlockNumber),
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    txLog.Address.Hex(),
			Args: &core.Json{
				"creditManager": mdl.GetCM(),
				"botList":       common.BytesToAddress(txLog.Data[:]).Hex(),
			},
			Type: schemas.CumulativeLossReset,
		})
	}
}

// select block_num, token , credit_manager , count(*) from (select * from (select block_num , credit_manager, token  from allowed_tokens ) a union all (select block_num, credit_manager, token from token_ltramp)) b group by b.block_num, b.token, b.credit_manager having count(*)>1;

type feesDS struct {
	FeeInterest                uint16
	FeeLiquidation             uint16
	LiquidationDiscount        uint16
	FeeLiquidationExpired      uint16
	LiquidationDiscountExpired uint16
}

func getF(values []interface{}) *feesDS {
	defer func() {
		if err := recover(); err != nil {
			log.Info(err)
			// return nil
		}
	}()
	return abi.ConvertType(values, new(feesDS)).(*feesDS)
}
