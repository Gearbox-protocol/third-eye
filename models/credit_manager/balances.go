package credit_manager

import (
	"math/big"
)

type EventBalance struct {
	BorrowedAmount *big.Int
	Transfers      map[string]*big.Int
	SessionId      string
	BlockNumber    int64
	Index          int64
	Clear          bool
}

func newEventBalance(blockNum uint64, index uint, sessionId string, borrowedAmount *big.Int, transfers map[string]*big.Int, clear bool) EventBalance {
	return EventBalance{
		BorrowedAmount: borrowedAmount,
		Transfers:      transfers,
		SessionId:      sessionId,
		BlockNumber:    int64(blockNum),
		Index:          int64(index),
		Clear:          clear,
	}
}

func (mdl *CreditManager) addEventBalance(eb EventBalance) {
	mdl.eventBalances = append(mdl.eventBalances, &eb)
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
