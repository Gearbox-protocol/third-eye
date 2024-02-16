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
	// pfversion -> token -> price feed object
	prevPriceFeeds map[schemas.PFVersion]map[string]*schemas.PriceFeed
	spotOracle     *priceFetcher.OneInchOracle
	mu             *sync.Mutex
}

func NewPrevPriceStore(client core.ClientI, tokensRepo *TokensRepo) *PrevPriceStore {
	chainId, err := client.ChainID(context.TODO())
	log.CheckFatal(err)

	store := &PrevPriceStore{
		prevPriceFeeds: map[schemas.PFVersion]map[string]*schemas.PriceFeed{},
		mu:             &sync.Mutex{},
	}
	if chainId.Int64() == 1 || chainId.Int64() == 7878 {
		store.spotOracle = ds.SetOneInchUpdater(client, tokensRepo)
	}
	return store
}

func (repo *PrevPriceStore) loadPrevPriceFeed(db *gorm.DB) {
	defer utils.Elapsed("loadPrevPriceFeed")()
	data := []*schemas.PriceFeed{}
	err := db.Raw(`SELECT * FROM 
		(SELECT distinct on(token, merged_pf_version) * FROM price_feeds ORDER BY token, merged_pf_version, block_num DESC) t ORDER BY block_num`).Find(&data).Error
	log.CheckFatal(err)
	for _, pf := range data {
		repo.isPFAdded(pf, false)
	}
}

// isUSD -> token -> feed -> price feed object
func (repo *PrevPriceStore) isPFAdded(pf *schemas.PriceFeed, save bool) bool {
	for _, pfVersion := range pf.MergedPFVersion.MergedPFVersionToList() {
		if repo.prevPriceFeeds[pfVersion] == nil {
			repo.prevPriceFeeds[pfVersion] = map[string]*schemas.PriceFeed{}
		}
		oldPF := repo.prevPriceFeeds[pfVersion][pf.Token]
		//
		price := pf.PriceBI.Convert().Int64()
		if oldPF != nil {
			// if the blocknum of new price is less than previous seenly price , ignore
			if oldPF.BlockNumber >= pf.BlockNumber && !(price == 0 || price == 100) {
				log.Fatalf("oldPF %s.\n NewPF %s.", oldPF, pf)
			}
			// same price then don't add
			if oldPF.PriceBI.Cmp(pf.PriceBI) == 0 {
				return false
			}
		}
		repo.prevPriceFeeds[pfVersion][pf.Token] = pf
	}
	return true
}

func getPrice(pfs map[string]*schemas.PriceFeed) (ans []*schemas.TokenCurrentPrice) {
	for _, pf := range pfs {
		ans = append(ans, &schemas.TokenCurrentPrice{
			PriceBI:  pf.PriceBI,
			Price:    pf.Price,
			BlockNum: pf.BlockNumber,
			Token:    pf.Token,
			PriceSrc: string(core.SOURCE_CHAINLINK),
		})
	}
	return ans
}

func (repo *PrevPriceStore) getCurrentPrice() (ans []*schemas.TokenCurrentPrice) {
	if repo.prevPriceFeeds[schemas.V3PF_MAIN] != nil {
		return getPrice(repo.prevPriceFeeds[schemas.V3PF_MAIN])
	}
	return getPrice(repo.prevPriceFeeds[schemas.V2PF])
}
func (repo *PrevPriceStore) saveCurrentPrices(client core.ClientI, tx *gorm.DB, blockNum int64) {
	// chainlink current prices to updated
	currentPricesToSync := repo.getCurrentPrice()
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
