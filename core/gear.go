package core

import (
	"github.com/Gearbox-protocol/third-eye/artifacts/dataCompressor/mainnet"
	"math/big"
)

type EngineI interface {
	SyncHandler()
	Sync(syncTill int64)
	UseThreads()
}

type Protocol struct {
	Id            string `gorm:"primaryKey;column:id;autoincrement:true" json:"-"`
	Protocol      string `gorm:"column:protocol" json:"protocol"`
	Adapter       string `gorm:"column:adapter" json:"adapter"`
	BlockNumber   int64  `gorm:"column:block_num" json:"blockNum"`
	CreditManager string `gorm:"column:credit_manager" json:"creditManager"`
}

func (Protocol) TableName() string {
	return "allowed_protocols"
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
	SetAndGetBlock(blockNum int64) *Block
	GetBlocks() map[int64]*Block
	GetTokenOracles() map[int16]map[string]*TokenOracle
	GetDisabledTokens() []*AllowedToken
	LoadBlocks(from, to int64)
	// credit account operations
	AddAccountOperation(accountOperation *AccountOperation)
	// for getting executeparser
	GetExecuteParser() ExecuteParserI
	// price feed/oracle funcs
	AddTokenOracle(*TokenOracle)
	AddPriceFeed(blockNum int64, pf *PriceFeed)
	// token funcs
	AddAllowedProtocol(logID uint, txHash, creditFilter string, p *Protocol)
	DisableProtocol(blockNum int64, logID uint, txHash, cm, creditFilter, protocol string)
	AddAllowedToken(logID uint, txHash, creditFilter string, atoken *AllowedToken)
	DisableAllowedToken(blockNum int64, logID uint, txHash string, creditManager, creditFilter, token string)
	// v2
	AddAllowedTokenV2(logID uint, txHash, creditFilter string, atoken *AllowedToken)
	UpdateLimits(logID uint, txHash, creditConfigurator string, params *Parameters)
	UpdateFees(logID uint, txHash, creditConfigurator string, params *Parameters)
	TransferAccountAllowed(*TransferAccountAllowed)
	GetPricesInUSD(blockNum int64, tokenAddrs []string) JsonFloatMap
	//
	AddToken(token string) *Token
	GetToken(addr string) *Token
	GetTokens() []string
	ConvertToBalanceWithMask(balances []mainnet.DataTypesTokenBalance, mask *big.Int) (*JsonBalance, error)
	// credit session funcs
	AddCreditSession(session *CreditSession, loadedFromDB bool, txHash string, logID uint)
	GetCreditSession(sessionId string) *CreditSession
	UpdateCreditSession(sessionId string, values map[string]interface{}) *CreditSession
	GetSessions() map[string]*CreditSession
	GetValueInCurrency(blockNum int64, version int16, token, currency string, amount *big.Int) *big.Int
	AddDieselToken(dieselToken, underlyingToken, pool string)
	// credit session snapshots funcs
	AddCreditSessionSnapshot(css *CreditSessionSnapshot)
	// dc
	GetDCWrapper() *DataCompressorWrapper
	AddDataCompressor(blockNum int64, addr string)
	// pools
	AddPoolStat(ps *PoolStat)
	AddPoolLedger(pl *PoolLedger)
	GetPoolUniqueUserLen(pool string) int
	IsDieselToken(token string) bool
	// weth
	SetWETHAddr(address string)
	GetWETHAddr() string
	GetUSDCAddr() string
	GetGearTokenAddr() string
	// credit manager
	AddAccountTokenTransfer(tt *TokenTransfer)
	AddCreditManagerToFilter(cmAddr, cfAddr string)
	GetMask(blockNum int64, cmAddr, accountAddr string, version int16) *big.Int
	AddCreditManagerStats(cms *CreditManagerStat)
	GetCMState(cmAddr string) *CreditManagerState
	GetUnderlyingDecimal(cmAddr string) int8
	AddRepayOnCM(blockNum int64, cm string, pnl PnlOnRepay)
	GetRepayOnCM(blockNum int64, cm string) *PnlOnRepay
	AddParameters(logID uint, txHash string, params *Parameters, token string)
	AddFastCheckParams(logID uint, txHash, cm, creditFilter string, fcParams *FastCheckParams)
	AfterSync(blockNum int64)
	GetAccountManager() *AccountTokenManager
	AddAccountAddr(account string)
	// dao
	AddDAOOperation(operation *DAOOperation)
	CalCurrentTreasuryValue(syncTill int64)
	AddTreasuryTransfer(blockNum int64, logID uint, token string, amount *big.Int)
	RecentEventMsg(blockNum int64, msg string, args ...interface{})
	//
	// oracle and uni
	AddUniswapPrices(prices *UniPoolPrices)
	GetYearnFeedAddrs() []string
	LoadLastDebtSync() int64
	LoadLastAdapterSync() int64
	Clear()
	// multicall
	GetUniPricesByToken(token string) []*UniPoolPrices
	AddUniPoolsForToken(blockNum int64, token string)
	AddUniPriceAndChainlinkRelation(relation *UniPriceAndChainlink)
	AddLastSyncForToken(token string, lastSync int64)
	// for testing
	AddTokenObj(token *Token)
	PrepareSyncAdapter(adapter *SyncAdapter) SyncAdapterI
}

type GearBalance struct {
	Balance *BigInt `gorm:"column:balance"`
	Updated bool    `gorm:"-"`
	User    string  `gorm:"column:user_address;primaryKey"`
}

func (GearBalance) TableName() string {
	return "gear_balances"
}
