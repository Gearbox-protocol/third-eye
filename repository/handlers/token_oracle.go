package handlers

import (
	"math/big"
	"sync"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/chainlink_price_feed"
	"gorm.io/gorm"
)

type TokenOracleRepo struct {
	// version  to token to oracle
	tokensCurrentOracle map[int16]map[string]*schemas.TokenOracle // done
	mu                  *sync.Mutex
	adapters            *SyncAdaptersRepo
	blocks              *BlocksRepo
	repo                ds.RepositoryI
	client              core.ClientI
	zeroPFs             map[string]bool
}

func NewTokenOracleRepo(adapters *SyncAdaptersRepo, blocks *BlocksRepo, repo ds.RepositoryI, client core.ClientI) *TokenOracleRepo {
	return &TokenOracleRepo{
		tokensCurrentOracle: make(map[int16]map[string]*schemas.TokenOracle),
		mu:                  &sync.Mutex{},
		adapters:            adapters,
		blocks:              blocks,
		repo:                repo,
		client:              client,
		zeroPFs:             map[string]bool{},
	}
}

// for price oracle/feeds
func (repo *TokenOracleRepo) LoadCurrentTokenOracle(db *gorm.DB) {
	defer utils.Elapsed("loadCurrentTokenOracle")()
	data := []*schemas.TokenOracle{}
	query := `SELECT token_oracle.* FROM token_oracle
	JOIN (SELECT max(block_num) AS bn, token FROM token_oracle GROUP BY token) AS max_to
	ON max_to.bn = token_oracle.block_num AND max_to.token = token_oracle.token`
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

func (repo *TokenOracleRepo) addTokenCurrentOracle(oracle *schemas.TokenOracle) {
	if repo.tokensCurrentOracle[oracle.Version] == nil {
		repo.tokensCurrentOracle[oracle.Version] = map[string]*schemas.TokenOracle{}
	}
	repo.tokensCurrentOracle[oracle.Version][oracle.Token] = oracle
}

func (repo *TokenOracleRepo) AddTokenOracle(newTokenOracle *schemas.TokenOracle, feedType string) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	log.Info(newTokenOracle.String())
	if repo.tokensCurrentOracle[newTokenOracle.Version] != nil &&
		repo.tokensCurrentOracle[newTokenOracle.Version][newTokenOracle.Token] != nil {
		oldTokenOracle := repo.tokensCurrentOracle[newTokenOracle.Version][newTokenOracle.Token]
		oldFeed := oldTokenOracle.Feed
		// log
		if feedType == ds.ChainlinkPriceFeed {
			if oldTokenOracle.Oracle != newTokenOracle.Oracle {
				log.Updatef("Chainlink proxy changed in gearbox protocol from (%s) to %s for token(%s) at %d",
					oldTokenOracle.Oracle, newTokenOracle.Oracle, newTokenOracle.Token, newTokenOracle.BlockNumber)
			} else if oldFeed != newTokenOracle.Feed {
				log.Updatef("Chainlink feed changed internally from (%s) to %s for token(%s) at %d",
					oldFeed, newTokenOracle.Feed, newTokenOracle.Token, newTokenOracle.BlockNumber)
			}
		} else {
			log.Updatef("%s changed from %s to %s for token(%s) at %d",
				feedType, oldFeed, newTokenOracle.Feed, newTokenOracle.Token, newTokenOracle.BlockNumber)
		}
		//
		adapter := repo.adapters.GetAdapter(oldFeed)
		if oldFeed == newTokenOracle.Feed {
			log.Warnf("Same feed(%s) added for token(%s)", newTokenOracle.Feed, newTokenOracle.Token)
			return
		}
		// disable the corresponding adapter
		if adapter == nil && repo.zeroPFs[oldFeed] {
			// no adapter is used for zeroPF as the price is always zero.
			// we can just work with 'adapter==nil' but we want to check if the adapter is null for other pricefeed by mistake. like disabled for chainlink etc.
		} else if adapter == nil {
			log.Error("Adapter not found for", oldFeed)
		} else if adapter.GetName() != ds.QueryPriceFeed {
			adapter.SetBlockToDisableOn(newTokenOracle.BlockNumber)
		} else {
			repo.adapters.AggregatedFeed.DisableYearnFeed(newTokenOracle.Token, oldFeed, newTokenOracle.BlockNumber)
		}
	}
	// set current state of oracle for token.
	repo.addTokenCurrentOracle(
		newTokenOracle,
	)
	// token oracle
	repo.blocks.SetAndGetBlock(newTokenOracle.BlockNumber).AddTokenOracle(
		newTokenOracle,
	)
}

// called from chainlink feed and price oracle
func (repo *TokenOracleRepo) AddTokenFeed(feedType, token, oracle string, discoveredAt int64, version int16) {
	switch feedType {
	case ds.CurvePF, ds.YearnPF, ds.ZeroPF, ds.AlmostZeroPF:
		// add token oracle for db
		// feed is also oracle address for yearn address
		// we don't relie on underlying feed
		repo.AddTokenOracle(&schemas.TokenOracle{
			Token:       token,
			Oracle:      oracle,
			Feed:        oracle, // feed is same as oracle
			BlockNumber: discoveredAt,
			Version:     version,
			FeedType:    feedType}, feedType)
		if feedType == ds.ZeroPF || feedType == ds.AlmostZeroPF {
			priceBI := new(big.Int)
			if ds.AlmostZeroPF == feedType {
				priceBI = big.NewInt(100)
			}
			repo.blocks.AddPriceFeed(&schemas.PriceFeed{
				BlockNumber: discoveredAt,
				Token:       token,
				Feed:        oracle,
				RoundId:     0,
				PriceBI:     (*core.BigInt)(priceBI),
				Price:       utils.GetFloat64Decimal(priceBI, 18),
			})
			repo.zeroPFs[oracle] = true // oracle and feed are same for non-chainlink price feed
		} else {
			repo.adapters.AggregatedFeed.AddFeedOrToken(token, oracle, feedType, discoveredAt, version)
		}
	case ds.ChainlinkPriceFeed:
		obj := chainlink_price_feed.NewChainlinkPriceFeed(token, oracle, discoveredAt, repo.client, repo.repo, version)
		if repo.tokensCurrentOracle[version] != nil && repo.tokensCurrentOracle[version][token] != nil {
			oldTokenOracle := repo.tokensCurrentOracle[version][token]
			if oldTokenOracle.Oracle == oracle && oldTokenOracle.Feed == obj.Address {
				log.Warnf("Same chainlinkfeed(%s) added for token(%s)", oldTokenOracle.Feed, token)
				return
			}
		}
		repo.AddTokenOracle(&schemas.TokenOracle{
			Token:       token,
			Oracle:      oracle,
			Feed:        obj.Address,
			BlockNumber: discoveredAt,
			Version:     version,
			FeedType:    feedType}, feedType)
		repo.adapters.AddSyncAdapter(obj)
	default:
		log.Fatal(feedType, "not handled")
	}
}

func (repo *TokenOracleRepo) GetTokenOracles() map[int16]map[string]*schemas.TokenOracle {
	return repo.tokensCurrentOracle
}
