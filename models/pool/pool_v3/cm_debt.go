package pool_v3

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

type ManageDebtType string

const (
	INCREASE_DEBT ManageDebtType = "IncreaseDebt(address,uint256)"
	DECREASE_DEBT ManageDebtType = "DecreaseDebt(address,uint256)"
)

type ManageDebt struct {
	BlockNum      int64
	TxHash        string
	Amount        *big.Int
	Account       string
	Type          ManageDebtType
	CreditManager string
	LogId         uint
}
type CMDebtHandler struct {
	// txhash
	events map[string][]ManageDebt
	v310   bool
}

func NewCMDebtHandler(v310 bool) CMDebtHandler {
	return CMDebtHandler{
		events: make(map[string][]ManageDebt),
		v310:   v310,
	}
}

func (handler *CMDebtHandler) AddDebt(e ManageDebt) {
	if !handler.v310 {
		return
	}
	handler.events[e.TxHash] = append(handler.events[e.TxHash], e)
}

func (handler *CMDebtHandler) Get(txHash string, cm string, lastLogId uint) []ManageDebt {
	events := handler.events[txHash]
	cmevents, others := []ManageDebt{}, []ManageDebt{}
	for _, entry := range events {
		if entry.CreditManager == cm && entry.LogId <= lastLogId {
			cmevents = append(cmevents, entry)
		} else {
			others = append(others, entry)
		}
	}
	handler.events[txHash] = others
	if len(handler.events[txHash]) == 0 {
		delete(handler.events, txHash)
	}
	return cmevents
}

func (handler CMDebtHandler) CheckIsClean() {
	if len(handler.events) != 0 {
		log.Fatal("events should be processed for v310 pool as creditmanager doesn't emit increase/decrease debt", utils.ToJson(handler.events))
	}
}
