package trace_service

import (
	"testing"

	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/stretchr/testify/require"
)

// doesn't work
func _TestParityFetcher(t *testing.T) {
	rpc := utils.GetEnvOrDefault("GOERLI_ETH_PROVIDER", "")
	if rpc == "" {
		return
	}
	txHash := "0x63c2fee94a94379de941aab4950c51a505eea652c89b9ad3757d1480092fb330"
	fetcher := NewParityFetcher(rpc)
	rpcTrace, err := fetcher.getData(txHash)
	if err != nil {
		t.Fatal(err)
	}
	tenderlyTrace := convertToTenderlyTrace(rpcTrace, txHash)
	input := TenderlySampleTestInput{}
	utils.ReadJsonAndSetInterface("../../inputs/execute_parser_transfers/get_transfers.json", &input)
	require.JSONEq(t, utils.ToJson(input.TenderlyTrace), utils.ToJson(tenderlyTrace))
}
