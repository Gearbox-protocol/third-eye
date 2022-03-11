package repository

import (
	"github.com/Gearbox-protocol/third-eye/artifacts/creditFilter"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// for credit filter
func (repo *Repository) AddAllowedProtocol(logID uint, txHash, creditFilter string, p *core.Protocol) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.blocks[p.BlockNumber].AddAllowedProtocol(p)
	args := core.Json{"adapter": p.Adapter, "protocol": p.Protocol, "creditManager": p.CreditManager}
	repo.addDAOOperation(&core.DAOOperation{
		BlockNumber: p.BlockNumber,
		LogID:       logID,
		TxHash:      txHash,
		Contract:    creditFilter,
		Type:        core.ContractAllowed,
		Args:        &args,
	})
}

func (repo *Repository) DisableProtocol(blockNum int64, logID uint, txHash, cm, creditFilter, protocol string) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	args := core.Json{"protocol": protocol, "creditManager": cm}
	repo.addDAOOperation(&core.DAOOperation{
		BlockNumber: blockNum,
		LogID:       logID,
		TxHash:      txHash,
		Contract:    creditFilter,
		Type:        core.ContractForbidden,
		Args:        &args,
	})
}

func (repo *Repository) AddAllowedToken(logID uint, txHash, creditFilter string, atoken *core.AllowedToken) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.setAndGetBlock(atoken.BlockNumber).AddAllowedToken(atoken)
	prevLiqThreshold := repo.GetPreviousLiqThreshold(atoken.CreditManager, atoken.Token)
	args := core.Json{
		"liquidityThreshold":       atoken.LiquidityThreshold,
		"token":                    atoken.Token,
		"creditManager":            atoken.CreditManager,
		"prevLiquidationThreshold": prevLiqThreshold,
	}
	repo.addAllowedTokenState(atoken, false)
	repo.addDAOOperation(&core.DAOOperation{
		BlockNumber: atoken.BlockNumber,
		LogID:       logID,
		TxHash:      txHash,
		Contract:    creditFilter,
		Type:        core.TokenAllowed,
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
// (c1)if the previous lt is not present only store dao operation.
// (c2)if previous lt is present, set the old lt to new lt. store dao operation, update allowed token state and add allowed token to table.
// if previous lt has disabledBlock,  set reenabled
// if previous lt hasn't disabledBlock, disable previous entry
// - liquiditythreshold emitted.
// (c3)store dao operation, update allowed token state and add allowed token to table.
func (repo *Repository) AddAllowedTokenV2(logID uint, txHash, creditFilter string, atoken *core.AllowedToken) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	cm := atoken.CreditManager
	token := atoken.Token
	prevLiqThreshold := repo.GetPreviousLiqThreshold(cm, token)
	if prevLiqThreshold.Convert().Int64() == 0 && atoken.LiquidityThreshold == nil { // c1
		repo.addDAOOperation(&core.DAOOperation{
			BlockNumber: atoken.BlockNumber,
			LogID:       logID,
			TxHash:      txHash,
			Contract:    creditFilter,
			Type:        core.TokenAllowedV2,
			Args:        &core.Json{
				"token":                    atoken.Token,
				"creditManager":            atoken.CreditManager,
			},
		})
		return
	}
	if atoken.LiquidityThreshold == nil { // c2
		atoken.LiquidityThreshold = prevLiqThreshold
	}
	reEnable := repo.isAllowedTokenDisabled(cm, token)
	repo.setAndGetBlock(atoken.BlockNumber).AddAllowedToken(atoken)
	args := core.Json{
		"liquidityThreshold":       atoken.LiquidityThreshold,
		"token":                    atoken.Token,
		"creditManager":            atoken.CreditManager,
		"prevLiquidationThreshold": prevLiqThreshold,
	}
	// previous allowed token disabled
	if reEnable {
		args["type"] = "reEnabled"
	} else if repo.allowedTokens[cm] != nil && repo.allowedTokens[cm][token] != nil {
		// and previous entries is present has disabledBlock set to 0
		// previous allowed token enabled
		prevToken := repo.allowedTokens[atoken.CreditManager][atoken.Token]
		prevToken.DisableBlock = atoken.BlockNumber
		repo.disabledTokens = append(repo.disabledTokens, prevToken)
	}
	repo.addAllowedTokenState(atoken, true)
	var daoEventType uint 
	if atoken.LiquidityThreshold == nil {
		daoEventType = core.TokenAllowedV2
	} else {
		daoEventType = core.LTUpdated
	}
	repo.addDAOOperation(&core.DAOOperation{
		BlockNumber: atoken.BlockNumber,
		LogID:       logID,
		TxHash:      txHash,
		Contract:    creditFilter,
		Type:        daoEventType,
		Args:        &args,
	})
}

func (repo *Repository) DisableAllowedToken(blockNum int64, logID uint, txHash, creditManager, creditFilter, token string) {
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
	repo.addDAOOperation(&core.DAOOperation{
		BlockNumber: atoken.DisableBlock,
		LogID:       logID,
		TxHash:      txHash,
		Contract:    creditFilter,
		Type:        core.TokenForbidden,
		Args:        &args,
	})
}

func (repo *Repository) AddCreditManagerToFilter(cmAddr, cfAddr string) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	cf, err := creditFilter.NewCreditFilter(common.HexToAddress(cfAddr), repo.client)
	log.CheckFatal(err)
	repo.creditManagerToFilter[cmAddr] = cf
}

func (repo *Repository) getCreditManagerToFilter(cmAddr string) *creditFilter.CreditFilter {
	cf := repo.creditManagerToFilter[cmAddr]
	if cf == nil {
		log.Fatalf("Credit filter not found for manager: %s", cmAddr)
	}
	return cf
}

func (repo *Repository) GetMask(blockNum int64, cmAddr, accountAddr string) *big.Int {
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
	}
	mask, err := repo.getCreditManagerToFilter(cmAddr).EnabledTokens(opts, common.HexToAddress(accountAddr))
	log.CheckFatal(err)
	return mask
}

func (repo *Repository) AddFastCheckParams(logID uint, txHash, cm, creditFilter string, fcParams *core.FastCheckParams) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.setAndGetBlock(fcParams.BlockNum).AddFastCheckParams(fcParams)
	// set the dao action
	oldFCParams := repo.cmFastCheckParams[fcParams.CreditManager]
	if oldFCParams == nil {
		oldFCParams = core.NewFastCheckParams()
	}
	args := oldFCParams.Diff(fcParams)
	(*args)["creditManager"] = cm
	repo.addDAOOperation(&core.DAOOperation{
		BlockNumber: fcParams.BlockNum,
		LogID:       logID,
		TxHash:      txHash,
		Contract:    creditFilter,
		Args:        args,
		Type:        core.NewFastCheckParameters,
	})
	//
	repo.cmFastCheckParams[fcParams.CreditManager] = fcParams
}

func (repo *Repository) AddConfiguratorUpdated(blockNum int64, logID uint, txHash, cm, oldconfigurator, configurator string) {
	repo.addDAOOperation(&core.DAOOperation{
		BlockNumber: blockNum,
		LogID:       logID,
		TxHash:      txHash,
		Contract:    oldconfigurator,
		Args:        &core.Json{"oldConfigurator": oldconfigurator, "configurator": configurator},
		Type:        core.NewFastCheckParameters,
	})
}

func (repo *Repository) AddFacadeUpdated(blockNum int64, logID uint, txHash, cm, oldconfigurator, facade string) {
	repo.addDAOOperation(&core.DAOOperation{
		BlockNumber: blockNum,
		LogID:       logID,
		TxHash:      txHash,
		Contract:    oldconfigurator,
		Args:        &core.Json{"facade": facade},
		Type:        core.NewFastCheckParameters,
	})
}