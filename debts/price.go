package debts

import (
	"fmt"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"gorm.io/gorm"
)


type PriceHandler struct {
	// v310
	poolToPriceOracle map[string]schemas.PriceOracleT
	cmToPool map[string]string
	feedLastPrice map[string]*schemas.PriceFeed
	// priceOracleToTokenFeed
	poTotokenOracle map[schemas.PriceOracleT]map[string]*schemas.TokenOracle
}


func NewPriceHandler() *PriceHandler {
	return &PriceHandler{
		//
		poolToPriceOracle: map[string]schemas.PriceOracleT{},
		cmToPool: map[string]string{},
		feedLastPrice:         make(map[string]*schemas.PriceFeed),
	}
}



func (eng *PriceHandler) loadPoolToPriceOracle(lastDebtSync int64, db *gorm.DB) {
	data := []schemas.Relation{}
	err := db.Raw(`(select * from relations where type='PoolOracle' where block_num <=? order by block_num)`,lastDebtSync).Find(&data).Error
	log.CheckFatal(err)
	for _, entry:= range data {
		eng.poolToPriceOracle[entry.Owner] = schemas.PriceOracleT(entry.Dependent)
	}

	cms := []schemas.CreditManagerState{}
	err = db.Raw(`(select address, pool from credit_managers)`).Find(&cms).Error
	log.CheckFatal(err)
	for _, entry:= range cms {
		eng.cmToPool[entry.Address] = entry.PoolAddress
	}
}


func (eng *PriceHandler) load(lastDebtSync int64, db *gorm.DB) {
	eng.loadPoolToPriceOracle(lastDebtSync, db)
	eng.loadTokenLastPrice(lastDebtSync, db)
	eng.loadTokenOracle(lastDebtSync, db)

}
// token price from feeds
func (eng *PriceHandler) loadTokenLastPrice(lastDebtSync int64, db *gorm.DB) {
	defer utils.Elapsed("Debt(loadTokenLastPrice)")()
	data := []*schemas.PriceFeed{}
	query := `select * from (SELECT distinct on (feed) * FROM price_feeds WHERE block_num <= ? ORDER BY feed, block_num DESC) t order by block_num;`
	err := db.Raw(query, lastDebtSync).Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, tokenPrice := range data {
		eng.AddTokenLastPrice(tokenPrice)
	}
}

func (eng *PriceHandler) loadTokenOracle(lastDebtSync int64, db *gorm.DB) {
	data := []*schemas.TokenOracle{}
	err := db.Raw(`select * from token_oracle where block_num < ? and reserve='f' order by block_num`, lastDebtSync).Find(&data).Error
	log.CheckFatal(err)
	for _, entry:= range data {
		if eng.poTotokenOracle[entry.PriceOracle] == nil {
			eng.poTotokenOracle[entry.PriceOracle] = map[string]*schemas.TokenOracle{}
		}
		eng.poTotokenOracle[entry.PriceOracle][entry.Token] = entry
	}
}

func (eng *PriceHandler) AddTokenLastPrice(pf *schemas.PriceFeed) {
	eng.feedLastPrice[pf.Feed] = pf
}

// pfVersion is used only for weth on v1
func (eng *PriceHandler) GetLastPriceFeed(cm string ,token string, version core.VersionType, dontFail ...bool) *schemas.PriceFeed {
	pool := eng.cmToPool[cm]
	priceOracle := eng.poolToPriceOracle[pool]
	feed := eng.poTotokenOracle[priceOracle][token]
	if feed != nil { // has feed
		return eng.feedLastPrice[feed.Feed]
		// feed.Feed		
	}
	// if eng.feedLastPrice[pfVersion][addr] != nil {
	// 	return eng.feedLastPrice[pfVersion][addr].PriceBI.Convert()
	// }
	//
	if len(dontFail) > 0 && dontFail[0] {
		return nil
	}
	log.Fatal(fmt.Sprintf("Price not found for %s pfversion: %d", token, version))
	return nil
}
func (eng *PriceHandler) GetLastPriceFeedByOracle(priceOracle schemas.PriceOracleT ,token string, version core.VersionType, dontFail ...bool) *schemas.PriceFeed {
	feed := eng.poTotokenOracle[priceOracle][token]
	if feed != nil { // has feed
		return eng.feedLastPrice[feed.Feed]
		// feed.Feed		
	}
	// if eng.feedLastPrice[pfVersion][addr] != nil {
	// 	return eng.feedLastPrice[pfVersion][addr].PriceBI.Convert()
	// }
	//
	if len(dontFail) > 0 && dontFail[0] {
		return nil
	}
	log.Fatal(fmt.Sprintf("Price not found for %s pfversion: %d", token, version))
	return nil
}
func (eng *PriceHandler) GetLastPrice(cm string ,token string, version core.VersionType, dontFail ...bool) *big.Int {
	if version.Eq(1) && "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" == token { // for mainnet on ethereum
		return core.WETHPrice
	}
	return eng.GetLastPriceFeed(cm, token, version).PriceBI.Convert()
}

func (eng *PriceHandler) requestPriceFeed(blockNum int64, client core.ClientI, retryFeed ds.QueryPriceFeedI, token string) {
	// defer func() {
	// 	if err:= recover(); err != nil {
	// 		log.Warn("err", err, "in getting yearn price feed in debt", feed, token, blockNum, pfVersion)
	// 	}
	// }()
	// PFFIX
	calls, isQueryable := retryFeed.GetCalls(blockNum)
	if !isQueryable {
		return
	}
	log.Info("getting price for ", token, "at", blockNum)
	results := core.MakeMultiCall(client, blockNum, false, calls)
	price := retryFeed.ProcessResult(blockNum, results, true)
	eng.AddTokenLastPrice(&schemas.PriceFeed{
		BlockNumber:     blockNum,
		Feed:            retryFeed.GetAddress(),
		RoundId:         price.RoundId,
		PriceBI:         price.PriceBI,
		Price:           price.Price,
	})
}