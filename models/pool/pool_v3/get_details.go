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
func (mdl *Poolv3) getdUSDC() string {
	return mdl.getAddrFromDetails("dUSDC")
}
func (mdl *Poolv3) getPoolv2() string {
	return mdl.getAddrFromDetails("poolv2")
}
func (mdl *Poolv3) getZapPoolv2() string {
	return mdl.getAddrFromDetails("dUSDC-farmedUSDCv3")
}
func (mdl *Poolv3) getAddrFromDetails(key string) string {
	if mdl.Details == nil || mdl.Details[key] == nil { // if zapper already set
		return ""
	}
	return mdl.GetDetailsByKey(key)
}

func (mdl *Poolv3) checkIfZapAddr(addr string) bool {
	return utils.Contains(mdl.zappers.GetZapper(), addr)
}

func (mdl *Poolv3) setZapper() {
	mdl.zappers.Load(mdl.Details)
	if len(*mdl.zappers) != 0 { // if zapper already set
		return
	}
	//
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
	// farmingPools := core.GetFarmingPoolsToSymbolByChainId(core.GetChainId(mdl.Client))
	var ETHAddr common.Address
	if poolToCheck.Underlying == syms.Tokens["WETH"] {
		// TODO ? why is eth address not set in tokens
		// ETHAddr = syms.Tokens["ETH"]
		ETHAddr = common.HexToAddress("0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee")
	}

	zappers := Zappers{}
	// out = farmedUSDCv3, dUSDCv3
	for _, zapper := range poolToCheck.Zappers {
		if zapper.TokenIn == poolToCheck.Underlying && zapper.TokenOut.Hex() != mdl.Address { // tokenIn = USDC,  tokenOut is farming Pool(!= dUSDCv3)
			zappers = append(zappers, Zapper{
				Zapper:  zapper.Zapper.Hex(),
				TokenIn: zapper.TokenIn.Hex(),
				Farm:    zapper.TokenOut.Hex(),
			})
		}
	}
	for _, zapper := range poolToCheck.Zappers { // for ETH
		if zapper.TokenIn == ETHAddr && utils.Contains(zappers.GetFarm(), zapper.TokenOut.Hex()) {
			zappers = append(zappers, Zapper{
				Zapper:  zapper.Zapper.Hex(),
				TokenIn: zapper.TokenIn.Hex(),
				Farm:    zapper.TokenOut.Hex(),
			})
		}
		if zapper.TokenIn != poolToCheck.Underlying && zapper.TokenIn != ETHAddr &&
			utils.Contains(zappers.GetFarm(), zapper.TokenOut.Hex()) { // tokenIn = dUSDC, tokenOut = farmedUSDCv3
			mdl.setDetailsByKey("dUSDC-farmedUSDCv3", zapper.Zapper.Hex())
			mdl.setDetailsByKey("dUSDC", zapper.TokenIn.Hex())
		}
	}
	(*mdl.zappers) = zappers
	//
	if len(*mdl.zappers) == 0 {
		log.Fatal("Can't get zapper for ", mdl.Address)
	}
	if log.GetBaseNet(core.GetChainId(mdl.Client)) == "MAINNET" && // only on mainnet
		mdl.getdUSDC() != "" { // is not null for only USDC, DAI, WETH and WBTC
		dieselTokenToPool := mdl.Repo.GetDieselTokens()
		pool, ok := dieselTokenToPool[mdl.GetDetailsByKey("dUSDC")]
		if !ok {
			log.Fatalf("Can't get poolv2(dieselToken: %s) from poolv3: %s ", mdl.GetDetailsByKey("dUSDC"), mdl.Address)
		}
		mdl.setDetailsByKey("poolv2", pool.Pool)
	}
}
