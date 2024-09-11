package main

import (
	"fmt"
	"time"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg"
	"github.com/Gearbox-protocol/sdk-go/pkg/priceFetcher"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/repository"
	"github.com/ethereum/go-ethereum/common"
)

// DATABASE_URL
// ETH_PROVIDER
// ETHERSCAN_KEY

type RewardAndTs struct {
	Value                  []float64
	// StartTs int64
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
	lastblockAllowed:= pkg.GetBlockNum( uint64( lastAllowedTs), core.GetChainId(client))
	if lastblockAllowed == 0 {
		log.Fatal("lastblockAllowed is zero")
	}
	data := []struct {
		User      string  `gorm:"column:user_address" json:"user"`
		Amount    float64 `gorm:"column:amount" json:"amount"`
		TimeStamp int64   `gorm:"column:timestamp"`
		Token     string  `gorm:"column:underlying_token"`
		Event     string  `gorm:"column:event" json:"event"`
		BlockNum int64 `gorm:"column:block_num"`
	}{}
	log.Info(lastblockAllowed)
	err := db.Raw(`select p.underlying_token, a.*, b.timestamp from (select * from pool_ledger where event in  ('AddLiquidity', 'RemoveLiquidity') and pool in (select address from pools where _version=300) and block_num<?) a join blocks b on b.id=a.block_num join pools p on p.address=a.pool order by a.block_num`, lastblockAllowed).Find(&data).Error
	log.CheckFatal(err)

	ans := map[string]map[string]float64{}
	syms := core.GetTokenToSymbolByChainId(core.GetChainId(client))
	final := map[UserDeposit]*RewardAndTs{}
	//
	ts:=utils.TimeToDateEndTime(time.Unix( data[0].TimeStamp,0)) 
	dataInd :=0
	for ts.Unix() <= lastAllowedTs {
		for {
			if dataInd >= len(data) {
				break
			}
			curAction :=data[dataInd]
			if curAction.TimeStamp > ts.Unix() {
				break
			}
			v := curAction.Amount  
			if curAction.Event == "RemoveLiquidity" {
				// log.Info("here")
				v = -v
			}
			if ans[curAction.Token] == nil {
				ans[curAction.Token] = map[string]float64{}
			}
			ans[curAction.Token][curAction.User] += v
			dataInd++
		}
		dayEndBlock := pkg.GetBlockNum(uint64( ts.Unix()), core.GetChainId(client))
		if dayEndBlock == 0 {
			log.Fatal("")
		}
		log.Info(ts.Format(time.DateOnly), dayEndBlock)
		for token, entry:= range ans {
			var t float64 = 0
			for user, e:= range entry {
				t+=e
				if final[UserDeposit{User: user, Token: token}] == nil {
					final[UserDeposit{User: user, Token: token}] = &RewardAndTs{}
				}
				val := e*getPrice(client, token, dayEndBlock, ts.Unix())
				if val > 10 {
					a := final[UserDeposit{User: user, Token: token}].Value
					a=   append(a, val)
					final[UserDeposit{User: user, Token: token}].Value = a
				}
			}
			price :=getPrice(client, token, dayEndBlock, ts.Unix())
			log.Info(syms[common.HexToAddress(token)], t, price*t)
		}
		log.Info("####")
		ts = ts.Add(time.Hour * 24)
	}


	//
	fmt.Println("User,Token,avgValue,noofDays")
	for user, entry := range final {
		// decimals := store.GetDecimals(common.HexToAddress(user.Token))
		// value := utils.GetFloat64Decimal(prices[user.Token].Convert(), 8) * entry.Value
		fmt.Printf("%s,%s,%f,%d\n", user.User, user.Token, sum(entry.Value)/float64(len(entry.Value)), len(entry.Value))
	}
}

func sum(a []float64) float64 {
	var s float64 = 0
	for _, v := range a {
		s += v
	}
	return s
}
var orachel *priceFetcher.OneInchOracle
var blockToPrice map[int64]map[string]*core.BigInt

func getPrice(client core.ClientI, token string, block, ts int64) float64 {
	if blockToPrice == nil {
		blockToPrice = map[int64]map[string]*core.BigInt{}
	}
	if blockToPrice[block] == nil {
		blockToPrice = map[int64]map[string]*core.BigInt{}
		if orachel == nil {
			store := priceFetcher.NewTokensStore(client)
			orachel = priceFetcher.New1InchOracle(client, store, priceFetcher.URLsAndResolve{})
		}
		calls := orachel.GetCalls()
		results := core.MakeMultiCall(client, block, false, calls)
		prices := orachel.GetPrices(results, block, uint64(ts))
		blockToPrice[block] = prices
	}
	return utils.GetFloat64Decimal(blockToPrice[block][token].Convert(), 8)
}