package core

import (
	"context"
	"fmt"
	"github.com/Gearbox-protocol/gearscan/ethclient"
	"github.com/Gearbox-protocol/gearscan/log"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"math/big"
)

const MaxUint = ^int64(0)

type SyncAdapter struct {
	DiscoveredAt int64             `gorm:"column:discovered_at"`
	FirstLogAt   int64             `gorm:"column:firstlog_at"`
	LastSync     int64             `gorm:"column:last_sync"`
	Address      string            `gorm:"primaryKey;column:address"`
	Disabled     bool              `gorm:"column:disabled"`
	Client       *ethclient.Client `gorm:"-"`
	Type         string            `gorm:"column:type"`
}

func (SyncAdapter) TableName() string {
	return "sync_adapters"
}

type SyncAdapterI interface {
	OnLog(txLog types.Log)
	GetLastSync() int64
	SetLastSync(int64)
	GetAdapterState() *SyncAdapter
	GetAddress() string
	FirstSync() bool
	GetType() string
}

func (s *SyncAdapter) SetLastSync(lastSync int64) {
	s.LastSync = lastSync
}

func (s *SyncAdapter) FirstSync() bool {
	return s.FirstLogAt == s.LastSync
}
func (s *SyncAdapter) GetAddress() string {
	return s.Address
}
func (s *SyncAdapter) GetType() string {
	return s.Type
}

// func NewSyncAdapter(addr string, client *ethclient.Client) SyncAdapterI {
// 	obj := &SyncAdapter{
// 			Type: "AddressProvider",
// 			Address: addr,
// 			Client: client,
// 		}
// 	firstDetection:= obj.DiscoverFirstLog()
// 	fmt.Println(firstDetection)
// 	obj.DiscoveredAt = firstDetection
// 	obj.FirstLogAt = firstDetection
// 	obj.LastSync = firstDetection
// 	return obj
// }
func (s *SyncAdapter) GetAdapterState() *SyncAdapter {
	return s
}
func (s *SyncAdapter) LoadState() {

}

func (s *SyncAdapter) init() {
	// s.Address
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

func (s *SyncAdapter) DiscoverFirstLog() int64 {

	// log.Debugf("Discovering first log of: %s\n", s.Address)
	lastBlock, err := s.Client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal("Cant get last block at discovery " + err.Error())
	}

	FirstLogAt, err := s.findFirstLogBound(1, int64(lastBlock))
	if err != nil {
		log.Fatal("Cant find deployment events " + err.Error())
	}

	return FirstLogAt
}

func (s *SyncAdapter) findFirstLogBound(fromBlock, toBlock int64) (int64, error) {

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(fromBlock),
		ToBlock:   big.NewInt(toBlock),
		Addresses: []common.Address{
			common.HexToAddress(s.Address),
		},
		Topics: [][]common.Hash{},
	}

	logs, err := s.Client.FilterLogs(context.Background(), query)
	if err != nil {
		if err.Error() == "query returned more than 10000 results" ||
			err.Error() == "Log response size exceeded. You can make eth_getLogs requests with up to a 2K block range and no limit on the response size, or you can request any block range with a cap of 10K logs in the response." {
			middle := (fromBlock + toBlock) / 2

			log.Verbosef("Run in range %d %d", fromBlock, middle-1)
			foundLow, err := s.findFirstLogBound(fromBlock, middle-1)
			if err != nil && err.Error() != "no events found" {
				return 0, err
			}

			log.Verbosef("Run in range %d %d", middle, toBlock)
			foundHigh, err := s.findFirstLogBound(middle, toBlock)
			if err != nil && err.Error() != "no events found" && err.Error() != "Cant find any events" {
				return 0, err
			}

			log.Verbosef("%d %d", foundLow, foundHigh)

			if foundLow == 0 && foundHigh == 0 {
				return 0, fmt.Errorf("No events was found for the contract")
			}

			if foundLow == 0 {
				return foundHigh, nil
			}

			return foundLow, nil

		}
		return 0, err
	}

	FirstLogAt := int64(0)

	for _, vLog := range logs {
		block := int64(vLog.BlockNumber)
		if block < FirstLogAt || FirstLogAt == 0 {
			FirstLogAt = block
		}
	}

	if FirstLogAt == MaxUint {
		return 0, fmt.Errorf("no events found")
	}

	return FirstLogAt, nil
}
