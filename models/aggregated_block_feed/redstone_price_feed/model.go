package redstone_price_feed

import (
	"encoding/hex"
	"math/big"
	"time"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg/priceFetcher"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/aggregated_block_feed/base_price_feed"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type RedstonePriceFeed struct {
	*base_price_feed.BasePriceFeed
}

func NewRedstonePriceFeed(token, oracle string, pfType string, discoveredAt int64, client core.ClientI, repo ds.RepositoryI, pfVersion schemas.PFVersion) *RedstonePriceFeed {
	adapter := base_price_feed.NewBasePriceFeed(token, oracle, pfType, discoveredAt, client, repo, pfVersion)
	return NewRedstonePriceFeedFromAdapter(adapter.SyncAdapter)
}

func (feed RedstonePriceFeed) GetRedstonePF() []*core.RedStonePF {
	return []*core.RedStonePF{
		feed.DetailsDS.Info[feed.GetAddress()],
	}
}

func NewRedstonePriceFeedFromAdapter(adapter *ds.SyncAdapter) *RedstonePriceFeed {
	obj := &RedstonePriceFeed{
		BasePriceFeed: base_price_feed.NewBasePriceFeedFromAdapter(adapter),
	}
	if obj.DetailsDS.Info[adapter.GetAddress()] == nil {
		feedToken, signThreshold, dataId := priceFetcher.RedstoneDetails(common.HexToAddress(adapter.GetAddress()), adapter.Client)
		//
		tokenDetails := &core.RedStonePF{
			Type:             15,
			DataServiceId:    "redstone-primary-prod",
			DataId:           dataId,
			SignersThreshold: signThreshold,
			UnderlyingToken:  feedToken,
		}
		obj.DetailsDS.Info[adapter.GetAddress()] = tokenDetails
	}
	return obj
}

func (obj *RedstonePriceFeed) GetCalls(blockNum int64) (calls []multicall.Multicall2Call, isQueryable bool) {
	data, _ := hex.DecodeString("feaf968c") // latestRounData
	return []multicall.Multicall2Call{
		{
			Target:   common.HexToAddress(obj.Address),
			CallData: data,
		},
	}, true
}

func (mdl *RedstonePriceFeed) ProcessResult(blockNum int64, results []multicall.Multicall2Result, force ...bool) *schemas.PriceFeed {
	validTokens := mdl.TokensValidAtBlock(blockNum)
	isPriceInUSD := mdl.GetVersion().IsPriceInUSD()
	{
		if len(results) != 1 {
			log.Fatal("wrong result")
		}
		if results[0].Success {
			value, err := core.GetAbi("YearnPriceFeed").Unpack("latestRoundData", results[0].ReturnData)
			log.CheckFatal(err)
			price := *abi.ConvertType(value[1], new(*big.Int)).(**big.Int)
			// log.Info("onchain price found for ", mdl.Address, "at", blockNum, price)
			return parsePriceForRedStone(price, isPriceInUSD)
		} else if time.Since(time.Unix(int64(mdl.Repo.SetAndGetBlock(blockNum).Timestamp), 0)) < time.Hour {
			// } else {
			// 	if (len(force) ==0 || !force[0] ) {
			// 		return nil
			// 	}
		}
	}
	{
		//
		priceBI := mdl.Repo.GetRedStonemgr().GetPrice(int64(mdl.Repo.SetAndGetBlock(blockNum).Timestamp), *mdl.DetailsDS.Info[mdl.GetAddress()])
		if priceBI.Cmp(new(big.Int)) == 0 {
			log.Warnf("RedStone price for %s at %d is %f", mdl.Repo.GetToken(validTokens[0].Token).Symbol, blockNum, priceBI)
			return nil
		}

		priceData := parsePriceForRedStone(priceBI, isPriceInUSD)
		log.Infof("RedStone price for %s at %d is %f", mdl.Repo.GetToken(validTokens[0].Token).Symbol, blockNum, priceData.Price)
		//
		return priceData
	}
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
