package pool_lmrewards

import "github.com/Gearbox-protocol/sdk-go/core"

// for both v2/v3
type LMReward struct {
	User   string       `gorm:"primaryKey;column:user_address"`
	Pool   string       `gorm:"primaryKey;column:pool"`
	Farm   string       `gorm:"-"`
	Reward *core.BigInt `gorm:"column:reward"`
}

func (LMReward) TableName() string {
	return "lm_rewards"
}
