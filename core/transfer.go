package core

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/Gearbox-protocol/third-eye/utils"
	"math/big"
	"strings"
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
	fmt.Println("aanil != nil")
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

func (z *Transfers) UnmarshalJSON(b []byte) error {
	obj := Json{}
	d := json.NewDecoder(bytes.NewReader(b))
	d.UseNumber()
	if err := d.Decode(&obj); err != nil {
		fmt.Println("error:", err)
	}
	transfers := map[string]*big.Int{}
	for token, amount := range obj {
		value, ok := amount.(string)
		if !ok {
			return fmt.Errorf("can unmarshal BigInt")
		}
		bigAmount := strings.Trim(value, "\"")
		transfers[token] = utils.StringToInt(bigAmount)
	}
	*z = transfers
	return nil
}
