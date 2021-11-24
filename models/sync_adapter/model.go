package sync_adapter

import (
	"github.com/Gearbox-protocol/gearscan/ethclient"
	"github.com/Gearbox-protocol/gearscan/log" 
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum"
	"context"
	"math/big"
	"fmt"
)
const MaxUint = ^int64(0)
type SyncAdapter struct {
	discoveredAt uint64
	deployedAt uint64
	LastSync uint64
	Address common.Address
	Type string
	StateLoaded bool
	Dormant bool
	client *ethclient.Client
}

type SyncAdapterI interface {
	OnLog()
	Sync()
	LoadState()
}

func (s *SyncAdapter) init() {
	// s.Address
}

func (s *SyncAdapter) discoverFirstLog(address string) int64 {

	log.Debugf("Discovering external contract/AP at: %s\n", address)
	lastBlock, err := s.client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal("Cant get last block at discovery " + err.Error())
	}

	deployedAt, err := s.findFirstLogBound(address, 1, int64(lastBlock))
	if err != nil {
		log.Fatal("Cant find deployment events " + err.Error())
	}

	return deployedAt
}

func (s *SyncAdapter) findFirstLogBound(address string, fromBlock, toBlock int64) (int64, error) {

	addressHex := common.HexToAddress(address)

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(fromBlock),
		ToBlock:   big.NewInt(toBlock),
		Addresses: []common.Address{
			addressHex,
		},
		Topics: [][]common.Hash{},
	}

	logs, err := s.client.FilterLogs(context.Background(), query)
	if err != nil {
		if err.Error() == "query returned more than 10000 results" ||
			err.Error() == "Log response size exceeded. You can make eth_getLogs requests with up to a 2K block range and no limit on the response size, or you can request any block range with a cap of 10K logs in the response." {
			middle := (fromBlock + toBlock) / 2

			log.Verbosef("Run in range %d %d", fromBlock, middle-1)
			foundLow, err := s.findFirstLogBound(address, fromBlock, middle-1)
			if err != nil && err.Error() != "no events found" {
				return 0, err
			}

			log.Verbosef("Run in range %d %d", middle, toBlock)
			foundHigh, err := s.findFirstLogBound(address, middle, toBlock)
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

	deployedAt := int64(0)

	for _, vLog := range logs {
		block := int64(vLog.BlockNumber)
		if block < deployedAt || deployedAt == 0 {
			deployedAt = block
		}
	}

	if deployedAt == MaxUint {
		return 0, fmt.Errorf("no events found")
	}

	return deployedAt, nil
}