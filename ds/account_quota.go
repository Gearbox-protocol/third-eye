package ds

import (
	"math/big"
	"sync"

	"github.com/Gearbox-protocol/sdk-go/artifacts/poolQuotaKeeperv3"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas/schemas_v3"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type AccountQuotaMgr struct {
	mu sync.Mutex
	// so that credit manager can create latest quota update event
	// account+ txHash
	events   map[string][]*updateQuotaEvent
	contract *poolQuotaKeeperv3.PoolQuotaKeeperv3
}

type updateQuotaEvent struct {
	CreditAccount common.Address
	Token         common.Address
	QuotaChange   *big.Int
	TxHash        string
	Index         uint
	BlockNumber   int64
}

func NewAccountQuotaMgr(client core.ClientI) *AccountQuotaMgr {
	return &AccountQuotaMgr{
		contract: func() *poolQuotaKeeperv3.PoolQuotaKeeperv3 {
			contract, err := poolQuotaKeeperv3.NewPoolQuotaKeeperv3(core.NULL_ADDR, client)
			log.CheckFatal(err)
			return contract
		}(),
		events: map[string][]*updateQuotaEvent{},
	}
}

func (mdl *AccountQuotaMgr) GetUpdateQuotaEventForAccount(txHash common.Hash, account string) []*updateQuotaEvent {
	mdl.mu.Lock()
	defer mdl.mu.Unlock()
	key := account + txHash.Hex()
	//
	ans := mdl.events[key]
	delete(mdl.events, key)
	return ans
}

func (mdl *AccountQuotaMgr) getUpdateQuotaEvent(txLog types.Log) *updateQuotaEvent {
	updateQuota, err := mdl.contract.ParseUpdateQuota(txLog)
	log.CheckFatal(err)
	return &updateQuotaEvent{
		CreditAccount: updateQuota.CreditAccount,
		Token:         updateQuota.Token,
		QuotaChange:   updateQuota.QuotaChange,
		TxHash:        txLog.TxHash.Hex(),
		Index:         txLog.TxIndex,
		BlockNumber:   int64(txLog.BlockNumber),
	}
}
func (mdl *AccountQuotaMgr) AddAccountQuota(blockNum int64,
	txLog types.Log) {
	mdl.mu.Lock()
	defer mdl.mu.Unlock()
	//
	updateQuota := mdl.getUpdateQuotaEvent(txLog)
	account := updateQuota.CreditAccount.Hex()
	key := account + txLog.TxHash.Hex()
	mdl.events[key] = append(mdl.events[key], updateQuota)
}

// utils

// func GetQuotaIndexCurrent(currentTs uint64, prevDetails *schemas_v3.QuotaDetails) *core.BigInt {
// 	prevCumIndex := prevDetails.CumQuotaIndex.Convert()

// 	interestInPeriod := func() *big.Int {
// 		rayByPrecent := utils.GetExpInt(27 - 4)
// 		interestInYear := new(big.Int).Mul(rayByPrecent, big.NewInt(int64(currentTs-prevDetails.Timestamp)*int64(prevDetails.Rate)))
// 		return new(big.Int).Quo(interestInYear, big.NewInt(core.SECONDS_PER_YEAR))
// 	}()
// 	return (*core.BigInt)(new(big.Int).Add(interestInPeriod, prevCumIndex))
// }

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
