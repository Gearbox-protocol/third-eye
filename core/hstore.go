package core

import (
	"database/sql/driver"
	"fmt"
)

type Hstore struct {
	store map[string]string
}

func (dst *Hstore) Scan(src interface{}) error {
	return nil
}

func NewHstore() Hstore {
	return Hstore{
		store: make(map[string]string),
	}
}

func (dst *Hstore) Get(key string) string {
	return dst.store[key]
}
func (dst *Hstore) Set(key, value string) {
	dst.store[key] = value
}

func (src Hstore) Value() (driver.Value, error) {
	var out string
	isFirst := true
	for k, v := range src.store {
		if !isFirst {
			out += ","
		} else {
			isFirst = false
		}
		out += k
		out += "=>"
		out += v

	}
	fmt.Println(out)
	return out, nil
}
