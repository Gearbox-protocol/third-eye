package getter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/common"
)

//
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

func (ep *TenderlyFetcher) getTenderlyData(txHash string) (*TenderlyTrace, error) {
	link := fmt.Sprintf("https://api.tenderly.co/api/v1/public-contract/%d/trace/%s", ep.ChainId, txHash)
	req, _ := http.NewRequest(http.MethodGet, link, nil)
	resp, err := ep.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	trace := &TenderlyTrace{}
	err = json.Unmarshal(body, trace)
	if err != nil {
		log.Fatal(err, " for ", txHash)
		return nil, err
	}
	return trace, nil
}

func (ep *TenderlyFetcher) _getTxTrace(txHash string) *TenderlyTrace {
	trace, err := ep.getTenderlyData(txHash)
	if err != nil {
		log.Fatal(err)
	}
	if trace.CallTrace == nil {
		log.Info("Call trace nil retrying in 30 sec")
		time.Sleep(30 * time.Second)
		trace, err = ep.getTenderlyData(txHash)
		if err != nil {
			log.Fatal(err)
		}
		if trace.CallTrace == nil {
			log.Fatal("Retry failed for tenderly: ", txHash)
		}
		return trace
	}
	return trace
}

type TenderlyFetcher struct {
	Client  http.Client
	ChainId int64
}

func NewTenderlyFetcher(client core.ClientI, chainId int64) TenderlyFetcher {
	return TenderlyFetcher{
		ChainId: chainId,
		Client:  http.Client{},
	}
}

// Tenderly test

type TenderlySampleTestInput struct {
	CallTrace       *TenderlyTrace   `json:"callTrace"`
	Account         string           `json:"account"`
	UnderlyingToken string           `json:"underlyingToken"`
	Users           ds.BorrowerAndTo `json:"users"`
}

///////////////////////////
// Fetcher
///////////////////////////

type InternalFetcher struct {
	txLogger             TxLogger
	internalCallsFetcher *InternalCallFetcher
	tenderlyFetcher      TenderlyFetcher
}

func NewInternalFetcher(cfg *config.Config, client core.ClientI) InternalFetcher {
	return InternalFetcher{
		txLogger:             NewTxLogger(client, cfg.BatchSizeForHistory),
		internalCallsFetcher: NewInternalCallFetcher(cfg.EthProvider),
		tenderlyFetcher:      NewTenderlyFetcher(client, cfg.ChainId),
	}
}

func (ep InternalFetcher) GetTxTrace(txHash string, canLoadLogsFromRPC bool) *TenderlyTrace {
	trace := ep.tenderlyFetcher._getTxTrace(txHash)
	if canLoadLogsFromRPC && len(trace.Logs) == 0 {
		trace.Logs = ep.txLogger.GetLogs(int(trace.BlockNumber), trace.TxHash)
	}
	return trace
}
