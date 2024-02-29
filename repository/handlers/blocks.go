package handlers

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/core/schemas/schemas_v3"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/ethereum/go-ethereum/core/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BlocksRepo struct {
	blocks *core.MutexDS[int64, *schemas.Block]
	// for treasury to get the date
	blockDatePairs *core.MutexDS[int64, *schemas.BlockDate]
	client         core.ClientI
	db             *gorm.DB
	prevStore      *PrevPriceStore
}

func NewBlocksRepo(db *gorm.DB, client core.ClientI, cfg *config.Config, tokensRepo *TokensRepo) *BlocksRepo {
	blocksRepo := &BlocksRepo{
		blocks:         core.NewMutexDS[int64, *schemas.Block](),
		blockDatePairs: core.NewMutexDS[int64, *schemas.BlockDate](),
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
		Preload("RebaseDetailsForDB").
		// v3
		Preload("QuotaDetails").
		Find(&data, "id > ? AND id <= ?", from, to).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, block := range data {
		repo.blocks.Set(block.BlockNumber, block)
	}
}

func (repo *BlocksRepo) Save(tx *gorm.DB, blockNum int64) {
	defer utils.Elapsed("blocks sql statements")()
	blocksToSync := make([]*schemas.Block, 0, len(repo.GetBlocks()))
	a := []*schemas.PriceFeed{}
	for _, block := range repo.GetBlocks() {
		a = append(a, block.PriceFeeds...)
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
func (repo *BlocksRepo) GetPrice(token string) *big.Int {
	store := repo.prevStore.prevPriceFeeds[schemas.V3PF_MAIN]
	if store != nil {
		return store[token].PriceBI.Convert()
	}

	store = repo.prevStore.prevPriceFeeds[schemas.V2PF]
	if store != nil {
		return store[token].PriceBI.Convert()
	}
	return nil
}

func (repo *BlocksRepo) GetBlocks() map[int64]*schemas.Block {
	return repo.blocks.GetInner()
}

func (repo *BlocksRepo) GetBlockDatePairs(ts int64) *schemas.BlockDate {
	return repo.blockDatePairs.Get(ts)
}

func (repo *BlocksRepo) fetchBlock(blockNum int64) *types.Block {
	b, err := repo.client.BlockByNumber(context.Background(), big.NewInt(blockNum))
	// if err != nil && err.Error() == "server returned empty uncle list but block header indicates uncles" {
	// 	repo.blocks[blockNum] = &core.Block{BlockNumber: blockNum}
	// 	return
	// }
	if err != nil {
		log.Fatalf("%s: %d", err, blockNum)
	}
	return b
}
func (repo *BlocksRepo) setBlock(blockNum int64) {
	if repo.blocks.Get(blockNum) == nil {
		b := repo.fetchBlock(blockNum)
		repo.blocks.Set(blockNum, &schemas.Block{BlockNumber: blockNum, Timestamp: b.Time()})
		repo.addBlockDate(&schemas.BlockDate{BlockNum: blockNum, Timestamp: int64(b.Time())})
	}
}

func (repo *BlocksRepo) SetBlock(blockNum int64) {
	repo.setBlock(blockNum)
}

func (repo *BlocksRepo) _setAndGetBlock(blockNum int64) *schemas.Block {
	repo.setBlock(blockNum)
	return repo.blocks.Get(blockNum)
}

func (repo *BlocksRepo) SetAndGetBlock(blockNum int64) *schemas.Block {
	return repo._setAndGetBlock(blockNum)
}

// for treasury calculations
func (repo *BlocksRepo) addBlockDate(entry *schemas.BlockDate) {
	ts := utils.TimeToDateEndTs(time.Unix(entry.Timestamp, 0))
	previousEntry := repo.blockDatePairs.Get(ts)
	if previousEntry == nil || previousEntry.BlockNum < entry.BlockNum {
		repo.blockDatePairs.Set(ts, entry)
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

func (repo *BlocksRepo) IsBlockRecent(block int64, dur time.Duration) bool {
	ts := repo._setAndGetBlock(block).Timestamp
	return time.Since(time.Unix(int64(ts), 0)) < dur
}
func (repo *BlocksRepo) RecentMsgf(header log.RiskHeader, msg string, args ...interface{}) {
	if !repo.IsBlockRecent(header.BlockNumber, time.Hour) {
		return
	}
	if header.EventCode == "AMQP" {
		log.AMQPMsgf(msg, args...)
	} else {
		log.SendRiskAlert(log.RiskAlert{
			Msg:        fmt.Sprintf(msg, args...),
			RiskHeader: header,
		})
	}
}

func (repo *BlocksRepo) Clear() {
	repo.blocks.Clear()
}

// setter
func (repo *BlocksRepo) AddPriceFeed(pf *schemas.PriceFeed) {

	if pf.MergedPFVersion == 0 {
		log.Fatal(utils.ToJson(pf))
	}
	if repo.prevStore.isPFAdded(pf) {
		repo.SetAndGetBlock(pf.BlockNumber).AddPriceFeed(pf)
	}
}

func (repo *BlocksRepo) AddDAOOperation(operation *schemas.DAOOperation) {
	repo.SetAndGetBlock(operation.BlockNumber).AddDAOOperation(operation)
}

func (repo *BlocksRepo) AddCreditManagerStats(cms *schemas.CreditManagerStat) {
	repo.SetAndGetBlock(cms.BlockNum).AddCreditManagerStats(cms)
}
func (repo *BlocksRepo) AddTokenLTRamp(details *schemas_v3.TokenLTRamp) {
	repo.SetAndGetBlock(details.BlockNum).AddTokenLTRamp(details)
}
func (repo *BlocksRepo) AddQuotaDetails(details *schemas_v3.QuotaDetails) {
	repo.SetAndGetBlock(details.BlockNum).AddQuotaDetails(details)
}

func (repo *BlocksRepo) AddPoolStat(ps *schemas.PoolStat) {
	repo.SetAndGetBlock(ps.BlockNum).AddPoolStat(ps)
}
func (repo *BlocksRepo) AddDieselTransfer(transfer *schemas.DieselTransfer) {
	repo.SetAndGetBlock(transfer.BlockNum).AddDieselTransfer(transfer)
}
func (repo *BlocksRepo) AddRebaseDetailsForDB(transfer *schemas.RebaseDetailsForDB) {
	log.Info("RebaseToken details added at", transfer.BlockNum)
	repo.SetAndGetBlock(transfer.BlockNum).AddRebaseDetailsForDB(transfer)
}

func (repo *BlocksRepo) TransferAccountAllowed(obj *schemas.TransferAccountAllowed) {
	repo.SetAndGetBlock(obj.BlockNumber).AddTransferAccountAllowed(obj)
}
