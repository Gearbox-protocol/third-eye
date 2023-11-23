package cm_wrapper

import (
	"sort"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/credit_manager/cm_v1"
	"github.com/Gearbox-protocol/third-eye/models/credit_manager/cm_v2"
	"github.com/Gearbox-protocol/third-eye/models/wrappers"
	"github.com/ethereum/go-ethereum/core/types"
)

type CMWrapper struct {
	*wrappers.SyncWrapper
}

func NewCMWrapper(client core.ClientI) *CMWrapper {
	w := &CMWrapper{
		SyncWrapper: wrappers.NewSyncWrapper(ds.CMWrapper, client),
	}
	w.ViaDataProcess = ds.ViaMultipleLogs
	return w
}

func (s CMWrapper) OnLogs(txLogs []types.Log) {
	ind := 0
	var lastBlockNum int64 = 0
	for ind < len(txLogs) {
		txLog := txLogs[ind]
		// check block change
		newBlockNum := int64(txLog.BlockNumber)
		if lastBlockNum == 0 {
			lastBlockNum = newBlockNum
		}
		if lastBlockNum != newBlockNum {
			s.onBlockChange(lastBlockNum, newBlockNum)
			lastBlockNum = newBlockNum
		}
		// process txLog
		s.OnLog(txLog)
		ind++
		// check if the addr changed in creditManager, if changed fetch all logs again for new addr set.

		if s.adapterAddrsChanged(txLog.Address.Hex()) {
			newTxLogs, err := pkg.Node{Client: s.Client}.GetLogs(int64(txLog.BlockNumber), s.WillSyncTill, s.GetAllAddrsForLogs(), nil)
			log.CheckFatal(err)
			splitInd := sort.Search(len(newTxLogs), func(i int) bool {
				return newTxLogs[i].BlockNumber > txLog.BlockNumber ||
					(newTxLogs[i].BlockNumber == txLog.BlockNumber && newTxLogs[i].Index > txLog.Index)
			})
			txLogs = newTxLogs[splitInd:]
			ind = 0
		}
	}
	// check block change
	if lastBlockNum != 0 {
		s.onBlockChange(lastBlockNum, s.WillSyncTill+1)
	}
}

type blockChangeI interface {
	OnBlockChange(int64) ([]multicall.Multicall2Call, []func(multicall.Multicall2Result))
	UpdateSessionWithDirectTokenTransferBefore(int64)
	IsAddrChanged() bool
}

func getCM(adapter ds.SyncAdapterI) (cm blockChangeI) {
	switch v := adapter.(type) {
	case *cm_v1.CMv1:
		cm = v
	case *cm_v2.CMv2:
		cm = v
	}
	return
}

func (s CMWrapper) onBlockChange(lastBlockNum, newBlockNum int64) {
	adapters := s.Adapters.GetAll()
	//
	calls := make([]multicall.Multicall2Call, 0, len(adapters))
	processFns := make([]func(multicall.Multicall2Result), 0, len(adapters))
	//
	for _, adapter := range adapters {
		if adapter.GetLastSync() >= lastBlockNum {
			continue
		}
		cm := getCM(adapter)
		call, processFn := cm.OnBlockChange(lastBlockNum)
		// if process fn is not null
		if processFn != nil {
			processFns = append(processFns, processFn...)
			calls = append(calls, call...)
		}
	}
	results := core.MakeMultiCall(s.Client, lastBlockNum, false, calls)
	for ind, result := range results {
		processFns[ind](result)
	}
	// update for direct token transfer
	for _, adapter := range adapters {
		if adapter.GetLastSync() >= lastBlockNum {
			continue
		}
		cm := getCM(adapter)
		cm.UpdateSessionWithDirectTokenTransferBefore(newBlockNum)
	}
}

func (s CMWrapper) adapterAddrsChanged(addr string) bool {
	adapter := s.Adapters.GetFromLogAddr(addr)
	changed := getCM(adapter).IsAddrChanged()
	if changed {
		s.Adapters.Add(adapter.GetAddress(), adapter.GetAllAddrsForLogs(), adapter)
	}
	return changed
}
