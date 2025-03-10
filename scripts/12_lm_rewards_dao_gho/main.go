package main

import (
	"context"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg"
	"github.com/Gearbox-protocol/sdk-go/pkg/priceFetcher"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/ds/dc_wrapper"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	v3 "github.com/Gearbox-protocol/third-eye/models/pool_lmrewards/v3"
	"github.com/Gearbox-protocol/third-eye/repository"
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm/utils"
)

// dUSDC - farmedUSDCv3 https://etherscan.io/tx/0x737fb7e55268d6ef95806c60e074948515fa19e1add8499f20f18ad3f62cf250
type Repo struct {
	ds.DummyRepo
	client       core.ClientI
	events       []*schemas.PoolLedger
	dc           *dc_wrapper.DataCompressorWrapper
	dieselTokens map[string]*schemas.UTokenAndPool
	tStore       *priceFetcher.TokensStore
}

func (r *Repo) AddPoolLedger(event *schemas.PoolLedger) {
	r.events = append(r.events, event)
}
func (r *Repo) GetToken(token string) *schemas.Token {
	x, err := r.tStore.GetToken(token)
	log.CheckFatal(err)
	return x
}
func (r *Repo) getPoolLedger() []*schemas.PoolLedger {
	ans := r.events
	r.events = nil
	return ans
}

func (r *Repo) AddDieselToken(dieselToken, underlyingToken, pool string, version core.VersionType) {
	r.dieselTokens[dieselToken] = &schemas.UTokenAndPool{
		UToken:  underlyingToken,
		Pool:    pool,
		Version: version,
	}
}

func (r *Repo) GetDieselTokens() map[string]*schemas.UTokenAndPool {
	return r.dieselTokens
}

func (repo *Repo) SetAndGetBlock(blockNum int64) *schemas.Block {
	b, err := repo.client.BlockByNumber(context.Background(), big.NewInt(blockNum))
	if err != nil {
		log.Fatalf("%s: %d", err, blockNum)
	}
	return &schemas.Block{Timestamp: b.Time(), BlockNumber: blockNum}
}
func (r *Repo) GetDCWrapper() *dc_wrapper.DataCompressorWrapper {
	return r.dc
}

func NewRepo(client core.ClientI) *Repo {
	r := Repo{dieselTokens: map[string]*schemas.UTokenAndPool{}, tStore: priceFetcher.NewTokensStore(client)}
	r.dc = dc_wrapper.NewDataCompressorWrapper(client)
	r.client = client
	r.dc.AddDataCompressorByVersion(core.NewVersion(300), "0x104c4e209329524adb0febE8b6481346a6eB75C6", 19363134)
	return &r
}

func getFarmsv3(client core.ClientI, repo *Repo) []*v3.Farmv3 {
	poolAndFarms := []*v3.Farmv3{}
	pools, found := repo.GetDCWrapper().GetPoolListv3()
	if found {
		farmingPools := map[common.Address]common.Address{}
		for _, pool := range pools {
			if !utils.Contains([]string{"0x4d56c9cBa373AD39dF69Eb18F076b7348000AE09", "0xe7146F53dBcae9D6Fa3555FE502648deb0B2F823"}, pool.Addr.Hex()) {
				continue
			}
			for _, zapper := range pool.Zappers {
				// can be diselToken zapperOut -- https://etherscan.io/address/0xcaa199f91294e6ee95f9ea90fe716cbd2f9f2900#code
				if _, ok := farmingPools[zapper.TokenOut]; ok && zapper.TokenIn == pool.Underlying && zapper.TokenOut != pool.DieselToken {
					if zapper.TokenOut.Hex() == "0x580e39ADb33E106fFc2712cBD57B9cE954dcfE75" { // GHO
						zapper.TokenOut = common.HexToAddress("0xE2037090f896A858E3168B978668F22026AC52e7")
					}
					if zapper.TokenOut.Hex() == "0x7aB44F17EE21A3D6Bb2aeb1c6cA8B875041608C4" { // DAI
						zapper.TokenOut = common.HexToAddress("0xC853E4DA38d9Bd1d01675355b8c8f3BBC1451973")
					}
					log.Info(zapper.TokenOut)
					poolAndFarms = append(poolAndFarms, &v3.Farmv3{
						Farm:        zapper.TokenOut.Hex(),
						Pool:        pool.Addr.Hex(),
						DieselToken: pool.DieselToken.Hex(),
						// initial
						Fpt:         (*core.BigInt)(new(big.Int)),
						TotalSupply: (*core.BigInt)(new(big.Int)),
						Reward:      (*core.BigInt)(new(big.Int)),
					})
				}
			}
		}
	}
	return poolAndFarms
}
func main() {
	cfg := config.NewConfig()
	db := repository.NewDBClient(cfg)
	client := ethclient.NewEthClient(cfg)
	r := NewRepo(client)
	var adapters []*ds.SyncAdapter
	{
		err := db.Raw(`SELECT * from sync_adapters where type='LMRewardsv3'`).Find(&adapters).Error
		log.CheckFatal(err)
	}

	{ // add diesel tokens
		states := []*schemas.PoolState{}
		err := db.Raw(`SELECT * from pools`).Find(&states).Error
		log.CheckFatal(err)
		for _, state := range states {
			r.AddDieselToken(state.DieselToken, state.UnderlyingToken, state.Address, state.Version)
		}
	}

	// create sync_adapters
	var lm *v3.LMRewardsv3
	for _, adapter := range adapters {
		adapter.Client = client
		adapter.Repo = r
		lm = v3.NewLMRewardsv3FromAdapter(adapter)
	}
	lm.SetUnderlyingState(getFarmsv3(client, r))
	lm.SetUnderlyingState([]*v3.UserLMDetails{})

	// update the logs
	txLogs, err := pkg.Node{Client: client}.GetLogs(lm.FirstLogAt, lm.LastSync,
		lm.GetAllAddrsForLogs(),
		[][]common.Hash{{
			core.Topic("Transfer(address,address,uint256)"),
			core.Topic("RewardUpdated(uint256,uint256)"),
		}})
	log.CheckFatal(err)
	for _, txLog := range txLogs {
		log.Info(txLog.BlockNumber)
		lm.OnLog(txLog)
	}
	//
	currentTs := r.SetAndGetBlock(lm.LastSync).Timestamp
	tx := db.Begin()
	lm.Save(tx, currentTs)
	info := tx.Commit()
	log.CheckFatal(info.Error)
}
