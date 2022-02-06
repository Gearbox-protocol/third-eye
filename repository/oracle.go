package repository

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
)

// for price oracle/feeds
func (repo *Repository) loadCurrentTokenOracle() {
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
	repo.tokensCurrentOracle[oracle.Token] = oracle
}

func (repo *Repository) AddTokenOracle(token, oracle, feed string, blockNum int64) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if repo.tokensCurrentOracle[token] != nil {
		currentFeed := repo.tokensCurrentOracle[token].Feed
		log.Warnf("New feed(%s) discovered at %d for token(%s) old feed: %s", feed, blockNum, token, currentFeed)
		repo.GetAdapter(currentFeed).SetBlockToDisableOn(blockNum)
	}
	// set current state of oracle for token.
	repo.addTokenCurrentOracle(
		&core.TokenOracle{Token: token, Oracle: oracle, Feed: feed, BlockNumber: blockNum},
	)
	// token oracle
	repo.setAndGetBlock(blockNum).AddTokenOracle(
		&core.TokenOracle{Token: token, Oracle: oracle, Feed: feed, BlockNumber: blockNum},
	)
}

func (repo *Repository) AddPriceFeed(blockNum int64, pf *core.PriceFeed) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.setAndGetBlock(blockNum).AddPriceFeed(pf)
}
