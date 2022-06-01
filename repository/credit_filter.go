package repository

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/creditFilter"
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditManagerv2"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// for credit filter
func (repo *Repository) AddAllowedProtocol(logID uint, txHash, creditFilter string, p *schemas.Protocol) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.blocks[p.BlockNumber].AddAllowedProtocol(p)
	args := core.Json{"adapter": p.Adapter, "protocol": p.Protocol, "creditManager": p.CreditManager}
	repo.addDAOOperation(&schemas.DAOOperation{
		BlockNumber: p.BlockNumber,
		LogID:       logID,
		TxHash:      txHash,
		Contract:    creditFilter,
		Type:        schemas.ContractAllowed,
		Args:        &args,
	})
}

func (repo *Repository) DisableProtocol(blockNum int64, logID uint, txHash, cm, creditFilter, protocol string) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	args := core.Json{"protocol": protocol, "creditManager": cm}
	repo.addDAOOperation(&schemas.DAOOperation{
		BlockNumber: blockNum,
		LogID:       logID,
		TxHash:      txHash,
		Contract:    creditFilter,
		Type:        schemas.ContractForbidden,
		Args:        &args,
	})
}

// for allowed token
func (repo *Repository) addAllowedToken(atoken *schemas.AllowedToken) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.addToken(atoken.Token)
	repo.setAndGetBlock(atoken.BlockNumber).AddAllowedToken(atoken)
}

func (repo *Repository) DisableAllowedToken(blockNum int64, logID uint, txHash, creditManager, creditFilter, token string) {
	daoOperation := repo.AllowedTokenRepo.DisableAllowedToken(blockNum, logID, txHash, creditManager, creditFilter, token)
	repo.addDAOOperation(daoOperation)
}

func (repo *Repository) AddAllowedToken(logID uint, txHash, creditFilter string, atoken *schemas.AllowedToken) {
	repo.addAllowedToken(atoken)
	daoOperation := repo.AllowedTokenRepo.AddAllowedToken(logID, txHash, creditFilter, atoken)
	repo.addDAOOperation(daoOperation)
}

func (repo *Repository) AddAllowedTokenV2(logID uint, txHash, creditFilter string, atoken *schemas.AllowedToken) {
	atoken, daoOperation := repo.AllowedTokenRepo.AddAllowedTokenV2(logID, txHash, creditFilter, atoken)
	if atoken != nil {
		repo.addAllowedToken(atoken)
	}
	repo.addDAOOperation(daoOperation)
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

func (repo *Repository) GetMask(blockNum int64, cmAddr, accountAddr string, version int16) *big.Int {
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
	}
	switch version {
	case 1:
		mask, err := repo.getCreditManagerToFilter(cmAddr).EnabledTokens(opts, common.HexToAddress(accountAddr))
		log.CheckFatal(err)
		return mask
	case 2:
		cm, err := creditManagerv2.NewCreditManagerv2(common.HexToAddress(cmAddr), repo.client)
		log.CheckFatal(err)
		mask, err := cm.EnabledTokensMap(opts, common.HexToAddress(accountAddr))
		log.CheckFatal(err)
		return mask
	}
	return nil
}

func (repo *Repository) AddFastCheckParams(logID uint, txHash, cm, creditFilter string, fcParams *schemas.FastCheckParams) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.setAndGetBlock(fcParams.BlockNum).AddFastCheckParams(fcParams)
	// set the dao action
	oldFCParams := repo.cmFastCheckParams[fcParams.CreditManager]
	if oldFCParams == nil {
		oldFCParams = schemas.NewFastCheckParams()
	}
	args := oldFCParams.Diff(fcParams)
	(*args)["creditManager"] = cm
	repo.addDAOOperation(&schemas.DAOOperation{
		BlockNumber: fcParams.BlockNum,
		LogID:       logID,
		TxHash:      txHash,
		Contract:    creditFilter,
		Args:        args,
		Type:        schemas.NewFastCheckParameters,
	})
	//
	repo.cmFastCheckParams[fcParams.CreditManager] = fcParams
}
