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
	if repo.config.ChainId != 1 {
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
	repo.aggregatedFeed.AddPools(token, &core.UniswapPools{
		V2:      poolv2Addr.Hex(),
		V3:      poolv3Addr.Hex(),
		Updated: true,
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

func (repo *Repository) GetUniPricesByToken(token string) []*core.UniPoolPrices {
	return repo.aggregatedFeed.GetUniPricesByToken(token)
}
