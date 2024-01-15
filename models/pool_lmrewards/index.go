package pool_lmrewards

import "github.com/Gearbox-protocol/sdk-go/core"

type LMReward struct {
	User   string       `gorm:"primaryKey;column:user_address"`
	Pool   string       `gorm:"primaryKey;column:pool"`
	Reward *core.BigInt `gorm:"column:reward"`
}

func (LMReward) TableName() string {
	return "lm_rewards"
}
