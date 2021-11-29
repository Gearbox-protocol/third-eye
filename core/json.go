package core

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Json map[string]string

func (j *Json) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (z *Json) Scan(value interface{}) error {
	out := map[string]string{}
	switch t := value.(type) {
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
