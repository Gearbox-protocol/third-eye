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

type HashAndAccounts struct {
	Hashs     []common.Hash
	Accounts  []string
	isAccount map[string]bool
}

func NewHashAndAddrs() HashAndAccounts {
	return HashAndAccounts{
		Hashs:     []common.Hash{},
		Accounts:  []string{},
		isAccount: map[string]bool{},
	}
}

func (ha *HashAndAccounts) Add(addr string) {
	if !ha.isAccount[addr] {
		ha.Accounts = append(ha.Accounts, addr)
		ha.Hashs = append(ha.Hashs, common.HexToHash(addr))
		ha.isAccount[addr] = true
	}
}

// for getting directtoken transfer
// AddAccount is called from account factory
// AddAccountTokenTransfer sets the data in the ds.DirectTransferManager
type AccountManager struct {
	*ds.SyncAdapter
	node     *pkg.Node
	accounts HashAndAccounts
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
		accounts:    NewHashAndAddrs(),
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
			IsFromAccount: mdl.accounts.isAccount[from],
			IsToAccount:   mdl.accounts.isAccount[to],
		}
		mdl.Repo.AddAccountTokenTransfer(tt)
	}
}

func (mdl *AccountManager) Query(queryTill int64) {
	queryFrom := mdl.GetLastSync() + 1
	tokenAddrs := mdl.Repo.GetTokens()
	hexAddrs := []common.Address{}
	for _, tokenAddr := range tokenAddrs {
		hexAddrs = append(hexAddrs, common.HexToAddress(tokenAddr))
	}
	if len(mdl.accounts.Accounts) == 0 {
		return
	}
	txLogs, err := mdl.node.GetLogsForTransfer(queryFrom, queryTill, hexAddrs, mdl.accounts.Hashs)
	log.Infof("Sync %s from %d to %d: %d. tokensAddrs %d accountHashes %d", mdl.GetName(), queryFrom, queryTill, len(txLogs), len(tokenAddrs), len(mdl.accounts.Accounts))
	if err != nil {
		if strings.Contains(err.Error(), "exceed max topics") && ds.IsTestnet(mdl.Client) { // anvil failure
			return
		}
		log.Fatal(err, "range ", queryFrom, queryTill, "tokenAddrs", len(tokenAddrs), "accountHashes", len(mdl.accounts.Accounts))
	}
	for _, txLog := range txLogs {
		mdl.OnLog(txLog)
	}
}

func (mdl *AccountManager) AddAccount(addr string) {
	mdl.accounts.Add(addr)
}

func (mdl *AccountManager) populateInternalData() {
	if mdl.Details == nil {
		mdl.Details = make(map[string]interface{})
	}
	var accountAddrs []string
	if mdl.Details["accounts"] != nil {
		accountAddrs = utils.ConvertToListOfString(mdl.Details["accounts"])
	}
	for _, addr := range accountAddrs {
		mdl.accounts.Add(addr)
	}
}

func (mdl *AccountManager) AfterSyncHook(syncedTill int64) {
	mdl.Details["accounts"] = mdl.accounts.Accounts
	mdl.Repo.GetAccountManager().Init()
	mdl.SyncAdapter.AfterSyncHook(syncedTill)
}
