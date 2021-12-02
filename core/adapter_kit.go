package core

import (
	"fmt"
	"github.com/Gearbox-protocol/third-eye/utils"
)

type AdapterKit struct {
	addressMap       map[string]SyncAdapterI
	order            []string
	adapters         map[string][]SyncAdapterI
	adapterTypeIndex int
	index            int
	len              int
}

func NewAdapterKit(order []string) *AdapterKit {
	adapterTypes := []string{
		"ACL",
		"AddressProvider",
		"AccountFactory",
		"Pool",
		"CreditManager",
		"PriceOracle",
		"PriceFeed",
		"ContractRegister",
		"CreditFilter",
	}
	otherAdapters := []string{}
	for _, name := range adapterTypes {
		if !utils.Contains(order, name) {
			otherAdapters = append(otherAdapters, name)
		}
	}
	return &AdapterKit{
		addressMap: make(map[string]SyncAdapterI),
		adapters:   make(map[string][]SyncAdapterI),
		order:      append(order, otherAdapters...),
		index:      -1,
	}
}

func (kit *AdapterKit) Add(adapter SyncAdapterI) {
	kit.addressMap[adapter.GetAddress()] = adapter
	kit.adapters[adapter.GetName()] = append(kit.adapters[adapter.GetName()], adapter)
	kit.len++
}

func (kit *AdapterKit) Next() bool {
	if len(kit.order) == kit.adapterTypeIndex {
		return false
	}
	name := kit.order[kit.adapterTypeIndex]
	kit.index++
	if kit.index < len(kit.adapters[name]) {
		return true
	}
	for len(kit.order) > kit.adapterTypeIndex && len(kit.adapters[name]) == kit.index {
		kit.adapterTypeIndex++
		if len(kit.order) > kit.adapterTypeIndex {
			name = kit.order[kit.adapterTypeIndex]
			kit.index = 0
		}
	}
	if len(kit.order) == kit.adapterTypeIndex {
		return false
	}
	return true
}

func (kit *AdapterKit) First() SyncAdapterI {
	if kit.Next() {
		v := kit.Get()
		kit.Reset()
		return v
	}
	kit.Reset()
	return nil
}

func (kit *AdapterKit) Get() SyncAdapterI {
	name := kit.order[kit.adapterTypeIndex]
	return kit.adapters[name][kit.index]
}

func (kit *AdapterKit) GetAdapter(addr string) SyncAdapterI {
	return kit.addressMap[addr]
}

func (kit *AdapterKit) Reset() {
	kit.adapterTypeIndex = 0
	kit.index = -1
}

func (kit *AdapterKit) DisableSyncAdapter(addr string) {
	adapter := kit.addressMap[addr]
	if adapter != nil {
		adapter.Disable()
	}
}

func (kit *AdapterKit) Len() int {
	return kit.len
}

func (kit *AdapterKit) Details() string {
	str := ""
	for _, name := range kit.order {
		str += fmt.Sprintf("%s: %d\n", name, len(kit.adapters[name]))
	}
	return str
}
