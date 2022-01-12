package credit_filter

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
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
	case core.Topic("ContractAllowed(address,address)"):
		contractAllowedEvent, err := mdl.contractETH.ParseContractAllowed(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack contract allowed event", err)
		}
		mdl.Repo.AddAllowedProtocol(txLog.Index, txLog.TxHash.Hex(), mdl.Address, &core.Protocol{
			BlockNumber:   blockNum,
			CreditManager: creditManager,
			Protocol:      contractAllowedEvent.Protocol.Hex(),
			Adapter:       contractAllowedEvent.Adapter.Hex(),
		})
	case core.Topic("TokenAllowed(address,uint256)"):
		tokenEvent, err := mdl.contractETH.ParseTokenAllowed(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack token allowed event", err)
		}
		mdl.Repo.AddAllowedToken(txLog.Index, txLog.TxHash.Hex(), mdl.Address, &core.AllowedToken{
			BlockNumber:        blockNum,
			CreditManager:      creditManager,
			Token:              tokenEvent.Token.Hex(),
			LiquidityThreshold: (*core.BigInt)(tokenEvent.LiquidityThreshold),
		})
		mdl.Repo.AddToken(tokenEvent.Token.Hex())
	case core.Topic("TokenForbidden(address)"):
		token := common.HexToAddress(txLog.Topics[1].Hex())
		mdl.Repo.DisableAllowedToken(blockNum, txLog.Index, txLog.TxHash.Hex(), creditManager, mdl.Address, token.Hex())
	case core.Topic("NewFastCheckParameters(uint256,uint256)"):
		fcParams, err := mdl.contractETH.ParseNewFastCheckParameters(txLog)
		log.CheckFatal(err)
		mdl.Repo.AddFastCheckParams(txLog.Index, txLog.TxHash.Hex(), mdl.GetAddress(), &core.FastCheckParams{
			BlockNum:        blockNum,
			CreditManager:   creditManager,
			ChiThreshold:    (*core.BigInt)(fcParams.ChiThreshold),
			HFCheckInterval: (*core.BigInt)(fcParams.FastCheckDelay),
		})
	case core.Topic("ContractForbidden(address)"):
		contractDisabledEvent, err := mdl.contractETH.ParseContractForbidden(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack contract forbidden event", err)
		}
		mdl.Repo.DisableProtocol(blockNum, txLog.Index, txLog.TxHash.Hex(), creditManager, mdl.Address, contractDisabledEvent.Protocol.Hex())
	case core.Topic("TransferPluginAllowed(address,bool)"):
		transferPlugin, err := mdl.contractETH.ParseTransferPluginAllowed(txLog)
		log.CheckFatal(err)
		args := &core.Json{"plugin": transferPlugin.Pugin, "state": transferPlugin.State}
		mdl.Repo.AddDAOOperation(&core.DAOOperation{
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			BlockNumber: blockNum,
			Contract:    mdl.Address,
			Type:        core.TransferPluginAllowed,
			Args:        args,
		})
	case core.Topic("PriceOracleUpdated(address)"):
		po, err := mdl.contractETH.ParsePriceOracleUpdated(txLog)
		log.CheckFatal(err)
		args := &core.Json{"newPriceOracle": po.NewPriceOracle, "creditManager": creditManager}
		mdl.Repo.AddDAOOperation(&core.DAOOperation{
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			BlockNumber: blockNum,
			Contract:    mdl.Address,
			Type:        core.PriceOracleUpdated,
			Args:        args,
		})
	}
}
