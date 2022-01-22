package core

import (
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/utils"
	"sort"
	"sync"
)

type SessionData struct {
	Account       string `gorm:"column:account"`
	Since         int64  `gorm:"column:since"`
	CreditManager string `gorm:"column:credit_manager"`
	SessionID     string `gorm:"column:id"`
	ClosedAt      int64  `gorm:"column:closed_at"`
}

type AccountData struct {
	Address      string
	blockNums    []int64
	blockPresent map[int64]bool
	transfers    map[int64]map[string][]*TokenTransfer
	Details      []*SessionData
}

func newAccountData(addr string) *AccountData {
	return &AccountData{
		Address:      addr,
		blockPresent: map[int64]bool{},
		transfers:    map[int64]map[string][]*TokenTransfer{},
	}
}

func (ad *AccountData) process(tt *TokenTransfer) {
	txHash := tt.TxHash
	if ad.transfers[tt.BlockNum] == nil {
		ad.transfers[tt.BlockNum] = make(map[string][]*TokenTransfer)
	}
	ad.transfers[tt.BlockNum][txHash] = append(ad.transfers[tt.BlockNum][txHash], tt)
	if !ad.blockPresent[tt.BlockNum] {
		ad.blockNums = append(ad.blockNums, tt.BlockNum)
	}
	ad.blockPresent[tt.BlockNum] = true
}

func (ad *AccountData) deleteTxHash(blockNum int64, txHash string) {
	if ad.transfers[blockNum] != nil {
		delete(ad.transfers[blockNum], txHash)
	}
}

func (ad *AccountData) Clear() {
	ad.blockNums = []int64{}
	ad.blockPresent = make(map[int64]bool)
	ad.transfers = make(map[int64]map[string][]*TokenTransfer)
}

func (ad *AccountData) init() {
	sort.Slice(ad.blockNums, func(i, j int) bool { return ad.blockNums[i] < ad.blockNums[j] })
}

func (ad *AccountData) AddDetails(sd *SessionData) {
	ad.Details = append(ad.Details, sd)
}

func (ad *AccountData) SetStatus(since int64, closedAt int64) {
	// log.Info(utils.ToJson(ad.Details))
	for _, details := range ad.Details {
		if since == details.Since {
			details.ClosedAt = closedAt
		}
	}
}

func (ad *AccountData) detailsAssigned() bool {
	// if the account details is not assigned then
	// for some other credit manager, when getting remaining transfers
	// we should skip for this account as it can be different credit manager
	// that detail is missing currently
	return len(ad.Details) != 0
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
			log.Fatalf(`Token transferred to account(%s) before borrower is assigned.
				CreditManager With Since/closed details: %s. 
				session transfers: %v
				BlockNum with transfers: %v`,
				ad.Address,
				utils.ToJson(ad.Details),
				ad.transfers[blockNum],
				ad.blockNums,
			)
		}
		details := ad.Details[detailsInd]
		if details.CreditManager != cm {
			continue
		}
		if blockNum >= details.Since && (details.ClosedAt == 0 || details.ClosedAt > blockNum) {
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
		accountToData:    make(map[string]*AccountData),
		mu:               &sync.Mutex{},
		txToTransfers:    make(map[string][]*TokenTransfer),
		txHashToAccounts: make(map[string][]string),
	}
}

func (mgr *AccountTokenManager) AddTokenTransfer(tt *TokenTransfer, isFromAccount, isToAccount bool) {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()
	if isFromAccount {
		if mgr.accountToData[tt.From] == nil {
			mgr.accountToData[tt.From] = newAccountData(tt.From)
		}
		mgr.txHashToAccounts[tt.TxHash] = append(mgr.txHashToAccounts[tt.TxHash], tt.From)
		mgr.accountToData[tt.From].process(tt)
	}
	if isToAccount {
		if mgr.accountToData[tt.To] == nil {
			mgr.accountToData[tt.To] = newAccountData(tt.To)
		}
		mgr.txHashToAccounts[tt.TxHash] = append(mgr.txHashToAccounts[tt.TxHash], tt.To)
		mgr.accountToData[tt.To].process(tt)
	}
	mgr.txToTransfers[tt.TxHash] = append(mgr.txToTransfers[tt.TxHash], tt)
}

func (mgr *AccountTokenManager) AddAccountDetails(sessionData *SessionData) {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()
	account := sessionData.Account
	if mgr.accountToData[account] == nil {
		mgr.accountToData[account] = newAccountData(account)
	}
	mgr.accountToData[account].AddDetails(sessionData)
}

func (mgr *AccountTokenManager) CloseAccountDetails(account string, since, closedAt int64) {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()
	mgr.accountToData[account].SetStatus(since, closedAt)
}

func (mgr *AccountTokenManager) CheckTokenTransfer(cm string, from, to int64) map[int64]map[string][]*TokenTransfer {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()
	result := map[int64]map[string][]*TokenTransfer{}
	for _, dataMdl := range mgr.accountToData {
		if !dataMdl.detailsAssigned() {
			continue
		}
		remainingTransfers, noSessionTxs := dataMdl.GetRemainingTransfer(cm, from, to)
		mgr.NoSessionTxs = append(mgr.NoSessionTxs, noSessionTxs...)
		for blockNum, data := range remainingTransfers {
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
	for _, data := range mgr.accountToData {
		data.Clear()
	}
	mgr.txHashToAccounts = make(map[string][]string)
	mgr.txToTransfers = make(map[string][]*TokenTransfer)
	mgr.NoSessionTxs = []string{}
}

func (mgr *AccountTokenManager) Init() {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()
	for _, accountData := range mgr.accountToData {
		accountData.init()
	}
}

func (mgr *AccountTokenManager) GetNoSessionTxs() (tts map[string][]*TokenTransfer) {
	tts = make(map[string][]*TokenTransfer)
	for _, txHash := range mgr.NoSessionTxs {
		tts[txHash] = mgr.txToTransfers[txHash]
	}
	log.Infof("len of nosessionTxs %d", len(tts))
	return
}
