package address_provider

import (
	"sort"
	"strconv"
	"strings"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/common"
)

type blockAndOracle struct {
	priceOracle schemas.PriceOracleT
	blockNum    int64
}
type AddressProvider struct {
	*ds.SyncAdapter
	priceOracles []blockAndOracle `json:"-"`
	otherAddrs   []common.Address `json:"-"`
	// hashToContractName map[common.Hash]string `json:"-"`
}

func GetAddressProvider(client core.ClientI, addressProviderAddrs string) (firstAddressProvider string, otherAddrs []common.Address) {
	minVersion := core.NewVersion(10_000)
	for _, addr := range strings.Split(addressProviderAddrs, ",") {
		version := core.FetchVersion(addr, 0, client)
		if version.LessThan(minVersion) {
			minVersion = version
			firstAddressProvider = addr
		}
	}
	for _, addr := range strings.Split(addressProviderAddrs, ",") {
		if addr != firstAddressProvider {
			otherAddrs = append(otherAddrs, common.HexToAddress(addr))
		}
	}
	firstAddressProvider = common.HexToAddress(firstAddressProvider).Hex()
	return
}

func NewAddressProvider(client core.ClientI, repo ds.RepositoryI, apAddrs string) *AddressProvider {
	firstAddressProvider, _ := GetAddressProvider(client, apAddrs)
	return NewAddressProviderFromAdapter(
		ds.NewSyncAdapter(firstAddressProvider, ds.AddressProvider, -1, client, repo),
		apAddrs,
	)
}

func NewAddressProviderFromAdapter(adapter *ds.SyncAdapter, apAddrs string) *AddressProvider {
	obj := &AddressProvider{
		SyncAdapter: adapter,
	}
	if obj.Details == nil {
		obj.Details = core.Json{}
	}
	if core.GetChainId(adapter.Client) == 1337 {
		apAddrs = adapter.GetAddress()
	}
	_, otherAddrProviders := GetAddressProvider(obj.Client, apAddrs)
	obj.Details["others"] = otherAddrProviders
	obj.otherAddrs = otherAddrProviders

	for _, mcaddr := range strings.Split(utils.GetEnvOrDefault("MARKET_CONFIGURATORS", ""), ",") {
		if mcaddr != "" {
			obj.addMarketConfig(adapter.GetLastSync(), common.HexToAddress(mcaddr))
		}
	}
	// if core.GetChainId(adapter.Client) != 1337 {
	// addrv310 := core.GetAddressProvider(core.GetChainId(adapter.Client), core.NewVersion(300))
	// obj.hashToContractName = pkg.Initv310ContractHashMap(adapter.Client, common.HexToAddress(addrv310))
	// }

	return obj
}

func (mdl *AddressProvider) GetAllAddrsForLogs() []common.Address {
	all := append(mdl.otherAddrs, common.HexToAddress(mdl.Address))
	if mdl.Details["MARKET_FACTORY"] != nil {
		addr := mdl.Details["MARKET_FACTORY"].(string)
		all = append(all, common.HexToAddress(addr))
	}
	return all
}

func (mdl *AddressProvider) GetDetailsByKey(strBlockNum string) string {
	blockNum, err := strconv.ParseInt(strBlockNum, 10, 64)
	if err != nil { // if input is not number make call on the embedded struct
		return mdl.SyncAdapter.GetDetailsByKey(strBlockNum)
	}
	//
	mdl.setPriceOracle()
	ind := sort.Search(len(mdl.priceOracles), func(i int) bool {
		return mdl.priceOracles[i].blockNum > blockNum
	})
	return string(mdl.priceOracles[ind-1].priceOracle)
}

func (mdl *AddressProvider) getPriceOracleMap() map[string]interface{} {
	if mdl.Details == nil {
		mdl.Details = make(map[string]interface{})
	}
	// price oracles
	priceOracles, ok := mdl.Details["priceOracles"].(map[string]interface{})
	if !ok {
		if priceOracles == nil {
			priceOracles = map[string]interface{}{}
		}
	}
	return priceOracles
}
