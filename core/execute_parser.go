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
	Depth     uint8     `json:"depth"`
	Name      string    `json:"name"`
	Args      *Json     `json:"args"`
	Transfers Transfers `json:"transfers"`
}
