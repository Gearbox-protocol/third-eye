package ds

type AdapterKit struct {
	addressMap         map[string]SyncAdapterI
	levels             []Level
	adapterNameToLevel map[string]int
	len                int
}

func (x AdapterKit) GetAddrsMap() map[string]SyncAdapterI {
	return x.addressMap
}
func (kit *AdapterKit) init() {
	kit.AddLevel([]string{AddressProvider})
	kit.AddLevel([]string{ContractRegister, PriceOracle, ACL, AccountFactory, GearToken})
	kit.AddLevel([]string{Pool, AccountManager, ChainlinkPriceFeed, CompositeChainlinkPF})
	kit.AddLevel([]string{CreditManager, AggregatedBlockFeed, PoolLMRewards})
	kit.AddLevel([]string{CreditFilter, CreditConfigurator, Treasury})
	// - we are dropping the uni check, so the dependency is reversed.
	//		(AggregatedBlockFeed => ChainlinkPriceFeed; so that deviation btw uniswap pool and chainlink can be calculated)
	//   Another reason being to get all the yearnPriceFeed in single go.
	// - CreditManager => CreditFilter/CreditConfigurator for creation only.
	// - AccountFactory => AccountManager => CreditManager; factory gets the accounts address to accountmanager for getting
	//   all token transfers, in CreditManager filter transfer related to events on creditmanager
	// - Pool -> CreditManager; for getting the session for borrow/repay event on Pool
	// - Treasury dependent on pools, so it is last
	// - acl, PriceOracle and geartoken are independent
	// - creditconfigurator and core.CreditFilter are same dependent on creditmanager
	// - pool -> dieseltokens -> PoolLMRewards
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
