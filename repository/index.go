package repository

import (
	"sync"
	"time"

	"fmt"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/repository/handlers"
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
	// mutex
	mu *sync.Mutex
	// object fx objects
	db            *gorm.DB
	client        core.ClientI
	config        *config.Config
	executeParser ds.ExecuteParserI
	// treasury
	treasurySnapshot *schemas.TreasurySnapshot
	lastTreasureTime time.Time
	accountManager   *ds.AccountTokenManager
	relations        []*schemas.UniPriceAndChainlink
}

func GetRepository(db *gorm.DB, client core.ClientI, cfg *config.Config, ep ds.ExecuteParserI) *Repository {
	blocksRepo := handlers.NewBlocksRepo(db, client)
	tokensRepo := handlers.NewTokensRepo(client)
	repo := &Repository{
		SessionRepo:      handlers.NewSessionRepo(),
		AllowedTokenRepo: handlers.NewAllowedTokenRepo(blocksRepo, tokensRepo),
		ParamsRepo:       handlers.NewParamsRepo(blocksRepo),
		PoolUsersRepo:    handlers.NewPoolUsersRepo(),
		TokensRepo:       tokensRepo,
		BlocksRepo:       blocksRepo,
		ExtrasRepo:       handlers.NewExtraRepo(client),
		mu:               &sync.Mutex{},
		db:               db,
		client:           client,
		config:           cfg,
		executeParser:    ep,
		accountManager:   ds.NewAccountTokenManager(),
	}
	repo.SyncAdaptersRepo = handlers.NewSyncAdaptersRepo(client, repo, cfg, repo.ExtrasRepo)
	repo.TokenOracleRepo = handlers.NewTokenOracleRepo(repo.SyncAdaptersRepo, blocksRepo, repo, client)
	return repo
}

func NewRepository(db *gorm.DB, client core.ClientI, config *config.Config, ep ds.ExecuteParserI) ds.RepositoryI {
	r := GetRepository(db, client, config, ep)
	r.init()
	return r
}

func (repo *Repository) GetExecuteParser() ds.ExecuteParserI {
	return repo.executeParser
}

func (repo *Repository) init() {
	lastDebtSync := repo.LoadLastDebtSync()
	// token should be loaded before syncAdapters as credit manager adapter uses underlying token details
	repo.TokensRepo.LoadTokens(repo.db)
	// syncadapter state for cm and pool is set after loading of pool/credit manager table data from db
	repo.SyncAdaptersRepo.LoadSyncAdapters(repo.db)
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
	repo.loadLastTreasuryTs()
	repo.loadTreasurySnapshot()
	// for direct token transfer
	repo.loadAccountLastSession()
	// credit_sessions
	repo.LoadCreditSessions(repo.db, lastDebtSync)
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
	repo.SetAndGetBlock(obj.BlockNumber).AddTransferAccountAllowed(obj)
}
