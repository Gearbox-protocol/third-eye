package repository

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
)

// For token with symbol/decimals
func (repo *Repository) AddToken(addr string) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if repo.tokens[addr] == nil {
		repo.tokens[addr] = core.NewToken(addr, repo.client)
	}
}

func (repo *Repository) GetToken(addr string) *core.Token {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	token := repo.tokens[addr]
	if token == nil {
		log.Fatal("token not found for address", addr)
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
func (repo *Repository) AddTokenOracle(token, oracle string, blockNum int64) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.blocks[blockNum].AddTokenOracle(
		&core.TokenOracle{Token: token, Oracle: oracle, BlockNumber: blockNum},
	)
}

func (repo *Repository) AddPriceFeed(blockNum int64, pf *core.PriceFeed) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.blocks[blockNum].AddPriceFeed(pf)
}
