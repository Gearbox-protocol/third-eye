package core

type AdapterKit struct {
	addressMap         map[string]SyncAdapterI
	levels             []Level
	adapterNameToLevel map[string]int
	len                int
}

func (kit *AdapterKit) init() {
	kit.AddLevel([]string{AddressProvider})
	kit.AddLevel([]string{ContractRegister, PriceOracle, ACL, AccountFactory, GearToken})
	kit.AddLevel([]string{Pool})
	kit.AddLevel([]string{CreditManager, ChainlinkPriceFeed, YearnPriceFeed})
	kit.AddLevel([]string{CreditFilter})
}

func (kit *AdapterKit) AddLevel(lvl []string) {
	lvlIndex := len(kit.levels)
	kit.levels = append(kit.levels, NewLevel(lvl))
	for _, adapterName := range lvl {
		kit.adapterNameToLevel[adapterName] = lvlIndex
	}
	kit.len++
}

func NewAdapterKit() *AdapterKit {
	kit := &AdapterKit{
		addressMap:         make(map[string]SyncAdapterI),
		adapterNameToLevel: make(map[string]int),
	}
	kit.init()
	return kit
}

func (kit *AdapterKit) Add(adapter SyncAdapterI) {
	adapterName := adapter.GetName()
	adapterAddress := adapter.GetAddress()
	kit.addressMap[adapterAddress] = adapter
	lvlIndex := kit.adapterNameToLevel[adapterName]
	kit.levels[lvlIndex].Add(adapterName, adapterAddress)
}

func (kit *AdapterKit) Get(lvl int) SyncAdapterI {
	adapterAddr := kit.levels[lvl].Get()
	return kit.addressMap[adapterAddr]
}

func (kit *AdapterKit) Next(lvl int) bool {
	return kit.levels[lvl].Next()
}

func (kit *AdapterKit) Len() int {
	return kit.len
}

func (kit *AdapterKit) Reset(lvl int) {
	kit.levels[lvl].Reset()
}

func (kit *AdapterKit) LenOfLevel(lvl int) int {
	return kit.levels[lvl].Len()
}

func (kit *AdapterKit) First(lvl int) SyncAdapterI {
	adapterAddr := kit.levels[lvl].First()
	return kit.addressMap[adapterAddr]
}

func (kit *AdapterKit) GetAdapter(addr string) SyncAdapterI {
	return kit.addressMap[addr]
}

func (kit *AdapterKit) DisableSyncAdapter(addr string) {
	adapter := kit.addressMap[addr]
	if adapter != nil {
		adapter.Disable()
	}
}

func (kit *AdapterKit) GetAdapterAddressByName(name string) []string {
	lvlIndex := kit.adapterNameToLevel[name]
	return kit.levels[lvlIndex].GetAddressByName(name)
}

func (kit *AdapterKit) Details() {
	for _, lvl := range kit.levels {
		lvl.Details()
	}
}
