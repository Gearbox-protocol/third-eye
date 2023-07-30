package debts

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/ds"
	"gorm.io/gorm"
)

type DebtEngine struct {
	repo             ds.RepositoryI
	db               *gorm.DB
	client           core.ClientI
	config           *config.Config
	lastCSS          map[string]*schemas.CreditSessionSnapshot
	tokenLastPrice   map[string]*schemas.PriceFeed
	tokenLastPriceV2 map[string]*schemas.PriceFeed
	//// credit_manager -> token -> liquidity threshold
	allowedTokensThreshold map[string]map[string]*core.BigInt
	poolLastInterestData   map[string]*schemas.PoolInterestData
	debts                  []*schemas.Debt
	tvlSnapshots           []*schemas.TvlSnapshots
	lastDebts              map[string]*schemas.Debt
	currentDebts           []*schemas.CurrentDebt
	liquidableBlockTracker map[string]*schemas.LiquidableAccount
	// cm to paramters
	lastParameters    map[string]*schemas.Parameters
	isTesting         bool
	farmingCalc       *FarmingCalculator
	lastTvlSnapshot   *schemas.TvlSnapshots
	lastRebaseDetails *schemas.RebaseDetailsForDB
}

func GetDebtEngine(db *gorm.DB, client core.ClientI, config *config.Config, repo ds.RepositoryI, testing bool) ds.DebtEngineI {
	return &DebtEngine{
		repo:                   repo,
		db:                     db,
		client:                 client,
		config:                 config,
		lastCSS:                make(map[string]*schemas.CreditSessionSnapshot),
		tokenLastPrice:         make(map[string]*schemas.PriceFeed),
		tokenLastPriceV2:       make(map[string]*schemas.PriceFeed),
		allowedTokensThreshold: make(map[string]map[string]*core.BigInt),
		poolLastInterestData:   make(map[string]*schemas.PoolInterestData),
		lastDebts:              make(map[string]*schemas.Debt),
		liquidableBlockTracker: make(map[string]*schemas.LiquidableAccount),
		lastParameters:         make(map[string]*schemas.Parameters),
		isTesting:              testing,
		farmingCalc:            NewFarmingCalculator(core.GetChainId(client), testing),
	}
}

func NewDebtEngine(db *gorm.DB, client core.ClientI, config *config.Config, repo ds.RepositoryI) ds.DebtEngineI {
	return GetDebtEngine(db, client, config, repo, false)
}

func (eng *DebtEngine) ProcessBackLogs() {
	// NOTE: while syncing from scratch for some adapter disable the debt engine
	// as it might happen that some of the components for calculating debt are missing
	// check if adapters are synchronised.
	if !eng.AreActiveAdapterSynchronized() {
		eng.config.DisableDebtEngine = true
	}
	if eng.config.DisableDebtEngine {
		return
	}
	log.Info("Debt engine started")
	// synced till
	lastDebtSynced := eng.repo.LoadLastDebtSync()
	eng.loadLastTvlSnapshot()
	eng.loadLastCSS(lastDebtSynced)
	eng.loadLastRebaseDetails(lastDebtSynced)
	eng.loadTokenLastPrice(lastDebtSynced)
	eng.loadAllowedTokenThreshold(lastDebtSynced)
	eng.loadPoolLastInterestData(lastDebtSynced)
	eng.loadLastDebts(lastDebtSynced)
	eng.loadParameters(lastDebtSynced)
	eng.loadLiquidableAccounts(lastDebtSynced)
	// process blocks for calculating debts
	adaptersSyncedTill := eng.repo.LoadLastAdapterSync()
	batchSize := eng.config.BatchSizeForHistory
	for ; lastDebtSynced+batchSize < adaptersSyncedTill; lastDebtSynced += batchSize {
		eng.processBlocksInBatch(lastDebtSynced, lastDebtSynced+batchSize)
	}
	eng.processBlocksInBatch(lastDebtSynced, adaptersSyncedTill)
}
func (eng *DebtEngine) loadLastTvlSnapshot() {
	lastTvlSnapshot := &schemas.TvlSnapshots{}
	if err := eng.db.Raw(`SELECT * FROM tvl_snapshots ORDER BY block_num DESC LIMIT 1`).Find(lastTvlSnapshot).Error; err != nil {
		log.Fatal(err)
	}
	eng.lastTvlSnapshot = lastTvlSnapshot
}

// load blocks from > and to <=
func (eng *DebtEngine) processBlocksInBatch(from, to int64) {
	if from == to {
		return
	}
	eng.repo.LoadBlocks(from, to)
	if len(eng.repo.GetBlocks()) > 0 {
		eng.CalculateDebtAndClear(to)
	}
}

// called for the engine/index.go and the debt engine
func (eng *DebtEngine) CalculateDebtAndClear(to int64) {
	if !eng.config.DisableDebtEngine {
		eng.CalculateDebt()
		eng.flushDebt(to)
		eng.CalCurrentDebts(to)
		eng.flushCurrentDebts(to)
	}
	eng.Clear()
}

func (eng *DebtEngine) Clear() {
	eng.debts = []*schemas.Debt{}
	// clear repo after calculating debt as debt uses repository for calculations
	eng.repo.Clear()
}

func (eng *DebtEngine) loadLiquidableAccounts(lastDebtSync int64) {
	defer utils.Elapsed("loadLiquidableAccounts")()
	data := []*schemas.LiquidableAccount{}
	query := `SELECT * FROM liquidable_accounts la JOIN credit_sessions cs ON la.session_id = cs.id WHERE cs.status not in (1,2);`
	err := eng.db.Raw(query).Find(&data).Error
	log.CheckFatal(err)
	for _, entry := range data {
		eng.liquidableBlockTracker[entry.SessionId] = entry
	}
}

func (eng *DebtEngine) addLiquidableAccount(sessionId string, newBlockNum int64) {
	liquidableAccount := eng.liquidableBlockTracker[sessionId]
	if liquidableAccount == nil {
		eng.liquidableBlockTracker[sessionId] = &schemas.LiquidableAccount{
			SessionId: sessionId,
			BlockNum:  newBlockNum,
			Updated:   true,
		}
	} else {
		liquidableAccount.BlockNum = newBlockNum
		liquidableAccount.Updated = true
	}
}

func (eng *DebtEngine) notifiedIfLiquidable(sessionId string, notified bool) {
	liquidableAccount := eng.liquidableBlockTracker[sessionId]
	liquidableAccount.NotifiedIfLiquidable = notified
	liquidableAccount.Updated = true
}

func (eng *DebtEngine) AreActiveAdapterSynchronized() bool {
	data := schemas.DebtSync{}
	query := "SELECT count(distinct last_sync) as last_calculated_at FROM sync_adapters where disabled=false"
	err := eng.db.Raw(query).Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	val := data.LastCalculatedAt <= 1
	if !val {
		log.Warn("DebtEngine disabled acitve adapters are not synchronised")
	}
	return val
}

type LiquidationTx struct {
	TxHash string `gorm:"column:tx_hash"`
}

func (eng *DebtEngine) GetLiquidationTx(sessionId string) string {
	if eng.isTesting {
		return ""
	}
	data := LiquidationTx{}
	query := `SELECT tx_hash from account_operations 
		WHERE session_id = ?  AND action like 'LiquidateCreditAccount%'`
	err := eng.db.Raw(query, sessionId).Find(&data).Error
	log.CheckFatal(err)
	return data.TxHash
}

func (eng *DebtEngine) GetDebts() core.Json {
	obj := core.Json{}
	obj["debts"] = eng.debts
	obj["currentDebts"] = eng.currentDebts
	return obj
}

func CompareBalance(a, b *core.BigInt, token *ds.CumIndexAndUToken) bool {
	precision := utils.GetPrecision(token.Symbol)
	return utils.AlmostSameBigInt(a.Convert(), b.Convert(), token.Decimals-precision)
}
