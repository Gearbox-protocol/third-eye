package main

import (
	"fmt"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg"
	"github.com/Gearbox-protocol/sdk-go/pkg/priceFetcher"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/repository"
)

// DATABASE_URL
// ETH_PROVIDER
// ETHERSCAN_KEY

type RewardAndTs struct {
	Value                  float64
	HowLongAgoFirstDeposit int64
}

type UserDeposit struct {
	User  string
	Token string
}

func main() {
	cfg := config.NewConfig()
	client := ethclient.NewEthClient(cfg)
	// node := pkg.Node{Client: client}
	db := repository.NewDBClient(cfg)
	var lastAllowedTs int64 = 1719792000
	lastblockAllowed, err := pkg.GetBlockNumForTs(utils.GetEnvOrDefault("ETHERSCAN_API_KEY", ""), core.GetChainId(client), lastAllowedTs)
	log.CheckFatal(err)
	data := []struct {
		User      string  `gorm:"column:user_address" json:"user"`
		Amount    float64 `gorm:"column:amount" json:"amount"`
		TimeStamp int64   `gorm:"column:timestamp"`
		Token     string  `gorm:"column:underlying_token"`
		Event     string  `gorm:"column:event" json:"event"`
	}{}
	log.Info(lastblockAllowed)
	err = db.Raw(`select p.underlying_token, a.*, b.timestamp from (select * from pool_ledger where event in  ('AddLiquidity', 'RemoveLiquidity') and pool in (select address from pools where _version=300) and block_num<?) a join blocks b on b.id=a.block_num join pools p on p.address=a.pool order by a.block_num`, lastblockAllowed).Find(&data).Error
	log.CheckFatal(err)
	ans := map[UserDeposit]*RewardAndTs{}
	for _, d := range data {
		v := float64(lastAllowedTs-d.TimeStamp) * d.Amount
		if d.Event == "RemoveLiquidity" {
			v = -v
		}
		key := UserDeposit{User: d.User, Token: d.Token}
		if _, ok := ans[key]; !ok {
			ans[key] = &RewardAndTs{
				HowLongAgoFirstDeposit: lastAllowedTs - d.TimeStamp,
			}
		}
		ans[key].Value += v
	}

	store := priceFetcher.NewTokensStore(client)
	orachel := priceFetcher.New1InchOracle(client, store, priceFetcher.URLsAndResolve{})
	calls := orachel.GetCalls()
	results := core.MakeMultiCall(client, lastblockAllowed, false, calls)
	prices := orachel.GetPrices(results, lastblockAllowed, uint64(lastAllowedTs))
	log.Info("here")
	//
	fmt.Println("User,Token,avgValue=sigma/HowLongAgoFirstDeposit,HowLongAgoFirstDeposit, sigma[deposit*(1julyts-depoitts)-withdraw*(1julyts-withdrawts)]")
	for user, entry := range ans {
		// decimals := store.GetDecimals(common.HexToAddress(user.Token))
		value := utils.GetFloat64Decimal(prices[user.Token].Convert(), 8) * entry.Value
		fmt.Printf("%s,%s,%f,%d,%f\n", user.User, user.Token, value/float64(entry.HowLongAgoFirstDeposit), entry.HowLongAgoFirstDeposit, value)
	}
}
