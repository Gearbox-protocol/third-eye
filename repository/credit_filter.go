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
	args := core.Json{"adapter": p.Adapter, "protocol": p.Protocol}
	repo.AddDAOOperation(&core.DAOOperation{
		BlockNumber: p.BlockNumber,
		LogID:       logID,
		TxHash:      txHash,
		Contract:    creditFilter,
		Type:        core.ContractAllowed,
		Args:        &args,
	})
}

func (repo *Repository) DisableProtocol(blockNum int64, logID uint, txHash, cm, creditFilter, protocol string) {
	args := core.Json{"protocol": protocol}
	repo.AddDAOOperation(&core.DAOOperation{
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
	repo.addAllowedTokenState(atoken)
	repo.blocks[atoken.BlockNumber].AddAllowedToken(atoken)
	args := core.Json{"liqThreshold": atoken.LiquidityThreshold}
	repo.AddDAOOperation(&core.DAOOperation{
		BlockNumber: atoken.BlockNumber,
		LogID:       logID,
		TxHash:      txHash,
		Contract:    creditFilter,
		Type:        core.TokenForbidden,
		Args:        &args,
	})
}

func (repo *Repository) DisableAllowedToken(blockNum int64, logID uint, txHash, creditFilter string, creditManager, token string) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	atoken := repo.allowedTokens[creditManager][token]
	atoken.DisableBlock = blockNum
	repo.disabledTokens = append(repo.disabledTokens, atoken)
	delete(repo.allowedTokens[creditManager], token)
	repo.AddDAOOperation(&core.DAOOperation{
		BlockNumber: atoken.DisableBlock,
		LogID:       logID,
		TxHash:      txHash,
		Contract:    creditFilter,
		Type:        core.TokenAllowed,
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
	repo.mu.Lock()
	defer repo.mu.Unlock()
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
	}
	mask, err := repo.getCreditManagerToFilter(cmAddr).EnabledTokens(opts, common.HexToAddress(accountAddr))
	log.CheckFatal(err)
	return mask
}

func (repo *Repository) AddDAOOperation(operation *core.DAOOperation) {
	repo.GetBlock(operation.BlockNumber).AddDAOOperation(operation)
}

func (repo *Repository) AddFastCheckParams(logID uint, txHash, creditFilter string, fcParams *core.FastCheckParams) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.GetBlock(fcParams.BlockNum).AddFastCheckParams(fcParams)
	// set the dao action
	oldFCParams := repo.cmFastCheckParams[fcParams.CreditManager]
	if oldFCParams == nil {
		oldFCParams = core.NewFastCheckParams()
	}
	args := oldFCParams.Diff(fcParams)
	repo.AddDAOOperation(&core.DAOOperation{
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
