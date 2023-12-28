package handlers

import (
	"context"
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
	prevPriceFeeds map[bool]map[string]map[string]*schemas.PriceFeed
	currentPrices  map[string]*schemas.TokenCurrentPrice
	spotOracle     *priceFetcher.OneInchOracle
	mu             *sync.Mutex
}

func NewPrevPriceStore(client core.ClientI, tokensRepo *TokensRepo) *PrevPriceStore {
	chainId, err := client.ChainID(context.TODO())
	log.CheckFatal(err)

	store := &PrevPriceStore{
		prevPriceFeeds: map[bool]map[string]map[string]*schemas.PriceFeed{},
		currentPrices:  map[string]*schemas.TokenCurrentPrice{},
		mu:             &sync.Mutex{},
	}
	if chainId.Int64() == 1 {
		store.spotOracle = ds.SetOneInchUpdater(client, tokensRepo)
	}
	return store
}

func (repo *PrevPriceStore) loadPrevPriceFeed(db *gorm.DB) {
	defer utils.Elapsed("loadPrevPriceFeed")()
	data := []*schemas.PriceFeed{}
	err := db.Raw("SELECT distinct on(token, price_in_usd)* FROM price_feeds ORDER BY token, price_in_usd, block_num DESC").Find(&data).Error
	log.CheckFatal(err)
	for _, pf := range data {
		repo.addPrevPriceFeed(pf)
		repo.addCurrentPrice(pf, false)
	}
}

// isUSD -> token -> feed -> price feed object
func (repo *PrevPriceStore) addPrevPriceFeed(pf *schemas.PriceFeed) {
	if repo.prevPriceFeeds[pf.IsPriceInUSD] == nil {
		repo.prevPriceFeeds[pf.IsPriceInUSD] = map[string]map[string]*schemas.PriceFeed{}
	}
	if repo.prevPriceFeeds[pf.IsPriceInUSD][pf.Token] == nil {
		repo.prevPriceFeeds[pf.IsPriceInUSD][pf.Token] = map[string]*schemas.PriceFeed{}
	}
	oldPF := repo.prevPriceFeeds[pf.IsPriceInUSD][pf.Token][pf.Feed]
	price := pf.PriceBI.Convert().Int64()
	if oldPF != nil && oldPF.BlockNumber >= pf.BlockNumber && !(price == 0 || price == 100) {
		log.Fatalf("oldPF %s.\n NewPF %s.", oldPF, pf)
	}
	//
	repo.prevPriceFeeds[pf.IsPriceInUSD][pf.Token][pf.Feed] = pf
}

func (repo *PrevPriceStore) canAddPF(pf *schemas.PriceFeed) bool {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if repo.prevPriceFeeds[pf.IsPriceInUSD] != nil &&
		repo.prevPriceFeeds[pf.IsPriceInUSD][pf.Token] != nil &&
		repo.prevPriceFeeds[pf.IsPriceInUSD][pf.Token][pf.Feed] != nil {
		prevPF := repo.prevPriceFeeds[pf.IsPriceInUSD][pf.Token][pf.Feed]
		if prevPF.BlockNumber >= pf.BlockNumber {
			log.Fatalf("oldPF %s.\n NewPF %s.", prevPF, pf)
		}
		if prevPF.PriceBI.Cmp(pf.PriceBI) == 0 {
			repo.addPrevPriceFeed(pf)
			return false
		}
	}
	repo.addPrevPriceFeed(pf)
	repo.addCurrentPrice(pf, true)
	return true
}

func (repo PrevPriceStore) addCurrentPrice(pf *schemas.PriceFeed, save bool) {
	if !pf.IsPriceInUSD {
		return
	}
	repo.currentPrices[pf.Token] = &schemas.TokenCurrentPrice{
		Save:     save,
		PriceBI:  pf.PriceBI,
		Price:    pf.Price,
		BlockNum: pf.BlockNumber,
		Token:    pf.Token,
		PriceSrc: string(core.SOURCE_CHAINLINK),
	}
}
func (repo *PrevPriceStore) saveCurrentPrices(client core.ClientI, tx *gorm.DB, blockNum int64) {
	// chainlink current prices to updated
	var currentPricesToSync []*schemas.TokenCurrentPrice
	for _, currentPrice := range repo.currentPrices {
		if currentPrice.Save {
			currentPrice.Save = false
			currentPricesToSync = append(currentPricesToSync, currentPrice)
		}
	}
	if len(currentPricesToSync) == 0 { // usd prices are set? only valid from v2
		// so if it's empty, we don't need to store currentPrice and nor fetch 1inch prices in usdc
		return
	}
	// spot prices to updated
	if repo.spotOracle != nil {
		calls := repo.spotOracle.GetCalls()
		results := core.MakeMultiCall(client, blockNum, false, calls)
		for token, priceBI := range repo.spotOracle.GetPrices(results, blockNum) {
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
