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

type repoI interface {
	GetAdapter(addr string) ds.SyncAdapterI
	// if returned value is nil, it means that token oracle hasn't been added yet.
	GetOracleForV2Token(token string) *schemas.TokenOracle
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
	repo repoI
	TokenSymMap
	mu     *sync.Mutex
	client core.ClientI
}

func NewQueryPFDepenencies(repo ds.RepositoryI, client core.ClientI) *QueryPFDependencies {
	chainId, err := client.ChainID(context.TODO())
	log.CheckFatal(err)
	depGraph := getDepGraph()
	return &QueryPFDependencies{
		depGraph:                    depGraph,
		ChainlinkSymToQueryPFSyms:   getInvertDependencyGraph(depGraph),
		ChainlinkSymToUpdatedBlocks: map[string][]int64{},
		//
		repo:        repo,
		mu:          &sync.Mutex{},
		TokenSymMap: newTokenSymMap(chainId.Int64()),
		client:      client,
	}
}

func (q *QueryPFDependencies) ChainlinkPriceUpdatedAt(token string, blockNums []int64) {
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
func getDepGraph() map[string][]string {
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
		log.Fatalf("Dep for query based price feed(%s) not found for token(%s) at %d", oracle, depQueryPFSym, blockNum)
	}
}

func (q *QueryPFDependencies) updateQueryPrices(pfs []*schemas.PriceFeed) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.depBasedExtraPrices = append(q.depBasedExtraPrices, pfs...)
}

// clearExtraBefore is used to remove price feed before the lastSync of aggregatedBlockFeed
func (q *QueryPFDependencies) extraPriceForQueryFeed(clearExtraBefore int64) []*schemas.PriceFeed {
	updates := q.getChainlinkBasedQueryUpdates(clearExtraBefore)
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
	// get the latestRoundData call data
	priceFeedABI := core.GetAbi("PriceFeed")
	data, err := priceFeedABI.Pack("latestRoundData")
	log.CheckFatal(err)

	// generate calls
	var calls []multicall.Multicall2Call
	var tokenOracles []*schemas.TokenOracle
	for token := range tokens {
		details := q.repo.GetOracleForV2Token(token)
		if details == nil {
			continue
		}
		tokenOracles = append(tokenOracles, details)
		call := multicall.Multicall2Call{
			Target:   common.HexToAddress(details.Feed),
			CallData: data,
		}
		calls = append(calls, call)
	}
	// get result
	results := core.MakeMultiCall(q.client, blockNum, false, calls, 30)
	// parse result and create PriceFeed obj
	var newPrices []*schemas.PriceFeed
	for ind, entry := range results {
		details := tokenOracles[ind]
		var newPrice *schemas.PriceFeed
		/// parse price
		// there is no check that call was made for feed that is existing for given blockNum
		if entry.Success && len(entry.ReturnData) != 0 {
			newPrice = parseRoundData(entry.ReturnData, true, details.Feed) // only valid for v2
		} else {
			// if failed check and pfType of the queryPrice is YearnPF
			adapterI := q.repo.GetAdapter(details.Feed)
			if adapter, ok := adapterI.(*QueryPriceFeed); !ok {
				log.Fatal("Conversion of adapter to queryPriceFeed failed ", details.Feed)
			} else if adapter.GetDetailsByKey("pfType") == ds.YearnPF {
				// if underlying price feed address is null, then don't set price
				if _newPrice, err := adapter.calculateYearnPFInternally(blockNum); err == nil {
					newPrice = _newPrice
				}
			}
		}
		if newPrice != nil {
			// add token and feed details
			newPrice.Token = details.Token
			newPrice.Feed = details.Feed
			newPrice.BlockNumber = blockNum
			newPrices = append(newPrices, newPrice)
		}
	}
	q.updateQueryPrices(newPrices)
	// sync control
	<-ch
	wg.Done()
}
