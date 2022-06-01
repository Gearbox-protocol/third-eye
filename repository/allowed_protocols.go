package repository

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
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
