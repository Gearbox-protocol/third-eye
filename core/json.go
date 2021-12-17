package core

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
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

func (j *JsonBalance) Copy() JsonBalance {
	var newJB = make(JsonBalance)
	for k, v := range newJB {
		newJB[k] = &BalanceType{
			BI: NewBigInt(v.BI),
			F:  v.F,
		}
	}
	return newJB
}
