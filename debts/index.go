package debts

import (
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/log"
	"gorm.io/gorm"
)

type DebtEngine struct {
	repo           core.RepositoryI
	db             *gorm.DB
	client         *ethclient.Client
	config         *config.Config
	lastCSS        map[string]*core.CreditSessionSnapshot
	tokenLastPrice map[string]*core.PriceFeed
	//// credit_manager -> token -> liquidity threshold
	allowedTokensThreshold map[string]map[string]*core.BigInt
	poolLastInterestData   map[string]*core.PoolInterestData
	debts                  []*core.Debt
	lastDebts              map[string]*core.Debt
	liquidableBlockTracker map[string]*core.LiquidableAccount
}

func NewDebtEngine(db *gorm.DB, client *ethclient.Client, config *config.Config, repo core.RepositoryI) core.DebtEngineI {
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
	}
}

func (eng *DebtEngine) Init() {
	if eng.config.DisableDebtEngine {
		return
	}
	lastDebtSync := eng.repo.LoadLastDebtSync()
	eng.loadLastCSS(lastDebtSync)
	eng.loadTokenLastPrice(lastDebtSync)
	eng.loadAllowedTokenThreshold(lastDebtSync)
	eng.loadPoolLastInterestData(lastDebtSync)
	eng.loadLastDebts()
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
		eng.CalculateDebtAndClear()
	}
}

func (eng *DebtEngine) CalculateDebtAndClear() {
	if !eng.config.DisableDebtEngine {
		eng.calculateDebt()
	}
	eng.Clear()
}

func (eng *DebtEngine) Clear() {
	eng.debts = []*core.Debt{}
	// clear repo after calculating debt as debt uses repository for calculations
	eng.repo.Clear()
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
