package repository

import (
	"gorm.io/gorm"
	"sync"
	"github.com/Gearbox-protocol/gearscan/core"
	"github.com/Gearbox-protocol/gearscan/ethclient"
	"github.com/Gearbox-protocol/gearscan/log"

	"math/big"
	"context"
)

type Repository struct {
	db *gorm.DB
	syncAdapters []core.SyncAdapterI
	mu    *sync.Mutex
	client *ethclient.Client
	blocks map[int64]*core.Block
}


func NewRepository(db *gorm.DB, client *ethclient.Client) core.RepositoryI {
	r := &Repository{
		db: db,
		mu:    &sync.Mutex{},
		client: client,
		blocks: make(map[int64]*core.Block),
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

func (repo *Repository) loadSyncAdapters()  {
	data := []*core.SyncAdapter{}
	err := repo.db.Find(&data, "disabled = ?", false).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, adapter := range data {
		repo.AddSyncAdapter(prepareSyncAdapter(adapter, repo, repo.client))
	}
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