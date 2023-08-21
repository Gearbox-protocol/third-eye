package ds

import (
	"math/big"
	"strings"

	"github.com/Gearbox-protocol/sdk-go/artifacts/poolQuotaKeeperv3"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas/schemas_v3"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

type AccountQuotaMgr struct {
	accounts map[string]map[string]*schemas_v3.AccountQuotaInfo
	entries  map[string][]*schemas_v3.AccountQuotaInfo
}

func NewAccountQuotaMgr() *AccountQuotaMgr {
	return &AccountQuotaMgr{
		accounts: map[string]map[string]*schemas_v3.AccountQuotaInfo{},
		entries:  map[string][]*schemas_v3.AccountQuotaInfo{},
	}
}

func (mdl AccountQuotaMgr) InitQuotas(data []*schemas_v3.AccountQuotaInfo) {
	for _, entry := range data {
		account := strings.Split(entry.SessionId, "_")[0]
		mdl.accounts[account][entry.Token] = entry
	}
}

func (mdl AccountQuotaMgr) GetUpdateQuotaEventForAccount(account string) *schemas_v3.AccountQuotaInfo {
	l := len(mdl.entries)
	if l == 0 {
		return nil
	} else {
		front := mdl.entries[account][0]
		mdl.entries[account] = mdl.entries[account][1:]
		return front
	}
}

func (mdl AccountQuotaMgr) AddAccountQuota(blockNum int64,
	keeper PoolKeeperI,
	updateQuota *poolQuotaKeeperv3.PoolQuotaKeeperv3UpdateQuota) {
	account := updateQuota.CreditAccount.Hex()
	token := updateQuota.Token.Hex()
	if mdl.accounts[account] == nil {
		mdl.accounts[account] = map[string]*schemas_v3.AccountQuotaInfo{}
	}
	if mdl.accounts[account][token] == nil {
		mdl.accounts[account][token] = schemas_v3.NotNilAccountQuotaInfo()
	}
	prevValues := mdl.accounts[account][token]
	currentTs := keeper.GetRepo().SetAndGetBlock(blockNum).Timestamp
	// calculate
	currentCumIndex := GetQuptaIndexCurrent(currentTs, keeper.GetQuotas(token))
	//
	newDetails := &schemas_v3.AccountQuotaInfo{
		Timestamp:       currentTs,
		BlockNum:        blockNum,
		Token:           token,
		PoolQuotaKeeper: keeper.GetAddress(),
		//
		QuotaIndex: currentCumIndex,
		Quota:      (*core.BigInt)(new(big.Int).Add(prevValues.Quota.Convert(), updateQuota.QuotaChange)),
		Fees: (*core.BigInt)(
			new(big.Int).Add(
				prevValues.Fees.Convert(),
				GetQuotaFee(updateQuota.QuotaChange, keeper.GetQuotas(token).IncreaseFee),
			),
		),
		Interest: (*core.BigInt)(
			new(big.Int).Add(
				prevValues.Interest.Convert(),
				GetQuotaInterest(keeper.GetQuotas(token).CumQuotaIndex, currentCumIndex, prevValues.Quota),
			),
		),
	}

	// updates
	mdl.entries[account] = append(mdl.entries[account], newDetails.Copy())
	mdl.accounts[account][token] = newDetails
	// delete if qutoa less than <=1
	if newDetails.IsDisabled() {
		delete(mdl.accounts[account], token)
	}
}

// utils

func GetQuptaIndexCurrent(currentTs uint64, prevDetails *schemas_v3.QuotaDetails) *core.BigInt {
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
