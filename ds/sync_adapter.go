package ds

import (
	"context"
	"math"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

type SyncAdapter struct {
	*schemas.SyncAdapterSchema
	UnderlyingStatePresent bool        `gorm:"-" json:"-"`
	Repo                   RepositoryI `gorm:"-" json:"-"`
	DataProcessType        int         `gorm:"-" json:"-"`
	WillSyncTill           int64       `gorm:"-" json:"-"`
}

func (SyncAdapter) TableName() string {
	return "sync_adapters"
}

func (s *SyncAdapter) GetVersion() core.VersionType {
	return s.V
}

// func (s *SyncAdapter) SetVersion(version int16) {
// 	s.V = core.NewVersion(version)
// }

const (
	ViaLog = iota
	ViaMultipleLogs
	ViaQuery
)

type SyncAdapterI interface {
	OnLog(txLog types.Log)           //
	OnLogs(txLog []types.Log)        //
	GetDataProcessType() int         //
	GetLastSync() int64              //
	GetAddress() string              //
	GetName() string                 //
	AfterSyncHook(syncTill int64)    //
	HasUnderlyingState() bool        //
	GetUnderlyingState() interface{} //
	SetUnderlyingState(obj interface{})
	GetAdapterState() *SyncAdapter //
	Query(queryTill int64)         //
	WillBeSyncedTo(blockNum int64) //
	IsDisabled() bool              //
	SetBlockToDisableOn(blockNum int64)
	GetBlockToDisableOn() int64
	GetDiscoveredAt() int64
	GetDetailsByKey(key string) string
	GetDetails() core.Json
	GetVersion() core.VersionType
	GetAllAddrsForLogs() []common.Address
	Topics() [][]common.Hash
}

func (s SyncAdapter) Topics() [][]common.Hash {
	return nil
}

func (s SyncAdapter) GetAllAddrsForLogs() (addrs []common.Address) {
	if s.GetAddress() != "" {
		addrs = append(addrs, common.HexToAddress(s.GetAddress()))
	}
	if s.Details == nil {
		return
	}
	secAddrs := s.Details["secAddrs"]
	if secAddrs == nil {
		return
	}
	m := secAddrs.(map[string]interface{})
	for _, v := range m {
		addrs = append(addrs, InterfaceToAddr(v))
	}
	return
}

func InterfaceToAddr(v interface{}) common.Address {
	switch obj := v.(type) {
	case string:
		return common.HexToAddress(obj)
	case common.Address:
		return obj
	}
	panic("")
}
func (s *SyncAdapter) WillBeSyncedTo(blockNum int64) {
	s.WillSyncTill = blockNum
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
func (s *SyncAdapter) GetDetails() core.Json {
	return s.Details
}

func (s *SyncAdapter) GetDataProcessType() int {
	return s.DataProcessType
}
func (s *SyncAdapter) OnLogs(txLog []types.Log) {

}

// disable if this block is disable block or above
func (s *SyncAdapter) disableOnBlock(currentBlock int64) {
	if s.BlockToDisableOn != 0 && currentBlock >= s.BlockToDisableOn {
		s.Disabled = true
	}
}

// will sync till this block.
// used in CreditFilter, priceOracle and priceFeed
func (s *SyncAdapter) SetBlockToDisableOn(blockNum int64) {
	s.BlockToDisableOn = blockNum
}

func (s *SyncAdapter) GetBlockToDisableOn() int64 {
	if s.BlockToDisableOn == 0 {
		return math.MaxInt64
	}
	return s.BlockToDisableOn
}

func (s *SyncAdapter) setLastSync(lastSync int64) {
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
	s.setLastSync(syncTill)
	s.disableOnBlock(syncTill)
}
func (s *SyncAdapter) Query(queryTill int64) {
}

func NewSyncAdapter(addr, name string, discoveredAt int64, client core.ClientI, repo RepositoryI) *SyncAdapter {
	obj := &SyncAdapter{
		SyncAdapterSchema: &schemas.SyncAdapterSchema{
			Contract: schemas.NewContract(addr, name, discoveredAt, client),
		},
		Repo: repo,
	}
	// for addressProvider discoveredAt is -1 but NewContract set it to firstLogAt so this will work
	obj.LastSync = obj.FirstLogAt - 1
	// version in pricefeed is not related to gearbox protocol
	if name != ChainlinkPriceFeed && name != QueryPriceFeed {
		obj.V = core.FetchVersion(addr, discoveredAt, client)
	}
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
