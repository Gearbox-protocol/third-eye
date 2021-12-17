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
