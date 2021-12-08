package engine

import (
	"context"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/models/address_provider"
	"github.com/Gearbox-protocol/third-eye/utils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strings"
	"sync"
	"time"
)

type Engine struct {
	config              *config.Config
	client              *ethclient.Client
	repo                core.RepositoryI
	syncBlockBatchSize  int64
	currentlySyncedTill int64
}

func NewEngine(config *config.Config,
	ec *ethclient.Client,
	repo core.RepositoryI) core.EngineI {
	return &Engine{
		config: config,
		client: ec,
		repo:   repo,
	}
}

func (e *Engine) init() {
	e.syncBlockBatchSize = 1000 * 5
	kit := e.repo.GetKit()
	kit.Details()
	if kit.LenOfLevel(0) == 0 {
		addr := common.HexToAddress(e.config.AddressProviderAddress).Hex()
		obj := address_provider.NewAddressProvider(addr, e.client, e.repo)
		e.repo.AddSyncAdapter(obj)
		e.currentlySyncedTill = obj.GetLastSync()
	} else {
		e.currentlySyncedTill = kit.First(0).GetLastSync()
	}
}

func (e *Engine) getLatestBlockNumber() int64 {
	latestBlockNum, err := e.client.BlockNumber(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Lastest blocknumber", latestBlockNum)
	return int64(latestBlockNum)
}
func (e *Engine) SyncHandler() {
	e.init()
	latestBlockNum := e.getLatestBlockNumber()
	e.syncLoop(latestBlockNum)
	for {
		log.Infof("Synced till %d sleeping for 5 mins", e.currentlySyncedTill)
		time.Sleep(5 * time.Minute) // on kovan 5 blocks in 1 min , sleep for 5 mins
		latestBlockNum = e.getLatestBlockNumber()
		e.sync(latestBlockNum)
	}
}

func (e *Engine) syncLoop(latestBlockNum int64) {
	syncTill := e.currentlySyncedTill
	syncStart := syncTill
	for syncTill <= latestBlockNum {
		roundStartTime := time.Now()
		e.sync(syncTill)
		roundSyncDur := (time.Now().Sub(roundStartTime).Minutes())
		syncTimePerBlock := roundSyncDur / float64(syncTill-syncStart)
		remainingTime := (syncTimePerBlock * float64(latestBlockNum-syncTill)) / (60)
		log.Infof("Synced till %d in %f .Remaining time %f hrs ", e.currentlySyncedTill, roundSyncDur, remainingTime)
		// new sync target
		syncTill += e.syncBlockBatchSize
	}
}

func (e *Engine) sync(syncTill int64) {
	kit := e.repo.GetKit()
	log.Info("Sync till", syncTill)
	for lvlIndex := 0; lvlIndex < kit.Len(); lvlIndex++ {
		wg := &sync.WaitGroup{}
		for kit.Next(lvlIndex) {
			adapter := kit.Get(lvlIndex)
			if !adapter.IsDisabled() {
				wg.Add(1)
				if adapter.OnlyQueryAllowed() {
					go adapter.Query(syncTill, wg)
				} else {
					go e.SyncModel(adapter, syncTill, wg)
				}
			}
		}
		kit.Reset(lvlIndex)
		wg.Wait()
	}
	e.repo.Flush()
	e.repo.CalculateDebt()
	e.currentlySyncedTill = syncTill
}

func (e *Engine) SyncModel(mdl core.SyncAdapterI, syncTill int64, wg *sync.WaitGroup) {
	defer wg.Done()
	syncFrom := mdl.GetLastSync() + 1
	if syncFrom > syncTill {
		return
	}

	log.Infof("Sync %s(%s) from %d to %d", mdl.GetName(), mdl.GetAddress(), syncFrom, syncTill)
	logs, err := e.GetLogs(syncFrom, syncTill, mdl.GetAddress())
	if err != nil {
		log.Fatal(err)
	}
	for _, log := range logs {
		blockNum := int64(log.BlockNumber)
		if mdl.GetBlockToDisableOn() < blockNum {
			break
		}
		e.repo.SetBlock(blockNum)
		mdl.OnLog(log)
	}
	// after sync
	mdl.AfterSyncHook(utils.Min(mdl.GetBlockToDisableOn(), syncTill))
}

func (e *Engine) GetLogs(fromBlock, toBlock int64, addr string) ([]types.Log, error) {
	query := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetInt64(fromBlock),
		ToBlock:   new(big.Int).SetInt64(toBlock),
		Addresses: []common.Address{common.HexToAddress(addr)},
	}
	var logs []types.Log
	var err error
	logs, err = e.client.FilterLogs(context.Background(), query)
	if err != nil {
		if err.Error() == "query returned more than 10000 results" ||
			strings.Contains(err.Error(), "Log response size exceeded. You can make eth_getLogs requests with up to a 2K block range and no limit on the response size, or you can request any block range with a cap of 10K logs in the response.") {
			middle := (fromBlock + toBlock) / 2
			log.Info(fromBlock, middle, toBlock)
			bottomHalfLogs, err := e.GetLogs(fromBlock, middle-1, addr)
			if err != nil {
				return []types.Log{}, err
			}
			logs = append(logs, bottomHalfLogs...)

			topHalfLogs, err := e.GetLogs(middle, toBlock, addr)
			if err != nil {
				return []types.Log{}, err
			}
			logs = append(logs, topHalfLogs...)
			return logs, nil
		}
	}
	return logs, err
}
