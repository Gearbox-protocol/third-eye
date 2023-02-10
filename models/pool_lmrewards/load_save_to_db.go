package pool_lmrewards

import (
	"log"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

type DieselBalance struct {
	BalanceBI *core.BigInt `gorm:"column:balance_bi"`
	Balance   float64      `gorm:"column:balance"`
	User      string       `gorm:"primaryKey;column:user_address"`
	Diesel    string       `gorm:"primaryKey;column:diesel_sym"`
}

func (DieselBalance) TableName() string {
	return "diesel_balances"
}

func (mdl PoolLMRewards) GetDieselBalances() (dieselBalances []DieselBalance) {
	for tokenSym, balances := range mdl.dieselBalances {
		decimals := mdl.decimalsAndPool[tokenSym].decimals
		for user, balanceBI := range balances {
			dieselBalances = append(dieselBalances, DieselBalance{
				BalanceBI: (*core.BigInt)(balanceBI),
				User:      user,
				Diesel:    tokenSym,
				Balance:   utils.GetFloat64Decimal(balanceBI, decimals),
			})
		}
	}
	return dieselBalances
}

func (mdl PoolLMRewards) LoadDieselBalances(dieselBalances []DieselBalance) {
	for _, dieselBalance := range dieselBalances {
		if _, ok := mdl.dieselBalances[dieselBalance.Diesel]; !ok {
			mdl.dieselBalances[dieselBalance.Diesel] = map[string]*big.Int{}
		}
		mdl.dieselBalances[dieselBalance.Diesel][dieselBalance.User] = dieselBalance.BalanceBI.Convert()
	}
}

type LMReward struct {
	User   string       `gorm:"primaryKey;column:user_address"`
	Pool   string       `gorm:"primaryKey;column:pool"`
	Reward *core.BigInt `gorm:"column:reward"`
}

func (LMReward) TableName() string {
	return "lm_rewards"
}

func (mdl PoolLMRewards) GetLMRewards() (rewards []LMReward) {
	for pool, rewardForUsers := range mdl.rewards {
		for user, reward := range rewardForUsers {
			rewards = append(rewards, LMReward{
				User:   user,
				Pool:   pool,
				Reward: (*core.BigInt)(reward),
			})
		}
	}
	return rewards
}

func (mdl PoolLMRewards) LoadLMRewards(rewards []LMReward) {
	for _, reward := range rewards {
		if mdl.rewards[reward.Pool] == nil {
			mdl.rewards[reward.Pool] = map[string]*big.Int{}
		}
		mdl.rewards[reward.Pool][reward.User] = reward.Reward.Convert()
	}
}

func (mdl *PoolLMRewards) totalSuppliesToDetails() {
	supplies := core.Json{}
	for tokenSym, totalSupply := range mdl.totalSupplies {
		supplies[tokenSym] = (*core.BigInt)(totalSupply)
	}
	mdl.Details = supplies
}

func toBigInt(x interface{}) *big.Int {
	switch v := x.(type) {
	case *big.Int:
		return v
	case string:
		num, b := new(big.Int).SetString(v, 10)
		if !b {
			log.Fatal("")
		}
		return num
	}
	return nil
}
func (mdl *PoolLMRewards) detailsToTotalSupplies() {
	supplies := map[string]*big.Int{}
	for tokenSym, totalSupply := range mdl.Details {
		supplies[tokenSym] = toBigInt(totalSupply)
	}
	mdl.totalSupplies = supplies
}
