package ds

type AdapterKit struct {
	// address or type to adapter
	addressMap         map[string]SyncAdapterI
	levels             []Level
	adapterNameToLevel map[string]int
	len                int
}

func (kit *AdapterKit) init() {
	// wrapper should have same topics
	kit.AddLevel([]string{AddressProvider})
	// REVERT_ADMIN_WRAPPER: ACL, AccountFactory, GearToken, ContractRegister, RebaseToken
	kit.AddLevel([]string{PriceOracle, AdminWrapper, RebaseToken})
	// REVERT_POOL_WRAPPER
	kit.AddLevel([]string{PoolWrapper, AccountManager, CompositeChainlinkPF})
	// another level created bcz of poolKeeper.
	kit.AddLevel([]string{PoolQuotaWrapper, ChainlinkWrapper}) // ChainlinkPriceFeed
	// REVERT_CM_WRAPPER
	kit.AddLevel([]string{CMWrapper, AggregatedQueryFeedWrapper, LMRewardsv2, LMRewardsv3})
	// REVERT_CF_WRAPPER
	kit.AddLevel([]string{CFWrapper, CreditConfigurator, Treasury})
	// - we are dropping the uni check, so the dependency is reversed.
	//		(AggregatedQueryFeedWrapper => ChainlinkPriceFeed; so that deviation btw uniswap pool and chainlink can be calculated)
	//   Another reason being to get all the yearnPriceFeed in single go.
	// - CreditManager => CreditFilter/CreditConfigurator for creation only.
	// - AccountFactory => AccountManager => CreditManager; factory gets the accounts address to accountmanager for getting
	// pool => poolKeeper => CreditManager for credit sessions with multicall updateQuota
	//   all token transfers, in CreditManager filter transfer related to events on creditmanager
	// - Pool -> CreditManager; for getting the session for borrow/repay event on Pool
	// - Treasury dependent on pools, so it is last
	// - acl, PriceOracle and geartoken are independent
	// - creditconfigurator and core.CreditFilter are same dependent on creditmanager
	// - pool -> dieseltokens -> LMRewardsv2 and LMRewardsv3
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
	return kit.GetAdapter(adapterAddr)
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

func (kit *AdapterKit) GetAdapterAddressByName(name string) []string {
	lvlIndex := kit.adapterNameToLevel[name]
	return kit.levels[lvlIndex].GetAddressByName(name)
}

func (kit *AdapterKit) Details() {
	for _, lvl := range kit.levels {
		lvl.Details()
	}
}
