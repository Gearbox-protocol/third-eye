package address_provider

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
)

func (mdl *AddressProvider) setPriceOracle() {
	priceOracles := mdl.getPriceOracleMap()
	if mdl.priceOracles == nil {
		for _strBlockNum, oracle := range priceOracles {
			oracleBlockNum, err := strconv.ParseInt(_strBlockNum, 10, 64)
			log.CheckFatal(err)
			mdl.priceOracles = append(mdl.priceOracles, blockAndOracle{
				blockNum:    oracleBlockNum,
				priceOracle: schemas.PriceOracleT(oracle.(string)),
			})
		}
		sort.SliceStable(mdl.priceOracles, func(i, j int) bool {
			return mdl.priceOracles[i].blockNum < mdl.priceOracles[j].blockNum
		})
	}
}

func (mdl *AddressProvider) addPriceOracleLegacy(blockNum int64, priceOracle schemas.PriceOracleT) {
	priceOraclesMap := mdl.getPriceOracleMap()
	priceOraclesMap[fmt.Sprintf("%d", blockNum)] = string(priceOracle)
	mdl.Details["priceOracles"] = priceOraclesMap
	//
	mdl.setPriceOracle()
	mdl.priceOracles = append(mdl.priceOracles, blockAndOracle{
		blockNum:    blockNum,
		priceOracle: priceOracle,
	})
}

// only valid for 1,2 version , can implement for v3 but not required . After 310 , each pool has different priceoracle
// and the address provider no longer register the priceoracle.
func (mdl *AddressProvider) GetPriceOracleLegacy(version core.VersionType) schemas.PriceOracleT {
	mdl.setPriceOracle()
	if version.Eq(1) {
		return mdl.priceOracles[0].priceOracle
	} else if version.Eq(2) {
		return mdl.priceOracles[1].priceOracle
	}
	log.Fatal("only valid for 1 and 2 version")
	return ""
}
