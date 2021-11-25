package core

type CreditManager struct {
	Address           string `gorm:"primaryKey"`
	PoolAddress       string `gorm:"column:pool_address"`
	UnderlyingToken   string `gorm:"column:underlying_token"`
	IsWETH            bool   `gorm:"column:is_weth"`
	MaxLeverageFactor int    `gorm:"column:max_leverage"`
	Sessions          Hstore `gorm:"column:sessions"`
}
