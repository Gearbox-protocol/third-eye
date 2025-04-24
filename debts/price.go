package debts

import (
	"fmt"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/address_provider"
	"gorm.io/gorm"
)

type PriceHandler struct {
	// v310
	poolToPriceOracle map[string]schemas.PriceOracleT
	cmToPool          map[string]string
	feedLastPrice     map[string]*schemas.PriceFeed
	// priceOracleToTokenFeed
	poTotokenOracle map[schemas.PriceOracleT]map[string]*schemas.TokenOracle
	repo            ds.RepositoryI
}

func NewPriceHandler(repo ds.RepositoryI) *PriceHandler {
	return &PriceHandler{
		repo: repo,
		//
		poolToPriceOracle: map[string]schemas.PriceOracleT{},
		cmToPool:          map[string]string{},
		feedLastPrice:     make(map[string]*schemas.PriceFeed),
		poTotokenOracle:   map[schemas.PriceOracleT]map[string]*schemas.TokenOracle{},
	}
}

func (eng *PriceHandler) loadPoolToPriceOracle(lastDebtSync int64, db *gorm.DB) {
	data := []schemas.Relation{}
	err := db.Raw(`(select * from relations where category='PoolOracle' and block_num <=? order by block_num)`, lastDebtSync).Find(&data).Error
	log.CheckFatal(err)
	for _, entry := range data {
		eng.poolToPriceOracle[entry.Owner] = schemas.PriceOracleT(entry.Dependent)
	}

	cms := []schemas.CreditManagerState{}
	err = db.Raw(`(select address, pool_address from credit_managers)`).Find(&cms).Error
	log.CheckFatal(err)
	for _, entry := range cms {
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
	for _, entry := range data {
		eng.AddTokenOracle(entry)
	}
}

func (eng *PriceHandler) AddTokenOracle(entry *schemas.TokenOracle) {
	if eng.poTotokenOracle[entry.PriceOracle] == nil {
		eng.poTotokenOracle[entry.PriceOracle] = map[string]*schemas.TokenOracle{}
	}
	eng.poTotokenOracle[entry.PriceOracle][entry.Token] = entry
}

func (eng *PriceHandler) AddTokenLastPrice(pf *schemas.PriceFeed) {
	eng.feedLastPrice[pf.Feed] = pf
}

// pfVersion is used only for weth on v1
func (eng *PriceHandler) GetLastPriceFeed(cm string, token string, version core.VersionType, dontFail ...bool) *schemas.PriceFeed {
	pool := eng.cmToPool[cm]
	priceOracle := eng.poolToPriceOracle[pool]
	if version.LessThan(core.NewVersion(300)) {
		addrs := eng.repo.GetAdapterAddressByName(ds.AddressProvider)
		adapter := eng.repo.GetAdapter(addrs[0])
		obj := adapter.(*address_provider.AddressProvider)
		priceOracle = obj.GetPriceOracleLegacy(version)
	}
	feed := eng.poTotokenOracle[priceOracle][token]
	if feed != nil && eng.feedLastPrice[feed.Feed] != nil { // has feed
		return eng.feedLastPrice[feed.Feed]
		// feed.Feed
	}
	// if eng.feedLastPrice[pfVersion][addr] != nil {
	// 	return eng.feedLastPrice[pfVersion][addr].PriceBI.Convert()
	// }
	//
	if len(dontFail) > 0 && dontFail[0] {
		log.Infof("Price not found for %s pfversion: %d, priceoracle:%s, feed:%s", token, version, priceOracle, feed.Feed)
		return nil
	}
	log.Fatal(fmt.Sprintf("Price not found for %s pfversion: %d, priceoracle:%s, feed:%s", token, version, priceOracle, utils.ToJson(feed)))
	return nil
}
func (eng *PriceHandler) GetLastPriceFeedByOracle(priceOracle schemas.PriceOracleT, token string, version core.VersionType, dontFail ...bool) *schemas.PriceFeed {
	feed := eng.poTotokenOracle[priceOracle][token]
	if feed != nil { // has feed
		return eng.feedLastPrice[feed.Feed]
		// feed.Feed
	}
	//
	if len(dontFail) > 0 && dontFail[0] {
		return nil
	}
	log.Fatal(fmt.Sprintf("Price not found for %s pfversion: %d", token, version))
	return nil
}
func (eng *PriceHandler) GetLastPrice(cm string, token string, version core.VersionType, dontFail ...bool) *big.Int {
	if version.Eq(1) && eng.repo.GetWETHAddr() == token { // for mainnet on ethereum
		return core.WETHPrice
	}
	a := eng.GetLastPriceFeed(cm, token, version)
	if a == nil {
		log.Fatal(cm, token, version)
	}

	return a.PriceBI.Convert()
}

func (eng *PriceHandler) requestPriceFeed(blockNum int64, client core.ClientI, retryFeed ds.QueryPriceFeedI, token string, misHFOrTValue bool, db *gorm.DB) {
	defer func() {
		// if err := recover(); err != nil {
		// 	log.Warn("err", err, "in getting query price feed in debt", token, blockNum)
		// }
	}()
	// PFFIX
	calls, isQueryable := retryFeed.GetCalls(blockNum)
	if !isQueryable {
		return
	}
	category := ""
	if misHFOrTValue {
		category = "due to missed hf/total_value"
	} else {
		category = "due to hf <1"
	}
	log.Info("getting price for ", token, "at", blockNum, category)
	results := core.MakeMultiCall(client, blockNum, false, calls)
	price := retryFeed.ProcessResult(blockNum, results, true)
	if price != nil {
		newPF := &schemas.PriceFeed{
			BlockNumber: blockNum,
			Feed:        retryFeed.GetAddress(),
			RoundId:     price.RoundId,
			PriceBI:     price.PriceBI,
			Price:       price.Price,
		}
		log.CheckFatal(db.Create(newPF).Error)
		eng.AddTokenLastPrice(newPF)
	}
}

func (eng *PriceHandler) GetPoolFromCM(cm string) string {
	return eng.cmToPool[cm]
}

func (eng *PriceHandler) init(repo ds.RepositoryI) {
	// poolToPriceOracle
	for _, pool := range repo.GetAdapterAddressByName(ds.Pool) {
		adapter := repo.GetAdapter(pool)
		state := adapter.GetUnderlyingState().(*schemas.PoolState)
		eng.poolToPriceOracle[pool] = state.PriceOracle
	}
	// cmToPool
	for _, cm := range repo.GetAdapterAddressByName(ds.CreditManager) {
		adapter := repo.GetAdapter(cm)
		state := adapter.GetUnderlyingState().(*schemas.CreditManagerState)
		eng.cmToPool[state.Address] = state.PoolAddress
	}
}
