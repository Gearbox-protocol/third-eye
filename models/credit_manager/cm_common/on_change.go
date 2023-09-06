package cm_common

import (
	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg/dc"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type OnDirectTokenTransferFn = func(ds.RepositoryI, *schemas.TokenTransfer, *schemas.CreditSession)
type onChangeDetails struct {
	lastEventBlock        int64
	lastTxHash            string
	lastTxHashCompleted   func(string)
	calculateCMStat       func(int64, dc.CMCallData)
	onDirectTokenTransfer OnDirectTokenTransferFn
}

func (details *onChangeDetails) SetLastTxHashCompleted(fn func(string)) {
	details.lastTxHashCompleted = fn
}
func (details *onChangeDetails) SetOnDirectTokenTransferFn(fn func(ds.RepositoryI, *schemas.TokenTransfer, *schemas.CreditSession)) {
	details.onDirectTokenTransfer = fn
}
func (details *onChangeDetails) SetCalculateCMStatFn(fn func(int64, dc.CMCallData)) {
	details.calculateCMStat = fn
}

// works for newBlockNum > mdl.lastEventBlock
func (mdl *CommonCMAdapter) OnBlockChange(lastBlockNum int64) (calls []multicall.Multicall2Call, processFns []func(multicall.Multicall2Result)) {
	// datacompressor works for cm address only after the address is registered with contractregister
	// i.e. discoveredAt
	if mdl.lastEventBlock != 0 && mdl.lastEventBlock == lastBlockNum && lastBlockNum >= mdl.DiscoveredAt {
		//// ON NEW TXHASH
		mdl.processLastTx("")
		// ON NEW BLOCK
		data := mdl.Repo.GetAccountManager().CheckTokenTransfer(mdl.GetAddress(), lastBlockNum, lastBlockNum+1)
		mdl.processDirectTransfersOnBlock(lastBlockNum, data[lastBlockNum])
		calls, processFns = mdl.FetchFromDCForChangedSessions(lastBlockNum)
		call, processFn := mdl.getCMCallAndProcessFn(lastBlockNum)
		if processFn != nil {
			calls = append(calls, call)
			processFns = append(processFns, processFn)
		}
		mdl.lastEventBlock = 0
	}
	return
}

func (mdl *CommonCMAdapter) getCMCallAndProcessFn(blockNum int64) (call multicall.Multicall2Call, processFn func(multicall.Multicall2Result)) {
	call, resultFn, err := mdl.Repo.GetDCWrapper().GetCreditManagerData(blockNum, common.HexToAddress(mdl.Address))
	if err != nil {
		log.Fatalf("[CM:%s] Failed preparing get Cm call %v", mdl.Address, err)
	}
	return call, func(result multicall.Multicall2Result) {
		state, err := resultFn(result.ReturnData)
		if err != nil {
			log.Fatalf("[CM:%s] Cant get data from data compressor", mdl.Address, err)
		}
		mdl.calculateCMStat(blockNum, state)
	}
}

// handles for v2(for multicalls) and v1 (for executeorder)
func (mdl *CommonCMAdapter) processLastTx(newTxHash string) {
	// on txHash
	if mdl.lastTxHash != "" && mdl.lastTxHash != newTxHash {
		mdl.lastTxHashCompleted(mdl.lastTxHash)
	}
	mdl.lastTxHash = newTxHash
}

func (mdl *CommonCMAdapter) PrefixOnLog(txLog types.Log) {
	mdl.processLastTx(txLog.TxHash.Hex())
	mdl.lastEventBlock = int64(txLog.BlockNumber)
	//
	mdl.Repo.GetAccountManager().DeleteTxHash(int64(txLog.BlockNumber), txLog.TxHash.Hex())
}
