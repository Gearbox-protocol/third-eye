package treasury

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg/redstone"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
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
	// for getting the block num from ts when it is missing from db
	//
	redstoneMgr redstone.RedStoneMgrI
}

func NewTreasuryRepo(tokens *handlers.TokensRepo, blocks *handlers.BlocksRepo, adapters *handlers.SyncAdaptersRepo, client core.ClientI, cfg *config.Config) *TreasuryRepo {
	return &TreasuryRepo{
		tokens:      tokens,
		client:      client,
		adapters:    adapters,
		blocks:      blocks,
		redstoneMgr: redstone.NewRedStoneMgr(client),
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
	if ss.OperationalBalances == nil {
		ss.OperationalBalances = &core.JsonFloatMap{}
	}
	repo.treasurySnapshot = &ss
}

func (repo *TreasuryRepo) LoadLastTreasuryTs(db *gorm.DB) {
	defer utils.Elapsed("loadLastTreasuryTs")()
	data := struct {
		LastCalculatedAt int64 `json:"last_calculated_at"`
	}{}
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
			DoUpdates: clause.AssignmentColumns([]string{"date_str", "prices_in_usd", "balances", "value_in_usd", "operational_value_in_usd", "operational_balances"}),
		}).Create(repo.treasurySnapshot).Error
		log.CheckFatal(err)
	}
}

// external funcs
func (repo *TreasuryRepo) AddTreasuryTransfer(blockNum int64, logID uint,
	token string, amount *big.Int, operationTransfer bool) {
	block := repo.blocks.SetAndGetBlock(blockNum)
	// treasury transfer
	block.AddTreasuryTransfer(&schemas.TreasuryTransfer{
		BlockNum:            blockNum,
		LogID:               logID,
		Token:               token,
		Amount:              (*core.BigInt)(amount),
		OperationalTransfer: operationTransfer,
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
		// SPECIAL CASE
		// for kovan, this check is there for edge case in the redeployment of gearbox for testing v2 on kovan.
		// the events and then chainlink/yearn feeds were missing for 29 june,
		// so we don't have the blockNum for that date, as a result the snapshot is missing for that date
		// ignore only this transfer of 30 june which tries to save snapshot of 29 june.
		chainId, err := repo.client.ChainID(context.TODO())
		log.CheckFatal(err)
		if !(blockNum == 32476006 && chainId.Int64() == 42) {
			repo.saveTreasurySnapshot()
		}
	}
	// set the current treasury snapshot fields
	repo.treasurySnapshot.Date = utils.TimeToDate(currentTime)
	tokenObj := repo.tokens.GetToken(token)
	amt := utils.GetFloat64Decimal(amount, tokenObj.Decimals)
	balance := (*repo.treasurySnapshot.Balances)[token]
	(*repo.treasurySnapshot.Balances)[token] = balance + amt
	if operationTransfer {
		operationalBalance := (*repo.treasurySnapshot.OperationalBalances)[token]
		(*repo.treasurySnapshot.OperationalBalances)[token] = operationalBalance + amt
	}
}

// saves snapshot for previous/last day
func (repo *TreasuryRepo) saveTreasurySnapshot() {
	ts := repo.lastTreasureTime.Unix()
	blockDate := repo.blocks.GetBlockDatePairs(ts)
	if blockDate == nil {
		key := fmt.Sprintf("%s_API_KEY", log.GetNetworkName(core.GetChainId(repo.client)))
		if utils.GetEnvOrDefault(key, "") != "" {
			if blockNum := core.GetBlockNum(uint64(ts), core.GetChainId(repo.client)); blockNum != 0 {
				repo.blocks.SetBlock(blockNum)
				blockDate = &schemas.BlockDate{
					BlockNum:  blockNum,
					Timestamp: ts,
				}
			} else {
				log.Fatal("Etherscan err")
			}
		} else {
			log.Fatalf("can't find the blocknum for ts(%d) and etherscan api key is also missing", ts)
		}
	}

	tss := &schemas.TreasurySnapshot{
		Date:                utils.TimeToDate(repo.lastTreasureTime),
		BlockNum:            blockDate.BlockNum,
		Balances:            repo.treasurySnapshot.Balances.Copy(),
		OperationalBalances: repo.treasurySnapshot.OperationalBalances.Copy(),
	}
	repo.calFieldsOfTreasurySnapshot(blockDate.BlockNum, tss)
	log.Info(utils.ToJson(tss))
	repo.blocks.SetAndGetBlock(blockDate.BlockNum).AddTreasurySnapshot(tss)
}

func (repo *TreasuryRepo) calFieldsOfTreasurySnapshot(blockNum int64, tss *schemas.TreasurySnapshot) {
	var tokenAddrs []string
	for token := range *tss.Balances {
		if token == repo.tokens.GetGearTokenAddr() {
			continue
		}
		tokenAddrs = append(tokenAddrs, token)
	}
	//
	priceOracle, version, _ := repo.adapters.GetActivePriceOracleByBlockNum(blockNum)
	prices := repo.getPricesInUSD(blockNum, priceOracle, version, tokenAddrs)
	//
	tss.PricesInUSD = &prices
	tss.ValueInUSD = tss.Balances.ValueInUSD(prices)
	tss.OperationalValueInUSD = tss.OperationalBalances.ValueInUSD(prices)
}

func (repo *TreasuryRepo) CalCurrentTreasuryValue(blockNum int64) {
	repo.calFieldsOfTreasurySnapshot(blockNum, repo.treasurySnapshot)
}
