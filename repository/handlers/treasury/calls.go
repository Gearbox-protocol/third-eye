package treasury

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/repository/handlers"
	"github.com/ethereum/go-ethereum/common"
)

func dieselCalls(poolForDieselRate []string) (calls []multicall.Multicall2Call) {
	poolABI := core.GetAbi(ds.Pool)
	for _, pool := range poolForDieselRate {
		data, err := poolABI.Pack("getDieselRate_RAY")
		log.CheckFatal(err)
		calls = append(calls, multicall.Multicall2Call{
			Target:   common.HexToAddress(pool),
			CallData: data,
		})
	}
	return
}

func dieselAnswers(entries []multicall.Multicall2Result) (dieselRates []*big.Int) {
	poolABI := core.GetAbi(ds.Pool)
	for _, entry := range entries {
		dieselRate := big.NewInt(0)
		if entry.Success {
			if len(entry.ReturnData) != 0 {
				value, err := poolABI.Unpack("getDieselRate_RAY", entry.ReturnData)
				log.CheckFatal(err)
				dieselRate = (value[0]).(*big.Int)
			}
		} else {
			log.Fatal("dieselRates fetching failed")
		}
		dieselRates = append(dieselRates, dieselRate)
	}
	return
}

func v1PriceCalls(oracle common.Address, tokenAddrs []string, repo *handlers.TokensRepo) (calls []multicall.Multicall2Call) {
	oracleABI := core.GetAbi("PriceOracle")
	usdcToken := common.HexToAddress(repo.GetUSDCAddr())
	for _, tokenAddr := range tokenAddrs {
		amount := utils.GetExpInt(repo.GetToken(tokenAddr).Decimals)
		data, err := oracleABI.Pack("convert", amount, common.HexToAddress(tokenAddr), usdcToken)
		log.CheckFatal(err)
		calls = append(calls, multicall.Multicall2Call{
			Target:   oracle,
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

func v2PriceCalls(oracle common.Address, tokenAddrs []string) (calls []multicall.Multicall2Call) {
	oracleABI := core.GetAbi("PriceOraclev2")
	zeroAddr := common.Address{}
	for _, token := range tokenAddrs {
		data, err := oracleABI.Pack("getPrice", zeroAddr, common.HexToAddress(token))
		log.CheckFatal(err)
		calls = append(calls, multicall.Multicall2Call{
			Target:   oracle,
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
