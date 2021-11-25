package credit_manager

import (
	"fmt"
	"github.com/Gearbox-protocol/gearscan/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
)

func (mdl *CreditManager) onOpenCreditAccount(txLog *types.Log, sender, onBehalfOf, account string,
	amount,
	borrowAmount,
	referralCode *big.Int) error {
	cmAddr := txLog.Address.Hex()
	sessionId := fmt.Sprintf("%s_%d_%d", account, txLog.BlockNumber, txLog.Index)
	action, args := mdl.ParseEvent("OpenCreditAccount", txLog)
	accountOperation := &core.AccountOperation{
		TxHash:      txLog.TxHash.Hex(),
		BlockNumber: int64(txLog.BlockNumber),
		LogId:       txLog.Index,
		Borrower:    onBehalfOf,
		SessionId:   sessionId,
		AdapterCall: false,
		Action:      action,
		Args:        args,
		Transfers:   "",
		Dapp:        cmAddr,
	}
	mdl.Repo.AddAccountOperation(accountOperation)
	mdl.Repo.AddCreditOwnerSession(cmAddr, onBehalfOf, sessionId)
	return nil
}

// onCloseCreditAccount handles CloseCreditAccount Event
func (mdl *CreditManager) onCloseCreditAccount(txLog *types.Log, owner, to string, remainingFunds *big.Int) error {
	cmAddr := txLog.Address.Hex()
	sessionId := mdl.Repo.GetCreditOwnerSession(cmAddr, owner)
	action, args := mdl.ParseEvent("CloseCreditAccount", txLog)
	accountOperation := &core.AccountOperation{
		TxHash:          txLog.TxHash.Hex(),
		BlockNumber: int64(txLog.BlockNumber),
		LogId: txLog.Index,
		Borrower: owner,
		SessionId: sessionId,
		AdapterCall: false,
		Action: action,
		Args: args,
		Transfers: "",
		Dapp: cmAddr,
	}
	mdl.Repo.AddAccountOperation(accountOperation)
	mdl.Repo.RemoveCreditOwnerSession(cmAddr, owner)
	return nil
}

func (mdl *CreditManager) onLiquidateCreditAccount(txLog *types.Log, owner, liquidator string, remainingFunds *big.Int) error {
	cmAddr := txLog.Address.Hex()
	sessionId := mdl.Repo.GetCreditOwnerSession(cmAddr, owner)
	action, args := mdl.ParseEvent("LiquidateCreditAccount", txLog)
	accountOperation := &core.AccountOperation{
		TxHash:          txLog.TxHash.Hex(),
		BlockNumber:        int64(txLog.BlockNumber),
		LogId: txLog.Index,
		Borrower: owner,
		SessionId: sessionId,
		AdapterCall: false,
		Action: action,
		Args: args,
		Transfers: "",
		Dapp: txLog.Address.Hex(),
	}
	mdl.Repo.AddAccountOperation(accountOperation)
	mdl.Repo.RemoveCreditOwnerSession(cmAddr, owner)
	return nil
}

func (mdl *CreditManager) onRepayCreditAccount(txLog *types.Log, owner, to string) error {
	cmAddr := txLog.Address.Hex()
	sessionId := mdl.Repo.GetCreditOwnerSession(cmAddr, owner)
	action, args := mdl.ParseEvent("RepayCreditAccount", txLog)
	accountOperation := &core.AccountOperation{
		TxHash:          txLog.TxHash.Hex(),
		BlockNumber:        int64(txLog.BlockNumber),
		LogId: txLog.Index,
		Borrower: owner,
		SessionId: sessionId,
		AdapterCall: false,
		Action: action,
		Args: args,
		Transfers: "",
		Dapp: cmAddr,
	}
	mdl.Repo.AddAccountOperation(accountOperation)
	mdl.Repo.RemoveCreditOwnerSession(cmAddr, owner)
	return nil
}

func (mdl *CreditManager) onAddCollateral(txLog *types.Log, onBehalfOf, token string, value *big.Int) error {
	cmAddr := txLog.Address.Hex()
	sessionId := mdl.Repo.GetCreditOwnerSession(cmAddr, onBehalfOf)
	action, args := mdl.ParseEvent("RepayCreditAccount", txLog)
	accountOperation := &core.AccountOperation{
		TxHash:          txLog.TxHash.Hex(),
		BlockNumber:        int64(txLog.BlockNumber),
		LogId: txLog.Index,
		Borrower: onBehalfOf,
		SessionId: sessionId,
		AdapterCall: false,
		Action: action,
		Args: args,
		Transfers: "",
		Dapp: txLog.Address.Hex(),
	}
	mdl.Repo.AddAccountOperation(accountOperation)
	mdl.Repo.RemoveCreditOwnerSession(cmAddr, onBehalfOf)
	return nil
}

func (mdl *CreditManager) onIncreaseBorrowedAmount(txLog *types.Log, borrower string, amount *big.Int) error {
	cmAddr := txLog.Address.Hex()
	sessionId := mdl.Repo.GetCreditOwnerSession(cmAddr, borrower)
	action, args := mdl.ParseEvent("IncreaseBorrowedAmount", txLog)
	accountOperation := &core.AccountOperation{
		TxHash:          txLog.TxHash.Hex(),
		BlockNumber:        int64(txLog.BlockNumber),
		LogId: txLog.Index,
		Borrower: borrower,
		SessionId: sessionId,
		AdapterCall: false,
		Action: action,
		Args: args,
		Transfers: "",
		Dapp: txLog.Address.Hex(),
	}
	mdl.Repo.AddAccountOperation(accountOperation)
	return nil
}

func (mdl *CreditManager) onTransferAccount(txLog *types.Log, owner, newOwner string) error {
	cmAddr := txLog.Address.Hex()
	sessionId := mdl.Repo.GetCreditOwnerSession(cmAddr, owner)
	action, args := mdl.ParseEvent("TransferAccount", txLog)
	accountOperation := &core.AccountOperation{
		TxHash:          txLog.TxHash.Hex(),
		BlockNumber:        int64(txLog.BlockNumber),
		LogId: txLog.Index,
		Borrower: owner,
		SessionId: sessionId,
		AdapterCall: false,
		Action: action,
		Args: args,
		Transfers: "",
		Dapp: txLog.Address.Hex(),
	}
	mdl.Repo.AddAccountOperation(accountOperation)
	mdl.Repo.RemoveCreditOwnerSession(cmAddr, owner)
	mdl.Repo.AddCreditOwnerSession(cmAddr, newOwner, sessionId)
	return nil
}

func (mdl *CreditManager) onExecuteOrder(txLog *types.Log,
	borrower,
	targetContract common.Address) error {
	cmAddr := txLog.Address.Hex()
	sessionId := mdl.Repo.GetCreditOwnerSession(cmAddr, borrower.Hex())
	creditAccount := strings.Split(sessionId, "_")[0]
	calls := mdl.Repo.GetExecuteParser().GetExecuteCalls(txLog, common.HexToAddress(creditAccount), targetContract)
	
	for _, call := range calls {
		accountOperation := &core.AccountOperation{
			BlockNumber: int64(txLog.BlockNumber),
			TxHash: txLog.TxHash.Hex(),
			LogId: uint(call.LogId),
			// owner/account data
			Borrower: borrower.Hex(),
			SessionId: sessionId,
			// dapp
			Dapp: targetContract.Hex(),
			// call/events data
			Action: call.Name,
			Args: call.Args,
			AdapterCall: true,
			Transfers: call.Balances.String(),
			// extras
			Depth: call.Depth,
		}
		mdl.Repo.AddAccountOperation(accountOperation)
	}
	return nil
}