package core

import (
	"math/big"
	"sort"
)

type (
	Block struct {
		BlockNumber       int64                    `gorm:"primaryKey;column:id"` // Block Number
		Timestamp         uint64                   `gorm:"column:timestamp"`
		AccountOperations []*AccountOperation      `gorm:"foreignKey:block_num"`
		TokenOracles      []*TokenOracle           `gorm:"foreignKey:block_num"`
		PriceFeeds        []*PriceFeed             `gorm:"foreignKey:block_num"`
		Protocols         []*Protocol              `gorm:"foreignKey:block_num"`
		CSS               []*CreditSessionSnapshot `gorm:"foreignKey:block_num"`
		PoolStats         []*PoolStat              `gorm:"foreignKey:block_num"`
		PoolLedgers       []*PoolLedger            `gorm:"foreignKey:block_num"`
		CMStats           []*CreditManagerStat     `gorm:"foreignKey:block_num"`
		AllowedTokens     []*AllowedToken          `gorm:"foreignKey:block_num"`
		DAOOperations     []*DAOOperation          `gorm:"foreignKey:block_num"`
		Params            []*Parameters            `gorm:"foreignKey:block_num"`
		FastCheckParams   []*FastCheckParams       `gorm:"foreignKey:block_num"`
		eventBalances     SortedEventbalances      `gorm:"-"`
		pnlOnCM           map[string]*PnlOnRepay   `gorm:"-"`
		treasuryTransfers []*TreasuryTransfer      `gorm:"foreignKey:block_num"`
		treasurySnapshots []*TreasurySnapshot      `gorm:"foreignKey:timestamp"`
	}
)

func (Block) TableName() string {
	return "blocks"
}

func (b *Block) AddAccountOperation(accountOperation *AccountOperation) {
	b.AccountOperations = append(b.AccountOperations, accountOperation)
}
func (b *Block) AddTokenOracle(tokenOracle *TokenOracle) {
	b.TokenOracles = append(b.TokenOracles, tokenOracle)
}
func (b *Block) AddPriceFeed(pf *PriceFeed) {
	b.PriceFeeds = append(b.PriceFeeds, pf)
}
func (b *Block) AddAllowedProtocol(p *Protocol) {
	b.Protocols = append(b.Protocols, p)
}

func (b *Block) AddAllowedToken(atoken *AllowedToken) {
	b.AllowedTokens = append(b.AllowedTokens, atoken)
}

func (b *Block) AddDAOOperation(operation *DAOOperation) {
	b.DAOOperations = append(b.DAOOperations, operation)
}

func (b *Block) AddCreditSessionSnapshot(css *CreditSessionSnapshot) {
	b.CSS = append(b.CSS, css)
}

func (b *Block) AddPoolStat(ps *PoolStat) {
	b.PoolStats = append(b.PoolStats, ps)
}
func (b *Block) AddPoolLedger(pl *PoolLedger) {
	b.PoolLedgers = append(b.PoolLedgers, pl)
}

func (b *Block) AddCreditManagerStats(cms *CreditManagerStat) {
	b.CMStats = append(b.CMStats, cms)
}

func (b *Block) AddTreasuryTransfer(tt *TreasuryTransfer) {
	b.treasuryTransfers = append(b.treasuryTransfers, tt)
}

func (b *Block) AddTreasurySnapshot(tss *TreasurySnapshot) {
	b.treasurySnapshots = append(b.treasurySnapshots, tss)
}

func (b *Block) GetAllowedTokens() []*AllowedToken {
	return b.AllowedTokens
}

func (b *Block) GetPriceFeeds() []*PriceFeed {
	return b.PriceFeeds
}

func (b *Block) AddEventBalance(eb *EventBalance) {
	b.eventBalances = append(b.eventBalances, eb)
}

func (b *Block) GetEventBalances() []*EventBalance {
	sort.Sort(b.eventBalances)
	return b.eventBalances
}

func (b *Block) AddParameters(params *Parameters) {
	b.Params = append(b.Params, params)
}

func (b *Block) AddFastCheckParams(params *FastCheckParams) {
	b.FastCheckParams = append(b.FastCheckParams, params)
}

func (b *Block) GetCSS() []*CreditSessionSnapshot {
	return b.CSS
}

func (b *Block) GetPoolStats() []*PoolStat {
	return b.PoolStats
}

func (b *Block) AddRepayOnCM(cmAddr string, pnl *PnlOnRepay) {
	if b.pnlOnCM == nil {
		b.pnlOnCM = make(map[string]*PnlOnRepay)
	}
	oldPnl := b.pnlOnCM[cmAddr]
	if oldPnl != nil {
		pnl.Profit = new(big.Int).Add(oldPnl.Profit, pnl.Profit)
		pnl.Loss = new(big.Int).Add(oldPnl.Loss, pnl.Loss)
		pnl.BorrowedAmount = new(big.Int).Add(oldPnl.BorrowedAmount, pnl.BorrowedAmount)
	}
	b.pnlOnCM[cmAddr] = pnl
}

func (b *Block) GetRepayOnCM(cmAddr string) *PnlOnRepay {
	if b.pnlOnCM == nil || b.pnlOnCM[cmAddr] == nil {
		return nil
	}
	return b.pnlOnCM[cmAddr]
}

func (b *Block) GetParams() []*Parameters {
	return b.Params
}
