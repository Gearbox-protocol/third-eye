package debts

import (
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/ethclient"
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
	}
}

func (eng *DebtEngine) Init() {
	lastDebtSync := eng.repo.LoadLastDebtSync()
	eng.loadLastCSS(lastDebtSync)
	eng.loadTokenLastPrice(lastDebtSync)
	eng.loadAllowedTokenThreshold(lastDebtSync)
	eng.loadPoolLastInterestData(lastDebtSync)
	eng.loadLastDebts()
	// process blocks for calculating debts
	adaptersSyncedTill := eng.repo.LoadLastAdapterSync()
	var batchSize int64 = 1000
	for ; lastDebtSync+batchSize < adaptersSyncedTill; lastDebtSync += batchSize {
		eng.processBlocksInBatch(lastDebtSync, lastDebtSync+batchSize)
	}
	eng.processBlocksInBatch(lastDebtSync, adaptersSyncedTill)
}

func (eng *DebtEngine) processBlocksInBatch(from, to int64) {
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
