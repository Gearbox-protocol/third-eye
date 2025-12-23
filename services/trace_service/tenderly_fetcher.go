package trace_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/common"
)

type Call struct {
	From     string  `json:"from"`
	To       string  `json:"to"`
	CallerOp string  `json:"caller_op"`
	Input    string  `json:"input"`
	Value    string  `json:"value"`
	Calls    []*Call `json:"calls"`
}

//

type RawLog struct {
	Address common.Address `json:"address"`
	Topics  []common.Hash  `json:"topics"`
	Data    string         `json:"data"`
}
type Log struct {
	Name string `json:"name"`
	Raw  RawLog `json:"raw"`
}

type TenderlyTrace struct {
	CallTrace   *Call  `json:"call_trace"`
	TxHash      string `json:"transaction_id"`
	Logs        []Log  `json:"logs"`
	BlockNumber int64  `json:"block_number"`
}

func readAndGetReader(a io.Reader) ([]byte, error) {
	b := bytes.NewBuffer(nil)
	_, err := b.ReadFrom(a)
	if err != nil {
		return nil, err
	}
	str := b.Bytes()
	return str, nil
}
func (ep *TenderlyFetcher) getTenderly(txHash string) (*TenderlyTrace, error) {
	link := fmt.Sprintf("https://api.tenderly.co/api/v1/public-contract/%d/trace/%s", ep.ChainId, txHash)
	return ep.getUrl(link)
}
func (ep *TenderlyFetcher) getProxy(txHash string) (*TenderlyTrace, error) {
	link := fmt.Sprintf("https://testnet.gearbox.foundation/redstone/tenderly/%d/%s", ep.ChainId, txHash)
	return ep.getUrl(link)
}

func (ep *TenderlyFetcher) getUrl(link string) (*TenderlyTrace, error) {
	req, _ := http.NewRequest(http.MethodGet, link, nil)
	resp, err := ep.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	msg, err := readAndGetReader(resp.Body)
	log.CheckFatal(err)
	trace := &TenderlyTrace{}

	err = json.Unmarshal(msg, trace)
	if err != nil {
		return nil, err
	}
	return trace, nil
}

func (ep *TenderlyFetcher) getTxTrace(txHash string) (*TenderlyTrace, error) {
	trace, err := ep.getTenderly(txHash)
	if err != nil {
		trace, err = ep.getProxy(txHash)
		if err != nil {
			return nil, log.WrapErrWithLine(err)
		}
	}
	if trace.CallTrace == nil {
		log.Info("Call trace nil retrying in 30 sec")
		time.Sleep(30 * time.Second)
		trace, err = ep.getTenderly(txHash)
		if err != nil {
			return nil, log.WrapErrWithLine(err)
		}
		if trace.CallTrace == nil {
			log.Fatal("Retry failed for tenderly: ", txHash)
		}
		return trace, nil
	}
	return trace, nil
}

type TenderlyFetcher struct {
	Client  http.Client
	ChainId int64
}

func NewTenderlyFetcher(chainId int64) TenderlyFetcher {
	return TenderlyFetcher{
		ChainId: chainId,
		Client:  http.Client{},
	}
}

// Tenderly test

type TenderlySampleTestInput struct {
	TenderlyTrace   *TenderlyTrace   `json:"callTrace"`
	Account         string           `json:"account"`
	UnderlyingToken string           `json:"underlyingToken"`
	Users           ds.BorrowerAndTo `json:"users"`
}

///////////////////////////
// Fetcher
///////////////////////////

type InternalFetcher struct {
	txLogger         TxLogger
	parityFetcher    *ParityFetcher
	tenderlyFetcher  TenderlyFetcher
	useTenderlyTrace bool
}

func NewInternalFetcher(cfg *config.Config, client core.ClientI) InternalFetcher {
	fetcher := InternalFetcher{
		txLogger:         NewTxLogger(client, cfg.BatchSizeForHistory),
		parityFetcher:    NewParityFetcher(cfg.EthProvider),
		tenderlyFetcher:  NewTenderlyFetcher(core.GetChainId(client)),
		useTenderlyTrace: cfg.UseTenderlyTrace,
	}
	fetcher.check()
	return fetcher
}

const alchemyTraceTransactionExpectedErr = "invalid argument 0: hex string has length 0, want 64 for common.Hash"
const anvilTraceTransactionExpectedErr = "invalid length 0, expected a (both 0x-prefixed or not) hex string or byte array containing 32 bytes"

func (ep InternalFetcher) check() {
	if !ep.useTenderlyTrace {
		_, err := ep.parityFetcher.getData("")
		endpointsSupportTrace := 0
		if err != nil {
			for _, err := range err.(utils.Errors) {
				if strings.Contains(err.Error(), alchemyTraceTransactionExpectedErr) || // on alchemy
					strings.Contains(err.Error(), anvilTraceTransactionExpectedErr) { // anvil {
					endpointsSupportTrace += 1
				}
			}
		}
		log.Info("Endpoints that support parity trace_transaction: ", endpointsSupportTrace)
	}
}

// {"method":"debug_traceTransaction","params":["%s", {"tracer": "callTracer"}], "id":1,"jsonrpc":"2.0"} for quicknode, like monad, etherlink and somnia
// {"id":1,"jsonrpc":"2.0","params":["%s"],"method":"trace_transaction"} on alchemy, rpcs
// tenderly for others
func (ep InternalFetcher) GetTxTrace(txHash string, canLoadLogsFromRPC bool) *TenderlyTrace {
	var trace *TenderlyTrace
	if ep.useTenderlyTrace {
		var err error
		trace, err = ep.tenderlyFetcher.getTxTrace(txHash)
		log.CheckFatal(err)
	} else if traceUrl := utils.GetEnvOrDefault("TRACE_URL", ""); traceUrl != "" {
		var err error
		trace, err = ep.parityFetcher.getTxTraceQuickNode(txHash, traceUrl)
		// log.Fatal(utils.ToJson(trace))
		if err != nil {
			log.Info("fallback on tenderly due to", err, " for ", txHash)
			trace, err = ep.tenderlyFetcher.getTxTrace(txHash)
		}
	} else {
		var err error
		trace, err = ep.parityFetcher.getTxTrace(txHash)
		if err != nil {
			log.Info("fallback on tenderly due to", err, " for ", txHash)
			trace, err = ep.tenderlyFetcher.getTxTrace(txHash)
		}
	}
	//
	if canLoadLogsFromRPC && len(trace.Logs) == 0 {
		trace.Logs = ep.txLogger.GetLogs(int(trace.BlockNumber), trace.TxHash)
	}
	return trace
}
