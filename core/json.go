package core

import (
	"encoding/json"
	"database/sql/driver"
	"fmt"
)
type Json map[string]string
func (j *Json) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (z *Json) Scan(value interface{}) error {
	out:=map[string]string{}
	switch t := value.(type) {
	case []byte:
		err := json.Unmarshal(value.([]byte), &out)
		*z = Json(out)
		return err
	default:
		return fmt.Errorf("Could not scan type %T", t)
	}
}