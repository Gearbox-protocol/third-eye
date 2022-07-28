package ds

import (
	"encoding/hex"

	"github.com/Gearbox-protocol/sdk-go/core"
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
type GBv2Multicall []struct {
	Target   common.Address `json:"target"`
	CallData []uint8        `json:"callData"`
}
type FuncWithMultiCall struct {
	Name       string        `json:"name"`
	MultiCalls GBv2Multicall `json:"-"`
	TestLen    int           `json:"len"`
}

func (f *FuncWithMultiCall) LenForBorrower(bwr string) (len int) {
	// for testing
	if f.TestLen != 0 {
		return f.TestLen
	}
	borrower := common.HexToAddress(bwr)
	for _, call := range f.MultiCalls {
		if hex.EncodeToString(call.CallData[:4]) == "59781034" { // addcollateral(address, address, uint256) call
			// if the onbehalf is different then borrower ignore
			if common.BytesToAddress(call.CallData[4:36]) != borrower {
				continue
			}
		}
		len++
	}
	return
}

type ExecuteParserI interface {
	GetExecuteCalls(txHash, creditManagerAddr string, paramsList []ExecuteParams) []*KnownCall
	GetMainEventLogs(txHash, creditFacade string) []*FuncWithMultiCall
	GetTransfers(txHash string, borrower, account, underlyingToken string, owner []string) core.Transfers
}

type KnownCall struct {
	// Input string
	Depth     uint8          `json:"depth"`
	Name      string         `json:"name"`
	Args      *core.Json     `json:"args"`
	Transfers core.Transfers `json:"transfers"`
}
