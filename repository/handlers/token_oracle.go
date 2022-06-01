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
}

func NewTokenOracleRepo(adapters *SyncAdaptersRepo, blocks *BlocksRepo, repo ds.RepositoryI, client core.ClientI) *TokenOracleRepo {
	return &TokenOracleRepo{
		tokensCurrentOracle: make(map[int16]map[string]*schemas.TokenOracle),
		mu:                  &sync.Mutex{},
		adapters:            adapters,
		blocks:              blocks,
		repo:                repo,
		client:              client,
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
	for _, oracle := range data {
		repo.addTokenCurrentOracle(oracle)
	}
}

func (repo *TokenOracleRepo) addTokenCurrentOracle(oracle *schemas.TokenOracle) {
	if repo.tokensCurrentOracle[oracle.Version] == nil {
		repo.tokensCurrentOracle[oracle.Version] = map[string]*schemas.TokenOracle{}
	}
	repo.tokensCurrentOracle[oracle.Version][oracle.Token] = oracle
}

func (repo *TokenOracleRepo) AddTokenOracle(tokenOracle *schemas.TokenOracle) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if repo.tokensCurrentOracle[tokenOracle.Version] != nil && repo.tokensCurrentOracle[tokenOracle.Version][tokenOracle.Token] != nil {
		currentFeed := repo.tokensCurrentOracle[tokenOracle.Version][tokenOracle.Token].Feed
		log.Warnf("New feed(%s) discovered at %d for token(%s) old feed: %s",
			tokenOracle.Feed, tokenOracle.BlockNumber, tokenOracle.Token, currentFeed)
		adapter := repo.adapters.GetAdapter(currentFeed)
		if currentFeed == tokenOracle.Feed {
			log.Warnf("Oracle Prev feed(%s) and new feed(%s) for token(%s) are same.",
				currentFeed, tokenOracle.Feed, tokenOracle.Token)
			return
		}
		if adapter.GetName() != ds.QueryPriceFeed {
			adapter.SetBlockToDisableOn(tokenOracle.BlockNumber)
		} else {
			repo.adapters.AggregatedFeed.DisableYearnFeed(tokenOracle.Token, currentFeed, tokenOracle.BlockNumber)
		}
	}
	// set current state of oracle for token.
	repo.addTokenCurrentOracle(
		tokenOracle,
	)
	// token oracle
	repo.blocks.SetAndGetBlock(tokenOracle.BlockNumber).AddTokenOracle(
		tokenOracle,
	)
}

// called from chainlink feed and price oracle
func (repo *TokenOracleRepo) AddTokenFeed(feedType, token, oracle string, discoveredAt int64, version int16) {
	switch feedType {
	case ds.CurvePF, ds.YearnPF, ds.ZeroPF:
		// add token oracle for db
		// feed is also oracle address for yearn address
		// we don't relie on underlying feed
		repo.AddTokenOracle(&schemas.TokenOracle{
			Token:       token,
			Oracle:      oracle,
			Feed:        oracle, // feed is same as oracle
			BlockNumber: discoveredAt,
			Version:     version})
		if feedType != ds.ZeroPF {
			repo.adapters.AggregatedFeed.AddFeedOrToken(token, oracle, feedType, discoveredAt, version)
		} else {
			repo.blocks.AddPriceFeed(&schemas.PriceFeed{
				BlockNumber: discoveredAt,
				Token:       token,
				Feed:        oracle,
				RoundId:     0,
				PriceBI:     (*core.BigInt)(new(big.Int)),
				Price:       0,
			})
		}
	case ds.ChainlinkPF:
		obj := chainlink_price_feed.NewChainlinkPriceFeed(token, oracle, discoveredAt, repo.client, repo.repo, version)
		if repo.tokensCurrentOracle[version] != nil && repo.tokensCurrentOracle[version][token] != nil {
			oldTokenOracle := repo.tokensCurrentOracle[version][token]
			if oldTokenOracle.Oracle == oracle && oldTokenOracle.Feed == obj.Address {
				log.Warnf("Oracle Prev feed(%s) and new feed(%s) for token(%s) are same.",
					oldTokenOracle.Feed, obj.Address, token)
				return
			}
		}
		repo.AddTokenOracle(&schemas.TokenOracle{
			Token:       token,
			Oracle:      oracle,
			Feed:        obj.Address,
			BlockNumber: discoveredAt,
			Version:     version})
		repo.adapters.AddSyncAdapter(obj)
	}
}

func (repo *TokenOracleRepo) GetTokenOracles() map[int16]map[string]*schemas.TokenOracle {
	return repo.tokensCurrentOracle
}
