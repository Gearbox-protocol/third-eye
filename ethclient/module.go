package ethclient

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/ethclient"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/config"
	"go.uber.org/fx"
)

func NewEthClient(config *config.Config) core.ClientI {
	client, err := ethclient.Dial(config.EthProvider)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

var Module = fx.Option(
	fx.Provide(NewEthClient))
