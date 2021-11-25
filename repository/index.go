package repository

import (
	"sync"

	"github.com/Gearbox-protocol/gearscan/core"
	"github.com/Gearbox-protocol/gearscan/ethclient"
	"github.com/Gearbox-protocol/gearscan/log"
	"github.com/Gearbox-protocol/gearscan/utils"
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
}

func NewRepository(db *gorm.DB, client *ethclient.Client, ep *utils.ExecuteParser) core.RepositoryI {
	r := &Repository{
		db:             db,
		mu:             &sync.Mutex{},
		client:         client,
		blocks:         make(map[int64]*core.Block),
		creditManagers: make(map[string]*core.CreditManager),
		executeParser: ep,
	}
	r.init()
	return r
}

func (repo *Repository) GetExecuteParser() *utils.ExecuteParser {
	return repo.executeParser
}

func (repo *Repository) init() {
	repo.loadSyncAdapters()
	repo.loadCreditManagers()
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
