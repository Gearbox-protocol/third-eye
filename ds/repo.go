package ds

import (
	"github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressor/mainnet"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"math/big"
)

type EngineI interface {
	SyncHandler()
	Sync(syncTill int64)
	UseThreads()
}

type RepositoryI interface {
	// sync adapters
	GetKit() *AdapterKit
	AddSyncAdapter(adapterI SyncAdapterI)
	InitChecks()
	GetChainId() uint
	// saving to the db
	Flush() error
	// adding block/timestamp
	SetBlock(blockNum int64)
	SetAndGetBlock(blockNum int64) *schemas.Block
	GetBlocks() map[int64]*schemas.Block
	GetTokenOracles() map[int16]map[string]*schemas.TokenOracle
	GetDisabledTokens() []*schemas.AllowedToken
	LoadBlocks(from, to int64)
	// credit account operations
	AddAccountOperation(accountOperation *schemas.AccountOperation)
	// for getting executeparser
	GetExecuteParser() ExecuteParserI
	// price feed/oracle funcs
	AddTokenOracle(*schemas.TokenOracle)
	AddPriceFeed(blockNum int64, pf *schemas.PriceFeed)
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
	AddToken(token string) *schemas.Token
	GetToken(addr string) *schemas.Token
	GetTokens() []string
	ConvertToBalanceWithMask(balances []mainnet.DataTypesTokenBalance, mask *big.Int) (*core.JsonBalance, error)
	// credit session funcs
	AddCreditSession(session *schemas.CreditSession, loadedFromDB bool, txHash string, logID uint)
	GetCreditSession(sessionId string) *schemas.CreditSession
	UpdateCreditSession(sessionId string, values map[string]interface{}) *schemas.CreditSession
	GetSessions() map[string]*schemas.CreditSession
	GetValueInCurrency(blockNum int64, version int16, token, currency string, amount *big.Int) *big.Int
	AddDieselToken(dieselToken, underlyingToken, pool string)
	// credit session snapshots funcs
	AddCreditSessionSnapshot(css *schemas.CreditSessionSnapshot)
	// dc
	GetDCWrapper() *DataCompressorWrapper
	AddDataCompressor(blockNum int64, addr string)
	// pools
	AddPoolStat(ps *schemas.PoolStat)
	AddPoolLedger(pl *schemas.PoolLedger)
	GetPoolUniqueUserLen(pool string) int
	IsDieselToken(token string) bool
	// weth
	SetWETHAddr(address string)
	GetWETHAddr() string
	GetUSDCAddr() string
	GetGearTokenAddr() string
	// credit manager
	AddAccountTokenTransfer(tt *schemas.TokenTransfer)
	AddCreditManagerToFilter(cmAddr, cfAddr string)
	GetMask(blockNum int64, cmAddr, accountAddr string, version int16) *big.Int
	AddCreditManagerStats(cms *schemas.CreditManagerStat)
	GetCMState(cmAddr string) *schemas.CreditManagerState
	GetUnderlyingDecimal(cmAddr string) int8
	AddRepayOnCM(blockNum int64, cm string, pnl schemas.PnlOnRepay)
	GetRepayOnCM(blockNum int64, cm string) *schemas.PnlOnRepay
	AddParameters(logID uint, txHash string, params *schemas.Parameters, token string)
	AddFastCheckParams(logID uint, txHash, cm, creditFilter string, fcParams *schemas.FastCheckParams)
	AfterSync(blockNum int64)
	GetAccountManager() *AccountTokenManager
	AddAccountAddr(account string)
	// dao
	AddDAOOperation(operation *schemas.DAOOperation)
	CalCurrentTreasuryValue(syncTill int64)
	AddTreasuryTransfer(blockNum int64, logID uint, token string, amount *big.Int)
	RecentEventMsg(blockNum int64, msg string, args ...interface{})
	//
	// oracle and uni
	AddUniswapPrices(prices *schemas.UniPoolPrices)
	GetYearnFeedAddrs() []string
	AddTokenFeed(feedType string, token, oracle, feed string, discoveredAt int64, version int16)
	LoadLastDebtSync() int64
	LoadLastAdapterSync() int64
	Clear()
	// multicall
	GetUniPricesByToken(token string) []*schemas.UniPoolPrices
	AddUniPoolsForToken(blockNum int64, token string)
	AddUniPriceAndChainlinkRelation(relation *schemas.UniPriceAndChainlink)
	AddLastSyncForToken(token string, lastSync int64)
	// for testing
	AddTokenObj(token *schemas.Token)
	PrepareSyncAdapter(adapter *SyncAdapter) SyncAdapterI
}
