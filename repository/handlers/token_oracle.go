package handlers

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/address_provider"
	"github.com/Gearbox-protocol/third-eye/models/chainlink_price_feed"
	"github.com/Gearbox-protocol/third-eye/models/composite_chainlink"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type tokenAndPriceOracle struct {
	_token       string
	_priceOracle schemas.PriceOracleT
}

func (z tokenAndPriceOracle) MarshalText() (text []byte, err error) {
	return []byte(fmt.Sprintf("{\"token\":\"%s\", \"priceOracle\":\"%s\"}", z._token, z._priceOracle)), nil
}

type inner struct {
	// priceOracle  to token to oracle
	tokensCurrentOracle map[schemas.PriceOracleT]map[string]*schemas.TokenOracle // done
	// feed to token to true
	feedToTokens   map[string]map[tokenAndPriceOracle]*schemas.TokenOracle
	disabledTokens []*schemas.TokenOracle
}

func (repo *inner) TokensValidAtBlock(feed string, blockNum int64) (valid []*schemas.TokenOracle) {
	for _, entry := range repo.feedToTokens[feed] {
		if entry.BlockNumber <= blockNum && (entry.DisabledAt == 0 || entry.DisabledAt > blockNum) {
			valid = append(valid, entry)
		}
	}
	return valid
}

func (repo *inner) TokenAddrsValidAtBlock(feed string, blockNum int64) (addr map[string]bool) {
	valid := repo.TokensValidAtBlock(feed, blockNum)
	addr = map[string]bool{}
	for _, entry := range valid {
		addr[entry.Token] = true
	}
	return addr
}

func (repo *inner) addTokenCurrentOracle(oracle *schemas.TokenOracle) {
	if repo.tokensCurrentOracle[oracle.PriceOracle] == nil {
		repo.tokensCurrentOracle[oracle.PriceOracle] = map[string]*schemas.TokenOracle{}
	}
	repo.tokensCurrentOracle[oracle.PriceOracle][oracle.Token] = oracle
	if repo.feedToTokens[oracle.Feed] == nil {
		repo.feedToTokens[oracle.Feed] = map[tokenAndPriceOracle]*schemas.TokenOracle{}
	}
	repo.feedToTokens[oracle.Feed][tokenAndPriceOracle{
		_token:       oracle.Token,
		_priceOracle: oracle.PriceOracle,
	}] = oracle
}

func (repo *TokenOracleRepo) Save(tx *gorm.DB, blockNum int64) {
	var v2CloseBlock int64 = 19752044
	if blockNum > v2CloseBlock { // disable v1 and v2
		addrs := repo.adapters.GetAdapterAddressByName(ds.AddressProvider)
		adapter := repo.adapters.GetAdapter(addrs[0]).(*address_provider.AddressProvider)
		for _, v := range []int16{2} { // 29 v1 accounts still open
			po := adapter.GetPriceOracleLegacy(core.NewVersion(v))
			for _, d := range repo.tokensCurrentOracle[po] {
				d.DisabledAt = v2CloseBlock
				repo.disabledTokens = append(repo.disabledTokens, d)
			}
		}
	}
	err := tx.Clauses(clause.OnConflict{UpdateAll: true}).CreateInBatches(repo.disabledTokens, 50).Error
	log.CheckFatal(err)
	repo.disabledTokens = nil
	//
	repo.blocks.prevStore.SaveCurrentPrices(repo.client, tx, blockNum, repo.blocks.SetAndGetBlock(blockNum).Timestamp, repo.tokensCurrentOracle)
}

func newinner() inner {
	return inner{
		tokensCurrentOracle: map[schemas.PriceOracleT]map[string]*schemas.TokenOracle{},
		feedToTokens:        map[string]map[tokenAndPriceOracle]*schemas.TokenOracle{},
	}
}

type TokenOracleRepo struct {
	inner
	mu       *sync.Mutex
	adapters *SyncAdaptersRepo
	blocks   *BlocksRepo
	repo     ds.RepositoryI
	client   core.ClientI
	zeroPFs  map[string]bool
}

func NewTokenOracleRepo(adapters *SyncAdaptersRepo, blocks *BlocksRepo, repo ds.RepositoryI, client core.ClientI) *TokenOracleRepo {
	return &TokenOracleRepo{
		inner:    newinner(),
		mu:       &sync.Mutex{},
		adapters: adapters,
		blocks:   blocks,
		repo:     repo,
		client:   client,
		zeroPFs:  map[string]bool{},
	}
}

// for price oracle/feeds
func (repo *TokenOracleRepo) LoadCurrentTokenOracle(db *gorm.DB) {
	defer utils.Elapsed("loadCurrentTokenOracle")()
	data := []*schemas.TokenOracle{}
	query := `SELECT distinct on (token, version, reserve) * FROM token_oracle order by token, version, reserve, block_num desc;`
	err := db.Raw(query).Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, tokenOracle := range data {
		repo.addTokenCurrentOracle(tokenOracle)
	}
	repo.loadZeroPFs(db)
}

func (repo *TokenOracleRepo) loadZeroPFs(db *gorm.DB) {
	data := []schemas.TokenOracle{}
	err := db.Where("feed_type in ('ZeroPF', 'AlmostZeroPF')").Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, tokenOracle := range data {
		repo.zeroPFs[tokenOracle.Feed] = true
	}
}

// if same feed is active for current token and version
func (repo *TokenOracleRepo) alreadyActiveFeedForToken(newTokenOracle *schemas.TokenOracle) bool {
	feedType := newTokenOracle.FeedType
	//
	if repo.tokensCurrentOracle[newTokenOracle.PriceOracle] != nil &&
		repo.tokensCurrentOracle[newTokenOracle.PriceOracle][newTokenOracle.Token] != nil {
		oldTokenOracle := repo.tokensCurrentOracle[newTokenOracle.PriceOracle][newTokenOracle.Token]

		if oldTokenOracle.Feed == newTokenOracle.Feed {
			log.Debugf("Same %s(%s) added for token(%s)", feedType, newTokenOracle.Feed, newTokenOracle.Token)
			return true
		}
	}
	return false
}

func (repo *inner) removeTokenLastOracle(newTokenOracle *schemas.TokenOracle) {
	if repo.tokensCurrentOracle[newTokenOracle.PriceOracle] != nil &&
		repo.tokensCurrentOracle[newTokenOracle.PriceOracle][newTokenOracle.Token] != nil {
		oldTokenOracle := repo.tokensCurrentOracle[newTokenOracle.PriceOracle][newTokenOracle.Token]
		// oldFeed := oldTokenOracle.Feed
		// delete(repo.feedToTokens[oldFeed], tokenAndPriceOracle{
		// 	_token:       newTokenOracle.Token,
		// 	_priceOracle: newTokenOracle.PriceOracle,
		// })
		oldTokenOracle.DisabledAt = newTokenOracle.BlockNumber
		repo.disabledTokens = append(repo.disabledTokens, oldTokenOracle)
	}
}

func (repo *TokenOracleRepo) DirectlyAddTokenOracleTest(newTokenOracle *schemas.TokenOracle) {
	repo.addTokenCurrentOracle(
		newTokenOracle,
	)
}
func (repo *TokenOracleRepo) disablePrevAdapterAndAddNewTokenOracle(newTokenOracle *schemas.TokenOracle) {
	repo.removeTokenLastOracle(
		newTokenOracle,
	)
	// set current state of oracle for token.
	repo.addTokenCurrentOracle(
		newTokenOracle,
	)
	// token oracle
	repo.blocks.SetAndGetBlock(newTokenOracle.BlockNumber).AddTokenOracle(
		newTokenOracle,
	)
}

func (repo *TokenOracleRepo) DirectlyAddTokenOracle(newTokenOracle *schemas.TokenOracle) {
	repo.disablePrevAdapterAndAddNewTokenOracle(newTokenOracle)
}

// called from chainlink feed and price oracle
func (repo *TokenOracleRepo) AddNewPriceOracleEvent(newTokenOracle *schemas.TokenOracle, forChainlinkNewFeed ...bool) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	// SPECIAL CASE
	// for token 0xf568F6C71aE0439B8d3FFD60Ceba9B1DcB5819bF on mainnet, while adding blocker token for v1
	// two feeds where added for same block(15371802)
	// ignore 0xBc1c306920309F795fB5A740083eCBf5057349e9 at log 202
	// use 0xAaaF70b91877966900F0EfC0f2E7296e4F86B119 at log 212
	if newTokenOracle.Feed == "0xBc1c306920309F795fB5A740083eCBf5057349e9" && newTokenOracle.BlockNumber == 15371802 {
		return
	}
	switch newTokenOracle.FeedType {
	case ds.ZeroPF, ds.AlmostZeroPF,
		ds.CurvePF, ds.SingleAssetPF, ds.YearnPF, ds.RedStonePF, ds.CompositeRedStonePF:
		if repo.alreadyActiveFeedForToken(newTokenOracle) {
			return
		}
		repo.disablePrevAdapterAndAddNewTokenOracle(newTokenOracle)
		//
		if newTokenOracle.FeedType == ds.ZeroPF || newTokenOracle.FeedType == ds.AlmostZeroPF {
			priceBI := new(big.Int)
			if ds.AlmostZeroPF == newTokenOracle.FeedType {
				priceBI = big.NewInt(100)
			}
			repo.blocks.AddPriceFeed(&schemas.PriceFeed{
				BlockNumber: newTokenOracle.BlockNumber,
				Feed:        newTokenOracle.Oracle,
				RoundId:     0,
				PriceBI:     (*core.BigInt)(priceBI),
				Price:       utils.GetFloat64Decimal(priceBI, newTokenOracle.Version.Decimals()),
			})
			repo.zeroPFs[newTokenOracle.Oracle] = true // oracle and feed are same for non-chainlink price feed
		} else {
			repo.adapters.GetAggregatedFeed().AddFeedOrToken(
				newTokenOracle.Token,
				newTokenOracle.Feed,
				newTokenOracle.FeedType,
				newTokenOracle.BlockNumber,
				newTokenOracle.Version,
				newTokenOracle.Underlyings,
			)
		}
	case ds.ChainlinkPriceFeed:
		obj := chainlink_price_feed.NewChainlinkPriceFeed(
			repo.client, repo.repo,
			newTokenOracle.Oracle,
			newTokenOracle.BlockNumber,
			newTokenOracle.Version,
			forChainlinkNewFeed...,
		)
		newTokenOracle.Feed = obj.Address
		//
		if repo.alreadyActiveFeedForToken(newTokenOracle) {
			return
		}
		repo.disablePrevAdapterAndAddNewTokenOracle(newTokenOracle)
		if newTokenOracle.Reserve {
			return
		}
		//
		if adapter := repo.adapters.GetAdapter(obj.Address); adapter != nil {
			return
		}
		// SPECIAL CASE
		// on goerli, there are two v2 priceoracles added
		// on first priceoracle cvx token is 0x9683a59Ad8D7B5ac3eD01e4cff1D1A2a51A8f1c0
		// and on second priceoracle it is 0x6D75eb70402CF06a0cB5B8fdc1836dAe29702B17
		// and both uses the same chainlink price feed
		// REASON: we only support 1 chainlink per token, two different tokens can't share chainlink feed
		// so ignore the first cvx token as it is not used anywhere
		if newTokenOracle.Token != "0x9683a59Ad8D7B5ac3eD01e4cff1D1A2a51A8f1c0" {
			repo.adapters.AddSyncAdapter(obj)
		}
	case ds.CompositeChainlinkPF:
		obj := composite_chainlink.NewCompositeChainlinkPF(
			newTokenOracle.Token,
			newTokenOracle.Oracle,
			newTokenOracle.BlockNumber,
			repo.client, repo.repo,
			newTokenOracle.Version,
			newTokenOracle.PriceOracle,
		)
		//
		if repo.alreadyActiveFeedForToken(newTokenOracle) {
			return
		}
		repo.disablePrevAdapterAndAddNewTokenOracle(newTokenOracle)
		if newTokenOracle.Reserve {
			return
		}

		if adapter := repo.adapters.GetAdapter(obj.Address); adapter != nil {
			return
		}
		repo.adapters.AddSyncAdapter(obj)
	default:
		log.Fatal(newTokenOracle.FeedType, "not handled")
	}
}

func (repo *TokenOracleRepo) GetTokenOracles() map[schemas.PriceOracleT]map[string]*schemas.TokenOracle {
	return repo.tokensCurrentOracle
}

// if returned value is nil, it means that token oracle hasn't been added yet.
// func (repo *TokenOracleRepo) GetOracleForV2Token(token string) *schemas.TokenOracle {
// 	obj := repo.tokensCurrentOracle[core.NewVersion(2)][token]
// 	return obj
// }
