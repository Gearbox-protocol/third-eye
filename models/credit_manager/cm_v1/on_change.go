package cm_v1

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/credit_manager/cm_common"
)

func OnDirectTokenTransfer(repo ds.RepositoryI, tx *schemas.TokenTransfer, session *schemas.CreditSession) {
	if tx == nil || session == nil {
		return
	}
	repo.RecentMsgf(log.RiskHeader{
		BlockNumber: tx.BlockNum,
		EventCode:   "AMQP",
	}, "Deposit: %s", cm_common.DirecTokenTransferString(repo, tx))
	amount := tx.Amount.Convert()
	repo.AddAccountOperation(&schemas.AccountOperation{
		TxHash:      tx.TxHash,
		BlockNumber: tx.BlockNum,
		LogId:       tx.LogID,
		Borrower:    session.Borrower,
		SessionId:   session.ID,
		Dapp:        tx.Token,
		Action:      "DirectTokenTransfer",
		Args:        &core.Json{"amount": amount, "to": tx.To, "from": tx.From},
		AdapterCall: false,
		Transfers:   &core.Transfers{tx.Token: amount},
	})
}

func (mdl *CMv1) lastTxHashCompleted(lastTxHash string) {
	if len(mdl.executeParams) > 0 {
		mdl.SaveExecuteEvents(lastTxHash, mdl.executeParams)
		mdl.executeParams = []ds.ExecuteParams{}
	}
}
func (mdl *CMv1) SetOnChange() {
	mdl.SetLastTxHashCompleted(mdl.lastTxHashCompleted)
	mdl.SetCalculateCMStatFn(mdl.CalculateCMStat)
	mdl.SetOnDirectTokenTransferFn(OnDirectTokenTransfer)
}
