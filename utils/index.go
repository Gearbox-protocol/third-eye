package utils

import (
	"math/big"
	"fmt"
	"strings"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/Gearbox-protocol/gearscan/artifacts/creditManager"
)

// maths 
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