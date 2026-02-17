package services

import (
	"context"
	"math/big"
	"testing"

	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/test"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/services/trace_service"
	"github.com/ethereum/go-ethereum/common"
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

func TestFile(t *testing.T) {
	paramsList := []ds.ExecuteParams{{
		SessionId:     "0x95AaAaaFB5132DAeEB30EE6c4F973e600428C763_54281036_78",
		Protocol:      common.HexToAddress("0x942644106b073e30d72c2c5d7529d5c296ea91ab"),
		CreditAccount: common.HexToAddress("0x95aaaaafb5132daeeb30ee6c4f973e600428c763"),
		Borrower:      common.HexToAddress("0xc6eb0bdc91b93f572f44018551780ae7665b647e"),
		Index:         9,
		BlockNumber:   54637760,
	}}
	creditManagerAddr := "0xe756919CC2e2B6E844a45dBBaCf566B85cB928ab"
	filter := ExecuteFilter{paramsList: paramsList,
		creditManager: common.HexToAddress(creditManagerAddr),
	}
	// log.Info(utils.ToJson(trace.CallTrace))
	// CallTrace :=
	call := &trace_service.Call{}
	err := utils.ReadJsonAndSetInterface("a.json", call)
	log.CheckFatal(err)
	calls := filter.getExecuteCalls(call)
	log.Info(utils.ToJson(calls))
}
