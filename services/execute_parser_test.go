package services

import (
	"context"
	"math/big"
	"testing"

	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/test"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/services/trace_service"
)

type TestClient struct {
	test.TestClient
	chainId int64
}

func (c TestClient) ChainID(ctx context.Context) (*big.Int, error) {
	return big.NewInt(c.chainId), nil
}
func TestGetMainEventLogs(t *testing.T) {
	t.Skip()
	ep := NewExecuteParser(&config.Config{UseTenderlyTrace: true}, &TestClient{chainId: 42})
	actionWithMulticall := ep.GetMainCalls("0xfbbfbca8d6300adc20c1fd9b2bf2074a14cad0873ad5ed8492ef226861f7c0cc", "0x5aacdab79aa2d30f4242898ecdafda2ed2216db2")
	if len(actionWithMulticall) != 1 || actionWithMulticall[0].Name != "FacadeOpenMulticall" || actionWithMulticall[0].LenOfMulticalls() != 1 {
		log.Fatal(utils.ToJson(actionWithMulticall))
	}
}

func TestGetTransfers(t *testing.T) {
	// create other variables
	input := trace_service.TenderlySampleTestInput{}
	utils.ReadJsonAndSetInterface("../inputs/execute_parser_transfers/get_transfers.json", &input)

	transfers := getCloseAccountv2Transfers(input.TenderlyTrace, input.Account, input.UnderlyingToken, input.Users)
	if len(transfers) != 1 || transfers[input.UnderlyingToken].String() != "1999963055379350458" {
		t.Fatal(utils.ToJson(transfers))
	}
}

func TestEmptyMulticall(t *testing.T) {
	t.Skip()
	ep := NewExecuteParser(&config.Config{EthProvider: ""}, &TestClient{chainId: 42})
	actionWithMulticall := ep.GetMainCalls("0xbac6102b402ce8d3afc6565308312bc0ca0bf626bb66d15bfeff9e3abf680452", "0xf6f4F24ae50206A07B8B32629AeB6cb1837d854F")
	t.Fatal(utils.ToJson(actionWithMulticall))
}
