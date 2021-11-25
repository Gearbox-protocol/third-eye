package credit_manager

import (
	"fmt"
	"github.com/Gearbox-protocol/gearscan/core"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

func (mdl *CreditManager) onOpenCreditAccount(txLog *types.Log, sender, onBehalfOf, account string,
	amount,
	borrowAmount,
	referralCode *big.Int) error {
	sessionId := fmt.Sprintf("%s_%d_%d", account, txLog.BlockNumber, txLog.Index)
	// action, args := mdl.ParseEvent("OpenCreditAccount", txLog)
	action, args := "", ""
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
		Dapp:        txLog.Address.Hex(),
	}
	mdl.Repo.AddAccountOperation(accountOperation)
	return nil
}
