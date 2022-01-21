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
		Preload("AllowedTokens").Preload("PriceFeeds").Preload("Params").
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

func (repo *Repository) setBlock(blockNum int64) {
	if repo.blocks[blockNum] == nil {
		b, err := repo.client.BlockByNumber(context.Background(), big.NewInt(blockNum))
		// if err != nil && err.Error() == "server returned empty uncle list but block header indicates uncles" {
		// 	repo.blocks[blockNum] = &core.Block{BlockNumber: blockNum}
		// 	return
		// }
		log.CheckFatal(err)
		repo.blocks[blockNum] = &core.Block{BlockNumber: blockNum, Timestamp: b.Time()}
		repo.addBlockDate(&core.BlockDate{BlockNum: blockNum, Timestamp: int64(b.Time())})
	}
}

func (repo *Repository) setAndGetBlock(blockNum int64) *core.Block {
	repo.setBlock(blockNum)
	return repo.blocks[blockNum]
}

func (repo *Repository) SetAndGetBlock(blockNum int64) *core.Block {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	return repo.setAndGetBlock(blockNum)
}

func (repo *Repository) SetBlock(blockNum int64) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.setBlock(blockNum)
}
