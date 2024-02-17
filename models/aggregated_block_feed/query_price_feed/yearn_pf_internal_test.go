package query_price_feed

import (
	"sync"
	"testing"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/ethclient"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/common"
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
	obj := &QueryPriceFeed{
		SyncAdapter: adapter,
		mu:          &sync.Mutex{},
		yearnPFInternal: yearnPFInternal{
			mainPFAddress: common.HexToAddress(adapter.Address),
			version:       core.NewVersion(300),
		}, // main price feed
	}
	pf, err := obj.CalculateYearnPFInternally(18631514)
	log.CheckFatal(err)
	log.Info(utils.ToJson(pf))
}
