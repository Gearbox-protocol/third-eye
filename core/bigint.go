/*
 * Gearbox monitoring
 * Copyright (c) 2021. Mikael Lazarev
 *
 */

package core

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
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
	return nil, nil
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
	return []byte(z.String()), nil
}

func (z *BigInt) UnmarshalJSON(b []byte) error {
	value, ok := new(big.Int).SetString(string(b), 10)
	if !ok {
		return errors.New("can unmarshap BigInt")
	}

	*z = *(*BigInt)(value)
	return nil

}
