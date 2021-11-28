package credit_filter

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/Gearbox-protocol/gearscan/core"
	"github.com/Gearbox-protocol/gearscan/log"
	
)

func (mdl *CreditFilter) OnLog(txLog types.Log) {
	blockNum := int64(txLog.BlockNumber)
	switch txLog.Topics[0] {
	case core.Topic("ContractAllowed(address,address)"):
		contractAllowedEvent, err := mdl.contractETH.ParseContractAllowed(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack contract allowed event", err)
		}
		mdl.Repo.AddAllowedProtocol(&core.Protocol{
			BlockNumber: blockNum, 
			CreditManager: mdl.Details["creditManager"],
			Protocol: contractAllowedEvent.Protocol.Hex(),
			Adapter: contractAllowedEvent.Adapter.Hex(),
		})
	case core.Topic("TokenAllowed(address,uint256)"):
		tokenEvent, err := mdl.contractETH.ParseTokenAllowed(txLog)
		if err != nil {
			log.Fatal("[CreditManagerModel]: Cant unpack token allowed event", err)
		}
		mdl.Repo.AddAllowedToken(&core.AllowedToken{
			CreditManager: mdl.Details["creditManager"],
			Token: tokenEvent.Token.Hex(),
			LiquidityThreshold: tokenEvent.LiquidityThreshold.String(),
		})
		mdl.Repo.AddToken(tokenEvent.Token.Hex())
	}
}
