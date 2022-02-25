package debts

import (
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/utils"
	"gorm.io/gorm"

	"math/big"
)

type DebtEngine struct {
	repo           core.RepositoryI
	db             *gorm.DB
	client         ethclient.ClientI
	config         *config.Config
	lastCSS        map[string]*core.CreditSessionSnapshot
	tokenLastPrice map[string]*core.PriceFeed
	//// credit_manager -> token -> liquidity threshold
	allowedTokensThreshold map[string]map[string]*core.BigInt
	poolLastInterestData   map[string]*core.PoolInterestData
	debts                  []*core.Debt
	lastDebts              map[string]*core.Debt
	currentDebts           []*core.CurrentDebt
	liquidableBlockTracker map[string]*core.LiquidableAccount
	// cm to paramters
	lastParameters map[string]*core.Parameters
	isTesting      bool
}

func GetDebtEngine(db *gorm.DB, client ethclient.ClientI, config *config.Config, repo core.RepositoryI, testing bool) core.DebtEngineI {
	return &DebtEngine{
		repo:                   repo,
		db:                     db,
		client:                 client,
		config:                 config,
		lastCSS:                make(map[string]*core.CreditSessionSnapshot),
		tokenLastPrice:         make(map[string]*core.PriceFeed),
		allowedTokensThreshold: make(map[string]map[string]*core.BigInt),
		poolLastInterestData:   make(map[string]*core.PoolInterestData),
		lastDebts:              make(map[string]*core.Debt),
		liquidableBlockTracker: make(map[string]*core.LiquidableAccount),
		lastParameters:         make(map[string]*core.Parameters),
		isTesting:              testing,
	}
}

func NewDebtEngine(db *gorm.DB, client ethclient.ClientI, config *config.Config, repo core.RepositoryI) core.DebtEngineI {
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
	lastDebtSync := eng.repo.LoadLastDebtSync()
	eng.loadLastCSS(lastDebtSync)
	eng.loadTokenLastPrice(lastDebtSync)
	eng.loadAllowedTokenThreshold(lastDebtSync)
	eng.loadPoolLastInterestData(lastDebtSync)
	eng.loadLastDebts()
	eng.loadParameters(lastDebtSync)
	eng.loadLiquidableAccounts(lastDebtSync)
	// process blocks for calculating debts
	adaptersSyncedTill := eng.repo.LoadLastAdapterSync()
	var batchSize int64 = 1000
	for ; lastDebtSync+batchSize < adaptersSyncedTill; lastDebtSync += batchSize {
		eng.processBlocksInBatch(lastDebtSync, lastDebtSync+batchSize)
	}
	eng.processBlocksInBatch(lastDebtSync, adaptersSyncedTill)
}

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
	eng.debts = []*core.Debt{}
	// clear repo after calculating debt as debt uses repository for calculations
	eng.repo.Clear()
}

func (eng *DebtEngine) calCloseAmount(creditManager string, totalValue *core.BigInt, isLiquidated bool, borrowedAmountWithInterest, borrowedAmount *big.Int) (amountToPool, profit, loss *big.Int) {
	params := eng.lastParameters[creditManager]
	loss = big.NewInt(0)
	profit = big.NewInt(0)
	var totalFunds *big.Int
	if isLiquidated {
		totalFunds = utils.PercentMul(totalValue.Convert(), params.LiquidationDiscount.Convert())
	} else {
		totalFunds = totalValue.Convert()
	}
	// borrow amt is greater than total funds
	if totalFunds.Cmp(borrowedAmountWithInterest) < 0 {
		amountToPool = new(big.Int).Sub(totalFunds, big.NewInt(1))
		loss = new(big.Int).Sub(borrowedAmountWithInterest, amountToPool)
	} else {
		if isLiquidated {
			amountToPool = utils.PercentMul(totalFunds, params.FeeLiquidation.Convert())
			amountToPool = new(big.Int).Add(borrowedAmountWithInterest, amountToPool)
		} else {
			interestAmt := new(big.Int).Sub(borrowedAmountWithInterest, borrowedAmount)
			fee := utils.PercentMul(interestAmt, params.FeeInterest.Convert())
			amountToPool = new(big.Int).Add(borrowedAmountWithInterest, fee)
		}

		if totalFunds.Cmp(amountToPool) <= 0 {
			amountToPool = new(big.Int).Sub(totalFunds, big.NewInt(1))
		}
		profit = new(big.Int).Sub(amountToPool, borrowedAmountWithInterest)
	}
	return
}

func (eng *DebtEngine) loadLiquidableAccounts(lastDebtSync int64) {
	data := []*core.LiquidableAccount{}
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
		eng.liquidableBlockTracker[sessionId] = &core.LiquidableAccount{
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
	data := core.DebtSync{}
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
