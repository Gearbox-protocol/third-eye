package pool_lmrewards

import (
	"log"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
)

type DieselBalance struct {
	Balance *core.BigInt `gorm:"column:balance"`
	User    string       `gorm:"primaryKey;column:user"`
	Diesel  string       `gorm:"primaryKey;column:diesel_sym"`
}

func (DieselBalance) TableName() string {
	return "diesel_balances"
}

func (mdl PoolLMRewards) GetDieselBalances() (dieselBalances []DieselBalance) {
	for tokenSym, balances := range mdl.dieselBalances {
		for user, balance := range balances {
			dieselBalances = append(dieselBalances, DieselBalance{
				Balance: (*core.BigInt)(balance),
				User:    user,
				Diesel:  tokenSym,
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
		mdl.dieselBalances[dieselBalance.Diesel][dieselBalance.User] = dieselBalance.Balance.Convert()
	}
}

type LMReward struct {
	User   string       `gorm:"primaryKey;column:user"`
	Reward *core.BigInt `gorm:"column:reward"`
}

func (LMReward) TableName() string {
	return "lm_rewards"
}

func (mdl PoolLMRewards) GetLMRewards() (rewards []LMReward) {
	for user, reward := range mdl.rewards {
		rewards = append(rewards, LMReward{
			User:   user,
			Reward: (*core.BigInt)(reward),
		})
	}
	return rewards
}

func (mdl PoolLMRewards) LoadLMRewards(rewards []LMReward) {
	for _, reward := range rewards {
		mdl.rewards[reward.User] = reward.Reward.Convert()
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
