package credit_filter

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (mdl *CreditFilter) GetCM() string {
	creditManager, ok := mdl.Details["creditManager"].(string)
	if !ok {
		log.Fatal("Failed in asserting credit manager(%v) for credit filter %s", mdl.Details["creditManager"], mdl.GetAddress())
	}
	return creditManager
}
func (mdl *CreditFilter) OnLog(txLog types.Log) {
	blockNum := int64(txLog.BlockNumber)
	creditManager := mdl.GetCM()
	switch txLog.Topics[0] {
	// common
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
	}
	mdl.OnLogv2(txLog)
}
