package pool_keeper

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas/schemas_v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func bytesToUInt16(data []byte) uint16 {
	return uint16(new(big.Int).SetBytes(data).Int64())
}
func (mdl PoolKeeper) updateQuotaDetails(token string, newDetails *schemas_v3.QuotaDetails) {
	if mdl.quotas[token] == nil {
		mdl.quotas[token] = &schemas_v3.QuotaDetails{
			PoolKeeper: mdl.Address,
			Pool:       mdl.GetDetailsByKey("pool"),
			Token:      token,
		}
	}
	details := mdl.quotas[token]
	if newDetails.Rate != 0 {
		details.Rate = newDetails.Rate
	}
	if newDetails.IncreaseFee != 0 {
		details.IncreaseFee = newDetails.IncreaseFee
	}
	if newDetails.Limit != nil {
		details.Limit = newDetails.Limit
	}
	details.IsDirty = true
}
func (mdl PoolKeeper) OnLog(txLog types.Log) {
	blockNum := int64(txLog.BlockNumber)
	mdl.lastBlockNum = blockNum
	switch txLog.Topics[0] {
	case core.Topic("SetQuotaIncreaseFee(address,uint16)"):
		mdl.updateQuotaDetails(common.BytesToAddress(txLog.Topics[1][:]).Hex(), &schemas_v3.QuotaDetails{
			IncreaseFee: bytesToUInt16(txLog.Data),
		})
	case core.Topic("SetTokenLimit(address,uint96)"):
		mdl.updateQuotaDetails(common.BytesToAddress(txLog.Topics[1][:]).Hex(), &schemas_v3.QuotaDetails{
			Limit: (*core.BigInt)(new(big.Int).SetBytes(txLog.Data)),
		})
	case core.Topic("UpdateTokenQuotaRate(address,uint16)"):
		mdl.updateQuotaDetails(common.BytesToAddress(txLog.Topics[1][:]).Hex(), &schemas_v3.QuotaDetails{
			Rate: bytesToUInt16(txLog.Data),
		})
	}
}
