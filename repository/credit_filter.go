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
	repo.SetAndGetBlock(p.BlockNumber).AddAllowedProtocol(p)
	args := core.Json{"adapter": p.Adapter, "protocol": p.Protocol, "creditManager": p.CreditManager}
	repo.AddDAOOperation(&schemas.DAOOperation{
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
	repo.AddDAOOperation(&schemas.DAOOperation{
		BlockNumber: blockNum,
		LogID:       logID,
		TxHash:      txHash,
		Contract:    creditFilter,
		Type:        schemas.ContractForbidden,
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
