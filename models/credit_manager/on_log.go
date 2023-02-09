package credit_manager

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/core/types"
)

func (mdl *CreditManager) processExecuteEvents() {
	if len(mdl.executeParams) > 0 {
		mdl.handleExecuteEvents(mdl.executeParams)
		mdl.executeParams = []ds.ExecuteParams{}
	}
}

// works for newBlockNum > mdl.lastEventBlock
//
func (mdl *CreditManager) onBlockChange(newBlockNum int64) {
	// on each new block
	mdl.ProcessAccountEvents(newBlockNum)
	// datacompressor works for cm address only after the address is registered with contractregister
	// i.e. discoveredAt
	// only after each event block.
	if mdl.lastEventBlock != 0 && mdl.lastEventBlock >= mdl.DiscoveredAt {
		mdl.calculateCMStat(mdl.lastEventBlock)
		mdl.lastEventBlock = 0
		// set dc data for credit manager to nil
	}
}

func bytesToUInt16(data []byte) uint16 {
	return uint16(new(big.Int).SetBytes(data).Int64())
}

func (mdl *CreditManager) OnLog(txLog types.Log) {
	extraLogs := mdl.getv2ExtraLogs(txLog)
	for _, extraLog := range extraLogs {
		mdl.logHandler(extraLog)
	}
	mdl.logHandler(txLog)
}
func (mdl *CreditManager) logHandler(txLog types.Log) {
	// creditConfigurator events for test
	// CreditFacadeUpgraded is emitted when creditconfigurator is initialized, so we will receive it on init
	// although we have already set creditfacadeUpgra
	if mdl.GetDetailsByKey("configurator") == txLog.Address.Hex() {
		switch txLog.Topics[0] {
		case core.Topic("CreditFacadeUpgraded(address)"):
			facade := utils.ChecksumAddr(txLog.Topics[1].Hex())
			mdl.setCreditFacadeSyncer(facade)
		case core.Topic("FeesUpdated(uint16,uint16,uint16,uint16,uint16)"):
			mdl.setParams(&schemas.Parameters{
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
	mdl.onTxHash(txLog.TxHash.Hex())
	// on new block
	// for credit manager stats
	blockNum := int64(txLog.BlockNumber)
	if mdl.lastEventBlock != blockNum {
		mdl.onBlockChange(blockNum)
	}
	mdl.lastEventBlock = blockNum
	//
	mdl.Repo.GetAccountManager().DeleteTxHash(blockNum, txLog.TxHash.Hex())
	switch mdl.GetVersion() {
	case 1:
		mdl.checkLogV1(txLog)
	case 2:
		mdl.checkLogV2(txLog)
	}
}

// handles for v2(for multicalls) and v1 (for executeorder)
func (mdl *CreditManager) onTxHash(newTxHash string) {
	// on txHash
	if mdl.LastTxHash != "" && mdl.LastTxHash != newTxHash {
		switch mdl.GetVersion() {
		case 1:
			// storing execute order in a single tx and processing them in a single go on next tx
			// for credit session stats
			//
			// execute events are matched with tenderly response to get transfers for each events
			mdl.processExecuteEvents()
		case 2:
			mdl.onNewTxHashV2()
		}
	}
	mdl.LastTxHash = newTxHash
}
