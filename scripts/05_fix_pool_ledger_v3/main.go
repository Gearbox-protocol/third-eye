package main

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg"
	"github.com/Gearbox-protocol/sdk-go/pkg/priceFetcher"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/ds/dc_wrapper"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/models/pool/pool_v3"
	"github.com/Gearbox-protocol/third-eye/repository"
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm/clause"
)

// dUSDC - farmedUSDCv3 https://etherscan.io/tx/0x737fb7e55268d6ef95806c60e074948515fa19e1add8499f20f18ad3f62cf250
type Repo struct {
	ds.DummyRepo
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

func (r *Repo) GetDCWrapper() *dc_wrapper.DataCompressorWrapper {
	return r.dc
}

func NewRepo(client core.ClientI) *Repo {
	r := Repo{dieselTokens: map[string]*schemas.UTokenAndPool{}, tStore: priceFetcher.NewTokensStore(client)}
	r.dc = dc_wrapper.NewDataCompressorWrapper(client)
	r.dc.AddDataCompressorv300(core.NewVersion(300), "0xc0101abAFce0BD3de10aa1F3dd827672B150436E", 18798875)
	return &r
}
func main() {
	cfg := config.NewConfig()
	db := repository.NewDBClient(cfg)
	client := ethclient.NewEthClient(cfg)
	r := NewRepo(client)

	{
		err := db.Exec(`UPDATE sync_adapters set details='{}' where type='Pool' and version=300`).Error
		log.CheckFatal(err)
	}
	var adapters []*ds.SyncAdapter
	err := db.Raw(`SELECT * from sync_adapters where type='Pool' and version=300`).Find(&adapters).Error
	log.CheckFatal(err)

	states := []*schemas.PoolState{}
	err = db.Raw(`SELECT * from pools where  _version=300`).Find(&states).Error
	log.CheckFatal(err)
	{
		states := []*schemas.PoolState{}
		err = db.Raw(`SELECT * from pools`).Find(&states).Error
		log.CheckFatal(err)
		for _, state := range states {
			r.AddDieselToken(state.DieselToken, state.UnderlyingToken, state.Address, state.Version)
		}
	}

	pools := map[string]*pool_v3.Poolv3{}
	for _, adapter := range adapters {
		adapter.Client = client
		adapter.Repo = r
		pool := pool_v3.NewPoolFromAdapter(adapter)
		pools[pool.GetAddress()] = pool
	}

	for _, state := range states {
		pool := pools[state.Address]
		pool.SetUnderlyingState(state)
	}

	for _, pool := range pools {
		txLogs, err := pkg.Node{Client: client}.GetLogs(pool.FirstLogAt, pool.LastSync,
			pool.GetAllAddrsForLogs(),
			[][]common.Hash{{
				core.Topic("Transfer(address,address,uint256)"),
				core.Topic("Deposit(address,address,uint256,uint256)"),
				core.Topic("Withdraw(address,address,address,uint256,uint256)"),
			}})
		log.CheckFatal(err)
		for _, txLog := range txLogs {
			// if txLog.BlockNumber == 18814974 {
			// 	log.Info("here")
			// }
			log.Info(pool.GetAddress(), txLog.BlockNumber)
			pool.OnLog(txLog)
		}
		err = db.Exec(`DELETE from pool_ledger where pool=? and event in ('AddLiquidity','RemoveLiquidity')`, pool.GetAddress()).Error
		log.CheckFatal(err)
		events := r.getPoolLedger()
		err = db.CreateInBatches(events, 500).Error
		log.CheckFatal(err)
		pool.UpdatePoolv2Ledger(db)
		err = db.Clauses(clause.OnConflict{UpdateAll: true}).Create(pool.SyncAdapter).Error
		log.CheckFatal(err)
	}
}
