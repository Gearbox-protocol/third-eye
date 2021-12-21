package repository

import (
	"context"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"math/big"
)

func (repo *Repository) LoadBlocks(from, to int64) {
	log.Infof("Loaded %d to %d blocks for debt", from, to)
	data := []*core.Block{}
	err := repo.db.Preload("CSS").Preload("PoolStats").
		Preload("AllowedTokens").Preload("PriceFeeds").
		Find(&data, "id > ? AND id <= ?", from, to).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, block := range data {
		repo.blocks[block.BlockNumber] = block
	}
}

func (repo *Repository) GetBlocks() map[int64]*core.Block {
	return repo.blocks
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
