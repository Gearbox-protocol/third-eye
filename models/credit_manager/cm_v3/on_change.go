package cm_v3

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/pkg/dc"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/models/credit_manager/cm_v1"
	"github.com/ethereum/go-ethereum/core/types"
)

func bytesToUInt16(data []byte) uint16 {
	return uint16(new(big.Int).SetBytes(data).Int64())
}

func (mdl *CMv3) isExpired(blockNum int64) bool {
	if mdl.expirationDate == 0 {
		return false
	}
	return mdl.Repo.SetAndGetBlock(blockNum).Timestamp >= mdl.expirationDate
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
		case core.Topic("SetBorrowingLimits(uint256,uint256)"):
			minDebt := new(big.Int).SetBytes(txLog.Data[:32])
			maxDebt := new(big.Int).SetBytes(txLog.Data[32:])
			mdl.State.MinAmount = (*core.BigInt)(minDebt)
			mdl.State.MaxAmount = (*core.BigInt)(maxDebt)
		case core.Topic("SetExpirationDate(uint40)"):
			mdl.expirationDate = uint64(new(big.Int).SetBytes(txLog.Data[:]).Int64())
		case core.Topic("UpdateFees(uint16,uint16,uint16,uint16,uint16)"):
			// feeInterest,
			// uint16 feeLiquidation,
			// uint16 liquidationPremium,
			// uint16 feeLiquidationExpired,
			// uint16 liquidationPremiumExpired
			params := &schemas.Parameters{
				BlockNum:                   int64(txLog.BlockNumber),
				CreditManager:              mdl.Address,
				FeeInterest:                bytesToUInt16(txLog.Data[:32]),
				FeeLiquidation:             bytesToUInt16(txLog.Data[32:64]),
				LiquidationDiscount:        10000 - bytesToUInt16(txLog.Data[64:96]), // 10000- liqPremium
				FeeLiquidationExpired:      bytesToUInt16(txLog.Data[96:128]),
				LiquidationDiscountExpired: 10000 - bytesToUInt16(txLog.Data[128:160]), // 10000- liqPremiumExpired
			}
			mdl.SetParams(params)
			mdl.Repo.UpdateFees(txLog.Index, txLog.TxHash.Hex(), mdl.GetAddress(), params)
		}
		return
	}
	//
	mdl.CommonCMAdapter.PrefixOnLog(txLog)
	mdl.checkLogV3(txLog)
}

func (mdl *CMv3) SetOnChangeFn() {
	mdl.SetLastTxHashCompleted(mdl.lastTxHashCompleted)
	mdl.SetCalculateCMStatFn(func(blockNum int64, state dc.CMCallData) {
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
