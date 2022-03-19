package core

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/utils"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
	"testing"
)

type Json map[string]interface{}

func (j *Json) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (z *Json) Scan(value interface{}) error {
	out := map[string]interface{}{}
	switch t := value.(type) {
	case string:
		err := json.Unmarshal([]byte(value.(string)), &out)
		*z = Json(out)
		return err
	case []byte:
		err := json.Unmarshal(value.([]byte), &out)
		*z = Json(out)
		return err
	default:
		return fmt.Errorf("Could not scan type %T", t)
	}
}

type BalanceType struct {
	BI *BigInt `json:"BI"`
	F  float64 `json:"F"`
	// linked get fetched for data compressor is if this token is allowed on credit manager
	// we change it to the mask bit
	Linked bool `json:"linked"` // there has been a credit manager event /swap/addcollateral/borrow for this token
}

type JsonBalance map[string]*BalanceType

func (j *JsonBalance) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (z *JsonBalance) Scan(value interface{}) error {
	out := JsonBalance{}
	switch t := value.(type) {
	case []byte:
		err := json.Unmarshal(value.([]byte), &out)
		*z = out
		return err
	default:
		return fmt.Errorf("Could not scan type %T", t)
	}
}

func (j *JsonBalance) Copy() *JsonBalance {
	var newJB = make(JsonBalance)
	for k, v := range (map[string]*BalanceType)(*j) {
		newJB[k] = &BalanceType{
			BI:     NewBigInt(v.BI),
			F:      v.F,
			Linked: v.Linked,
		}
	}
	return &newJB
}

func (j *JsonBalance) ValueInUnderlying(underlyingToken string, uDecimals int8, prices JsonFloatMap) *big.Int {
	var total float64
	priceOfUnderlying := prices[underlyingToken]
	for token, bal := range *j {
		tokenPrice := prices[token]
		value := (bal.F * tokenPrice) / priceOfUnderlying
		total += value
	}
	valueInFloat := new(big.Float).Mul(big.NewFloat(total), utils.GetExpFloat(uDecimals))
	remainingFunds, _ := valueInFloat.Int(nil)
	return remainingFunds
}

type JsonBigIntMap map[string]*BigInt

func (j *JsonBigIntMap) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (z *JsonBigIntMap) Scan(value interface{}) error {
	out := JsonBigIntMap{}
	switch t := value.(type) {
	case string:
		err := json.Unmarshal([]byte(value.(string)), &out)
		*z = out
		return err
	case []byte:
		err := json.Unmarshal(value.([]byte), &out)
		*z = out
		return err
	default:
		return fmt.Errorf("Could not scan type %T", t)
	}
}

type JsonFloatMap map[string]float64

func (j *JsonFloatMap) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (z *JsonFloatMap) Scan(value interface{}) error {
	out := JsonFloatMap{}
	switch t := value.(type) {
	case string:
		err := json.Unmarshal([]byte(value.(string)), &out)
		*z = out
		return err
	case []byte:
		err := json.Unmarshal(value.([]byte), &out)
		*z = out
		return err
	default:
		return fmt.Errorf("Could not scan type %T", t)
	}
}

func (addrs AddressMap) checkIfAddress(v string) string {
	if strings.HasPrefix(v, "#") {
		key := strings.Trim(v, "#")
		if addrs[key] == "" {
			addrs[key] = utils.RandomAddr()
		}
		return addrs[key]
	} else if strings.HasPrefix(v, "!#") {
		key := strings.Trim(v, "!#")
		if addrs[key] == "" {
			addrs[key] = utils.RandomHash()
		}
		return addrs[key]
	} else {
		return v
	}
}

type AddressMap map[string]string

func (addrs AddressMap) checkInterface(data interface{}, t *testing.T) interface{} {
	switch data.(type) {
	case string:
		value, ok := data.(string)
		if !ok {
			t.Error("string parsing failed")
		}
		return addrs.checkIfAddress(value)
	case []interface{}:
		value, ok := data.([]interface{})
		if !ok {
			t.Error("[]interface{} parsing failed")
		}
		for i, entry := range value {
			value[i] = addrs.checkInterface(entry, t)
		}
		return value
	case map[string]interface{}:
		value, ok := data.(map[string]interface{})
		if !ok {
			t.Error("map[string]interface{} parsing failed")
		}
		for key, entry := range value {
			newKey := addrs.checkIfAddress(key)
			if newKey != key {
				delete(value, key)
				key = newKey
			}
			value[key] = addrs.checkInterface(entry, t)
		}
		return value
	default:
		return data
	}
}

func (z *Json) ParseAddress(t *testing.T, addrs AddressMap) {
	for key, data := range *z {
		newKey := addrs.checkIfAddress(key)
		if newKey != key {
			delete(*z, key)
			key = newKey
		}
		(*z)[key] = addrs.checkInterface(data, t)
	}
}

func (z *Json) UnmarshalJSON(b []byte) (err error) {
	tmpMap := map[string]interface{}{}
	err = json.Unmarshal(b, &tmpMap)
	if err == nil {
		(*z) = tmpMap
	}
	return
}

func (addrs AddressMap) ReplaceWithVariable(key string, data interface{}) interface{} {
	if key == "sessionId" {
		value, ok := data.(string)
		if !ok {
				log.Error("string parsing failed")
		}
		splits := strings.Split(value, "_")
		splits[0] = addrs[splits[0]]
		return strings.Join(splits, "_")
	}
	switch data.(type) {
	case string:
		value, ok := data.(string)
		if !ok {
			log.Error("string parsing failed")
		}
		answer := addrs[value]
		if answer != "" {
			return answer
		}
	case []interface{}:
		value, ok := data.([]interface{})
		if !ok {
			log.Error("[]interface{} parsing failed")
		}
		for i, entry := range value {
			value[i] = addrs.ReplaceWithVariable(key, entry)
		}
		return value
	case map[string]interface{}:
		value, ok := data.(map[string]interface{})
		if !ok {
			log.Error("map[string]interface{} parsing failed")
		}
		obj := Json(value)
		obj.ReplaceWithVariable(addrs)
		return obj
	default:
		return data
	}
	return data
}

func (z *Json) ReplaceWithVariable(addrToVariable AddressMap) {
	for key, value := range *z {
		if value == nil {
			delete(*z, key)
			continue
		}
		if addrToVariable[key] != "" {
			delete(*z, key)
			variable := addrToVariable[key]
			(*z)[variable] = addrToVariable.ReplaceWithVariable(key, value)
		} else {
			(*z)[key] = addrToVariable.ReplaceWithVariable(key, value)
		}
	}
}

func (z *Json) CheckSumAddress() {
	for key, value := range *z {
		(*z)[key] = fixAddress(value)
	}
}

func fixAddress(data interface{}) interface{} {
	switch data.(type) {
	case common.Address:
		value, ok := data.(common.Address)
		if !ok {
			log.Error("common.Address parsing failed")
		}
		return value.Hex()
	case string:
		value, ok := data.(string)
		if !ok {
			log.Error("string parsing failed")
		}
		if len(value) == 42 && value[:2] == "0x" {
			log.Info(value)
			return common.HexToAddress(value).Hex()
		}
	case []interface{}:
		value, ok := data.([]interface{})
		if !ok {
			log.Error("[]interface{} parsing failed")
		}
		for i, entry := range value {
			value[i] = fixAddress(entry)
		}
		return value
	case map[string]interface{}:
		value, ok := data.(map[string]interface{})
		if !ok {
			log.Error("map[string]interface{} parsing failed")
		}
		for key, entry := range value {
			value[key] = fixAddress(entry)
		}
		return value
	}
	return data
}
