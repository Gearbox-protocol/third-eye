package core

import (
	"math/big"
)

type LatestRounData struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}

type UniPoolPrices struct {
	ID                   int64   `gorm:"primaryKey;column:id;autoIncrement:true"`
	PriceV2              float64 `gorm:"column:uniswapv2_price"`
	TwapV3               float64 `gorm:"column:uniswapv3_price"`
	PriceV3              float64 `gorm:"column:uniswapv3_twap"`
	PriceV2Success       bool    `gorm:"-"`
	TwapV3Success        bool    `gorm:"-"`
	PriceV3Success       bool    `gorm:"-"`
	BlockNum             int64   `gorm:"column:block_num"`
	ChainlinkBlockNumber int64   `gorm:"column:chainlink_block_num"`
	Token                string  `gorm:"column:token"`
}

func (UniPoolPrices) TableName() string {
	return "uniswap_pool_prices"
}

type UniswapPools struct {
	V2       string
	V3       string
	Decimals int8
	LastSync int64
}
