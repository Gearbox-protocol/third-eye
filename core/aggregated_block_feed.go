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

type PoolPrices struct {
	PriceV2  float64
	TwapV3   float64
	PriceV3  float64
	PriceV2Success bool
	TwapV3Success  bool
	PriceV3Success bool
	BlockNum int64
}

type UniswapPools struct {
	V2       string
	V3       string
	Decimals int8
	LastSync int64
}
