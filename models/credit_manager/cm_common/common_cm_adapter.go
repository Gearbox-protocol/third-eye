package cm_common

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/core/types"
)

type CommonCMAdapter struct {
	*ds.SyncAdapter
	State                  *schemas.CreditManagerState
	borrowedAmountForBlock *big.Int
	MulticallMgr           ds.MultiCallProcessor
	//
	onChangeDetails
	//
	// tmp storage
	PnlOnCM                     *PnlCM
	params                      *schemas.Parameters
	DontGetSessionFromDCForTest bool
	//
	// calculating credit session stats
	UpdatedSessions map[string]int
	ClosedSessions  map[string]*SessionCloseDetails
}

func NewCommonCMAdapter(adapter *ds.SyncAdapter) CommonCMAdapter {
	return CommonCMAdapter{
		SyncAdapter: adapter,
		//
		PnlOnCM: NewPnlCM(),
		//
		UpdatedSessions: make(map[string]int),
		ClosedSessions:  make(map[string]*SessionCloseDetails),
	}
}

func (mdl *CommonCMAdapter) GetUnderlyingToken() string {
	return mdl.State.UnderlyingToken
}
func (mdl *CommonCMAdapter) GetUnderlyingDecimal() int8 {
	decimals := mdl.Repo.GetToken(mdl.GetUnderlyingToken()).Decimals
	return decimals
}

func (mdl *CommonCMAdapter) PoolBorrow(txLog *types.Log, sessionId, borrower string, amount *big.Int) {
	mdl.Repo.AddPoolLedger(&schemas.PoolLedger{
		LogId:       txLog.Index,
		BlockNumber: int64(txLog.BlockNumber),
		TxHash:      txLog.TxHash.Hex(),
		Pool:        mdl.State.PoolAddress,
		Event:       "Borrow",
		User:        borrower,
		SessionId:   sessionId,
		AmountBI:    (*core.BigInt)(amount),
		Amount:      utils.GetFloat64Decimal(amount, mdl.GetUnderlyingDecimal()),
	})
}

func (mdl *CommonCMAdapter) PoolRepay(blockNum int64, logId uint, txHash, sessionId, borrower string, amount *big.Int) {
	mdl.Repo.AddPoolLedger(&schemas.PoolLedger{
		LogId:       logId,
		BlockNumber: blockNum,
		TxHash:      txHash,
		Pool:        mdl.State.PoolAddress,
		Event:       "Repay",
		User:        borrower,
		SessionId:   sessionId,
		AmountBI:    (*core.BigInt)(amount),
		Amount:      utils.GetFloat64Decimal(amount, mdl.GetUnderlyingDecimal()),
	})
}

func (mdl *CommonCMAdapter) GetUnderlyingState() interface{} {
	return mdl.State
}

// for params
func (mdl *CommonCMAdapter) SetParams(params *schemas.Parameters) {
	mdl.params = params
}

func (mdl *CommonCMAdapter) CMStatsOnOpenAccount(borrowAmount *big.Int) {
	// manager state
	mdl.State.TotalOpenedAccounts++
	mdl.State.OpenedAccountsCount++
	mdl.AddBorrowAmountForBlock(borrowAmount)
}

// borroweAmount can't be negative
func (mdl *CommonCMAdapter) AddBorrowAmountForBlock(borrowAmount *big.Int) {
	if borrowAmount.Sign() < 0 {
		log.Fatal("Borrowed Amount can't be negative. As repaid amount is tracked on pool")
	}
	if mdl.borrowedAmountForBlock == nil {
		mdl.borrowedAmountForBlock = new(big.Int)
	}
	mdl.borrowedAmountForBlock = new(big.Int).Add(mdl.borrowedAmountForBlock, borrowAmount)
}

func (mdl *CommonCMAdapter) GetBorrowAmountForBlockAndClear() *big.Int {
	if mdl.borrowedAmountForBlock == nil {
		return new(big.Int)
	}
	lastValue := mdl.borrowedAmountForBlock
	mdl.borrowedAmountForBlock = new(big.Int)
	return lastValue
}

func (mdl CommonCMAdapter) AddAccountOperation(accountOperation *schemas.AccountOperation) {
	mdl.Repo.AddAccountOperation(accountOperation)
}

func (mdl CommonCMAdapter) CloseAccount(sessionID string, blockNum int64, txHash string, logID uint) {
	session := mdl.Repo.GetCreditSession(sessionID)
	mdl.Repo.GetAccountManager().CloseAccountDetails(session.Account, session.Since, blockNum, txHash, logID)
}

func (mdl CommonCMAdapter) SaveExecuteEvents(lastTxHash string, executeParams []ds.ExecuteParams) {
	// credit manager has the execute event
	calls := mdl.Repo.GetExecuteParser().GetExecuteCalls(lastTxHash, mdl.Address, executeParams)

	for i, call := range calls {
		params := executeParams[i]

		accountOperation := &schemas.AccountOperation{
			BlockNumber: params.BlockNumber,
			TxHash:      lastTxHash,
			LogId:       params.Index,
			// owner/account data
			Borrower:  params.Borrower.Hex(),
			SessionId: params.SessionId,
			// dapp
			Dapp: params.Protocol.Hex(),
			// call/events data
			Action:      call.Name,
			Args:        call.Args,
			AdapterCall: true,
			Transfers:   &call.Transfers,
		}
		mdl.AddAccountOperation(accountOperation)
		mdl.SetSessionIsUpdated(params.SessionId)
	}
}
