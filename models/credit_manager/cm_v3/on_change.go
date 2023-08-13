package cm_v3

import (
	"math/big"

	dcv2 "github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressor/dataCompressorv2"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/models/credit_manager/cm_v1"
	"github.com/ethereum/go-ethereum/core/types"
)

func bytesToUInt16(data []byte) uint16 {
	return uint16(new(big.Int).SetBytes(data).Int64())
}

func (mdl *CMv3) OnLog(txLog types.Log) {
	// creditConfigurator events for test
	// CreditFacadeUpgraded is emitted when creditconfigurator is initialized, so we will receive it on init
	// although we have already set creditfacadeUpgra
	if mdl.GetDetailsByKey("configurator") == txLog.Address.Hex() {
		switch txLog.Topics[0] {
		case core.Topic("SetCreditFacade(address)"):
			facade := utils.ChecksumAddr(txLog.Topics[1].Hex())
			mdl.setCreditFacadeSyncer(facade)
		case core.Topic("UpdateFees(uint16,uint16,uint16,uint16,uint16)"):
			mdl.SetParams(&schemas.Parameters{
				BlockNum:                   int64(txLog.BlockNumber),
				CreditManager:              mdl.Address,
				FeeInterest:                bytesToUInt16(txLog.Data[:32]),
				FeeLiquidation:             bytesToUInt16(txLog.Data[32:64]),
				LiquidationDiscount:        10000 - bytesToUInt16(txLog.Data[64:96]), // 10000- liqPremium
				FeeLiquidationExpired:      bytesToUInt16(txLog.Data[96:128]),
				LiquidationDiscountExpired: 10000 - bytesToUInt16(txLog.Data[128:160]), // 10000- liqPremiumExpired
			})
		}
		return
	}
	//
	mdl.checkLogV3(txLog)
}

func (mdl *CMv3) SetOnChangeFn() {
	mdl.SetLastTxHashCompleted(mdl.lastTxHashCompleted)
	mdl.SetCalculateCMStatFn(func(blockNum int64, state dcv2.CreditManagerData) {
		// mdl.addProtocolAdapters(state)
		mdl.CalculateCMStat(blockNum, state)
	})
	mdl.SetOnDirectTokenTransferFn(cm_v1.OnDirectTokenTransfer)
}

// /////////
// On TxHash
// /////////
func (mdl *CMv3) lastTxHashCompleted(lastTxHash string) {
	nonMulticallExecuteEvents := mdl.ProcessNonMultiCalls()
	mdl.ProcessRemainingMultiCalls(mdl.GetVersion(), lastTxHash, nonMulticallExecuteEvents)
}
