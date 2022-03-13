package repository

import (
	"fmt"
	"math/big"

	"github.com/Gearbox-protocol/third-eye/artifacts/priceOracle"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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
		repo.AddTokenObj(token)
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
		repo.AddTokenObj(token)
	}
}

// not to be called directly
// only exposed for testing framework
func (repo *Repository) AddTokenObj(t *core.Token) {
	// set usdc addr in repo
	if t.Symbol == "USDC" {
		repo.USDCAddr = t.Address
	}
	if t.Symbol == "WETH" {
		repo.SetWETHAddr(t.Address)
	}
	if repo.tokens[t.Address] == nil {
		repo.tokens[t.Address] = t
	}
}

func (repo *Repository) loadAllowedTokensState() {
	data := []*core.AllowedToken{}
	// v1 query
	// err := repo.db.Raw("SELECT * FROM allowed_tokens where disable_block = 0 order by block_num").Find(&data).Error
	// v2 query
	err := repo.db.Raw("SELECT distinct on (credit_manager, token) * FROM allowed_tokens order by credit_manager, token, block_num DESC").Find(&data).Error
	log.CheckFatal(err)
	for _, entry := range data {
		repo.addAllowedTokenState(entry, false)
	}
}

func (repo *Repository) addAllowedTokenState(entry *core.AllowedToken, usingV2 bool) {
	tokensForCM := repo.allowedTokens[entry.CreditManager]
	if tokensForCM == nil {
		repo.allowedTokens[entry.CreditManager] = make(map[string]*core.AllowedToken)
		tokensForCM = repo.allowedTokens[entry.CreditManager]
	}
	if tokensForCM[entry.Token] != nil && !usingV2 {
		log.Warnf("Token already enabled: new %#v, previous entry: %#v", entry, tokensForCM[entry.Token])
	}
	tokensForCM[entry.Token] = entry
}

func (repo *Repository) GetPreviousLiqThreshold(cm, token string) *core.BigInt {
	if repo.allowedTokens[cm] == nil || repo.allowedTokens[cm][token] == nil {
		return (*core.BigInt)(new(big.Int))
	}
	return repo.allowedTokens[cm][token].LiquidityThreshold
}
func (repo *Repository) isAllowedTokenDisabled(cm, token string) bool {
	if repo.allowedTokens[cm] == nil || repo.allowedTokens[cm][token] == nil {
		return false
	}
	return repo.allowedTokens[cm][token].DisableBlock != 0
}

func (repo *Repository) GetActivePriceOracleByBlockNum(blockNum int64) (string, error) {
	oracles := repo.kit.GetAdapterAddressByName(core.PriceOracle)
	for _, addr := range oracles {
		oracleAdapter := repo.GetAdapter(addr)
		if !oracleAdapter.IsDisabled() && oracleAdapter.GetDiscoveredAt() <= blockNum {
			return addr, nil
		}
	}
	return "", fmt.Errorf("Not Found")
}
func (repo *Repository) GetPriceOracleByVersion(version int16) (string, error) {
	addrProviderAddr := repo.kit.GetAdapterAddressByName(core.AddressProvider)
	addrProvider := repo.kit.GetAdapter(addrProviderAddr[0])
	details := addrProvider.GetDetails()
	if details != nil {
		priceOracles := details["priceOracles"]
		if priceOracles != nil {
			return utils.ConvertToListOfString(priceOracles)[version-1], nil
		}
	}
	return "", fmt.Errorf("Not Found")
}

// This function is used for getting the collateral value in usd and underlying
func (repo *Repository) GetValueInCurrency(blockNum int64, version int16, token, currency string, amount *big.Int) *big.Int {
	oracle, err := repo.GetPriceOracleByVersion(version)
	log.CheckFatal(err)
	poContract, err := priceOracle.NewPriceOracle(common.HexToAddress(oracle), repo.client)
	log.CheckFatal(err)
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
	}
	currencyAddr := common.HexToAddress(repo.USDCAddr)
	if currency != "USDC" {
		currencyAddr = common.HexToAddress(currency)
	}
	usdcAmount, err := poContract.Convert(opts, amount, common.HexToAddress(token), currencyAddr)
	log.CheckFatal(err)
	// convert to 8 decimals
	return usdcAmount
}

func (repo *Repository) GetTokens() []string {
	tokens := []string{}
	for addr, _ := range repo.tokens {
		tokens = append(tokens, addr)
	}
	if repo.GearTokenAddr != "" {
		tokens = append(tokens, repo.GearTokenAddr)
	}
	return tokens
}
