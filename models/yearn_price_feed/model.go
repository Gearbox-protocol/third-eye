package yearn_price_feed

import (
	"github.com/Gearbox-protocol/third-eye/artifacts/yearnPriceFeed"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/ethereum/go-ethereum/common"
)

type YearnPriceFeed struct {
	*core.SyncAdapter
	contractETH *yearnPriceFeed.YearnPriceFeed
}

func NewYearnPriceFeed(token, oracle string, discoveredAt int64, client *ethclient.Client, repo core.RepositoryI) *YearnPriceFeed {
	syncAdapter := &core.SyncAdapter{
		Contract: &core.Contract{
			Address:      oracle,
			DiscoveredAt: discoveredAt,
			FirstLogAt:   discoveredAt,
			ContractName: core.YearnPriceFeed,
			Client:       client,
		},
		Details:  map[string]string{"token": token},
		LastSync: discoveredAt - 1,
		Repo:     repo,
	}
	// add token oracle for db
	// feed is also oracle address for yearn address
	// we don't relie on underlying feed
	repo.AddTokenOracle(token, oracle, oracle, discoveredAt)
	return NewYearnPriceFeedFromAdapter(
		syncAdapter,
	)
}

func NewYearnPriceFeedFromAdapter(adapter *core.SyncAdapter) *YearnPriceFeed {
	yearnPFContract, err := yearnPriceFeed.NewYearnPriceFeed(common.HexToAddress(adapter.Address), adapter.Client)
	if err != nil {
		log.Fatal(err)
	}
	obj := &YearnPriceFeed{
		SyncAdapter: adapter,
		contractETH: yearnPFContract,
	}
	obj.OnlyQuery = true
	return obj
}
