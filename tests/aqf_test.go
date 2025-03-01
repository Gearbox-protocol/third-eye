package tests

import (
	"sort"
	"testing"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/test"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/aggregated_block_feed"
	"github.com/Gearbox-protocol/third-eye/tests/framework"
	"github.com/ethereum/go-ethereum/common"
)

func TestAQFWrapper(t *testing.T) {
	log.SetTestLogging(t)
	client := test.NewTestClient()
	inputFile, addressMap := framework.ParseMockClientInput(t, client, []string{"aqf/input.json"})
	//
	r := &ds.DummyRepo{}
	aqf := aggregated_block_feed.NewAQFWrapper(client, r, 25) // 25 is the sync interval
	updateAQF(t, aqf, addressMap, inputFile, client)
	//
	//
	//
	aqf.Query(60)
	reverseAddrMap := reverseMap(addressMap)
	sort.Slice(r.PFs, func(a, b int) bool {
		aBlock := r.PFs[a].BlockNumber
		bBlock := r.PFs[b].BlockNumber
		return aBlock < bBlock || (aBlock == bBlock && reverseAddrMap[r.PFs[a].Token] < reverseAddrMap[r.PFs[b].Token])
	})

	framework.Check(t, addressMap, map[string]interface{}{"data": r.PFs}, "aqf/blocks.json")
	// framework.Print(t, addressMap, map[string]interface{}{"data": r.PFs})
}

func updateAQF(t *testing.T, aqf *aggregated_block_feed.AQFWrapper, addressMap map[string]string, inputFile *framework.TestInput3Eye, client core.ClientI) {

	//
	syncAdapterObj := inputFile.GetSyncAdapter(addressMap, t)
	// set feed to token
	if syncAdapterObj != nil {
		for _, adapter := range syncAdapterObj.Adapters {
			adapter.Client = client
			aqf.AddQueryPriceFeed(aggregated_block_feed.NewQueryPriceFeedFromAdapter(adapter))
		}
	}
	log.Info(addressMap)

	tokenSymMap := aggregated_block_feed.TokenSymMapFromchainId(1)
	for tokenVar, sym := range map[string]string{
		"Token_1": "OHM",
		"Token_2": "OHMFRAXBP",
		"Token_3": "cvxOHMFRAXBP",
		"Token_4": "stkcvxOHMFRAXBP",
	} {
		addr := common.HexToAddress(addressMap[tokenVar])
		tokenSymMap.UpdateForTest(sym, addr)
	}
	aqf.GetDepFetcher().TokenSymMap = tokenSymMap
	aqf.ChainlinkPriceUpdatedAt(addressMap["Token_1"], []int64{4, 11, 26, 51, 53, 58})
	//
	aqf.DisableYearnFeed(addressMap["Token_4"], addressMap["YearnFeed_3"], 56, schemas.V2PF)
	aqf.AddFeedOrToken(addressMap["Token_4"], addressMap["YearnFeed_4"], ds.YearnPF, 56, schemas.V2PF, nil)
}

func reverseMap(in map[string]string) (r map[string]string) {
	r = map[string]string{}
	for k, v := range in {
		r[v] = k
	}
	return
}
