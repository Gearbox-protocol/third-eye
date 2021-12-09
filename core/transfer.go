package core

import (
	"database/sql/driver"
	"fmt"
	"math/big"
)

type Transfers map[string]*big.Int

func (bal *Transfers) String() string {
	var str string
	first := true
	for addr, amt := range (map[string]*big.Int)(*bal) {
		if !first {
			str += ","
		}
		str += fmt.Sprintf("%s=>%s", addr, amt.String())
		first = false
	}
	return str
}

func (src *Transfers) Value() (driver.Value, error) {
	if src == nil {
		return nil, nil
	}
	return src.String(), nil
}

func (dst *Transfers) Scan(value interface{}) error {
	hstore := NewHstore()
	if err := hstore.Scan(value); err != nil {
		return err
	}
	var transfers map[string]*big.Int
	for key, val := range hstore.GetMap() {
		if intval, ok := new(big.Int).SetString(val, 10); !ok {
			return fmt.Errorf("")
		} else {
			transfers[key] = intval
		}
	}
	*dst = (Transfers)(transfers)
	return nil
}
