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
	tradingTokens := []string{"WETH", "stETH", "DAI", "SUSD", "FRAX", "GUSD", "LUSD", "WBTC", "USDC", "USDT", "LQTY", "LDO", "CVX", "FXS", "CRV", "SNX"}
	tradingTokensMap := map[string]bool{}
	syms := core.GetSymToAddrByChainId(chainId)
	for _, tradingSym := range tradingTokens {
		if addr, ok := syms.Tokens[tradingSym]; ok {
			tradingTokensMap[addr.Hex()] = true
		} else {
			log.Fatalf("Trading token(%s) for tf_index missing from sdk config:", tradingSym, addr)
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
			if session.Version == 1 {
				priceDecimals = 18
			}
			farmingVal += balance.F * utils.GetFloat64Decimal(priceStore.GetPrices(token, session.Version), priceDecimals)
		}
	}
	if session.Version == 1 {
		farmingVal = farmingVal / utils.GetFloat64Decimal(priceStore.GetPrices(calc.usdc, session.Version), 18) // convert to usd
		// by dividing by usdc price in eth
	}
	// farming val is zero for closed accounts
	debt.FarmingValUSD = farmingVal
}
