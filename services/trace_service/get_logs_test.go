package trace_service

import (
	"testing"

	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func TestInsertInSlice(t *testing.T) {
	a := make([]int, 0, 5)
	a = append(a, []int{2, 3, 6}...)

	a = insertInSlice(a, 0)
	if len(a) != 4 || cap(a) != 5 || !cmp.Equal(a, []int{0, 2, 3, 6}) {
		t.Fatal(a, len(a), cap(a))
	}
	a = insertInSlice(a, 7)
	if len(a) != 5 || cap(a) != 5 || !cmp.Equal(a, []int{0, 2, 3, 6, 7}) {
		t.Fatal(a, len(a), cap(a))
	}

	a = insertInSlice(a, 5)
	if len(a) != 6 || !cmp.Equal(a, []int{0, 2, 3, 5, 6, 7}) {
		t.Fatal(a, len(a), cap(a))
	}
	a = insertInSlice(a, 8)
	if len(a) != 7 || !cmp.Equal(a, []int{0, 2, 3, 5, 6, 7, 8}) {
		t.Fatal(a, len(a), cap(a))
	}
}

func TestDeleteInSlice(t *testing.T) {
	a := make([]int, 0, 5)
	a = append(a, []int{0, 2, 3, 6, 7}...)

	a = deleteInSlice(a, 5, map[int]map[string][]Log{})

	if len(a) != 2 || cap(a) != 5 || !cmp.Equal(a, []int{6, 7}) {
		t.Fatal(a, len(a), cap(a))
	}
}

func TestTxLogger(t *testing.T) {
	// create eth client rpc
	url := utils.GetEnvOrDefault("GOERLI_ETH_PROVIDER", "")
	if url == "" {
		return
	}
	client := ethclient.NewEthClient(&config.Config{EthProvider: url})
	fetcher := NewInternalFetcher(&config.Config{BatchSizeForHistory: 10, UseTenderlyTrace: true}, client)
	// create other variables
	input := TenderlySampleTestInput{}
	utils.ReadJsonAndSetInterface("../../inputs/execute_parser_transfers/get_transfers.json", &input)
	trace := input.TenderlyTrace
	// check 1
	logs := fetcher.txLogger.GetLogs(int(trace.BlockNumber), trace.TxHash)
	require.JSONEq(t, utils.ToJson(trace.Logs), utils.ToJson(logs))
}
