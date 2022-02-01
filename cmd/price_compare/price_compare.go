package main

import (
	"context"

	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/models/aggregated_block_feed"
	"github.com/Gearbox-protocol/third-eye/utils"
	"math"
	"sort"
	"time"

	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/repository"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type DBhandler struct {
	db                 *gorm.DB
	Chainlinks         map[string]*core.SyncAdapter
	ChainlinkPrices    map[string]core.SortedPriceFeed
	UniPoolPrices      map[string]core.SortedUniPoolPrices
	Relations          []*core.UniPriceAndChainlink
	lastBlockForToken  map[string]int64
	ChainlinkFeedIndex map[string]int
	StartFrom          int64
	blockFeed          *aggregated_block_feed.AggregatedBlockFeed
}

func (handler *DBhandler) populateBlockFeed(obj *aggregated_block_feed.AggregatedBlockFeed) {
	tokens := []*core.Token{}
	err := handler.db.Raw(`SELECT * FROM tokens`).Find(&tokens).Error
	tokenMap := map[string]*core.Token{}
	for _, entry := range tokens {
		tokenMap[entry.Address] = entry
	}
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
	}
	obj.populateBlockFeed(blockFeed)
	return obj
}

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
	err := handler.db.Raw(`(SELECT address FROM sync_adapters where type='ChainlinkPriceFeed')`).Find(&data).Error
	log.CheckFatal(err)
	for _, entry := range data {
		handler.Chainlinks[entry.Address] = entry
	}
}

func (handler *DBhandler) getUniPrices(from, to int64) bool {
	log.Info("loaded from %d to %d", from, to)
	data := []*core.UniPoolPrices{}
	err := handler.db.Raw(`SELECT * FROM uniswap_pool_prices 
		WHERE block_num >= ? AND block_num < to ORDER BY block_num`, from, to).Find(&data).Error
	log.CheckFatal(err)
	for _, entry := range data {
		if entry.BlockNum == 1+handler.lastBlockForToken[entry.Token] || handler.lastBlockForToken[entry.Token] == 0 {
		} else {
			weth := "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
			startBlock := handler.lastBlockForToken[entry.Token] + 1
			for ; startBlock < entry.BlockNum; startBlock++ {
				_, uniPrices := handler.blockFeed.QueryData(startBlock, weth)
				for _, entry2 := range uniPrices {
					handler.UniPoolPrices[entry2.Token] = append(handler.UniPoolPrices[entry2.Token], entry2)
				}
			}
		}
		handler.lastBlockForToken[entry.Token] = entry.BlockNum
		handler.UniPoolPrices[entry.Token] = append(handler.UniPoolPrices[entry.Token], entry)
	}
	for _, entries := range handler.UniPoolPrices {
		sort.Sort(entries)
	}
	return len(data) > 0
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
	data := LastBlock{}
	err := handler.db.Raw(`SELECT max(block_num) as block_num FROM uniswap_chainlink_relations`).Find(&data).Error
	log.CheckFatal(err)
	if data.BlockNum != 0 {
		handler.StartFrom = utils.Max(handler.StartFrom, data.BlockNum+1)
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
				log.Fatal("token parse failed", feed)
			}
			handler.findRelations(adapter, handler.UniPoolPrices[token], handler.ChainlinkPrices[feed])
		}
	}
	handler.clearUniPrices()
	handler.saveRelations()
}

func (h *DBhandler) saveRelations() {
	err := h.db.CreateInBatches(h.Relations, 50).Error
	log.CheckFatal(err)
	h.Relations = []*core.UniPriceAndChainlink{}
}

func (h *DBhandler) findRelations(adapter *core.SyncAdapter, uniPrices core.SortedUniPoolPrices, chainlinkPrices core.SortedPriceFeed) {
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
			h.compareDiff(chainPrevPrices, uniPrices[uniPriceInd])
		}
	}
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
			go func() {
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
