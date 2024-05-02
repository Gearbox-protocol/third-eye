package curve_price_feed

import (
	"math/big"
	"time"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/aggregated_block_feed/base_price_feed"
	"github.com/ethereum/go-ethereum/common"
)

type CurvePriceFeed struct {
	*base_price_feed.BasePriceFeed
}

func NewCurvePriceFeed(token, oracle string, pfType string, discoveredAt int64, client core.ClientI, repo ds.RepositoryI, pfVersion schemas.PFVersion) *CurvePriceFeed {
	adapter := base_price_feed.NewBasePriceFeed(token, oracle, pfType, discoveredAt, client, repo, pfVersion)
	return NewCurvePriceFeedFromAdapter(adapter.SyncAdapter)
}

func NewCurvePriceFeedFromAdapter(adapter *ds.SyncAdapter) *CurvePriceFeed {
	return &CurvePriceFeed{
		BasePriceFeed: base_price_feed.NewBasePriceFeedFromAdapter(adapter),
	}
}

// same as query price feed
// func (*CurvePriceFeed) GetCalls(blockNum int64) (calls []multicall.Multicall2Call, isQueryable bool) {

var curvePFLatestRoundDataTimer = map[string]log.TimerFn{}

func (adapter *CurvePriceFeed) ProcessResult(blockNum int64, results []multicall.Multicall2Result) *schemas.PriceFeed {
	if !results[0].Success {
		if adapter.GetVersion().LessThan(core.NewVersion(300)) { // failed and
			// if virtualprice of pool for this oracle is not within lowerBound and upperBound , ignore the price
			oracleAddr := common.HexToAddress(adapter.GetAddress())
			virtualPrice := GetCurveVirtualPrice(blockNum, oracleAddr, adapter.GetVersion(), adapter.Client)
			//
			withinLimits := func() bool {
				lowerLimit, err := core.CallFuncWithExtraBytes(adapter.Client, "a384d6ff", oracleAddr, blockNum, nil) // lowerBound
				log.CheckFatal(err)
				upperLimit, err := core.CallFuncWithExtraBytes(adapter.Client, "b09ad8a0", oracleAddr, blockNum, nil) // upperBound
				log.CheckFatal(err)
				return new(big.Int).SetBytes(lowerLimit).Cmp(virtualPrice) < 0 &&
					new(big.Int).SetBytes(upperLimit).Cmp(virtualPrice) > 0
			}()
			if curvePFLatestRoundDataTimer[adapter.GetAddress()] == nil {
				curvePFLatestRoundDataTimer[adapter.GetAddress()] = log.GetRiskMsgTimer()
			}
			var msg string
			if !withinLimits {
				msg = "virtual price is not within limits for " + adapter.GetAddress()
			} else {
				msg = "failing due to unknown reason maybe underlying pricefeed of curve pool token is failing for curve adapter" + adapter.GetAddress()
			}
			log.SendRiskAlertPerTimer(
				log.RiskAlert{
					Msg: msg,
					RiskHeader: log.RiskHeader{
						BlockNumber: blockNum,
						EventCode:   "CURVE_LATEST_ROUNDDATA_FAIL",
					},
				},
				curvePFLatestRoundDataTimer[adapter.GetAddress()],
				86400*time.Second,
			)
			return nil
		} else if core.GetChainId(adapter.Client) == 7878 {
			return nil
		}
	}
	isPriceInUSD := adapter.GetVersion().IsPriceInUSD()
	return base_price_feed.ParseQueryRoundData(results[0].ReturnData, isPriceInUSD, adapter.GetAddress(), blockNum)
}

func GetCurveVirtualPrice(blockNum int64, oracleAddr common.Address, version core.VersionType, client core.ClientI) *big.Int {
	curvePool := func() common.Address {
		if !version.MoreThanEq(core.NewVersion(300)) {
			curvePoolBytes, err := core.CallFuncWithExtraBytes(client, "218751b2", oracleAddr, blockNum, nil) // curvePool from curvev1Adapter abi
			log.CheckFatal(err)
			return common.BytesToAddress(curvePoolBytes)
		} else {
			// LPCONTRACT_LOGIC
			lpCOntractBytes, err := core.CallFuncWithExtraBytes(client, "8acee3cf", oracleAddr, blockNum, nil) // lpContract
			log.CheckFatal(err)
			return common.BytesToAddress(lpCOntractBytes)
		}
	}()
	virtualPrice, err := core.CallFuncWithExtraBytes(client, "bb7b8b80", curvePool, blockNum, nil) // getVirtualPrice
	log.CheckFatal(err)
	return new(big.Int).SetBytes(virtualPrice)
}
