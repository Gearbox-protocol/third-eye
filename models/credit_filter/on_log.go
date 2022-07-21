package credit_filter

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (mdl *CreditFilter) OnLog(txLog types.Log) {
	blockNum := int64(txLog.BlockNumber)
	creditManager, ok := mdl.Details["creditManager"].(string)
	if !ok {
		log.Fatal("Failed in asserting credit manager(%v) for credit filter %s", mdl.Details["creditManager"], mdl.GetAddress())
	}
	switch txLog.Topics[0] {
	// common
	case core.Topic("PriceOracleUpdated(address)"):
		po, err := mdl.filterContract.ParsePriceOracleUpdated(txLog)
		log.CheckFatal(err)
		args := &core.Json{"newPriceOracle": po.NewPriceOracle.Hex(), "creditManager": creditManager}
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			BlockNumber: blockNum,
			Contract:    mdl.Address,
			Type:        schemas.PriceOracleUpdated,
			Args:        args,
		})
	case core.Topic("ContractAllowed(address,address)"):
		contractAllowedEvent, err := mdl.filterContract.ParseContractAllowed(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack contract allowed event", err)
		}
		mdl.Repo.AddAllowedProtocol(txLog.Index, txLog.TxHash.Hex(), mdl.Address, &schemas.Protocol{
			BlockNumber:   blockNum,
			CreditManager: creditManager,
			Protocol:      contractAllowedEvent.Protocol.Hex(),
			Adapter:       contractAllowedEvent.Adapter.Hex(),
		})
	case core.Topic("TokenForbidden(address)"):
		token := common.HexToAddress(txLog.Topics[1].Hex())
		mdl.Repo.DisableAllowedToken(blockNum, txLog.Index, txLog.TxHash.Hex(), creditManager, mdl.Address, token.Hex())
	case core.Topic("ContractForbidden(address)"):
		contractDisabledEvent, err := mdl.filterContract.ParseContractForbidden(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack contract forbidden event", err)
		}
		mdl.Repo.DisableProtocol(blockNum, txLog.Index, txLog.TxHash.Hex(), creditManager, mdl.Address, contractDisabledEvent.Protocol.Hex())
	case core.Topic("NewFastCheckParameters(uint256,uint256)"):
		fcParams, err := mdl.filterContract.ParseNewFastCheckParameters(txLog)
		log.CheckFatal(err)
		mdl.Repo.AddFastCheckParams(txLog.Index, txLog.TxHash.Hex(), creditManager, mdl.GetAddress(), &schemas.FastCheckParams{
			BlockNum:        blockNum,
			CreditManager:   creditManager,
			ChiThreshold:    (*core.BigInt)(fcParams.ChiThreshold),
			HFCheckInterval: (*core.BigInt)(fcParams.FastCheckDelay),
		})
	////////////////////////
	// credit filter events
	////////////////////////
	case core.Topic("TokenAllowed(address,uint256)"):
		tokenEvent, err := mdl.filterContract.ParseTokenAllowed(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack token allowed event", err)
		}
		mdl.Repo.AddAllowedToken(txLog.Index, txLog.TxHash.Hex(), mdl.Address, &schemas.AllowedToken{
			BlockNumber:        blockNum,
			CreditManager:      creditManager,
			Token:              tokenEvent.Token.Hex(),
			LiquidityThreshold: (*core.BigInt)(tokenEvent.LiquidityThreshold),
		})
		mdl.Repo.GetToken(tokenEvent.Token.Hex())
	case core.Topic("TransferPluginAllowed(address,bool)"):
		transferPlugin, err := mdl.filterContract.ParseTransferPluginAllowed(txLog)
		log.CheckFatal(err)
		args := &core.Json{"plugin": transferPlugin.Pugin.Hex(), "state": transferPlugin.State, "creditManager": creditManager}
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			BlockNumber: blockNum,
			Contract:    mdl.Address,
			Type:        schemas.TransferPluginAllowed,
			Args:        args,
		})
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
		})
	case core.Topic("TokenLiquidationThresholdUpdated(address,uint16)"):
		tokenEvent, err := mdl.cfgContract.ParseTokenLiquidationThresholdUpdated(txLog)
		log.CheckFatal(err)
		mdl.Repo.AddAllowedTokenV2(txLog.Index, txLog.TxHash.Hex(), mdl.Address, &schemas.AllowedToken{
			BlockNumber:        blockNum,
			CreditManager:      creditManager,
			Token:              tokenEvent.Token.Hex(),
			LiquidityThreshold: (*core.BigInt)(big.NewInt(int64(tokenEvent.LiquidityThreshold))),
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
	case core.Topic("FeesUpdated(uint16,uint16,uint16)"):
		feesEvent, err := mdl.cfgContract.ParseFeesUpdated(txLog)
		log.CheckFatal(err)
		if mdl.underlyingToken != nil {
			mdl.Repo.AddAllowedTokenV2(txLog.Index, txLog.TxHash.Hex(), mdl.Address, &schemas.AllowedToken{
				BlockNumber:        blockNum,
				CreditManager:      creditManager,
				Token:              mdl.underlyingToken.Hex(),
				LiquidityThreshold: (*core.BigInt)(big.NewInt(int64(feesEvent.LiquidationPremium - feesEvent.FeeLiquidation))),
			})
		}
		mdl.Repo.UpdateFees(txLog.Index, txLog.TxHash.Hex(), mdl.GetAddress(), &schemas.Parameters{
			BlockNum:            int64(txLog.BlockNumber),
			CreditManager:       creditManager,
			FeeInterest:         (*core.BigInt)(big.NewInt(int64(feesEvent.FeeInterest))),
			FeeLiquidation:      (*core.BigInt)(big.NewInt(int64(feesEvent.FeeLiquidation))),
			LiquidationDiscount: (*core.BigInt)(big.NewInt(int64(feesEvent.LiquidationPremium))),
		})
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
		newFacade := utils.ChecksumAddr(txLog.Topics[1].Hex())
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: int64(txLog.BlockNumber),
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    txLog.Address.Hex(),
			Args:        &core.Json{"facade": newFacade, "creditManager": creditManager},
			Type:        schemas.CreditFacadeUpgraded,
		})
	}
}
