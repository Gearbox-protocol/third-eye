package repository

import (
	"math/big"
	"time"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

func (repo *Repository) AddTreasuryTransfer(blockNum int64, logID uint, token string, amount *big.Int) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	block := repo.SetAndGetBlock(blockNum)
	// treasury transfer
	block.AddTreasuryTransfer(&schemas.TreasuryTransfer{
		BlockNum: blockNum,
		LogID:    logID,
		Token:    token,
		Amount:   (*core.BigInt)(amount),
	})
	// treasury snapshots
	currentTime := time.Unix(int64(block.Timestamp), 0).UTC()
	// log.Info(utils.TimeToDate(currentTime), utils.TimeToDate(repo.lastTreasureTime))
	for currentTime.Sub(repo.lastTreasureTime) >= 24*time.Hour {
		if repo.lastTreasureTime.Unix() == 0 {
			repo.lastTreasureTime = utils.TimeToDateEndTime(currentTime.AddDate(0, 0, -1))
		} else {
			repo.lastTreasureTime = utils.TimeToDateEndTime(repo.lastTreasureTime.AddDate(0, 0, 1))
		}
		repo.saveTreasurySnapshot()
	}
	// set the current treasury snapshot fields
	repo.treasurySnapshot.Date = utils.TimeToDate(currentTime)
	balance := (*repo.treasurySnapshot.Balances)[token]
	tokenObj := repo.GetToken(token)
	amt := utils.GetFloat64Decimal(amount, tokenObj.Decimals)
	(*repo.treasurySnapshot.Balances)[token] = balance + amt
}

func (repo *Repository) saveTreasurySnapshot() {
	ts := repo.lastTreasureTime.Unix()
	blockDate := repo.GetBlockDatePairs(ts)
	balances := core.JsonFloatMap{}
	for token, amt := range *repo.treasurySnapshot.Balances {
		balances[token] = amt
	}
	tss := &schemas.TreasurySnapshot{
		Date:     utils.TimeToDate(repo.lastTreasureTime),
		BlockNum: blockDate.BlockNum,
		Balances: &balances,
	}
	repo.CalFieldsOfTreasurySnapshot(blockDate.BlockNum, tss)
	log.Info(utils.ToJson(tss))
	repo.SetAndGetBlock(blockDate.BlockNum).AddTreasurySnapshot(tss)
}

func (repo *Repository) CalFieldsOfTreasurySnapshot(blockNum int64, tss *schemas.TreasurySnapshot) {
	var totalValueInUSD float64
	var tokenAddrs []string
	for token := range *tss.Balances {
		if token == repo.GetGearTokenAddr() || token == repo.GetWETHAddr() {
			continue
		}
		tokenAddrs = append(tokenAddrs, token)
	}
	prices := repo.GetPricesInUSD(blockNum, tokenAddrs)
	for token, amt := range *tss.Balances {
		totalValueInUSD += amt * prices[token]
	}
	tss.PricesInUSD = &prices
	tss.ValueInUSD = totalValueInUSD
}

func (repo *Repository) AfterSync(syncTill int64) {
	// for direct token transfer
	for _, txs := range repo.accountManager.GetNoSessionTxs() {
		for _, tx := range txs {
			repo.RecentEventMsg(tx.BlockNum, "No session account token transfer: %v", tx)
			repo.SetAndGetBlock(tx.BlockNum).AddNoSessionTx(tx)
		}
	}
	// for direct token transfer
	repo.accountManager.Clear()
	// chainlink and uniswap prices
	repo.AggregatedFeed.Clear()
}
func (repo *Repository) CalCurrentTreasuryValue(blockNum int64) {
	repo.CalFieldsOfTreasurySnapshot(blockNum, repo.treasurySnapshot)
}

// used for treasury calculation and for remainingFunds on close v2
func (repo *Repository) GetPricesInUSD(blockNum int64, tokenAddrs []string) core.JsonFloatMap {
	priceByToken := core.JsonFloatMap{}
	var tokenForCalls []string
	var poolForDieselRate []string
	for _, token := range tokenAddrs {
		uTokenAndPool := repo.GetDieselToken(token)
		if uTokenAndPool != nil {
			tokenForCalls = append(tokenForCalls, uTokenAndPool.UToken)
			poolForDieselRate = append(poolForDieselRate, uTokenAndPool.Pool)
		} else {
			tokenForCalls = append(tokenForCalls, token)
		}
	}
	priceOracle, _ := repo.GetActivePriceOracleByBlockNum(blockNum)
	prices, dieselRates := repo.getPricesInBatch(priceOracle, blockNum, false, tokenForCalls, poolForDieselRate)
	var poolIndex int
	for i, token := range tokenAddrs {
		var price *big.Int
		if repo.IsDieselToken(token) {
			dieselRate := dieselRates[poolIndex]
			poolIndex++
			price = new(big.Int).Mul(dieselRate, prices[i])
			price = utils.GetInt64(price, 27)
		} else {
			price = prices[i]
		}
		priceByToken[token] = utils.GetFloat64Decimal(price, 8)
	}
	if repo.GetKit().GetAdapter(priceOracle).GetVersion() == 1 {
		priceByToken[repo.GetWETHAddr()] = 1
	}
	return priceByToken
}

func (repo *Repository) loadLastTreasuryTs() {
	defer utils.Elapsed("loadLastTreasuryTs")()
	data := schemas.DebtSync{}
	if err := repo.db.Raw(`SELECT timestamp AS last_calculated_at FROM blocks 
		WHERE id in (SELECT max(block_num) FROM treasury_snapshots)`).Find(&data).Error; err != nil {
		log.Fatal(err)
	}
	repo.lastTreasureTime = time.Unix(data.LastCalculatedAt, 0)
	if repo.lastTreasureTime.Unix() != 0 {
		repo.lastTreasureTime = utils.TimeToDateEndTime(repo.lastTreasureTime)
	}
}

func (repo *Repository) loadTreasurySnapshot() {
	defer utils.Elapsed("loadTreasurySnapshot")()
	ss := schemas.TreasurySnapshot{}
	sql := `SELECT * FROM treasury_snapshots WHERE block_num=0`
	if err := repo.db.Raw(sql).Find(&ss).Error; err != nil {
		log.Fatal(err)
	}
	if ss.Balances == nil {
		ss.Balances = &core.JsonFloatMap{}
	}
	repo.treasurySnapshot = &ss
}
