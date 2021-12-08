package core

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type ExecuteParams struct {
	SessionId             string
	Protocol              common.Address
	CreditAccount         common.Address
	Borrower              common.Address
	Index                 uint
	BlockNumber           int64
	CumulativeIndexAtOpen *big.Int
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
	Args      string
	Transfers Transfers
}
