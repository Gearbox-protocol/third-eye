package aggregated_block_feed

import (
	"strings"
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
	// chainlink symbol updated at these blocks, we need to fetch price for dependent query pf at these block numbers
	ChainlinkSymToUpdatedBlocks map[string][]int64
	// blocks to remove
	aggregatedFetchedBlocks []int64
	// yearn feed prices to be ordered
	depBasedExtraPrices []*schemas.PriceFeed
	TestI               map[string][]string
	//
	aqf *AQFWrapper
	//
	repo   repoI
	mu     *sync.Mutex
	client core.ClientI
}

func NewQueryPFDepenencies(repo ds.RepositoryI, client core.ClientI) *QueryPFDependencies {
	return &QueryPFDependencies{
		ChainlinkSymToUpdatedBlocks: map[string][]int64{},
		//
		repo:   repo,
		mu:     &sync.Mutex{},
		client: client,
	}
}

func (q *QueryPFDependencies) chainlinkPriceUpdatedAt(token string, blockNums []int64) {
	chainlinkSym := q.repo.GetToken(token).Symbol
	q.mu.Lock()
	defer q.mu.Unlock()
	q.ChainlinkSymToUpdatedBlocks[chainlinkSym] = blockNums
}

func (q *QueryPFDependencies) getChainlinkBasedQueryUpdates(clearExtraBefore int64) map[int64]map[string]bool {
	//  blockNum to QueryPFToken
	updates := map[int64]map[string]bool{}
	var updatedChainlinkSym []string
	chainlinkToQueryTokens := q.GetChainlinkTokenToUpdateToken()
	for chainlinkSym, blockNums := range q.ChainlinkSymToUpdatedBlocks {
		updatedChainlinkSym = append(updatedChainlinkSym, chainlinkSym)
		//
		for _, dependentAddr := range chainlinkToQueryTokens[chainlinkSym] {
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
				updates[blockNum][dependentAddr] = true
			}
		}
	}
	if updatedChainlinkSym != nil {
		log.Info("Updated chainlinks are:", updatedChainlinkSym)
	}
	q.ChainlinkSymToUpdatedBlocks = map[string][]int64{}
	return updates
}

var base = []string{"WETH", "WBTC", "DAI", "USDC", "USDT", "USDC", "OHM"}
var combo = map[string][]string{
	"3crv": {"DAI", "USDC", "USDT"},
}

// {"USDT": {"3crv:address"}, "USDC": {"3crv:address"}, "DAI": {"3crv:address"}, "FRAX": {"crvFRAX"}}
func (q *QueryPFDependencies) GetChainlinkTokenToUpdateToken() map[string][]string {
	if core.GetChainId(q.client) == 1337 {
		return q.TestI
	}
	tokens := []*schemas.Token{}
	for _, token := range q.repo.GetTokens() {
		tokens = append(tokens, q.repo.GetToken(token))
	}
	ans := map[string][]string{}
	for _, token := range tokens {
		for _, sym := range base {
			if strings.Contains(token.Symbol, sym) && token.Symbol != sym {
				ans[token.Symbol] = append(ans[token.Symbol], token.Address)
			}
			if strings.Contains(strings.ToLower(token.Symbol), "3crv") {
				for _, underlyingsym := range combo["3crv"] {
					ans[underlyingsym] = append(ans[underlyingsym], token.Address)
				}
			}
		}
	}
	return ans
}

// func (q *QueryPFDependencies) checkInDepGraph(token, oracle string, blockNum int64) {
// 	depQueryPFSym := q.getTokenSym(token)
// 	if q.depGraph[depQueryPFSym] == nil {
// 		log.Infof("Warn: Dep for query based price feed(%s) not found for token(%s) at %d", oracle, depQueryPFSym, blockNum)
// 	}
// }

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
