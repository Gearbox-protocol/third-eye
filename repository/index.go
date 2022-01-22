package repository

import (
	"sync"
	"time"

	"fmt"

	"github.com/Gearbox-protocol/third-eye/artifacts/creditFilter"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/utils"
	"gorm.io/gorm"
)

type Repository struct {
	// mutex
	mu *sync.Mutex
	// object fx objects
	WETHAddr              string
	USDCAddr              string
	GearTokenAddr         string
	db                    *gorm.DB
	client                *ethclient.Client
	config                *config.Config
	kit                   *core.AdapterKit
	executeParser         core.ExecuteParserI
	dcWrapper             *core.DataCompressorWrapper
	creditManagerToFilter map[string]*creditFilter.CreditFilter
	allowedTokens         map[string]map[string]*core.AllowedToken
	disabledTokens        []*core.AllowedToken
	// blocks/token
	blocks map[int64]*core.Block
	tokens map[string]*core.Token
	// changed during syncing
	sessions            map[string]*core.CreditSession
	poolUniqueUsers     map[string]map[string]bool
	tokensCurrentOracle map[string]*core.TokenOracle
	// for params diff calculation
	cmParams          map[string]*core.Parameters
	cmFastCheckParams map[string]*core.FastCheckParams
	// treasury
	treasurySnapshot *core.TreasurySnapshot
	lastTreasureTime time.Time
	BlockDatePairs   map[int64]*core.BlockDate
	dieselTokens     map[string]*core.UTokenAndPool
	accountManager   *core.AccountTokenManager
}

func NewRepository(db *gorm.DB, client *ethclient.Client, config *config.Config, ep core.ExecuteParserI) core.RepositoryI {
	r := &Repository{
		mu:                    &sync.Mutex{},
		db:                    db,
		client:                client,
		config:                config,
		blocks:                make(map[int64]*core.Block),
		executeParser:         ep,
		kit:                   core.NewAdapterKit(),
		tokens:                make(map[string]*core.Token),
		sessions:              make(map[string]*core.CreditSession),
		poolUniqueUsers:       make(map[string]map[string]bool),
		tokensCurrentOracle:   make(map[string]*core.TokenOracle),
		dcWrapper:             core.NewDataCompressorWrapper(client),
		creditManagerToFilter: make(map[string]*creditFilter.CreditFilter),
		allowedTokens:         make(map[string]map[string]*core.AllowedToken),
		cmParams:              make(map[string]*core.Parameters),
		cmFastCheckParams:     make(map[string]*core.FastCheckParams),
		BlockDatePairs:        make(map[int64]*core.BlockDate),
		dieselTokens:          make(map[string]*core.UTokenAndPool),
		accountManager:        core.NewAccountTokenManager(),
	}
	r.init()
	return r
}

func (repo *Repository) GetDCWrapper() *core.DataCompressorWrapper {
	return repo.dcWrapper
}

func (repo *Repository) GetExecuteParser() core.ExecuteParserI {
	return repo.executeParser
}

func (repo *Repository) GetKit() *core.AdapterKit {
	return repo.kit
}

func (repo *Repository) init() {
	lastDebtSync := repo.LoadLastDebtSync()
	// token should be loaded before syncAdapters as credit manager adapter uses underlying token details
	repo.loadToken()
	// syncadapter state for cm and pool is set after loading of pool/credit manager table data from db
	repo.loadSyncAdapters()
	// for disabling previous token oracle if new oracle is set
	repo.loadCurrentTokenOracle()
	// load state for sync_adpters
	repo.loadPool()
	repo.loadCreditManagers()
	repo.loadGearBalances()
	// required for disabling allowed tokens
	repo.loadAllowedTokensState()
	// fastcheck and new parameters
	repo.loadAllParams()
	// treasury funcs
	repo.loadBlockDatePair()
	repo.loadLastTreasuryTs()
	repo.loadTreasurySnapshot()
	// for direct token transfer
	repo.loadAccountLastSession()
	// credit_sessions
	repo.loadCreditSessions(lastDebtSync)
}

func (repo *Repository) AddAccountOperation(accountOperation *core.AccountOperation) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
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

// redundant
func (repo *Repository) AddEventBalance(eb core.EventBalance) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.setAndGetBlock(eb.BlockNumber).AddEventBalance(&eb)
}

func (repo *Repository) CallRankingProcedure() {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if err := repo.db.Raw("CALL rankings()").Error; err != nil {
		log.CheckFatal(err)
	}
	log.Info("Refreshed rankings by 7/20 days")
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
func (eng *Repository) InitChecks() {
	data := []*LastSyncAndType{}
	err := eng.db.Raw(`SELECT type, address,  last_sync AS last_calculated_at 
		FROM sync_adapters 
		WHERE type IN ('AccountManager', 'CreditManager','AccountFactory')`).Find(&data).Error
	log.CheckFatal(err)
	var accountManagerLastSync, accountFactoryLastSync int64
	var cmLastSync int64
	var str string
	for _, entry := range data {
		str += entry.String()
		switch entry.Type {
		case core.AccountFactory:
			accountFactoryLastSync = entry.LastSync
		case core.AccountManager:
			accountManagerLastSync = entry.LastSync
		case core.CreditManager:
			cmLastSync = utils.Min(entry.LastSync, cmLastSync)
		}
	}
	if accountFactoryLastSync != accountManagerLastSync ||
		cmLastSync < accountManagerLastSync {
		log.Fatal("Account manager/credit manager/AccountFactory are not synchronised: ", str)
	}
}
