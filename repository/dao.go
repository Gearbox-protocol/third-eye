package repository

import (
	"github.com/Gearbox-protocol/third-eye/artifacts/poolService"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"time"
)

func (repo *Repository) AddDAOOperation(operation *core.DAOOperation) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.addDAOOperation(operation)
}

func (repo *Repository) addDAOOperation(operation *core.DAOOperation) {
	repo.setAndGetBlock(operation.BlockNumber).AddDAOOperation(operation)
}

func (repo *Repository) AddTreasuryTransfer(blockNum int64, logID uint, token string, amount *big.Int) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	block := repo.setAndGetBlock(blockNum)
	// treasury transfer
	block.AddTreasuryTransfer(&core.TreasuryTransfer{
		BlockNum: blockNum,
		LogID:    logID,
		Token:    token,
		Amount:   (*core.BigInt)(amount),
	})
	// treasury snapshots
	currentTime := time.Unix(int64(block.Timestamp), 0)
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
	tokenObj, err := repo.getTokenWithError(token)
	log.CheckFatal(err)
	amt := utils.GetFloat64Decimal(amount, tokenObj.Decimals)
	(*repo.treasurySnapshot.Balances)[token] = balance + amt
}

func (repo *Repository) saveTreasurySnapshot() {
	ts := repo.lastTreasureTime.Unix()
	blockDate := repo.BlockDatePairs[ts]
	balances := core.JsonFloatMap{}
	for token, amt := range *repo.treasurySnapshot.Balances {
		balances[token] = amt
	}
	tss := &core.TreasurySnapshot{
		Date:     utils.TimeToDate(repo.lastTreasureTime),
		BlockNum: blockDate.BlockNum,
		Balances: &balances,
	}
	repo.CalFieldsOfTreasurySnapshot(blockDate.BlockNum, tss)
	log.Info(utils.ToJson(tss))
	repo.setAndGetBlock(blockDate.BlockNum).AddTreasurySnapshot(tss)
}

func (repo *Repository) CalFieldsOfTreasurySnapshot(blockNum int64, tss *core.TreasurySnapshot) {
	var totalValueInUSD float64
	prices := core.JsonFloatMap{}
	for token, amt := range *tss.Balances {
		price := repo.GetPriceInUSD(blockNum, token)
		prices[token] = utils.GetFloat64Decimal(price, 8)
		totalValueInUSD += amt * prices[token]
	}
	tss.PricesInUSD = &prices
	tss.ValueInUSD = totalValueInUSD
}

func (repo *Repository) CalCurrentTreasuryValue(blockNum int64) {
	repo.CalFieldsOfTreasurySnapshot(blockNum, repo.treasurySnapshot)
}

func (repo *Repository) GetPriceInUSD(blockNum int64, token string) (price *big.Int) {
	tokenObj, err := repo.getTokenWithError(token)
	log.CheckFatal(err)
	uTokenAndPool := repo.dieselTokens[token]
	if token == repo.GearTokenAddr {
		price = big.NewInt(0)
	} else if uTokenAndPool != nil {
		price = repo.GetValueInUSD(blockNum, uTokenAndPool.UToken, utils.GetExpInt(tokenObj.Decimals))
		pool, err := poolService.NewPoolService(common.HexToAddress(uTokenAndPool.Pool), repo.client)
		log.CheckFatal(err)
		opts := &bind.CallOpts{
			BlockNumber: big.NewInt(blockNum),
		}
		dieselRate, err := pool.GetDieselRateRAY(opts)
		log.CheckFatal(err)
		price = new(big.Int).Mul(dieselRate, price)
		price = utils.GetInt64(price, 27)
	} else {
		price = repo.GetValueInUSD(blockNum, token, utils.GetExpInt(tokenObj.Decimals))
	}
	return
}

func (repo *Repository) loadLastTreasuryTs() {
	data := core.DebtSync{}
	if err := repo.db.Raw(`SELECT timestamp AS last_calculated_at FROM blocks 
		WHERE id in (SELECT max(block_num) FROM treasury_snapshots)`).Find(&data).Error; err != nil {
		log.Fatal(err)
	}
	repo.lastTreasureTime = utils.TimeToDateEndTime(time.Unix(data.LastCalculatedAt, 0))
}

func (repo *Repository) loadBlockDatePair() {
	data := []*core.BlockDate{}
	sql := `select b.*, a.timestamp from blocks a 
	JOIN (select timezone('UTC', to_timestamp(timestamp))::date as date,max(id) as block_num from blocks group by date) b 
	ON a.id=b.block_num order by block_num`
	if err := repo.db.Raw(sql).Find(&data).Error; err != nil {
		log.Fatal(err)
	}
	for _, entry := range data {
		repo.addBlockDate(entry)
	}
}

func (repo *Repository) addBlockDate(entry *core.BlockDate) {
	ts := utils.TimeToDateEndTs(time.Unix(entry.Timestamp, 0))
	previousEntry := repo.BlockDatePairs[ts]
	if previousEntry == nil || previousEntry.BlockNum < entry.BlockNum {
		repo.BlockDatePairs[ts] = entry
	}
}

func (repo *Repository) loadTreasurySnapshot() {
	ss := core.TreasurySnapshot{}
	sql := `SELECT * FROM treasury_snapshots WHERE block_num=0`
	if err := repo.db.Raw(sql).Find(&ss).Error; err != nil {
		log.Fatal(err)
	}
	if ss.Balances == nil {
		ss.Balances = &core.JsonFloatMap{}
	}
	repo.treasurySnapshot = &ss
}
