package services

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Gearbox-protocol/sdk-go/artifacts/creditFacade"
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditFacadev3"
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditManager"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/services/trace_service"

	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
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

type ExecuteParser struct {
	IgnoreCMEventIds map[common.Hash]bool
	trace_service.InternalFetcher
}

func getCMEventIds() map[common.Hash]bool {
	ids := map[common.Hash]bool{}
	for _, contractABI := range []string{creditManager.CreditManagerABI, // v1 has ExecuteOrder event which has same sig as  v2'ExecuteOrder and so we are able to separate transfer in  batches for ExecuteOrder
		creditFacade.CreditFacadeABI, creditFacadev3.CreditFacadev3ABI} {
		if abiObj, err := abi.JSON(strings.NewReader(contractABI)); err == nil {
			for _, event := range abiObj.Events {
				ids[event.ID] = true
			}
		}
	}
	return ids
}

func NewExecuteParser(cfg *config.Config, client core.ClientI) ds.ExecuteParserI {
	return &ExecuteParser{
		IgnoreCMEventIds: getCMEventIds(),
		InternalFetcher:  trace_service.NewInternalFetcher(cfg, client),
	}
}

// used for adding to multicalls of v2/v3 or directly add AccountOperation to db for v1/v3
// used at 2 places, both are in this function, ProcessRemainingMultiCalls
func (ep *ExecuteParser) GetExecuteCalls(version core.VersionType, txHash, creditManagerAddr string, paramsList []ds.ExecuteParams) []*ds.KnownCall {
	if len(paramsList) == 0 {
		return nil
	}
	trace := ep.GetTxTrace(txHash, true)
	filter := ExecuteFilter{paramsList: paramsList,
		creditManager: common.HexToAddress(creditManagerAddr),
	}
	calls := filter.getExecuteCalls(trace.CallTrace)

	var executeTransfers []core.Transfers
	if version.MoreThanEq(core.NewVersion(300)) {
		executeTransfers = filter.getExecuteTransfersv3(trace.Logs, ep.IgnoreCMEventIds)
	} else {
		executeTransfers = filter.getExecuteTransfersv2(trace.Logs, ep.IgnoreCMEventIds)
	}
	// log.Info(utils.ToJson(trace.Logs))
	// log.Info(utils.ToJson(executeTransfers))

	// check if parsed execute Order currently
	if len(calls) == len(executeTransfers) {
		for i, call := range calls {
			call.Transfers = executeTransfers[i]
		}
	} else {
		fileName := fmt.Sprintf("trace-%s-%s.json", txHash, time.Now())
		os.WriteFile(fileName, []byte(utils.ToJson(trace)), os.ModePerm)
		log.Fatalf("Calls %d execute details %d tx:%s creditManager:%s",
			len(calls), len(executeTransfers), txHash, creditManagerAddr)
	}
	return calls
}

// ////////////////////////
// parser functions for v2
// ////////////////////////

// GetTransfersAtClosev2
// currently only valid for closeCreditAccount v2
func (ep *ExecuteParser) GetTransfersAtClosev2(txHash, account, underlyingToken string, users ds.BorrowerAndTo) core.Transfers {
	trace := ep.GetTxTrace(txHash, true)
	return getCloseAccountv2Transfers(trace, account, underlyingToken, users)
}

// currently only valid for closeCreditAccount v2
func getCloseAccountv2Transfers(trace *trace_service.TenderlyTrace, account, underlyingToken string, users ds.BorrowerAndTo) core.Transfers {
	transfers := getTransfersToUser(trace.Logs, account, underlyingToken, users)
	// convertWETH is set, only valid for closecreditaccountv2
	convertWETHInd := 2 + 8 + 64 + 64 + 64
	// for close call if convertEThInd is true
	if trace.CallTrace.Input[:10] == "0x5f73fbec" && trace.CallTrace.Input[convertWETHInd-1] == '1' {
		ethAmount := ethTransferDueToConvertWETH(trace.CallTrace, users)
		if ethAmount == nil {
			// log.Msgf("Can't get unwrapped WETH amount at closeCreditAccount(%s) sent to user. Tx: %s.", account, users.Borrower, trace.TxHash)
			ethAmount = new(big.Int)
		}
		if transfers[underlyingToken] == nil {
			transfers[underlyingToken] = new(big.Int)
		}
		transfers[underlyingToken] = new(big.Int).Add(transfers[underlyingToken], ethAmount)
	}
	return transfers
}

// eth transfer due to convertWETH
func ethTransferDueToConvertWETH(call *trace_service.Call, users ds.BorrowerAndTo) (ethAmount *big.Int) {
	if len(call.Input) == 10+64*2 && call.Input[:10] == "0x5869dba8" && common.HexToAddress(call.Input[10:74]) == users.To {
		ethAmount, _ := new(big.Int).SetString(call.Input[74:], 16)
		return ethAmount
	}
	for _, innerCall := range call.Calls {
		if ethAmount := ethTransferDueToConvertWETH(innerCall, users); ethAmount != nil {
			return ethAmount
		}
	}
	return nil
}

// is valid for closeCreditAccount v2
// tenderly has logs for events(we mainly use for Transfer on token) and calls( for unwrapETH on wethgateway)
// wrapWETH is also present in closecreditaccount, but it sends the wrapped eth back to user and then the user has approval on weth for creditmanager so in second step the weth is transferred
// handling native eth refund is only needed when convertETH is true
// native eth transfer from account is handled in parent function, not in this function
func getTransfersToUser(txLogs []trace_service.Log, account, underlyingToken string, users ds.BorrowerAndTo) core.Transfers {
	transfers := core.Transfers{}
	for _, raw := range txLogs {
		eventLog := raw.Raw
		if eventLog.Topics[0] == core.Topic("Transfer(address,address,uint256)") { // transfer event
			from := common.BytesToAddress(eventLog.Topics[1][:])
			to := common.BytesToAddress(eventLog.Topics[2][:])
			token := eventLog.Address.Hex()
			var sign *big.Int
			if from == users.Borrower && to.Hex() == account && token == underlyingToken {
				sign = big.NewInt(-1)
			} else {
				if !(to == users.Borrower || to == users.To) {
					continue
				}
				sign = big.NewInt(1)
			}
			amt, b := new(big.Int).SetString(eventLog.Data[2:], 16)
			if !b {
				log.Fatal("failed at serializing transfer data in int")
			}
			amt = new(big.Int).Mul(sign, amt)
			oldBalance := new(big.Int)
			if transfers[token] != nil {
				oldBalance = transfers[token]
			}
			transfers[token] = new(big.Int).Add(amt, oldBalance)
		}
	}
	return transfers
}
