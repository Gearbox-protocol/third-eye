package repository

import (
	"fmt"
	"math"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/priceOracle"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// For token with symbol/decimals
func (repo *Repository) AddToken(addr string) *schemas.Token {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	token, err := repo.addToken(addr)
	if err != nil {
		log.Fatal("Adding token failed for", token)
	}
	return token
}

func (repo *Repository) addToken(addr string) (*schemas.Token, error) {
	if repo.tokens[addr] == nil {
		token, err := schemas.NewToken(addr, repo.client)
		if err != nil {
			return nil, err
		}
		repo.AddTokenObj(token)
	}
	return repo.tokens[addr], nil
}

func (repo *Repository) GetToken(addr string) *schemas.Token {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	token, err := repo.getTokenWithError(addr)
	log.CheckFatal(err)
	return token
}

func (repo *Repository) getTokenWithError(addr string) (*schemas.Token, error) {
	token := repo.tokens[addr]
	if token == nil {
		return repo.addToken(addr)
	}
	return token, nil
}

func (repo *Repository) loadToken() {
	defer utils.Elapsed("loadToken")()
	data := []*schemas.Token{}
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
func (repo *Repository) AddTokenObj(t *schemas.Token) {
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
	defer utils.Elapsed("loadAllowedTokensState")()
	data := []*schemas.AllowedToken{}
	// v1 query
	// err := repo.db.Raw("SELECT * FROM allowed_tokens where disable_block = 0 order by block_num").Find(&data).Error
	// v2 query
	err := repo.db.Raw("SELECT distinct on (credit_manager, token) * FROM allowed_tokens order by credit_manager, token, block_num DESC").Find(&data).Error
	log.CheckFatal(err)
	for _, entry := range data {
		repo.addAllowedTokenState(entry, false)
	}
}

func (repo *Repository) addAllowedTokenState(entry *schemas.AllowedToken, usingV2 bool) {
	tokensForCM := repo.allowedTokens[entry.CreditManager]
	if tokensForCM == nil {
		repo.allowedTokens[entry.CreditManager] = make(map[string]*schemas.AllowedToken)
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

// return the active first oracle under blockNum
// if all disabled return the last one
func (repo *Repository) GetActivePriceOracleByBlockNum(blockNum int64) (string, error) {
	var disabledLastOracle, activeFirstOracle string
	var disabledOracleBlock, activeOracleBlock int64
	activeOracleBlock = math.MaxInt64
	oracles := repo.kit.GetAdapterAddressByName(ds.PriceOracle)
	for _, addr := range oracles {
		oracleAdapter := repo.GetAdapter(addr)
		if oracleAdapter.GetDiscoveredAt() <= blockNum {
			if oracleAdapter.IsDisabled() {
				if disabledOracleBlock < oracleAdapter.GetDiscoveredAt() {
					disabledOracleBlock = oracleAdapter.GetDiscoveredAt()
					disabledLastOracle = addr
				}
			} else {
				if activeOracleBlock > oracleAdapter.GetDiscoveredAt() {
					activeOracleBlock = oracleAdapter.GetDiscoveredAt()
					activeFirstOracle = addr
				}
			}
		}
	}
	if activeFirstOracle != "" {
		return activeFirstOracle, nil
	} else if disabledLastOracle != "" {
		return disabledLastOracle, nil
	} else {
		return "", fmt.Errorf("Not Found")
	}
}
func (repo *Repository) GetPriceOracleByVersion(version int16) (string, error) {
	addrProviderAddr := repo.kit.GetAdapterAddressByName(ds.AddressProvider)
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
