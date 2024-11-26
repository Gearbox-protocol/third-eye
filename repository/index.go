package repository

import (
	"sync"

	"fmt"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/repository/handlers"
	"github.com/Gearbox-protocol/third-eye/repository/handlers/treasury"
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository struct {
	// repos
	*handlers.SessionRepo
	*handlers.AllowedTokenRepo
	*handlers.ParamsRepo
	*handlers.PoolUsersRepo
	*handlers.BlocksRepo
	*handlers.TokensRepo
	*handlers.ExtrasRepo
	*handlers.SyncAdaptersRepo
	*handlers.TokenOracleRepo
	*treasury.TreasuryRepo
	// mutex
	mu *sync.Mutex
	// object fx objects
	db              *gorm.DB
	client          core.ClientI
	config          *config.Config
	accountManager  *ds.DirectTransferManager
	AccountQuotaMgr *ds.AccountQuotaMgr
	//
	feedToTicker map[string]common.Address // feed to ticker
}

func GetRepository(db *gorm.DB, client core.ClientI, cfg *config.Config, extras *handlers.ExtrasRepo) *Repository {
	tokensRepo := handlers.NewTokensRepo(client)
	blocksRepo := handlers.NewBlocksRepo(db, client, cfg, tokensRepo)
	repo := &Repository{
		SessionRepo:      handlers.NewSessionRepo(),
		AllowedTokenRepo: handlers.NewAllowedTokenRepo(blocksRepo, tokensRepo),
		ParamsRepo:       handlers.NewParamsRepo(blocksRepo),
		PoolUsersRepo:    handlers.NewPoolUsersRepo(),
		TokensRepo:       tokensRepo,
		BlocksRepo:       blocksRepo,
		ExtrasRepo:       extras,
		mu:               &sync.Mutex{},
		db:               db,
		client:           client,
		config:           cfg,
		accountManager:   ds.NewDirectTransferManager(),
		feedToTicker:     map[string]common.Address{},
	}
	repo.SyncAdaptersRepo = handlers.NewSyncAdaptersRepo(client, repo, cfg, extras)
	repo.TokenOracleRepo = handlers.NewTokenOracleRepo(repo.SyncAdaptersRepo, blocksRepo, repo, client)
	repo.TreasuryRepo = treasury.NewTreasuryRepo(tokensRepo, blocksRepo, repo.SyncAdaptersRepo, client, cfg)
	repo.AccountQuotaMgr = ds.NewAccountQuotaMgr(repo.client)
	return repo
}

func NewRepository(db *gorm.DB, client core.ClientI, config *config.Config, ep *handlers.ExtrasRepo) ds.RepositoryI {
	r := GetRepository(db, client, config, ep)
	return r
}

func (r Repository) GetDB() *gorm.DB {
	return r.db
}

func (repo *Repository) Init() {
	// lastdebtsync is required to load credit session which are active or closed after lastdebtsync block number
	lastDebtSync := repo.LoadLastDebtSync()
	// token should be loaded before syncAdapters as credit manager adapter uses underlying token details
	repo.TokensRepo.LoadTokens(repo.db)

	repo.loadDieselToken()
	// syncadapter state for cm and pool is set after loading of pool/credit manager table data from db
	repo.SyncAdaptersRepo.LoadSyncAdapters(repo.db)
	// load poolLMrewards
	repo.loadLMRewardDetailsv2()
	repo.loadLMRewardDetailsv3()
	//
	// for disabling previous token oracle if new oracle is set
	repo.LoadCurrentTokenOracle(repo.db)
	// load state for sync_adapters
	repo.loadPool()
	repo.LoadPoolUniqueUsers(repo.db)
	repo.loadQuotaDetails()
	// repo.loadAccountQuotaInfo()
	// load credit manager
	repo.loadCreditManagers()
	repo.loadGearBalances()
	// required for disabling allowed tokens
	repo.LoadAllowedTokensState(repo.db)
	// fastcheck and new parameters
	repo.ParamsRepo.LoadAllParams(repo.db)
	// blocks load block ts to date and prevPrice by (token/feed)and currentPrice by(token)
	repo.BlocksRepo.Load()
	// treasury funcs
	repo.LoadLastTreasuryTs(repo.db)
	repo.TreasuryRepo.LoadTreasurySnapshot(repo.db)
	// for direct token transfer
	repo.loadAccountLastSession()
	// credit_sessions
	repo.LoadCreditSessions(repo.db, lastDebtSync)
	repo.loadTicker()

	repo.initChecks()
}

func (repo *Repository) AddAccountOperation(accountOperation *schemas.AccountOperation) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if accountOperation.SessionId == "" {
		panic(utils.ToJson(accountOperation))
	}
	repo.SetAndGetBlock(accountOperation.BlockNumber).AddAccountOperation(accountOperation)
}

type LastSyncAndType struct {
	Type     string `gorm:"column:type"`
	LastSync int64  `gorm:"column:last_sync"`
	Address  string `gorm:"column:address"`
}

func (obj *LastSyncAndType) String() string {
	return fmt.Sprintf("%s(%s):%d", obj.Type, obj.Address, obj.LastSync)

}
func (repo *Repository) initChecks() {
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

func (repo *Repository) AfterSync(syncTill int64) {
	// for direct token transfer
	for _, txs := range repo.accountManager.GetNoSessionTxs() {
		for _, tx := range txs {
			repo.RecentMsgf(log.RiskHeader{
				BlockNumber: tx.BlockNum,
				EventCode:   "AMQP",
			}, "No session account token transfer: %v", tx)
			repo.SetAndGetBlock(tx.BlockNum).AddNoSessionTx(tx)
		}
	}
	// for direct token transfer
	repo.accountManager.Clear()
	// chainlink and uniswap prices
}

func (repo *Repository) ChainlinkPriceUpdatedAt(token string, blockNums []int64) {
	repo.GetAggregatedFeed().ChainlinkPriceUpdatedAt(token, blockNums)
}

func (repo *Repository) GetFeedToTicker(feed string) common.Address {
	return repo.feedToTicker[feed]
}
func (repo *Repository) AddFeedToTicker(feed string, ticker common.Address) {
	repo.feedToTicker[feed] = ticker
}
func (repo *Repository) saveTicker(tx *gorm.DB) {
	data := []ticker{}
	for feed, token := range repo.feedToTicker {
		data = append(data, ticker{
			Feed:   feed,
			Ticker: token.Hex(),
		})
	}
	err := tx.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(data, 10).Error
	log.CheckFatal(err)
	//
}

type ticker struct {
	Feed   string `gorm:"column:feed"`
	Ticker string `gorm:"column:ticker;primaryKey"`
}

func (repo *Repository) loadTicker() {
	data := []ticker{}
	err := repo.db.Raw(`select * from tickers`).Find(&data).Error
	log.CheckFatal(err)
	for _, entry := range data {
		repo.feedToTicker[entry.Feed] = common.HexToAddress(entry.Ticker)
	}
}
