package base_price_feed

import (
	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/third-eye/ds"
)

type QueryPriceFeedI interface {
	TokensValidAtBlock(blockNum int64) []schemas.TokenAndMergedPFVersion
	GetPFType() string
	ds.SyncAdapterI
	GetCalls(blockNum int64) (calls []multicall.Multicall2Call, isQueryable bool)
	ProcessResult(blockNum int64, results []multicall.Multicall2Result) *schemas.PriceFeed
	DisableToken(token string, disabledAt int64, pfVersion schemas.PFVersion)
	AddToken(token string, discoveredAt int64, pfVersion schemas.PFVersion)
	GetTokens() map[string]map[schemas.PFVersion][]int64
}
