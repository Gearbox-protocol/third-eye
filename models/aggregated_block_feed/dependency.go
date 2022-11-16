package aggregated_block_feed

import (
	"context"
	"sync"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/common"
)

// ignoring dependencies for v1
type QueryPFDependencies struct {
	// chainlink symbol to dependent query pf
	ChainlinkSymToQueryPFSyms map[string][]string
	// chainlink symbol updated at these blocks, we need to fetch price for dependent query pf at these block numbers
	ChainlinkSymToUpdatedBlocks map[string][]int64
	// yearn/curve symbol to chainlink dependency symbol
	depGraph map[string][]string
	// blocks to remove
	aggregatedFetchedBlocks []int64
	// yearn feed prices to be ordered
	depBasedExtraPrices []*schemas.PriceFeed
	//
	repo   ds.RepositoryI
	mu     *sync.Mutex
	client core.ClientI
}

func NewQueryPFDepenencies(repo ds.RepositoryI, client core.ClientI) *QueryPFDependencies {
	chainId, err := client.ChainID(context.TODO())
	log.CheckFatal(err)
	depGraph := getDepGraph(chainId.Int64())
	return &QueryPFDependencies{
		depGraph:                    depGraph,
		ChainlinkSymToQueryPFSyms:   getInvertDependencyGraph(depGraph),
		ChainlinkSymToUpdatedBlocks: map[string][]int64{},
		//
		repo:   repo,
		mu:     &sync.Mutex{},
		client: client,
	}
}

func (q *QueryPFDependencies) ChainlinkPriceUpdatedAt(token string, blockNums []int64) {
	chainlinkSym := q.repo.GetToken(token).Symbol
	q.ChainlinkSymToUpdatedBlocks[chainlinkSym] = blockNums

}

func (q *QueryPFDependencies) getChainlinkBasedQueryUpdates() map[int64]map[string]bool {
	//  blockNum to QueryPFToken
	updates := map[int64]map[string]bool{}
	var updatedChainlinkSym []string
	for chainlinkSym, blockNums := range q.ChainlinkSymToUpdatedBlocks {
		updatedChainlinkSym = append(updatedChainlinkSym, chainlinkSym)
		//
		for _, dependentSym := range q.ChainlinkSymToQueryPFSyms[chainlinkSym] {
			depAddr := q.repo.GetAddressBySymbol(dependentSym)
			for _, blockNum := range blockNums {
				if updates[blockNum] == nil {
					updates[blockNum] = map[string]bool{}
				}
				updates[blockNum][depAddr] = true
			}
		}
	}
	if updatedChainlinkSym != nil {
		log.Info("Updated chainlinks are:", updatedChainlinkSym)
	}
	q.ChainlinkSymToUpdatedBlocks = map[string][]int64{}
	return updates
}

func getDepGraph(chainId int64) map[string][]string {
	depGraph := map[string][]string{
		// frax and curve
		"yvCurve-FRAX":     {"FRAX", "USDC", "USDT", "DAI"},
		"FRAX3CRV-f":       {"FRAX", "USDC", "USDT", "DAI"},
		"stkcvxFRAX3CRV-f": {"FRAX", "USDC", "USDT", "DAI"},
		"cvxFRAX3CRV-f":    {"FRAX", "USDC", "USDT", "DAI"},
		// frax and usdc
		"crvFRAX":       {"USDC", "FRAX"},
		"stkcvxcrvFRAX": {"USDC", "FRAX"},

		// yearn
		"yvWBTC": {"WBTC"},
		"yvWETH": {"WETH"},
		"yvDAI":  {"DAI"},
		"yvUSDC": {"USDC"},
		"wstETH": {"stETH"},

		// diesel tokens
		// "dUSDC":   {"USDC"},
		// "dDAI":    {"DAI"},
		// "dWETH":   {"ETH"},
		// "dwstETH": {"stETH"},
		// "dWBTC":   {"WBTC"},
		// 3 crv
		"3Crv":       {"USDC", "USDT", "DAI"},
		"stkcvx3Crv": {"USDC", "USDT", "DAI"},
		"cvx3Crv":    {"USDC", "USDT", "DAI"},
		// lusd and 3crv
		"LUSD3CRV-f":       {"LUSD", "DAI", "USDC", "USDT"},
		"stkcvxLUSD3CRV-f": {"LUSD", "DAI", "USDC", "USDT"},
		"cvxLUSD3CRV-f":    {"LUSD", "DAI", "USDC", "USDT"},
		// susd and 3crv
		"cvxcrvPlain3andSUSD":    {"SUSD", "DAI", "USDC", "USDT"},
		"stkcvxcrvPlain3andSUSD": {"SUSD", "DAI", "USDC", "USDT"},
		"crvPlain3andSUSD":       {"SUSD", "DAI", "USDC", "USDT"},

		// gusd and 3crv
		"stkcvxgusd3CRV": {"GUSD", "DAI", "USDC", "USDT"},
		"cvxgusd3CRV":    {"GUSD", "DAI", "USDC", "USDT"},
		// steth/eth
		"steCRV":        {"stETH", "ETH"},
		"yvCurve-stETH": {"stETH", "ETH"},
		"cvxsteCRV":     {"stETH", "ETH"},
	}
	if chainId == 5 {
		for _, token := range []string{"yvWBTC", "yvCurve-FRAX", "yvDAI", "yvCurve-stETH", "yvUSDC"} {
			depGraph["Yearn "+token] = depGraph[token]
			delete(depGraph, token)
		}
		delete(depGraph, "yvWETH")
	}
	return depGraph
}

func getInvertDependencyGraph(depGraph map[string][]string) map[string][]string {
	invertedGraph := map[string][]string{}
	for token, deps := range depGraph {
		for _, chainlinkSym := range deps {
			invertedGraph[chainlinkSym] = append(invertedGraph[chainlinkSym], token)
		}
	}
	return invertedGraph
}

func (q *QueryPFDependencies) checkInDepGraph(token, oracle string, blockNum int64) {
	depQueryPFSym := q.repo.GetToken(token).Symbol
	if q.depGraph[depQueryPFSym] == nil {
		log.Fatalf("Dep for query based price feed(%s) not found for token(%s) at %d", oracle, depQueryPFSym, blockNum)
	}
}

func (q *QueryPFDependencies) updateQueryPrices(pfs []*schemas.PriceFeed) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.depBasedExtraPrices = append(q.depBasedExtraPrices, pfs...)
}

func (q *QueryPFDependencies) extraPriceForQueryFeed() []*schemas.PriceFeed {
	updates := q.getChainlinkBasedQueryUpdates()
	for _, blockToDelete := range q.aggregatedFetchedBlocks {
		delete(updates, blockToDelete)
	}
	ch := make(chan int, 4)
	wg := &sync.WaitGroup{}
	for blockNum, tokens := range updates {
		wg.Add(1)
		ch <- 1
		go q.fetchRoundData(blockNum, tokens, ch, wg)
	}
	wg.Wait()
	pfs := q.depBasedExtraPrices
	q.depBasedExtraPrices = nil
	log.Info(len(pfs), "extra price feed fetched due to chainlink updates")
	return pfs
}

func (q *QueryPFDependencies) fetchRoundData(blockNum int64, tokens map[string]bool,
	ch <-chan int, wg *sync.WaitGroup) {
	priceFeedABI := core.GetAbi("PriceFeed")
	var calls []multicall.Multicall2Call

	data, err := priceFeedABI.Pack("latestRoundData")
	log.CheckFatal(err)
	var tokenOracles []*schemas.TokenOracle
	for token := range tokens {
		details := q.repo.GetOracleForV2Token(token)
		tokenOracles = append(tokenOracles, details)
		call := multicall.Multicall2Call{
			Target:   common.HexToAddress(details.Feed),
			CallData: data,
		}
		calls = append(calls, call)
	}
	results := core.MakeMultiCall(q.client, blockNum, false, calls, 30)
	var newPrices []*schemas.PriceFeed
	for ind, entry := range results {
		details := tokenOracles[ind]
		var newPrice *schemas.PriceFeed
		/// parse price
		if entry.Success {
			newPrice = parseRoundData(entry.ReturnData, true) // only valid for v2
		} else {
			// if failed check if pfType of the queryPrice is YearnPF
			adapterI := q.repo.GetKit().GetAdapter(tokenOracles[ind].Feed)
			adapter := adapterI.(*QueryPriceFeed)
			switch adapter.GetDetailsByKey("pfType") {
			case ds.YearnPF:
				newPrice = adapter.calculateYearnPFInternally(blockNum)
			}
		}
		// add token and feed details
		newPrice.Token = details.Token
		newPrice.Feed = details.Feed
		newPrice.BlockNumber = blockNum
		newPrices = append(newPrices, newPrice)
	}
	q.updateQueryPrices(newPrices)
	// sync control
	<-ch
	wg.Done()
}
