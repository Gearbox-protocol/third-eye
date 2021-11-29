package core

import (
	"database/sql/driver"
	"log"
	"strings"
)

type Hstore struct {
	store map[string]string
}

func (dst *Hstore) Scan(value interface{}) error {
	dst.store = make(map[string]string)
	if value == nil {
		return nil
	}
	var b byte
	pair := [][]byte{{}, {}}
	pi := 0
	inQuote := false
	didQuote := false
	sawSlash := false
	bindex := 0
	valueStr, ok := value.(string)
	if !ok {
		log.Fatal("hstore scan failed")
	}
	for bindex, b = range []byte(valueStr) {
		if sawSlash {
			pair[pi] = append(pair[pi], b)
			sawSlash = false
			continue
		}

		switch b {
		case '\\':
			sawSlash = true
			continue
		case '"':
			inQuote = !inQuote
			if !didQuote {
				didQuote = true
			}
			continue
		default:
			if !inQuote {
				switch b {
				case ' ', '\t', '\n', '\r':
					continue
				case '=':
					continue
				case '>':
					pi = 1
					didQuote = false
					continue
				case ',':
					s := string(pair[1])
					if !didQuote && len(s) == 4 && strings.ToLower(s) == "null" {
						dst.store[string(pair[0])] = ""
					} else {
						dst.store[string(pair[0])] = string(pair[1])
					}
					pair[0] = []byte{}
					pair[1] = []byte{}
					pi = 0
					continue
				}
			}
		}
		pair[pi] = append(pair[pi], b)
	}
	if bindex > 0 {
		s := string(pair[1])
		if !didQuote && len(s) == 4 && strings.ToLower(s) == "null" {
			dst.store[string(pair[0])] = ""
		} else {
			dst.store[string(pair[0])] = string(pair[1])
		}
	}
	return nil
}

func NewHstore() Hstore {
	return Hstore{
		store: make(map[string]string),
	}
}

func HstoreFromMap(m map[string]string) Hstore {
	return Hstore{
		store: m,
	}
}

func (dst *Hstore) GetMap() map[string]string {
	return dst.store
}

func (dst *Hstore) Get(key string) string {
	return dst.store[key]
}

func (dst *Hstore) Set(key, value string) {
	dst.store[key] = value
}

func (dst *Hstore) Remove(key string) {
	delete(dst.store, key)
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
	return out, nil
}
