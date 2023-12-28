package ds

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/pkg/priceFetcher"
)

var _oneInchUpdater *priceFetcher.OneInchOracle

func SetOneInchUpdater(client core.ClientI, tStore priceFetcher.DecimalStoreI) *priceFetcher.OneInchOracle {
	if _oneInchUpdater == nil {
		_oneInchUpdater = priceFetcher.New1InchOracle(client, core.GetChainId(client), tStore)
	}
	return _oneInchUpdater
}
func GetOneInchUpdater() *priceFetcher.OneInchOracle {
	return _oneInchUpdater
}
