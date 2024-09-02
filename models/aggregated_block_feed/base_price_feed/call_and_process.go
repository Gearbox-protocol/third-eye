package base_price_feed

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func (mdl *BasePriceFeed) GetCalls(blockNum int64) (calls []multicall.Multicall2Call, isQueryable bool) {
	priceFeedABI := core.GetAbi("PriceFeed")
	// reduntant check already in the aqfwrapper, CHECK_RED_QUERY_ADAPTER
	// this check is not needed as in dependency based fetching, lastSync will be set to willsyncTo, and will be updated in the next block
	// if blockNum <= mdl.GetLastSync() || len(mdl.TokensValidAtBlock(blockNum)) == 0 {
	// 	return
	// }
	data, err := priceFeedABI.Pack("latestRoundData")
	log.CheckFatal(err)
	return []multicall.Multicall2Call{{
		Target:   common.HexToAddress(mdl.GetAddress()),
		CallData: data,
	}}, true
}

func ParseQueryRoundData(returnData []byte, isPriceInUSD bool, feed string, blockNum int64) *schemas.PriceFeed {
	priceFeedABI := core.GetAbi("PriceFeed")
	roundData := schemas.LatestRounData{}
	value, err := priceFeedABI.Unpack("latestRoundData", returnData)
	if err != nil {
		log.Warnf("For feed(%s) can't get the latestRounData: %s at %d", feed, err, blockNum)
		return nil
	}
	roundData.RoundId = *abi.ConvertType(value[0], new(*big.Int)).(**big.Int)
	roundData.Answer = *abi.ConvertType(value[1], new(*big.Int)).(**big.Int)
	// roundData.StartedAt = *abi.ConvertType(value[2], new(*big.Int)).(**big.Int)
	// roundData.UpdatedAt = *abi.ConvertType(value[3], new(*big.Int)).(**big.Int)
	// roundData.AnsweredInRound = *abi.ConvertType(value[4], new(*big.Int)).(**big.Int)
	var decimals int8 = 18 // for eth
	if isPriceInUSD {
		decimals = 8 // for usd
	}
	return &schemas.PriceFeed{
		RoundId: roundData.RoundId.Int64(),
		PriceBI: (*core.BigInt)(roundData.Answer),
		Price:   utils.GetFloat64Decimal(roundData.Answer, decimals),
	}
}

func (mdl *BasePriceFeed) ProcessResult(blockNum int64, results []multicall.Multicall2Result, force ...bool) *schemas.PriceFeed {
	isPriceInUSD := mdl.GetVersion().IsPriceInUSD()
	if !results[0].Success {
		if mdl.GetVersion().MoreThanEq(core.NewVersion(300)) {
			if core.GetChainId(mdl.Client) == 7878 {
				return nil
			} else if utils.Contains([]string{"0xCbeCfA4017965939805Da5a2150E3DB1BeDD0364", "0x814E6564e8cda436c1ab25041C10bfdb21dEC519"},
				mdl.GetAddress()) { // arbitrum redstone composite feed // reserve price feeds
				return nil
			} else {
				log.Warnf("Can't get latestRounData in AQFWrapper for %s(%s) at %d",
					mdl.GetDetailsByKey("pfType"), mdl.GetAddress(), blockNum)
				return nil
			}
		}
	}
	return ParseQueryRoundData(results[0].ReturnData, isPriceInUSD, mdl.GetAddress(), blockNum)
	//
}
