package core

import (
	"github.com/Gearbox-protocol/third-eye/log"
)

type Level struct {
	order            []string
	adapters         map[string][]string
	adapterTypeIndex int
	index            int
	len              int
}

func NewLevel(lvl []string) Level {
	return Level{
		order:    lvl,
		adapters: make(map[string][]string),
		index:    -1,
	}
}

func (lvl *Level) Add(name, address string) {
	lvl.adapters[name] = append(lvl.adapters[name], address)
	lvl.len++
}

func (lvl *Level) Next() bool {
	if len(lvl.order) == lvl.adapterTypeIndex {
		return false
	}
	name := lvl.order[lvl.adapterTypeIndex]
	lvl.index++
	if lvl.index < len(lvl.adapters[name]) {
		return true
	}
	for len(lvl.order) > lvl.adapterTypeIndex && len(lvl.adapters[name]) == lvl.index {
		lvl.adapterTypeIndex++
		if len(lvl.order) > lvl.adapterTypeIndex {
			name = lvl.order[lvl.adapterTypeIndex]
			lvl.index = 0
		}
	}
	if len(lvl.order) == lvl.adapterTypeIndex {
		return false
	}
	return true
}

func (lvl *Level) Get() string {
	name := lvl.order[lvl.adapterTypeIndex]
	return lvl.adapters[name][lvl.index]
}

func (lvl *Level) Reset() {
	lvl.adapterTypeIndex = 0
	lvl.index = -1
}

func (lvl *Level) Len() int {
	return lvl.len
}

func (lvl *Level) First() string {
	defer lvl.Reset()
	if lvl.Next() {
		v := lvl.Get()
		return v
	}
	return ""
}

func (lvl *Level) Details() {
	for _, name := range lvl.order {
		log.Info(name, len(lvl.adapters[name]))
	}
}

func (lvl *Level) GetAddressByName(name string) []string {
	return lvl.adapters[name]
}
