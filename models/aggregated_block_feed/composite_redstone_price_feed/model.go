package composite_redstone_price_feed

import (
	"math/big"
	"time"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/aggregated_block_feed/base_price_feed"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type CompositeRedStonePriceFeed struct {
	*base_price_feed.BasePriceFeed
	priceFeed0 common.Address
	priceFeed1 common.Address
	Decimals   int8
}

func NewRedstonePriceFeed(token, oracle string, pfType string, discoveredAt int64, client core.ClientI, repo ds.RepositoryI, pfVersion schemas.PFVersion) *CompositeRedStonePriceFeed {
	adapter := base_price_feed.NewBasePriceFeed(token, oracle, pfType, discoveredAt, client, repo, pfVersion)
	return NewRedstonePriceFeedFromAdapter(adapter.SyncAdapter)
}

func NewRedstonePriceFeedFromAdapter(adapter *ds.SyncAdapter) *CompositeRedStonePriceFeed {
	pf0, err := core.CallFuncWithExtraBytes(adapter.Client, "385aee1b", common.HexToAddress(adapter.Address), 0, nil) // priceFeed0
	log.CheckFatal(err)
	pf1, err := core.CallFuncWithExtraBytes(adapter.Client, "ab0ca0e1", common.HexToAddress(adapter.Address), 0, nil) // priceFeed1
	log.CheckFatal(err)
	//
	decimals, err := core.CallFuncWithExtraBytes(adapter.Client, "313ce567", common.BytesToAddress(pf0), 0, nil) // decimals
	log.CheckFatal(err)
	return &CompositeRedStonePriceFeed{
		BasePriceFeed: base_price_feed.NewBasePriceFeedFromAdapter(adapter),
		priceFeed0:    common.BytesToAddress(pf0),
		priceFeed1:    common.BytesToAddress(pf1),
		Decimals:      int8(new(big.Int).SetBytes(decimals).Int64()),
	}
}

func (mdl *CompositeRedStonePriceFeed) GetCalls(blockNum int64) (calls []multicall.Multicall2Call, isQueryable bool) {
	priceFeedABI := core.GetAbi("PriceFeed")
	data, err := priceFeedABI.Pack("latestRoundData")
	log.CheckFatal(err)
	return []multicall.Multicall2Call{{
		Target:   common.HexToAddress(mdl.Address),
		CallData: data,
	},{
		Target:   mdl.priceFeed1,
		CallData: data,
	}}, true
}

func (mdl *CompositeRedStonePriceFeed) ProcessResult(blockNum int64, results []multicall.Multicall2Result, force ...bool) *schemas.PriceFeed {
	if !results[1].Success {
		return nil
	}
	isPriceInUSD := mdl.GetVersion().IsPriceInUSD() // should be always true
	{
		if results[0].Success {
			value, err := core.GetAbi("YearnPriceFeed").Unpack("latestRoundData", results[0].ReturnData)
			log.CheckFatal(err)
			price :=  *abi.ConvertType(value[1], new(*big.Int)).(**big.Int)
			log.Info("onchain price found for ", mdl.Address, "at", blockNum, price)
			return parsePriceForRedStone(price, isPriceInUSD)
		} else if time.Since(time.Unix(int64(mdl.Repo.SetAndGetBlock(blockNum).Timestamp),0)) > time.Hour {
			if (len(force) ==0 || !force[0] ) {
				return nil
			}
		}
	}
	validTokens := mdl.TokensValidAtBlock(blockNum)
	// log.Info(mdl.Repo.SetAndGetBlock(blockNum).Timestamp, validTokens, utils.ToJson(mdl.DetailsDS))
	targetPrice := mdl.Repo.GetRedStonemgr().GetPrice(int64(mdl.Repo.SetAndGetBlock(blockNum).Timestamp), validTokens[0].Token, true)
	if targetPrice.Cmp(new(big.Int)) == 0 {
		log.Warnf("RedStone composite targetprice for %s at %d is %f", mdl.Repo.GetToken(validTokens[0].Token).Symbol, blockNum, targetPrice)
		return nil
	}
	//
	basePrice := func() *big.Int {
		values, err := core.GetAbi("PriceFeed").Unpack("latestRoundData", results[1].ReturnData)
		if err != nil {
			log.Warnf("Can't get the lastestRounData: %s at %d for mdl.priceFeed1(%s)", err, blockNum, mdl.priceFeed1)
			return nil
		}
		return *abi.ConvertType(values[1], new(*big.Int)).(**big.Int)
	}()
	log.Infof("RedStone composite targetprice for %s at %d is %f, basePrice, %s", mdl.Repo.GetToken(validTokens[0].Token).Symbol, blockNum, utils.GetFloat64Decimal(targetPrice, mdl.Decimals), basePrice)
	if basePrice == nil {
		return nil
	}
	// calculate price
	price := utils.GetInt64(new(big.Int).Mul(targetPrice, basePrice), mdl.Decimals)
	return parsePriceForRedStone(price, isPriceInUSD)
}

func parsePriceForRedStone(price *big.Int, isPriceInUSD bool) *schemas.PriceFeed {
	var decimals int8 = 18 // for eth
	if isPriceInUSD {
		decimals = 8 // for usd
	}
	return &schemas.PriceFeed{
		RoundId: 0,
		PriceBI: (*core.BigInt)(price),
		Price:   utils.GetFloat64Decimal(price, decimals),
	}
}
