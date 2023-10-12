package ds

import (
	"math/big"
	"sync"

	"github.com/Gearbox-protocol/sdk-go/artifacts/poolQuotaKeeperv3"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas/schemas_v3"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

type AccountQuotaMgr struct {
	mu sync.Mutex
	// so that credit manager can create latest quota update event
	events map[string][]*poolQuotaKeeperv3.PoolQuotaKeeperv3UpdateQuota
}

func NewAccountQuotaMgr() *AccountQuotaMgr {
	return &AccountQuotaMgr{
		events: map[string][]*poolQuotaKeeperv3.PoolQuotaKeeperv3UpdateQuota{},
	}
}

func (mdl *AccountQuotaMgr) GetUpdateQuotaEventForAccount(account string) *poolQuotaKeeperv3.PoolQuotaKeeperv3UpdateQuota {
	mdl.mu.Lock()
	defer mdl.mu.Unlock()
	l := len(mdl.events[account])
	if l == 0 {
		return nil
	} else {
		front := mdl.events[account][0]
		mdl.events[account] = mdl.events[account][1:]
		return front
	}
}

func (mdl *AccountQuotaMgr) AddAccountQuota(blockNum int64,
	updateQuota *poolQuotaKeeperv3.PoolQuotaKeeperv3UpdateQuota) {
	mdl.mu.Lock()
	defer mdl.mu.Unlock()
	//
	account := updateQuota.CreditAccount.Hex()
	mdl.events[account] = append(mdl.events[account], updateQuota)
}

// utils

func GetQuotaIndexCurrent(currentTs uint64, prevDetails *schemas_v3.QuotaDetails) *core.BigInt {
	prevCumIndex := prevDetails.CumQuotaIndex.Convert()

	interestInPeriod := func() *big.Int {
		rayByPrecent := utils.GetExpInt(27 - 4)
		interestInYear := new(big.Int).Mul(rayByPrecent, big.NewInt(int64(currentTs-prevDetails.Timestamp)*int64(prevDetails.Rate)))
		return new(big.Int).Quo(interestInYear, big.NewInt(365*24*60*60))
	}()
	return (*core.BigInt)(new(big.Int).Add(interestInPeriod, prevCumIndex))
}

func GetQuotaInterest(oldCum, newCum, quoted *core.BigInt) *big.Int {
	cumDiff := new(big.Int).Sub(newCum.Convert(), oldCum.Convert())
	interstRay := new(big.Int).Mul(quoted.Convert(), cumDiff)
	return new(big.Int).Quo(interstRay, utils.GetExpInt(27))
}

func GetQuotaFee(change *big.Int, increaseFee uint16) *big.Int {

	val := new(big.Int).Mul(change, big.NewInt(int64(increaseFee)))
	return new(big.Int).Quo(val, utils.GetExpInt(4))
}

type PoolKeeperI interface {
	GetRepo() RepositoryI
	GetQuotas(token string) *schemas_v3.QuotaDetails
	GetAddress() string
}
