package ds

import (
	"math/big"
	"time"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/core/schemas/schemas_v3"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg/redstone"
	"github.com/Gearbox-protocol/third-eye/ds/dc_wrapper"
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"
)

type DummyRepo struct {
	PFs []*schemas.PriceFeed
}

func (DummyRepo) Init() {

}
func (DummyRepo) GetDB() *gorm.DB {
	return nil
}

// sync adapters
func (DummyRepo) GetAdapter(addr string) SyncAdapterI {
	return nil
}

func (DummyRepo) GetAdapterAddressByName(name string) []string {
	return nil
}

func (DummyRepo) AddSyncAdapter(adapterI SyncAdapterI) {
}
func (DummyRepo) GetChainId() uint {
	return 0
}

// saving to the db
func (DummyRepo) Flush(syncTill int64) error {
	return nil
}

// adding block/timestamp
func (DummyRepo) SetBlock(blockNum int64) {
}
func (DummyRepo) SetAndGetBlock(blockNum int64) *schemas.Block {
	return nil
}
func (DummyRepo) GetBlocks() map[int64]*schemas.Block {
	return nil
}
func (DummyRepo) GetMainTokenOracles() map[schemas.PriceOracleT]map[string]*schemas.TokenOracle {
	return nil
}
func (DummyRepo) GetDisabledTokens() []*schemas.AllowedToken {
	return nil
}
func (DummyRepo) LoadBlocks(from, to int64) {
}

// credit account operations
func (DummyRepo) AddAccountOperation(accountOperation *schemas.AccountOperation) {
}

// for getting executeparser
func (DummyRepo) GetExecuteParser() ExecuteParserI {
	return nil
}

// price feed/oracle funcs
func (DummyRepo) DirectlyAddTokenOracle(tokenOracle *schemas.TokenOracle) {
}
func (DummyRepo) GetPriceInUSD(blockNum int64, _ string, tokenAddrs string) *big.Int {
	return nil
}
func (DummyRepo) GetActivePriceOracleByBlockNum(blockNum int64) (schemas.PriceOracleT, core.VersionType, error) {
	return "", core.VersionType{}, nil
}
func (r *DummyRepo) AddPriceFeed(pf *schemas.PriceFeed) {
	r.PFs = append(r.PFs, pf)
}

// token funcs
func (DummyRepo) AddAllowedProtocol(logID uint, txHash, creditFilter string, p *schemas.Protocol) {
}
func (DummyRepo) DisableProtocol(blockNum int64, logID uint, txHash, cm, creditFilter, protocol string) {
}
func (DummyRepo) AddAllowedToken(logID uint, txHash, creditFilter string, atoken *schemas.AllowedToken) {
}
func (DummyRepo) DisableAllowedToken(blockNum int64, logID uint, txHash string, creditManager, creditFilter, token string) {
}

func (DummyRepo) GetFeedToTicker(feed string, _ string) common.Address {
	return core.NULL_ADDR
}

func (DummyRepo) AddFeedToTicker(feed string, ticker common.Address) {
}

// v2
func (DummyRepo) AddAllowedTokenV2(logID uint, txHash, creditFilter string, atoken *schemas.AllowedToken) {
}
func (DummyRepo) UpdateLimits(logID uint, txHash, creditConfigurator string, params *schemas.Parameters) {
}
func (DummyRepo) UpdateFees(logID uint, txHash, creditConfigurator string, params *schemas.Parameters) {
}
func (DummyRepo) UpdateEmergencyLiqDiscount(logID uint, txHash, creditConfigurator string, params *schemas.Parameters) {
}
func (DummyRepo) TransferAccountAllowed(*schemas.TransferAccountAllowed) {
}
func (DummyRepo) GetPricesInUSD(blockNum int64, _ string, tokenAddrs []string) core.JsonFloatMap {
	return nil
}

func (DummyRepo) GetToken(addr string) *schemas.Token {
	return nil
}
func (DummyRepo) GetTokens() []string {
	return nil
}

// credit session funcs
func (DummyRepo) AddCreditSession(session *schemas.CreditSession, loadedFromDB bool, txHash string, logID uint) {
}
func (DummyRepo) GetCreditSession(sessionId string) *schemas.CreditSession {
	return nil
}
func (DummyRepo) UpdateCreditSession(sessionId string, values map[string]interface{}) *schemas.CreditSession {
	return nil
}
func (DummyRepo) GetSessions() map[string]*schemas.CreditSession {
	return nil
}
func (DummyRepo) GetValueInCurrency(blockNum int64, version core.VersionType, token, currency string, amount *big.Int) (*big.Int, float64) {
	return nil, 0
}
func (DummyRepo) AddDieselToken(dieselToken, underlyingToken, pool string, version core.VersionType) {
}
func (DummyRepo) GetDieselTokens() map[string]*schemas.UTokenAndPool {
	return nil
}

// credit session snapshots funcs
func (DummyRepo) AddCreditSessionSnapshot(css *schemas.CreditSessionSnapshot) {
}

// dc
func (DummyRepo) GetDCWrapper() *dc_wrapper.DataCompressorWrapper {
	return nil
}

// pools
func (DummyRepo) AddPoolStat(ps *schemas.PoolStat) {
}
func (DummyRepo) AddDieselTransfer(dt *schemas.DieselTransfer) {
}

var Count int64

func (DummyRepo) AddRebaseDetailsForDB(transfer *schemas.RebaseDetailsForDB) {
	Count += 1
}
func (DummyRepo) AddPoolLedger(pl *schemas.PoolLedger) {
}
func (DummyRepo) GetPoolUniqueUserLen(pool string) int {
	return 0
}
func (DummyRepo) IsDieselToken(token string) bool {
	return false
}
func (DummyRepo) GetWETHAddr() string {
	return ""
}
func (DummyRepo) GetUSD() common.Address {
	return common.Address{}
}
func (DummyRepo) GetGearTokenAddr() string {
	return ""
}

// credit manager
func (DummyRepo) AddAccountTokenTransfer(tt *schemas.TokenTransfer) {
}
func (DummyRepo) AddCreditManagerStats(cms *schemas.CreditManagerStat) {
}
func (DummyRepo) GetCMState(cmAddr string) *schemas.CreditManagerState {
	return nil
}
func (DummyRepo) GetUnderlyingDecimal(cmAddr string) int8 {
	return 0
}
func (DummyRepo) AddRepayOnCM(cm string, pnl schemas.PnlOnRepay) {
}
func (DummyRepo) AddParameters(logID uint, txHash string, params *schemas.Parameters, token string) {
}
func (DummyRepo) AddFastCheckParams(logID uint, txHash, cm, creditFilter string, fcParams *schemas.FastCheckParams) {
}
func (DummyRepo) AfterSync(blockNum int64) {
}
func (DummyRepo) GetAccountManager() *DirectTransferManager {
	return nil
}
func (DummyRepo) AddAccountAddr(account string) {
}

// dao
func (DummyRepo) AddDAOOperation(operation *schemas.DAOOperation) {
}
func (DummyRepo) CalCurrentTreasuryValue(syncTill int64) {
}
func (DummyRepo) AddTreasuryTransfer(blockNum int64, logID uint, token string, amount *big.Int, operationTransfer bool) {
}
func (DummyRepo) RecentMsgf(headers log.RiskHeader, msg string, args ...interface{}) {
}

// oracle
func (DummyRepo) GetRetryFeedForDebts() []QueryPriceFeedI {
	return nil
}

// has mutex lock
func (DummyRepo) AddNewPriceOracleEvent(tokenOracle *schemas.TokenOracle, forChainlinkNewFeed ...bool) {
}

func (DummyRepo) LoadLastDebtSync() schemas.LastSync {
	return schemas.LastSync{}
}
func (DummyRepo) LoadLastAdapterSync() int64 {
	return 0
}
func (DummyRepo) Clear() {
}

// multicall
func (DummyRepo) ChainlinkPriceUpdatedAt(token string, blockNums []int64) {
}

// for testing
func (DummyRepo) AddTokenObj(token *schemas.Token) {
}
func (DummyRepo) PrepareSyncAdapter(adapter *SyncAdapter) SyncAdapterI {
	return nil
}

func (DummyRepo) AddTokenLTRamp(*schemas_v3.TokenLTRamp)   {}
func (DummyRepo) AddQuotaDetails(*schemas_v3.QuotaDetails) {}

func (DummyRepo) AddRelation(details *schemas.Relation) {}

func (DummyRepo) GetAccountQuotaMgr() *AccountQuotaMgr              { return nil }
func (DummyRepo) IsBlockRecent(block int64, dur time.Duration) bool { return false }
func (DummyRepo) GetRedStonemgr() redstone.RedStoneMgrI {
	return nil
}

type DieselBalance struct {
	BalanceBI *core.BigInt `gorm:"column:balance_bi"`
	Balance   float64      `gorm:"column:balance"`
	User      string       `gorm:"primaryKey;column:user_address"`
	Updated   bool         `gorm:"-"`
	// Diesel    string       `gorm:"primaryKey;column:diesel_sym"`
	Pool string `gorm:"primaryKey;column:pool"`
}

func (DieselBalance) TableName() string {
	return "diesel_balances"
}

func (DummyRepo) TokensValidAtBlock(string, int64) []*schemas.TokenOracle {
	return nil
}
func (DummyRepo) TokenAddrsValidAtBlock(string, int64) map[string]bool {
	return nil
}

func (DummyRepo) GetPrevPriceFeed(feed string) *schemas.PriceFeed {
	return nil
}

type QueryPriceFeedI interface {
	// TokensValidAtBlock(blockNum int64) []schemas.TokenAndMergedPFVersion
	GetPFType() string
	SyncAdapterI
	GetCalls(blockNum int64) (calls []multicall.Multicall2Call, isQueryable bool)
	ProcessResult(blockNum int64, results []multicall.Multicall2Result, token string, force ...bool) *schemas.PriceFeed
	// DisableToken(token string, disabledAt int64, pfVersion schemas.PFVersion)
	// AddToken(token string, discoveredAt int64, pfVersion schemas.PFVersion)
	// GetTokens() map[string]map[schemas.PFVersion][]int64
	GetRedstonePF() []*core.RedStonePF
}
