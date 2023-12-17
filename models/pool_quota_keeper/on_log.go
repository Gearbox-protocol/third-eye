package pool_quota_keeper

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas/schemas_v3"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (mdl PoolQuotaKeeper) updateQuotaDetails(blockNum int64, token string, newDetails *schemas_v3.QuotaDetails) {
	details := mdl.quotas[token]
	if newDetails.Rate != 0 {
		currentTs := mdl.Repo.SetAndGetBlock(blockNum).Timestamp
		details.CumQuotaIndex = (*core.BigInt)(details.GetCumulativeIndexAt(currentTs))
		// update rates and ts later
		details.Timestamp = currentTs
		details.Rate = newDetails.Rate
	}
	if newDetails.IncreaseFee != 0 {
		details.IncreaseFee = newDetails.IncreaseFee
	}
	if newDetails.Limit != nil {
		details.Limit = newDetails.Limit
	}
	details.IsDirty = true
	details.BlockNum = blockNum
	mdl.quotas[token] = details
}

func (mdl PoolQuotaKeeper) addToken(blockNum int64, token string) {
	details := &schemas_v3.QuotaDetails{
		PoolQuotaKeeper: mdl.Address,
		BlockNum:        blockNum,
		Token:           token,
		//
		Timestamp:     mdl.Repo.SetAndGetBlock(blockNum).Timestamp,
		Pool:          mdl.GetDetailsByKey("pool"),
		CumQuotaIndex: (*core.BigInt)(big.NewInt(1)),
		//
		IsDirty: true,
	}
	mdl.quotas[token] = details
}
func (mdl *PoolQuotaKeeper) OnLog(txLog types.Log) {
	blockNum := int64(txLog.BlockNumber)
	mdl.lastBlockNum = blockNum
	switch txLog.Topics[0] {
	case core.Topic("SetQuotaIncreaseFee(address,uint16)"):
		mdl.updateQuotaDetails(blockNum, common.BytesToAddress(txLog.Topics[1][:]).Hex(), &schemas_v3.QuotaDetails{
			IncreaseFee: utils.BytesToUInt16(txLog.Data),
		})
	case core.Topic("SetTokenLimit(address,uint96)"):
		mdl.updateQuotaDetails(blockNum, common.BytesToAddress(txLog.Topics[1][:]).Hex(), &schemas_v3.QuotaDetails{
			Limit: (*core.BigInt)(new(big.Int).SetBytes(txLog.Data)),
		})
	case core.Topic("UpdateTokenQuotaRate(address,uint16)"):
		mdl.updateQuotaDetails(blockNum, common.BytesToAddress(txLog.Topics[1][:]).Hex(), &schemas_v3.QuotaDetails{
			Rate: utils.BytesToUInt16(txLog.Data),
		})
	case core.Topic("AddQuotaToken(address)"):
		mdl.addToken(blockNum, common.BytesToAddress(txLog.Topics[1][:]).Hex())
	case core.Topic("UpdateQuota(address,address,int96)"):
		mdl.mgr.AddAccountQuota(blockNum, txLog)
	}
}

func (mdl PoolQuotaKeeper) GetRepo() ds.RepositoryI {
	return mdl.Repo
}
func (mdl PoolQuotaKeeper) GetQuotas(token string) *schemas_v3.QuotaDetails {
	return mdl.quotas[token]
}
