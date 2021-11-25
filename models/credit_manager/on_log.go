package credit_manager

import (
	"github.com/Gearbox-protocol/gearscan/core"
	"github.com/ethereum/go-ethereum/core/types"

)

func (mdl *CreditManager) OnLog(txLog types.Log){
	blockNum := int64(txLog.BlockNumber)
	switch txLog.Topics[0] {
		case core.Topic("OpenCreditAccount(address,address,address,uint256,uint256,uint256)"):
			blockNum+=1
		case core.Topic("CloseCreditAccount(address,address,uint256)"):
			blockNum+=1
		case core.Topic("LiquidateCreditAccount(address,address,uint256)"):
			blockNum+=1
		case core.Topic("RepayCreditAccount(address,address)"):
			blockNum+=1
		case core.Topic("IncreaseBorrowedAmount(address,uint256)"):
			blockNum+=1
		case core.Topic("AddCollateral(address,address,uint256)"):
			blockNum+=1
		case core.Topic("ExecuteOrder(address,address)"):
			blockNum+=1
		case core.Topic("TransferAccount(address,address)"):
			blockNum+=1
	}
}