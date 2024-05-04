package treasury

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/artifacts/priceOraclev3"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg/redstone"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/aggregated_block_feed"
	"github.com/Gearbox-protocol/third-eye/models/aggregated_block_feed/base_price_feed"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// used for treasury calculation and for remainingFunds on close v2
func (repo *TreasuryRepo) GetPricesInUSD(blockNum int64, tokenAddrs []string) core.JsonFloatMap {
	priceByToken := core.JsonFloatMap{}
	var tokenForCalls []string
	var poolForDieselRate []*schemas.UTokenAndPool
	for _, token := range tokenAddrs {
		uTokenAndPool := repo.tokens.GetDieselToken(token)
		if uTokenAndPool != nil {
			tokenForCalls = append(tokenForCalls, uTokenAndPool.UToken)
			poolForDieselRate = append(poolForDieselRate, uTokenAndPool)
		} else {
			tokenForCalls = append(tokenForCalls, token)
		}
	}
	priceOracle, version, _ := repo.adapters.GetActivePriceOracleByBlockNum(blockNum)
	prices, dieselRates := repo.getPricesInBatch(priceOracle, version, blockNum, false, tokenForCalls, poolForDieselRate)
	var poolIndex int
	for i, token := range tokenAddrs {
		var price *big.Int
		if repo.tokens.IsDieselToken(token) {
			dieselRate := dieselRates[poolIndex]
			poolIndex++
			price = new(big.Int).Mul(dieselRate, prices[i])
			price = utils.GetInt64(price, 27)
		} else {
			price = prices[i]
		}
		priceByToken[token] = utils.GetFloat64Decimal(price, 8)
	}
	return priceByToken
}

// multicall for getting price in batch
// For only getting the prices for calculating the treasury value
func (repo *TreasuryRepo) getPricesInBatch(oracle string, version core.VersionType, blockNum int64, successRequired bool, tokenAddrs []string, poolForDieselRate []*schemas.UTokenAndPool) (prices []*big.Int, dieselRates []*big.Int) {
	// base case
	if oracle == "" {
		for range tokenAddrs {
			prices = append(prices, new(big.Int))
		}
		for range poolForDieselRate {
			dieselRates = append(dieselRates, new(big.Int))
		}
		return
	}
	//
	// make calls
	calls := make([]multicall.Multicall2Call, 0, len(tokenAddrs)+len(poolForDieselRate))
	if version.Eq(1) {
		calls = append(calls, v1PriceCalls(common.HexToAddress(oracle), tokenAddrs, repo.tokens)...)
	} else if version.IsPriceInUSD() {
		calls = append(calls, v2PriceCalls(common.HexToAddress(oracle), tokenAddrs)...)
	}
	calls = append(calls, dieselCalls(poolForDieselRate)...)
	//
	// get response
	result := core.MakeMultiCall(repo.client, blockNum, successRequired, calls)

	// parse result
	if version.Eq(1) {
		prices = v1PriceAnswers(result[:len(tokenAddrs)])
	} else if version.IsPriceInUSD() {
		prices = v2PriceAnswers(result[:len(tokenAddrs)])
	}
	for ind, token := range tokenAddrs {
		if prices[ind] == nil {
			if price := repo.GetRedStonePrice(blockNum, oracle, token); price != nil {
				prices[ind] = price
			}
		}
	}
	dieselRates = dieselAnswers(result[len(tokenAddrs):])
	return
}

func (repo TreasuryRepo) GetRedStonemgr() redstone.RedStoneMgrI {
	return repo.redstoneMgr
}

func (repo TreasuryRepo) GetRedStonePrice(blockNum int64, oracle, token string) *big.Int {
	if adapter := repo.IsRedStoneAdapter(blockNum, oracle, token); adapter != nil {
		call, isQueryable := adapter.GetCalls(blockNum)
		if !isQueryable {
			return nil
		}
		results := core.MakeMultiCall(repo.client, blockNum, false, call)
		price := adapter.ProcessResult(blockNum, results)
		return price.PriceBI.Convert()
	}
	return nil
}

func (repo TreasuryRepo) IsRedStoneAdapter(blockNum int64, oracle string, token string) base_price_feed.QueryPriceFeedI {
	pon, err := priceOraclev3.NewPriceOraclev3(common.HexToAddress(oracle), repo.client)
	log.CheckFatal(err)
	priceFeed, err := pon.PriceFeeds(&bind.CallOpts{BlockNumber: big.NewInt(blockNum)}, common.HexToAddress(token))
	if err != nil {
		return nil
	}
	adapter := repo.adapters.GetAdapter(priceFeed.Hex())
	log.Info(adapter, token, oracle, priceFeed)
	if adapter != nil && // for chainlink or composite chainlink oracle
		adapter.GetName() == ds.QueryPriceFeed &&
		utils.Contains([]string{ds.RedStonePF, ds.CompositeRedStonePF}, adapter.GetDetailsByKey("pfType")) {
		return aggregated_block_feed.FromAdapter(adapter)
	}
	//
	return nil
}
