package wrappers

import (
	"math"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type OrderedMap[T any] struct {
	m map[string]T
	a []T
}

func NewOrderedMap[T any]() OrderedMap[T] {
	return OrderedMap[T]{
		m: make(map[string]T),
		a: make([]T, 0),
	}
}

func (x OrderedMap[T]) Get(name string) T {
	return x.m[name]
}
func (x *OrderedMap[T]) Add(name string, val T) {
	x.m[name] = val
	x.a = append(x.a, val)
}

func (x OrderedMap[T]) GetAll() []T {
	return x.a
}

// we are creating sync wrappers to wrap , chainlink, creditfilter, credit manager and pools to reduce the number of rpc calls
// only for HasOnLog = true
type SyncWrapper struct {
	Adapters       OrderedMap[ds.SyncAdapterI]
	viaDataProcess int
	name           string
	lastSync       int64
}

func NewSyncWrapper(name string) *SyncWrapper {
	return &SyncWrapper{
		Adapters:       NewOrderedMap[ds.SyncAdapterI](),
		viaDataProcess: -1,
		name:           name,
		lastSync:       math.MaxInt64 - 10,
	}
}

// extra methods
func (w SyncWrapper) GetAdapter(addr string) ds.SyncAdapterI {
	return w.Adapters.Get(addr)
}

func (w *SyncWrapper) AddSyncAdapter(adapter ds.SyncAdapterI) {
	if w.viaDataProcess == -1 {
		w.viaDataProcess = adapter.GetDataProcessType()
	}
	if adapter.GetDataProcessType() != w.viaDataProcess {
		log.Fatalf("adapter(%s) have different input data process(%d) than %d", adapter.GetAddress(), adapter.GetDataProcessType(), w.viaDataProcess)
	}
	w.Adapters.Add(adapter.GetAddress(), adapter)
	w.lastSync = utils.Min(adapter.GetLastSync(), w.lastSync)
}

func (w *SyncWrapper) GetUnderlyingAdapterAddrs() (addrs []string) {
	for _, cf := range w.Adapters.GetAll() {
		if !cf.IsDisabled() {
			addrs = append(addrs, cf.GetAddress())
		}
	}
	return
}

// //////////
// //////////
func (s SyncWrapper) Topics() [][]common.Hash {
	return nil
}

func (w *SyncWrapper) GetDataProcessType() int {
	if w.viaDataProcess == -1 {
		return ds.ViaLog
	}
	return w.viaDataProcess
}

func (s SyncWrapper) GetName() string {
	return s.name
}
func (s SyncWrapper) GetAddress() string {
	return s.name
}
func (s SyncWrapper) OnLogs(txLog []types.Log) {
}

func (SyncWrapper) HasUnderlyingState() bool {
	return false
}

func (SyncWrapper) GetUnderlyingState() interface{} {
	return nil
}

func (SyncWrapper) Query(queryTill int64) {
}

func (SyncWrapper) Version() int16 {
	return 1
}
func (SyncWrapper) GetDetails() core.Json {
	return nil
}

func (SyncWrapper) GetDetailsByKey(key string) string {
	return ""
}

func (SyncWrapper) GetDiscoveredAt() int64 {
	return 0
}
func (SyncWrapper) GetBlockToDisableOn() int64 {
	return math.MaxInt64
}
func (SyncWrapper) IsDisabled() bool {
	return false
}

func (SyncWrapper) SetBlockToDisableOn(int64) {
}

// /
func (SyncWrapper) GetVersion() int16 {
	return 1
}
func (w SyncWrapper) GetLastSync() int64 {
	return w.lastSync
}

func (s SyncWrapper) OnLog(txLog types.Log) {
	adapter := s.Adapters.Get(txLog.Address.Hex())
	if adapter.GetLastSync() < int64(txLog.BlockNumber) {
		adapter.OnLog(txLog)
	}
}

func (s SyncWrapper) GetAdapterState() (states []*ds.SyncAdapter) {
	adapters := s.Adapters.GetAll()
	states = make([]*ds.SyncAdapter, 0, len(adapters))
	for _, adapter := range adapters {
		states = append(states, adapter.GetAdapterState()...)
	}
	return
}

// ///////
// if not disabled, then do the operation on the underlying sync adapter
// ///////
func (w *SyncWrapper) GetAllAddrsForLogs() (addrs []common.Address) {
	adapters := w.Adapters.GetAll()
	addrs = make([]common.Address, 0, len(adapters))
	for _, cf := range adapters {
		if !cf.IsDisabled() {
			addrs = append(addrs, common.HexToAddress(cf.GetAddress()))
		}
	}
	return
}

func (s SyncWrapper) AfterSyncHook(syncTill int64) {
	adapters := s.Adapters.GetAll()
	for _, cf := range adapters {
		if !cf.IsDisabled() {
			cf.AfterSyncHook(syncTill)
		}
	}
}

func (s SyncWrapper) WillBeSyncedTo(blockNum int64) {
	adapters := s.Adapters.GetAll()
	for _, adapter := range adapters {
		// if last sync is smaller then new sync till
		if adapter.GetLastSync() < blockNum && !adapter.IsDisabled() {
			adapter.WillBeSyncedTo(blockNum)
		}
	}
}
