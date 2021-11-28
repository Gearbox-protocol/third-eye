package repository

import (
	"sync"

	"github.com/Gearbox-protocol/gearscan/core"
	"github.com/Gearbox-protocol/gearscan/log"
	"github.com/Gearbox-protocol/gearscan/utils"
	"github.com/Gearbox-protocol/gearscan/ethclient"
	"github.com/Gearbox-protocol/gearscan/artifacts/dataCompressor"
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
	executeParser *utils.ExecuteParser
	tokens        map[string]*core.Token
	allowedTokens []*core.AllowedToken
	pools         map[string]*core.Pool
	dc            *dataCompressor.DataCompressor
	sessions            map[string]*core.CreditSession
	lastCSS        map[string]*core.CreditSessionSnapshot
}

func NewRepository(db *gorm.DB, client *ethclient.Client, ep *utils.ExecuteParser) core.RepositoryI {
	r := &Repository{
		db:             db,
		mu:             &sync.Mutex{},
		client:         client,
		blocks:         make(map[int64]*core.Block),
		creditManagers: make(map[string]*core.CreditManager),
		executeParser: ep,
		tokens: make(map[string]*core.Token),
		pools: make(map[string]*core.Pool),
		sessions: make(map[string]*core.CreditSession),
		lastCSS: make(map[string]*core.CreditSessionSnapshot),
	}
	r.init()
	return r
}

func (repo *Repository) GetExecuteParser() *utils.ExecuteParser {
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
