package cm_v2

import (
	"math/big"
	"testing"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/credit_manager/cm_mvp"
)

type repo struct {
	ds.DummyRepo
	mgr               *ds.DirectTransferManager
	accountOperations []*schemas.AccountOperation
	sessions          map[string]*schemas.CreditSession
	dieselTokens      map[string]bool
}

func (r repo) GetAccountManager() *ds.DirectTransferManager {
	return r.mgr
}
func (r *repo) AddAccountOperation(op *schemas.AccountOperation) {
	r.accountOperations = append(r.accountOperations, op)
}
func (r repo) GetCreditSession(sessionId string) *schemas.CreditSession {
	return r.sessions[sessionId]
}

func (r repo) IsDieselToken(token string) bool {
	return r.dieselTokens[token]
}

func getAccountWithAddrAndCM(addr, cm string, openBlock, closeBlock int64, sessionId string) *ds.SessionData {
	return &ds.SessionData{
		CreditManager: cm,
		Account:       addr,
		Since:         openBlock,
		ClosedAt:      closeBlock,
		OpenTxHash:    utils.RandomHash(),
		ClosedTxHash:  utils.RandomHash(),
		SessionID:     sessionId,
	}
}

func (r *repo) getTokenTransferWithLogId(account string, blockNum int64, logID uint, fromAddr string, txHash string) *schemas.TokenTransfer {
	token := utils.RandomAddr()
	r.dieselTokens[token] = true
	return &schemas.TokenTransfer{
		BlockNum:      blockNum,
		LogID:         logID,
		TxHash:        txHash,
		Token:         token,
		From:          fromAddr,
		To:            account,
		Amount:        (*core.BigInt)(big.NewInt(1)),
		IsFromAccount: false,
		IsToAccount:   true,
	}
}

func newRepo(sessionAndAccount map[string]string) *repo {
	r := &repo{
		mgr:          ds.NewDirectTransferManager(),
		sessions:     map[string]*schemas.CreditSession{},
		dieselTokens: map[string]bool{},
	}

	for session, account := range sessionAndAccount {
		r.sessions[session] = &schemas.CreditSession{Account: account}
	}
	return r
}

func TestRewardClaimed(t *testing.T) {
	cm := utils.RandomAddr()
	account := utils.RandomAddr()
	allowedProtocol, allowedProtocol2 := utils.RandomAddr(), utils.RandomAddr()
	sessionId := utils.RandomHash()
	//
	r := newRepo(map[string]string{sessionId: account})
	session := getAccountWithAddrAndCM(account, cm, 10, 20, sessionId)
	txHash := utils.RandomHash()
	transfer1 := r.getTokenTransferWithLogId(account, 11, 0, core.NULL_ADDR.Hex(), txHash) // from 0x0
	transfer2 := r.getTokenTransferWithLogId(account, 11, 1, allowedProtocol, txHash)
	transfer3 := r.getTokenTransferWithLogId(account, 12, 2, allowedProtocol2, utils.RandomHash())
	//
	r.mgr.AddAccountDetails(session)
	r.mgr.AddTokenTransfer(transfer1)
	r.mgr.AddTokenTransfer(transfer2)
	r.mgr.AddTokenTransfer(transfer3)

	adapter := &ds.SyncAdapter{
		Repo: r,
		SyncAdapterSchema: &schemas.SyncAdapterSchema{
			Contract: &schemas.Contract{
				Address: cm,
			},
		},
	}
	cmModel := CMv2{
		Cmv1v2: cm_mvp.NewCMv1v2(adapter),
		allowedProtocols: map[string]bool{
			allowedProtocol:  true,
			allowedProtocol2: true,
		},
	}
	cmModel.DontGetSessionFromDCForTest = true
	cmModel.SetOnDirectTokenTransferFn(cmModel.getDirectTokenTransferFn())
	//
	cmModel.UpdateSessionWithDirectTokenTransferBefore(20)
	if len(r.accountOperations) != 2 {
		t.Fatal("Improper account operations", utils.ToJson(r.accountOperations))
	}
	if len(*r.accountOperations[0].Args) != 2 || len(*r.accountOperations[0].Transfers) != 2 {
		t.Fatal("First accoutOperation is wrong", utils.ToJson(r.accountOperations[0]))
	}
	if len(*r.accountOperations[1].Args) != 1 || len(*r.accountOperations[1].Transfers) != 1 {
		t.Fatal("Second accoutOperation is wrong", utils.ToJson(r.accountOperations[0]))
	}
}
