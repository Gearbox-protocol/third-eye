package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/repository"
)

type BorrowAndValue struct {
	TotalValue float64
	ClosedTs   int64
	StartedTs  int64
	LastTs     int64
	LastValue  float64
	schemas.CreditSession
}

func GetPureSessions( client core.ClientI) map[string]bool {
	chainId :=core.GetChainId(client)
	net:= string(log.GetBaseNet(chainId))
	
	f, err := os.OpenFile(fmt.Sprintf("scripts/points/trading_rewards/%s_pure.csv", strings.ToLower(net) ), os.O_RDONLY, 0)
	log.CheckFatal(err)
	scan := bufio.NewScanner(f)
	ans := map[string]bool{}
	for scan.Scan() {
		id := scan.Text()
		// log.Info(len(id))
		ans[id] = true
	}
	return ans
}
func main() {
	cfg := config.NewConfig()
	client := ethclient.NewEthClient(cfg)
	// node := pkg.Node{Client: client}
	db := repository.NewDBClient(cfg)

	//
	var lastAllowedTs int64 = 1719792000
	lastblockAllowed := pkg.GetBlockNum(uint64(lastAllowedTs), core.GetChainId(client))
	if lastblockAllowed == 0 {
		log.Fatal("lastblockAllowed is zero")
	}
	log.Info(lastblockAllowed)

	sessions := []struct {
		schemas.CreditSession
		ClosedTs int64 `gorm:"column:end_ts"`
	}{}
	err := db.Raw(`select a.*, b.timestamp end_ts from (select * from credit_sessions where version=300 and since < ? and credit_manager in (select address from credit_managers where _version=300 and name like '%Trade%' )) a left join blocks b on b.id=a.closed_at `, lastblockAllowed).Find(&sessions).Error
	log.Info("got sessions")
	log.CheckFatal(err)
	data := map[string]*BorrowAndValue{}

	for _, session := range sessions {
		data[session.ID] = &BorrowAndValue{CreditSession: session.CreditSession, ClosedTs: session.ClosedTs}
	}

	debts := []struct {
		schemas.Debt
		Ts int64 `gorm:"column:timestamp"`
	}{}
	err = db.Raw(`with cm as (select id from credit_sessions where version=300 and credit_manager in (select address from credit_managers where _version=300 and  name like '%Trade%')) select a.*, b.timestamp from (select * from debts where session_id in (select id from cm) and block_num <?) a join blocks b on a.block_num= b.id order by block_num `, lastblockAllowed).Find(&debts).Error
	log.Info("got debts")
	log.CheckFatal(err)
	for _, debt := range debts {
		a := data[debt.SessionId]
		if debt.CalHealthFactor.Convert().Cmp(big.NewInt(65535)) == 0 {
			if a.LastTs != 0 {
				a.TotalValue += float64(debt.Ts-a.LastTs) * a.LastValue
				a.LastTs = 0
				a.ClosedTs = debt.Ts
			}
			continue
		}
		if a.LastTs != 0 {
			a.TotalValue += float64(debt.Ts-a.LastTs) * a.LastValue
		} else {
			a.StartedTs = debt.Ts
		}
		a.LastValue = debt.TotalValueInUSD
		a.LastTs = debt.Ts
	}
	valid := GetPureSessions(client)
	fmt.Println("session, user, avg_total_value, started_ts, closed_ts, usd_value*holder_period ")
	for _, v := range data {
		if !valid[v.ID] {
			continue
		}
		lastts := lastAllowedTs
		if v.ClosedTs != 0 && v.ClosedTs < lastAllowedTs {
			lastts = v.ClosedTs
		}
		if v.LastTs != 0 {
			v.TotalValue += float64(lastts-v.LastTs) * v.LastValue
		}
		fmt.Printf("%s, %s, %f, %d, %d, %f\n", v.ID, v.Borrower, v.TotalValue/float64(lastts-v.StartedTs), v.StartedTs, v.ClosedTs, v.TotalValue)
	}
}
