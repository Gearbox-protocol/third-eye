package main

import (
	"context"
	"flag"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/models/aggregated_block_feed"
	"github.com/Gearbox-protocol/third-eye/utils"
	"gorm.io/gorm/clause"
	"math"
	"math/big"
	"sort"
	"time"

	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/repository"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type DBhandler struct {
	client                  *ethclient.Client
	db                      *gorm.DB
	Chainlinks              map[string]*core.SyncAdapter
	ChainlinkPrices         map[string]core.SortedPriceFeed
	UniPoolPrices           map[string]core.SortedUniPoolPrices
	Relations               []*core.UniPriceAndChainlink
	blocks                  map[int64]*core.Block
	lastBlockForToken       map[string]int64
	ChainlinkFeedIndex      map[string]int
	StartFrom               int64
	tokens                  map[string]*core.Token
	blockFeed               *aggregated_block_feed.AggregatedBlockFeed
	UpdatesForUniPoolPrices []*core.UniPoolPrices
}

func (h *DBhandler) AddUniPrices(prices *core.UniPoolPrices) {
	log.Infof("%+v", prices)
	blockNum := prices.BlockNum
	if h.blocks[blockNum] == nil {
		b, err := h.client.BlockByNumber(context.Background(), big.NewInt(blockNum))
		if err != nil {
			panic(err)
		}
		log.CheckFatal(err)
		h.blocks[blockNum] = &core.Block{BlockNumber: blockNum, Timestamp: b.Time()}
	}
	h.blocks[blockNum].AddUniswapPrices(prices)
}

func (handler *DBhandler) populateBlockFeed(obj *aggregated_block_feed.AggregatedBlockFeed) {
	tokens := []*core.Token{}
	err := handler.db.Raw(`SELECT * FROM tokens`).Find(&tokens).Error
	tokenMap := map[string]*core.Token{}
	for _, entry := range tokens {
		tokenMap[entry.Address] = entry
	}
	handler.tokens = tokenMap
	log.CheckFatal(err)
	uniswapPools := []*core.UniswapPools{}
	err = handler.db.Raw(`SELECT * FROM tokens`).Find(&tokens).Error
	log.CheckFatal(err)
	for _, entry := range uniswapPools {
		obj.AddPools(tokenMap[entry.Token], entry)
	}
}

func NewDBhandler(db *gorm.DB, client *ethclient.Client) *DBhandler {
	blockFeed := aggregated_block_feed.NewAggregatedBlockFeed(client, nil, 1)
	obj := &DBhandler{db: db,
		ChainlinkPrices:    map[string]core.SortedPriceFeed{},
		Chainlinks:         map[string]*core.SyncAdapter{},
		UniPoolPrices:      map[string]core.SortedUniPoolPrices{},
		ChainlinkFeedIndex: map[string]int{},
		lastBlockForToken:  map[string]int64{},
		blockFeed:          blockFeed,
		client:             client,
		StartFrom:          math.MaxInt64,
		tokens:             map[string]*core.Token{},
	}
	obj.populateBlockFeed(blockFeed)
	return obj
}

var WETHAddr string = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"

func (handler *DBhandler) getChainlinkPrices() {
	data := []*core.PriceFeed{}
	err := handler.db.Raw(`SELECT * FROM price_feeds 
		WHERE feed in (SELECT address FROM sync_adapters where type='ChainlinkPriceFeed')
		ORDER BY block_num`).Find(&data).Error
	log.CheckFatal(err)
	for _, entry := range data {
		handler.ChainlinkPrices[entry.Feed] = append(handler.ChainlinkPrices[entry.Feed], entry)
		handler.StartFrom = utils.Min(handler.StartFrom, entry.BlockNumber)
	}
	for _, entries := range handler.ChainlinkPrices {
		sort.Sort(entries)
	}
}

func (handler *DBhandler) getChainlinks() {
	data := []*core.SyncAdapter{}
	err := handler.db.Raw(`(SELECT * FROM sync_adapters where type='ChainlinkPriceFeed')`).Find(&data).Error
	log.CheckFatal(err)
	for _, entry := range data {
		handler.Chainlinks[entry.Address] = entry
	}
}

func (handler *DBhandler) getUniPrices(from, to int64) bool {
	log.Infof("loaded from %d to %d", from, to)
	data := []*core.UniPoolPrices{}
	err := handler.db.Raw(`SELECT * FROM uniswap_pool_prices 
		WHERE block_num >= ? AND block_num < ? ORDER BY block_num`, from, to).Find(&data).Error
	log.CheckFatal(err)
	for _, entry := range data {
		if entry.BlockNum == 1+handler.lastBlockForToken[entry.Token] || handler.lastBlockForToken[entry.Token] == 0 {
		} else {

			startBlock := handler.lastBlockForToken[entry.Token] + 1
			for ; startBlock < entry.BlockNum; startBlock++ {
				_, uniPrices := handler.blockFeed.QueryData(startBlock, WETHAddr)
				for _, entry2 := range uniPrices {
					handler.AddUniPrices(entry2)
					handler.addUniPoolPrices(entry2)
				}
			}
		}
		handler.lastBlockForToken[entry.Token] = entry.BlockNum
		handler.addUniPoolPrices(entry)
	}
	for _, entries := range handler.UniPoolPrices {
		sort.Sort(entries)
	}
	return len(data) > 0
}

func (handler *DBhandler) addUniPoolPrices(entry *core.UniPoolPrices) {
	if entry.PriceV2 == 0 || entry.PriceV3 == 0 || entry.TwapV3 == 0 {
		_, uniPrices := handler.blockFeed.QueryData(entry.BlockNum, WETHAddr)
		log.Info(entry)
		for _, newPrices := range uniPrices {
			handler.UpdatesForUniPoolPrices = append(handler.UpdatesForUniPoolPrices, newPrices)
			if newPrices.PriceV2 == 0 || newPrices.PriceV3 == 0 || newPrices.TwapV3 == 0 {
				log.Fatal(uniPrices)
			}
		}
	}
	handler.UniPoolPrices[entry.Token] = append(handler.UniPoolPrices[entry.Token], entry)
}

func (handler *DBhandler) clearUniPrices() {
	for token := range handler.UniPoolPrices {
		handler.UniPoolPrices[token] = core.SortedUniPoolPrices{}
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
	tx := h.db.Begin()
	//
	now := time.Now()

	blocks := []*core.Block{}
	for _, block := range h.blocks {
		blocks = append(blocks, block)
	}
	err := tx.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(blocks, 100).Error
	log.CheckFatal(err)
	h.blocks = map[int64]*core.Block{}
	log.Infof("created missing uniswap_pool_prices in %f sec in blocks: %d", time.Now().Sub(now).Seconds(), len(blocks))
	now = time.Now()
	//
	if SaveRelations {
		err = tx.CreateInBatches(h.Relations, 4000).Error
		log.CheckFatal(err)
	}
	h.Relations = []*core.UniPriceAndChainlink{}

	err = tx.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(h.UpdatesForUniPoolPrices, 100).Error
	log.CheckFatal(err)
	h.UpdatesForUniPoolPrices = []*core.UniPoolPrices{}
	//
	info := tx.Commit()
	log.CheckFatal(info.Error)
	log.Infof("created uniswap_chainlink_relations in %f sec", time.Now().Sub(now).Seconds())
}

func (h *DBhandler) findRelations(adapter *core.SyncAdapter, uniPrices core.SortedUniPoolPrices, chainlinkPrices core.SortedPriceFeed) (relations int64) {
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
		var chainCurrentPrices *core.PriceFeed
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

func (h *DBhandler) compareDiff(pf *core.PriceFeed, uniPoolPrices *core.UniPoolPrices) {
	// previous pricefeed can be nil
	if pf == nil {
		return
	}
	h.AddUniPriceAndChainlinkRelation(&core.UniPriceAndChainlink{
		UniBlockNum:          uniPoolPrices.BlockNum,
		ChainlinkBlockNumber: pf.BlockNumber,
		Token:                pf.Token,
		Feed:                 pf.Feed,
	})
	if (uniPoolPrices.PriceV2Success && greaterFluctuation(uniPoolPrices.PriceV2, pf.PriceETH)) ||
		(uniPoolPrices.PriceV3Success && greaterFluctuation(uniPoolPrices.PriceV3, pf.PriceETH)) ||
		(uniPoolPrices.TwapV3Success && greaterFluctuation(uniPoolPrices.TwapV3, pf.PriceETH)) {
		// if !mdl.isNotified() {
		// mdl.uniPriceVariationNotify(pf, uniPoolPrices)
		// mdl.Details["notified"] = true
		// }
	} else {
	}
}

func (h *DBhandler) AddUniPriceAndChainlinkRelation(relation *core.UniPriceAndChainlink) {
	h.Relations = append(h.Relations, relation)
}

func greaterFluctuation(a, b float64) bool {
	return math.Abs((a-b)/a) > 0.03
}

var SaveRelations bool

func StartServer(lc fx.Lifecycle, handler *DBhandler, shutdowner fx.Shutdowner) {

	// Starting server
	lc.Append(fx.Hook{
		// To mitigate the impact of deadlocks in application startup and
		// shutdown, Fx imposes a time limit on OnStart and OnStop hooks. By
		// default, hooks have a total of 15 seconds to complete. Timeouts are
		// passed via Go's usual context.Context.
		OnStart: func(context.Context) error {
			// In production, we'd want to separate the Listen and Serve phases for
			// better error-handling.
			flag.BoolVar(&SaveRelations, "save", false, "where to save relations or not.")
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
