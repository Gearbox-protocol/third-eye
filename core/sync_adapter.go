package core

import (
	"context"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"math/big"
)

const MaxUint = ^int64(0)

type SyncAdapter struct {
	*Contract
	LastSync               int64       `gorm:"column:last_sync"`
	Details                Json        `gorm:"column:details"`
	UnderlyingStatePresent bool        `gorm:"-"`
	Error                  string      `gorm:"column:error"`
	Repo                   RepositoryI `gorm:"-"`
	OnlyQuery              bool        `gorm:"-"`
}

func (SyncAdapter) TableName() string {
	return "sync_adapters"
}

type SyncAdapterI interface {
	OnLog(txLog types.Log)
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
	GetAdapterState() *SyncAdapter
	OnlyQueryAllowed() bool
	Query(queryTill int64)
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
}
func (s *SyncAdapter) Query(queryTill int64) {
}

func NewSyncAdapter(addr, name string, discoveredAt int64, client *ethclient.Client, repo RepositoryI) *SyncAdapter {
	obj := &SyncAdapter{
		Contract: NewContract(addr, name, discoveredAt, client),
		Repo:     repo,
	}
	obj.LastSync = obj.FirstLogAt - 1
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
