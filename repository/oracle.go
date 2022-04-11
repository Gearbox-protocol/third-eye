package repository

import (
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
)

// for price oracle/feeds
func (repo *Repository) loadCurrentTokenOracle() {
	defer utils.Elapsed("loadCurrentTokenOracle")()
	data := []*schemas.TokenOracle{}
	query := `SELECT token_oracle.* FROM token_oracle
	JOIN (SELECT max(block_num) AS bn, token FROM token_oracle GROUP BY token) AS max_to
	ON max_to.bn = token_oracle.block_num AND max_to.token = token_oracle.token`
	err := repo.db.Raw(query).Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, oracle := range data {
		repo.addTokenCurrentOracle(oracle)
	}
}

func (repo *Repository) addTokenCurrentOracle(oracle *schemas.TokenOracle) {
	if repo.tokensCurrentOracle[oracle.Version] == nil {
		repo.tokensCurrentOracle[oracle.Version] = map[string]*schemas.TokenOracle{}
	}
	repo.tokensCurrentOracle[oracle.Version][oracle.Token] = oracle
}

func (repo *Repository) AddTokenOracle(tokenOracle *schemas.TokenOracle) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if repo.tokensCurrentOracle[tokenOracle.Version] != nil && repo.tokensCurrentOracle[tokenOracle.Version][tokenOracle.Token] != nil {
		currentFeed := repo.tokensCurrentOracle[tokenOracle.Version][tokenOracle.Token].Feed
		log.Warnf("New feed(%s) discovered at %d for token(%s) old feed: %s",
			tokenOracle.Feed, tokenOracle.BlockNumber, tokenOracle.Token, currentFeed)
		adapter := repo.GetAdapter(currentFeed)
		if currentFeed == tokenOracle.Feed {
			log.Warnf("Oracle Prev feed(%s) and new feed(%s) for token(%s) are same.",
				currentFeed, tokenOracle.Feed, tokenOracle.Token)
			return
		}
		if adapter.GetName() != ds.YearnPriceFeed {
			adapter.SetBlockToDisableOn(tokenOracle.BlockNumber)
		} else {
			repo.aggregatedFeed.DisableYearnFeed(tokenOracle.Token, currentFeed, tokenOracle.BlockNumber)
		}
	}
	// set current state of oracle for token.
	repo.addTokenCurrentOracle(
		tokenOracle,
	)
	// token oracle
	repo.setAndGetBlock(tokenOracle.BlockNumber).AddTokenOracle(
		tokenOracle,
	)
}

func (repo *Repository) AddPriceFeed(blockNum int64, pf *schemas.PriceFeed) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.setAndGetBlock(blockNum).AddPriceFeed(pf)
}

func (repo *Repository) AddYearnFeed(token, oracle string, discoveredAt int64, version int16) {
	// add token oracle for db
	// feed is also oracle address for yearn address
	// we don't relie on underlying feed
	repo.AddTokenOracle(&schemas.TokenOracle{
		Token:       token,
		Oracle:      oracle,
		Feed:        oracle,
		BlockNumber: discoveredAt,
		Version:     version})
	repo.aggregatedFeed.AddYearnFeedOrToken(token, oracle, discoveredAt, version)
}
