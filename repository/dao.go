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
			year, month, day := currentTime.Date()
			repo.lastTreasureTime = time.Date(year, month, day-1, 0, 0, 0, 0, time.UTC)
		} else {
			year, month, day := repo.lastTreasureTime.Date()
			repo.lastTreasureTime = time.Date(year, month, day+1, 0, 0, 0, 0, time.UTC)
		}
		repo.saveTreasurySnapshot()
	}
	// set the current treasury snapshot fields
	repo.treasurySnapshot.Date = utils.TimeToDate(currentTime)
	balance := (*repo.treasurySnapshot.Balances)[token]
	(*repo.treasurySnapshot.Balances)[token] = core.AddCoreAndInt(balance, amount)
}

func (repo *Repository) saveTreasurySnapshot() {
	ts := repo.lastTreasureTime.Unix()
	blockDate := repo.BlockDatePairs[ts]
	balances := core.JsonBigIntMap{}
	for token, amt := range *repo.treasurySnapshot.Balances {
		balances[token] = core.NewBigInt(amt)
	}
	tss := &core.TreasurySnapshot{
		Date:      utils.TimeToDate(repo.lastTreasureTime),
		Timestamp: blockDate.Timestamp,
		Balances:  &balances,
	}
	repo.CalFieldsOfTreasurySnapshot(blockDate.BlockNum, tss)
	repo.setAndGetBlock(blockDate.BlockNum).AddTreasurySnapshot(tss)
}

func (repo *Repository) CalFieldsOfTreasurySnapshot(blockNum int64, tss *core.TreasurySnapshot) {
	var totalValueInUSD float64
	prices := core.JsonBigIntMap{}
	for token, amt := range *tss.Balances {
		price, value := repo.GetPriceAndValInUSD(blockNum, token, amt.Convert())
		prices[token] = (*core.BigInt)(price)
		valueInUSD := utils.GetFloat64Decimal(value, 8)
		totalValueInUSD += valueInUSD
	}
	tss.ValueInUSD = totalValueInUSD
}

func (repo *Repository) CalCurrentTreasuryValue(blockNum int64) {
	repo.CalFieldsOfTreasurySnapshot(blockNum, repo.treasurySnapshot)
}

func (repo *Repository) GetPriceAndValInUSD(blockNum int64, token string, amount *big.Int) (price, value *big.Int) {
	tokenObj := repo.GetToken(token)
	uTokenAndPool := repo.dieselTokens[token]
	if uTokenAndPool != nil {
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
	value = utils.GetInt64(new(big.Int).Mul(price, amount), 8)
	return
}

func (repo *Repository) loadLastTreasuryTs() {
	data := core.DebtSync{}
	if err := repo.db.Raw("SELECT max(ts) AS last_calculated_at FROM treasury_snapshots").Find(&data).Error; err != nil {
		log.Fatal(err)
	}
	repo.lastTreasureTime = time.Unix(data.LastCalculatedAt, 0)
}

func (repo *Repository) loadBlockDatePair() {
	data := []*core.BlockDate{}
	sql := `select b.*, a.timestamp from blocks a 
	JOIN (select to_timestamp(timestamp)::date as date,max(id) as block_num from blocks group by date) b 
	ON a.id=b.block_num order by block_num`
	if err := repo.db.Raw(sql).Find(&data).Error; err != nil {
		log.Fatal(err)
	}
	for _, entry := range data {
		repo.addBlockDate(entry)
	}
}

func (repo *Repository) addBlockDate(entry *core.BlockDate) {
	repo.BlockDatePairs[utils.TsToDateStartTs(entry.Timestamp)] = entry
}

func (repo *Repository) loadTreasurySnapshot() {
	ss := core.TreasurySnapshot{}
	sql := `SELECT * FROM treasury_snapshots WHERE ts=0`
	if err := repo.db.Raw(sql).Find(&ss).Error; err != nil {
		log.Fatal(err)
	}
	ss.Balances = &core.JsonBigIntMap{}
	repo.treasurySnapshot = &ss
}
