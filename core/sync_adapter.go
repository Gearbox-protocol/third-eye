package core

import (
	"context"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"math"
	"math/big"
)

const MaxUint = ^int64(0)

type SyncAdapter struct {
	*Contract
	LastSync               int64       `gorm:"column:last_sync" json:"lastSync"`
	Details                Json        `gorm:"column:details" json:"details"`
	UnderlyingStatePresent bool        `gorm:"-" json:"-"`
	Error                  string      `gorm:"column:error" json:"error"`
	Repo                   RepositoryI `gorm:"-" json:"-"`
	OnlyQuery              bool        `gorm:"-" json:"-"`
	BlockToDisableOn       int64       `gorm:"column:disabled_at" json:"disabled_at"`
	HasOnLogs              bool        `gorm:"-" json:"-"`
	V                      int16       `gorm:"column:version" json:"version"`
}

func (SyncAdapter) TableName() string {
	return "sync_adapters"
}
func (s *SyncAdapter) GetVersion() int16 {
	return s.V
}
func (s *SyncAdapter) SetVersion(version int16) {
	s.V = version
}

type SyncAdapterI interface {
	OnLog(txLog types.Log)
	OnLogs(txLog []types.Log)
	GetHasOnLogs() bool
	GetLastSync() int64
	SetLastSync(int64)
	GetAddress() string
	GetName() string
	AfterSyncHook(syncTill int64)
	IsDisabled() bool
	Disable()
	GetFirstLog() int64
	HasUnderlyingState() bool
	GetUnderlyingState() interface{}
	SetUnderlyingState(obj interface{})
	SetDetails(obj interface{})
	GetAdapterState() *SyncAdapter
	OnlyQueryAllowed() bool
	Query(queryTill int64)
	DisableOnBlock(currentBlock int64)
	SetBlockToDisableOn(blockNum int64)
	GetBlockToDisableOn() int64
	GetDiscoveredAt() int64
	GetDetailsByKey(key string) string
	GetDetails() Json
	FetchVersion(blockNum int64) int16
	GetVersion() int16
}

func (s *SyncAdapter) SetDetails(obj interface{}) {
}

func (s *SyncAdapter) GetDetailsByKey(key string) string {
	if s.Details == nil {
		return ""
	}
	value, ok := s.Details[key].(string)
	if !ok {
		log.Fatalf("Not able to parse detail field `%s` in %+v", key, s.Details)
	}
	return value
}
func (s *SyncAdapter) GetDetails() Json {
	return s.Details
}
func (s *SyncAdapter) GetHasOnLogs() bool {
	return s.HasOnLogs
}

func (s *SyncAdapter) OnLogs(txLog []types.Log) {

}

func (s *SyncAdapter) DisableOnBlock(currentBlock int64) {
	if s.BlockToDisableOn != 0 && currentBlock >= s.BlockToDisableOn {
		log.Warnf("DisableOnBlock at currentBlock(%d) and s.BlockToDisableOn(%d) for %s(%s)",
			currentBlock, s.BlockToDisableOn, s.ContractName, s.Address)
		s.Disable()
	}
}

func (s *SyncAdapter) SetBlockToDisableOn(blockNum int64) {
	s.BlockToDisableOn = blockNum
	log.Warnf("Block to disable on set for %s(%s) at %d", s.GetName(), s.GetAddress(), blockNum)
}

func (s *SyncAdapter) GetBlockToDisableOn() int64 {
	if s.BlockToDisableOn == 0 {
		return math.MaxInt64
	}
	return s.BlockToDisableOn
}

func (s *SyncAdapter) OnlyQueryAllowed() bool {
	return s.OnlyQuery
}

func (s *SyncAdapter) SetLastSync(lastSync int64) {
	s.LastSync = lastSync
}

func (s *SyncAdapter) SetError(err error) {
	s.Disabled = true
	msg := err.Error()
	msgLen := len(msg)
	if msgLen > 200 {
		msgLen = 200
	}
	s.Error = err.Error()[:msgLen]
}

func (s *SyncAdapter) AfterSyncHook(syncTill int64) {
	s.SetLastSync(syncTill)
	s.DisableOnBlock(syncTill)
}
func (s *SyncAdapter) Query(queryTill int64) {
}

func NewSyncAdapter(addr, name string, discoveredAt int64, client ethclient.ClientI, repo RepositoryI) *SyncAdapter {
	obj := &SyncAdapter{
		Contract: NewContract(addr, name, discoveredAt, client),
		Repo:     repo,
	}
	obj.LastSync = obj.FirstLogAt - 1
	// for addressProvider discoveredAt is -1 but NewContract set it to firstLogAt so this will work
	obj.V = obj.FetchVersion(discoveredAt)
	return obj
}

func (s *SyncAdapter) SetUnderlyingState(obj interface{}) {
}

func (s *SyncAdapter) GetUnderlyingState() interface{} {
	return nil
}

func (s *SyncAdapter) HasUnderlyingState() bool {
	return s.UnderlyingStatePresent
}

func (s *SyncAdapter) GetAdapterState() *SyncAdapter {
	return s
}

// func (mdl *SyncAdapter) OnLog(txLog types.Log) {
// 	log.Infof("%s\n", reflect.TypeOf(mdl))
// 	log.Infof("%+v\n", txLog)
// }

func (mdl *SyncAdapter) GetLastSync() int64 {
	return mdl.LastSync
}

func (mdl *SyncAdapter) Monitor(startBlock, endBlock int64) (chan types.Log, event.Subscription, error) {
	query := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetInt64(startBlock),
		ToBlock:   new(big.Int).SetInt64(endBlock),
		Addresses: []common.Address{common.HexToAddress(mdl.Address)},
	}
	var logs = make(chan types.Log, 2)
	s, err := mdl.Client.SubscribeFilterLogs(context.TODO(), query, logs)
	if err != nil {
		return logs, s, err
	}
	return logs, s, nil
}
