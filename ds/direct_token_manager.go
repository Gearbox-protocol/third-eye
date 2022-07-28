package ds

import (
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"

	// "github.com/Gearbox-protocol/sdk-go/utils"
	"sort"
	"sync"
)

type SessionData struct {
	Account       string
	Since         int64
	CreditManager string
	SessionID     string
	ClosedAt      int64
	OpenTxHash    string
	OpenLogId     uint
	ClosedTxHash  string
	ClosedLogId   uint
}

type AccountData struct {
	Address      string
	blockNums    []int64
	blockPresent map[int64]bool
	transfers    map[int64]map[string][]*schemas.TokenTransfer
	Details      []*SessionData
}

func newAccountData(addr string) *AccountData {
	return &AccountData{
		Address:      addr,
		blockPresent: map[int64]bool{},
		transfers:    map[int64]map[string][]*schemas.TokenTransfer{},
	}
}

func (ad *AccountData) process(tt *schemas.TokenTransfer) {
	txHash := tt.TxHash
	if ad.transfers[tt.BlockNum] == nil {
		ad.transfers[tt.BlockNum] = make(map[string][]*schemas.TokenTransfer)
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
	ad.transfers = make(map[int64]map[string][]*schemas.TokenTransfer)
}

func (ad *AccountData) init() {
	sort.Slice(ad.blockNums, func(i, j int) bool { return ad.blockNums[i] < ad.blockNums[j] })
}

func (ad *AccountData) AddDetails(sd *SessionData) {
	ad.Details = append(ad.Details, sd)
}

func (ad *AccountData) SetClose(since int64, closedAt int64, closeTxHash string, closeLogID uint) {
	// log.Info(utils.ToJson(ad.Details))
	for _, details := range ad.Details {
		if since == details.Since {
			details.ClosedAt = closedAt
			details.ClosedTxHash = closeTxHash
			details.ClosedLogId = closeLogID
		}
	}
}

func (ad *AccountData) detailsAssigned() bool {
	// if the account details is not assigned then
	// for some other credit manager, when getting remaining transfers
	// we should skip for this account as it can be different credit manager
	// for which detail might be missing currently
	return len(ad.Details) != 0
}

// process the transfer events for from <= block < to
func (ad *AccountData) GetRemainingTransfer(cm string, from, to int64) map[int64]map[string][]*schemas.TokenTransfer {
	// blockNum => sessionID => tokentranfers
	extraTokenTransfers := map[int64]map[string][]*schemas.TokenTransfer{}
	detailsInd := len(ad.Details) - 1
	blockInd := len(ad.blockNums) - 1
	// find first blockNum less than `to`
	for ; blockInd >= 0 && ad.blockNums[blockInd] >= to; blockInd-- {
	}
	for ; blockInd >= 0 && ad.blockNums[blockInd] >= from; blockInd-- {
		blockNum := ad.blockNums[blockInd]
		// find the account valid at
		for ; detailsInd >= 0 && ad.Details[detailsInd].Since > blockNum; detailsInd-- {

		}
		if detailsInd < 0 {
			break
		}
		details := ad.Details[detailsInd]
		if details.CreditManager != cm {
			continue
		}
		if blockNum >= details.Since && (details.ClosedAt == 0 || details.ClosedAt >= blockNum) {
			if extraTokenTransfers[blockNum] == nil {
				extraTokenTransfers[blockNum] = make(map[string][]*schemas.TokenTransfer)
			}
			for txHash, txs := range ad.transfers[blockNum] {
				deleteTxHash := false
				for _, transfer := range txs {
					// transfer is less than the txhash for open of credit account
					if (transfer.BlockNum == details.Since &&
						(transfer.LogID < details.OpenLogId && transfer.TxHash != details.OpenTxHash)) ||
						// transfer is more than the txhash for open of credit account
						(details.ClosedAt != 0 && transfer.BlockNum == details.ClosedAt &&
							(transfer.LogID > details.ClosedLogId && transfer.TxHash != details.ClosedTxHash)) {
						continue
					}
					extraTokenTransfers[blockNum][details.SessionID] = append(
						extraTokenTransfers[blockNum][details.SessionID], transfer)
					deleteTxHash = true
				}
				if deleteTxHash {
					delete(ad.transfers[blockNum], txHash)
				}
			}
		}
	}
	return extraTokenTransfers
}

func (ad *AccountData) GetNoSessionTxs() (noSessionTxs []string) {
	for _, txs := range ad.transfers {
		for txHash := range txs {
			noSessionTxs = append(noSessionTxs, txHash)
		}
	}
	return
}

//
type DirectTransferManager struct {
	// account => blockNum => txhash => transfers
	accountToData map[string]*AccountData
	// txhash to account which transferred asset
	txHashToAccounts map[string][]string
	mu               *sync.Mutex
	txToTransfers    map[string][]*schemas.TokenTransfer
}

func NewDirectTransferManager() *DirectTransferManager {
	return &DirectTransferManager{
		accountToData:    make(map[string]*AccountData),
		mu:               &sync.Mutex{},
		txToTransfers:    make(map[string][]*schemas.TokenTransfer),
		txHashToAccounts: make(map[string][]string),
	}
}

func (mgr *DirectTransferManager) AddTokenTransfer(tt *schemas.TokenTransfer) {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()
	if tt.IsFromAccount {
		if mgr.accountToData[tt.From] == nil {
			mgr.accountToData[tt.From] = newAccountData(tt.From)
		}
		mgr.txHashToAccounts[tt.TxHash] = append(mgr.txHashToAccounts[tt.TxHash], tt.From)
		mgr.accountToData[tt.From].process(tt)
	}
	if tt.IsToAccount {
		if mgr.accountToData[tt.To] == nil {
			mgr.accountToData[tt.To] = newAccountData(tt.To)
		}
		mgr.txHashToAccounts[tt.TxHash] = append(mgr.txHashToAccounts[tt.TxHash], tt.To)
		mgr.accountToData[tt.To].process(tt)
	}
	mgr.txToTransfers[tt.TxHash] = append(mgr.txToTransfers[tt.TxHash], tt)
}

func (mgr *DirectTransferManager) AddAccountDetails(sessionData *SessionData) {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()
	account := sessionData.Account
	if mgr.accountToData[account] == nil {
		mgr.accountToData[account] = newAccountData(account)
	}
	mgr.accountToData[account].AddDetails(sessionData)
}

func (mgr *DirectTransferManager) CloseAccountDetails(account string, since, closedAt int64, closeTxHash string, logID uint) {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()
	mgr.accountToData[account].SetClose(since, closedAt, closeTxHash, logID)
}

func (mgr *DirectTransferManager) CheckTokenTransfer(cm string, from, to int64) map[int64]map[string][]*schemas.TokenTransfer {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()
	result := map[int64]map[string][]*schemas.TokenTransfer{}
	for _, dataMdl := range mgr.accountToData {
		if !dataMdl.detailsAssigned() {
			continue
		}
		remainingTransfers := dataMdl.GetRemainingTransfer(cm, from, to)
		// if len(noSessionTxs) > 0{
		// 	log.Info(cm, dataMdl.Address, from, to, noSessionTxs)
		// 	log.Info(utils.ToJson(dataMdl.Details))
		// }
		for blockNum, data := range remainingTransfers {
			if result[blockNum] == nil {
				result[blockNum] = make(map[string][]*schemas.TokenTransfer)
			}
			for sessionID, transfers := range data {
				result[blockNum][sessionID] = transfers
			}
		}
	}
	return result
}

func (mgr *DirectTransferManager) DeleteTxHash(blockNum int64, txHash string) {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()
	for _, account := range mgr.txHashToAccounts[txHash] {
		mgr.accountToData[account].deleteTxHash(blockNum, txHash)
	}
}

func (mgr *DirectTransferManager) Clear() {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()
	for _, data := range mgr.accountToData {
		data.Clear()
	}
	mgr.txHashToAccounts = make(map[string][]string)
	mgr.txToTransfers = make(map[string][]*schemas.TokenTransfer)
}

func (mgr *DirectTransferManager) Init() {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()
	for _, accountData := range mgr.accountToData {
		accountData.init()
	}
}

func (mgr *DirectTransferManager) GetNoSessionTxs() (tts map[string][]*schemas.TokenTransfer) {
	tts = make(map[string][]*schemas.TokenTransfer)
	for _, dataMdl := range mgr.accountToData {
		for _, txHash := range dataMdl.GetNoSessionTxs() {
			tts[txHash] = mgr.txToTransfers[txHash]
		}
	}
	log.Infof("len of nosessionTxs %d", len(tts))
	return
}
