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
func (repo *Repository) AddAllowedProtocol(p *core.Protocol) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.blocks[p.BlockNumber].AddAllowedProtocol(p)
}

func (repo *Repository) AddAllowedToken(atoken *core.AllowedToken) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.blocks[atoken.BlockNumber].AddAllowedToken(atoken)
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
