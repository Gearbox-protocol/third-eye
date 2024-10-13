package market_configurator

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/models/credit_manager/cm_v3"
	"github.com/Gearbox-protocol/third-eye/models/pool/pool_v3"
	"github.com/Gearbox-protocol/third-eye/models/price_oracle/po_v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (mdl *MarketConfigurator) OnLog(txLog types.Log) {
	blockNum := int64(txLog.BlockNumber)
	switch txLog.Topics[0] {
	case core.Topic("CreateMarket(address,address,string,string)"):
		// details := mdl.GetDetails()
		// if val := details["legacy"]; val == nil {
		// 	details["legacy"] = "set"
		// 	mdl.Details = details
		// 	return
		// } // first one is the legacy one is no need to set that.
		poolAddr := common.BytesToAddress(txLog.Topics[1][:]).Hex()
		poolAdapter := mdl.Repo.GetAdapter(poolAddr)
		priceOracleAddr := mdl.GetPriceOracle(poolAddr, blockNum)
		mdl.Repo.AddRelation(&schemas.Relation{
			Type:"MarketPool",
			Owner: mdl.Address,
			Dependent: poolAddr,
			BlockNum: blockNum,
		})
		mdl.Repo.AddRelation(&schemas.Relation{
			Type:"PoolOracle",
			Owner: poolAddr,
			Dependent: priceOracleAddr,
			BlockNum: blockNum,
		})
		if priceOracle := mdl.Repo.GetAdapter(priceOracleAddr); priceOracle == nil {
			po := po_v3.NewPriceOracle(priceOracleAddr, blockNum, mdl.Client, mdl.Repo)
			mdl.Repo.AddSyncAdapter(po)
		}
		if poolAdapter == nil {
			pool := pool_v3.NewPool(poolAddr, mdl.Client, mdl.Repo, blockNum, mdl.Address, schemas.PriceOracleT(priceOracleAddr), core.NewVersion(310))
			mdl.Repo.AddSyncAdapter(pool)
		} else {
			if poolv3, ok := poolAdapter.(*pool_v3.Poolv3); ok {
				poolv3.State.Market = mdl.Address
				poolv3.State.PriceOracle = schemas.PriceOracleT(priceOracleAddr)
			} else {
				log.Fatalf("the pool by the address(%s) is not poolv3", poolAddr)
			}
		}
	case core.Topic("CreateCreditManager(address)"):
		cm :=cm_v3.NewCMv3(common.BytesToAddress(txLog.Topics[1][:]).Hex(), mdl.Client, mdl.Repo, blockNum)
		mdl.Repo.AddSyncAdapter(cm)
	}
}

func (mdl *MarketConfigurator) GetPriceOracle(pool string, blockNum int64) string {
	hash :=common.HexToHash(pool)
	data , err:=core.CallFuncWithExtraBytes(mdl.Client, "01374518", common.HexToAddress(mdl.Address) ,blockNum, hash[:])
	log.CheckFatal(err)
	return  common.BytesToAddress(data).Hex()
}