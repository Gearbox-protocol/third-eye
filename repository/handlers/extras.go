package handlers

import (
	"math/big"
	"sync"

	"github.com/Gearbox-protocol/sdk-go/artifacts/creditFilter"
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditManagerv2"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds/dc_wrapper"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type ExtrasRepo struct {
	dcWrapper *dc_wrapper.DataCompressorWrapper
	// AggregatedFeed        *aggregated_block_feed.AggregatedBlockFeed
	creditManagerToFilter map[string]*creditFilter.CreditFilter
	client                core.ClientI
	mu                    *sync.Mutex
}

func NewExtraRepo(client core.ClientI) *ExtrasRepo {
	return &ExtrasRepo{
		dcWrapper:             dc_wrapper.NewDataCompressorWrapper(client),
		creditManagerToFilter: make(map[string]*creditFilter.CreditFilter),
		client:                client,
		mu:                    &sync.Mutex{},
	}
}

func (repo *ExtrasRepo) GetDCWrapper() *dc_wrapper.DataCompressorWrapper {
	return repo.dcWrapper
}

func (repo *ExtrasRepo) AddCreditManagerToFilter(cmAddr, cfAddr string) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	cf, err := creditFilter.NewCreditFilter(common.HexToAddress(cfAddr), repo.client)
	log.CheckFatal(err)
	repo.creditManagerToFilter[cmAddr] = cf
}

func (repo *ExtrasRepo) getCreditManagerToFilter(cmAddr string) *creditFilter.CreditFilter {
	cf := repo.creditManagerToFilter[cmAddr]
	if cf == nil {
		log.Fatalf("Credit filter not found for manager: %s", cmAddr)
	}
	return cf
}

func (repo *ExtrasRepo) GetMask(blockNum int64, cmAddr, accountAddr string, version int16) *big.Int {
	repo.mu.Lock()
	defer repo.mu.Unlock()
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
