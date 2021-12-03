package engine

import (
	"context"
	"math/big"
	"time"

	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/models/address_provider"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
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
	log.Info("init sync adapters", kit.Details())
	if kit.Len() == 0 {
		addr := common.HexToAddress(e.config.AddressProviderAddress).Hex()
		obj := address_provider.NewAddressProvider(addr, e.client, e.repo)
		e.repo.AddSyncAdapter(obj)
		e.currentlySyncedTill = obj.GetLastSync()
	} else {
		e.currentlySyncedTill = kit.First().GetLastSync()
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
	for syncTill <= latestBlockNum {
		e.sync(syncTill)
		syncTill += e.syncBlockBatchSize
	}
}

func (e *Engine) sync(syncTill int64) {
	kit := e.repo.GetKit()
	log.Info("Sync till", syncTill)
	for kit.Next() {
		e.SyncModel(kit.Get(), syncTill)
	}
	kit.Reset()
	e.repo.Flush()
	e.currentlySyncedTill = syncTill
}

func (e *Engine) SyncModel(mdl core.SyncAdapterI, syncTill int64) {
	if mdl.IsDisabled() {
		return
	}
	syncFrom := mdl.GetLastSync() + 1
	if syncFrom > syncTill {
		return
	}

	log.Infof("Sync %s(%s) from %d to %d", mdl.GetName(), mdl.GetAddress(), syncFrom, syncTill)
	query := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetInt64(syncFrom),
		ToBlock:   new(big.Int).SetInt64(syncTill),
		Addresses: []common.Address{common.HexToAddress(mdl.GetAddress())},
	}
	logs, err := e.client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	for _, log := range logs {
		e.repo.SetBlock(int64(log.BlockNumber))
		mdl.OnLog(log)
	}
	// after sync
	mdl.AfterSyncHook(syncTill)
}
