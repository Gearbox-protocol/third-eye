package ds

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds/dc_wrapper"
)

type EngineI interface {
	SyncHandler()
	Sync(syncTill int64)
	LastSyncedBlock() int64
	UseThreads()
}

type RepositoryI interface {
	Init()
	// sync adapters
	GetAdapter(addr string) SyncAdapterI
	GetAdapterAddressByName(name string) []string
	AddSyncAdapter(adapterI SyncAdapterI)
	// saving to the db
	Flush(syncTill int64) error
	// adding block/timestamp
	SetBlock(blockNum int64)
	SetAndGetBlock(blockNum int64) *schemas.Block
	GetBlocks() map[int64]*schemas.Block
	GetDisabledTokens() []*schemas.AllowedToken
	LoadBlocks(from, to int64)
	// credit account operations
	AddAccountOperation(accountOperation *schemas.AccountOperation)
	// for getting executeparser
	GetExecuteParser() ExecuteParserI
	// price feed/oracle funcs
	GetTokenOracles() map[int16]map[string]*schemas.TokenOracle
	// if returned value is nil, it means that token oracle hasn't been added yet.
	GetOracleForV2Token(token string) *schemas.TokenOracle
	DirectlyAddTokenOracle(tokenOracle *schemas.TokenOracle)
	AddNewPriceOracleEvent(tokenOracle *schemas.TokenOracle, bounded bool)
	//
	AddPriceFeed(pf *schemas.PriceFeed)
	// token funcs
	AddAllowedProtocol(logID uint, txHash, creditFilter string, p *schemas.Protocol)
	DisableProtocol(blockNum int64, logID uint, txHash, cm, creditFilter, protocol string)
	AddAllowedToken(logID uint, txHash, creditFilter string, atoken *schemas.AllowedToken)
	DisableAllowedToken(blockNum int64, logID uint, txHash string, creditManager, creditFilter, token string)
	// v2
	AddAllowedTokenV2(logID uint, txHash, creditFilter string, atoken *schemas.AllowedToken)
	UpdateLimits(logID uint, txHash, creditConfigurator string, params *schemas.Parameters)
	UpdateFees(logID uint, txHash, creditConfigurator string, params *schemas.Parameters)
	TransferAccountAllowed(*schemas.TransferAccountAllowed)
	GetPricesInUSD(blockNum int64, tokenAddrs []string) core.JsonFloatMap
	//
	GetToken(addr string) *schemas.Token
	GetTokens() []string
	// credit session funcs
	AddCreditSession(session *schemas.CreditSession, loadedFromDB bool, txHash string, logID uint)
	GetCreditSession(sessionId string) *schemas.CreditSession
	UpdateCreditSession(sessionId string, values map[string]interface{}) *schemas.CreditSession
	GetSessions() map[string]*schemas.CreditSession
	GetValueInCurrency(blockNum int64, version int16, token, currency string, amount *big.Int) *big.Int
	AddDieselToken(dieselToken, underlyingToken, pool string)
	GetDieselTokens() map[string]*schemas.UTokenAndPool
	// credit session snapshots funcs
	AddCreditSessionSnapshot(css *schemas.CreditSessionSnapshot)
	// dc
	GetDCWrapper() *dc_wrapper.DataCompressorWrapper
	// pools
	AddPoolStat(ps *schemas.PoolStat)
	AddDieselTransfer(dt *schemas.DieselTransfer)
	AddPoolLedger(pl *schemas.PoolLedger)
	GetPoolUniqueUserLen(pool string) int
	IsDieselToken(token string) bool
	GetWETHAddr() string
	GetUSDCAddr() string
	GetGearTokenAddr() string
	// credit manager
	AddAccountTokenTransfer(tt *schemas.TokenTransfer)
	AddCreditManagerStats(cms *schemas.CreditManagerStat)
	GetCMState(cmAddr string) *schemas.CreditManagerState
	GetUnderlyingDecimal(cmAddr string) int8
	AddRepayOnCM(cm string, pnl schemas.PnlOnRepay)
	AddParameters(logID uint, txHash string, params *schemas.Parameters, token string)
	AddFastCheckParams(logID uint, txHash, cm, creditFilter string, fcParams *schemas.FastCheckParams)
	AfterSync(blockNum int64)
	GetAccountManager() *DirectTransferManager
	AddAccountAddr(account string)
	// dao
	AddDAOOperation(operation *schemas.DAOOperation)
	CalCurrentTreasuryValue(syncTill int64)
	AddTreasuryTransfer(blockNum int64, logID uint, token string, amount *big.Int, operationTransfer bool)
	RecentMsgf(headers log.RiskHeader, msg string, args ...interface{})
	//
	// oracle and uni
	GetYearnFeedAddrs() []string
	//
	LoadLastDebtSync() int64
	LoadLastAdapterSync() int64
	Clear()
	// multicall
	ChainlinkPriceUpdatedAt(token string, blockNums []int64)
	// for testing
	AddTokenObj(token *schemas.Token)
	PrepareSyncAdapter(adapter *SyncAdapter) SyncAdapterI
}
