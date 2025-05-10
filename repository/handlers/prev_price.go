package handlers

import (
	"sync"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg/priceFetcher"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PrevPriceStore struct {
	// for prevently duplicate query price feed already with same price for a token
	// token to feed
	// feed -> price feed object
	prevPriceFeeds map[string]*schemas.PriceFeed
	spotOracle     *priceFetcher.OneInchOracle
	mu             *sync.Mutex
	db             *gorm.DB
}

func NewPrevPriceStore(client core.ClientI, tokensRepo *TokensRepo, db *gorm.DB) *PrevPriceStore {
	chainId := core.GetChainId(client)

	store := &PrevPriceStore{
		prevPriceFeeds: map[string]*schemas.PriceFeed{},
		mu:             &sync.Mutex{},
		db:             db,
	}
	if !utils.Contains([]log.NETWORK{"TEST", "SONIC"}, log.GetBaseNet(chainId)) {
		store.spotOracle = ds.SetOneInchUpdater(client, tokensRepo)
	}
	return store
}

func (repo *PrevPriceStore) loadPrevPriceFeed(db *gorm.DB) {
	defer utils.Elapsed("loadPrevPriceFeed")()
	data := []*schemas.PriceFeed{}
	err := db.Raw(`SELECT * FROM 
		(SELECT distinct on(feed) * FROM price_feeds ORDER BY feed, block_num DESC) t ORDER BY block_num`).Find(&data).Error
	log.CheckFatal(err)
	for _, pf := range data {
		_ = repo.isPFAdded(pf)
	}
}

// isUSD -> token -> feed -> price feed object
func (repo *PrevPriceStore) isPFAdded(pf *schemas.PriceFeed) bool {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	oldPF := repo.prevPriceFeeds[pf.Feed]
	if oldPF != nil {
		if oldPF.BlockNumber >= pf.BlockNumber {
			log.Warnf("oldPF %s.\n NewPF %s.", oldPF, pf)
			return false
		}
		if oldPF.PriceBI.Cmp(pf.PriceBI) == 0 {
			return false
		}
	}
	repo.prevPriceFeeds[pf.Feed] = pf
	return true
}

func (repo *PrevPriceStore) getCurrentPrice(tokenOracleToFeed map[schemas.PriceOracleT]map[string]*schemas.TokenOracle) (ans []*schemas.TokenCurrentPrice) {
	for priceOracle, tokenToFeed := range tokenOracleToFeed {
		for _, entry := range tokenToFeed {
			pf := repo.prevPriceFeeds[entry.Feed]
			if pf == nil {
				continue
			}
			ans = append(ans, &schemas.TokenCurrentPrice{
				PriceBI:     pf.PriceBI,
				Price:       pf.Price,
				BlockNum:    pf.BlockNumber,
				Token:       entry.Token,
				PriceOracle: priceOracle,
				PriceSrc:    string(core.SOURCE_GEARBOX),
			})
			ans = append(ans, &schemas.TokenCurrentPrice{
				PriceBI:     pf.PriceBI,
				Price:       pf.Price,
				BlockNum:    pf.BlockNumber,
				Token:       entry.Token,
				PriceOracle: priceOracle,
				PriceSrc:    "chainlink",
			})
		}
	}
	return
}
func (repo *PrevPriceStore) SaveCurrentPrices(client core.ClientI, tx *gorm.DB, blockNum int64, ts uint64, tokenToFeed map[schemas.PriceOracleT]map[string]*schemas.TokenOracle) {
	{
		a := struct {
			BlockNum int64 `gorm:"column:block_num"`
		}{}
		err := repo.db.Raw(`select max(id) block_num from blocks`).Find(&a).Error
		log.CheckFatal(err)
		if a.BlockNum >= blockNum {
			return
		}
	}
	// chainlink current prices to updated
	currentPricesToSync := repo.getCurrentPrice(tokenToFeed)
	if len(currentPricesToSync) == 0 { // usd prices are set? only valid from v2
		// so if it's empty, we don't need to store currentPrice and nor fetch 1inch prices in usdc
		return
	}
	// if log.GetBaseNet(core.GetChainId(client)) == "ARBITRUM" {
	// 	spot := []*schemas.TokenCurrentPrice{}
	// 	for _, price := range currentPricesToSync {
	// 		spot = append(spot, &schemas.TokenCurrentPrice{
	// 			PriceBI:  price.PriceBI,
	// 			Price:    utils.GetFloat64Decimal(price.PriceBI.Convert(), 8),
	// 			BlockNum: blockNum,
	// 			Token:    price.Token,
	// 			PriceSrc: string(core.SOURCE_SPOT),
	// 		})
	// 	}
	// 	currentPricesToSync = append(currentPricesToSync, spot...)
	// }
	// spot prices to updated
	if repo.spotOracle != nil {
		calls := repo.spotOracle.GetCalls()
		results := core.MakeMultiCall(client, blockNum, false, calls)
		for token, priceBI := range repo.spotOracle.GetPrices(results, blockNum, ts) {
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
