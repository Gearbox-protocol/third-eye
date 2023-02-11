package pool

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

// 2 funcs
// - storing all removeLiqEvents for a block
// - fix using gateway WithDrawWETH user address in RemoveLiquidity event

// when liquidity is removed from the pool, like WETH and if there is a gateway then the user in RemoveLiquidity Event is the gateway address, for getting correct user we need to check the WithdrawalWETH event
type GatewayHandler struct {
	lastEntry   *schemas.PoolLedger
	gatewayAddr common.Address
	// for calculating the remove liq amount in underlying instead of diesel token amount which is present in the emitted removeLiquidity event
	removeLiqEvents []*schemas.PoolLedger
}

func NewGatewayHandler(gatewayAddr common.Address) GatewayHandler {
	return GatewayHandler{gatewayAddr: gatewayAddr}
}

func (g *GatewayHandler) addRemoveLiqEvent(entry *schemas.PoolLedger) {
	g.addLastEventToArr()
	g.lastEntry = entry
}

func (g *GatewayHandler) addLastEventToArr() {
	if g.lastEntry != nil {
		g.removeLiqEvents = append(g.removeLiqEvents, g.lastEntry)
	}
	g.lastEntry = nil
}

func (g *GatewayHandler) getRemoveLiqEventsAndClear() []*schemas.PoolLedger {
	g.addLastEventToArr()
	events := g.removeLiqEvents
	g.removeLiqEvents = nil
	return events
}

func (g *GatewayHandler) checkWithdrawETH(blockNum, ind int64, pool, user string) {
	if g.gatewayAddr == core.NULL_ADDR {
		return
	}
	if g.lastEntry != nil && g.lastEntry.BlockNumber == blockNum &&
		g.lastEntry.User == g.gatewayAddr.Hex() && g.lastEntry.Pool == pool {
		g.lastEntry.User = user
	}
	log.Warn("WithdrawalWETH event on gateway@(%d,%d) but no matching last pool Remove Liquidity %s", blockNum, ind, utils.ToJson(g.lastEntry))
}
