package ds

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/pkg/priceFetcher"
)

var _oneInchUpdater *priceFetcher.OneInchOracle

func SetOneInchUpdater(client core.ClientI, tStore priceFetcher.DecimalStoreI) *priceFetcher.OneInchOracle {
	if _oneInchUpdater == nil {
		_oneInchUpdater = priceFetcher.New1InchOracle(client, tStore, priceFetcher.URLsAndResolve{Resolve: false})
	}
	return _oneInchUpdater
}
func GetOneInchUpdater() *priceFetcher.OneInchOracle {
	return _oneInchUpdater
}
