package pool_common

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
)

func CheckIfAmountMoreThan1Mil(client core.ClientI, repo ds.RepositoryI, state *schemas.PoolState, amount *big.Int, blockNum int64, txHash string, operation string) {
	token := state.UnderlyingToken
	priceInUSD := repo.GetPrice(token)
	if priceInUSD == nil {
		return
	}
	value := utils.GetFloat64Decimal(new(big.Int).Mul(priceInUSD, amount), repo.GetToken(token).Decimals+8)
	if value > 1_000_000 {
		urls := core.NetworkUIUrl(core.GetChainId(client))
		repo.RecentMsgf(log.RiskHeader{
			BlockNumber: blockNum,
			EventCode:   "AMQP",
		}, "Pool %s in %s is more than 1Million USD, calculated value is %f", operation, urls.ExplorerHashUrl(txHash), value)
	}
}
