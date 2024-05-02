package account_manager

import (
	"strings"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"

	//
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// for getting directtoken transfer
// AddAccount is called from account factory
// AddAccountTokenTransfer sets the data in the ds.DirectTransferManager
type AccountManager struct {
	*ds.SyncAdapter
	node          *pkg.Node
	AccountHashes []common.Hash
	isAccount     map[string]bool
}

func NewAccountManager(addr string, discoveredAt int64, client core.ClientI, repo ds.RepositoryI) *AccountManager {
	syncAdapter := &ds.SyncAdapter{
		SyncAdapterSchema: &schemas.SyncAdapterSchema{
			Contract: &schemas.Contract{
				Address:      addr,
				DiscoveredAt: discoveredAt,
				FirstLogAt:   discoveredAt,
				ContractName: ds.AccountManager,
				Client:       client,
			},
			LastSync: discoveredAt - 1,
			V:        core.NewVersion(1),
		},
		Repo: repo,
	}
	return NewAccountManagerFromAdapter(
		syncAdapter,
	)
}

func NewAccountManagerFromAdapter(adapter *ds.SyncAdapter) *AccountManager {
	obj := &AccountManager{
		SyncAdapter: adapter,
		isAccount:   map[string]bool{},
	}
	obj.DataProcessType = ds.ViaQuery
	obj.node = &pkg.Node{
		Client: adapter.Client,
	}
	obj.populateInternalData()
	return obj
}

func (mdl *AccountManager) OnLog(txLog types.Log) {
	switch txLog.Topics[0] {
	case core.Topic("Transfer(address,address,uint256)"):
		from := common.HexToAddress(txLog.Topics[1].Hex()).Hex()
		to := common.HexToAddress(txLog.Topics[2].Hex()).Hex()
		value, ok := new(big.Int).SetString(common.BytesToHash(txLog.Data).Hex()[2:], 16)
		if !ok {
			log.Fatal("Failed parsing value")
		}
		tt := &schemas.TokenTransfer{
			BlockNum:      int64(txLog.BlockNumber),
			LogID:         txLog.Index,
			Token:         txLog.Address.Hex(),
			TxHash:        txLog.TxHash.Hex(),
			From:          from,
			To:            to,
			Amount:        (*core.BigInt)(value),
			IsFromAccount: mdl.isAccount[from],
			IsToAccount:   mdl.isAccount[to],
		}
		mdl.Repo.AddAccountTokenTransfer(tt)
	}
}

func (mdl *AccountManager) Query(queryTill int64) {
	queryFrom := mdl.GetLastSync() + 1
	log.Infof("Sync %s from %d to %d", mdl.GetName(), queryFrom, queryTill)
	tokenAddrs := mdl.Repo.GetTokens()
	hexAddrs := []common.Address{}
	for _, tokenAddr := range tokenAddrs {
		hexAddrs = append(hexAddrs, common.HexToAddress(tokenAddr))
	}
	if len(mdl.AccountHashes) == 0 {
		return
	}
	txLogs, err := mdl.node.GetLogsForTransfer(queryFrom, queryTill, hexAddrs, mdl.AccountHashes)
	if err != nil {
		if strings.Contains(err.Error(), "exceed max topics") && log.GetNetworkName(core.GetChainId(mdl.Client)) != log.GetBaseNet(core.GetChainId(mdl.Client)) { // testnet
			return
		}
		log.Fatal(err, "range ", queryFrom, queryTill, "tokenAddrs", len(tokenAddrs), "accountHashes", len(mdl.AccountHashes))
	}
	for _, txLog := range txLogs {
		mdl.OnLog(txLog)
	}
}

func (mdl *AccountManager) AddAccount(addr string) {
	accounts := mdl.getAccountAddrs()
	mdl.Details["accounts"] = append(accounts, addr)
	mdl.populateInternalData()
}

func (mdl *AccountManager) populateInternalData() {
	accountHashes := []common.Hash{}
	for _, accountAddr := range mdl.getAccountAddrs() {
		mdl.isAccount[accountAddr] = true
		accountHashes = append(accountHashes, common.HexToHash(accountAddr))
	}
	mdl.AccountHashes = accountHashes
}

func (mdl *AccountManager) getAccountAddrs() []string {
	if mdl.Details == nil {
		mdl.Details = make(map[string]interface{})
	}
	var accountAddrs []string
	if mdl.Details["accounts"] != nil {
		accountAddrs = utils.ConvertToListOfString(mdl.Details["accounts"])
	}
	return accountAddrs
}

func (mdl *AccountManager) AfterSyncHook(syncedTill int64) {
	mdl.Repo.GetAccountManager().Init()
	mdl.SyncAdapter.AfterSyncHook(syncedTill)
}
