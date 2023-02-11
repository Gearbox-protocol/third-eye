package pool

import (
	"context"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
)

type GatewayDetails struct {
	Gateway    common.Address
	Token      common.Address
	Sym        string
	UserCantBe common.Address
}

// 2 funcs
// - storing all removeLiqEvents for a block
// - fix using gateway WithDrawWETH user address in RemoveLiquidity event

// when liquidity is removed from the pool, like WETH and if there is a gateway then the user in RemoveLiquidity Event is the gateway address, for getting correct user we need to check the WithdrawalWETH event
type GatewayHandler struct {
	lastEntry *schemas.PoolLedger
	GatewayDetails
	// for calculating the remove liq amount in underlying instead of diesel token amount which is present in the emitted removeLiquidity event
	removeLiqEvents []*schemas.PoolLedger
}

func NewGatewayHandler(details GatewayDetails) GatewayHandler {
	return GatewayHandler{GatewayDetails: details}
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
	if g.Gateway == core.NULL_ADDR {
		return
	}
	if g.lastEntry != nil && g.lastEntry.BlockNumber == blockNum &&
		g.lastEntry.User == g.Gateway.Hex() && g.lastEntry.Pool == pool {
		g.lastEntry.User = user
	} else {
		log.Fatalf(`WithdrawalWETH event on gateway@(%d,%d) 
			but no matching last pool Remove Liquidity %s`,
			blockNum, ind, utils.ToJson(g.lastEntry))
	}
}

// utils
func GetPoolGateways(client core.ClientI) map[common.Address]GatewayDetails {
	_chainId, err := client.ChainID(context.Background())
	log.CheckFatal(err)
	chainId := _chainId.Int64()
	if !(chainId == 1 || chainId == 5) {
		return map[common.Address]GatewayDetails{}
	}
	symToAddrStore := core.GetSymToAddrByChainId(chainId)
	return map[common.Address]GatewayDetails{
		symToAddrStore.Exchanges["GEARBOX_WETH_POOL"]: {
			Gateway: symToAddrStore.Exchanges["WETH_GATEWAY"],
			Sym:     "WETH",
		},
		symToAddrStore.Exchanges["GEARBOX_WSTETH_POOL"]: {
			Gateway:    symToAddrStore.Exchanges["WSTETH_GATEWAY"],
			Token:      symToAddrStore.Tokens["stETH"],
			UserCantBe: symToAddrStore.Tokens["wstETH"],
			Sym:        "WSTETH",
		},
	}
}
