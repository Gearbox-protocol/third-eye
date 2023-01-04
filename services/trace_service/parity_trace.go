package trace_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"

	"github.com/Gearbox-protocol/sdk-go/log"
)

// https://docs.alchemy.com/reference/trace-api
type ParityFetcher struct {
	rpc    string
	client http.Client
}

func NewParityFetcher(rpc string) *ParityFetcher {
	return &ParityFetcher{
		rpc:    rpc,
		client: http.Client{},
	}
}

type traceResp struct {
	Error struct {
		Code    int64  `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
	Result []RPCTrace
}

// https://docs.alchemy.com/reference/trace-transaction
func (app ParityFetcher) getData(txHash string) ([]RPCTrace, error) {
	format := `{"id":1,"jsonrpc":"2.0","params":["%s"],"method":"trace_transaction"}`
	params := fmt.Sprintf(format, txHash)
	//
	buf := &bytes.Buffer{}
	buf.WriteString(params)
	req, _ := http.NewRequest(http.MethodPost, app.rpc, buf)
	resp, err := app.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("While making request %s", err)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("While reading body from response %s", err)
	}
	traceObj := traceResp{}
	err = json.Unmarshal(data, &traceObj)
	if err != nil {
		return nil, fmt.Errorf("While unmarshaling %s", err)
	}
	if traceObj.Error.Code != 0 {
		return nil, fmt.Errorf(traceObj.Error.Message)
	}
	return traceObj.Result, nil
}

func (app ParityFetcher) getTxTrace(txHash string) *TenderlyTrace {
	rpcTrace, err := app.getData(txHash)
	log.CheckFatal(err)
	return convertToTenderlyTrace(rpcTrace, txHash)
}

type RPCTrace struct {
	Action struct {
		From     string `json:"from"`
		To       string `json:"to"`
		CallType string `json:"callType"`
		Value    string `json:"value"`
		Input    string `json:"input"`
	} `json:"action"`
	BlockNumber  int64 `json:"blockNumber"`
	Subtraces    int   `json:"subtraces"`
	TraceAddress []int `json:"traceAddress"`
}

func convertToTenderlyTrace(old []RPCTrace, txHash string) *TenderlyTrace {
	call, _ := toTenderlyCall(old[0], txHash)
	for _, rpcEntry := range old[1:] {
		nextCall, path := toTenderlyCall(rpcEntry, txHash)
		cur := call
		for _, step := range path[:len(path)-1] {
			cur = cur.Calls[step]
		}
		cur.Calls[path[len(path)-1]] = nextCall
	}
	return &TenderlyTrace{
		CallTrace:   call,
		BlockNumber: old[0].BlockNumber,
		TxHash:      txHash,
	}
}

func toTenderlyCall(old RPCTrace, txHash string) (*Call, []int) {
	callerOp := strings.ToUpper(old.Action.CallType)
	var valueStr string
	if callerOp != "STATICCALL" {
		value, ok := new(big.Int).SetString(old.Action.Value[2:], 16)
		if !ok {
			log.Fatal("For txhash (%s) can't parse ethValue %s", txHash, old.Action.Value)
		}
		valueStr = value.String()
	}
	var calls []*Call
	if old.Subtraces > 0 {
		calls = make([]*Call, old.Subtraces)
	}
	return &Call{
		From:     old.Action.From,
		To:       old.Action.To,
		CallerOp: strings.ToUpper(old.Action.CallType),
		Input:    old.Action.Input,
		Value:    valueStr,
		Calls:    calls,
	}, old.TraceAddress
}
