/*
 * Gearbox monitoring
 * Copyright (c) 2021. Mikael Lazarev
 *
 */

package core

import (
	"database/sql/driver"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"strings"
)

func Topic(topic string) common.Hash {
	return crypto.Keccak256Hash([]byte(topic))
}

type BigInt big.Int

func (z *BigInt) Convert() *big.Int {
	return (*big.Int)(z)
}

func (z *BigInt) String() string {
	return z.Convert().String()
}

func (z *BigInt) Value() (driver.Value, error) {
	if z != nil {
		return (*big.Int)(z).String(), nil
	}
	return "0", nil
}
func NewBigInt(bi *BigInt) *BigInt {
	if bi == nil {
		return (*BigInt)(big.NewInt(0))
	}
	obj := new(big.Int).Mul(bi.Convert(), big.NewInt(1))
	return (*BigInt)(obj)
}

func (a *BigInt) Cmp(b *BigInt) int {
	return a.Convert().Cmp(b.Convert())
}

func DiffMoreThanFraction(oldValue, newValue *BigInt, diff *big.Float) bool {
	newFloat := new(big.Float).SetInt(newValue.Convert())
	oldFloat := new(big.Float).SetInt(oldValue.Convert())
	fractionalChange := new(big.Float).Quo(
		new(big.Float).Sub(newFloat, oldFloat),
		oldFloat)
	return new(big.Float).Abs(fractionalChange).Cmp(diff) >= 1
}
func ValueDifferSideOf10000(a, b *BigInt) bool {
	return (IntGreaterThanEqualTo(a, 10000) != IntGreaterThanEqualTo(b, 10000))
}
func IntGreaterThanEqualTo(value *BigInt, cmp int64) bool {
	return value.Convert().Cmp(big.NewInt(cmp)) >= 0
}

func AddCoreAndInt(a *BigInt, b *big.Int) *BigInt {
	if a != nil {
		return (*BigInt)(new(big.Int).Add(a.Convert(), b))
	}
	return NewBigInt((*BigInt)(b))
}
func SubCoreAndInt(a *BigInt, b *big.Int) *BigInt {
	if a != nil {
		return (*BigInt)(new(big.Int).Sub(a.Convert(), b))
	}
	return NewBigInt((*BigInt)(b))
}

func (z *BigInt) Scan(value interface{}) error {
	if value == nil {
		z = nil
	}
	switch t := value.(type) {
	case string:
		v, ok := new(big.Int).SetString(value.(string), 10)
		if !ok {
			return fmt.Errorf("Could not scan type %T into BigInt", t)
		}
		*z = *(*BigInt)(v)
	default:
		return fmt.Errorf("Could not scan type %T into BigInt", t)
	}

	return nil
}

func (z *BigInt) MarshalJSON() ([]byte, error) {
	return []byte("\"" + z.String() + "\""), nil
}

func (z *BigInt) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), "\"")
	value, ok := new(big.Int).SetString(str, 10)
	if !ok {
		return fmt.Errorf("can unmarshal BigInt")
	}

	*z = *(*BigInt)(value)
	return nil

}
