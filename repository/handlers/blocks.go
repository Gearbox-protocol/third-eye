package handlers

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/core/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BlocksRepo struct {
	blocks map[int64]*schemas.Block
	// for treasury to get the date
	blockDatePairs map[int64]*schemas.BlockDate
	mu             *sync.Mutex
	client         core.ClientI
	db             *gorm.DB
}

func NewBlocksRepo(db *gorm.DB, client core.ClientI) *BlocksRepo {
	return &BlocksRepo{
		blocks:         make(map[int64]*schemas.Block),
		blockDatePairs: map[int64]*schemas.BlockDate{},
		//
		mu:     &sync.Mutex{},
		client: client,
		db:     db,
	}
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

func (repo *BlocksRepo) Save(tx *gorm.DB) {
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
}

// external funcs
func (repo *BlocksRepo) GetBlocks() map[int64]*schemas.Block {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	return repo.blocks
}

func (repo *BlocksRepo) GetBlockDatePairs(ts int64) *schemas.BlockDate {
	repo.mu.Lock()
	defer repo.mu.Unlock()
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

func (repo *BlocksRepo) LoadBlockDatePair() {
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

func (repo *BlocksRepo) RecentEventMsg(blockNum int64, msg string, args ...interface{}) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	ts := repo._setAndGetBlock(blockNum).Timestamp
	if time.Now().Sub(time.Unix(int64(ts), 0)) < time.Hour {
		log.Msgf(msg, args...)
	}
}

func (repo *BlocksRepo) Clear() {
	repo.blocks = map[int64]*schemas.Block{}
}

// setter
func (repo *BlocksRepo) AddDAOOperation(operation *schemas.DAOOperation) {
	repo.SetAndGetBlock(operation.BlockNumber).AddDAOOperation(operation)
}

func (repo *BlocksRepo) AddCreditManagerStats(cms *schemas.CreditManagerStat) {
	repo.SetAndGetBlock(cms.BlockNum).AddCreditManagerStats(cms)
}

func (repo *BlocksRepo) AddRepayOnCM(blockNum int64, cmAddr string, pnlOnRepay schemas.PnlOnRepay) {
	repo.SetAndGetBlock(blockNum).AddRepayOnCM(cmAddr, &pnlOnRepay)
}

func (repo *BlocksRepo) AddPriceFeed(pf *schemas.PriceFeed) {
	repo.SetAndGetBlock(pf.BlockNumber).AddPriceFeed(pf)
}

func (repo *BlocksRepo) AddUniswapPrices(prices *schemas.UniPoolPrices) {
	repo.SetAndGetBlock(prices.BlockNum).AddUniswapPrices(prices)
}

func (repo *BlocksRepo) AddPoolStat(ps *schemas.PoolStat) {
	repo.SetAndGetBlock(ps.BlockNum).AddPoolStat(ps)
}

func (repo *BlocksRepo) TransferAccountAllowed(obj *schemas.TransferAccountAllowed) {
	repo.SetAndGetBlock(obj.BlockNumber).AddTransferAccountAllowed(obj)
}

// getter
func (repo *BlocksRepo) GetRepayOnCM(blockNum int64, cmAddr string) *schemas.PnlOnRepay {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	return repo.blocks[blockNum].GetRepayOnCM(cmAddr)
}
