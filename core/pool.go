package core

type Pool struct {
	Address         string `gorm:"primaryKey"`
	UnderlyingToken string `gorm:"column:underlying_token"`
	DieselToken     string `gorm:"column:diesel_token"`
	IsWETH          bool   `gorm:"column:is_weth"`
}
