package debts

type caValueAndAvailLiquidity struct {
	caTotalValue            float64
	totalAvailableLiquidity float64
	expectedLiq             float64
}

type MarketToTvl map[string]*caValueAndAvailLiquidity

func (ds *MarketToTvl) add(pool string, caValue float64, availLiq float64, expectedLiq float64) {
	if x := (*ds)[pool]; x != nil {
		x.caTotalValue += caValue
		x.totalAvailableLiquidity += availLiq
		x.expectedLiq += expectedLiq
	} else {
		(*ds)[pool] = &caValueAndAvailLiquidity{
			caTotalValue:            caValue,
			totalAvailableLiquidity: availLiq,
			expectedLiq:             expectedLiq,
		}
	}
}
