package ds

import (
	"math/big"
	"testing"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

func getAccountWithAddr(addr string, openBlock, closeBlock int64) *SessionData {
	return getAccountWithAddrAndCM(addr, utils.RandomAddr(), openBlock, closeBlock)
}
func getAccountWithAddrAndCM(addr, cm string, openBlock, closeBlock int64) *SessionData {
	return &SessionData{
		CreditManager: cm,
		Account:       addr,
		Since:         openBlock,
		ClosedAt:      closeBlock,
		OpenTxHash:    utils.RandomHash(),
		ClosedTxHash:  utils.RandomHash(),
		SessionID:     utils.RandomHash(),
	}
}
func getAccount(openBlock, closeBlock int64) *SessionData {
	return getAccountWithAddr(utils.RandomAddr(), openBlock, closeBlock)
}
func getAccountWithLogIds(addr string, openBlock int64, openLogID uint, closeBlock int64, closeLogID uint) *SessionData {
	return &SessionData{
		CreditManager: utils.RandomAddr(),
		Account:       addr,
		Since:         openBlock,
		ClosedAt:      closeBlock,
		OpenLogId:     openLogID,
		ClosedLogId:   closeLogID,
		OpenTxHash:    utils.RandomHash(),
		ClosedTxHash:  utils.RandomHash(),
		SessionID:     utils.RandomHash(),
	}
}

func getTokenTransferWithLogId(account string, blockNum int64, logID uint) *schemas.TokenTransfer {
	txHash := utils.RandomHash()
	token := utils.RandomAddr()
	tmpAddr := utils.RandomAddr()
	return &schemas.TokenTransfer{
		BlockNum:      blockNum,
		LogID:         logID,
		TxHash:        txHash,
		Token:         token,
		From:          tmpAddr,
		To:            account,
		Amount:        (*core.BigInt)(big.NewInt(1)),
		IsFromAccount: false,
		IsToAccount:   true,
	}
}
func getTokenTransfer(account string, blockNum int64) *schemas.TokenTransfer {
	return getTokenTransferWithLogId(account, blockNum, 0)
}

// TEST
func TestAccountManagerTxHashDelete(t *testing.T) {
	mgr := NewDirectTransferManager()
	account1 := getAccount(1, 3)
	transfer1 := getTokenTransfer(account1.Account, 2)
	mgr.AddAccountDetails(account1)
	mgr.AddTokenTransfer(transfer1)
	mgr.DeleteTxHash(transfer1.BlockNum, transfer1.TxHash)
	directTransfers := mgr.CheckTokenTransfer(account1.CreditManager, 2, 3)
	if len(directTransfers[transfer1.BlockNum][account1.SessionID]) != 0 {
		t.Error("transfer not deleted")
	}
}

func TestAccountManagerDetectTransferAtClose(t *testing.T) {
	mgr := NewDirectTransferManager()
	account1 := getAccount(1, 3)
	transfer1 := getTokenTransfer(account1.Account, 3)
	mgr.AddAccountDetails(account1)
	mgr.AddTokenTransfer(transfer1)
	directTransfers := mgr.CheckTokenTransfer(account1.CreditManager, 1, 4)
	txs := directTransfers[transfer1.BlockNum][account1.SessionID]
	if !txs[0].Equal(transfer1) {
		t.Error("transfer1 not detected for account1")
	}
}
func TestAccountManagerDetectTransferAtEndOfRange(t *testing.T) {
	mgr := NewDirectTransferManager()
	account1 := getAccount(1, 3)
	transfer1 := getTokenTransfer(account1.Account, 3)
	mgr.AddAccountDetails(account1)
	mgr.AddTokenTransfer(transfer1)
	directTransfers := mgr.CheckTokenTransfer(account1.CreditManager, 1, 3)
	if len(directTransfers) != 0 {
		t.Error("transfer detected at end of search range")
	}
}

func TestAccountManagerOpenCloseAccountWithDiffSession(t *testing.T) {
	mgr := NewDirectTransferManager()

	account1 := getAccountWithLogIds(utils.RandomAddr(), 1, 1000, 2, 1000)
	account2 := getAccountWithLogIds(account1.Account, 2, 1001, 3, 3000)
	mgr.AddAccountDetails(account1)
	mgr.AddAccountDetails(account2)
	transfer0 := getTokenTransferWithLogId(account1.Account, 1, 999)
	transfer1 := getTokenTransferWithLogId(account1.Account, 1, 1001)
	transfer2 := getTokenTransferWithLogId(account1.Account, 2, 1001)
	mgr.AddTokenTransfer(transfer0)
	mgr.AddTokenTransfer(transfer1)
	mgr.AddTokenTransfer(transfer2)
	directTransfers := mgr.CheckTokenTransfer(account1.CreditManager, 1, 10)
	txs := directTransfers[transfer1.BlockNum][account1.SessionID]
	if !txs[0].Equal(transfer1) || len(directTransfers[transfer1.BlockNum]) != 1 {
		t.Error("transfer1 not detected for account1")
	}

	directTransfers = mgr.CheckTokenTransfer(account2.CreditManager, 2, 10)
	txs = directTransfers[transfer2.BlockNum][account2.SessionID]
	if !txs[0].Equal(transfer2) || len(directTransfers[transfer2.BlockNum]) != 1 {
		t.Error("transfer2 not detected for account2")
	}
	if !mgr.GetNoSessionTxs()[transfer0.TxHash][0].Equal(transfer0) {
		t.Errorf("transfer0 without session of account1 is not detected.")
	}
}

func TestDirectTransferManager(t *testing.T) {
	mgr := NewDirectTransferManager()

	account1 := getAccount(1, 3)
	account2 := getAccountWithAddr(account1.Account, 4, 5)
	account3 := getAccountWithAddrAndCM(utils.RandomAddr(), account1.CreditManager, 1, 10)
	transfer1 := getTokenTransfer(account1.Account, 2)
	transfer2 := getTokenTransfer(account1.Account, 4)
	transfer3 := getTokenTransfer(account3.Account, 5)
	mgr.AddAccountDetails(account1)
	mgr.AddAccountDetails(account2)
	mgr.AddAccountDetails(account3)
	mgr.AddTokenTransfer(transfer1)
	mgr.AddTokenTransfer(transfer2)
	mgr.AddTokenTransfer(transfer3)

	directTransfers := mgr.CheckTokenTransfer(account1.CreditManager, 2, 10)
	txs := directTransfers[transfer1.BlockNum][account1.SessionID]
	if !txs[0].Equal(transfer1) {
		t.Error("transfer1 not detected for account1")
	}
	t.Log(utils.ToJson(directTransfers))
	txs = directTransfers[transfer3.BlockNum][account3.SessionID]
	if !txs[0].Equal(transfer3) {
		t.Error("transfer3 not detected for account3 and same credit manager as account 1")
	}

	directTransfers = mgr.CheckTokenTransfer(account2.CreditManager, 2, 3)
	if len(directTransfers) != 0 {
		t.Error("account2 creditmanager detected transfer")
	}
	directTransfers = mgr.CheckTokenTransfer(account2.CreditManager, 4, 5)
	txs = directTransfers[transfer2.BlockNum][account2.SessionID]
	if !txs[0].Equal(transfer2) {
		t.Error("transfer2 not detected for account2 with same account as account1")
	}
	if len(mgr.GetNoSessionTxs()) != 0 {
		t.Error("no sessionTx detected.")
	}

	transfer4 := getTokenTransfer(account3.Account, 11)
	mgr.AddTokenTransfer(transfer4)
	if !mgr.GetNoSessionTxs()[transfer4.TxHash][0].Equal(transfer4) {
		t.Error("transfer 4 without any session not detected.")
	}
}
