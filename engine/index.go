package engine

import (
	"context"
	"math/big"
	"time"

	"github.com/Gearbox-protocol/gearscan/config"
	"github.com/Gearbox-protocol/gearscan/core"
	"github.com/Gearbox-protocol/gearscan/ethclient"
	"github.com/Gearbox-protocol/gearscan/log"
	"github.com/Gearbox-protocol/gearscan/models/address_provider"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

type Engine struct {
	config       *config.Config
	client       *ethclient.Client
	repo         core.RepositoryI
	blockPerSync int64
	nextSyncStop int64
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
	e.blockPerSync = 400 * 5
	adapters := e.repo.GetSyncAdapters()
	log.Info("init sync adapters", adapters)
	if len(adapters) == 0 {
		addr := common.HexToAddress(e.config.AddressProviderAddress).Hex()
		obj := address_provider.NewAddressProvider(addr, e.client, e.repo)
		e.repo.AddSyncAdapter(obj)
		e.nextSyncStop = obj.GetLastSync() + e.blockPerSync
	} else {
		e.nextSyncStop = adapters[0].GetLastSync() + e.blockPerSync
	}
}

func (e *Engine) SyncHandler() {
	e.init()
	for {
		latestBlockNum, err := e.client.BlockNumber(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
		log.Info("Lastest blocknumber",latestBlockNum)
		e.sync(int64(latestBlockNum))
		log.Infof("Synced till %d sleeping for 5 mins", latestBlockNum)
		time.Sleep(5 * time.Minute)
		e.blockPerSync = 5 * 5 // on kovan 5 blocks in 1 min , sleep for 5 mins
	}
}
func (e *Engine) sync(latestBlockNum int64) {
	for ;e.nextSyncStop < latestBlockNum; {
		log.Info("Sync till", e.nextSyncStop)
		for _, adapter := range e.repo.GetSyncAdapters() {
			e.SyncModel(adapter, e.nextSyncStop)
		}
		e.repo.Flush()
		e.nextSyncStop += e.blockPerSync
	}
}

func (e *Engine) SyncModel(mdl core.SyncAdapterI, syncTill int64) {
	if mdl.IsDisabled() {
		return
	}
	syncFrom := mdl.GetLastSync()
	if !mdl.FirstSync() {
		syncFrom += 1
	}
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
