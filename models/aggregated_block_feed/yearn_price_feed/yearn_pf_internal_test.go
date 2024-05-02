package yearn_price_feed

import (
	"testing"

	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/ethclient"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
)

func TestInternalyearnPrice(t *testing.T) {
	t.Skip()
	client, err := ethclient.Dial("")
	log.CheckFatal(err)
	adapter := &ds.SyncAdapter{
		SyncAdapterSchema: &schemas.SyncAdapterSchema{
			Contract: &schemas.Contract{
				Address: "0xCbE1965888aCd7799Ecf59E8E8D184A7Abed9fc4",
				Client:  client,
			},
		},
	}
	obj := NewYearnPriceFeedFromAdapter(adapter)
	pf, err := obj.CalculateYearnPFInternally(18631514)
	log.CheckFatal(err)
	log.Info(utils.ToJson(pf))
}
