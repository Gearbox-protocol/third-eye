package repository

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/utils"
)

// for price oracle/feeds
func (repo *Repository) loadCurrentTokenOracle() {
	defer utils.Elapsed("loadCurrentTokenOracle")()
	data := []*core.TokenOracle{}
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

func (repo *Repository) addTokenCurrentOracle(oracle *core.TokenOracle) {
	if repo.tokensCurrentOracle[oracle.Version] == nil {
		repo.tokensCurrentOracle[oracle.Version] = map[string]*core.TokenOracle{}
	}
	repo.tokensCurrentOracle[oracle.Version][oracle.Token] = oracle
}

func (repo *Repository) AddTokenOracle(tokenOracle *core.TokenOracle) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if repo.tokensCurrentOracle[tokenOracle.Version] != nil && repo.tokensCurrentOracle[tokenOracle.Version][tokenOracle.Token] != nil {
		currentFeed := repo.tokensCurrentOracle[tokenOracle.Version][tokenOracle.Token].Feed
		log.Warnf("New feed(%s) discovered at %d for token(%s) old feed: %s",
			tokenOracle.Feed, tokenOracle.BlockNumber, tokenOracle.Token, currentFeed)
		repo.GetAdapter(currentFeed).SetBlockToDisableOn(tokenOracle.BlockNumber)
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

func (repo *Repository) AddPriceFeed(blockNum int64, pf *core.PriceFeed) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.setAndGetBlock(blockNum).AddPriceFeed(pf)
}
