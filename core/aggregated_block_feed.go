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
	BlockNum int64
}
