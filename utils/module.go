package utils

import (
	"go.uber.org/fx"
	"math/big"
	"fmt"
)

var Module = fx.Option(
	fx.Provide(NewExecuteParser))


func GetExpFloat(decimals int64) *big.Float {
	if decimals < 0 {
		panic(fmt.Sprintf("GetExpFloat received pow:%d", decimals))
	}
	bigIntDecimal := new(big.Int).Exp(big.NewInt(10), new(big.Int).SetInt64(decimals), big.NewInt(0))
	return new(big.Float).SetInt(bigIntDecimal)
}
func IntToFloat(amt *big.Int) *big.Float {
	return new(big.Float).SetInt(amt)
}

func GetFloat64Decimal(num *big.Int, decimals uint8) float64 {
	floatBorrowedAmount, _ := new(big.Float).Quo(
		IntToFloat(num),
		GetExpFloat(int64(decimals)),
	).Float64()
	return floatBorrowedAmount
}