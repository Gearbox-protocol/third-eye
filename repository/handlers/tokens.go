package handlers

import (
	"sync"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TokensRepo struct {
	// export vars
	wethAddr      string
	usdcAddr      string
	gearTokenAddr string
	// blocks/token
	tokens map[string]*schemas.Token
	// diesel tokens to pool and underlying
	dieselTokens map[string]*schemas.UTokenAndPool
	//
	mu             *sync.Mutex
	client         core.ClientI
	symToAddrToken map[string]common.Address
}

func NewTokensRepo(client core.ClientI) *TokensRepo {
	return &TokensRepo{
		mu:     &sync.Mutex{},
		client: client,
		//
		tokens: map[string]*schemas.Token{},
		// for getting the diesel tokens
		dieselTokens:   make(map[string]*schemas.UTokenAndPool),
		symToAddrToken: core.GetSymToAddrByChainId(core.GetChainId(client)).Tokens,
	}
}

func (repo *TokensRepo) GetTokenFromSdk(symbol string) string {
	if addr, ok := repo.symToAddrToken[symbol]; ok {
		return addr.Hex()
	}
	log.Fatalf("Can't get token(%s) from sdk", symbol)
	return ""
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
	defer utils.Elapsed("tokens sql statements")()
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
	if t.Address == "0xEe8Adf657c5EF8e10622b6B47014D2C6f6993E5E" { // in goerli , for yvWETH the symbol is set to WETH.
		t.Symbol = "yvWETH"
	}
	// set usdc addr in repo
	if t.Symbol == "USDC" {
		repo.usdcAddr = t.Address
	}
	if t.Symbol == "WETH" {
		repo.wethAddr = t.Address
	}
	if t.Symbol == "GEAR" {
		repo.gearTokenAddr = t.Address
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
func (repo *TokensRepo) GetDecimalsForList([]common.Address) {
}

func (repo *TokensRepo) GetDecimals(addr common.Address) int8 {
	return repo.GetToken(addr.Hex()).Decimals
}

func (repo *TokensRepo) GetTokens() []string {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	tokens := []string{}
	for addr, _ := range repo.tokens {
		tokens = append(tokens, addr)
	}
	if repo.gearTokenAddr != "" {
		tokens = append(tokens, repo.gearTokenAddr)
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

func (repo *TokensRepo) GetDieselToken(dieselToken string) *schemas.UTokenAndPool {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	return repo.dieselTokens[dieselToken]
}

// testing purpose
func (repo *TokensRepo) AddTokenObj(obj *schemas.Token) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.addTokenObj(obj)
}

// get specific tokens
func (repo *TokensRepo) setwethAddr(addr string) {
	repo.wethAddr = addr
}

func (repo *TokensRepo) GetWETHAddr() string {
	return repo.wethAddr
}
func (repo *TokensRepo) GetUSDCAddr() string {
	return repo.usdcAddr
}
func (repo *TokensRepo) GetGearTokenAddr() string {
	return repo.gearTokenAddr
}

func (repo *TokensRepo) IsDieselToken(token string) bool {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	return repo.dieselTokens[token] != nil
}

func (repo *TokensRepo) GetDieselTokens() map[string]*schemas.UTokenAndPool {
	return repo.dieselTokens
}
