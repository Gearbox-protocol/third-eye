package repository

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/uniswapv2Factory"
	"github.com/Gearbox-protocol/sdk-go/artifacts/uniswapv2Router"
	"github.com/Gearbox-protocol/sdk-go/artifacts/uniswapv3Factory"
	"github.com/Gearbox-protocol/sdk-go/artifacts/uniswapv3Router"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (repo *Repository) AddUniPoolsForToken(blockNum int64, token string) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if repo.GetWETHAddr() == token || // if the token is weth don't run for weth/weth pool
		repo.config.ChainId != 1 || // if not mainnet
		repo.AggregatedFeed.UniPoolByToken[token] != nil { // if the uni v2/v3 pool details already present don't add again
		return
	}
	v2FactoryAddr := repo.GetFactoryv2Address(blockNum)
	v3FactoryAddr := repo.GetFactoryv3Address(blockNum)
	v2Factory, err := uniswapv2Factory.NewUniswapv2Factory(v2FactoryAddr, repo.client)
	log.CheckFatal(err)
	poolv2Addr, err := v2Factory.GetPair(nil, common.HexToAddress(token), common.HexToAddress(repo.GetWETHAddr()))
	log.CheckFatal(err)
	//
	v3Factory, err := uniswapv3Factory.NewUniswapv3Factory(v3FactoryAddr, repo.client)
	log.CheckFatal(err)
	poolv3Addr, err := v3Factory.GetPool(nil, common.HexToAddress(token), common.HexToAddress(repo.GetWETHAddr()), big.NewInt(3000))
	log.CheckFatal(err)
	if poolv2Addr.Hex() == "0x0000000000000000000000000000000000000000" ||
		poolv3Addr.Hex() == "0x0000000000000000000000000000000000000000" {
		log.Fatalf("pool not fetched for v2/v3: %s/%s", token, repo.GetWETHAddr())
	}
	tokenInfo := repo.GetToken(token)
	repo.AggregatedFeed.AddUniPools(tokenInfo, &schemas.UniswapPools{
		V2:      poolv2Addr.Hex(),
		V3:      poolv3Addr.Hex(),
		Updated: true,
		Token:   token,
	})
}

func (repo *Repository) AddLastSyncForToken(token string, lastSync int64) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.AggregatedFeed.AddLastSyncForToken(token, lastSync)
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

func (repo *Repository) loadUniswapPools() {
	defer utils.Elapsed("loadUniswapPools")()
	data := []*schemas.UniswapPools{}
	err := repo.db.Raw(`SELECT * from uniswap_pools`).Find(&data).Error
	log.CheckFatal(err)
	for _, entry := range data {
		tokenInfo := repo.GetToken(entry.Token)
		repo.AggregatedFeed.AddUniPools(tokenInfo, entry)
	}
}

func (repo *Repository) loadChainlinkPrevState() {
	defer utils.Elapsed("loadChainlinkPrevState")()
	data := []*schemas.PriceFeed{}
	err := repo.db.Raw(`SELECT distinct on (feed)* from price_feeds order by feed, block_num DESC`).Find(&data).Error
	log.CheckFatal(err)
	for _, entry := range data {
		if adapter := repo.GetAdapter(entry.Feed); adapter != nil && adapter.GetName() == ds.ChainlinkPriceFeed {
			adapter.SetUnderlyingState(entry)
		}
	}
}

func (repo *Repository) GetUniPricesByToken(token string) []*schemas.UniPoolPrices {
	return repo.AggregatedFeed.GetUniPricesByToken(token)
}
func (repo *Repository) AddUniPriceAndChainlinkRelation(relation *schemas.UniPriceAndChainlink) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.relations = append(repo.relations, relation)
}
