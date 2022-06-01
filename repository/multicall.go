package repository

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/common"
)

type PriceCallParams struct {
	Address common.Address
}

// multicall for getting price in batch
// For only getting the prices for calculating the treasury value
func (repo *Repository) getPricesInBatch(oracle string, blockNum int64, successRequired bool, tokenAddrs []string, poolForDieselRate []string) (prices []*big.Int, dieselRates []*big.Int) {
	calls := []multicall.Multicall2Call{}

	if oracle == "" {
		for _ = range tokenAddrs {
			prices = append(prices, new(big.Int))
		}
		for _ = range poolForDieselRate {
			dieselRates = append(dieselRates, new(big.Int))
		}
		return
	}
	oracleABI := schemas.GetAbi(ds.PriceOracle)
	for _, token := range tokenAddrs {
		tokenObj := repo.GetToken(token)
		amount := utils.GetExpInt(tokenObj.Decimals)
		data, err := oracleABI.Pack("convert", amount, common.HexToAddress(token), common.HexToAddress(repo.GetUSDCAddr()))
		log.CheckFatal(err)
		calls = append(calls, multicall.Multicall2Call{
			Target:   common.HexToAddress(oracle),
			CallData: data,
		})
	}

	poolABI := schemas.GetAbi(ds.Pool)
	for _, pool := range poolForDieselRate {
		data, err := poolABI.Pack("getDieselRate_RAY")
		log.CheckFatal(err)
		calls = append(calls, multicall.Multicall2Call{
			Target:   common.HexToAddress(pool),
			CallData: data,
		})
	}
	// call
	result := core.MakeMultiCall(repo.client, blockNum, successRequired, calls)

	for i, entry := range result {
		// token price
		if i < len(tokenAddrs) {
			price := big.NewInt(0)
			if entry.Success {
				value, err := oracleABI.Unpack("convert", entry.ReturnData)
				log.CheckFatal(err)
				price = (value[0]).(*big.Int)
				price = new(big.Int).Mul(price, big.NewInt(100))
			}
			prices = append(prices, price)
		} else {
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
	}
	return
}
