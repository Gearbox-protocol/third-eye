package aggregated_block_feed

import (
	"sync"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
)

type repoI interface {
	GetAdapter(addr string) ds.SyncAdapterI
	// if returned value is nil, it means that token oracle hasn't been added yet.
	GetTokens() []string
	GetToken(string) *schemas.Token
}

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
	aqf *AQFWrapper
	//
	repo repoI
	TokenSymMap
	mu     *sync.Mutex
	client core.ClientI
}

func NewQueryPFDepenencies(repo ds.RepositoryI, client core.ClientI) *QueryPFDependencies {
	chainId := core.GetChainId(client)
	depGraph := getDepGraph(chainId)
	return &QueryPFDependencies{
		depGraph:                    depGraph,
		ChainlinkSymToQueryPFSyms:   getInvertDependencyGraph(depGraph),
		ChainlinkSymToUpdatedBlocks: map[string][]int64{},
		//
		repo:        repo,
		mu:          &sync.Mutex{},
		TokenSymMap: newTokenSymMap(chainId),
		client:      client,
	}
}

func (q *QueryPFDependencies) chainlinkPriceUpdatedAt(token string, blockNums []int64) {
	q.updateIfTest(q.repo)
	chainlinkSym := q.getTokenSym(token)
	q.mu.Lock()
	defer q.mu.Unlock()
	q.ChainlinkSymToUpdatedBlocks[chainlinkSym] = blockNums
}

func (q *QueryPFDependencies) getChainlinkBasedQueryUpdates(clearExtraBefore int64) map[int64]map[string]bool {
	//  blockNum to QueryPFToken
	updates := map[int64]map[string]bool{}
	var updatedChainlinkSym []string
	for chainlinkSym, blockNums := range q.ChainlinkSymToUpdatedBlocks {
		updatedChainlinkSym = append(updatedChainlinkSym, chainlinkSym)
		//
		for _, dependentSym := range q.ChainlinkSymToQueryPFSyms[chainlinkSym] {
			depAddr := q.getTokenAddr(dependentSym)
			if depAddr == "" {
				continue
			}
			for _, blockNum := range blockNums {
				// if a new chainlink price oracle is added it will create initial pf entry for lower block number
				// and queryPFDependency will try to fetch dependent query token's pf for this lower block number.
				// but the last pf fetched for query token saved in db can have higher block number than this lower block number.
				if blockNum <= clearExtraBefore {
					continue
				}
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

// token to its dependencies
func getDepGraph(chainId int64) map[string][]string {
	depGraph := map[string][]string{
		// frax and curve
		"yvCURVE_FRAX":   {"FRAX", "USDC", "USDT", "DAI"},
		"FRAX3CRV":       {"FRAX", "USDC", "USDT", "DAI"},
		"stkcvxFRAX3CRV": {"FRAX", "USDC", "USDT", "DAI"},
		"cvxFRAX3CRV":    {"FRAX", "USDC", "USDT", "DAI"},
		// frax and usdc
		"crvFRAX":       {"USDC", "FRAX"},
		"cvxcrvFRAX":    {"USDC", "FRAX"},
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
		// "dWETH":   {"WETH"},
		// "dwstETH": {"stETH"},
		// "dWBTC":   {"WBTC"},
		// 3 crv
		"3CRV":       {"USDC", "USDT", "DAI"},
		"stkcvx3Crv": {"USDC", "USDT", "DAI"},
		"cvx3Crv":    {"USDC", "USDT", "DAI"},
		// lusd and 3crv
		"LUSD3CRV":       {"LUSD", "DAI", "USDC", "USDT"},
		"stkcvxLUSD3CRV": {"LUSD", "DAI", "USDC", "USDT"},
		"cvxLUSD3CRV":    {"LUSD", "DAI", "USDC", "USDT"},
		// susd and 3crv
		"cvxcrvPlain3andSUSD":    {"SUSD", "DAI", "USDC", "USDT"},
		"stkcvxcrvPlain3andSUSD": {"SUSD", "DAI", "USDC", "USDT"},
		"crvPlain3andSUSD":       {"SUSD", "DAI", "USDC", "USDT"},

		// gusd and 3crv
		"stkcvxgusd3CRV": {"GUSD", "DAI", "USDC", "USDT"},
		"cvxgusd3CRV":    {"GUSD", "DAI", "USDC", "USDT"},
		"GUSD3CRV":       {"GUSD", "DAI", "USDC", "USDT"},
		// steth/eth
		"stkcvxsteCRV":  {"stETH", "WETH"}, // phantom convex on mainnet
		"steCRV":        {"stETH", "WETH"}, // curve steth
		"yvCurve_stETH": {"stETH", "WETH"}, // yearn
		"cvxsteCRV":     {"stETH", "WETH"}, // convex token for steth
		//
		// new v3 pools
		"OHMFRAXBP":       {"OHM", "FRAX", "USDC"},
		"cvxOHMFRAXBP":    {"OHM", "FRAX", "USDC"},
		"stkcvxOHMFRAXBP": {"OHM", "FRAX", "USDC"},
		//
		"MIM_3LP3CRV":       {"USDC", "USDT", "DAI", "MIM"},
		"cvxMIM_3LP3CRV":    {"USDC", "USDT", "DAI", "MIM"},
		"stkcvxMIM_3LP3CRV": {"USDC", "USDT", "DAI", "MIM"},
		//
		"crvCRVETH":       {"CRV", "WETH"},
		"cvxcrvCRVETH":    {"CRV", "WETH"},
		"stkcvxcrvCRVETH": {"CRV", "WETH"},
		//
		"crvCVXETH":       {"CVX", "WETH"},
		"cvxcrvCVXETH":    {"CVX", "WETH"},
		"stkcvxcrvCVXETH": {"CVX", "WETH"},
		//
		"crvUSDTWBTCWETH":       {"USDT", "WBTC", "WETH"},
		"cvxcrvUSDTWBTCWETH":    {"USDT", "WBTC", "WETH"},
		"stkcvxcrvUSDTWBTCWETH": {"USDT", "WBTC", "WETH"},
		//
		"LDOETH":       {"LDO", "WETH"},
		"cvxLDOETH":    {"LDO", "WETH"},
		"stkcvxLDOETH": {"LDO", "WETH"},

		//
		"crvUSDUSDC":         {"crvUSD", "USDC"},
		"crvUSDUSDT":         {"crvUSD", "USDT"},
		"crvUSDFRAX":         {"crvUSD", "WETH", "FRAX"},
		"crvUSDETHCRV":       {"crvUSD", "WETH", "CRV"},
		"cvxcrvUSDETHCRV":    {"crvUSD", "WETH", "CRV"},
		"stkcvxcrvUSDETHCRV": {"crvUSD", "WETH", "CRV"},

		// due to v3 compatibility
		"rETH_f":                  {"WETH"},
		"cLINK":                   {"LINK"},
		"sDAI":                    {"DAI"},
		"YieldETH":                {},
		"USDC_DAI_USDT":           {},
		"B_rETH_STABLE":           {},
		"auraB_rETH_STABLE":       {},
		"auraB_rETH_STABLE_vault": {},
		// redstones
		"weETH": {"WETH"},
		// "ezETH":  {"WETH"},
		"rswETH": {"WETH"},
		"pufETH": {"WETH"},
		"rsETH":  {"WETH"},
		"pzETH":  {"stETH"},
		//
		"steakLRT": {"stETH"},
		"eBTC":     {"WBTC"},
	}
	if log.GetBaseNet(chainId) != "MAINNET" {
		for sym, deps := range depGraph {
			x := make([]string, 0, len(deps))
			for _, d := range deps {
				if d != "stETH" {
					x = append(x, d)
				}
			}
			depGraph[sym] = x
		}
		delete(depGraph, "cLINK") // for non mainnet remove stETH.
	}
	return depGraph
}

// token to token dependent on it
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
	depQueryPFSym := q.getTokenSym(token)
	if q.depGraph[depQueryPFSym] == nil {
		log.Infof("Warn: Dep for query based price feed(%s) not found for token(%s) at %d", oracle, depQueryPFSym, blockNum)
	}
}

func (q *QueryPFDependencies) updateQueryPrices(pfs []*schemas.PriceFeed) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.depBasedExtraPrices = append(q.depBasedExtraPrices, pfs...)
}

// clearExtraBefore is used to remove price feed before the lastSync of aggregatedBlockFeed
func (q *QueryPFDependencies) extraPriceForQueryFeed(deleteExtraBefore int64) []*schemas.PriceFeed {
	updates := q.getChainlinkBasedQueryUpdates(deleteExtraBefore)
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
	var calls []multicall.Multicall2Call
	queryAdapters := []adapterAndNoCall{}
	//
	for _, adapter := range q.aqf.getFeedAdapters(blockNum, tokens) {
		otherCalls, isQueryable := adapter.GetCalls(blockNum)
		if isQueryable {
			calls = append(calls, otherCalls...)
			queryAdapters = append(queryAdapters, adapterAndNoCall{
				adapter: adapter,
				nocalls: len(otherCalls),
			})
		}
	}
	//
	_results := core.MakeMultiCall(q.client, blockNum, false, calls, 30)
	iterator := core.NewMulticallResultIterator(_results)
	//
	for _, entry := range queryAdapters {
		var results []multicall.Multicall2Result
		for i := 0; i < entry.nocalls; i++ {
			results = append(results, iterator.Next())
		}
		prices := processRoundDataWithAdapterTokens(blockNum, entry.adapter, results)
		q.updateQueryPrices(prices)
	}
	// sync control
	<-ch
	wg.Done()
}
