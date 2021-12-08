package core

import (
	"math/big"
)

type EventBalance struct {
	BorrowedAmount *big.Int
	Transfers      Transfers
	SessionId      string
	BlockNumber    int64
	Index          int64
	Clear          bool
	CreditManager  string
	Borrower       string
}

func NewEventBalance(blockNum int64, index uint, sessionId string, borrowedAmount *big.Int, transfers Transfers, clear bool, cm string) EventBalance {
	return EventBalance{
		BorrowedAmount: borrowedAmount,
		Transfers:      transfers,
		SessionId:      sessionId,
		BlockNumber:    int64(blockNum),
		Index:          int64(index),
		Clear:          clear,
		CreditManager:  cm,
	}
}

// sort event balances by block number/log id
type SortedEventbalances []*EventBalance

func (ts SortedEventbalances) Len() int {
	return len(ts)
}
func (ts SortedEventbalances) Swap(i, j int) {
	ts[i], ts[j] = ts[j], ts[i]
}

// sort in increasing order by blockNumber,index
func (ts SortedEventbalances) Less(i, j int) bool {
	return ts[i].BlockNumber < ts[i].BlockNumber || ts[i].Index < ts[i].Index
}
