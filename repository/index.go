package repository

import (
	"sync"

	"github.com/Gearbox-protocol/third-eye/artifacts/dataCompressor"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/services"
	"gorm.io/gorm"

	"context"
	"math/big"
)

type Repository struct {
	db             *gorm.DB
	syncAdapters   map[string]core.SyncAdapterI
	mu             *sync.Mutex
	client         *ethclient.Client
	blocks         map[int64]*core.Block
	executeParser  *services.ExecuteParser
	tokens         map[string]*core.Token
	allowedTokens  []*core.AllowedToken
	pools          map[string]*core.Pool
	dc             map[int64]*dataCompressor.DataCompressor
	dcBlockNum     []int64
	sessions       map[string]*core.CreditSession
	lastCSS        map[string]*core.CreditSessionSnapshot
	poolUniqueUsers map[string]map[string]bool
}

func NewRepository(db *gorm.DB, client *ethclient.Client, ep *services.ExecuteParser) core.RepositoryI {
	r := &Repository{
		db:             db,
		mu:             &sync.Mutex{},
		client:         client,
		blocks:         make(map[int64]*core.Block),
		executeParser:  ep,
		syncAdapters: make(map[string]core.SyncAdapterI),
		dc:             make(map[int64]*dataCompressor.DataCompressor),
		tokens:         make(map[string]*core.Token),
		pools:          make(map[string]*core.Pool),
		sessions:       make(map[string]*core.CreditSession),
		lastCSS:        make(map[string]*core.CreditSessionSnapshot),
		poolUniqueUsers: make(map[string]map[string]bool),
	}
	r.init()
	return r
}

func (repo *Repository) GetExecuteParser() *services.ExecuteParser {
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