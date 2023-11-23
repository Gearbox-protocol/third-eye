package ds

import (
	"encoding/hex"
	"strings"
	"testing"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
)

type ExecuteParserLenTester struct {
	TxHash string                      `json:"txHash"`
	Calls  []string                    `json:"calls"`
	Events []*schemas.AccountOperation `json:"events"`
}

func (tester ExecuteParserLenTester) GetCalls(t *testing.T) FacadeCallNameWithMulticall {
	var multicalls []multicall.Multicall2Call
	for _, call := range tester.Calls {
		_s := strings.Split(call, "@")
		bytes, err := hex.DecodeString(_s[1])
		if err != nil {
			t.Fatal(err)
		}
		multicalls = append(multicalls, multicall.Multicall2Call{
			Target:   common.HexToAddress(_s[0]),
			CallData: bytes,
		})
	}
	return FacadeCallNameWithMulticall{
		Name:       "test",
		multiCalls: multicalls,
	}
}

// checks failed tokeDisabled call
func Test_Check1(t *testing.T) {
	data := ExecuteParserLenTester{}
	utils.ReadJsonAndSetInterface("execute_parser/check_failed_token_disabled.json", &data)

	calls := data.GetCalls(t)
	if !calls.SameMulticallLenAsEvents(data.Events) {
		t.Fatalf("expected %d multicalls, but third-eye detected %d. Events: %s. Calls: %s. txhash: %s",
			calls.LenOfMulticalls(), len(data.Events),
			utils.ToJson(data.Events), calls.String(), data.TxHash)
	}
}

// checks if the events len is zero, can func handle it?
func Test_Check2(t *testing.T) {
	data := ExecuteParserLenTester{}
	utils.ReadJsonAndSetInterface("execute_parser/check_event_len_0.json", &data)

	calls := data.GetCalls(t)
	if !calls.SameMulticallLenAsEvents(data.Events) {
		t.Fatalf("expected %d multicalls, but third-eye detected %d. Events: %s. Calls: %s. txhash: %s",
			calls.LenOfMulticalls(), len(data.Events),
			utils.ToJson(data.Events), calls.String(), data.TxHash)
	}
}
