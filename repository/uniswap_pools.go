package repository

import (
	"math/big"

	"github.com/Gearbox-protocol/third-eye/artifacts/uniswapv2Factory"
	"github.com/Gearbox-protocol/third-eye/artifacts/uniswapv2Router"
	"github.com/Gearbox-protocol/third-eye/artifacts/uniswapv3Factory"
	"github.com/Gearbox-protocol/third-eye/artifacts/uniswapv3Router"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (repo *Repository) AddPoolsForToken(blockNum int64, token string) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if repo.config.ChainId != 1 || repo.aggregatedFeed.UniPoolByToken[token] != nil {
		return
	}
	v2FactoryAddr := repo.GetFactoryv2Address(blockNum)
	v3FactoryAddr := repo.GetFactoryv3Address(blockNum)
	v2Factory, err := uniswapv2Factory.NewUniswapv2Factory(v2FactoryAddr, repo.client)
	log.CheckFatal(err)
	poolv2Addr, err := v2Factory.GetPair(nil, common.HexToAddress(token), common.HexToAddress(repo.WETHAddr))
	log.CheckFatal(err)
	//
	v3Factory, err := uniswapv3Factory.NewUniswapv3Factory(v3FactoryAddr, repo.client)
	log.CheckFatal(err)
	poolv3Addr, err := v3Factory.GetPool(nil, common.HexToAddress(token), common.HexToAddress(repo.WETHAddr), big.NewInt(3000))
	log.CheckFatal(err)
	if poolv2Addr.Hex() == "0x0000000000000000000000000000000000000000" ||
		poolv3Addr.Hex() == "0x0000000000000000000000000000000000000000" {
		log.Fatalf("pool not fetched for v2/v3: %s/%s", token, repo.WETHAddr)
	}
	tokenInfo, err := repo.getTokenWithError(token)
	log.CheckFatal(err)
	repo.aggregatedFeed.AddPools(tokenInfo, &core.UniswapPools{
		V2:      poolv2Addr.Hex(),
		V3:      poolv3Addr.Hex(),
		Updated: true,
		Token:   token,
	})
}

func (repo *Repository) AddLastSyncForToken(token string, lastSync int64) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.aggregatedFeed.AddLastSyncForToken(token, lastSync)
}

func (repo *Repository) GetFactoryv2Address(blockNum int64) common.Address {
	v2Router, err := uniswapv2Router.NewUniswapv2Router(common.HexToAddress(repo.config.Uniswapv2Router), repo.client)
	if err != nil {
		log.CheckFatal(err)
	}
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
	}
	v2Factory, err := v2Router.Factory(opts)
	return v2Factory
}

func (repo *Repository) GetFactoryv3Address(blockNum int64) common.Address {
	v3Router, err := uniswapv3Router.NewUniswapv3Router(common.HexToAddress(repo.config.Uniswapv3Router), repo.client)
	if err != nil {
		log.CheckFatal(err)
	}
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
	}
	v3Factory, err := v3Router.Factory(opts)
	return v3Factory
}

func (repo *Repository) AddUniswapPrices(prices *core.UniPoolPrices) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.setAndGetBlock(prices.BlockNum).AddUniswapPrices(prices)
}

func (repo *Repository) loadUniswapPools() {
	data := []*core.UniswapPools{}
	err := repo.db.Raw(`SELECT * from uniswap_pools`).Find(&data).Error
	log.CheckFatal(err)
	for _, entry := range data {
		tokenInfo, err := repo.getTokenWithError(entry.Token)
		log.CheckFatal(err)
		repo.aggregatedFeed.AddPools(tokenInfo, entry)
	}
}

func (repo *Repository) loadChainlinkPrevState() {
	data := []*core.PriceFeed{}
	err := repo.db.Raw(`SELECT distinct on (feed)* from price_feeds order by feed, block_num DESC`).Find(&data).Error
	log.CheckFatal(err)
	for _, entry := range data {
		if adapter := repo.kit.GetAdapter(entry.Feed); adapter != nil && adapter.GetName() == core.ChainlinkPriceFeed {
			adapter.SetUnderlyingState(entry)
		}
	}
}

func (repo *Repository) GetUniPricesByToken(token string) []*core.UniPoolPrices {
	return repo.aggregatedFeed.GetUniPricesByToken(token)
}
func (repo *Repository) AddUniPriceAndChainlinkRelation(relation *core.UniPriceAndChainlink) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.relations = append(repo.relations, relation)
}

func (repo *Repository) GetYearnFeedAddrs() (addrs []string) {
	feeds := repo.aggregatedFeed.GetYearnFeeds()
	for _, adapter := range feeds {
		addrs = append(addrs, adapter.GetAddress())
	}
	return
}
