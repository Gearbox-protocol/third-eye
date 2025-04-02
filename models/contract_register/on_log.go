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

func (mdl *ContractRegister) addCM(blockNum int64, cmAddr string) {
	cmAdapter := mdl.Repo.GetAdapter(cmAddr)
	if cmAdapter == nil {
		cm := NewCM(cmAddr, mdl.Client, mdl.Repo, blockNum)
		mdl.Repo.AddSyncAdapter(cm)
	}
}
func (mdl *ContractRegister) poolToMarketAndPriceOracleRelation(blockNum int64, poolAddr, priceOracleAddr string) {
	marketAddr := mdl.GetDetailsByKey("MARKET")
	mdl.Repo.AddRelation(&schemas.Relation{
		Type:      "MarketPool",
		Owner:     marketAddr, //
		Dependent: poolAddr,
		BlockNum:  blockNum,
	})
	mdl.Repo.AddRelation(&schemas.Relation{
		Type:      "PoolOracle",
		Owner:     poolAddr,
		Dependent: priceOracleAddr,
		BlockNum:  blockNum,
	})
	poolAdapter := mdl.Repo.GetAdapter(poolAddr)
	if poolv3, ok := poolAdapter.(*pool_v3.Poolv3); ok {
		log.Infof("Setting on pool %s, market: %s , po: %s ", poolv3.State.Name, mdl.GetDetailsByKey("MARKET"), priceOracleAddr)
		poolv3.State.Market = mdl.Address
		poolv3.State.PriceOracle = schemas.PriceOracleT(priceOracleAddr)
	} else {
		log.Fatalf("the pool by the address(%s) is not poolv3", poolAddr)
	}
	if priceOracle := mdl.Repo.GetAdapter(priceOracleAddr); priceOracle == nil {
		po := po_v3.NewPriceOracle(priceOracleAddr, blockNum, mdl.Client, mdl.Repo)
		mdl.Repo.AddSyncAdapter(po)
	}
}
func (mdl *ContractRegister) OnLog(txLog types.Log) {
	blockNum := int64(txLog.BlockNumber)
	switch txLog.Topics[0] {
	case core.Topic("NewPoolAdded(address)"):
		address := common.HexToAddress(txLog.Topics[1].Hex()).Hex()
		mdl.addNewPool(address, blockNum, core.NULL_ADDR.Hex(), core.NULL_ADDR)
		//
	case core.Topic("NewCreditManagerAdded(address)"):
		address := common.HexToAddress(txLog.Topics[1].Hex()).Hex()
		mdl.addCM(blockNum, address)
		// for v310
	case core.Topic("RegisterMarket(address,address,address)"):
		// create marketpool and pooloracle relations
		// add priceoracle adapter and add pool is not present
		poolAddr := common.BytesToAddress(txLog.Topics[1][:]).Hex()
		priceOracleAddr := common.BytesToAddress(txLog.Topics[2][:])
		mdl.addNewPool(poolAddr, blockNum, mdl.GetDetailsByKey("MARKET"), priceOracleAddr)
		//
		mdl.poolToMarketAndPriceOracleRelation(blockNum, poolAddr, priceOracleAddr.Hex())
	case core.Topic("RegisterCreditSuite(address,address)"): // pool, cm
		cmAddr := common.BytesToAddress(txLog.Topics[2][:])
		mdl.addCM(blockNum, cmAddr.Hex())
	case core.Topic("SetPriceOracle(address,address)"): // pool, priceoracle
		poolAddr := common.BytesToAddress(txLog.Topics[1][:]).Hex()
		priceOracle := common.BytesToAddress(txLog.Topics[2][:]).Hex()
		mdl.poolToMarketAndPriceOracleRelation(blockNum, poolAddr, priceOracle)
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

// if already present doesn't add.
func (mdl *ContractRegister) addNewPool(addr string, blockNum int64, market string, pOracle common.Address) {
	version := core.FetchVersion(addr, blockNum, mdl.Client)
	switch version {
	case core.NewVersion(1), core.NewVersion(2):
		obj := pool_v2.NewPool(addr, mdl.Client, mdl.Repo, blockNum)
		mdl.Repo.AddSyncAdapter(obj)
	default:
		// add pool to the lmrewards so that farm_v3 table entry can be created.
		if mdl.Repo.GetAdapter(addr) != nil {
			return
		}
		adapters := mdl.Repo.GetAdapterAddressByName(ds.LMRewardsv3)
		lmRewards := mdl.Repo.GetAdapter(adapters[0])
		lmRewards.(*v3.LMRewardsv3).AddPoolv3(blockNum, addr)
		// add pool
		if pOracle == core.NULL_ADDR {
			po, version, err := mdl.Repo.GetActivePriceOracleByBlockNum(blockNum)
			if version == core.NewVersion(300) {
				log.Fatal()
			}
			log.CheckFatal(err)
			pOracle = common.HexToAddress(string(po))
		}
		obj := pool_v3.NewPool(addr, mdl.Client, mdl.Repo, blockNum, market, schemas.PriceOracleT(pOracle.Hex())) // can be 310, 300, 2, 1
		mdl.Repo.AddSyncAdapter(obj)
		//
		mdl.poolToMarketAndPriceOracleRelation(blockNum, addr, pOracle.Hex())
	}
}
