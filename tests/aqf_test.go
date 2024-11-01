package tests

import (
	"sort"
	"testing"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/test"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/aggregated_block_feed"
	"github.com/Gearbox-protocol/third-eye/tests/framework"
	"github.com/ethereum/go-ethereum/common"
)

type tr struct {
	ds.DummyRepo
	addressMap map[string]string
	// t          handlers.TokenOracleRepo
}

// TokensValidAtBlock not implemented
func (mdl tr) TokenAddrsValidAtBlock(addr string, blockNum int64) map[string]bool {
	switch addr {
	case mdl.addressMap["YearnFeed_1"]:
		if blockNum >= 50 {
			return nil
		}
		return map[string]bool{mdl.addressMap["Token_2"]: true, mdl.addressMap["Token_3"]: true, mdl.addressMap["Token_4"]: true}
	case mdl.addressMap["YearnFeed_3"]:
		if blockNum >= 56 {
			return map[string]bool{mdl.addressMap["Token_2"]: true, mdl.addressMap["Token_3"]: true}
		}
		if blockNum >= 50 {
			return map[string]bool{mdl.addressMap["Token_2"]: true, mdl.addressMap["Token_3"]: true, mdl.addressMap["Token_4"]: true}
		}
		return nil
	case mdl.addressMap["YearnFeed_4"]:
		if blockNum >= 56 {
			return map[string]bool{mdl.addressMap["Token_4"]: true}
		}
		return nil
	}
	return nil
}

func TestAQFWrapper(t *testing.T) {
	log.SetTestLogging(t)
	client := test.NewTestClient()
	inputFile, addressMap := framework.ParseMockClientInput(t, client, []string{"aqf/input.json"})
	//
	r := &tr{addressMap: addressMap}
	aqf := aggregated_block_feed.NewAQFWrapper(client, r, 25) // 25 is the sync interval
	updateAQF(t, aqf, addressMap, inputFile, client)
	//
	//
	//
	aqf.Query(60)
	sort.Slice(r.PFs, func(a, b int) bool {
		aBlock := r.PFs[a].BlockNumber
		bBlock := r.PFs[b].BlockNumber
		return aBlock < bBlock || (aBlock == bBlock && r.PFs[a].RoundId < r.PFs[b].RoundId)
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
	// aqf.DisableYearnFeed(addressMap["Token_4"], addressMap["YearnFeed_3"], 56, schemas.V2PF)
	// aqf.AddFeedOrToken(addressMap["Token_4"], addressMap["YearnFeed_4"], ds.YearnPF, 56, schemas.V2PF)
}
