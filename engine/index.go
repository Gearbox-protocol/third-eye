package engine

import (
	"context"
	"github.com/Gearbox-protocol/gearscan/config"
	"github.com/Gearbox-protocol/gearscan/core"
	"github.com/Gearbox-protocol/gearscan/ethclient"
	"github.com/Gearbox-protocol/gearscan/log"
	"github.com/Gearbox-protocol/gearscan/models/address_provider"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Engine struct {
	config       *config.Config
	client       *ethclient.Client
	repo         core.RepositoryI
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

var blockPerSync int64 = 2000

func (e *Engine) init() {
	adapters := e.repo.GetSyncAdapters()
	log.Info("init sync adapters", adapters)
	if len(adapters) == 0 {
		addr := common.HexToAddress(e.config.AddressProviderAddress).Hex()
		obj := address_provider.NewAddressProvider(addr, e.client, e.repo)
		e.repo.AddSyncAdapter(obj)
		e.nextSyncStop = obj.GetLastSync() + blockPerSync
	} else {
		e.nextSyncStop = adapters[0].GetLastSync() + blockPerSync
	}
}
func (e *Engine) Sync() {
	e.init()
	// for i := 0; i < 2; i++ {
	for {
		log.Info("Sync till", e.nextSyncStop)
		for _, adapter := range e.repo.GetSyncAdapters() {
			e.SyncModel(adapter, e.nextSyncStop)
		}
		e.repo.Flush()
		// log.Fatal("end")
		e.nextSyncStop += blockPerSync
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
