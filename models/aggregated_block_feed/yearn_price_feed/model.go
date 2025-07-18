package yearn_price_feed

import (
	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/aggregated_block_feed/base_price_feed"
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm/utils"
)

type YearnPriceFeed struct {
	*base_price_feed.BasePriceFeed
	yearnPFInternal
}

func newYearnPriceFeed(token, oracle string, pfType string, discoveredAt int64, client core.ClientI, repo ds.RepositoryI, version core.VersionType, underlyings []string) *YearnPriceFeed {
	adapter := base_price_feed.NewBasePriceFeed(token, oracle, pfType, discoveredAt, client, repo, version, underlyings)
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

func (mdl *YearnPriceFeed) ProcessResult(blockNum int64, results []multicall.Multicall2Result, _ string, force ...bool) *schemas.PriceFeed {
	if !results[0].Success {
		if utils.Contains([]string{
			"0x628539959F3B3bb0cFe2102dCaa659cf1E8D19EB",
			"0xf4bc1D894F23e85bF666dA647CA573fB13109811",
			"0x5e6ee42dD1D1A8299CB0aC4C7641F597C434aC5e",
			"0x228C64cA6ece0ECeB3593e9838996CD5851e3797",                 // reverse v300 feed failing as the underlying token price is not updated. redstone price feed on mainnet
			"0x77BB9d857a5FfB6Db01581887312F68B1C9832A6"}, mdl.Address) { // https://optimistic.etherscan.io/address/0x628539959F3B3bb0cFe2102dCaa659cf1E8D19EB // yvWETH, v3
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
