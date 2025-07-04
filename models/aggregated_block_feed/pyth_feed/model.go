package pyth_feed

import (
	"encoding/hex"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/aggregated_block_feed/base_price_feed"
	"github.com/ethereum/go-ethereum/common"
)

type PythPriceFeed struct {
	*base_price_feed.BasePriceFeed
}

func NewPythPriceFeed(token, oracle string, pfType string, discoveredAt int64, client core.ClientI, repo ds.RepositoryI, version core.VersionType, underlyings []string) *PythPriceFeed {
	adapter := base_price_feed.NewBasePriceFeed(token, oracle, pfType, discoveredAt, client, repo, version, underlyings)
	return NewPythPriceFeedFromAdapter(adapter.SyncAdapter)
}

func NewPythPriceFeedFromAdapter(adapter *ds.SyncAdapter) *PythPriceFeed {
	return &PythPriceFeed{
		BasePriceFeed: base_price_feed.NewBasePriceFeedFromAdapter(adapter),
	}
}

func (feed PythPriceFeed) GetCalls(blockNum int64) (calls []multicall.Multicall2Call, isQueryable bool) {
	data, _ := hex.DecodeString("feaf968c") // latestRounData
	return []multicall.Multicall2Call{
		{
			Target:   common.HexToAddress(feed.Address),
			CallData: data,
		},
	}, true // pyth price feed is queryable
}

func (adapter *PythPriceFeed) ProcessResult(blockNum int64, results []multicall.Multicall2Result, _ string, force ...bool) *schemas.PriceFeed {
	result := results[len(results)-1]
	if !result.Success {
		ts := adapter.Repo.SetAndGetBlock(blockNum).Timestamp
		obj, err := pkg.GetPrice(adapter.DetailsDS.Underlyings[0], int64(ts))
		if err != nil {
			log.Fatal("Pyth price feed", adapter.GetAddress(), " failed at block: ", blockNum, err)
		}
		return &schemas.PriceFeed{
			RoundId:     0,
			PriceBI:     obj.Price,
			Price:       obj.F,
			BlockNumber: blockNum,
		}
	}
	isPriceInUSD := adapter.GetVersion().IsPriceInUSD()
	return base_price_feed.ParseQueryRoundData(result.ReturnData, isPriceInUSD, adapter.GetAddress(), blockNum)
}
