package repository

import (
	"math/big"

	"github.com/Gearbox-protocol/third-eye/artifacts/multicall"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func getMultiCallAddr() string {
	return "0x5BA1e12693Dc8F9c48aAD8770482f4739bEeD696"
}

type PriceCallParams struct {
	Address common.Address
}

func (repo *Repository) MakeMultiCall(blockNum int64, successRequired bool, calls []multicall.Multicall2Call) []multicall.Multicall2Result {
	contract := getMultiCallContract(repo.client)
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
	}
	var result []multicall.Multicall2Result
	var tmpCalls []multicall.Multicall2Call
	callsInd := 0
	callsLen := len(calls)
	for callsInd < callsLen {
		for i := 0; i < 10 && callsInd < callsLen; i++ {
			tmpCalls = append(tmpCalls, calls[callsInd])
			callsInd++
		}
		tmpResult, err := contract.TryAggregate(opts, successRequired, tmpCalls)
		log.CheckFatal(err)
		result = append(result, tmpResult...)
		tmpCalls = []multicall.Multicall2Call{}
	}
	return result
}

func (repo *Repository) getPricesInBatch(blockNum int64, successRequired bool, tokenAddrs []string, poolForDieselRate []string) (prices []*big.Int, dieselRates []*big.Int) {
	calls := []multicall.Multicall2Call{}

	oracle, err := repo.GetActivePriceOracle()
	log.CheckFatal(err)
	oracleABI := core.GetAbi(core.PriceOracle)
	for _, token := range tokenAddrs {
		tokenObj, err := repo.getTokenWithError(token)
		log.CheckFatal(err)
		amount := utils.GetExpInt(tokenObj.Decimals)
		data, err := oracleABI.Pack("convert", amount, common.HexToAddress(token), common.HexToAddress(repo.USDCAddr))
		log.CheckFatal(err)
		calls = append(calls, multicall.Multicall2Call{
			Target:   common.HexToAddress(oracle),
			CallData: data,
		})
	}

	poolABI := core.GetAbi(core.Pool)
	for _, pool := range poolForDieselRate {
		data, err := poolABI.Pack("getDieselRate_RAY")
		log.CheckFatal(err)
		calls = append(calls, multicall.Multicall2Call{
			Target:   common.HexToAddress(pool),
			CallData: data,
		})
	}
	// call
	result := repo.MakeMultiCall(blockNum, successRequired, calls)

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
				value, err := poolABI.Unpack("getDieselRate_RAY", entry.ReturnData)
				log.CheckFatal(err)
				dieselRate = (value[0]).(*big.Int)
			} else {
				log.Fatal("dieselRates fetching failed")
			}
			dieselRates = append(dieselRates, dieselRate)
		}
	}
	return
}

func getMultiCallContract(client *ethclient.Client) *multicall.Multicall {
	contract, err := multicall.NewMulticall(common.HexToAddress(getMultiCallAddr()), client)
	log.CheckFatal(err)
	return contract
}
