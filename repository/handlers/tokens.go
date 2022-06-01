package handlers

import (
	"sync"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TokensRepo struct {
	// blocks/token
	tokens        map[string]*schemas.Token
	WETHAddr      string
	USDCAddr      string
	GearTokenAddr string
	dieselTokens  map[string]*schemas.UTokenAndPool
	//
	mu     *sync.Mutex
	client core.ClientI
}

func NewTokensRepo(client core.ClientI) *TokensRepo {
	return &TokensRepo{
		mu:     &sync.Mutex{},
		client: client,
		//
		tokens: map[string]*schemas.Token{},
		// for getting the diesel tokens
		dieselTokens: make(map[string]*schemas.UTokenAndPool),
	}
}

// load/save tokens
func (repo *TokensRepo) LoadTokens(db *gorm.DB) {
	defer utils.Elapsed("loadToken")()
	data := []*schemas.Token{}
	err := db.Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, token := range data {
		repo.AddTokenObj(token)
	}
}

func (repo *TokensRepo) Save(tx *gorm.DB) {
	tokens := make([]*schemas.Token, 0, len(repo.tokens))
	for _, token := range repo.tokens {
		tokens = append(tokens, token)
	}
	err := tx.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(tokens, 50).Error
	log.CheckFatal(err)
}

// not to be called directly
// only exposed for testing framework
func (repo *TokensRepo) addTokenObj(t *schemas.Token) {
	// set usdc addr in repo
	if t.Symbol == "USDC" {
		repo.USDCAddr = t.Address
	}
	if t.Symbol == "WETH" {
		repo.WETHAddr = t.Address
	}
	if repo.tokens[t.Address] == nil {
		repo.tokens[t.Address] = t
	}
}

func (repo *TokensRepo) addToken(addr string) (*schemas.Token, error) {
	if repo.tokens[addr] == nil {
		token, err := schemas.NewToken(addr, repo.client)
		if err != nil {
			return nil, err
		}
		repo.addTokenObj(token)
	}
	return repo.tokens[addr], nil
}

func (repo *TokensRepo) getTokenWithError(addr string) (*schemas.Token, error) {
	token := repo.tokens[addr]
	if token == nil {
		return repo.addToken(addr)
	}
	return token, nil
}

// external funcs
func (repo *TokensRepo) GetToken(addr string) *schemas.Token {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	token, err := repo.getTokenWithError(addr)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	return token
}

func (repo *TokensRepo) GetTokens() []string {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	tokens := []string{}
	for addr, _ := range repo.tokens {
		tokens = append(tokens, addr)
	}
	if repo.GearTokenAddr != "" {
		tokens = append(tokens, repo.GearTokenAddr)
	}
	return tokens
}

func (repo *TokensRepo) AddDieselToken(dieselToken, underlyingToken, pool string) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.dieselTokens[dieselToken] = &schemas.UTokenAndPool{
		UToken: underlyingToken,
		Pool:   pool,
	}
	repo.addToken(dieselToken)
}

// testing purpose
func (repo *TokensRepo) AddTokenObj(obj *schemas.Token) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.addTokenObj(obj)
}

// get specific tokens
func (repo *TokensRepo) setWETHAddr(addr string) {
	repo.WETHAddr = addr
}

func (repo *TokensRepo) GetWETHAddr() string {
	return repo.WETHAddr
}
func (repo *TokensRepo) GetUSDCAddr() string {
	return repo.USDCAddr
}
func (repo *TokensRepo) GetGearTokenAddr() string {
	return repo.GearTokenAddr
}
