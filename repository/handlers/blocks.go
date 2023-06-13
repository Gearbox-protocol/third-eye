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
	"github.com/Gearbox-protocol/sdk-go/pkg/priceFetcher"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BlocksRepo struct {
	blocks map[int64]*schemas.Block
	// for treasury to get the date
	blockDatePairs map[int64]*schemas.BlockDate
	// for prevently duplicate query price feed already with same price for a token
	// token to feed
	prevPriceFeeds map[bool]map[string]*schemas.PriceFeed
	mu             *sync.Mutex
	client         core.ClientI
	db             *gorm.DB
	spotPrices     *priceFetcher.OneInchOracle
}

func NewBlocksRepo(db *gorm.DB, client core.ClientI, cfg *config.Config, tokensRepo *TokensRepo) *BlocksRepo {
	blocksRepo := &BlocksRepo{
		blocks:         make(map[int64]*schemas.Block),
		blockDatePairs: map[int64]*schemas.BlockDate{},
		prevPriceFeeds: map[bool]map[string]*schemas.PriceFeed{},
		//
		mu:     &sync.Mutex{},
		client: client,
		db:     db,
	}
	chainId, err := client.ChainID(context.TODO())
	log.CheckFatal(err)
	if chainId.Int64() == 1 {
		blocksRepo.spotPrices = priceFetcher.New1InchOracle(client, chainId.Int64(),
			common.HexToAddress("0x07D91f5fb9Bf7798734C3f606dB065549F6893bb"), tokensRepo)

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

var lastSaveBlockForTCP int64 = 0

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

	if blockNum-lastSaveBlockForTCP > core.NoOfBlocksPerMin*5 {
		repo.saveCurrentPrices(tx, blockNum)
		lastSaveBlockForTCP = blockNum
	}
}

func (repo *BlocksRepo) saveCurrentPrices(tx *gorm.DB, blockNum int64) {
	// chainlink current prices to updated
	var currentPricesToSync []*schemas.TokenCurrentPrice
	for _, tokenPrice := range repo.prevPriceFeeds[true] {
		if tokenPrice.SaveCurrentPrice {
			tokenPrice.SaveCurrentPrice = false
			currentPricesToSync = append(currentPricesToSync, &schemas.TokenCurrentPrice{
				PriceBI:  tokenPrice.PriceBI,
				Price:    tokenPrice.Price,
				BlockNum: tokenPrice.BlockNumber,
				Token:    tokenPrice.Token,
				PriceSrc: string(core.SOURCE_CHAINLINK),
			})
		}
	}
	if len(currentPricesToSync) == 0 { // usd prices are set? only valid from v2
		// so if it's empty, we don't need to store currentPrice and nor fetch 1inch prices in usdc
		return
	}
	// spot prices to updated
	if repo.spotPrices != nil {
		calls := repo.spotPrices.GetCalls()
		results := core.MakeMultiCall(repo.client, blockNum, false, calls)
		for token, priceBI := range repo.spotPrices.GetPrices(results, blockNum) {
			currentPricesToSync = append(currentPricesToSync, &schemas.TokenCurrentPrice{
				PriceBI:  priceBI,
				Price:    utils.GetFloat64Decimal(priceBI.Convert(), 8),
				BlockNum: blockNum,
				Token:    token,
				PriceSrc: string(core.SOURCE_SPOT),
			})
		}
	}
	err := tx.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(currentPricesToSync, 100).Error
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
	repo.loadPrevPriceFeed()
}

func (repo *BlocksRepo) loadPrevPriceFeed() {
	defer utils.Elapsed("loadPrevPriceFeed")()
	data := []*schemas.PriceFeed{}
	err := repo.db.Raw("SELECT distinct on(token)* FROM price_feeds ORDER BY token, block_num DESC").Find(&data).Error
	log.CheckFatal(err)
	for _, pf := range data {
		repo.addPrevPriceFeed(pf)
	}
}

// isUSD -> token -> feed -> price feed object
func (repo *BlocksRepo) addPrevPriceFeed(pf *schemas.PriceFeed) {
	if repo.prevPriceFeeds[pf.IsPriceInUSD] == nil {
		repo.prevPriceFeeds[pf.IsPriceInUSD] = map[string]*schemas.PriceFeed{}
	}
	oldPF := repo.prevPriceFeeds[pf.IsPriceInUSD][pf.Token]
	price := pf.PriceBI.Convert().Int64()
	if oldPF != nil && oldPF.BlockNumber >= pf.BlockNumber && !(price == 0 || price == 100) {
		log.Fatalf("oldPF %s.\n NewPF %s.", oldPF, pf)
	}
	//
	repo.prevPriceFeeds[pf.IsPriceInUSD][pf.Token] = pf
}

func (repo *BlocksRepo) AddPriceFeed(pf *schemas.PriceFeed) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if repo.prevPriceFeeds[pf.IsPriceInUSD] != nil &&
		repo.prevPriceFeeds[pf.IsPriceInUSD][pf.Token] != nil &&
		repo.prevPriceFeeds[pf.IsPriceInUSD][pf.Token].Feed == pf.Feed {
		prevPF := repo.prevPriceFeeds[pf.IsPriceInUSD][pf.Token]
		if prevPF.BlockNumber >= pf.BlockNumber {
			log.Fatalf("oldPF %s.\n NewPF %s.", prevPF, pf)
		}
		if prevPF.PriceBI.Cmp(pf.PriceBI) == 0 {
			repo.addPrevPriceFeed(pf)
			return
		}
	}
	pf.SaveCurrentPrice = true
	repo.addPrevPriceFeed(pf)
	repo._setAndGetBlock(pf.BlockNumber).AddPriceFeed(pf)
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
