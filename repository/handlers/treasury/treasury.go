package treasury

import (
	"math/big"
	"time"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/repository/handlers"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TreasuryRepo struct {
	// treasury
	treasurySnapshot *schemas.TreasurySnapshot
	lastTreasureTime time.Time
	tokens           *handlers.TokensRepo
	client           core.ClientI
	adapters         *handlers.SyncAdaptersRepo
	blocks           *handlers.BlocksRepo
}

func NewTreasuryRepo(tokens *handlers.TokensRepo, blocks *handlers.BlocksRepo, adapters *handlers.SyncAdaptersRepo, client core.ClientI) *TreasuryRepo {
	return &TreasuryRepo{
		tokens:   tokens,
		client:   client,
		adapters: adapters,
		blocks:   blocks,
	}
}

// load /save

func (repo *TreasuryRepo) LoadTreasurySnapshot(db *gorm.DB) {
	defer utils.Elapsed("loadTreasurySnapshot")()
	ss := schemas.TreasurySnapshot{}
	sql := `SELECT * FROM treasury_snapshots WHERE block_num=0`
	if err := db.Raw(sql).Find(&ss).Error; err != nil {
		log.Fatal(err)
	}
	if ss.Balances == nil {
		ss.Balances = &core.JsonFloatMap{}
	}
	repo.treasurySnapshot = &ss
}

func (repo *TreasuryRepo) LoadLastTreasuryTs(db *gorm.DB) {
	defer utils.Elapsed("loadLastTreasuryTs")()
	data := schemas.DebtSync{}
	if err := db.Raw(`SELECT timestamp AS last_calculated_at FROM blocks 
		WHERE id in (SELECT max(block_num) FROM treasury_snapshots)`).Find(&data).Error; err != nil {
		log.Fatal(err)
	}
	repo.lastTreasureTime = time.Unix(data.LastCalculatedAt, 0)
	if repo.lastTreasureTime.Unix() != 0 {
		repo.lastTreasureTime = utils.TimeToDateEndTime(repo.lastTreasureTime)
	}
}

func (repo *TreasuryRepo) Save(tx *gorm.DB) {
	utils.Elapsed("treasury sql statements")()
	if repo.treasurySnapshot.Date != "" {
		err := tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "block_num"}},
			DoUpdates: clause.AssignmentColumns([]string{"date_str", "prices_in_usd", "balances", "value_in_usd"}),
		}).Create(repo.treasurySnapshot).Error
		log.CheckFatal(err)
	}
}

// external funcs
func (repo *TreasuryRepo) AddTreasuryTransfer(blockNum int64, logID uint, token string, amount *big.Int) {
	block := repo.blocks.SetAndGetBlock(blockNum)
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
	tokenObj := repo.tokens.GetToken(token)
	amt := utils.GetFloat64Decimal(amount, tokenObj.Decimals)
	(*repo.treasurySnapshot.Balances)[token] = balance + amt
}

func (repo *TreasuryRepo) saveTreasurySnapshot() {
	ts := repo.lastTreasureTime.Unix()
	blockDate := repo.blocks.GetBlockDatePairs(ts)
	balances := core.JsonFloatMap{}
	for token, amt := range *repo.treasurySnapshot.Balances {
		balances[token] = amt
	}
	tss := &schemas.TreasurySnapshot{
		Date:     utils.TimeToDate(repo.lastTreasureTime),
		BlockNum: blockDate.BlockNum,
		Balances: &balances,
	}
	repo.calFieldsOfTreasurySnapshot(blockDate.BlockNum, tss)
	log.Info(utils.ToJson(tss))
	repo.blocks.SetAndGetBlock(blockDate.BlockNum).AddTreasurySnapshot(tss)
}

func (repo *TreasuryRepo) calFieldsOfTreasurySnapshot(blockNum int64, tss *schemas.TreasurySnapshot) {
	var totalValueInUSD float64
	var tokenAddrs []string
	for token := range *tss.Balances {
		if token == repo.tokens.GetGearTokenAddr() || token == repo.tokens.GetWETHAddr() {
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

func (repo *TreasuryRepo) CalCurrentTreasuryValue(blockNum int64) {
	repo.calFieldsOfTreasurySnapshot(blockNum, repo.treasurySnapshot)
}
