package handlers

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/ethereum/go-ethereum/core/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BlocksRepo struct {
	mu     *sync.RWMutex
	blocks map[int64]*schemas.Block
	// for treasury to get the date
	blockDatePairs map[int64]*schemas.BlockDate
	client         core.ClientI
	db             *gorm.DB
	prevStore      *PrevPriceStore
}

func NewBlocksRepo(db *gorm.DB, client core.ClientI, cfg *config.Config, tokensRepo *TokensRepo) *BlocksRepo {
	blocksRepo := &BlocksRepo{
		blocks:         make(map[int64]*schemas.Block),
		blockDatePairs: map[int64]*schemas.BlockDate{},
		mu:             &sync.RWMutex{},
		//
		client:    client,
		db:        db,
		prevStore: NewPrevPriceStore(client, tokensRepo),
	}
	return blocksRepo
}

func (repo *BlocksRepo) LoadBlocks(from, to int64) {
	log.Infof("Loaded %d to %d blocks for debt", from, to)
	data := []*schemas.Block{}
	err := repo.db.Preload("CSS").Preload("PoolStats").
		Preload("AllowedTokens").Preload("PriceFeeds").Preload("Params").
		Find(&data, "id > ? AND id <= ?", from, to).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, block := range data {
		repo.blocks[block.BlockNumber] = block
	}
}

func (repo *BlocksRepo) Save(tx *gorm.DB, blockNum int64) {
	defer utils.Elapsed("blocks sql statements")()
	blocksToSync := make([]*schemas.Block, 0, len(repo.GetBlocks()))
	for _, block := range repo.GetBlocks() {
		blocksToSync = append(blocksToSync, block)
	}
	// clauses not needed here
	err := tx.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(blocksToSync, 100).Error
	log.CheckFatal(err)

	repo.prevStore.saveCurrentPrices(repo.client, tx, blockNum)
}

// external funcs
func (repo *BlocksRepo) GetBlocks() map[int64]*schemas.Block {
	repo.mu.RLock()
	defer repo.mu.RUnlock()
	return repo.blocks
}

func (repo *BlocksRepo) GetBlockDatePairs(ts int64) *schemas.BlockDate {
	repo.mu.RLock()
	defer repo.mu.RUnlock()
	return repo.blockDatePairs[ts]
}

func (repo *BlocksRepo) fetchBlock(blockNum int64) *types.Block {
	b, err := repo.client.BlockByNumber(context.Background(), big.NewInt(blockNum))
	// if err != nil && err.Error() == "server returned empty uncle list but block header indicates uncles" {
	// 	repo.blocks[blockNum] = &core.Block{BlockNumber: blockNum}
	// 	return
	// }
	log.CheckFatal(err)
	return b
}
func (repo *BlocksRepo) setBlock(blockNum int64) {
	if repo.blocks[blockNum] == nil {
		b := repo.fetchBlock(blockNum)
		repo.blocks[blockNum] = &schemas.Block{BlockNumber: blockNum, Timestamp: b.Time()}
		repo.addBlockDate(&schemas.BlockDate{BlockNum: blockNum, Timestamp: int64(b.Time())})
	}
}

func (repo *BlocksRepo) SetBlock(blockNum int64) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if 7954375 <= blockNum {
		log.Fatal("")
	}
	repo.setBlock(blockNum)
}

func (repo *BlocksRepo) _setAndGetBlock(blockNum int64) *schemas.Block {
	repo.setBlock(blockNum)
	return repo.blocks[blockNum]
}

func (repo *BlocksRepo) SetAndGetBlock(blockNum int64) *schemas.Block {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	return repo._setAndGetBlock(blockNum)
}

// for treasury calculations
func (repo *BlocksRepo) addBlockDate(entry *schemas.BlockDate) {
	ts := utils.TimeToDateEndTs(time.Unix(entry.Timestamp, 0))
	previousEntry := repo.blockDatePairs[ts]
	if previousEntry == nil || previousEntry.BlockNum < entry.BlockNum {
		repo.blockDatePairs[ts] = entry
	}
}

func (repo *BlocksRepo) loadBlockDatePair() {
	defer utils.Elapsed("loadBlockDatePair")()
	data := []*schemas.BlockDate{}
	sql := `select b.*, a.timestamp from blocks a 
	JOIN (select timezone('UTC', to_timestamp(timestamp))::date as date,max(id) as block_num from blocks group by date) b 
	ON a.id=b.block_num order by block_num`
	if err := repo.db.Raw(sql).Find(&data).Error; err != nil {
		log.Fatal(err)
	}
	for _, entry := range data {
		repo.addBlockDate(entry)
	}
}

func (repo *BlocksRepo) Load() {
	repo.loadBlockDatePair()
	repo.prevStore.loadPrevPriceFeed(repo.db)
}

func (repo *BlocksRepo) RecentMsgf(header log.RiskHeader, msg string, args ...interface{}) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	ts := repo._setAndGetBlock(header.BlockNumber).Timestamp
	if time.Since(time.Unix(int64(ts), 0)) < time.Hour {
		if header.EventCode == "AMQP" {
			log.AMQPMsgf(msg, args...)
		} else {
			log.SendRiskAlert(log.RiskAlert{
				Msg:        fmt.Sprintf(msg, args...),
				RiskHeader: header,
			})
		}
	}
}

func (repo *BlocksRepo) Clear() {
	repo.blocks = map[int64]*schemas.Block{}
}

// setter
func (repo *BlocksRepo) AddPriceFeed(pf *schemas.PriceFeed) {
	if repo.prevStore.canAddPF(pf) {
		repo.SetAndGetBlock(pf.BlockNumber).AddPriceFeed(pf)
	}
}

func (repo *BlocksRepo) AddDAOOperation(operation *schemas.DAOOperation) {
	repo.SetAndGetBlock(operation.BlockNumber).AddDAOOperation(operation)
}

func (repo *BlocksRepo) AddCreditManagerStats(cms *schemas.CreditManagerStat) {
	repo.SetAndGetBlock(cms.BlockNum).AddCreditManagerStats(cms)
}

func (repo *BlocksRepo) AddPoolStat(ps *schemas.PoolStat) {
	repo.SetAndGetBlock(ps.BlockNum).AddPoolStat(ps)
}
func (repo *BlocksRepo) AddDieselTransfer(transfer *schemas.DieselTransfer) {
	repo.SetAndGetBlock(transfer.BlockNum).AddDieselTransfer(transfer)
}

func (repo *BlocksRepo) TransferAccountAllowed(obj *schemas.TransferAccountAllowed) {
	repo.SetAndGetBlock(obj.BlockNumber).AddTransferAccountAllowed(obj)
}
