package repository

import (
	"github.com/Gearbox-protocol/third-eye/artifacts/creditFilter"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"gorm.io/gorm"
	"sync"
)

type Repository struct {
	// mutex
	mu *sync.Mutex
	// object fx objects
	WETHAddr              string
	db                    *gorm.DB
	client                *ethclient.Client
	config                *config.Config
	kit                   *core.AdapterKit
	executeParser         core.ExecuteParserI
	dcWrapper             *core.DataCompressorWrapper
	creditManagerToFilter map[string]*creditFilter.CreditFilter
	// blocks/token
	blocks map[int64]*core.Block
	tokens map[string]*core.Token
	// changed during syncing
	sessions            map[string]*core.CreditSession
	poolUniqueUsers     map[string]map[string]bool
	tokensCurrentOracle map[string]*core.TokenOracle
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
	repo.loadCurrentTokenOracle()
	repo.loadPool()
	repo.loadCreditManagers()
	repo.loadCreditSessions(lastDebtSync)
}

func (repo *Repository) AddAccountOperation(accountOperation *core.AccountOperation) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.blocks[accountOperation.BlockNumber].AddAccountOperation(accountOperation)
}

func (repo *Repository) SetWETHAddr(addr string) {
	repo.WETHAddr = addr
}

func (repo *Repository) GetWETHAddr() string {
	return repo.WETHAddr
}

// redundant
func (repo *Repository) AddEventBalance(eb core.EventBalance) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.blocks[eb.BlockNumber].AddEventBalance(&eb)
}
