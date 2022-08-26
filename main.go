package main

import (
	"github.com/Gearbox-protocol/sdk-go/artifacts/priceOracle"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/ethclient"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/ac3e2f11fafb420c94cc0a357084fb85")
	log.CheckFatal(err)
	priceOracle.NewPriceOracle(common.HexToAddress(""), client)
	n := core.Node{Client: client}
	txLogs, err := n.GetLogs(0,
		15375264, []common.Address{common.HexToAddress("0x0e74a08443c5E39108520589176Ac12EF65AB080")},
		[][]common.Hash{{core.Topic("NewPriceFeed(address,address)")}})
	log.CheckFatal(err)
	for _, txLog := range txLogs {
		log.Info(txLog.BlockNumber, toAddr(txLog.Topics[1]), toAddr(txLog.Topics[2]))
	}
}

func toAddr(item common.Hash) string {
	return common.BytesToAddress(item[:]).Hex()
}
