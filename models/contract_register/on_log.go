package contract_register

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/credit_manager/cm_v1"
	"github.com/Gearbox-protocol/third-eye/models/credit_manager/cm_v2"
	"github.com/Gearbox-protocol/third-eye/models/credit_manager/cm_v3"
	"github.com/Gearbox-protocol/third-eye/models/pool/pool_v2"
	"github.com/Gearbox-protocol/third-eye/models/pool/pool_v3"
	v3 "github.com/Gearbox-protocol/third-eye/models/pool_lmrewards/v3"
	"github.com/Gearbox-protocol/third-eye/models/price_oracle/po_v3"
)

func (mdl *ContractRegister) OnLog(txLog types.Log) {
	blockNum := int64(txLog.BlockNumber)
	switch txLog.Topics[0] {
	case core.Topic("NewPoolAdded(address)"):
		address := common.HexToAddress(txLog.Topics[1].Hex()).Hex()
		// log.Info("new pool", address)
		obj := NewPool(address, mdl.SyncAdapter.Client, mdl.Repo, blockNum)
		mdl.Repo.AddSyncAdapter(obj)
	case core.Topic("NewCreditManagerAdded(address)"):
		address := common.HexToAddress(txLog.Topics[1].Hex()).Hex()
		// log.Info("new cm", address)
		cm := NewCM(address, mdl.SyncAdapter.Client, mdl.Repo, blockNum)
		mdl.Repo.AddSyncAdapter(cm)
		// for v310
	case core.Topic("CreateMarket(address,address,string,string)"):
		// create marketpool and pooloracle relations
		// add priceoracle adapter and add pool is not present
		poolAddr := common.BytesToAddress(txLog.Topics[1][:]).Hex()

		priceOracleAddr := mdl.GetPriceOracle(poolAddr, blockNum)
		mdl.Repo.AddRelation(&schemas.Relation{
			Type:      "MarketPool",
			Owner:     mdl.Address,
			Dependent: poolAddr,
			BlockNum:  blockNum,
		})
		mdl.Repo.AddRelation(&schemas.Relation{
			Type:      "PoolOracle",
			Owner:     poolAddr,
			Dependent: priceOracleAddr,
			BlockNum:  blockNum,
		})
		//
		if priceOracle := mdl.Repo.GetAdapter(priceOracleAddr); priceOracle == nil {
			po := po_v3.NewPriceOracle(priceOracleAddr, blockNum, mdl.Client, mdl.Repo)
			mdl.Repo.AddSyncAdapter(po)
		}
		//
		poolAdapter := mdl.Repo.GetAdapter(poolAddr)
		if poolAdapter == nil {
			pool := pool_v3.NewPool(poolAddr, mdl.Client, mdl.Repo, blockNum, mdl.Address, schemas.PriceOracleT(priceOracleAddr), 310)
			mdl.Repo.AddSyncAdapter(pool)
		} else {
			if poolv3, ok := poolAdapter.(*pool_v3.Poolv3); ok {
				poolv3.State.Market = mdl.Address
				poolv3.State.PriceOracle = schemas.PriceOracleT(priceOracleAddr)
			} else {
				log.Fatalf("the pool by the address(%s) is not poolv3", poolAddr)
			}
		}
	case core.Topic("RegisterCreditSuite(address,address)"):
		cm := cm_v3.NewCMv3(common.BytesToAddress(txLog.Topics[1][:]).Hex(), mdl.Client, mdl.Repo, blockNum)
		mdl.Repo.AddSyncAdapter(cm)
	}
}

func NewCM(addr string, client core.ClientI, repo ds.RepositoryI, blockNum int64) ds.SyncAdapterI {
	version := core.FetchVersion(addr, blockNum, client)
	switch version {
	case core.NewVersion(1):
		return cm_v1.NewCMv1(addr, client, repo, blockNum)
	case core.NewVersion(2):
		return cm_v2.NewCMv2(addr, client, repo, blockNum)
	default:
		if version.MoreThanEq(core.NewVersion(300)) {
			return cm_v3.NewCMv3(addr, client, repo, blockNum)
		}
	}
	log.Fatalf("Version(%d) of cm can't be created.", version)
	return nil
}

func NewPool(addr string, client core.ClientI, repo ds.RepositoryI, blockNum int64) ds.SyncAdapterI {
	version := core.FetchVersion(addr, blockNum, client)
	switch version {
	case core.NewVersion(1), core.NewVersion(2):
		return pool_v2.NewPool(addr, client, repo, blockNum)
	default:
		if version.MoreThanEq(core.NewVersion(300)) {
			// add pool to the lmrewards so that farm_v3 table entry can be created.
			adapters := repo.GetAdapterAddressByName(ds.LMRewardsv3)
			lmRewards := repo.GetAdapter(adapters[0])
			lmRewards.(*v3.LMRewardsv3).AddPoolv3(blockNum, addr)
			// add pool
			return pool_v3.NewPool(addr, client, repo, blockNum, core.NULL_ADDR.Hex(), schemas.PriceOracleT(core.NULL_ADDR.Hex()), 300)
		}
	}
	log.Fatalf("Version(%d) of pool can't be created.", version)
	return nil
}

func (mdl *ContractRegister) GetPriceOracle(pool string, blockNum int64) string {
	hash := common.HexToHash(pool)
	data, err := core.CallFuncGetSingleValue(mdl.Client, "01374518", common.HexToAddress(mdl.Address), blockNum, hash[:])
	log.CheckFatal(err)
	return common.BytesToAddress(data).Hex()
}
