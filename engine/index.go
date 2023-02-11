package engine

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/address_provider"
	"github.com/Gearbox-protocol/third-eye/models/pool_lmrewards"
	"github.com/Gearbox-protocol/third-eye/repository"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Engine struct {
	*core.Node
	config              *config.Config
	repo                *repository.Repository
	debtEng             ds.DebtEngineI
	syncedBlock         atomic.Value
	batchSizeForHistory int64
	UsingThreads        bool
}

// var syncBlockBatchSize = 1000 * core.NoOfBlocksPerMin

func NewEngine(config *config.Config,
	ec core.ClientI,
	debtEng ds.DebtEngineI,
	repo ds.RepositoryI) ds.EngineI {
	eng := &Engine{
		debtEng:             debtEng,
		config:              config,
		repo:                repo.(*repository.Repository),
		batchSizeForHistory: config.BatchSizeForHistory,
		Node: &core.Node{
			Client: ec,
		},
	}
	return eng
}

func (e *Engine) UseThreads() {
	e.UsingThreads = true
}

func (e *Engine) init() {
	log.AMQPMsg("Starting Third-eye")
	e.repo.Init()
	// debt engine initialisation
	e.debtEng.ProcessBackLogs()
}

func (e *Engine) getLastSyncedTill() int64 {
	kit := e.repo.GetKit()
	kit.Details()
	if kit.LenOfLevel(0) == 0 {
		addr := common.HexToAddress(e.config.AddressProviderAddress).Hex()
		obj := address_provider.NewAddressProvider(addr, e.Client, e.repo)
		e.repo.AddSyncAdapter(obj)
		poolLMRewardObj := pool_lmrewards.NewPoolLMRewards(addr, obj.LastSync, e.Client, e.repo)
		e.repo.AddSyncAdapter(poolLMRewardObj)
		return obj.GetLastSync()
	} else {
		// it will allow syncing from scratch of least synced adapter in batches
		// NOTE: while syncing from scratch for some adapter disable the debt engine
		// as it might happen that some of the components for calculating debt are missing
		return e.repo.LoadLastAdapterSync()
	}
}

func (e *Engine) SyncHandler() {
	e.init()
	latestBlockNum := e.GetLatestBlockNumber()
	lastSyncedTill := e.getLastSyncedTill()
	e.syncedBlock.Store(lastSyncedTill)
	//
	// only do batch sync if latestblock is far from currently synced block
	if lastSyncedTill+e.batchSizeForHistory <= latestBlockNum {
		syncedTill := e.syncLoop(lastSyncedTill, latestBlockNum)
		log.Infof("Synced till %d sleeping for 5 mins", syncedTill)
	}
	for {
		latestBlockNum = e.GetLatestBlockNumber()
		e.SyncAndFlush(latestBlockNum)
		log.Infof("Synced till %d sleeping for 5 mins", latestBlockNum)
		time.Sleep(5 * time.Minute) // on kovan 5 blocks in 1 min , sleep for 5 mins
	}
}

func (e *Engine) syncLoop(syncedTill, latestBlockNum int64) int64 {
	loopStartBlock := syncedTill
	loopStartTime := time.Now()
	//
	for syncTarget := syncedTill + e.batchSizeForHistory; syncTarget <= latestBlockNum; syncTarget += e.batchSizeForHistory {
		roundStartTime := time.Now()
		e.SyncAndFlush(syncTarget)
		syncedTill = syncTarget
		//
		roundSyncDur := (time.Since(roundStartTime).Minutes()) // Since ==  Now().Sub
		timePerBlock := time.Since(loopStartTime).Minutes() / float64(syncedTill-loopStartBlock)
		remainingTime := (timePerBlock * float64(latestBlockNum-syncedTill)) / (60)
		log.Infof("Synced till %d in %f mins. Remaining time %f hrs ", syncedTill, roundSyncDur, remainingTime)
	}
	return syncedTill
}

func (e *Engine) SyncAndFlush(syncTill int64) {
	e.Sync(syncTill)
	e.repo.Flush()
	e.debtEng.CalculateDebtAndClear(syncTill)
	e.syncedBlock.Store(syncTill)
}

func (e *Engine) LastSyncedBlock() int64 {
	v := e.syncedBlock.Load()
	if v == nil {
		return 0
	}
	return v.(int64)
}

func (e *Engine) Sync(syncTill int64) {
	kit := e.repo.GetKit()
	log.Info("Sync till", syncTill)
	for lvlIndex := 0; lvlIndex < kit.Len(); lvlIndex++ {
		wg := &sync.WaitGroup{}
		for kit.Next(lvlIndex) {
			adapter := kit.Get(lvlIndex)
			// if utils.Contains([]string{core.AccountFactory, core.YearnPriceFeed, core.ChainlinkPriceFeed}, adapter.GetName()) {
			// 	continue
			// }
			if !adapter.IsDisabled() {
				wg.Add(1)
				if adapter.OnlyQueryAllowed() {
					if e.UsingThreads {
						go e.QueryModel(adapter, syncTill, wg)
					} else {
						e.QueryModel(adapter, syncTill, wg)
					}
				} else {
					if e.UsingThreads {
						go e.SyncModel(adapter, syncTill, wg)
					} else {
						e.SyncModel(adapter, syncTill, wg)
					}
				}
			}
		}
		kit.Reset(lvlIndex)
		wg.Wait()
	}
	e.repo.AfterSync(syncTill)
}

func (e *Engine) SyncModel(mdl ds.SyncAdapterI, syncTill int64, wg *sync.WaitGroup) {
	defer wg.Done()
	syncFrom := mdl.GetLastSync() + 1
	if syncFrom > syncTill {
		return
	}
	syncTill = utils.Min(mdl.GetBlockToDisableOn(), syncTill)
	addrsForLogs := []common.Address{common.HexToAddress(mdl.GetAddress())}
	addrsForLogs = append(addrsForLogs, mdl.GetOtherAddrsForLogs()...) // currently being used for composite eth/usd
	mdl.WillBeSyncedTo(syncTill)
	//
	txLogs, err := e.GetLogs(syncFrom, syncTill, addrsForLogs, mdl.Topics())
	log.Infof("Sync %s(%s) from %d to %d: no: %d", mdl.GetName(), mdl.GetAddress(), syncFrom, syncTill, len(txLogs))
	if err != nil {
		log.Fatal(err)
	}
	if mdl.GetHasOnLogs() {
		for _, txLog := range txLogs {
			e.isEventPausedOrUnParsed(txLog)
		}
		mdl.OnLogs(txLogs)
	} else {
		for _, txLog := range txLogs {
			blockNum := int64(txLog.BlockNumber)
			// if mdl.GetBlockToDisableOn() < blockNum {
			// 	break
			// }
			e.repo.SetBlock(blockNum)
			// parse and unpause events
			e.isEventPausedOrUnParsed(txLog)
			// pass the event to the onlog handler
			mdl.OnLog(txLog)

		}
	}
	// after sync
	mdl.AfterSyncHook(utils.Min(mdl.GetBlockToDisableOn(), syncTill))
	// log.Infof("synced %s(%s)", mdl.GetName(), mdl.GetAddress())
}

func (e *Engine) isEventPausedOrUnParsed(txLog types.Log) bool {
	switch txLog.Topics[0] {
	case core.Topic("Paused(address)"):
		e.repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: int64(txLog.BlockNumber),
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    txLog.Address.Hex(),
			Type:        schemas.Paused,
			Args:        &core.Json{"account": common.BytesToAddress(txLog.Data).Hex()},
		})
		return true
	case core.Topic("Unpaused(address)"):
		e.repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: int64(txLog.BlockNumber),
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    txLog.Address.Hex(),
			Type:        schemas.UnPaused,
			Args:        &core.Json{"account": common.BytesToAddress(txLog.Data).Hex()},
		})
		return true
	default:
		return false
	}
}

func (e *Engine) QueryModel(mdl ds.SyncAdapterI, queryTill int64, wg *sync.WaitGroup) {
	defer wg.Done()
	if mdl.GetLastSync()+1 > queryTill {
		return
	}
	// if disable block is set disable after that.
	queryTill = utils.Min(mdl.GetBlockToDisableOn(), queryTill)
	mdl.Query(queryTill)
	// after sync
	mdl.AfterSyncHook(queryTill)
}
