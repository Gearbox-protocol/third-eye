package utils

import (
	"context"
	"fmt"
	"github.com/Gearbox-protocol/third-eye/artifacts/creditManager"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"strings"
	"time"
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

func GetFloat64Decimal(num *big.Int, decimals int8) float64 {
	floatBorrowedAmount, _ := GetFloat64(num, decimals).Float64()
	return floatBorrowedAmount
}

func GetInt64Decimal(num *big.Int, decimals int8) *big.Int {
	if decimals > 0 {
		return new(big.Int).Quo(
			num,
			GetExpInt(decimals))
	} else {
		return new(big.Int).Mul(
			num,
			GetExpInt(decimals))
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
			GetExpFloat(decimals))
	}
}

func AlmostSameBigInt(a, b *big.Int, decimals int8, precision int8) bool {
	// diff should be less than 100
	return new(big.Int).Sub(a, b).CmpAbs(GetExpInt(decimals-precision)) <= 0
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

func GetCreditManagerEventIds() []string {
	var ids []string
	if a, err := abi.JSON(strings.NewReader(creditManager.CreditManagerABI)); err == nil {
		for _, event := range a.Events {
			// fmt.Println(event.RawName, event.ID.Hex())
			// if event.RawName != "ExecuteOrder" {
			ids = append(ids, event.ID.Hex())
			// }
		}
	}
	return ids
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func GetTimeoutCtx(sec int) (context.Context, context.CancelFunc) {
	//https://blog.golang.org/context
	timeout, err := time.ParseDuration(fmt.Sprintf("%ds", sec))
	if err != nil {
		log.Error(err)
	}
	ctx, cancel := context.WithTimeout(context.TODO(), timeout*time.Second)
	return ctx, cancel
}
func GetTimeoutOpts(blockNum int64) (*bind.CallOpts, context.CancelFunc) {
	ctx, cancel := GetTimeoutCtx(20)
	return &bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
		Context:     ctx,
	}, cancel
}

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
