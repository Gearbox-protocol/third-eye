package handlers

import (
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"gorm.io/gorm"
)

type PoolUsersRepo struct {
	// changed during syncing
	poolUniqueUsers map[string]map[string]bool
}

func NewPoolUsersRepo() *PoolUsersRepo {
	return &PoolUsersRepo{
		poolUniqueUsers: make(map[string]map[string]bool),
	}
}

func (repo *PoolUsersRepo) LoadPoolUniqueUsers(db *gorm.DB) {
	defer utils.Elapsed("loadPoolUniqueUsers")()
	query := "select distinct pool, user_address from pool_ledger WHERE event = 'AddLiquidity';"
	data := []*schemas.PoolLedger{}
	err := db.Raw(query).Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range data {
		repo.AddPoolUniqueUser(entry.Pool, entry.User)
	}
}

func (repo *PoolUsersRepo) AddPoolUniqueUser(pool, user string) {
	if repo.poolUniqueUsers[pool] == nil {
		repo.poolUniqueUsers[pool] = make(map[string]bool)
	}
	repo.poolUniqueUsers[pool][user] = true
}

func (repo *PoolUsersRepo) GetPoolUniqueUserLen(pool string) int {
	return len(repo.poolUniqueUsers[pool])
}
