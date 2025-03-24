package chainlink_wrapper

import (
	"encoding/hex"
	"sort"

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
	adapters     []ds.SyncAdapterI
	storeLocally bool
}

func NewChainlinkWrapper(client core.ClientI) *ChainlinkWrapper {
	w := &ChainlinkWrapper{
		SyncWrapper:  wrappers.NewSyncWrapper(ds.ChainlinkWrapper, client),
		storeLocally: false,
	}
	// not using onLogs
	w.ViaDataProcess = ds.ViaMultipleLogs
	return w
}

func (c *ChainlinkWrapper) AddSyncAdapter(adap ds.SyncAdapterI) {
	if c.storeLocally {
		log.Info(adap.GetAddress())
		c.adapters = append(c.adapters, adap)
	} else {
		c.SyncWrapper.AddSyncAdapter(adap)
	}
}

func (w *ChainlinkWrapper) OnLogs(txLogs []types.Log) {
	w.storeLocally = true // send new adapter to buffer
	logsByAdapter := map[common.Address][]types.Log{}
	for _, txLog := range txLogs {
		logsByAdapter[txLog.Address] = append(logsByAdapter[txLog.Address], txLog)
	}
	for addr, txLogs := range logsByAdapter {
		adapter := w.Adapters.Get(addr.Hex())
		lastSync := adapter.GetLastSync()
		n := sort.Search(len(txLogs), func(i int) bool {
			return lastSync < int64(txLogs[i].BlockNumber)
		})
		txLogs = txLogs[n:]
		adapter.OnLogs(txLogs)
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
	w.SyncWrapper.SetLastSync(syncedTill)
	//
	for _, adap := range w.adapters {
		w.SyncWrapper.AddSyncAdapter(adap)
	}
	w.adapters = nil
}
