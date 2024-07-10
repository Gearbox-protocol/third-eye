package v2

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/models/pool_lmrewards"
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

func (mdl LMRewardsv2) GetDieselBalances() (dieselBalances []DieselBalance) {
	if !mdl.hasDataToSave {
		return
	}
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

func (mdl LMRewardsv2) LoadDieselBalances(dieselBalances []DieselBalance) {
	for _, dieselBalance := range dieselBalances {
		if _, ok := mdl.dieselBalances[dieselBalance.Diesel]; !ok {
			mdl.dieselBalances[dieselBalance.Diesel] = map[string]*big.Int{}
		}
		mdl.dieselBalances[dieselBalance.Diesel][dieselBalance.User] = dieselBalance.BalanceBI.Convert()
	}
}

func (mdl LMRewardsv2) GetLMRewards() (rewards []pool_lmrewards.LMReward) {
	if !mdl.hasDataToSave {
		return
	}
	// only for v2
	gearToken := core.GetSymToAddrByChainId(mdl.chainId).Tokens["GEAR"]
	for pool, rewardForUsers := range mdl.rewards {
		for user, reward := range rewardForUsers {
			rewards = append(rewards, pool_lmrewards.LMReward{
				User:        user,
				Pool:        pool,
				Reward:      (*core.BigInt)(reward),
				RewardToken: gearToken.Hex(),
			})
		}
	}
	return rewards
}

func (mdl *LMRewardsv2) SyncComplete() {
	mdl.hasDataToSave = false
}

func (mdl LMRewardsv2) LoadLMRewards(rewards []pool_lmrewards.LMReward) {
	for _, reward := range rewards {
		if mdl.rewards[reward.Pool] == nil {
			mdl.rewards[reward.Pool] = map[string]*big.Int{}
		}
		mdl.rewards[reward.Pool][reward.User] = reward.Reward.Convert()
	}
}

func (mdl *LMRewardsv2) totalSuppliesToDetails() {
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
func (mdl *LMRewardsv2) detailsToTotalSupplies() {
	supplies := map[string]*big.Int{}
	for tokenSym, totalSupply := range mdl.Details {
		supplies[tokenSym] = toBigInt(totalSupply)
	}
	mdl.totalSupplies = supplies
}
