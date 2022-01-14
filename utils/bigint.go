package utils

import (
	"fmt"
	"math/big"
)

// maths
func GetExpFloat(decimals int8) *big.Float {
	if decimals < 0 {
		panic(fmt.Sprintf("GetExpFloat received pow:%d", decimals))
	}
	bigIntDecimal := new(big.Int).Exp(big.NewInt(10), new(big.Int).SetInt64(int64(decimals)), big.NewInt(0))
	return new(big.Float).SetInt(bigIntDecimal)
}

func GetExpInt(decimals int8) *big.Int {
	if decimals < 0 {
		panic(fmt.Sprintf("GetInt received pow:%d", decimals))
	}
	return new(big.Int).Exp(big.NewInt(10), new(big.Int).SetInt64(int64(decimals)), big.NewInt(0))
}

func IntToFloat(amt *big.Int) *big.Float {
	return new(big.Float).SetInt(amt)
}

func PercentMul(a, b *big.Int) *big.Int {
	ans := new(big.Int).Mul(a, b)
	return new(big.Int).Quo(ans, big.NewInt(10000))
}

func GetFloat64Decimal(num *big.Int, decimals int8) float64 {
	floatBorrowedAmount, _ := GetFloat64(num, decimals).Float64()
	return floatBorrowedAmount
}

func GetInt64(num *big.Int, decimals int8) *big.Int {
	if decimals > 0 {
		return new(big.Int).Quo(
			num,
			GetExpInt(decimals))
	} else {
		return new(big.Int).Mul(
			num,
			GetExpInt(-decimals))
	}
}

func StringToInt(v string) *big.Int {
	value, ok := new(big.Int).SetString(v, 10)
	if !ok {
		panic("Parsing string to big.int failed")
	}
	return value
}

func GetFloat64(num *big.Int, decimals int8) *big.Float {
	if decimals > 0 {
		return new(big.Float).Quo(
			IntToFloat(num),
			GetExpFloat(decimals))
	} else {
		return new(big.Float).Mul(
			IntToFloat(num),
			GetExpFloat(-decimals))
	}
}

func AlmostSameBigInt(a, b *big.Int, noOFZeroIndiff int8) bool {
	// diff should be less than 100
	return new(big.Int).Sub(a, b).CmpAbs(GetExpInt(noOFZeroIndiff)) <= 0
}

func Min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// others

func GetPrecision(symbol string) int8 {
	switch symbol {
	case "USDC":
		return 0
	case "DAI":
		return 0
	case "WBTC":
		return 5
	case "LINK":
		return 2
	case "SNX":
		return 3
	}
	return 0
}

func absInt64(a int64) int64 {
	if a > 0 {
		return a
	}
	return -1 * a
}

func IntDiffMoreThanFraction(oldValue, newValue, diff int64) bool {
	return absInt64((newValue-oldValue)/oldValue) > diff
}
