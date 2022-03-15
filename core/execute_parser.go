package core

import (
	"github.com/Gearbox-protocol/third-eye/artifacts/creditFacade"
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

type FuncWithMultiCall struct {
	Name       string
	MultiCalls []*creditFacade.MultiCall
}

func (f *FuncWithMultiCall) LenOfMultiCalls() int {
	return len(f.MultiCalls)
}

type ExecuteParserI interface {
	GetExecuteCalls(txHash, creditManagerAddr string, paramsList []ExecuteParams) []*KnownCall
	GetMainEventLogs(txHash, creditFacade string) []*FuncWithMultiCall
	GetTransfers(txHash string, owner []string) Transfers
}

type KnownCall struct {
	// Input string
	Depth     uint8     `json:"depth"`
	Name      string    `json:"name"`
	Args      *Json     `json:"args"`
	Transfers Transfers `json:"transfers"`
}
