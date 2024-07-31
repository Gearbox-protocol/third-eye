package tests

import (
	"sort"
	"testing"

	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg/redstone"
	"github.com/Gearbox-protocol/sdk-go/test"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/aggregated_block_feed"
	"github.com/Gearbox-protocol/third-eye/tests/framework"
)

type repoForAQF struct {
	mgr redstone.RedStoneMgrI
	ds.DummyRepo
	tokens map[string]*schemas.Token
}

func NewrepoForAQF(mgr redstone.RedStoneMgrI) *repoForAQF {
	return &repoForAQF{
		mgr:    mgr,
		tokens: map[string]*schemas.Token{},
	}
}
func (r *repoForAQF) GetRedStonemgr() redstone.RedStoneMgrI {
	return r.mgr
}

func (r *repoForAQF) SetAndGetBlock(blockNum int64) *schemas.Block {
	if blockNum == 26 {
		return &schemas.Block{
			BlockNumber: blockNum,
			Timestamp:   1722371001, // 2024, 30 jun starting
		}
	}
	return nil
}

func (r *repoForAQF) GetToken(token string) *schemas.Token {
	return r.tokens[token]
}

func TestAQFMultipleAdapter(t *testing.T) {
	log.SetTestLogging(t)
	client := test.NewTestClient()
	r := NewrepoForAQF(redstone.NewRedStoneMgr(client))
	//
	inputFile, addressMap := framework.ParseMockClientInput(t, client, []string{"aqf_multiple_adapters/input.json"})
	for _, token := range client.GetToken() {
		r.tokens[token.Address] = token
	}
	aqf := aggregated_block_feed.NewAQFWrapper(client, r, 25) // 25 is the sync interval
	//
	log.Info(addressMap)
	syncAdapterObj := inputFile.GetSyncAdapter(addressMap, t)
	// set feed to token
	if syncAdapterObj != nil {
		for _, adapter := range syncAdapterObj.Adapters {
			adapter.Client = client
			adapter.Repo = r
			aqf.AddQueryPriceFeed(aggregated_block_feed.NewQueryPriceFeedFromAdapter(adapter))
		}
	}
	//
	aqf.Query(50)
	reverseAddrMap := _reverseMap(addressMap)
	sort.Slice(r.PFs, func(a, b int) bool {
		aBlock := r.PFs[a].BlockNumber
		bBlock := r.PFs[b].BlockNumber
		return aBlock < bBlock || (aBlock == bBlock && reverseAddrMap[r.PFs[a].Feed] < reverseAddrMap[r.PFs[b].Feed])
	})

	// framework.Print(t, addressMap, map[string]interface{}{"data": r.PFs})
	framework.Check(t, addressMap, map[string]interface{}{"data": r.PFs}, "aqf_multiple_adapters/blocks.json")
}

func _reverseMap(in map[string]string) (r map[string]string) {
	r = map[string]string{}
	for k, v := range in {
		r[v] = k
	}
	return
}
