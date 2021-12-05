package repository

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
)

// For token with symbol/decimals
func (repo *Repository) AddToken(addr string) *core.Token {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if repo.tokens[addr] == nil {
		repo.tokens[addr] = core.NewToken(addr, repo.client)
	}
	return repo.tokens[addr]
}

func (repo *Repository) GetToken(addr string) *core.Token {
	token := repo.tokens[addr]
	if token == nil {
		log.Info("token not found for address", addr)
		return repo.AddToken(addr)
	}
	return token
}

func (repo *Repository) loadToken() {
	data := []*core.Token{}
	err := repo.db.Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, token := range data {
		repo.AddTokenObj(token)
	}
}

func (repo *Repository) AddTokenObj(t *core.Token) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if repo.tokens[t.Address] == nil {
		repo.tokens[t.Address] = t
	}
}

// for credit filter
func (repo *Repository) AddAllowedProtocol(p *core.Protocol) {
	repo.blocks[p.BlockNumber].AddAllowedProtocol(p)
}

func (repo *Repository) AddAllowedToken(atoken *core.AllowedToken) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.allowedTokens = append(repo.allowedTokens, atoken)
}

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
		repo.AddTokenCurrentOracle(oracle)
	}
}

func (repo *Repository) AddTokenCurrentOracle(oracle *core.TokenOracle) {
	repo.tokensCurrentOracle[oracle.Token] = oracle
}

func (repo *Repository) AddTokenOracle(token, oracle, feed string, blockNum int64) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if repo.tokensCurrentOracle[token] != nil {
		currentFeed := repo.tokensCurrentOracle[token].Feed
		log.Warnf("New feed(%s) discovered at %d for token(%s) old feed: %s", feed, blockNum, token, currentFeed)
		repo.kit.GetAdapter(currentFeed).SetBlockToDisableOn(blockNum)
	}
	// set current state of oracle for token.
	repo.AddTokenCurrentOracle(
		&core.TokenOracle{Token: token, Oracle: oracle, Feed: feed, BlockNumber: blockNum},
	)
	// token oracle
	repo.blocks[blockNum].AddTokenOracle(
		&core.TokenOracle{Token: token, Oracle: oracle, Feed: feed, BlockNumber: blockNum},
	)
}

func (repo *Repository) AddPriceFeed(blockNum int64, pf *core.PriceFeed) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.blocks[blockNum].AddPriceFeed(pf)
}
