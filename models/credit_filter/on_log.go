package credit_filter

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
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
		mdl.Repo.AddAllowedProtocol(&core.Protocol{
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
		mdl.Repo.AddAllowedToken(&core.AllowedToken{
			BlockNumber:        int64(txLog.BlockNumber),
			CreditManager:      creditManager,
			Token:              tokenEvent.Token.Hex(),
			LiquidityThreshold: (*core.BigInt)(tokenEvent.LiquidityThreshold),
		})
		mdl.Repo.AddToken(tokenEvent.Token.Hex())
	}
}
