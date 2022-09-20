package address_provider

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
)

type blockAndOracle struct {
	priceOracle string
	blockNum    int64
}
type AddressProvider struct {
	*ds.SyncAdapter
	priceOracles []blockAndOracle `json:"-"`
}

func NewAddressProvider(addr string, client core.ClientI, repo ds.RepositoryI) *AddressProvider {
	return NewAddressProviderFromAdapter(
		ds.NewSyncAdapter(addr, ds.AddressProvider, -1, client, repo),
	)
}

func NewAddressProviderFromAdapter(adapter *ds.SyncAdapter) *AddressProvider {
	obj := &AddressProvider{
		SyncAdapter: adapter,
	}
	return obj
}

func (mdl *AddressProvider) GetDetailsByKey(strBlockNum string) string {
	blockNum, err := strconv.ParseInt(strBlockNum, 10, 64)
	if err != nil { // if input is not number make call on the embedded struct
		return mdl.SyncAdapter.GetDetailsByKey(strBlockNum)
	}
	priceOracles := mdl.getPriceOracleMap()
	if mdl.priceOracles == nil {
		for strBlockNum, oracle := range priceOracles {
			oracleBlockNum, err := strconv.ParseInt(strBlockNum, 10, 64)
			log.CheckFatal(err)
			mdl.priceOracles = append(mdl.priceOracles, blockAndOracle{
				blockNum:    oracleBlockNum,
				priceOracle: oracle.(string),
			})
		}
		sort.SliceStable(mdl.priceOracles, func(i, j int) bool {
			return mdl.priceOracles[i].blockNum < mdl.priceOracles[j].blockNum
		})
	}
	ind := sort.Search(len(mdl.priceOracles), func(i int) bool {
		return mdl.priceOracles[i].blockNum > blockNum
	})
	return mdl.priceOracles[ind-1].priceOracle
}

func (mdl *AddressProvider) addPriceOracle(blockNum int64, priceOracle string) {
	priceOraclesMap := mdl.getPriceOracleMap()
	priceOraclesMap[fmt.Sprintf("%d", blockNum)] = priceOracle
	mdl.Details["priceOracles"] = priceOraclesMap
	//
	mdl.priceOracles = append(mdl.priceOracles, blockAndOracle{
		blockNum:    blockNum,
		priceOracle: priceOracle,
	})
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
