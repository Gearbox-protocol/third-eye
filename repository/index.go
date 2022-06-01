package repository

import (
	"sync"
	"time"

	"fmt"

	"github.com/Gearbox-protocol/sdk-go/artifacts/creditFilter"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/ds/dc_wrapper"
	"github.com/Gearbox-protocol/third-eye/models/aggregated_block_feed"
	"github.com/Gearbox-protocol/third-eye/repository/handlers"
	"gorm.io/gorm"
)

type Repository struct {
	// repos
	*handlers.SessionRepo
	*handlers.AllowedTokenRepo
	*handlers.ParamsRepo
	*handlers.PoolUsersRepo
	// mutex
	mu *sync.Mutex
	// object fx objects
	WETHAddr              string
	USDCAddr              string
	GearTokenAddr         string
	db                    *gorm.DB
	client                core.ClientI
	config                *config.Config
	kit                   *ds.AdapterKit
	executeParser         ds.ExecuteParserI
	dcWrapper             *dc_wrapper.DataCompressorWrapper
	aggregatedFeed        *aggregated_block_feed.AggregatedBlockFeed
	creditManagerToFilter map[string]*creditFilter.CreditFilter
	// blocks/token
	blocks map[int64]*schemas.Block
	tokens map[string]*schemas.Token
	// version  to token to oracle
	tokensCurrentOracle map[int16]map[string]*schemas.TokenOracle // done
	// treasury
	treasurySnapshot *schemas.TreasurySnapshot
	lastTreasureTime time.Time
	BlockDatePairs   map[int64]*schemas.BlockDate
	dieselTokens     map[string]*schemas.UTokenAndPool
	accountManager   *ds.AccountTokenManager
	relations        []*schemas.UniPriceAndChainlink
}

func GetRepository(db *gorm.DB, client core.ClientI, config *config.Config, ep ds.ExecuteParserI) *Repository {
	repo := &Repository{
		SessionRepo:           handlers.NewSessionRepo(),
		AllowedTokenRepo:      handlers.NewAllowedTokenRepo(),
		ParamsRepo:            handlers.NewParamsRepo(),
		PoolUsersRepo:         handlers.NewPoolUsersRepo(),
		mu:                    &sync.Mutex{},
		db:                    db,
		client:                client,
		config:                config,
		blocks:                make(map[int64]*schemas.Block),
		executeParser:         ep,
		kit:                   ds.NewAdapterKit(),
		tokens:                make(map[string]*schemas.Token),
		tokensCurrentOracle:   make(map[int16]map[string]*schemas.TokenOracle),
		dcWrapper:             dc_wrapper.NewDataCompressorWrapper(client),
		creditManagerToFilter: make(map[string]*creditFilter.CreditFilter),
		// for treasury to get the date
		BlockDatePairs: make(map[int64]*schemas.BlockDate),
		// for getting the diesel tokens
		dieselTokens:   make(map[string]*schemas.UTokenAndPool),
		accountManager: ds.NewAccountTokenManager(),
	}
	// aggregated block feed
	repo.aggregatedFeed = aggregated_block_feed.NewAggregatedBlockFeed(repo.client, repo, repo.config.Interval)
	repo.kit.Add(repo.aggregatedFeed)
	return repo
}

func NewRepository(db *gorm.DB, client core.ClientI, config *config.Config, ep ds.ExecuteParserI) ds.RepositoryI {
	r := GetRepository(db, client, config, ep)
	r.init()
	return r
}

func (repo *Repository) GetDCWrapper() *dc_wrapper.DataCompressorWrapper {
	return repo.dcWrapper
}

func (repo *Repository) GetExecuteParser() ds.ExecuteParserI {
	return repo.executeParser
}

func (repo *Repository) GetKit() *ds.AdapterKit {
	return repo.kit
}

func (repo *Repository) init() {
	lastDebtSync := repo.LoadLastDebtSync()
	// token should be loaded before syncAdapters as credit manager adapter uses underlying token details
	repo.loadToken()
	// syncadapter state for cm and pool is set after loading of pool/credit manager table data from db
	repo.loadSyncAdapters()
	repo.loadChainlinkPrevState()
	//
	repo.loadUniswapPools()
	// for disabling previous token oracle if new oracle is set
	repo.loadCurrentTokenOracle()
	// load state for sync_adapters
	repo.loadPool()
	repo.LoadPoolUniqueUsers(repo.db)
	// load credit manager
	repo.loadCreditManagers()
	repo.loadGearBalances()
	// required for disabling allowed tokens
	repo.loadAllowedTokensState()
	// fastcheck and new parameters
	repo.ParamsRepo.LoadAllParams(repo.db)
	// treasury funcs
	repo.loadBlockDatePair()
	repo.loadLastTreasuryTs()
	repo.loadTreasurySnapshot()
	// for direct token transfer
	repo.loadAccountLastSession()
	// credit_sessions
	repo.loadCreditSessions(lastDebtSync)
}

func (repo *Repository) AddAccountOperation(accountOperation *schemas.AccountOperation) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if accountOperation.SessionId == "" {
		panic(utils.ToJson(accountOperation))
	}
	repo.setAndGetBlock(accountOperation.BlockNumber).AddAccountOperation(accountOperation)
}

func (repo *Repository) SetWETHAddr(addr string) {
	repo.WETHAddr = addr
}

func (repo *Repository) GetWETHAddr() string {
	return repo.WETHAddr
}
func (repo *Repository) GetUSDCAddr() string {
	return repo.USDCAddr
}
func (repo *Repository) GetGearTokenAddr() string {
	return repo.GearTokenAddr
}

func (eng *Repository) RecentEventMsg(blockNum int64, msg string, args ...interface{}) {
	ts := eng.SetAndGetBlock(blockNum).Timestamp
	if time.Now().Sub(time.Unix(int64(ts), 0)) < time.Hour {
		log.Msgf(msg, args...)
	}
}

type LastSyncAndType struct {
	Type     string `gorm:"column:type"`
	LastSync int64  `gorm:"column:last_sync"`
	Address  string `gorm:"column:address"`
}

func (obj *LastSyncAndType) String() string {
	return fmt.Sprintf("%s(%s):%d", obj.Type, obj.Address, obj.LastSync)

}
func (repo *Repository) InitChecks() {
	data := []*LastSyncAndType{}
	err := repo.db.Raw(`SELECT type, address,  last_sync AS last_calculated_at 
		FROM sync_adapters 
		WHERE type IN ('AccountManager', 'CreditManager','AccountFactory')`).Find(&data).Error
	log.CheckFatal(err)
	var accountManagerLastSync, accountFactoryLastSync int64
	var cmLastSync int64
	var str string
	for _, entry := range data {
		str += entry.String()
		switch entry.Type {
		case ds.AccountFactory:
			accountFactoryLastSync = entry.LastSync
		case ds.AccountManager:
			accountManagerLastSync = entry.LastSync
		case ds.CreditManager:
			cmLastSync = utils.Min(entry.LastSync, cmLastSync)
		}
	}
	if accountFactoryLastSync != accountManagerLastSync ||
		cmLastSync < accountManagerLastSync {
		log.Fatal("Account manager/credit manager/AccountFactory are not synchronised: ", str)
	}
}

func (repo *Repository) GetChainId() uint {
	return repo.config.ChainId
}

func (repo *Repository) TransferAccountAllowed(obj *schemas.TransferAccountAllowed) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.setAndGetBlock(obj.BlockNumber).AddTransferAccountAllowed(obj)
}
