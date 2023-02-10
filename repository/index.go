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
	"gorm.io/gorm"
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
	db             *gorm.DB
	client         core.ClientI
	config         *config.Config
	accountManager *ds.DirectTransferManager
	relations      []*schemas.UniPriceAndChainlink
}

func GetRepository(db *gorm.DB, client core.ClientI, cfg *config.Config, extras *handlers.ExtrasRepo) *Repository {
	blocksRepo := handlers.NewBlocksRepo(db, client, cfg)
	tokensRepo := handlers.NewTokensRepo(client)
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
	}
	repo.SyncAdaptersRepo = handlers.NewSyncAdaptersRepo(client, repo, cfg, extras)
	repo.TokenOracleRepo = handlers.NewTokenOracleRepo(repo.SyncAdaptersRepo, blocksRepo, repo, client)
	repo.TreasuryRepo = treasury.NewTreasuryRepo(tokensRepo, blocksRepo, repo.SyncAdaptersRepo, client)
	return repo
}

func NewRepository(db *gorm.DB, client core.ClientI, config *config.Config, ep *handlers.ExtrasRepo) ds.RepositoryI {
	r := GetRepository(db, client, config, ep)
	return r
}

func (repo *Repository) Init() {
	// lastdebtsync is required to load credit session which are active or closed after lastdebtsync block number
	lastDebtSync := repo.LoadLastDebtSync()
	// token should be loaded before syncAdapters as credit manager adapter uses underlying token details
	repo.TokensRepo.LoadTokens(repo.db)
	// syncadapter state for cm and pool is set after loading of pool/credit manager table data from db
	repo.SyncAdaptersRepo.LoadSyncAdapters(repo.db)
	// load poolLMrewards
	repo.loadLMRewardDetails()
	repo.loadChainlinkPrevState()
	//
	repo.loadUniswapPools()
	// for disabling previous token oracle if new oracle is set
	repo.LoadCurrentTokenOracle(repo.db)
	// load state for sync_adapters
	repo.loadPool()
	repo.LoadPoolUniqueUsers(repo.db)
	// load credit manager
	repo.loadCreditManagers()
	repo.loadGearBalances()
	// required for disabling allowed tokens
	repo.LoadAllowedTokensState(repo.db)
	// fastcheck and new parameters
	repo.ParamsRepo.LoadAllParams(repo.db)
	// treasury funcs
	repo.BlocksRepo.LoadBlockDatePair()
	repo.LoadLastTreasuryTs(repo.db)
	repo.TreasuryRepo.LoadTreasurySnapshot(repo.db)
	// for direct token transfer
	repo.loadAccountLastSession()
	// credit_sessions
	repo.LoadCreditSessions(repo.db, lastDebtSync)

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
	repo.AggregatedFeed.Clear()
}
