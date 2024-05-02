package redstone_price_feed

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/aggregated_block_feed/base_price_feed"
)

type RedstonePriceFeed struct {
	*base_price_feed.BasePriceFeed
}

func NewRedstonePriceFeed(token, oracle string, pfType string, discoveredAt int64, client core.ClientI, repo ds.RepositoryI, pfVersion schemas.PFVersion) *RedstonePriceFeed {
	adapter := base_price_feed.NewBasePriceFeed(token, oracle, pfType, discoveredAt, client, repo, pfVersion)
	return NewRedstonePriceFeedFromAdapter(adapter.SyncAdapter)
}

func NewRedstonePriceFeedFromAdapter(adapter *ds.SyncAdapter) *RedstonePriceFeed {
	return &RedstonePriceFeed{
		BasePriceFeed: base_price_feed.NewBasePriceFeedFromAdapter(adapter),
	}
}

func (*RedstonePriceFeed) GetCalls(blockNum int64) (calls []multicall.Multicall2Call, isQueryable bool) {
	return nil, true
}

func (mdl *RedstonePriceFeed) ProcessResult(blockNum int64, results []multicall.Multicall2Result) *schemas.PriceFeed {
	validTokens := mdl.TokensValidAtBlock(blockNum)
	priceBI := mdl.Repo.GetRedStonemgr().GetPrice(int64(mdl.Repo.SetAndGetBlock(blockNum).Timestamp), validTokens[0].Token, false)
	//
	isPriceInUSD := mdl.GetVersion().IsPriceInUSD() // should be always true

	priceData := parsePriceForRedStone(priceBI, isPriceInUSD)
	log.Infof("RedStone price for %s at %d is %f", mdl.Repo.GetToken(validTokens[0].Token).Symbol, blockNum, priceData.Price)
	//
	return priceData
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
