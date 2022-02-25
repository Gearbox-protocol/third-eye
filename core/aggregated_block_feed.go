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
	ID             int64   `gorm:"primaryKey;column:id;autoIncrement:true"`
	PriceV2        float64 `gorm:"column:uniswapv2_price"`
	TwapV3         float64 `gorm:"column:uniswapv3_price"`
	PriceV3        float64 `gorm:"column:uniswapv3_twap"`
	PriceV2Success bool    `gorm:"-"`
	TwapV3Success  bool    `gorm:"-"`
	PriceV3Success bool    `gorm:"-"`
	BlockNum       int64   `gorm:"column:block_num"`
	Token          string  `gorm:"column:token"`
}

func (UniPoolPrices) TableName() string {
	return "uniswap_pool_prices"
}

type UniPriceAndChainlink struct {
	ChainlinkBlockNumber int64  `gorm:"column:chainlink_block_num"`
	Token                string `gorm:"column:token"`
	UniBlockNum          int64  `gorm:"column:block_num"`
	Feed                 string `gorm:"column:feed"`
}

func (UniPriceAndChainlink) TableName() string {
	return "uniswap_chainlink_relations"
}

type UniswapPools struct {
	V2      string `gorm:"column:pool_v2"`
	V3      string `gorm:"column:pool_v3"`
	Token   string `gorm:"column:token;primaryKey"`
	Updated bool   `gorm:"-"`
}

type TokenSyncDetails struct {
	Decimals int8
	LastSync int64
}

// sort event balances by block number/log id
type SortedUniPoolPrices []*UniPoolPrices

func (ts SortedUniPoolPrices) Len() int {
	return len(ts)
}
func (ts SortedUniPoolPrices) Swap(i, j int) {
	ts[i], ts[j] = ts[j], ts[i]
}

// sort in increasing order by blockNumber,index
func (ts SortedUniPoolPrices) Less(i, j int) bool {
	return ts[i].BlockNum < ts[j].BlockNum
}
