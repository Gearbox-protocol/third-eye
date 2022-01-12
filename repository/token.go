package repository

import (
	"fmt"
	"github.com/Gearbox-protocol/third-eye/artifacts/priceOracle"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// For token with symbol/decimals
func (repo *Repository) AddToken(addr string) *core.Token {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	token, err := repo.addToken(addr)
	if err != nil {
		log.Fatal("Adding token failed for", token)
	}
	return token
}

func (repo *Repository) addToken(addr string) (*core.Token, error) {
	if repo.tokens[addr] == nil {
		token, err := core.NewToken(addr, repo.client)
		if err != nil {
			return nil, err
		}
		repo.addTokenObj(token)
	}
	return repo.tokens[addr], nil
}

func (repo *Repository) GetToken(addr string) *core.Token {
	repo.mu.Lock()
	defer repo.mu.Unlock()
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
		repo.addTokenObj(token)
	}
}

func (repo *Repository) addTokenObj(t *core.Token) {
	// set usdc addr in repo
	if t.Symbol == "USDC" {
		repo.USDCAddr = t.Address
	}
	if repo.tokens[t.Address] == nil {
		repo.tokens[t.Address] = t
	}
}

func (repo *Repository) loadAllowedTokensState() {
	data := []*core.AllowedToken{}
	err := repo.db.Raw("SELECT * FROM allowed_tokens where disable_block = 0 order by block_num").Find(&data).Error
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
	}
	tokensForCM[entry.Token] = entry
}

func (repo *Repository) GetActivePriceOracle() (string, error) {
	oracles := repo.kit.GetAdapterAddressByName(core.PriceOracle)
	for _, addr := range oracles {
		oracleAdapter := repo.kit.GetAdapter(addr)
		if !oracleAdapter.IsDisabled() {
			return addr, nil
		}
	}
	return "", fmt.Errorf("Not Found")
}

func (repo *Repository) GetPriceInUSD(blockNum int64, token string, amount *big.Int) *big.Int {
	oracle, err := repo.GetActivePriceOracle()
	log.CheckFatal(err)
	poContract, err := priceOracle.NewPriceOracle(common.HexToAddress(oracle), repo.client)
	log.CheckFatal(err)
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
	}
	usdcAmount, err := poContract.Convert(opts, amount, common.HexToAddress(token), common.HexToAddress(repo.USDCAddr))
	log.CheckFatal(err)
	return usdcAmount
}
