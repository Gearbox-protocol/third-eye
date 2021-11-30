package repository

import (
	"sync"

	"github.com/Gearbox-protocol/gearscan/artifacts/dataCompressor"
	"github.com/Gearbox-protocol/gearscan/core"
	"github.com/Gearbox-protocol/gearscan/ethclient"
	"github.com/Gearbox-protocol/gearscan/log"
	"github.com/Gearbox-protocol/gearscan/services"
	"gorm.io/gorm"

	"context"
	"math/big"
)

type Repository struct {
	db             *gorm.DB
	syncAdapters   []core.SyncAdapterI
	mu             *sync.Mutex
	client         *ethclient.Client
	blocks         map[int64]*core.Block
	creditManagers map[string]*core.CreditManager
	executeParser  *services.ExecuteParser
	tokens         map[string]*core.Token
	allowedTokens  []*core.AllowedToken
	pools          map[string]*core.Pool
	dc             map[int64]*dataCompressor.DataCompressor
	dcBlockNum     []int64
	sessions       map[string]*core.CreditSession
	lastCSS        map[string]*core.CreditSessionSnapshot
}

func NewRepository(db *gorm.DB, client *ethclient.Client, ep *services.ExecuteParser) core.RepositoryI {
	r := &Repository{
		db:             db,
		mu:             &sync.Mutex{},
		client:         client,
		blocks:         make(map[int64]*core.Block),
		creditManagers: make(map[string]*core.CreditManager),
		executeParser:  ep,
		dc: make(map[int64]*dataCompressor.DataCompressor),
		tokens:         make(map[string]*core.Token),
		pools:          make(map[string]*core.Pool),
		sessions:       make(map[string]*core.CreditSession),
		lastCSS:        make(map[string]*core.CreditSessionSnapshot),
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
	repo.loadPool()
	repo.loadCreditManagers()
	repo.loadCreditSessions()
	repo.loadLastCSS()
	repo.loadSyncAdapters()
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
