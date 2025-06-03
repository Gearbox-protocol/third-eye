package treasury

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/repository/handlers"
	"github.com/ethereum/go-ethereum/common"
)

func dieselCalls(poolForDieselRate []*schemas.UTokenAndPool) (calls []multicall.Multicall2Call) {
	poolv2GetDiesleRateData := func() []byte {
		poolABI := core.GetAbi(ds.Pool)
		data, err := poolABI.Pack("getDieselRate_RAY")
		log.CheckFatal(err)
		return data
	}()
	poolv3ABI := core.GetAbi("Poolv3")
	for _, pool := range poolForDieselRate {
		var data []byte
		// for 300
		if pool.Version.MoreThanEq(core.NewVersion(300)) {
			value, err := poolv3ABI.Pack("convertToAssets", core.RAY)
			log.CheckFatal(err)
			data = value
		} else { // for v1, v2
			data = poolv2GetDiesleRateData
		}
		// for all versions
		calls = append(calls, multicall.Multicall2Call{
			Target:   common.HexToAddress(pool.Pool),
			CallData: data,
		})
	}
	return
}

func dieselAnswers(entries []multicall.Multicall2Result) (dieselRates []*big.Int) {
	for _, entry := range entries {
		dieselRate := big.NewInt(0)
		if entry.Success && len(entry.ReturnData) != 0 {
			dieselRate = new(big.Int).SetBytes(entry.ReturnData) // all return data is 32 bytes
		} else {
			log.Fatal("dieselRates fetching failed", entry.Success, len(entry.ReturnData))
		}
		dieselRates = append(dieselRates, dieselRate)
	}
	return
}

func v1PriceCalls(oracle schemas.PriceOracleT, tokenAddrs []string, repo *handlers.TokensRepo) (calls []multicall.Multicall2Call) {
	oracleABI := core.GetAbi("PriceOracle")
	usdcToken := core.GetToken(core.GetBaseChainId(repo.GetClient()), "USDC")
	for _, tokenAddr := range tokenAddrs {
		amount := utils.GetExpInt(repo.GetToken(tokenAddr).Decimals)
		data, err := oracleABI.Pack("convert", amount, common.HexToAddress(tokenAddr), usdcToken)
		log.CheckFatal(err)
		calls = append(calls, multicall.Multicall2Call{
			Target:   oracle.Hex(),
			CallData: data,
		})
	}
	return
}

func v1PriceAnswers(entries []multicall.Multicall2Result) (prices []*big.Int) {
	oracleABI := core.GetAbi("PriceOracle")
	for _, entry := range entries {
		price := big.NewInt(0)
		if entry.Success {
			value, err := oracleABI.Unpack("convert", entry.ReturnData)
			log.CheckFatal(err)
			price = (value[0]).(*big.Int)
			price = new(big.Int).Mul(price, big.NewInt(100))
		}
		prices = append(prices, price)
	}
	return
}

func v2PriceCalls(oracle schemas.PriceOracleT, tokenAddrs []string) (calls []multicall.Multicall2Call) {
	oracleABI := core.GetAbi("PriceOraclev2")
	for _, token := range tokenAddrs {
		data, err := oracleABI.Pack("getPrice", common.HexToAddress(token))
		log.CheckFatal(err)
		calls = append(calls, multicall.Multicall2Call{
			Target:   oracle.Hex(),
			CallData: data,
		})
	}
	return
}

func v2PriceAnswers(entries []multicall.Multicall2Result) (prices []*big.Int) {
	oracleABI := core.GetAbi("PriceOraclev2")
	for _, entry := range entries {
		price := big.NewInt(0)
		if entry.Success {
			value, err := oracleABI.Unpack("getPrice", entry.ReturnData)
			log.CheckFatal(err)
			price = (value[0]).(*big.Int)
		}
		prices = append(prices, price)
	}
	return
}
