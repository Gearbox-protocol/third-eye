package rebase_token

import (
	"testing"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/ethclient"
)

func TestRebaseToken(t *testing.T) {
	provider := utils.GetEnvOrDefault("ETH_PROVIDER", "")
	if provider == "" {
		t.Skip()
	}
	log.SetTestLogging(t)
	client := ethclient.NewEthClient(&config.Config{EthProvider: provider})
	stETh := core.GetSymToAddrByChainId(core.GetChainId(client)).Tokens["stETH"]
	mdl := NewRebaseToken(stETh.Hex(), client, ds.DummyRepo{})
	//
	start := mdl.LastSync + 1
	// var endRange int64 = 13855093
	var endRange int64 = 0
	switch core.GetChainId(client) {
	case 1:
		endRange = 17572153
	case 5:
		endRange = 9253954
	}
	//
	logs, err := pkg.Node{Client: client}.GetLogs(start, endRange, mdl.GetAllAddrsForLogs(), mdl.Topics())
	log.CheckFatal(err)
	log.Info(len(logs))
	for _, txLog := range logs {
		// log.Info(txLog.BlockNumber)
		mdl.OnLog(txLog)
	}
	log.Info(ds.Count)
}
