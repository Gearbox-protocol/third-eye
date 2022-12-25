package services

import (
	"testing"

	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/ds"
)

func TestGetMainEventLogs(t *testing.T) {
	ep := NewExecuteParser(&config.Config{ChainId: 42}, nil)
	actionWithMulticall := ep.GetMainCalls("0xfbbfbca8d6300adc20c1fd9b2bf2074a14cad0873ad5ed8492ef226861f7c0cc", "0x5aacdab79aa2d30f4242898ecdafda2ed2216db2")
	if len(actionWithMulticall) != 1 || actionWithMulticall[0].Name != "FacadeOpenMulticall" || actionWithMulticall[0].LenOfMulticalls() != 1 {
		log.Fatal(utils.ToJson(actionWithMulticall))
	}
}

type getTransferTestInput struct {
	CallTrace       *TxTrace         `json:"callTrace"`
	Account         string           `json:"account"`
	UnderlyingToken string           `json:"underlyingToken"`
	Users           ds.BorrowerAndTo `json:"users"`
}

func TestGetTransfers(t *testing.T) {

	// create other variables
	input := getTransferTestInput{}
	utils.ReadJsonAndSetInterface("../inputs/execute_parser_transfers/get_transfers.json", &input)

	transfers := getCloseAccountv2Transfers(input.CallTrace, input.Account, input.UnderlyingToken, input.Users)
	if len(transfers) != 1 || transfers[input.UnderlyingToken].String() != "1999963055379350458" {
		t.Fatal(utils.ToJson(transfers))
	}
}
