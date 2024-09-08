package ds

import (
	"math/big"
	"time"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/core/schemas/schemas_v3"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg/redstone"
	"github.com/Gearbox-protocol/third-eye/ds/dc_wrapper"
	"gorm.io/gorm"
)

// type PriceInUSDType bool

// func (z *PriceInUSDType) MarshalJSON() ([]byte, error) {
// 	return []byte(fmt.Sprintf("%v", *z)), nil
// }

// func (z *PriceInUSDType) UnmarshalJSON(b []byte) error {
// 	str := strings.Trim(string(b), "\"")
// 	*z = (str == "true")
// 	return nil

// }
// func (z PriceInUSDType) MarshalText() (text []byte, err error) {
// 	return z.MarshalJSON()
// }

// func (s *PriceInUSDType) UnmarshalText(text []byte) error {
// 	return s.UnmarshalJSON(text)
// }

type EngineI interface {
	SyncHandler()
	Sync(syncTill int64)
	LastSyncedBlock() (int64, uint64)
	UseThreads()
}

type RepositoryI interface {
	GetDB() *gorm.DB
	GetRedStonemgr() redstone.RedStoneMgrI
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
	GetTokenOracles() map[schemas.PFVersion]map[string]*schemas.TokenOracle
	DirectlyAddTokenOracle(tokenOracle *schemas.TokenOracle)
	AddNewPriceOracleEvent(tokenOracle *schemas.TokenOracle, bounded bool, forChainlinkNewFeed ...bool)
	//
	GetPrice(token string) *big.Int
	AddPriceFeed(pf *schemas.PriceFeed)
	// token funcs
	AddAllowedProtocol(logID uint, txHash, creditFilter string, p *schemas.Protocol)
	DisableProtocol(blockNum int64, logID uint, txHash, cm, creditFilter, protocol string)
	AddAllowedToken(logID uint, txHash, creditFilter string, atoken *schemas.AllowedToken)
	DisableAllowedToken(blockNum int64, logID uint, txHash string, creditManager, creditFilter, token string)
	// v2
	AddAllowedTokenV2(logID uint, txHash, creditFilter string, atoken *schemas.AllowedToken)
	UpdateLimits(logID uint, txHash, creditConfigurator string, params *schemas.Parameters)
	UpdateEmergencyLiqDiscount(logID uint, txHash, creditConfigurator string, params *schemas.Parameters)
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
	GetValueInCurrency(blockNum int64, version core.VersionType, token, currency string, amount *big.Int) *big.Int
	AddDieselToken(dieselToken, underlyingToken, pool string, version core.VersionType)
	GetDieselTokens() map[string]*schemas.UTokenAndPool
	// credit session snapshots funcs
	AddCreditSessionSnapshot(css *schemas.CreditSessionSnapshot)
	// dc
	GetDCWrapper() *dc_wrapper.DataCompressorWrapper
	// pools
	AddPoolStat(ps *schemas.PoolStat)
	AddDieselTransfer(dt *schemas.DieselTransfer)
	AddRebaseDetailsForDB(transfer *schemas.RebaseDetailsForDB)
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
	GetRetryFeedForDebts() []QueryPriceFeedI
	//
	LoadLastDebtSync() int64
	LoadLastAdapterSync() int64
	Clear()
	// multicall
	ChainlinkPriceUpdatedAt(token string, blockNums []int64)
	// for testing
	AddTokenObj(token *schemas.Token)
	PrepareSyncAdapter(adapter *SyncAdapter) SyncAdapterI
	//
	GetTokenFromSdk(symbol string) string

	// v3 events
	AddTokenLTRamp(*schemas_v3.TokenLTRamp)
	AddQuotaDetails(*schemas_v3.QuotaDetails)
	GetAccountQuotaMgr() *AccountQuotaMgr
	IsBlockRecent(block int64, dur time.Duration) bool
}

func IsTestnet(client core.ClientI) bool {
	chainid := core.GetChainId(client)
	return log.GetNetworkName(chainid) != log.GetBaseNet(chainid)
}
