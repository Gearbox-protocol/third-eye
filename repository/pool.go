package repository

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
)

func (repo *Repository) AddPool(pool *core.Pool) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if repo.pools[pool.Address] == nil {
		repo.pools[pool.Address] = pool
	}
}

func (repo *Repository) loadPool() {
	data := []*core.Pool{}
	err := repo.db.Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, cm := range data {
		repo.AddPool(cm)
	}
}
