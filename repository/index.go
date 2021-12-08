package repository

import (
	"sync"

	"github.com/Gearbox-protocol/third-eye/artifacts/dataCompressor"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/log"
	"gorm.io/gorm"

	"context"
	"math/big"
)

type Repository struct {
	// mutex
	mu *sync.Mutex
	// object fx objects
	db            *gorm.DB
	kit           *core.AdapterKit
	client        *ethclient.Client
	executeParser core.ExecuteParserI
	dc            map[int64]*dataCompressor.DataCompressor
	// blocks/token
	blocks map[int64]*core.Block
	tokens map[string]*core.Token
	// blockNumbers of dc in asc order
	dcBlockNum []int64
	// changed during syncing
	sessions            map[string]*core.CreditSession
	poolUniqueUsers     map[string]map[string]bool
	tokensCurrentOracle map[string]*core.TokenOracle
	// modified after sync loop
	lastCSS        map[string]*core.CreditSessionSnapshot
	tokenLastPrice map[string]*core.PriceFeed
	//// token -> credit_manager -> liquidity threshold
	allowedTokensThreshold map[string]map[string]*core.BigInt
	poolLastInterestData   map[string]*core.PoolInterestData
}

func NewRepository(db *gorm.DB, client *ethclient.Client, ep core.ExecuteParserI) core.RepositoryI {
	r := &Repository{
		db:                     db,
		mu:                     &sync.Mutex{},
		client:                 client,
		blocks:                 make(map[int64]*core.Block),
		executeParser:          ep,
		kit:                    core.NewAdapterKit(),
		dc:                     make(map[int64]*dataCompressor.DataCompressor),
		tokens:                 make(map[string]*core.Token),
		sessions:               make(map[string]*core.CreditSession),
		lastCSS:                make(map[string]*core.CreditSessionSnapshot),
		poolUniqueUsers:        make(map[string]map[string]bool),
		tokensCurrentOracle:    make(map[string]*core.TokenOracle),
		tokenLastPrice:         make(map[string]*core.PriceFeed),
		allowedTokensThreshold: make(map[string]map[string]*core.BigInt),
		poolLastInterestData:   make(map[string]*core.PoolInterestData),
	}
	r.init()
	return r
}

func (repo *Repository) GetExecuteParser() core.ExecuteParserI {
	return repo.executeParser
}

func (repo *Repository) init() {
	// token should be loaded before syncAdapters as credit manager adapter uses underlying token details
	repo.loadToken()
	// syncadapter state for cm and pool is set after loading of pool/credit manager table data from db
	repo.loadSyncAdapters()
	repo.loadPool()
	repo.loadCreditManagers()
	repo.loadCreditSessions()
	repo.loadLastCSS()
	repo.loadCurrentTokenOracle()
	repo.loadTokenLastPrice()
	repo.loadAllowedTokenThreshold()
}

func (repo *Repository) AddAccountOperation(accountOperation *core.AccountOperation) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.blocks[accountOperation.BlockNumber].AddAccountOperation(accountOperation)
}

func (repo *Repository) SetBlock(blockNum int64) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if repo.blocks[blockNum] == nil {
		b, err := repo.client.BlockByNumber(context.Background(), big.NewInt(blockNum))
		if err != nil {
			log.Fatal(err)
		}
		repo.blocks[blockNum] = &core.Block{BlockNumber: blockNum, Timestamp: b.Time()}
	}
}

func (repo *Repository) AddCreditManagerStats(cms *core.CreditManagerStat) {
	repo.blocks[cms.BlockNum].AddCreditManagerStats(cms)
}

func (repo *Repository) GetKit() *core.AdapterKit {
	return repo.kit
}

func (repo *Repository) AddEventBalance(eb core.EventBalance) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.blocks[eb.BlockNumber].AddEventBalance(&eb)
}
