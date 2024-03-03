package pool_v3

import (
	dcv3 "github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressorv3"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm/utils"
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
	return mdl.getAddrFromDetails("PoolKeeper")
}
func (mdl *Poolv3) getFarmedUSDCv3() string {
	return mdl.getAddrFromDetails("farmedUSDCv3")
}
func (mdl *Poolv3) getdUSDC() string {
	return mdl.getAddrFromDetails("dUSDC")
}
func (mdl *Poolv3) getPoolv2() string {
	return mdl.getAddrFromDetails("poolv2")
}
func (mdl *Poolv3) getZapUnderlying() string {
	return mdl.getAddrFromDetails("USDC-farmedUSDCv3")
}
func (mdl *Poolv3) getZapPoolv2() string {
	return mdl.getAddrFromDetails("dUSDC-farmedUSDCv3")
}
func (mdl *Poolv3) getZapETH() string {
	return mdl.getAddrFromDetails("ETH-farmedETHv3")
}
func (mdl *Poolv3) getAddrFromDetails(key string) string {
	if mdl.Details == nil || mdl.Details[key] == nil { // if zapper already set
		return ""
	}
	return mdl.GetDetailsByKey(key)
}

func (mdl *Poolv3) checkIfZapAddr(addr string) bool {
	return utils.Contains([]string{mdl.getZapUnderlying(), mdl.getZapPoolv2(), mdl.getZapETH()}, addr)
}
func (mdl *Poolv3) setZapper() {
	if mdl.getAddrFromDetails("farmedUSDCv3") != "" { // if zapper already set
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

	// eth contract has 3 zappers.
	// 1. USDC-farmedUSDCv3
	// 2. dUSDC-farmedUSDCv3
	// 3. ETH-farmedETHv3
	syms := core.GetSymToAddrByChainId(core.GetChainId(mdl.Client))
	farmingPools := core.GetFarmingPoolsToSymbolByChainId(core.GetChainId(mdl.Client))
	var ETHAddr common.Address
	if poolToCheck.Underlying == syms.Tokens["WETH"] {
		// TODO ? why is eth address not set in tokens
		// ETHAddr = syms.Tokens["ETH"]
		ETHAddr = common.HexToAddress("0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee")
	}

	// out = farmedUSDCv3, dUSDCv3
	for _, zapper := range poolToCheck.Zappers {
		if zapper.TokenIn == poolToCheck.Underlying && farmingPools[zapper.TokenOut] != "" { // tokenIn = USDC,  tokenOut is farming Pool(!= dUSDCv3)
			mdl.setDetailsByKey("USDC-farmedUSDCv3", zapper.Zapper.Hex())
			mdl.setDetailsByKey("farmedUSDCv3", zapper.TokenOut.Hex())
		}
	}
	for _, zapper := range poolToCheck.Zappers {
		if zapper.TokenIn == ETHAddr && zapper.TokenOut.Hex() == mdl.GetDetailsByKey("farmedUSDCv3") {
			mdl.setDetailsByKey("ETH-farmedETHv3", zapper.Zapper.Hex())
		}
		if zapper.TokenIn != poolToCheck.Underlying && zapper.TokenIn != ETHAddr &&
			zapper.TokenOut.Hex() == mdl.GetDetailsByKey("farmedUSDCv3") { // tokenIn = dUSDC, tokenOut = farmedUSDCv3
			mdl.setDetailsByKey("dUSDC-farmedUSDCv3", zapper.Zapper.Hex())
			mdl.setDetailsByKey("dUSDC", zapper.TokenIn.Hex())
		}
	}
	if mdl.Details["dUSDC-farmedUSDCv3"] != nil {
		log.Fatal("Can't get dUSDC from zapper for ", mdl.Address)
	}
	if log.GetBaseNet(core.GetChainId(mdl.Client)) == "MAINNET" { // only on mainnet
		if mdl.GetDetailsByKey("dUSDC") == "" {
			log.Fatal("Can't get dUSDC from zapper for ", mdl.Address)
		}
		dieselTokenToPool := mdl.Repo.GetDieselTokens()
		pool, ok := dieselTokenToPool[mdl.GetDetailsByKey("dUSDC")]
		if !ok {
			log.Fatalf("Can't get poolv2(dieselToken: %s) from poolv3: %s ", mdl.GetDetailsByKey("dUSDC"), mdl.Address)
		}
		mdl.setDetailsByKey("poolv2", pool.Pool)
	}
}
