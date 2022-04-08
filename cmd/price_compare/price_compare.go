package main

import (
	"context"
	"flag"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/models/aggregated_block_feed"
	"gorm.io/gorm/clause"
	"math"
	"math/big"
	"sort"
	"time"

	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/repository"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type DBhandler struct {
	client                  core.ClientI
	db                      *gorm.DB
	Chainlinks              map[string]*ds.SyncAdapter
	ChainlinkPrices         map[string]schemas.SortedPriceFeed
	UniPoolPrices           map[string]schemas.SortedUniPoolPrices
	tokenStartBlock         map[string]int64
	Relations               []*schemas.UniPriceAndChainlink
	blocks                  map[int64]*schemas.Block
	lastBlockForToken       map[string]int64
	ChainlinkFeedIndex      map[string]int
	StartFrom               int64
	tokens                  map[string]*schemas.Token
	blockFeed               *aggregated_block_feed.AggregatedBlockFeed
	UpdatesForUniPoolPrices []*schemas.UniPoolPrices
}

func (h *DBhandler) AddUniPricesToDB(prices *schemas.UniPoolPrices) {
	blockNum := prices.BlockNum
	if h.blocks[blockNum] == nil {
		b, err := h.client.BlockByNumber(context.Background(), big.NewInt(blockNum))
		if err != nil {
			panic(err)
		}
		log.CheckFatal(err)
		h.blocks[blockNum] = &schemas.Block{BlockNumber: blockNum, Timestamp: b.Time()}
	}
	h.blocks[blockNum].AddUniswapPrices(prices)
}

func (handler *DBhandler) populateBlockFeed(obj *aggregated_block_feed.AggregatedBlockFeed) {
	tokens := []*schemas.Token{}
	err := handler.db.Raw(`SELECT * FROM tokens`).Find(&tokens).Error
	tokenMap := map[string]*schemas.Token{}
	for _, entry := range tokens {
		tokenMap[entry.Address] = entry
	}
	handler.tokens = tokenMap
	log.CheckFatal(err)
	uniswapPools := []*schemas.UniswapPools{}
	err = handler.db.Raw(`SELECT * FROM uniswap_pools`).Find(&uniswapPools).Error
	log.CheckFatal(err)
	for _, entry := range uniswapPools {
		obj.AddUniPools(tokenMap[entry.Token], entry)
	}
}

func NewDBhandler(db *gorm.DB, client *ethclient.Client) *DBhandler {
	blockFeed := aggregated_block_feed.NewAggregatedBlockFeed(client, nil, 1)
	obj := &DBhandler{db: db,
		ChainlinkPrices:    map[string]schemas.SortedPriceFeed{},
		Chainlinks:         map[string]*ds.SyncAdapter{},
		UniPoolPrices:      map[string]schemas.SortedUniPoolPrices{},
		ChainlinkFeedIndex: map[string]int{},
		lastBlockForToken:  map[string]int64{},
		blockFeed:          blockFeed,
		client:             client,
		StartFrom:          math.MaxInt64,
		tokens:             map[string]*schemas.Token{},
		blocks:             map[int64]*schemas.Block{},
		tokenStartBlock:    map[string]int64{},
	}
	obj.populateBlockFeed(blockFeed)
	return obj
}

var WETHAddr string = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"

func (handler *DBhandler) getChainlinkPrices() {
	data := []*schemas.PriceFeed{}
	err := handler.db.Raw(`SELECT * FROM price_feeds 
		WHERE feed in (SELECT address FROM sync_adapters where type='ChainlinkPriceFeed')
		ORDER BY block_num`).Find(&data).Error
	log.CheckFatal(err)
	for _, entry := range data {
		handler.ChainlinkPrices[entry.Feed] = append(handler.ChainlinkPrices[entry.Feed], entry)
		handler.StartFrom = utils.Min(handler.StartFrom, entry.BlockNumber)
		if handler.tokenStartBlock[entry.Token] == 0 {
			handler.tokenStartBlock[entry.Token] = entry.BlockNumber
		}
		handler.tokenStartBlock[entry.Token] = utils.Min(handler.tokenStartBlock[entry.Token], entry.BlockNumber)
	}
	for _, entries := range handler.ChainlinkPrices {
		sort.Sort(entries)
	}
}

func (handler *DBhandler) getChainlinks() {
	data := []*ds.SyncAdapter{}
	err := handler.db.Raw(`(SELECT * FROM sync_adapters where type='ChainlinkPriceFeed')`).Find(&data).Error
	log.CheckFatal(err)
	for _, entry := range data {
		handler.Chainlinks[entry.Address] = entry
	}
}

func (handler *DBhandler) getUniPrices(from, to int64) bool {
	log.Infof("loaded from %d to %d", from, to)
	data := []*schemas.UniPoolPrices{}
	err := handler.db.Raw(`SELECT * FROM uniswap_pool_prices 
		WHERE block_num >= ? AND block_num < ? ORDER BY block_num`, from, to).Find(&data).Error
	log.CheckFatal(err)
	for _, entry := range data {
		// don't get uni prices before the first token price from chainlink
		if handler.tokenStartBlock[entry.Token] >= entry.BlockNum ||
			// the next uni price is in increment of 1
			entry.BlockNum == 1+handler.lastBlockForToken[entry.Token] ||
			// if the lastBlock is not set check that it is equal to from for this batch
			(handler.lastBlockForToken[entry.Token] == 0 && entry.BlockNum == from) {
		} else {
			log.Info(entry.Token, entry.BlockNum)
			startBlock := handler.lastBlockForToken[entry.Token] + 1
			// if the first log in from to to is not equal to the from
			// this means that data is missing from `from` to first log.
			if handler.lastBlockForToken[entry.Token] == 0 {
				startBlock = from
			}
			log.Info(entry.Token, startBlock, entry.BlockNum, handler.tokenStartBlock[entry.Token])
			for ; startBlock < entry.BlockNum; startBlock++ {
				_, uniPrices := handler.blockFeed.QueryData(startBlock, WETHAddr, entry.Token)
				for _, entry2 := range uniPrices {
					entry2.Token = entry.Token
					handler.AddUniPricesToDB(entry2)
					handler.addUniPoolPrices(entry2)
				}
			}
		}
		handler.addUniPoolPrices(entry)
		handler.lastBlockForToken[entry.Token] = entry.BlockNum
	}
	for _, entries := range handler.UniPoolPrices {
		sort.Sort(entries)
	}
	return len(data) > 0
}

func (handler *DBhandler) addUniPoolPrices(entry *schemas.UniPoolPrices) {
	// if entry.PriceV2 == 0 || entry.PriceV3 == 0 || entry.TwapV3 == 0 {
	// 	log.Infof("%+v\n", entry)
	// 	_, allUniPrices := handler.blockFeed.QueryData(entry.BlockNum, WETHAddr, entry.Token)
	// 	for _, newPrices := range allUniPrices {
	// 		handler.UpdatesForUniPoolPrices = append(handler.UpdatesForUniPoolPrices, newPrices)
	// 		if newPrices.PriceV2 == 0 || newPrices.PriceV3 == 0 || newPrices.TwapV3 == 0 {
	// 			log.Fatalf("%+v\n", newPrices)
	// 		}
	// 	}
	// }
	handler.UniPoolPrices[entry.Token] = append(handler.UniPoolPrices[entry.Token], entry)
}

func (handler *DBhandler) clearUniPrices() {
	for token := range handler.UniPoolPrices {
		handler.UniPoolPrices[token] = schemas.SortedUniPoolPrices{}
	}
}

type LastBlock struct {
	BlockNum int64 `gorm:"column:block_num"`
}

func (handler *DBhandler) setStartBlock() {
	if SaveRelations {
		data := LastBlock{}
		err := handler.db.Raw(`SELECT max(block_num) as block_num FROM uniswap_chainlink_relations`).Find(&data).Error
		log.CheckFatal(err)
		if data.BlockNum != 0 {
			handler.StartFrom = utils.Max(handler.StartFrom, data.BlockNum+1)
		}
	}
}

func (handler *DBhandler) Run() {
	var batchSize int64 = 5000
	handler.getChainlinks()
	handler.getChainlinkPrices()
	handler.setStartBlock()
	for ; handler.getUniPrices(handler.StartFrom, handler.StartFrom+batchSize); handler.StartFrom = handler.StartFrom + batchSize {
		for feed, adapter := range handler.Chainlinks {
			token, ok := adapter.Details["token"].(string)
			if !ok {
				log.Fatal("token parse failed ", feed, utils.ToJson(adapter))
			}
			relations := handler.findRelations(adapter, handler.UniPoolPrices[token], handler.ChainlinkPrices[feed])
			log.Info("Feed:", feed, handler.tokens[token].Symbol, "Comparsions: ", relations)
		}
		handler.save()
		handler.clearUniPrices()
	}
}

func (h *DBhandler) save() {
	if !SaveAtAll {
		return
	}
	tx := h.db.Begin()
	//
	now := time.Now()

	blocks := []*schemas.Block{}
	for _, block := range h.blocks {
		blocks = append(blocks, block)
	}
	err := tx.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(blocks, 100).Error
	log.CheckFatal(err)
	h.blocks = map[int64]*schemas.Block{}
	log.Infof("created missing uniswap_pool_prices in %f sec in blocks: %d", time.Now().Sub(now).Seconds(), len(blocks))
	now = time.Now()
	//
	if SaveRelations {
		err = tx.CreateInBatches(h.Relations, 4000).Error
		log.CheckFatal(err)
	}
	h.Relations = []*schemas.UniPriceAndChainlink{}

	err = tx.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(h.UpdatesForUniPoolPrices, 100).Error
	log.CheckFatal(err)
	h.UpdatesForUniPoolPrices = []*schemas.UniPoolPrices{}
	//
	info := tx.Commit()
	log.CheckFatal(info.Error)
	log.Infof("created uniswap_chainlink_relations in %f sec", time.Now().Sub(now).Seconds())
}

func (h *DBhandler) findRelations(adapter *ds.SyncAdapter, uniPrices schemas.SortedUniPoolPrices, chainlinkPrices schemas.SortedPriceFeed) (relations int64) {
	uniPriceInd := 0
	feed := adapter.Address
	chainlinkPriceInd := h.ChainlinkFeedIndex[feed]
	for chainlinkPriceInd < len(chainlinkPrices) &&
		uniPriceInd < len(uniPrices) &&
		chainlinkPrices[chainlinkPriceInd].BlockNumber <= uniPrices[uniPriceInd].BlockNum {
		h.ChainlinkFeedIndex[feed] = chainlinkPriceInd
		chainlinkPriceInd++
	}
	index := h.ChainlinkFeedIndex[feed]
	for ; uniPriceInd < len(uniPrices) && index < len(chainlinkPrices); index++ {
		chainPrevPrices := chainlinkPrices[index]
		var chainCurrentPrices *schemas.PriceFeed
		if index+1 < len(chainlinkPrices) {
			chainCurrentPrices = chainlinkPrices[index+1]
		}
		for ; uniPriceInd < len(uniPrices) &&
			uniPrices[uniPriceInd].BlockNum < adapter.GetBlockToDisableOn() &&
			(chainCurrentPrices == nil ||
				uniPrices[uniPriceInd].BlockNum < chainCurrentPrices.BlockNumber); uniPriceInd++ {
			relations++
			h.compareDiff(chainPrevPrices, uniPrices[uniPriceInd])
		}
	}
	return
}

func (h *DBhandler) compareDiff(pf *schemas.PriceFeed, uniPoolPrices *schemas.UniPoolPrices) {
	// previous pricefeed can be nil
	if pf == nil {
		return
	}
	h.AddUniPriceAndChainlinkRelation(&schemas.UniPriceAndChainlink{
		UniBlockNum:          uniPoolPrices.BlockNum,
		ChainlinkBlockNumber: pf.BlockNumber,
		Token:                pf.Token,
		Feed:                 pf.Feed,
	})
	if (uniPoolPrices.PriceV2 != 0 && greaterFluctuation(uniPoolPrices.PriceV2, pf.Price)) ||
		(uniPoolPrices.PriceV3 != 0 && greaterFluctuation(uniPoolPrices.PriceV3, pf.Price)) ||
		(uniPoolPrices.TwapV3 != 0 && greaterFluctuation(uniPoolPrices.TwapV3, pf.Price)) {
		// if !mdl.isNotified() {
		// mdl.uniPriceVariationNotify(pf, uniPoolPrices)
		// mdl.Details["notified"] = true
		// }
	} else {
	}
}

func (h *DBhandler) AddUniPriceAndChainlinkRelation(relation *schemas.UniPriceAndChainlink) {
	h.Relations = append(h.Relations, relation)
}

func greaterFluctuation(a, b float64) bool {
	return math.Abs((a-b)/a) > 0.03
}

var SaveRelations, SaveAtAll bool

func StartServer(lc fx.Lifecycle, handler *DBhandler, config *config.Config, shutdowner fx.Shutdowner) {
	log.NewAMQPService(config.ChainId, config.AMPQEnable, config.AMPQUrl)
	// Starting server
	lc.Append(fx.Hook{
		// To mitigate the impact of deadlocks in application startup and
		// shutdown, Fx imposes a time limit on OnStart and OnStop hooks. By
		// default, hooks have a total of 15 seconds to complete. Timeouts are
		// passed via Go's usual context.Context.
		OnStart: func(context.Context) error {
			// In production, we'd want to separate the Listen and Serve phases for
			// better error-handling.
			flag.BoolVar(&SaveRelations, "relations", false, "where to save relations or not.")
			flag.BoolVar(&SaveAtAll, "save", false, "where to save relations or not.")
			go func() {
				flag.Parse()
				handler.Run()
				shutdowner.Shutdown()
			}()
			return nil
		},
	})
}

func main() {
	app := fx.New(
		config.Module,
		fx.NopLogger,
		fx.Provide(
			ethclient.NewEthClient,
			NewDBhandler,
			repository.NewDBClient,
		),
		fx.Invoke(StartServer),
	)
	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}
	<-app.Done()
}
