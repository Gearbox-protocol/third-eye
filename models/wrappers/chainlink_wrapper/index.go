package chainlink_wrapper

import (
	"encoding/hex"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/wrappers"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type ChainlinkWrapper struct {
	*wrappers.SyncWrapper
}

func NewChainlinkWrapper(client core.ClientI) *ChainlinkWrapper {
	w := &ChainlinkWrapper{
		SyncWrapper: wrappers.NewSyncWrapper(ds.ChainlinkWrapper, client),
	}
	// not using onLogs
	w.ViaDataProcess = ds.ViaMultipleLogs
	return w
}

func (w *ChainlinkWrapper) OnLogs(txLogs []types.Log) {
	logsByAdapter := map[common.Address][]types.Log{}
	for _, txLog := range txLogs {
		logsByAdapter[txLog.Address] = append(logsByAdapter[txLog.Address], txLog)
	}
	for adapter, txLogs := range logsByAdapter {
		w.Adapters.Get(adapter.String()).OnLogs(txLogs)
	}
}

func (w *ChainlinkWrapper) AfterSyncHook(syncedTill int64) {
	adapters := w.Adapters.GetAll()
	calls := []multicall.Multicall2Call{}
	aggregatorBytes, err := hex.DecodeString("245a7bfc") // aggregator on chainlink oracle
	log.CheckFatal(err)
	for _, adap := range adapters {
		calls = append(calls, multicall.Multicall2Call{
			Target:   common.HexToAddress(adap.GetDetailsByKey("oracle")),
			CallData: aggregatorBytes,
		})
	}
	results := core.MakeMultiCall(w.Client, syncedTill, false, calls)
	for ind, cf := range adapters {
		newAddr, ok := core.MulticallAnsAddress(results[ind])
		if !ok {
			log.Fatal("failed to get aggregator address", cf.GetDetailsByKey("oracle"), syncedTill)
		}
		if !cf.IsDisabled() {
			cf.(interface {
				AfterSyncHookWithPF(syncedTill int64, newPriceFeed common.Address)
			}).AfterSyncHookWithPF(syncedTill, newAddr)
		}
	}
}
