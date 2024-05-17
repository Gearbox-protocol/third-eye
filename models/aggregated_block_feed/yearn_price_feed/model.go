package yearn_price_feed

import (
	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/aggregated_block_feed/base_price_feed"
	"github.com/ethereum/go-ethereum/common"
)

type YearnPriceFeed struct {
	*base_price_feed.BasePriceFeed
	yearnPFInternal
}

func NewYearnPriceFeed(token, oracle string, pfType string, discoveredAt int64, client core.ClientI, repo ds.RepositoryI, pfVersion schemas.PFVersion) *YearnPriceFeed {
	adapter := base_price_feed.NewBasePriceFeed(token, oracle, pfType, discoveredAt, client, repo, pfVersion)
	return NewYearnPriceFeedFromAdapter(adapter.SyncAdapter)
}

func NewYearnPriceFeedFromAdapter(adapter *ds.SyncAdapter) *YearnPriceFeed {
	return &YearnPriceFeed{
		BasePriceFeed: base_price_feed.NewBasePriceFeedFromAdapter(adapter),
		yearnPFInternal: yearnPFInternal{
			mainPFAddress: common.HexToAddress(adapter.Address),
			version:       adapter.GetVersion(),
		},
	}
}

// same as query price feed
// func (*YearnPriceFeed) GetCalls(blockNum int64) (calls []multicall.Multicall2Call, isQueryable bool) {

func (mdl *YearnPriceFeed) ProcessResult(blockNum int64, results []multicall.Multicall2Result) *schemas.PriceFeed {
	if !results[0].Success {
		if mdl.Address == "0x628539959F3B3bb0cFe2102dCaa659cf1E8D19EB" { // https://optimistic.etherscan.io/address/0x628539959F3B3bb0cFe2102dCaa659cf1E8D19EB // yvWETH, v3
			return nil
		}
		//
		if !mdl.GetVersion().MoreThanEq(core.NewVersion(300)) { // v1,v2
			priceData, err := mdl.CalculateYearnPFInternally(blockNum)
			if err != nil {
				log.Fatalf("at %d can't calculate yearnfeed(%s)'s price internally: %s",
					blockNum,
					mdl.GetAddress(), err.Error())
			}
			log.Warnf("Can't get latestRounData for YearnModule in AQFWrapper for %s(%s) at %d",
				mdl.GetDetailsByKey("pfType"), mdl.GetAddress(), blockNum)
			return priceData
		} else { // v3
			log.Warnf("Can't get latestRounData for YearnModule in AQFWrapper for %s(%s) at %d",
				mdl.GetDetailsByKey("pfType"), mdl.GetAddress(), blockNum)
			return nil
		}
	}
	isPriceInUSD := mdl.GetVersion().IsPriceInUSD()
	return base_price_feed.ParseQueryRoundData(results[0].ReturnData, isPriceInUSD, mdl.GetAddress(), blockNum)
}
