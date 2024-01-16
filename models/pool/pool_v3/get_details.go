package pool_v3

import (
	dcv3 "github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressorv3"
	"github.com/Gearbox-protocol/sdk-go/log"
)

func (mdl *Poolv3) setDetailsByKey(key, value string) {
	if mdl.Details == nil {
		mdl.Details = make(map[string]interface{})
	}
	mdl.Details[key] = value
}
func (mdl *Poolv3) setPoolQuotaKeeper() {
	if mdl.getPoolKeeper() != "" {
		return
	}
	poolKeeper, err := mdl.contract.PoolQuotaKeeper(nil)
	log.CheckFatal(err)
	mdl.setDetailsByKey("PoolKeeper", poolKeeper.Hex())
}

func (mdl *Poolv3) getPoolKeeper() string {
	return mdl.GetDetailsByKey("PoolKeeper")
}
func (mdl *Poolv3) getFarmedUSDCv3() string {
	return mdl.GetDetailsByKey("farmedUSDCv3")
}
func (mdl *Poolv3) getdUSDC() string {
	return mdl.GetDetailsByKey("dUSDC")
}
func (mdl *Poolv3) getPoolv2() string {
	return mdl.GetDetailsByKey("poolv2")
}
func (mdl *Poolv3) getZapUnderlying() string {
	return mdl.GetDetailsByKey("USDC-farmedUSDCv3")
}
func (mdl *Poolv3) getZapPoolv2() string {
	return mdl.GetDetailsByKey("dUSDC-farmedUSDCv3")
}

func (mdl *Poolv3) setZapper() {
	if mdl.GetDetailsByKey("farmedUSDCv3") != "" { // if zapper already set
		return
	}
	pools, found := mdl.Repo.GetDCWrapper().GetPoolListv3()
	if !found {
		return
	}
	var poolToCheck dcv3.PoolData
	for _, pool := range pools {
		if pool.Addr.Hex() == mdl.Address {
			poolToCheck = pool
			break
		}
	}

	// out = farmedUSDCv3, dUSDCv3
	for _, zapper := range poolToCheck.Zappers {
		if zapper.TokenIn == poolToCheck.Underlying && zapper.TokenOut != poolToCheck.DieselToken { // tokenIn = USDC, tokenOut != dUSDCv3
			mdl.setDetailsByKey("USDC-farmedUSDCv3", zapper.Zapper.Hex())
			mdl.setDetailsByKey("farmedUSDCv3", zapper.TokenOut.Hex())
		}
	}
	for _, zapper := range poolToCheck.Zappers {
		if zapper.TokenIn != poolToCheck.Underlying &&
			zapper.TokenOut.Hex() == mdl.GetDetailsByKey("farmedUSDCv3") { // tokenIn = dUSDC, tokenOut = farmedUSDCv3
			mdl.setDetailsByKey("dUSDC-farmedUSDCv3", zapper.Zapper.Hex())
			mdl.setDetailsByKey("dUSDC", zapper.TokenIn.Hex())
		}
	}

	if mdl.GetDetailsByKey("dUSDC") == "" {
		log.Fatal("Can't get dUSDC from zapper for ", mdl.Address)
	}
	dieselTokenToPool := mdl.Repo.GetDieselTokens()
	pool, ok := dieselTokenToPool[mdl.GetDetailsByKey("dUSDC")]
	if !ok {
		log.Fatal("Can't get poolv2 from poolv3 ", mdl.Address)
	}
	mdl.setDetailsByKey("poolv2", pool.Pool)
}
