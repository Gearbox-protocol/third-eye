package pool_common

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

func (g *GatewayHandler) AddRemoveLiqEvent(entry *schemas.PoolLedger) {
	g.addLastEventToArr()
	g.lastEntry = entry
}

func (g *GatewayHandler) addLastEventToArr() {
	if g.lastEntry != nil {
		g.removeLiqEvents = append(g.removeLiqEvents, g.lastEntry)
	}
	g.lastEntry = nil
}

func (g *GatewayHandler) GetRemoveLiqEventsAndClear() []*schemas.PoolLedger {
	g.addLastEventToArr()
	events := g.removeLiqEvents
	g.removeLiqEvents = nil
	return events
}

func (g *GatewayHandler) CheckWithdrawETH(txHash string, blockNum, ind int64, pool, user string) {
	if g.Gateway == core.NULL_ADDR {
		return
	}
	if g.lastEntry != nil && g.lastEntry.BlockNumber == blockNum &&
		g.lastEntry.User == g.Gateway.Hex() && g.lastEntry.Pool == pool {
		g.lastEntry.Executor = g.lastEntry.User // executor in this case is the gateway address for removeLiquidity
		g.lastEntry.User = user
	} else {
		log.Fatalf(`WithdrawalWETH event on gateway@(%d,%d) 
			but no matching last pool Remove Liquidity %s. TxHash: %s`,
			blockNum, ind, utils.ToJson(g.lastEntry), txHash)
	}
}

// utils
func GetPoolGateways(client core.ClientI) map[common.Address]GatewayDetails {
	_chainId, err := client.ChainID(context.Background())
	log.CheckFatal(err)
	chainId := _chainId.Int64()
	if log.GetNetworkName(chainId) == "TEST" {
		return map[common.Address]GatewayDetails{}
	}
	symToAddrStore := core.GetSymToAddrByChainId(chainId)
	if log.GetBaseNet(chainId) == "ARBITRUM" {
		return map[common.Address]GatewayDetails{
			symToAddrStore.Exchanges["GEARBOX_WETH_POOL"]: {
				Gateway: symToAddrStore.Exchanges["WETH_GATEWAY"],
				Sym:     "WETH",
			},
		}
	}
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
