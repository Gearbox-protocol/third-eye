package core

import (
	"github.com/Gearbox-protocol/third-eye/log"
	"sort"
	"sync"
)

type SessionData struct {
	Account       string `gorm:"column:account"`
	Since         int64  `gorm:"column:since"`
	CreditManager string `gorm:"column:credit_manager"`
	Status        int    `gorm:"column:status"`
	SessionID     string `gorm:"column:id"`
	ClosedAt      int64  `json:"column:closed_at"`
}

type AccountData struct {
	blockNums    []int64
	blockToIndex map[int64]int
	transfers    map[int64]map[string][]*TokenTransfer
	Details      []*SessionData
}

func newAccountData() *AccountData {
	return &AccountData{
		blockToIndex: map[int64]int{},
		transfers:    map[int64]map[string][]*TokenTransfer{},
	}
}

func (ad *AccountData) process(tt *TokenTransfer) {
	txHash := tt.TxHash
	if ad.transfers[tt.BlockNum] == nil {
		ad.transfers[tt.BlockNum] = make(map[string][]*TokenTransfer)
	}
	ad.transfers[tt.BlockNum][txHash] = append(ad.transfers[tt.BlockNum][txHash], tt)
	if ad.blockToIndex[tt.BlockNum] == 0 {
		ad.blockNums = append(ad.blockNums, tt.BlockNum)
	}
	ad.blockToIndex[tt.BlockNum] = 1
}

func (ad *AccountData) deleteTxHash(blockNum int64, txHash string) {
	if ad.transfers[blockNum] != nil {
		delete(ad.transfers[blockNum], txHash)
	}
}

func (ad *AccountData) init() {
	sort.Slice(ad.blockNums, func(i, j int) bool { return ad.blockNums[i] < ad.blockNums[j] })
	for i, blockNum := range ad.blockNums {
		ad.blockToIndex[blockNum] = i
	}
}

func (ad *AccountData) AddDetails(sd *SessionData) {
	ad.Details = append(ad.Details, sd)
}

func (ad *AccountData) SetStatus(since int64, status int, closedAt int64) {
	for _, details := range ad.Details {
		if since == details.Since {
			details.Status = status
			details.ClosedAt = closedAt
		}
	}
}

// process the transfer events for from <= block < to
func (ad *AccountData) GetRemainingTransfer(cm string, from, to int64) (map[int64]map[string][]*TokenTransfer, []string) {
	// blockNum => sessionID => tokentranfers
	extraTokenTransfers := map[int64]map[string][]*TokenTransfer{}
	noSessionTxs := []string{}
	detailsInd := len(ad.Details) - 1
	blockInd := len(ad.blockNums) - 1
	// find first blockNum less than to
	for ; blockInd >= 0 && ad.blockNums[blockInd] >= to; blockInd-- {
	}
	for ; blockInd >= 0 && ad.blockNums[blockInd] >= from; blockInd-- {
		blockNum := ad.blockNums[blockInd]
		// find the account valid at
		for ; detailsInd >= 0 && ad.Details[detailsInd].Since > blockNum; detailsInd-- {

		}
		if detailsInd < 0 {
			log.Fatal("Token transferred to account before anyone is assigned")
		}
		details := ad.Details[detailsInd]
		if details.CreditManager != cm {
			continue
		}
		if blockNum >= details.Since && (details.Status == 0 || details.ClosedAt > blockNum) {
			if extraTokenTransfers[blockNum] == nil {
				extraTokenTransfers[blockNum] = make(map[string][]*TokenTransfer)
			}
			for _, txs := range ad.transfers[blockNum] {
				extraTokenTransfers[blockNum][details.SessionID] = append(
					extraTokenTransfers[blockNum][details.SessionID], txs...)
			}
		} else {
			for _, txs := range ad.transfers[blockNum] {
				for _, tx := range txs {
					noSessionTxs = append(noSessionTxs, tx.TxHash)
				}
			}
		}
	}
	return extraTokenTransfers, noSessionTxs
}

//
type AccountTokenManager struct {
	// account => blockNum => txhash => transfers
	accountToData map[string]*AccountData
	// txhash to account which transferred asset
	txHashToAccounts map[string][]string
	mu               *sync.Mutex
	NoSessionTxs     []string
	txToTransfers    map[string][]*TokenTransfer
}

func NewAccountTokenManager() *AccountTokenManager {
	return &AccountTokenManager{
		accountToData: make(map[string]*AccountData),
		mu:            &sync.Mutex{},
	}
}

func (mgr *AccountTokenManager) AddTokenTransfer(tt *TokenTransfer, isFromAccount, isToAccount bool) {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()
	if isFromAccount {
		if mgr.accountToData[tt.From] == nil {
			mgr.accountToData[tt.From] = newAccountData()
		}
		mgr.accountToData[tt.From].process(tt)
	}
	if isToAccount {
		if mgr.accountToData[tt.To] == nil {
			mgr.accountToData[tt.To] = newAccountData()
		}
		mgr.accountToData[tt.To].process(tt)
	}
	mgr.txToTransfers[tt.TxHash] = append(mgr.txToTransfers[tt.TxHash], tt)
}

func (mgr *AccountTokenManager) AddAccountDetails(sessionData *SessionData) {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()
	account := sessionData.Account
	if mgr.accountToData[account] == nil {
		mgr.accountToData[account] = newAccountData()
	}
	mgr.accountToData[account].AddDetails(sessionData)
}

func (mgr *AccountTokenManager) CloseAccountDetails(account string, status int, since, closedAt int64) {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()
	mgr.accountToData[account].SetStatus(since, status, closedAt)
}

func (mgr *AccountTokenManager) CheckTokenTransfer(cm string, from, to int64) map[int64]map[string][]*TokenTransfer {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()
	result := map[int64]map[string][]*TokenTransfer{}
	for _, dataMdl := range mgr.accountToData {
		answer, noSessionTxs := dataMdl.GetRemainingTransfer(cm, from, to)
		mgr.NoSessionTxs = append(mgr.NoSessionTxs, noSessionTxs...)
		for blockNum, data := range answer {
			if result[blockNum] == nil {
				result[blockNum] = make(map[string][]*TokenTransfer)
			}
			for sessionID, transfers := range data {
				result[blockNum][sessionID] = transfers
			}
		}
	}
	return result
}

func (mgr *AccountTokenManager) DeleteTxHash(blockNum int64, txHash string) {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()
	for _, account := range mgr.txHashToAccounts[txHash] {
		mgr.accountToData[account].deleteTxHash(blockNum, txHash)
	}
}

func (mgr *AccountTokenManager) Clear() {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()
	mgr.accountToData = make(map[string]*AccountData)
	mgr.txHashToAccounts = make(map[string][]string)
}

func (mgr *AccountTokenManager) Init() {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()
	for _, accountData := range mgr.accountToData {
		accountData.init()
	}
}

func (mgr *AccountTokenManager) GetNoSessionTxs() (tts map[string][]*TokenTransfer) {
	for _, txHash := range mgr.NoSessionTxs {
		tts[txHash] = mgr.txToTransfers[txHash]
	}
	return
}
