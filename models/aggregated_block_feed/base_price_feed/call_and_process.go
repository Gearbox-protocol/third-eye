package base_price_feed

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/artifacts/redstone"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg/priceFetcher"
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

var failedLatestRoundDataHandler = log.SendMsgIfCountMoreThan(time.Minute*30, 4) // if for the same feed, failed 4 times in 30 mins

// used in yearn and curve
func ParseQueryRoundData(returnData []byte, isPriceInUSD bool, feed string, blockNum int64) *schemas.PriceFeed {
	priceFeedABI := core.GetAbi("PriceFeed")
	roundData := schemas.LatestRounData{}
	value, err := priceFeedABI.Unpack("latestRoundData", returnData)
	if err != nil {
		if !utils.Contains([]string{"0x7B7C81748f311Cf3B9dfe90Ec7F23e9F06813323", "0x2E65c16Fe6CFd0519Ae1F80448FCa0E0B07c1911"}, feed) { // only for curve
			msg := fmt.Sprintf("For feed(%s) can't get the latestRounData: %s at %d", feed, err, blockNum)
			failedLatestRoundDataHandler(feed, msg)
		}
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
		RoundId:     roundData.RoundId.Int64(),
		PriceBI:     (*core.BigInt)(roundData.Answer),
		Price:       utils.GetFloat64Decimal(roundData.Answer, decimals),
		BlockNumber: blockNum,
	}
}

func (mdl *BasePriceFeed) ProcessResult(blockNum int64, results []multicall.Multicall2Result, _ string, force ...bool) *schemas.PriceFeed {
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
func (mdl *BasePriceFeed) GetRedStoneUnderlyings() []string {
	return mdl.DetailsDS.Underlyings
}
func (mdl *BasePriceFeed) GetUnderlyingCalls(blockNum int64) (calls []multicall.Multicall2Call, isQueryable bool) {
	updateABI := core.GetAbi("UpdatePriceFeed")
	for _, entry := range mdl.DetailsDS.Underlyings {
		contract, err := redstone.NewRedstone(common.HexToAddress(entry), mdl.Client)
		log.CheckFatal(err)
		var tokenDetails *core.RedStonePF
		if _, ok := mdl.DetailsDS.Info[entry]; ok {
			tokenDetails = mdl.DetailsDS.Info[entry]
			mdl.DetailsDS.FetchedInfo = true // if the redstone info wasn't there for feed, then we were still calling dataFeedId even though we tried to fetch that info before.
			// to prevent this we added fetchedInfo , this is just to provide the consistency for fetchedInfo as for feed where the info is already set, fetchedInfo should be true
		} else if !mdl.DetailsDS.FetchedInfo {
			if _, err := contract.DataFeedId(nil); err == nil {
				feedToken, signThreshold, dataId := priceFetcher.RedstoneDetails(common.HexToAddress(entry), mdl.Client)
				//
				tokenDetails = &core.RedStonePF{
					Type:             15,
					DataServiceId:    "redstone-primary-prod",
					DataId:           dataId,
					SignersThreshold: signThreshold,
					UnderlyingToken:  feedToken,
					Feed:             common.HexToAddress(mdl.Address),
				}
				mdl.DetailsDS.Info[entry] = tokenDetails
			}
			mdl.DetailsDS.FetchedInfo = true
		}
		if tokenDetails != nil {
			pod := mdl.Repo.GetRedStonemgr().GetPodSignWithRedstoneToken(int64(mdl.Repo.SetAndGetBlock(blockNum).Timestamp), *tokenDetails)
			// log.Info(pod.CallData)
			update, err := updateABI.Pack("updatePrice", pod.CallData)
			log.CheckFatal(err)
			calls = append(calls, multicall.Multicall2Call{
				Target:   common.HexToAddress(entry),
				CallData: update,
			})
		}
	}
	b, err := hex.DecodeString("feaf968c") // // latestRounData
	log.CheckFatal(err)
	calls = append(calls, multicall.Multicall2Call{
		Target:   common.HexToAddress(mdl.Address),
		CallData: b,
	})
	return calls, true
}
