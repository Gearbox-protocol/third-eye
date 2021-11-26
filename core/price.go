

package core

import (
	"math/big"
)

type TokenOracle struct {
	BlockNumber int64   `gorm:"column:block_num"`
	Token string   `gorm:"column:token"`
	Oracle string   `gorm:"column:oracle"`
}

func (TokenOracle) TableName() string {
	return "token_oracle"
}

type PriceFeed struct {
	BlockNumber int64   `gorm:"column:block_num"`
	Token string   `gorm:"column:token"`
	Feed string `gorm:"column:feed"`
	RoundId *big.Int `gorm:"column:round_id"`
	PriceETHBI *big.Int `gorm:"column:price_eth_bi"`
	PriceETH float64   `gorm:"column:price_eth"`
}

func (PriceFeed) TableName() string {
	return "price_feeds"
}

