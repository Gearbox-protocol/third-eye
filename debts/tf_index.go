package debts

import (
	"github.com/Gearbox-protocol/sdk-go/calc"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

type FarmingCalculator struct {
	tradingTokensMap map[string]bool
	usdc             string
	testing          bool
}

func NewFarmingCalculator(chainId int64, testing bool) *FarmingCalculator {
	tradingTokens := core.AllTradingSymbolForDBWithW()
	//
	tradingTokensMap := map[string]bool{}
	syms := core.GetSymToAddrByChainId(chainId)
	for _, tradingSym := range tradingTokens {
		if addr, ok := syms.Tokens[string(tradingSym)]; ok {
			tradingTokensMap[addr.Hex()] = true
		} else if (tradingSym == "stETH" || tradingSym == "GUSD") && log.GetBaseNet(chainId) == "ARBITRUM" { // these tokens are not present on arbitrum
			// ignore
		} else if utils.Contains([]core.Symbol{"stETH", "GUSD", "LQTY", "CVX", "yvUSDC", "GMX", "ARB", "MKR", "sDAI", "APE"}, tradingSym) && log.GetBaseNet(chainId) == "OPTIMISM" {
		} else {
			log.Warnf("Trading token(%s) for tf_index missing from sdk config: %s", tradingSym, addr)
		}
	}
	return &FarmingCalculator{
		tradingTokensMap: tradingTokensMap,
		usdc:             syms.Tokens["USDC"].Hex(),
		testing:          testing,
	}
}

func (calc FarmingCalculator) addFarmingVal(debt *schemas.Debt, session *schemas.CreditSession, css *schemas.CreditSessionSnapshot, priceStore calc.TokenDetailsForCalcI) {
	if calc.testing || session.Status != schemas.Active {
		return
	}
	var farmingVal float64 = 0
	for token, balance := range *css.Balances {
		if balance.IsEnabled && balance.HasBalanceMoreThanOne() && !calc.tradingTokensMap[token] {
			var priceDecimals int8 = 8
			if session.Version.Eq(1) {
				priceDecimals = 18
			}
			farmingVal += balance.F * utils.GetFloat64Decimal(priceStore.GetPriceOnBlock(session.CreditManager, token, session.Version), priceDecimals)
		}
	}
	if session.Version.Eq(1) {
		farmingVal = farmingVal / utils.GetFloat64Decimal(priceStore.GetPriceOnBlock(session.CreditManager, calc.usdc, session.Version), 18) // convert to usd
		// by dividing by usdc price in eth
	}
	// farming val is zero for closed accounts
	debt.FarmingValUSD = farmingVal
}
