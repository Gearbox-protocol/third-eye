package debts

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/core/schemas/schemas_v3"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/ds"
	"gorm.io/gorm"
)

type DebtEngine struct {
	repo    ds.RepositoryI
	db      *gorm.DB
	client  core.ClientI
	config  *config.Config
	lastCSS map[string]*schemas.CreditSessionSnapshot

	//// credit_manager -> token -> liquidity threshold
	poolLastInterestData   map[string]*schemas.PoolInterestData
	debts                  []*schemas.Debt
	tvlSnapshots           []*schemas.TvlSnapshots
	lastDebts              map[string]*schemas.Debt
	currentDebts           []*schemas.CurrentDebt
	liquidableBlockTracker map[string]*schemas.LiquidableAccount
	// cm to paramters
	lastParameters       map[string]*schemas.Parameters
	isTesting            bool
	farmingCalc          *FarmingCalculator
	marketTolastTvlBlock map[string]int64
	lastRebaseDetails    *schemas.RebaseDetailsForDB
	// used for v3 calc account fields
	currentTs uint64
	v3DebtDetails
	tokenLTRamp  map[string]map[string]*schemas_v3.TokenLTRamp
	priceHandler *PriceHandler
}

func GetDebtEngine(db *gorm.DB, client core.ClientI, config *config.Config, repo ds.RepositoryI, testing bool) ds.DebtEngineI {
	return &DebtEngine{
		repo:                   repo,
		db:                     db,
		client:                 client,
		config:                 config,
		lastCSS:                make(map[string]*schemas.CreditSessionSnapshot),
		poolLastInterestData:   make(map[string]*schemas.PoolInterestData),
		lastDebts:              make(map[string]*schemas.Debt),
		liquidableBlockTracker: make(map[string]*schemas.LiquidableAccount),
		lastParameters:         make(map[string]*schemas.Parameters),
		isTesting:              testing,
		farmingCalc:            NewFarmingCalculator(core.GetChainId(client), testing),
		v3DebtDetails:          Newv3DebtDetails(),
		tokenLTRamp:            map[string]map[string]*schemas_v3.TokenLTRamp{},
		priceHandler:           NewPriceHandler(repo),
		marketTolastTvlBlock:   make(map[string]int64),
	}
}

func (eng *DebtEngine) InitTest() {
	eng.priceHandler.poTotokenOracle = eng.repo.GetTokenOracles()
	eng.priceHandler.init(eng.repo)
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
	// synced till
	lastSync := eng.repo.LoadLastDebtSync()
	minSynced := lastSync.Min()
	// lastDebtSynced = 227143579
	log.Info("Debt engine started, from", minSynced)
	eng.loadLastTvlSnapshot()
	eng.loadLastCSS(minSynced)
	eng.loadLastRebaseDetails(minSynced)
	eng.loadAllowedTokenThreshold(minSynced)
	eng.loadLastLTRamp(minSynced)
	eng.loadPoolLastInterestData(minSynced)
	eng.loadLastDebts(minSynced)
	eng.loadParameters(minSynced)
	eng.loadLiquidableAccounts(minSynced)
	//
	eng.priceHandler.load(minSynced, eng.db)
	// v3
	// eng.loadAccounQuotaInfo(lastDebtSynced, eng.db)
	eng.loadPoolQuotaDetails(minSynced, eng.db)
	//
	// process blocks for calculating debts
	adaptersSyncedTill := eng.repo.LoadLastAdapterSync()
	// adaptersSyncedTill = 227143580
	batchSize := eng.config.BatchSizeForHistory
	for ; minSynced+batchSize < adaptersSyncedTill; minSynced += batchSize {
		eng.processBlocksInBatch(minSynced, minSynced+batchSize, lastSync)
	}
	eng.processBlocksInBatch(minSynced, adaptersSyncedTill, lastSync)
}
func (eng *DebtEngine) loadLastTvlSnapshot() {
	tvlsnaps := []*schemas.TvlSnapshots{}
	if err := eng.db.Raw(`SELECT * FROM tvl_snapshots ORDER BY block_num,market DESC LIMIT 1`).Find(&tvlsnaps).Error; err != nil {
		log.Fatal(err)
	}
	for _, entry := range tvlsnaps {
		eng.marketTolastTvlBlock[entry.Market] = entry.BlockNum
	}
}

// load blocks from > and to <=
func (eng *DebtEngine) processBlocksInBatch(from, to int64, lastSync schemas.LastSync) {
	if from == to {
		return
	}
	eng.repo.LoadBlocks(from, to)
	if len(eng.repo.GetBlocks()) > 0 {
		eng.CalculateDebtAndClear(to, lastSync)
	}
}

// called for the engine/index.go and the debt engine
func (eng *DebtEngine) CalculateDebtAndClear(to int64, lastSync schemas.LastSync) {
	if !eng.config.DisableDebtEngine {
		eng.CalculateDebt()
		//
		tx := eng.db.Begin()
		eng.flushDebt(to, tx, lastSync)
		eng.flushTvl(to, tx, lastSync)
		if info := tx.Commit(); info.Error != nil {
			log.Fatal(info.Error)
		}
		eng.tvlSnapshots = []*schemas.TvlSnapshots{}
		eng.debts = []*schemas.Debt{}
		//
		if to > lastSync.Debt {
			eng.CalCurrentDebts(to)
			eng.flushCurrentDebts(to)
		}
	}
	eng.Clear()
}

func (eng *DebtEngine) Clear() {
	eng.debts = []*schemas.Debt{}
	// clear repo after calculating debt as debt uses repository for calculations
	eng.repo.Clear()
}

func (eng *DebtEngine) loadLiquidableAccounts(_lastDebtSync int64) {
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

// QueryPriceFeed is updated only till the lastFetchedBlock, not the syncTill that is provided to the aqfwrapper's aftersynchook from engine/index.go in the syncmodel. So, ignore that for updating the debts.
func (eng *DebtEngine) AreActiveAdapterSynchronized() bool {
	data := struct {
		LastSync int64 `json:"last_sync"`
	}{}
	query := `SELECT count(distinct last_sync) as last_sync FROM sync_adapters 
	WHERE disabled=false AND type NOT IN ('QueryPriceFeed','RebaseToken','Treasury','LMRewardsv2','LMRewardsv3','GearToken')`
	err := eng.db.Raw(query).Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	val := data.LastSync <= 1
	if !val {
		log.Warn("DebtEngine disabled active adapters are not synchronised")
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

func IsChangeMoreThanFraction(a, b *core.BigInt, diff *big.Float) bool {
	if a.Cmp(b) > 0 {
		return IsChangeMoreThanFraction(b, a, diff)
	}
	// b < a
	return utils.DiffMoreThanFraction(a.Convert(), b.Convert(), diff)
}
