package repository

import (
	"github.com/Gearbox-protocol/gearscan/core"
	"github.com/Gearbox-protocol/gearscan/ethclient"
	"github.com/Gearbox-protocol/gearscan/log"
	"gorm.io/gorm"
	"sync"

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
}

func NewRepository(db *gorm.DB, client *ethclient.Client) core.RepositoryI {
	r := &Repository{
		db:             db,
		mu:             &sync.Mutex{},
		client:         client,
		blocks:         make(map[int64]*core.Block),
		creditManagers: make(map[string]*core.CreditManager),
	}
	r.init()
	return r
}

func (repo *Repository) init() {
	repo.loadSyncAdapters()
}

func (repo *Repository) GetSyncAdapters() []core.SyncAdapterI {
	return repo.syncAdapters
}

func (repo *Repository) AddSyncAdapter(adapterI core.SyncAdapterI) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.syncAdapters = append(repo.syncAdapters, adapterI)
}

func (repo *Repository) AddCreditManager(cm *core.CreditManager) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if repo.creditManagers[cm.Address] != nil {
		log.Fatal("credit manager already set")
	}
	repo.creditManagers[cm.Address] = cm
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
