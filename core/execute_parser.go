package core

import (
	"github.com/ethereum/go-ethereum/common"
)

type ExecuteParams struct {
	SessionId     string
	Protocol      common.Address
	CreditAccount common.Address
	Borrower      common.Address
	Index         uint
	BlockNumber   int64
}

type ExecuteParserI interface {
	GetExecuteCalls(txHash, creditManagerAddr string, paramsList []ExecuteParams) []*KnownCall
}

type KnownCall struct {
	// Input string
	From      common.Address
	To        common.Address
	Depth     uint8
	Name      string
	Args      *Json
	Transfers Transfers
}
