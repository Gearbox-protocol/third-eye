package handlers

import (
	"math/big"
	"sync"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AllowedTokenRepo struct {
	allowedTokens  map[string]map[string]*schemas.AllowedToken
	disabledTokens []*schemas.AllowedToken
	blocks         *BlocksRepo
	mu             *sync.Mutex
	tokens         *TokensRepo
}

func NewAllowedTokenRepo(blocks *BlocksRepo, tokens *TokensRepo) *AllowedTokenRepo {
	return &AllowedTokenRepo{
		allowedTokens: make(map[string]map[string]*schemas.AllowedToken),
		mu:            &sync.Mutex{},
		blocks:        blocks,
		tokens:        tokens,
	}
}

// internal functions
func (repo *AllowedTokenRepo) addAllowedTokenState(entry *schemas.AllowedToken, usingV2 bool) {
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

func (repo *AllowedTokenRepo) getPreviousLiqThreshold(cm, token string) *core.BigInt {
	if repo.allowedTokens[cm] == nil || repo.allowedTokens[cm][token] == nil {
		return (*core.BigInt)(new(big.Int))
	}
	return repo.allowedTokens[cm][token].LiquidityThreshold
}

// if entry is not present or if token is not disabled returns false
func (repo *AllowedTokenRepo) isAllowedTokenDisabled(cm, token string) bool {
	if repo.allowedTokens[cm] == nil || repo.allowedTokens[cm][token] == nil {
		return false
	}
	return repo.allowedTokens[cm][token].DisableBlock != 0
}

// for allowed token
func (repo *AllowedTokenRepo) addAllowedToken(atoken *schemas.AllowedToken) {
	repo.tokens.GetToken(atoken.Token)
	repo.blocks.SetAndGetBlock(atoken.BlockNumber).AddAllowedToken(atoken)
}

// external functions
func (repo *AllowedTokenRepo) LoadAllowedTokensState(db *gorm.DB) {
	defer utils.Elapsed("loadAllowedTokensState")()
	data := []*schemas.AllowedToken{}
	// v1 query
	// err := repo.db.Raw("SELECT * FROM allowed_tokens where disable_block = 0 order by block_num").Find(&data).Error
	// v2 query
	err := db.Raw("SELECT distinct on (credit_manager, token) * FROM allowed_tokens order by credit_manager, token, block_num DESC").Find(&data).Error
	log.CheckFatal(err)
	for _, entry := range data {
		repo.addAllowedTokenState(entry, false)
	}
}

func (repo *AllowedTokenRepo) Save(tx *gorm.DB) {
	defer utils.Elapsed("allowed token sql statements")()
	// add disabled tokens after the block num and allowed tokens are synced to db
	if len(repo.disabledTokens) > 0 {
		err := tx.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).CreateInBatches(repo.disabledTokens, 50).Error
		log.CheckFatal(err)
		repo.disabledTokens = []*schemas.AllowedToken{}
	}
}

// external funcs

func (repo *AllowedTokenRepo) AddAllowedToken(logID uint, txHash, creditFilter string, atoken *schemas.AllowedToken) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.addAllowedToken(atoken)
	prevLiqThreshold := repo.getPreviousLiqThreshold(atoken.CreditManager, atoken.Token)
	args := core.Json{
		"liquidityThreshold":       atoken.LiquidityThreshold,
		"token":                    atoken.Token,
		"creditManager":            atoken.CreditManager,
		"prevLiquidationThreshold": prevLiqThreshold,
	}
	repo.addAllowedTokenState(atoken, false)
	repo.blocks.AddDAOOperation(&schemas.DAOOperation{
		BlockNumber: atoken.BlockNumber,
		LogID:       logID,
		TxHash:      txHash,
		Contract:    creditFilter,
		Type:        schemas.TokenAllowed,
		Args:        &args,
	})
}

// allowed token

// v1 logic
// - token, threshold emitted.
// Take the difference from the previous lt present for this token.
// Store dao operation, add allowed token state with new lt. Add allowed token to table.

// v2 logic
// - token emitted.
// (c1) current lt is not set only store dao operation.
// (c2)if previous lt is present, set the old lt to new lt. store dao operation, update allowed token state and add allowed token to table.
// if previous lt has disabledBlock,  set reenabled
// if previous lt hasn't disabledBlock, disable previous entry
// - liquiditythreshold emitted.
// (c3)store dao operation, update allowed token state and add allowed token to table.
func (repo *AllowedTokenRepo) AddAllowedTokenV2(logID uint, txHash, creditFilter string, atoken *schemas.AllowedToken) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	cm := atoken.CreditManager
	token := atoken.Token
	prevLiqThreshold := repo.getPreviousLiqThreshold(cm, token)
	canReEnable := repo.isAllowedTokenDisabled(cm, token)

	if (prevLiqThreshold.Convert().Int64() == 0 || !canReEnable) &&
		// prevLiq is checked for 0 and current lt is also nil, then only add dao event
		// if prevLiq is not 0 but prevLiq event is also not disabled and only AllowToken is emitted then only add dao event, prevLiq will be used automatically
		atoken.LiquidityThreshold == nil { // c1,
		repo.blocks.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: atoken.BlockNumber,
			LogID:       logID,
			TxHash:      txHash,
			Contract:    creditFilter,
			Type:        schemas.TokenAllowedV2,
			Args: &core.Json{
				"token":                    atoken.Token,
				"creditManager":            atoken.CreditManager,
				"prevLiquidationThreshold": "0",
			},
		})
		return
	}
	isAllowTokenEvent := atoken.LiquidityThreshold == nil
	if isAllowTokenEvent { // c2
		atoken.LiquidityThreshold = prevLiqThreshold
	}
	repo.addAllowedToken(atoken)
	//
	args := core.Json{
		"liquidityThreshold":       atoken.LiquidityThreshold,
		"token":                    atoken.Token,
		"creditManager":            atoken.CreditManager,
		"prevLiquidationThreshold": prevLiqThreshold,
	}
	// previous allowed token disabled
	if canReEnable {
		args["type"] = "reEnabled"
	} else if repo.allowedTokens[cm] != nil && repo.allowedTokens[cm][token] != nil {
		// and previous entries is present has disabledBlock set to 0
		// previous allowed token enabled
		prevToken := repo.allowedTokens[atoken.CreditManager][atoken.Token]
		prevToken.DisableBlock = atoken.BlockNumber
		repo.blocks.SetBlock(atoken.BlockNumber)
		repo.disabledTokens = append(repo.disabledTokens, prevToken)
	}
	repo.addAllowedTokenState(atoken, true)
	var daoEventType uint
	if isAllowTokenEvent {
		daoEventType = schemas.TokenAllowedV2
	} else {
		daoEventType = schemas.LTUpdated
	}
	repo.blocks.AddDAOOperation(&schemas.DAOOperation{
		BlockNumber: atoken.BlockNumber,
		LogID:       logID,
		TxHash:      txHash,
		Contract:    creditFilter,
		Type:        daoEventType,
		Args:        &args,
	})
}

func (repo *AllowedTokenRepo) DisableAllowedToken(blockNum int64, logID uint, txHash, creditManager, creditFilter, token string) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	atoken := repo.allowedTokens[creditManager][token]
	atoken.DisableBlock = blockNum
	repo.disabledTokens = append(repo.disabledTokens, atoken)
	args := core.Json{
		"token":         token,
		"creditManager": creditManager,
	}
	// for v2 we shouldn't delete the previous state as it will be required for lt if only token is emitted.
	// delete(repo.allowedTokens[creditManager], token)
	repo.blocks.AddDAOOperation(&schemas.DAOOperation{
		BlockNumber: atoken.DisableBlock,
		LogID:       logID,
		TxHash:      txHash,
		Contract:    creditFilter,
		Type:        schemas.TokenForbidden,
		Args:        &args,
	})
}

func (repo *AllowedTokenRepo) GetDisabledTokens() []*schemas.AllowedToken {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	return repo.disabledTokens
}
