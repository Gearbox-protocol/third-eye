package aggregated_block_feed

import (
	"sort"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/common"
)

// INSTRUCTIONS: TO FIX
//
// - UniTokenLastSync is not updated, as AddLastSyncForUniToken is commented
// - AddUniPools is not getting called as in chainlinkPriceFeed removed link to repo method.AddUniPoolsForToken
//
type PriceOnUNIFetcher struct {
	UniswapPools      []string
	UniPoolByToken    map[string]*schemas.UniswapPools
	UniPricesByTokens map[string]schemas.SortedUniPoolPrices
	// uniswap price related data structures
	tokenInfos map[string]*schemas.Token
	//
	UniTokenLastSync map[string]int64
}

func NewPriceOnUNIFetcher() *PriceOnUNIFetcher {
	return &PriceOnUNIFetcher{
		UniPoolByToken:    map[string]*schemas.UniswapPools{},
		UniPricesByTokens: map[string]schemas.SortedUniPoolPrices{},
		tokenInfos:        map[string]*schemas.Token{},
		UniTokenLastSync:  map[string]int64{},
	}
}

func (mdl *PriceOnUNIFetcher) AddUniPools(token *schemas.Token, uniswapPools *schemas.UniswapPools) {
	if mdl.UniPoolByToken[uniswapPools.Token] == nil {
		mdl.UniPoolByToken[uniswapPools.Token] = uniswapPools
	}
	mdl.tokenInfos[token.Address] = token
}

func (mdl *PriceOnUNIFetcher) GetUniswapPools() (updatedPools []*schemas.UniswapPools) {
	for _, entry := range mdl.UniPoolByToken {
		if entry.Updated {
			updatedPools = append(updatedPools, entry)
		}
		entry.Updated = false
	}
	return
}

func (mdl *PriceOnUNIFetcher) Clear() {
	mdl.UniPricesByTokens = map[string]schemas.SortedUniPoolPrices{}
}

// for getting the uniswap prices for chainlink token/usdc uniswap pairs.
// func (mdl *PriceOnUNIFetcher) AddLastSyncForUniToken(token string, lastSync int64) {
// 	// there is new oracle/feed added for a token
// 	if mdl.UniTokenLastSync[token] == 0 {
// 		mdl.UniTokenLastSync[token] = lastSync
// 	}
// 	mdl.UniTokenLastSync[token] = utils.Min(mdl.UniTokenLastSync[token], lastSync)
// }

func (mdl *PriceOnUNIFetcher) getUniswapPoolCalls(blockNum int64, whatToQuery string) (calls []multicall.Multicall2Call, tokens []string) {
	v2ABI := core.GetAbi("Uniswapv2Pool")
	v3ABI := core.GetAbi("Uniswapv3Pool")
	for token, pools := range mdl.UniPoolByToken {
		if whatToQuery != "all" && whatToQuery != token {
			continue
		}
		// only sync uniswap pool price for token that have last sync
		if mdl.UniTokenLastSync[token] >= blockNum {
			continue
		}
		uniswapv2Price, err := v2ABI.Pack("getReserves")
		log.CheckFatal(err)
		calls = append(calls, multicall.Multicall2Call{
			Target:   common.HexToAddress(pools.V2),
			CallData: uniswapv2Price,
		})
		uniswapv3Price, err := v3ABI.Pack("slot0")
		log.CheckFatal(err)
		calls = append(calls, multicall.Multicall2Call{
			Target:   common.HexToAddress(pools.V3),
			CallData: uniswapv3Price,
		})
		uniswapv3Twap, err := v3ABI.Pack("observe", []uint32{0, 600})
		log.CheckFatal(err)
		calls = append(calls, multicall.Multicall2Call{
			Target:   common.HexToAddress(pools.V3),
			CallData: uniswapv3Twap,
		})
		tokens = append(tokens, token)
	}
	return
}

// called on next level in the adapter kit
// so mu is not required as write operation is not performed at that levelAggre
// func (mdl PriceOnUNIFetcher) GetUniPricesByToken(token string) []*schemas.UniPoolPrices {
// 	return mdl.UniPricesByTokens[token]
// }

func (mdl *PriceOnUNIFetcher) sortUniPrices() {
	// uni prices
	for _, prices := range mdl.UniPricesByTokens {
		sort.Sort(prices)
	}
}
