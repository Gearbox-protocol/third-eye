package repository

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
)

// For token with symbol/decimals
func (repo *Repository) AddToken(addr string) *core.Token {
	token, err := repo.addToken(addr)
	if err != nil {
		log.Fatal("Adding token failed for", token)
	}
	return token
}

func (repo *Repository) addToken(addr string) (*core.Token, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if repo.tokens[addr] == nil {
		var err error
		repo.tokens[addr], err = core.NewToken(addr, repo.client)
		if err != nil {
			return nil, err
		}
	}
	return repo.tokens[addr], nil
}

func (repo *Repository) GetToken(addr string) *core.Token {
	token, err := repo.getTokenWithError(addr)
	log.CheckFatal(err)
	return token
}

func (repo *Repository) getTokenWithError(addr string) (*core.Token, error) {
	token := repo.tokens[addr]
	if token == nil {
		return repo.addToken(addr)
	}
	return token, nil
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

func (repo *Repository) loadAllowedTokensState() {
	data := []*core.AllowedToken{}
	err := repo.db.Raw("SELECT * FROM allowed_tokens where disable_block = 0").Find(&data).Error
	log.CheckFatal(err)
	for _, entry := range data {
		repo.addAllowedTokenState(entry)
	}
}

func (repo *Repository) addAllowedTokenState(entry *core.AllowedToken) {
	tokensForCM := repo.allowedTokens[entry.CreditManager]
	if tokensForCM == nil {
		repo.allowedTokens[entry.CreditManager] = make(map[string]*core.AllowedToken)
		tokensForCM = repo.allowedTokens[entry.CreditManager]
	}
	if tokensForCM[entry.Token] != nil {
		log.Warnf("Token already enabled: new %#v, previous entry: %#v", entry, tokensForCM[entry.Token])
	} else {
		tokensForCM[entry.Token] = entry
	}
}
