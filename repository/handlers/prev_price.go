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
	// pfversion -> token -> price feed object
	prevPriceFeeds map[schemas.PFVersion]map[string]*schemas.PriceFeed
	spotOracle     *priceFetcher.OneInchOracle
	mu             *sync.Mutex
	db             *gorm.DB
}

func NewPrevPriceStore(client core.ClientI, tokensRepo *TokensRepo, db *gorm.DB) *PrevPriceStore {
	chainId := core.GetChainId(client)

	store := &PrevPriceStore{
		prevPriceFeeds: map[schemas.PFVersion]map[string]*schemas.PriceFeed{},
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
		(SELECT distinct on(token, merged_pf_version) * FROM price_feeds ORDER BY token, merged_pf_version, block_num DESC) t ORDER BY block_num`).Find(&data).Error
	log.CheckFatal(err)
	for _, pf := range data {
		repo.isPFAdded(pf)
	}
}

// isUSD -> token -> feed -> price feed object
func (repo *PrevPriceStore) isPFAdded(pf *schemas.PriceFeed) bool {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if pf.Feed == "0xD478C66Baf4B9ED1185Db9886DdCA1e1403e1C4e" {
		pf.Token = "0x0000000000000000000000000000000000001234"
	}
	for _, pfVersion := range pf.MergedPFVersion.MergedPFVersionToList() {
		if repo.prevPriceFeeds[pfVersion] == nil {
			repo.prevPriceFeeds[pfVersion] = map[string]*schemas.PriceFeed{}
		}
		oldPF := repo.prevPriceFeeds[pfVersion][pf.Token]
		//
		if oldPF != nil {
			// old.pf price isn't zero, new price is zero and old.block> new.block
			if oldPF.BlockNumber >= pf.BlockNumber && // if old pf has block number more than new pf
				(pf.Price == 0 || pf.Price == 100) { // and new pf price is not 0 or 100
				log.Warnf("Only for dev.Edge case: oldPF %s.\n NewPF %s.", oldPF, pf)
			}
			// old.block > new.block but none of the price is zero
			if oldPF.BlockNumber >= pf.BlockNumber && // if old pf has block number more than new pf
				!(pf.Price == 0 || pf.Price == 100) && // and new pf price is not 0 or 100
				!(oldPF.Price == 0 || oldPF.Price == 100) { // and old pf price is not 0 or 100
				if !(pf.Token == "0x2416092f143378750bb29b79eD961ab195CcEea5" && oldPF.BlockNumber == 228761216) { // for ezETH on arbitrum
					log.Warnf("oldPF %s.\n NewPF %s.", oldPF, pf)
					return false
				}
			}
			// if the blocknum of new price is less than previous seenly price , ignore
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
			PriceSrc: string(core.SOURCE_GEARBOX),
		})
		ans = append(ans, &schemas.TokenCurrentPrice{
			PriceBI:  pf.PriceBI,
			Price:    pf.Price,
			BlockNum: pf.BlockNumber,
			Token:    pf.Token,
			PriceSrc: "chainlink",
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
func (repo *PrevPriceStore) saveCurrentPrices(client core.ClientI, tx *gorm.DB, blockNum int64, ts uint64) {
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
	currentPricesToSync := repo.getCurrentPrice()
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
