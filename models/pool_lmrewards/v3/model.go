package v3

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/common"
)

// farmingPool https://etherscan.io/address/0x9ef444a6d7f4a5adcd68fd5329aa5240c90e14d2#code
// farmAccounting
// userAccounting
// farmingLib
type UserLMDetails struct {
	Correction      *core.BigInt `gorm:"column:correction"`
	FarmedBalanceBI *core.BigInt `gorm:"column:farmed_balance_bi"`
	FarmedBalance   float64      `gorm:"column:farmed_balance"`
	Account         string       `gorm:"column:account;primaryKey"`
	Farm            string       `gorm:"column:farm;primaryKey"`
	DieselSym       string       `gorm:"column:diesel_sym"`
	updated         bool         `gorm:"-"`
	DieselBalanceBI *core.BigInt `gorm:"column:diesel_balance"`
}

func (UserLMDetails) TableName() string {
	return "user_lmdetails_v3"
}

func (user UserLMDetails) GetPoints(farm *Farmv3, currentTs uint64) *big.Int {
	fpt := farm.calcFarmedPerToken(currentTs)
	num := new(big.Int).Mul(user.FarmedBalanceBI.Convert(), fpt)

	//
	return new(big.Int).Quo(new(big.Int).Sub(num, user.Correction.Convert()), _SCALE)
}

func (details *UserLMDetails) AddBalances(amount *big.Int, decimals int8) {
	details.updated = true
	details.FarmedBalanceBI = (*core.BigInt)(new(big.Int).Add(details.FarmedBalanceBI.Convert(), amount))
	details.FarmedBalance = utils.GetFloat64Decimal(details.FarmedBalanceBI.Convert(), decimals)
}
func (details *UserLMDetails) SubBalances(amount *big.Int, decimals int8) {
	details.updated = true
	details.AddBalances(new(big.Int).Neg(amount), decimals)
}
func (details *UserLMDetails) AddCorrection(amount *big.Int) {
	details.updated = true
	details.Correction = (*core.BigInt)(new(big.Int).Add(details.Correction.Convert(), amount))
}
func (details *UserLMDetails) SubCorrection(amount *big.Int) {
	details.updated = true
	details.AddCorrection(new(big.Int).Neg(amount))
}

type LMRewardsv3 struct {
	*ds.SyncAdapter
	// farm by farm address
	farms map[string]*Farmv3
	pools map[common.Address]string
	// farmv3 to user to balance
	users map[common.Address]map[string]*UserLMDetails
}

func NewLMRewardsv3(addr string, syncedTill int64, client core.ClientI, repo ds.RepositoryI) *LMRewardsv3 {
	return NewLMRewardsv3FromAdapter(
		&ds.SyncAdapter{
			SyncAdapterSchema: &schemas.SyncAdapterSchema{
				LastSync: syncedTill,
				Contract: &schemas.Contract{
					ContractName: ds.LMRewardsv3,
					Address:      addr,
					Client:       client,
				},
				V: core.NewVersion(300),
			},
			Repo: repo,
		},
	)
}

func NewLMRewardsv3FromAdapter(adapter *ds.SyncAdapter) *LMRewardsv3 {
	// chainId, err := adapter.Client.ChainID(context.Background())
	// log.CheckFatal(err)
	obj := &LMRewardsv3{
		SyncAdapter: adapter,
	}
	return obj
}

func (mdl *LMRewardsv3) AfterSyncHook(syncedTill int64) {
	mdl.SyncAdapter.AfterSyncHook(syncedTill)
}
